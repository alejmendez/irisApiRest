package controllers

import (
	"github.com/alejmendez/goApiRest/core/config"
	"github.com/gofiber/fiber/v2"
)

func Health(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"version": config.Conf.AppVersion})
}
