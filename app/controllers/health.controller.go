package controllers

import (
	"github.com/alejmendez/goApiRest/core/config"
	"github.com/gofiber/fiber/v2"
)

func NewHealthController(conf *config.Config) HealthController {
	return &healthController{
		Conf: conf,
	}
}

type HealthController interface {
	Get(ctx *fiber.Ctx) error
}

type healthController struct {
	Conf *config.Config
}

func (c *healthController) Get(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"version": c.Conf.AppVersion})
}
