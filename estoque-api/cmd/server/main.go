// @title           Estoque API
// @version         1.0
// @description     API REST para Sistema de Controle de Estoque 
// @host            localhost:8080
// @BasePath        /
// @securityDefinitions.apikey BearerAuth
// @in              header
// @name            Authorization
// @description     Informe: Bearer {token}
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/seu-usuario/estoque-api/config"
	"github.com/seu-usuario/estoque-api/middleware"
	"github.com/seu-usuario/estoque-api/routes"
	"github.com/seu-usuario/estoque-api/seeders"

	_ "github.com/seu-usuario/estoque-api/docs"
)

func main() {
	config.LoadConfig()
	config.InitLogger()
	defer config.Logger.Sync()

	config.ConnectDatabase()
	seeders.Run(config.DB)

	if config.AppConfig.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.LoggerMiddleware())
	r.Use(gin.Recovery())

	routes.SetupRoutes(r)

	r.Run(":" + config.AppConfig.Port)
}
