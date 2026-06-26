package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seu-usuario/estoque-api/dto"
	"github.com/seu-usuario/estoque-api/services"
	"github.com/seu-usuario/estoque-api/utils"
)

type AuthController struct {
	service *services.AuthService
}

func NewAuthController() *AuthController {
	return &AuthController{service: services.NewAuthService()}
}

// Login godoc
// @Summary      Login
// @Description  Autenticação de usuário, retorna token JWT
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body dto.LoginRequest true "Credenciais de acesso"
// @Success      200 {object} utils.Response
// @Failure      401 {object} utils.Response
// @Failure      422 {object} utils.ValidationResponse
// @Router       /auth/login [post]
func (ctrl *AuthController) Login(c *gin.Context) {
	var req dto.LoginRequest
	if !utils.Validate(c, &req, utils.Rules{
		"email": "required|email",
		"senha": "required|min:6",
	}) {
		return
	}

	resp, err := ctrl.service.Login(req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "login realizado com sucesso", resp)
}

// Register godoc
// @Summary      Registrar usuário
// @Description  Cria um novo usuário no sistema
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body dto.RegisterRequest true "Dados do usuário"
// @Success      201 {object} utils.Response
// @Failure      409 {object} utils.Response
// @Failure      422 {object} utils.ValidationResponse
// @Router       /auth/register [post]
func (ctrl *AuthController) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if !utils.Validate(c, &req, utils.Rules{
		"nome":   "required",
		"email":  "required|email|unique:usuarios,email",
		"senha":  "required|min:6",
		"perfil": "required|in:administrador,gerente,estoquista",
	}) {
		return
	}

	usuario, err := ctrl.service.Register(req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusConflict, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "usuário criado com sucesso", usuario)
}

// Me godoc
// @Summary      Usuário autenticado
// @Description  Retorna os dados do usuário logado
// @Tags         Auth
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} utils.Response
// @Failure      401 {object} utils.Response
// @Router       /auth/me [get]
func (ctrl *AuthController) Me(c *gin.Context) {
	userID, _ := c.Get("user_id")
	email, _ := c.Get("user_email")
	perfil, _ := c.Get("user_perfil")

	utils.SuccessResponse(c, http.StatusOK, "", gin.H{
		"id":     userID,
		"email":  email,
		"perfil": perfil,
	})
}
