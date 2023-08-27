package routers

import (
	"github.com/cetinboran/gosecAPI/api"
	"github.com/gofiber/fiber/v2"
)

func SetRouters(app *fiber.App) {
	Users := app.Group("/users")

	Users.Get("/", api.Users)
	Users.Get("/:id", api.UsersWId)

	Config := app.Group("/config")

	Config.Get("/", api.Configs)
	Config.Get("/:id", api.ConfigWUserId)

	Passwords := app.Group("/passwords")

	Passwords.Get("/:id", api.PasswordWUserId)
}
