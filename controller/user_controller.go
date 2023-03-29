package controller

import (
	"github.com/go-chi/chi/middleware"
	"github.com/gofiber/fiber/v2"
	"gitlab.com/learn-micorservices/user-service/middleware"
	"gitlab.com/learn-micorservices/user-service/model/web"
)

type UserController struct {
}

func (controller *UserController) NewUserRouter(app *fiber.App) {
	user := app.Group("/")

	user.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
			Code: fiber.StatusOK,
			Status: true,
			Message: "ok",
		})
	})

	user.Use(middleware.IsAuthenticated)
	user.Get("/", controller.)
}

func (controller *UserController) CreateUser (ctx *fiber.Ctx) error {
	
}