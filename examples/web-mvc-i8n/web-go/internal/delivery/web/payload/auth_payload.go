package payload

type SignUpRequest struct {
	Name     string `form:"name"     binding:"required,min=2,max=100"`
	Email    string `form:"email"    binding:"required,email,max=150"`
	Password string `form:"password" binding:"required,min=8,max=72"`
}

type SignInRequest struct {
	Email    string `form:"email"    binding:"required,email,max=150"`
	Password string `form:"password" binding:"required,min=8,max=72"`
}
