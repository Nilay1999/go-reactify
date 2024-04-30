package types

type UserType struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Gender   string `json:"gender" binding:"required"`
}
