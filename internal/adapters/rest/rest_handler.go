package rest

import (
	"github.com/B1gdawg0/se-project-backend/internal/adapters/rest/handlers"
)

type Handler struct{
	User *handlers.UserRestHandler
	Auth *handlers.AuthRestHandler
	Table *handlers.TableRestHandler
	Order *handlers.OrderRestHandler
}

func ProvideHandler(user *handlers.UserRestHandler, auth *handlers.AuthRestHandler, table *handlers.TableRestHandler, order *handlers.OrderRestHandler) *Handler{
	return &Handler{
		User: user,
		Auth: auth,
		Table: table,
		Order: order,
	}
}