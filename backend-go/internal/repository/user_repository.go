package repository

//go:generate mockery --name=IUserRepository
type IUserRepository interface {
}

type userRepository struct {
}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}
