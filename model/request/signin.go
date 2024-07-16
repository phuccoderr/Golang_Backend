package request

type SignIn struct {
	Username string `json:"username" binding:"required,email,min=8,max=255"`
	Password string `json:"password" binding:"required,min=4,max=20"`
}
