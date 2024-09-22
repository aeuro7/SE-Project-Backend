package repositories

import (
	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/queries"
	"github.com/B1gdawg0/se-project-backend/internal/transaction/requests"
	"github.com/B1gdawg0/se-project-backend/internal/transaction/response"
	"github.com/B1gdawg0/se-project-backend/internal/usecases/user"
	"gorm.io/gorm"
)


type UserRepositoryImpl struct{
	Queries queries.Database
}

func ProvideUserRepository(db *gorm.DB) user.UserRepository{
	return &UserRepositoryImpl{
		Queries: queries.New(db),
	}
}

func (uri *UserRepositoryImpl) FindAll() (*response.GetUsersResponse, error){
	response, err := uri.Queries.GetUsers()

	if err != nil{
		return nil, err
	}

	return response, nil
}


func (uri *UserRepositoryImpl) Save(rq *requests.CreateUserRequest) (*response.UserCreateResponse, error){
	response, err := uri.Queries.CreateUser(rq)

	if err != nil{
		return nil, err
	}

	return response, nil

}