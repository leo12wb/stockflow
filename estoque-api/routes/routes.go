package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/seu-usuario/estoque-api/controllers"
	"github.com/seu-usuario/estoque-api/middleware"
)

func SetupRoutes(r *gin.Engine) {
	authCtrl := controllers.NewAuthController()
	produtoCtrl := controllers.NewProdutoController()
	fornecedorCtrl := controllers.NewFornecedorController()
	movCtrl := controllers.NewMovimentacaoController()
	relCtrl := controllers.NewRelatorioController()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := r.Group("/auth")
	{
		auth.POST("/login", authCtrl.Login)
		auth.POST("/register", authCtrl.Register)
		auth.GET("/me", middleware.AuthMiddleware(), authCtrl.Me)
	}

	api := r.Group("/")
	api.Use(middleware.AuthMiddleware())
	{
		produtos := api.Group("/produtos")
		{
			produtos.GET("", produtoCtrl.FindAll)
			produtos.GET("/codigo/:codigo", produtoCtrl.FindByCodigoBarras)
			produtos.GET("/:id", produtoCtrl.FindByID)
			produtos.POST("", produtoCtrl.Create)
			produtos.PUT("/:id", produtoCtrl.Update)
			produtos.DELETE("/:id", produtoCtrl.Delete)
			produtos.PATCH("/:id/restaurar", produtoCtrl.Restore)
			produtos.DELETE("/:id/force", produtoCtrl.ForceDelete)
		}

		fornecedores := api.Group("/fornecedores")
		{
			fornecedores.GET("", fornecedorCtrl.FindAll)
			fornecedores.GET("/:id", fornecedorCtrl.FindByID)
			fornecedores.POST("", fornecedorCtrl.Create)
			fornecedores.PUT("/:id", fornecedorCtrl.Update)
			fornecedores.DELETE("/:id", fornecedorCtrl.Delete)
			fornecedores.PATCH("/:id/restaurar", fornecedorCtrl.Restore)
			fornecedores.DELETE("/:id/force", fornecedorCtrl.ForceDelete)
		}

		movimentacoes := api.Group("/movimentacoes")
		{
			movimentacoes.GET("", movCtrl.FindAll)
			movimentacoes.POST("", movCtrl.Create)
		}

		relatorios := api.Group("/relatorios")
		{
			relatorios.GET("/inventario", relCtrl.Inventario)
			relatorios.GET("/estoque-baixo", relCtrl.EstoqueBaixo)
			relatorios.GET("/entradas", relCtrl.Entradas)
			relatorios.GET("/saidas", relCtrl.Saidas)
			relatorios.GET("/mais-vendidos", relCtrl.MaisVendidos)
			relatorios.GET("/movimentacoes", relCtrl.Movimentacoes)
		}
	}
}
