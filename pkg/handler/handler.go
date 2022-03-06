package handler

import (
	"github.com/Up2110H/e-wallet/pkg/middleware"
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *gin.Engine {
	middlewares := middleware.NewMiddleware()
	router := gin.New()

	wallet := router.Group("wallet").Use(middlewares.WalletMiddleware())
	{
		wallet.GET("hello", h.hello)
	}

	return router
}
