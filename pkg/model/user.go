package model

import "github.com/lib/pq"

type User struct {
	ID           int
	Email        string
	Key          string
	IdentifiedAt pq.NullTime
	Wallet       *Wallet `json:"-"`
}
