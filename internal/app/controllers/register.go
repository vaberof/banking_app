package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vaberof/banking_app/internal/app/constants"
	"github.com/vaberof/banking_app/internal/app/model"
	"github.com/vaberof/banking_app/internal/app/service"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	inputUsername := data["username"]

	_, err = service.GetUser(data)
	if err == nil {
		c.Status(fiber.StatusConflict)
		return c.JSON(fiber.Map{
			"message": constants.UserAlreadyExists,
		})
	}

	inputPassword := data["password"]
	hashedPassword := service.HashPassword(inputPassword)

	user := service.CreateUser(inputUsername, hashedPassword)

	service.CreateUserInDatabase(user)

	userAccount := model.NewAccount()

	userAccount.SetUserID(user.ID)
	userAccount.SetOwner(user.Username)
	userAccount.SetMainAccountType()
	userAccount.SetInitialBalance()

	service.CreateAccountInDatabase(userAccount)

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": constants.Success,
	})
}
