package requests

type BalanceRequest struct {
	UserId int `header:"X-UserId"`
}
