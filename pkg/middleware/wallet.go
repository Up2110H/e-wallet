package middleware

import (
	"github.com/Up2110H/e-wallet/pkg/requests"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (m *Middleware) WalletMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var x requests.X
		if err := c.ShouldBindHeader(&x); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		if err := m.validate.Struct(&x); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.Next()
	}
}
