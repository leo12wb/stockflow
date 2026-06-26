package models

type Movimentacao struct {
	Base
	ProdutoID  uint    `gorm:"not null" json:"produto_id"`
	Produto    Produto `gorm:"foreignKey:ProdutoID" json:"produto,omitempty"`
	UsuarioID  uint    `gorm:"not null" json:"usuario_id"`
	Usuario    Usuario `gorm:"foreignKey:UsuarioID" json:"usuario,omitempty"`
	Tipo       string  `gorm:"not null" json:"tipo"`
	Quantidade int     `gorm:"not null" json:"quantidade"`
	Observacao string  `json:"observacao"`
}
