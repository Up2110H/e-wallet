package requests

type TransactionRequest struct {
	UserId int     `header:"X-UserId"`
	Amount float64 `json:"amount" validate:"required,gt=0"`
}
