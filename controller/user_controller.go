package controller

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/learn-micorservices/user-service/config"
	"gitlab.com/learn-micorservices/user-service/exception"
	"gitlab.com/learn-micorservices/user-service/helper"
	"gitlab.com/learn-micorservices/user-service/middleware"
	"gitlab.com/learn-micorservices/user-service/model/web"
	"gitlab.com/learn-micorservices/user-service/service"
)

type UserController interface {
	NewUserRouter(app *fiber.App)
}
type userController struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		UserService: userService,
	}
}

func (controller *userController) NewUserRouter(app *fiber.App) {
	user := app.Group(config.EndpointPrefix)

	user.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
			Code: fiber.StatusOK,
			Status: true,
			Message: "ok",
		})
	})

	user.Use(middleware.IsAuthenticated)
	user.Get("/", controller.GetAllUsers)
	user.Get("/:id", controller.GetUserbyID)
	user.Post("/create", controller.CreateUser)
	user.Put("/:id", controller.UpdateUser)
	user.Delete("/:id", controller.DeleteUser)
}

func (controller *userController) CreateUser (ctx *fiber.Ctx) error {
	claims := ctx.Locals("claims").(helper.JWTClaims)
	
	request := new(web.CreateUserRequest)
	err := ctx.BodyParser(&request)
	helper.FatalIfError(err)

	userResponse, err := controller.UserService.CreateUser(ctx.Context(), claims, *request)
	if err != nil {
		return exception.ErrorHandler(ctx, err)
	}

	// KAFKA

	return ctx.Status(fiber.StatusCreated).JSON(web.WebResponse{
		Code: fiber.StatusCreated,
		Status: true,
		Message: "success",
		Data: userResponse,
	})
}

func (controller *userController) GetAllUsers (ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusCreated).JSON(web.WebResponse{
		Code: fiber.StatusCreated,
		Status: true,
		Message: "success",
		Data: nil,
	})
}

func (controller *userController) GetUserbyID (ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusCreated).JSON(web.WebResponse{
		Code: fiber.StatusCreated,
		Status: true,
		Message: "success",
		Data: nil,
	})
}


func (controller *userController) UpdateUser (ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusCreated).JSON(web.WebResponse{
		Code: fiber.StatusCreated,
		Status: true,
		Message: "success",
		Data: nil,
	})
}

func (controller *userController) DeleteUser (ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusCreated).JSON(web.WebResponse{
		Code: fiber.StatusCreated,
		Status: true,
		Message: "success",
		Data: nil,
	})
}