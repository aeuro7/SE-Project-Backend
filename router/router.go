package router

import (
	"github.com/B1gdawg0/se-project-backend/internal/adapters/rest"
	"github.com/B1gdawg0/se-project-backend/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterApiRouter(app *fiber.App, rqHandler *rest.Handler){
	user := app.Group("/users")
	auth := app.Group("/auth")

	auth.Post("/login", rqHandler.Auth.Login)
	auth.Post("/register", rqHandler.Auth.Register)

	user.Use(middleware.CheckJWT)
	user.Get("", rqHandler.User.GetUsers)
	user.Get("/id=:id", rqHandler.User.GetUserByID)
	user.Get("/email=:email",rqHandler.User.GetUserByEmail)
}