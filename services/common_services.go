package services

import (
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
			HttpResponseCode: 400,
			Code: "02",
			Message: "validation error",
			Errors: &errstr,
		}
	}
	return nil
	
}