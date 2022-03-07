package model

type Wallet struct {
	ID           int
	UserId       int
	Amount       float64
	User         *User          `json:"-"`
	Transactions []*Transaction `json:"-"`
}
