package router

import (
	"github.com/B1gdawg0/se-project-backend/internal/adapters/rest"
	"github.com/gofiber/fiber/v2"
)

func RegisterApiRouter(app *fiber.App, rqHandler *rest.Handler){
	user := app.Group("user")
	
	user.Get("", rqHandler.User.FindAll)

	user.Post("/create", rqHandler.User.CreateUser)
}