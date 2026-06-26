package services

import (
	"errors"

	"github.com/seu-usuario/estoque-api/dto"
	"github.com/seu-usuario/estoque-api/models"
	"github.com/seu-usuario/estoque-api/repositories"
)

type FornecedorService struct {
	repo *repositories.FornecedorRepository
}

func NewFornecedorService() *FornecedorService {
	return &FornecedorService{repo: &repositories.FornecedorRepository{}}
}

func (s *FornecedorService) FindAll() ([]models.Fornecedor, error) {
	return s.repo.FindAll()
}

func (s *FornecedorService) FindAllPaginated(offset, limit int) ([]models.Fornecedor, int64, error) {
	return s.repo.FindAllPaginated(offset, limit)
}

func (s *FornecedorService) FindByID(id uint) (*models.Fornecedor, error) {
	return s.repo.FindByID(id)
}

func (s *FornecedorService) Create(req dto.FornecedorRequest) (*models.Fornecedor, error) {
	fornecedor := &models.Fornecedor{
		Nome:     req.Nome,
		CNPJ:     req.CNPJ,
		Email:    req.Email,
		Telefone: req.Telefone,
		Endereco: req.Endereco,
	}

	if err := s.repo.Create(fornecedor); err != nil {
		return nil, errors.New("erro ao criar fornecedor, CNPJ já existe")
	}

	return fornecedor, nil
}

func (s *FornecedorService) Update(id uint, req dto.FornecedorRequest) (*models.Fornecedor, error) {
	fornecedor, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("fornecedor não encontrado")
	}

	fornecedor.Nome = req.Nome
	fornecedor.CNPJ = req.CNPJ
	fornecedor.Email = req.Email
	fornecedor.Telefone = req.Telefone
	fornecedor.Endereco = req.Endereco

	if err := s.repo.Update(fornecedor); err != nil {
		return nil, errors.New("erro ao atualizar fornecedor")
	}

	return fornecedor, nil
}

func (s *FornecedorService) Delete(id uint) error {
	if _, err := s.repo.FindByID(id); err != nil {
		return errors.New("fornecedor não encontrado")
	}
	return s.repo.Delete(id)
}

func (s *FornecedorService) ForceDelete(id uint) error {
	return s.repo.ForceDelete(id)
}

func (s *FornecedorService) Restore(id uint) error {
	return s.repo.Restore(id)
}
