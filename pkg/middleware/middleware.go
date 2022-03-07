package middleware

import "github.com/go-playground/validator/v10"

type Middleware struct {
	validate *validator.Validate
}

func NewMiddleware(validate *validator.Validate) *Middleware {
	return &Middleware{validate}
}
