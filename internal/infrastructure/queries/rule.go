package queries

import (
	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"
	"github.com/B1gdawg0/se-project-backend/internal/transaction/response"
	"github.com/jackc/pgx/v5/pgtype"
)

type Database interface{
	CreateUser(rq *entities.User) (*response.CreateUserResponse, error)
	FindUserByEmail(email string) (*response.FindUserResponse, error)
	FindUserByID(id pgtype.UUID) (*response.FindUserResponse, error)
	FindAll() (*response.FindUsersResponse, error)
} 