package repositories

import (
	"time"

	"github.com/seu-usuario/estoque-api/config"
	"github.com/seu-usuario/estoque-api/models"
)

type MovimentacaoRepository struct{}

func (r *MovimentacaoRepository) FindAll() ([]models.Movimentacao, error) {
	var movimentacoes []models.Movimentacao
	result := config.DB.Preload("Produto").Preload("Usuario").Order("created_at desc").Find(&movimentacoes)
	return movimentacoes, result.Error
}

func (r *MovimentacaoRepository) FindAllPaginated(offset, limit int) ([]models.Movimentacao, int64, error) {
	var movimentacoes []models.Movimentacao
	var total int64
	config.DB.Model(&models.Movimentacao{}).Count(&total)
	result := config.DB.Preload("Produto").Preload("Usuario").Order("created_at desc").Offset(offset).Limit(limit).Find(&movimentacoes)
	return movimentacoes, total, result.Error
}

func (r *MovimentacaoRepository) Create(mov *models.Movimentacao) error {
	return config.DB.Create(mov).Error
}

func (r *MovimentacaoRepository) FindByTipo(tipo string) ([]models.Movimentacao, error) {
	var movimentacoes []models.Movimentacao
	result := config.DB.Preload("Produto").Preload("Usuario").Where("tipo = ?", tipo).Order("created_at desc").Find(&movimentacoes)
	return movimentacoes, result.Error
}

func (r *MovimentacaoRepository) FindMaisVendidos() ([]models.Produto, error) {
	var produtos []models.Produto
	result := config.DB.Raw(`
		SELECT p.* FROM produtos p
		INNER JOIN movimentacoes m ON m.produto_id = p.id
		WHERE m.tipo = 'SAIDA' AND m.deleted_at IS NULL AND p.deleted_at IS NULL
		GROUP BY p.id
		ORDER BY SUM(m.quantidade) DESC
		LIMIT 10
	`).Scan(&produtos)
	return produtos, result.Error
}

func (r *MovimentacaoRepository) FindByPeriodo(inicio, fim time.Time) ([]models.Movimentacao, error) {
	var movimentacoes []models.Movimentacao
	result := config.DB.Preload("Produto").Preload("Usuario").
		Where("created_at BETWEEN ? AND ?", inicio, fim).
		Order("created_at desc").Find(&movimentacoes)
	return movimentacoes, result.Error
}
