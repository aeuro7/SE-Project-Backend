package rest

import (
	"github.com/B1gdawg0/se-project-backend/internal/adapters/rest/handlers"
)

type Handler struct{
	User *handlers.UserRestHandler
}

func ProvideHandler(user *handlers.UserRestHandler) *Handler{
	return &Handler{
		User: user,
	}
}