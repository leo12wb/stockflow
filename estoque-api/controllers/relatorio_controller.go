package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seu-usuario/estoque-api/services"
	"github.com/seu-usuario/estoque-api/utils"
)

type RelatorioController struct {
	service *services.RelatorioService
}

func NewRelatorioController() *RelatorioController {
	return &RelatorioController{service: services.NewRelatorioService()}
}

// Inventario godoc
// @Summary      Inventário geral
// @Description  Retorna todos os produtos com quantidades atuais
// @Tags         Relatórios
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} utils.Response
// @Router       /relatorios/inventario [get]
func (ctrl *RelatorioController) Inventario(c *gin.Context) {
	data, err := ctrl.service.Inventario()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "erro ao gerar inventário")
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "inventário", data)
}

// EstoqueBaixo godoc
// @Summary      Estoque abaixo do mínimo
// @Description  Retorna produtos com quantidade abaixo do estoque mínimo
// @Tags         Relatórios
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} utils.Response
// @Router       /relatorios/estoque-baixo [get]
func (ctrl *RelatorioController) EstoqueBaixo(c *gin.Context) {
	data, err := ctrl.service.EstoqueBaixo()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "erro ao buscar estoque baixo")
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "estoque abaixo do mínimo", data)
}

// Entradas godoc
// @Summary      Relatório de entradas
// @Description  Retorna todas as movimentações do tipo ENTRADA
// @Tags         Relatórios
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} utils.Response
// @Router       /relatorios/entradas [get]
func (ctrl *RelatorioController) Entradas(c *gin.Context) {
	data, err := ctrl.service.Entradas()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "erro ao buscar entradas")
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "relatório de entradas", data)
}

// Saidas godoc
// @Summary      Relatório de saídas
// @Description  Retorna todas as movimentações do tipo SAIDA
// @Tags         Relatórios
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} utils.Response
// @Router       /relatorios/saidas [get]
func (ctrl *RelatorioController) Saidas(c *gin.Context) {
	data, err := ctrl.service.Saidas()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "erro ao buscar saídas")
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "relatório de saídas", data)
}

// MaisVendidos godoc
// @Summary      Produtos mais vendidos
// @Description  Retorna os 10 produtos com maior volume de saídas
// @Tags         Relatórios
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} utils.Response
// @Router       /relatorios/mais-vendidos [get]
func (ctrl *RelatorioController) MaisVendidos(c *gin.Context) {
	data, err := ctrl.service.MaisVendidos()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "erro ao buscar mais vendidos")
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "produtos mais vendidos", data)
}

// Movimentacoes godoc
// @Summary      Relatório de movimentações
// @Description  Retorna todas as movimentações de estoque
// @Tags         Relatórios
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} utils.Response
// @Router       /relatorios/movimentacoes [get]
func (ctrl *RelatorioController) Movimentacoes(c *gin.Context) {
	data, err := ctrl.service.Movimentacoes()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "erro ao buscar movimentações")
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "relatório de movimentações", data)
}
