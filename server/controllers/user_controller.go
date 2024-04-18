package controllers

import (
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"result": "parse error",
			"code": "01",
			"error": err.Error(),
		})
	}

	rsp:=u.userService.Add(dto)
	return c.Status(rsp.HttpResponseCode).JSON(rsp)
}