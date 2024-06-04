package types

type AuthType struct {
	Identifier string `json:"identifier" binding:"required" required:"username or email is required"`
	Password   string `json:"password" binding:"required" required:"password is required"`
}
