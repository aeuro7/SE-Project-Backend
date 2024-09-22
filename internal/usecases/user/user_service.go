package user

import (
	"errors"
	"github.com/B1gdawg0/se-project-backend/internal/transaction/requests"
	"github.com/B1gdawg0/se-project-backend/internal/transaction/response"
)

type UserUseCase interface{
	CreateUser(rq *requests.CreateUserRequest) (*response.UserCreateResponse, error)
	FindAll()(*response.GetUsersResponse, error)
}


type UserService struct{
	repo UserRepository
}


func ProvideUserService(repo UserRepository) UserUseCase{
	return &UserService{
		repo: repo,
	}
}

func (uc *UserService) FindAll() (*response.GetUsersResponse, error){
	list, err := uc.repo.FindAll()

	if err != nil{
		return nil, err
	}

	return list, nil
} 

func (uc *UserService) CreateUser(rq *requests.CreateUserRequest) (*response.UserCreateResponse, error){
	if rq.Fname == "" || rq.Lname == ""{
		return nil, errors.New("value can't be empty string")
	}

	user, err := uc.repo.Save(rq)

	if err != nil{
		return nil, err
	}

	return  user, nil
}