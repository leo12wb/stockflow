package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seu-usuario/estoque-api/dto"
	"github.com/seu-usuario/estoque-api/services"
	"github.com/seu-usuario/estoque-api/utils"
)

type MovimentacaoController struct {
	service *services.MovimentacaoService
}

func NewMovimentacaoController() *MovimentacaoController {
	return &MovimentacaoController{service: services.NewMovimentacaoService()}
}

// FindAll godoc
// @Summary      Listar movimentações
// @Description  Retorna todas as movimentações de estoque
// @Tags         Movimentações
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} utils.Response
// @Router       /movimentacoes [get]
func (ctrl *MovimentacaoController) FindAll(c *gin.Context) {
	p := utils.ParsePagination(c)
	movimentacoes, total, err := ctrl.service.FindAllPaginated(p.Offset, p.PerPage)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "erro ao buscar movimentações")
		return
	}
	utils.PagedSuccessResponse(c, movimentacoes, total, p)
}

// Create godoc
// @Summary      Registrar movimentação
// @Description  Registra uma entrada, saída ou ajuste de estoque
// @Tags         Movimentações
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body dto.MovimentacaoRequest true "Dados da movimentação"
// @Success      201 {object} utils.Response
// @Failure      400 {object} utils.Response
// @Failure      422 {object} utils.ValidationResponse
// @Router       /movimentacoes [post]
func (ctrl *MovimentacaoController) Create(c *gin.Context) {
	var req dto.MovimentacaoRequest
	if !utils.Validate(c, &req, utils.Rules{
		"produto_id": "required|numeric|exists:produtos,id",
		"tipo":       "required|in:ENTRADA,SAIDA,AJUSTE",
		"quantidade": "required|numeric|min:1",
		"observacao": "nullable",
	}) {
		return
	}

	userID, _ := c.Get("user_id")
	mov, err := ctrl.service.Create(req, userID.(uint))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "movimentação registrada com sucesso", mov)
}
