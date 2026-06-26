package repositories

import (
	"github.com/seu-usuario/estoque-api/config"
	"github.com/seu-usuario/estoque-api/models"
)

type ProdutoRepository struct{}

func (r *ProdutoRepository) FindAll() ([]models.Produto, error) {
	var produtos []models.Produto
	result := config.DB.Preload("Fornecedor").Find(&produtos)
	return produtos, result.Error
}

func (r *ProdutoRepository) FindAllPaginated(offset, limit int, search string) ([]models.Produto, int64, error) {
	var produtos []models.Produto
	var total int64
	q := config.DB.Model(&models.Produto{})
	if search != "" {
		q = q.Where("nome ILIKE ? OR codigo_barras ILIKE ? OR categoria ILIKE ?",
			"%"+search+"%", "%"+search+"%", "%"+search+"%")
	}
	q.Count(&total)
	result := q.Preload("Fornecedor").Offset(offset).Limit(limit).Find(&produtos)
	return produtos, total, result.Error
}

func (r *ProdutoRepository) FindByID(id uint) (*models.Produto, error) {
	var produto models.Produto
	result := config.DB.Preload("Fornecedor").First(&produto, id)
	return &produto, result.Error
}

func (r *ProdutoRepository) FindByCodigoBarras(codigo string) (*models.Produto, error) {
	var produto models.Produto
	result := config.DB.Preload("Fornecedor").Where("codigo_barras = ?", codigo).First(&produto)
	return &produto, result.Error
}

func (r *ProdutoRepository) Create(produto *models.Produto) error {
	return config.DB.Create(produto).Error
}

func (r *ProdutoRepository) Update(produto *models.Produto) error {
	return config.DB.Save(produto).Error
}

func (r *ProdutoRepository) Delete(id uint) error {
	return config.DB.Delete(&models.Produto{}, id).Error
}

func (r *ProdutoRepository) ForceDelete(id uint) error {
	return config.DB.Unscoped().Delete(&models.Produto{}, id).Error
}

func (r *ProdutoRepository) Restore(id uint) error {
	return config.DB.Unscoped().Model(&models.Produto{}).Where("id = ?", id).Update("deleted_at", nil).Error
}

func (r *ProdutoRepository) FindEstoqueBaixo() ([]models.Produto, error) {
	var produtos []models.Produto
	result := config.DB.Preload("Fornecedor").Where("quantidade <= estoque_minimo").Find(&produtos)
	return produtos, result.Error
}
