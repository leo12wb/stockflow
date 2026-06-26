package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/seu-usuario/estoque-api/dto"
	"github.com/seu-usuario/estoque-api/services"
	"github.com/seu-usuario/estoque-api/utils"
)

type ProdutoController struct {
	service *services.ProdutoService
}

func NewProdutoController() *ProdutoController {
	return &ProdutoController{service: services.NewProdutoService()}
}

// FindAll godoc
// @Summary      Listar produtos
// @Description  Retorna todos os produtos ativos
// @Tags         Produtos
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} utils.Response
// @Failure      401 {object} utils.Response
// @Router       /produtos [get]
func (ctrl *ProdutoController) FindAll(c *gin.Context) {
	p := utils.ParsePagination(c)
	search := c.Query("search")
	produtos, total, err := ctrl.service.FindAllPaginated(p.Offset, p.PerPage, search)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "erro ao buscar produtos")
		return
	}
	utils.PagedSuccessResponse(c, produtos, total, p)
}

// FindByID godoc
// @Summary      Buscar produto por ID
// @Description  Retorna um produto pelo ID
// @Tags         Produtos
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "ID do produto"
// @Success      200 {object} utils.Response
// @Failure      404 {object} utils.Response
// @Router       /produtos/{id} [get]
func (ctrl *ProdutoController) FindByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "id inválido")
		return
	}

	produto, err := ctrl.service.FindByID(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "produto não encontrado")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "", produto)
}

// FindByCodigoBarras godoc
// @Summary      Buscar produto por código de barras
// @Description  Retorna um produto pelo código de barras
// @Tags         Produtos
// @Produce      json
// @Security     BearerAuth
// @Param        codigo path string true "Código de barras"
// @Success      200 {object} utils.Response
// @Failure      404 {object} utils.Response
// @Router       /produtos/codigo/{codigo} [get]
func (ctrl *ProdutoController) FindByCodigoBarras(c *gin.Context) {
	codigo := c.Param("codigo")
	produto, err := ctrl.service.FindByCodigoBarras(codigo)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "produto não encontrado")
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "", produto)
}

// Create godoc
// @Summary      Criar produto
// @Description  Cadastra um novo produto
// @Tags         Produtos
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body dto.ProdutoRequest true "Dados do produto"
// @Success      201 {object} utils.Response
// @Failure      422 {object} utils.ValidationResponse
// @Router       /produtos [post]
func (ctrl *ProdutoController) Create(c *gin.Context) {
	var req dto.ProdutoRequest
	if !utils.Validate(c, &req, utils.Rules{
		"nome":           "required",
		"codigo_barras":  "nullable|unique:produtos,codigo_barras",
		"categoria":      "nullable",
		"descricao":      "nullable",
		"preco_compra":   "required|numeric",
		"preco_venda":    "required|numeric",
		"quantidade":     "required|numeric",
		"estoque_minimo": "nullable|numeric",
		"fornecedor_id":  "nullable|numeric|exists:fornecedores,id",
	}) {
		return
	}

	produto, err := ctrl.service.Create(req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "produto criado com sucesso", produto)
}

// Update godoc
// @Summary      Atualizar produto
// @Description  Atualiza os dados de um produto
// @Tags         Produtos
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "ID do produto"
// @Param        request body dto.ProdutoRequest true "Dados do produto"
// @Success      200 {object} utils.Response
// @Failure      422 {object} utils.ValidationResponse
// @Router       /produtos/{id} [put]
func (ctrl *ProdutoController) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "id inválido")
		return
	}

	var req dto.ProdutoRequest
	if !utils.Validate(c, &req, utils.Rules{
		"nome":           "required",
		"categoria":      "nullable",
		"descricao":      "nullable",
		"preco_compra":   "required|numeric",
		"preco_venda":    "required|numeric",
		"quantidade":     "required|numeric",
		"estoque_minimo": "nullable|numeric",
		"fornecedor_id":  "nullable|numeric|exists:fornecedores,id",
	}) {
		return
	}

	produto, err := ctrl.service.Update(uint(id), req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "produto atualizado com sucesso", produto)
}

// Delete godoc
// @Summary      Remover produto (soft delete)
// @Description  Remove um produto com soft delete
// @Tags         Produtos
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "ID do produto"
// @Success      200 {object} utils.Response
// @Failure      404 {object} utils.Response
// @Router       /produtos/{id} [delete]
func (ctrl *ProdutoController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "id inválido")
		return
	}

	if err := ctrl.service.Delete(uint(id)); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "produto removido com sucesso", nil)
}

// ForceDelete godoc
// @Summary      Excluir produto permanentemente
// @Description  Exclui um produto de forma permanente (irreversível)
// @Tags         Produtos
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "ID do produto"
// @Success      200 {object} utils.Response
// @Router       /produtos/{id}/force [delete]
func (ctrl *ProdutoController) ForceDelete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "id inválido")
		return
	}

	if err := ctrl.service.ForceDelete(uint(id)); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "produto excluído permanentemente", nil)
}

// Restore godoc
// @Summary      Restaurar produto
// @Description  Restaura um produto removido via soft delete
// @Tags         Produtos
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "ID do produto"
// @Success      200 {object} utils.Response
// @Router       /produtos/{id}/restaurar [patch]
func (ctrl *ProdutoController) Restore(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "id inválido")
		return
	}

	if err := ctrl.service.Restore(uint(id)); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "produto restaurado com sucesso", nil)
}
