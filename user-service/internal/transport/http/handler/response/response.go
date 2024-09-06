package response

type CreateUserResponse struct {
	ID       int    `json:"id"  binding:"required"`
	Email    string `json:"email"  binding:"required"`
	Username string `json:"username"  binding:"required"`
}
