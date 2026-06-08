package dto

type SignUpRequest struct {
	Name     string `json:"name"     binding:"required,min=2,max=100"`
	Email    string `json:"email"    binding:"required,email,max=150"`
	Password string `json:"password" binding:"required,min=8,max=72"`
}

type SignInRequest struct {
	Email    string `json:"email"    binding:"required,email,max=150"`
	Password string `json:"password" binding:"required,min=8,max=72"`
}

type AuthResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}
