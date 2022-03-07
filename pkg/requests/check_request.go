package requests

type CheckRequest struct {
	UserId int    `header:"X-UserId"`
	Email  string `json:"email" validate:"required,email"`
}
