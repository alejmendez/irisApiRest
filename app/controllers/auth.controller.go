package controllers

import (
	"github.com/alejmendez/goApiRest/app/dto"
	"github.com/alejmendez/goApiRest/app/services"

	"github.com/gofiber/fiber/v2"
)

func NewAuthController(service services.AuthServices) AuthController {
	return &authController{
		Service: service,
	}
}

type AuthController interface {
	Login(ctx *fiber.Ctx) error
}

type authController struct {
	Service services.AuthServices
}

// Login get user and password
func (c *authController) Login(ctx *fiber.Ctx) error {
	var input dto.LoginDto

	if err := ctx.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Error on login request")
	}

	token, err := c.Service.GenerateToken(input.Email, input.Password)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.JSON(fiber.Map{"token": token})
}
