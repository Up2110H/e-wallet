package main

import (
	eWallet "github.com/Up2110H/e-wallet"
)

func main() {
	server := eWallet.NewServer()
	server.Run()
}
