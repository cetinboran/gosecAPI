package main

import (
	"github.com/cetinboran/gosecAPI/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./views", ".html")

	// Şimdilik kalsın ama json yolladığım için endpointlere büyük ihtimalle tepplate engine e gerek yok.
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins: "http://localhost:3000", // Eklentinin olduğu etki alanını buraya ekleyin
	// 	AllowHeaders: "Origin, Content-Type, Accept",
	// 	AllowMethods: "GET, POST, PUT, DELETE",
	// }))

	app.Static("/", "./views")
	routers.SetRouters(app)

	app.Listen(":3000")
}
