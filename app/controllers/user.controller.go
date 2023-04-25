package controllers

import (
	"github.com/alejmendez/goApiRest/app/dto"
	model "github.com/alejmendez/goApiRest/app/models"
	"github.com/alejmendez/goApiRest/app/services"
	"github.com/alejmendez/goApiRest/app/utils"

	"github.com/gofiber/fiber/v2"
)

func NewUserController(userService services.UserService) UserController {
	return &userController{
		sevice: userService,
	}
}

type UserController interface {
	Get(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type userController struct {
	sevice services.UserService
}

func responseUser(user *model.User) *dto.UserResponse {
	return &dto.UserResponse{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}
}

// GetUser get a user
func (c *userController) Get(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user, err := c.sevice.Get(id)

	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "No user found with ID")
	}

	return ctx.JSON(responseUser(user))
}

// CreateUser new user
func (c *userController) Create(ctx *fiber.Ctx) error {
	UserRequest := new(dto.UserRequest)

	if err := utils.ParseBodyAndValidate(ctx, UserRequest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Review your input")
	}

	user := &model.User{
		Username: UserRequest.Username,
		Email:    UserRequest.Email,
		Password: UserRequest.Password,
	}

	user, err := c.sevice.Create(user)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Couldn't create user")
	}

	return ctx.JSON(responseUser(user))
}

// UpdateUser update user
func (c *userController) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	userRequest := new(dto.UserUpdateRequest)

	if err := utils.ParseBodyAndValidate(ctx, userRequest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Review your input")
	}

	user := &model.User{
		Username: userRequest.Username,
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}

	user, err := c.sevice.Update(id, user)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Couldn't create user")
	}

	return ctx.JSON(responseUser(user))
}

// DeleteUser delete user
func (c *userController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	success, err := c.sevice.Delete(id)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Couldn't delete the record")
	}

	if !success {
		return fiber.NewError(fiber.StatusInternalServerError, "Couldn't delete the record")
	}

	return ctx.JSON(fiber.Map{"message": "User successfully deleted"})
}
