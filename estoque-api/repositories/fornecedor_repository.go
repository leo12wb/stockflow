package repositories

import (
	"github.com/seu-usuario/estoque-api/config"
	"github.com/seu-usuario/estoque-api/models"
)

type FornecedorRepository struct{}

func (r *FornecedorRepository) FindAll() ([]models.Fornecedor, error) {
	var fornecedores []models.Fornecedor
	result := config.DB.Find(&fornecedores)
	return fornecedores, result.Error
}

func (r *FornecedorRepository) FindAllPaginated(offset, limit int) ([]models.Fornecedor, int64, error) {
	var fornecedores []models.Fornecedor
	var total int64
	config.DB.Model(&models.Fornecedor{}).Count(&total)
	result := config.DB.Offset(offset).Limit(limit).Find(&fornecedores)
	return fornecedores, total, result.Error
}

func (r *FornecedorRepository) FindByID(id uint) (*models.Fornecedor, error) {
	var fornecedor models.Fornecedor
	result := config.DB.First(&fornecedor, id)
	return &fornecedor, result.Error
}

func (r *FornecedorRepository) Create(fornecedor *models.Fornecedor) error {
	return config.DB.Create(fornecedor).Error
}

func (r *FornecedorRepository) Update(fornecedor *models.Fornecedor) error {
	return config.DB.Save(fornecedor).Error
}

func (r *FornecedorRepository) Delete(id uint) error {
	return config.DB.Delete(&models.Fornecedor{}, id).Error
}

func (r *FornecedorRepository) ForceDelete(id uint) error {
	return config.DB.Unscoped().Delete(&models.Fornecedor{}, id).Error
}

func (r *FornecedorRepository) Restore(id uint) error {
	return config.DB.Unscoped().Model(&models.Fornecedor{}).Where("id = ?", id).Update("deleted_at", nil).Error
}
