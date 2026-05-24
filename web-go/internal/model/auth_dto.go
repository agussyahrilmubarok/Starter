package model

type SignUpRequest struct {
	Name     string `form:"name" validate:"required,min=2,max=100"`
	Email    string `form:"email" validate:"required,email,max=100"`
	Password string `form:"password" validate:"required,min=8,max=255"`
}

type SignInRequest struct {
	Email    string `form:"email" validate:"required,email,max=100"`
	Password string `form:"password" validate:"required,min=8,max=255"`
}
