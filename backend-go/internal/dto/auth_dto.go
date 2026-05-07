package dto

type AuthDTO struct {
	Token string  `json:"token"`
	User  UserDTO `json:"user"`
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
