package repositories

import (
	"github.com/RVND-90/startup-apps-go/interfaces"
	"github.com/RVND-90/startup-apps-go/models/entities"
)

type UserRepository struct {
}

func NewUserRepository() interfaces.IUserRepository {
	return &UserRepository{}
}
// CreateUser implements interfaces.IUserRepository.
func (*UserRepository) CreateUser(user *entities.User) error {
	panic("unimplemented")
}

