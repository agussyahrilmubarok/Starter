package service

//go:generate mockery --name=IUserService
type IUserService interface {
}

type userService struct {
}

func NewUserService() IUserService {
	return &userService{}
}
