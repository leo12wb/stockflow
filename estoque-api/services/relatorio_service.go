package services

import (
	"github.com/seu-usuario/estoque-api/models"
	"github.com/seu-usuario/estoque-api/repositories"
)

type RelatorioService struct {
	produtoRepo      *repositories.ProdutoRepository
	movimentacaoRepo *repositories.MovimentacaoRepository
}

func NewRelatorioService() *RelatorioService {
	return &RelatorioService{
		produtoRepo:      &repositories.ProdutoRepository{},
		movimentacaoRepo: &repositories.MovimentacaoRepository{},
	}
}

func (s *RelatorioService) Inventario() ([]models.Produto, error) {
	return s.produtoRepo.FindAll()
}

func (s *RelatorioService) EstoqueBaixo() ([]models.Produto, error) {
	return s.produtoRepo.FindEstoqueBaixo()
}

func (s *RelatorioService) Entradas() ([]models.Movimentacao, error) {
	return s.movimentacaoRepo.FindByTipo("ENTRADA")
}

func (s *RelatorioService) Saidas() ([]models.Movimentacao, error) {
	return s.movimentacaoRepo.FindByTipo("SAIDA")
}

func (s *RelatorioService) MaisVendidos() ([]models.Produto, error) {
	return s.movimentacaoRepo.FindMaisVendidos()
}

func (s *RelatorioService) Movimentacoes() ([]models.Movimentacao, error) {
	return s.movimentacaoRepo.FindAll()
}
