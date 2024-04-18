package interfaces

import "github.com/RVND-90/startup-apps-go/models/dto"

type IUserService interface {
	Add(create dto.CreateUserDto) *dto.ResponseDto
}