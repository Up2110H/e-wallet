package model

import "github.com/lib/pq"

type Transaction struct {
	ID        int
	WalletId  int
	Amount    float64
	CreatedAt pq.NullTime
	Wallet    *Wallet `json:"-"`
}
