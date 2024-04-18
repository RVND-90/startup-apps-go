package services

import (
	"github.com/RVND-90/startup-apps-go/interfaces"
	"github.com/RVND-90/startup-apps-go/models/dto"
)

type UserService struct {
	repo interfaces.IUserRepository
	validator *ValidationService
}

func NewUserService(repo interfaces.IUserRepository, validator *ValidationService) interfaces.IUserService {
	return &UserService{repo, validator}
}

// Add implements interfaces.IUserService.
func (s *UserService) Add(create dto.CreateUserDto) *dto.ResponseDto {
	errValidator:=s.validator.Validate(create)
	if errValidator!=nil {
		return errValidator
	}
	
	return &dto.ResponseDto{
		Data: &dto.CreateSuccessResponse{
			Id: 1,
		},
		HttpResponseCode: 200,
		Code: "00",
		Message: "success",
	}
}
