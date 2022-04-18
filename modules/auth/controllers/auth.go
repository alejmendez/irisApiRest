package controllers

import (
	dtoAuth "github.com/alejmendez/goApiRest/modules/auth/dto"
	authService "github.com/alejmendez/goApiRest/modules/auth/services"

	"github.com/gofiber/fiber/v2"
)

// Login get user and password
func Login(c *fiber.Ctx) error {
	var input dtoAuth.LoginDto

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err})
	}
	identity := input.Identity
	pass := input.Password

	token, err := authService.GenerateToken(identity, pass)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": token})
}
