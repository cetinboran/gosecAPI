package main

import (
	"github.com/cetinboran/gosecAPI/routers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{})

	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins: "http://localhost:3000", // Eklentinin olduğu etki alanını buraya ekleyin
	// 	AllowHeaders: "Origin, Content-Type, Accept",
	// 	AllowMethods: "GET, POST, PUT, DELETE",
	// }))

	routers.SetRouters(app)

	app.Listen(":3000")
}
