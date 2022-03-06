package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!")
}
