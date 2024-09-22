package queries

import (
	"github.com/B1gdawg0/se-project-backend/internal/transaction/requests"
	"github.com/B1gdawg0/se-project-backend/internal/transaction/response"
)

type Database interface{
	CreateUser(rq *requests.CreateUserRequest) (*response.UserCreateResponse, error)
	GetUsers() (*response.GetUsersResponse, error)
}