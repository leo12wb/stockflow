package models

type Fornecedor struct {
	Base
	Nome     string    `gorm:"not null" json:"nome"`
	CNPJ     string    `gorm:"unique;not null" json:"cnpj"`
	Email    string    `json:"email"`
	Telefone string    `json:"telefone"`
	Endereco string    `json:"endereco"`
	Produtos []Produto `gorm:"foreignKey:FornecedorID" json:"produtos,omitempty"`
}
