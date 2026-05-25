package model

type AuthResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

type SignUpRequest struct {
	Name     string `json:"name" binding:"required,min=2,max=100"`
	Email    string `json:"email" binding:"required,email,max=100"`
	Password string `json:"password" binding:"required,min=8,max=255"`
}

type SignInRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
