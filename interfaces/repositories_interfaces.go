package interfaces

import "github.com/RVND-90/startup-apps-go/models/entities"

type IUserRepository interface {
	CreateUser(user *entities.User) error
}