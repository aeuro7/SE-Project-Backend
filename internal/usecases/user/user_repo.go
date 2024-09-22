package user

import (
	"github.com/B1gdawg0/se-project-backend/internal/transaction/requests"
	"github.com/B1gdawg0/se-project-backend/internal/transaction/response"
)

type UserRepository interface{
	FindAll() (*response.GetUsersResponse, error)
	Save(rq  *requests.CreateUserRequest) (*response.UserCreateResponse, error)
}