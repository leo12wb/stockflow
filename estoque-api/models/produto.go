package models

type Produto struct {
	Base
	Nome          string     `gorm:"not null" json:"nome"`
	Descricao     string     `json:"descricao"`
	CodigoBarras  string     `gorm:"unique" json:"codigo_barras"`
	Categoria     string     `json:"categoria"`
	PrecoCompra   float64    `gorm:"not null;default:0" json:"preco_compra"`
	PrecoVenda    float64    `gorm:"not null;default:0" json:"preco_venda"`
	Quantidade    int        `gorm:"not null;default:0" json:"quantidade"`
	EstoqueMinimo int        `gorm:"not null;default:0" json:"estoque_minimo"`
	FornecedorID  uint       `json:"fornecedor_id"`
	Fornecedor    Fornecedor `gorm:"foreignKey:FornecedorID" json:"fornecedor,omitempty"`
}
