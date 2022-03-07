package handler

import (
	"github.com/Up2110H/e-wallet/pkg/middleware"
	"github.com/Up2110H/e-wallet/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	validate *validator.Validate
	store    *store.Store
}

func NewHandler(store *store.Store) *Handler {
	return &Handler{validator.New(), store}
}

func (h *Handler) InitRoutes() *gin.Engine {
	middlewares := middleware.NewMiddleware(h.validate)
	router := gin.New()

	wallet := router.Group("wallet").Use(middlewares.WalletMiddleware())
	{
		wallet.POST("check", h.check)
		wallet.POST("new-transaction", h.newTransaction)
		wallet.POST("monthly-transaction-info", h.getMonthlyTransactionsInfo)
		wallet.POST("balance", h.getBalance)
	}

	return router
}
