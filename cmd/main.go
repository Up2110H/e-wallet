package main

import (
	eWallet "github.com/Up2110H/e-wallet"
	"github.com/Up2110H/e-wallet/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	server := new(eWallet.Server)
	if err := server.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Ошибка при запуске сервера %s", err.Error())
	}
}
