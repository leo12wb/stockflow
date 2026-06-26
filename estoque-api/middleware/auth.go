package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/seu-usuario/estoque-api/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.ErrorResponse(c, http.StatusUnauthorized, "token não fornecido")
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.ErrorResponse(c, http.StatusUnauthorized, "formato de token inválido")
			c.Abort()
			return
		}

		claims, err := utils.ValidateToken(parts[1])
		if err != nil {
			utils.ErrorResponse(c, http.StatusUnauthorized, "token inválido ou expirado")
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("user_email", claims.Email)
		c.Set("user_perfil", claims.Perfil)
		c.Next()
	}
}

func RequirePerfil(perfis ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		perfil, exists := c.Get("user_perfil")
		if !exists {
			utils.ErrorResponse(c, http.StatusForbidden, "acesso negado")
			c.Abort()
			return
		}

		for _, p := range perfis {
			if perfil == p {
				c.Next()
				return
			}
		}

		utils.ErrorResponse(c, http.StatusForbidden, "sem permissão para esta ação")
		c.Abort()
	}
}
