package server

import (
	"fmt"

	"github.com/RVND-90/startup-apps-go/repositories"
	"github.com/RVND-90/startup-apps-go/server/controllers"
	"github.com/RVND-90/startup-apps-go/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AppRoute struct {
	//all of controller
	userController controllers.UserController
}

func ServerStartup() {
	app:= fiber.New()
	Bootstrap(app)
}

func Bootstrap(app *fiber.App) {
	vl := validator.New()
	valdationservice:=services.NewValidationService(vl)
	userrepository:=repositories.NewUserRepository()
	userservice:=services.NewUserService(userrepository, valdationservice)

	appRouter := AppRoute{}
	appRouter.userController = *controllers.NewUserController(userservice)

	//Route
	appRouter.setup(app)


	fmt.Println("server running on port 3000")
	app.Listen(":3000")
	
	
}


func (router *AppRoute) setup(app *fiber.App) {
	//user
	app.Post("/user", router.userController.Add)
}


