package requests

type X struct {
	UserId int    `header:"X-UserId" binding:"required" validate:"required,gt=0"`
	Digest string `header:"X-Digest"`
}
