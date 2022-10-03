package controllers

import (
	"github.com/alejmendez/goApiRest/app/dto"
	"github.com/alejmendez/goApiRest/app/services"

	"github.com/gofiber/fiber/v2"
)

// Login get user and password
func Login(c *fiber.Ctx) error {
	var input dto.LoginDto

	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Error on login request")
	}
	identity := input.Identity
	pass := input.Password

	token, err := services.GenerateToken(identity, pass)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": token})
}
