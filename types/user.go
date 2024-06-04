package types

type UserType struct {
	Username string `json:"username" binding:"required" required:"username is required"`
	Password string `json:"password" binding:"required" required:"password is required"`
	Email    string `json:"email" binding:"required" required:"email is required"`
	Gender   string `json:"gender" binding:"required,oneof=male female" required:"gender is required" oneof:"Gender must be either 'male' or 'female'"`
	Age      uint   `json:"age" binding:"required,gte=12" gte:"Minimum age is 12"`
}
