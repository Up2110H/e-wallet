package handler

import (
	"github.com/Up2110H/e-wallet/pkg/middleware"
	"github.com/Up2110H/e-wallet/pkg/model"
	"github.com/Up2110H/e-wallet/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Handler struct {
	validate *validator.Validate
	store    *store.Store
}

func NewHandler(store *store.Store) *Handler {
	return &Handler{validator.New(), store}
}

func (h *Handler) InitRoutes() *gin.Engine {
	middlewares := middleware.NewMiddleware(h.validate, h.store)
	router := gin.New()

	wallet := router.Group("wallet").Use(middlewares.WalletMiddleware())
	{
		wallet.POST("check", h.check)
		wallet.POST("new-transaction", h.newTransaction)
		wallet.POST("monthly-transaction-info", h.getMonthlyTransactionsInfo)
		wallet.POST("balance", h.getBalance)
	}

	// Only for testing
	router.POST("get-digest", h.getDigest)
	router.POST("seed-users", h.seedUsers)

	return router
}

// TODO: remove
func (h *Handler) getDigest(c *gin.Context) {
	var request struct {
		UserId int `header:"X-UserId" binding:"required"`
	}

	if err := c.ShouldBindHeader(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, err := h.store.User().GetById(request.UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	body, _ := c.GetRawData()

	c.JSON(http.StatusOK, gin.H{"x-digest": []byte(middleware.MACEncrypt(body, []byte(user.Key)))})
}

// TODO: write normal seeder
func (h *Handler) seedUsers(c *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	users := make([]*model.User, 0)
	for i := 0; i < 10; i++ {
		t := rand.Intn(2)
		identified := false
		limit, _ := strconv.Atoi(os.Getenv("UNIDENTIFIED_USER_LIMIT"))
		if t == 1 {
			identified = true
			limit, _ = strconv.Atoi(os.Getenv("IDENTIFIED_USER_LIMIT"))
		}
		amount := rand.Intn(limit)
		generatedStr := ""

		for j := 0; j < 12; j++ {
			generatedStr += string('a' + rune(rand.Intn(26)))
		}

		user := &model.User{
			Email:      generatedStr + "@gmail.com",
			Key:        generatedStr,
			Identified: identified,
		}

		_, err := h.store.User().Create(user, float64(amount))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		users = append(users, user)
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}
