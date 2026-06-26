package services

import (
	"errors"

	"github.com/seu-usuario/estoque-api/dto"
	"github.com/seu-usuario/estoque-api/models"
	"github.com/seu-usuario/estoque-api/repositories"
	"github.com/seu-usuario/estoque-api/utils"
)

type ProdutoService struct {
	repo *repositories.ProdutoRepository
}

func NewProdutoService() *ProdutoService {
	return &ProdutoService{repo: &repositories.ProdutoRepository{}}
}

func (s *ProdutoService) FindAll() ([]models.Produto, error) {
	return s.repo.FindAll()
}

func (s *ProdutoService) FindAllPaginated(offset, limit int, search string) ([]models.Produto, int64, error) {
	return s.repo.FindAllPaginated(offset, limit, search)
}

func (s *ProdutoService) FindByID(id uint) (*models.Produto, error) {
	return s.repo.FindByID(id)
}

func (s *ProdutoService) FindByCodigoBarras(codigo string) (*models.Produto, error) {
	return s.repo.FindByCodigoBarras(codigo)
}

func (s *ProdutoService) Create(req dto.ProdutoRequest) (*models.Produto, error) {
	codigoBarras := req.CodigoBarras
	if codigoBarras == "" {
		codigoBarras = utils.GenerateBarcode()
	}

	produto := &models.Produto{
		Nome:          req.Nome,
		Descricao:     req.Descricao,
		CodigoBarras:  codigoBarras,
		Categoria:     req.Categoria,
		PrecoCompra:   req.PrecoCompra,
		PrecoVenda:    req.PrecoVenda,
		Quantidade:    req.Quantidade,
		EstoqueMinimo: req.EstoqueMinimo,
		FornecedorID:  req.FornecedorID,
	}

	if err := s.repo.Create(produto); err != nil {
		return nil, errors.New("erro ao criar produto")
	}

	return produto, nil
}

func (s *ProdutoService) Update(id uint, req dto.ProdutoRequest) (*models.Produto, error) {
	produto, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("produto não encontrado")
	}

	produto.Nome = req.Nome
	produto.Descricao = req.Descricao
	produto.Categoria = req.Categoria
	produto.PrecoCompra = req.PrecoCompra
	produto.PrecoVenda = req.PrecoVenda
	produto.Quantidade = req.Quantidade
	produto.EstoqueMinimo = req.EstoqueMinimo
	produto.FornecedorID = req.FornecedorID

	if req.CodigoBarras != "" {
		produto.CodigoBarras = req.CodigoBarras
	}

	if err := s.repo.Update(produto); err != nil {
		return nil, errors.New("erro ao atualizar produto")
	}

	return produto, nil
}

func (s *ProdutoService) Delete(id uint) error {
	if _, err := s.repo.FindByID(id); err != nil {
		return errors.New("produto não encontrado")
	}
	return s.repo.Delete(id)
}

func (s *ProdutoService) ForceDelete(id uint) error {
	return s.repo.ForceDelete(id)
}

func (s *ProdutoService) Restore(id uint) error {
	return s.repo.Restore(id)
}

func (s *ProdutoService) FindEstoqueBaixo() ([]models.Produto, error) {
	return s.repo.FindEstoqueBaixo()
}
