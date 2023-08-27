package api

import (
	"github.com/cetinboran/gosecAPI/database/config"
	"github.com/cetinboran/gosecAPI/database/password"
	"github.com/cetinboran/gosecAPI/database/user"
	"github.com/gofiber/fiber/v2"
)

func Users(c *fiber.Ctx) error {
	return c.JSON(user.GetAllUsers())
}

func UsersWId(c *fiber.Ctx) error {
	idStr := c.Params("id")
	user, err := user.GetUsersById(idStr)

	// Eğer user nil değil ise user döndür nil ise hatayı döndür.

	if user != nil {
		return c.JSON(user)
	}

	return c.JSON(err)
}

func Configs(c *fiber.Ctx) error {
	return c.JSON(config.GetAllConfigs())
}

func ConfigWUserId(c *fiber.Ctx) error {
	idStr := c.Params("id")
	config, err := config.GetConfigByUserId(idStr)

	// Eğer user nil değil ise user döndür nil ise hatayı döndür.

	if config != nil {
		return c.JSON(config)
	}

	return c.JSON(err)
}

func PasswordWUserId(c *fiber.Ctx) error {
	idStr := c.Params("id")
	config, err := password.GetPasswordsByUserId(idStr)

	// Eğer user nil değil ise user döndür nil ise hatayı döndür.

	if config != nil {
		return c.JSON(config)
	}

	return c.JSON(err)
}
