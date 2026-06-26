package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/seu-usuario/estoque-api/dto"
	"github.com/seu-usuario/estoque-api/services"
	"github.com/seu-usuario/estoque-api/utils"
)

type FornecedorController struct {
	service *services.FornecedorService
}

func NewFornecedorController() *FornecedorController {
	return &FornecedorController{service: services.NewFornecedorService()}
}

// FindAll godoc
// @Summary      Listar fornecedores
// @Description  Retorna todos os fornecedores ativos
// @Tags         Fornecedores
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} utils.Response
// @Router       /fornecedores [get]
func (ctrl *FornecedorController) FindAll(c *gin.Context) {
	p := utils.ParsePagination(c)
	fornecedores, total, err := ctrl.service.FindAllPaginated(p.Offset, p.PerPage)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "erro ao buscar fornecedores")
		return
	}
	utils.PagedSuccessResponse(c, fornecedores, total, p)
}

// FindByID godoc
// @Summary      Buscar fornecedor por ID
// @Description  Retorna um fornecedor pelo ID
// @Tags         Fornecedores
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "ID do fornecedor"
// @Success      200 {object} utils.Response
// @Failure      404 {object} utils.Response
// @Router       /fornecedores/{id} [get]
func (ctrl *FornecedorController) FindByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "id inválido")
		return
	}

	fornecedor, err := ctrl.service.FindByID(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "fornecedor não encontrado")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "", fornecedor)
}

// Create godoc
// @Summary      Criar fornecedor
// @Description  Cadastra um novo fornecedor
// @Tags         Fornecedores
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body dto.FornecedorRequest true "Dados do fornecedor"
// @Success      201 {object} utils.Response
// @Failure      422 {object} utils.ValidationResponse
// @Router       /fornecedores [post]
func (ctrl *FornecedorController) Create(c *gin.Context) {
	var req dto.FornecedorRequest
	if !utils.Validate(c, &req, utils.Rules{
		"nome":     "required",
		"cnpj":     "required|unique:fornecedores,cnpj",
		"email":    "nullable|email",
		"telefone": "nullable",
		"endereco": "nullable",
	}) {
		return
	}

	fornecedor, err := ctrl.service.Create(req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusConflict, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "fornecedor criado com sucesso", fornecedor)
}

// Update godoc
// @Summary      Atualizar fornecedor
// @Description  Atualiza os dados de um fornecedor
// @Tags         Fornecedores
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "ID do fornecedor"
// @Param        request body dto.FornecedorRequest true "Dados do fornecedor"
// @Success      200 {object} utils.Response
// @Failure      422 {object} utils.ValidationResponse
// @Router       /fornecedores/{id} [put]
func (ctrl *FornecedorController) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "id inválido")
		return
	}

	var req dto.FornecedorRequest
	if !utils.Validate(c, &req, utils.Rules{
		"nome":     "required",
		"cnpj":     "required",
		"email":    "nullable|email",
		"telefone": "nullable",
		"endereco": "nullable",
	}) {
		return
	}

	fornecedor, err := ctrl.service.Update(uint(id), req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "fornecedor atualizado com sucesso", fornecedor)
}

// Delete godoc
// @Summary      Remover fornecedor (soft delete)
// @Description  Remove um fornecedor com soft delete
// @Tags         Fornecedores
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "ID do fornecedor"
// @Success      200 {object} utils.Response
// @Router       /fornecedores/{id} [delete]
func (ctrl *FornecedorController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "id inválido")
		return
	}

	if err := ctrl.service.Delete(uint(id)); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "fornecedor removido com sucesso", nil)
}

// ForceDelete godoc
// @Summary      Excluir fornecedor permanentemente
// @Description  Exclui um fornecedor de forma permanente
// @Tags         Fornecedores
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "ID do fornecedor"
// @Success      200 {object} utils.Response
// @Router       /fornecedores/{id}/force [delete]
func (ctrl *FornecedorController) ForceDelete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "id inválido")
		return
	}

	if err := ctrl.service.ForceDelete(uint(id)); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "fornecedor excluído permanentemente", nil)
}

// Restore godoc
// @Summary      Restaurar fornecedor
// @Description  Restaura um fornecedor removido via soft delete
// @Tags         Fornecedores
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "ID do fornecedor"
// @Success      200 {object} utils.Response
// @Router       /fornecedores/{id}/restaurar [patch]
func (ctrl *FornecedorController) Restore(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "id inválido")
		return
	}

	if err := ctrl.service.Restore(uint(id)); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "fornecedor restaurado com sucesso", nil)
}
