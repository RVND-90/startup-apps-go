package services

import (
	"github.com/RVND-90/startup-apps-go/constants"
	"github.com/RVND-90/startup-apps-go/models/dto"
	"github.com/go-playground/validator/v10"
)

type ValidationService struct {
	validator *validator.Validate
}

func NewValidationService(validator *validator.Validate) *ValidationService {
	return &ValidationService{validator}
}

func (s *ValidationService) Validate(i interface{}) *dto.ResponseDto {
	err:=s.validator.Struct(i)
	
	if err!=nil {
		errs:=err.(validator.ValidationErrors)
		errstr:=[]string{}
		for _, e:=range errs {
			errstr=append(errstr, e.Error())
		}
		return &dto.ResponseDto{
			RC: constants.GetRC(constants.ERROR_VALIDATION_CODE),
			Errors: &errstr,
		}
	}
	return nil
	
}