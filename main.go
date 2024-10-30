package main

import (
	"github.com/B1gdawg0/se-project-backend/internal/wire"
	"github.com/B1gdawg0/se-project-backend/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main(){
	handler := wire.InitHandler()

	app := fiber.New()

	app.Use(
		cors.New(cors.Config{
			AllowOrigins: "http://localhost:3000",
			AllowHeaders: "Origin, Content-Type, Accept",
			AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
		}),
	)

	router.RegisterApiRouter(app, handler)


	app.Listen(":8000")
}