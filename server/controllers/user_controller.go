package controllers

import (
	"github.com/RVND-90/startup-apps-go/constants"
	"github.com/RVND-90/startup-apps-go/interfaces"
	"github.com/RVND-90/startup-apps-go/models/dto"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService interfaces.IUserService
	
}

func NewUserController(svc interfaces.IUserService) *UserController {
	return &UserController{ svc }
}

func (u *UserController) Add(c *fiber.Ctx) error {
	var dto dto.CreateUserDto
	err:=c.BodyParser(&dto)
	if err!=nil {
		return c.Status(fiber.StatusBadRequest).JSON(constants.GetRC(constants.ERROR_PARSE_CODE))
	}

	rsp:=u.userService.Add(dto)
	return c.Status(rsp.HttpResponseCode).JSON(rsp)
}