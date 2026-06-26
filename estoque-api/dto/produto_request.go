package dto

type ProdutoRequest struct {
	Nome          string  `json:"nome" validate:"required"`
	Descricao     string  `json:"descricao"`
	CodigoBarras  string  `json:"codigo_barras"`
	Categoria     string  `json:"categoria"`
	PrecoCompra   float64 `json:"preco_compra" validate:"gte=0"`
	PrecoVenda    float64 `json:"preco_venda" validate:"gte=0"`
	Quantidade    int     `json:"quantidade" validate:"gte=0"`
	EstoqueMinimo int     `json:"estoque_minimo" validate:"gte=0"`
	FornecedorID  uint    `json:"fornecedor_id"`
}
