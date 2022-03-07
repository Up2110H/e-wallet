package middleware

import (
	"github.com/Up2110H/e-wallet/pkg/store"
	"github.com/go-playground/validator/v10"
)

type Middleware struct {
	validate *validator.Validate
	store    *store.Store
}

func NewMiddleware(validate *validator.Validate, store *store.Store) *Middleware {
	return &Middleware{validate, store}
}
