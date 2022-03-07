package handler

import (
	"github.com/Up2110H/e-wallet/pkg/model"
	"github.com/Up2110H/e-wallet/pkg/requests"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"time"
)

func (h *Handler) check(c *gin.Context) {
	var request requests.CheckRequest
	c.BindHeader(&request)

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := h.validate.Struct(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, err := h.store.User().GetById(request.UserId)

	c.JSON(http.StatusOK, gin.H{"is_exist": err == nil && user.Email == request.Email})
}

func (h *Handler) newTransaction(c *gin.Context) {
	var request requests.TransactionRequest
	c.BindHeader(&request)

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := h.validate.Struct(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, err := h.store.User().GetById(request.UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	limit, _ := strconv.ParseFloat(os.Getenv("UNIDENTIFIED_USER_LIMIT"), 64)

	if user.Identified {
		limit, _ = strconv.ParseFloat(os.Getenv("IDENTIFIED_USER_LIMIT"), 64)
	}

	if limit < user.Wallet.Amount+request.Amount {
		c.JSON(http.StatusOK, gin.H{"message": "Limit is exceeded"})
		return
	}

	transaction := &model.Transaction{
		WalletId: user.Wallet.ID,
		Amount:   request.Amount,
	}

	_, err = h.store.Transaction().Create(transaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func (h *Handler) getMonthlyTransactionsInfo(c *gin.Context) {
	var request requests.TransactionInfoRequest
	c.BindHeader(&request)

	user, err := h.store.User().GetById(request.UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	now := time.Now()
	var count int
	var total float64
	count, total, err = h.store.User().GetTransactionsInfo(user, now.Add(-time.Hour*24*30), now)
	c.JSON(http.StatusOK, gin.H{"count": count, "total": total})
}

func (h *Handler) getBalance(c *gin.Context) {
	var request requests.BalanceRequest
	c.BindHeader(&request)

	user, err := h.store.User().GetById(request.UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"balance": user.Wallet.Amount})
}
