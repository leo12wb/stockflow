package dto

type FornecedorRequest struct {
	Nome     string `json:"nome" validate:"required"`
	CNPJ     string `json:"cnpj" validate:"required"`
	Email    string `json:"email" validate:"omitempty,email"`
	Telefone string `json:"telefone"`
	Endereco string `json:"endereco"`
}
