package params

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	RePassword string `json:"re_password"`
	Email string `json:"email"`
}
