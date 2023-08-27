package api

import (
	"github.com/cetinboran/gosecAPI/database/auth"
	"github.com/gofiber/fiber/v2"
)

type Login struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func AuthCheck(c *fiber.Ctx) error {
	newLogin := Login{}
	if err := c.BodyParser(&newLogin); err != nil {
		return err
	}

	user, err := auth.Check(newLogin.Username, newLogin.Password)

	if err != nil {
		return c.JSON(err)
	}

	return c.JSON(user)
}
