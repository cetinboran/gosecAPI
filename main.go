package main

import (
	"github.com/cetinboran/gosecAPI/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New(fiber.Config{})

	// Live server ip sini alloworigins e yazdığımda sorun çözüldü fetch atabiliyorum
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Eklentinin olduğu etki alanını buraya ekleyin
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	routers.SetRouters(app)

	app.Listen(":3000")
}
