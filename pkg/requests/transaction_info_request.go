package requests

type TransactionInfoRequest struct {
	UserId int `header:"X-UserId"`
}
