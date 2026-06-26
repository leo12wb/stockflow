package services

import (
	"errors"

	"github.com/seu-usuario/estoque-api/config"
	"github.com/seu-usuario/estoque-api/dto"
	"github.com/seu-usuario/estoque-api/models"
	"github.com/seu-usuario/estoque-api/repositories"
)

type MovimentacaoService struct {
	repo        *repositories.MovimentacaoRepository
	produtoRepo *repositories.ProdutoRepository
}

func NewMovimentacaoService() *MovimentacaoService {
	return &MovimentacaoService{
		repo:        &repositories.MovimentacaoRepository{},
		produtoRepo: &repositories.ProdutoRepository{},
	}
}

func (s *MovimentacaoService) FindAll() ([]models.Movimentacao, error) {
	return s.repo.FindAll()
}

func (s *MovimentacaoService) FindAllPaginated(offset, limit int) ([]models.Movimentacao, int64, error) {
	return s.repo.FindAllPaginated(offset, limit)
}

func (s *MovimentacaoService) Create(req dto.MovimentacaoRequest, usuarioID uint) (*models.Movimentacao, error) {
	produto, err := s.produtoRepo.FindByID(req.ProdutoID)
	if err != nil {
		return nil, errors.New("produto não encontrado")
	}

	tx := config.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	switch req.Tipo {
	case "ENTRADA":
		produto.Quantidade += req.Quantidade
	case "SAIDA":
		if produto.Quantidade < req.Quantidade {
			tx.Rollback()
			return nil, errors.New("estoque insuficiente")
		}
		produto.Quantidade -= req.Quantidade
	case "AJUSTE":
		produto.Quantidade = req.Quantidade
	default:
		tx.Rollback()
		return nil, errors.New("tipo de movimentação inválido")
	}

	if err := tx.Save(produto).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("erro ao atualizar estoque")
	}

	mov := &models.Movimentacao{
		ProdutoID:  req.ProdutoID,
		UsuarioID:  usuarioID,
		Tipo:       req.Tipo,
		Quantidade: req.Quantidade,
		Observacao: req.Observacao,
	}

	if err := tx.Create(mov).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("erro ao registrar movimentação")
	}

	tx.Commit()
	return mov, nil
}
