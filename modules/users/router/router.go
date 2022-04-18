package router

import (
	"github.com/alejmendez/goApiRest/middleware"
	controller "github.com/alejmendez/goApiRest/modules/users/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(api fiber.Router) {
	user := api.Group("/users")
	user.Get("/:id", controller.GetUser)
	user.Post("/", controller.CreateUser)
	user.Patch("/:id", middleware.Protected(), controller.UpdateUser)
	user.Delete("/:id", middleware.Protected(), controller.DeleteUser)
}
