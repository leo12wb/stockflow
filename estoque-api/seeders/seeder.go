package seeders

import (
	"log"

	"github.com/seu-usuario/estoque-api/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	seedUsuarios(db)
	seedFornecedores(db)
	seedProdutos(db)
	seedMovimentacoes(db)
	log.Println("Seeders executados com sucesso!")
}

func seedUsuarios(db *gorm.DB) {
	var count int64
	db.Model(&models.Usuario{}).Count(&count)
	if count > 0 {
		return
	}

	senha, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)

	usuarios := []models.Usuario{
		{Nome: "Administrador", Email: "admin@estoque.com", Senha: string(senha), Perfil: "administrador"},
		{Nome: "Gerente Silva", Email: "gerente@estoque.com", Senha: string(senha), Perfil: "gerente"},
		{Nome: "Estoquista João", Email: "estoquista@estoque.com", Senha: string(senha), Perfil: "estoquista"},
	}

	db.Create(&usuarios)
	log.Println("Usuários criados.")
}

func seedFornecedores(db *gorm.DB) {
	var count int64
	db.Model(&models.Fornecedor{}).Count(&count)
	if count > 0 {
		return
	}

	fornecedores := []models.Fornecedor{
		{
			Nome:     "Distribuidora Alpha Ltda",
			CNPJ:     "12.345.678/0001-90",
			Email:    "contato@alpha.com.br",
			Telefone: "(11) 98765-4321",
			Endereco: "Rua das Indústrias, 100 - São Paulo/SP",
		},
		{
			Nome:     "Atacado Beta S/A",
			CNPJ:     "98.765.432/0001-10",
			Email:    "vendas@beta.com.br",
			Telefone: "(21) 91234-5678",
			Endereco: "Av. Comercial, 500 - Rio de Janeiro/RJ",
		},
		{
			Nome:     "Fornecedor Gamma ME",
			CNPJ:     "11.222.333/0001-44",
			Email:    "gamma@fornecedor.com",
			Telefone: "(31) 99876-5432",
			Endereco: "Rua do Comércio, 200 - Belo Horizonte/MG",
		},
	}

	db.Create(&fornecedores)
	log.Println("Fornecedores criados.")
}

func seedProdutos(db *gorm.DB) {
	var count int64
	db.Model(&models.Produto{}).Count(&count)
	if count > 0 {
		return
	}

	produtos := []models.Produto{
		{
			Nome:          "Notebook Dell Inspiron 15",
			Descricao:     "Notebook 15.6\" Intel Core i5, 8GB RAM, 256GB SSD",
			CodigoBarras:  "7891234567890",
			Categoria:     "Informática",
			PrecoCompra:   2800.00,
			PrecoVenda:    3499.90,
			Quantidade:    15,
			EstoqueMinimo: 3,
			FornecedorID:  1,
		},
		{
			Nome:          "Mouse Logitech MX Master 3",
			Descricao:     "Mouse sem fio ergonômico, 4000 DPI",
			CodigoBarras:  "7899876543210",
			Categoria:     "Informática",
			PrecoCompra:   250.00,
			PrecoVenda:    389.90,
			Quantidade:    40,
			EstoqueMinimo: 10,
			FornecedorID:  1,
		},
		{
			Nome:          "Teclado Mecânico Redragon",
			Descricao:     "Teclado mecânico RGB switch Red",
			CodigoBarras:  "7891111222333",
			Categoria:     "Informática",
			PrecoCompra:   180.00,
			PrecoVenda:    279.90,
			Quantidade:    25,
			EstoqueMinimo: 5,
			FornecedorID:  2,
		},
		{
			Nome:          "Monitor LG 24\" Full HD",
			Descricao:     "Monitor IPS 24 polegadas 75Hz HDMI",
			CodigoBarras:  "7894444555666",
			Categoria:     "Informática",
			PrecoCompra:   700.00,
			PrecoVenda:    999.90,
			Quantidade:    8,
			EstoqueMinimo: 2,
			FornecedorID:  2,
		},
		{
			Nome:          "Cadeira Gamer ThunderX3",
			Descricao:     "Cadeira ergonômica reclinável 180°",
			CodigoBarras:  "7897777888999",
			Categoria:     "Mobiliário",
			PrecoCompra:   850.00,
			PrecoVenda:    1299.90,
			Quantidade:    2,
			EstoqueMinimo: 3,
			FornecedorID:  3,
		},
		{
			Nome:          "Headset HyperX Cloud II",
			Descricao:     "Headset gamer 7.1 surround, microfone removível",
			CodigoBarras:  "7890001112223",
			Categoria:     "Áudio",
			PrecoCompra:   320.00,
			PrecoVenda:    499.90,
			Quantidade:    18,
			EstoqueMinimo: 5,
			FornecedorID:  1,
		},
	}

	db.Create(&produtos)
	log.Println("Produtos criados.")
}

func seedMovimentacoes(db *gorm.DB) {
	var count int64
	db.Model(&models.Movimentacao{}).Count(&count)
	if count > 0 {
		return
	}

	movimentacoes := []models.Movimentacao{
		{ProdutoID: 1, UsuarioID: 1, Tipo: "ENTRADA", Quantidade: 15, Observacao: "Compra inicial de estoque"},
		{ProdutoID: 2, UsuarioID: 1, Tipo: "ENTRADA", Quantidade: 40, Observacao: "Compra inicial de estoque"},
		{ProdutoID: 3, UsuarioID: 1, Tipo: "ENTRADA", Quantidade: 25, Observacao: "Compra inicial de estoque"},
		{ProdutoID: 4, UsuarioID: 1, Tipo: "ENTRADA", Quantidade: 8, Observacao: "Compra inicial de estoque"},
		{ProdutoID: 5, UsuarioID: 1, Tipo: "ENTRADA", Quantidade: 2, Observacao: "Compra inicial de estoque"},
		{ProdutoID: 6, UsuarioID: 1, Tipo: "ENTRADA", Quantidade: 18, Observacao: "Compra inicial de estoque"},
		{ProdutoID: 1, UsuarioID: 2, Tipo: "SAIDA", Quantidade: 2, Observacao: "Venda para cliente"},
		{ProdutoID: 2, UsuarioID: 2, Tipo: "SAIDA", Quantidade: 5, Observacao: "Venda para cliente"},
		{ProdutoID: 3, UsuarioID: 3, Tipo: "SAIDA", Quantidade: 3, Observacao: "Venda para cliente"},
		{ProdutoID: 6, UsuarioID: 2, Tipo: "SAIDA", Quantidade: 4, Observacao: "Venda para cliente"},
	}

	db.Create(&movimentacoes)
	log.Println("Movimentações criadas.")
}
