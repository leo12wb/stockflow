package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/seu-usuario/estoque-api/config"
)

type Rules map[string]string

type ValidationResponse struct {
	Success bool                `json:"success"`
	Message string              `json:"message"`
	Errors  map[string][]string `json:"errors"`
}

var emailValidator = validator.New()

// Validate lê o body, popula a struct e valida as regras inline — igual Laravel
func Validate(c *gin.Context, obj interface{}, rules Rules) bool {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, "erro ao ler requisição")
		return false
	}

	// popula a struct com os dados do body
	if len(body) > 0 {
		if err := json.Unmarshal(body, obj); err != nil {
			ErrorResponse(c, http.StatusBadRequest, "JSON inválido: "+err.Error())
			return false
		}
	}

	// mapa para aplicar as regras campo a campo
	data := make(map[string]interface{})
	if len(body) > 0 {
		json.Unmarshal(body, &data)
	}

	errors := make(map[string][]string)

	for field, ruleStr := range rules {
		value, exists := data[field]
		ruleList := strings.Split(ruleStr, "|")

		isNullable := false
		for _, r := range ruleList {
			if r == "nullable" {
				isNullable = true
				break
			}
		}

		// nullable e vazio → ignora as demais regras
		if isNullable && (!exists || value == nil || fmt.Sprintf("%v", value) == "") {
			continue
		}

		for _, rule := range ruleList {
			if rule == "nullable" {
				continue
			}

			parts := strings.SplitN(rule, ":", 2)
			ruleName := parts[0]
			param := ""
			if len(parts) > 1 {
				param = parts[1]
			}

			msg := applyRule(field, value, exists, ruleName, param)
			if msg != "" {
				errors[field] = append(errors[field], msg)
				break // para no primeiro erro do campo, igual Laravel
			}
		}
	}

	if len(errors) > 0 {
		c.JSON(http.StatusUnprocessableEntity, ValidationResponse{
			Success: false,
			Message: "Os dados informados são inválidos.",
			Errors:  errors,
		})
		return false
	}

	return true
}

func applyRule(field string, value interface{}, exists bool, rule, param string) string {
	strVal := ""
	if value != nil {
		strVal = fmt.Sprintf("%v", value)
	}

	switch rule {
	case "required":
		if !exists || value == nil || strVal == "" {
			return fmt.Sprintf("O campo %s é obrigatório.", field)
		}

	case "string":
		if exists && value != nil {
			if _, ok := value.(string); !ok {
				return fmt.Sprintf("O campo %s deve ser um texto.", field)
			}
		}

	case "numeric":
		if exists && value != nil && strVal != "" {
			if _, err := strconv.ParseFloat(strVal, 64); err != nil {
				return fmt.Sprintf("O campo %s deve ser um número.", field)
			}
		}

	case "boolean":
		if exists && value != nil {
			if _, ok := value.(bool); !ok {
				return fmt.Sprintf("O campo %s deve ser verdadeiro ou falso.", field)
			}
		}

	case "email":
		if exists && strVal != "" {
			if err := emailValidator.Var(strVal, "email"); err != nil {
				return fmt.Sprintf("O campo %s deve ser um e-mail válido.", field)
			}
		}

	case "min":
		if exists && strVal != "" {
			n, _ := strconv.Atoi(param)
			if _, err := strconv.ParseFloat(strVal, 64); err == nil {
				// numérico
				f, _ := strconv.ParseFloat(strVal, 64)
				if f < float64(n) {
					return fmt.Sprintf("O campo %s deve ser no mínimo %s.", field, param)
				}
			} else {
				// string
				if len(strVal) < n {
					return fmt.Sprintf("O campo %s deve ter no mínimo %s caracteres.", field, param)
				}
			}
		}

	case "max":
		if exists && strVal != "" {
			n, _ := strconv.Atoi(param)
			if _, err := strconv.ParseFloat(strVal, 64); err == nil {
				f, _ := strconv.ParseFloat(strVal, 64)
				if f > float64(n) {
					return fmt.Sprintf("O campo %s deve ser no máximo %s.", field, param)
				}
			} else {
				if len(strVal) > n {
					return fmt.Sprintf("O campo %s deve ter no máximo %s caracteres.", field, param)
				}
			}
		}

	case "in":
		if exists && strVal != "" {
			options := strings.Split(param, ",")
			for _, opt := range options {
				if strVal == strings.TrimSpace(opt) {
					return ""
				}
			}
			return fmt.Sprintf("O campo %s deve ser um dos valores: %s.", field, strings.Join(strings.Split(param, ","), ", "))
		}

	case "unique":
		// unique:tabela ou unique:tabela,coluna
		if exists && strVal != "" {
			parts := strings.SplitN(param, ",", 2)
			table := parts[0]
			column := field
			if len(parts) > 1 {
				column = parts[1]
			}
			var count int64
			config.DB.Table(table).Where(column+" = ?", strVal).Count(&count)
			if count > 0 {
				return fmt.Sprintf("O campo %s já está em uso.", field)
			}
		}

	case "exists":
		// exists:tabela,coluna
		if exists && strVal != "" {
			parts := strings.SplitN(param, ",", 2)
			table := parts[0]
			column := field
			if len(parts) > 1 {
				column = parts[1]
			}
			var count int64
			config.DB.Table(table).Where(column+" = ?", strVal).Count(&count)
			if count == 0 {
				return fmt.Sprintf("O valor informado em %s não existe.", field)
			}
		}
	}

	return ""
}
