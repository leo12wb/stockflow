package dto

type MovimentacaoRequest struct {
	ProdutoID  uint   `json:"produto_id" validate:"required"`
	Tipo       string `json:"tipo" validate:"required,oneof=ENTRADA SAIDA AJUSTE"`
	Quantidade int    `json:"quantidade" validate:"required,gt=0"`
	Observacao string `json:"observacao"`
}
