package middleware

import "github.com/gin-gonic/gin"

func (m *Middleware) WalletMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
