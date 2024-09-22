package queries

import (
	"math/rand"

	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"
	"github.com/B1gdawg0/se-project-backend/internal/transaction/requests"
	"github.com/B1gdawg0/se-project-backend/internal/transaction/response"
	"gorm.io/gorm"
)

type PGGormDB struct{
	db *gorm.DB
}

func New(db *gorm.DB) Database{
	return &PGGormDB{
		db: db,
	}
}

func (pg *PGGormDB) GetUsers() (*response.GetUsersResponse, error){
	var users = new([]entities.User)

	if err := pg.db.Find(users).Error; err != nil{
		return nil, err
	}

	list := make([]response.GetUserResponse, len(*users))
	for i, ctx := range *users{
		list[i] = response.GetUserResponse{
			ID: ctx.ID,
			Fname: ctx.Fname,
			Lname: ctx.Lname,
		}
	}

	usersResponse := &response.GetUsersResponse{
        Users: list,
    }

	return usersResponse, nil
}


func (pg *PGGormDB)CreateUser(rq *requests.CreateUserRequest) (*response.UserCreateResponse, error){
	ID := rand.Intn(201)
	
	newUser := &entities.User{
		ID: ID,
		Fname: rq.Fname,
		Lname: rq.Lname,
	}

	if err := pg.db.Create(newUser).Error; err != nil{
		return nil, err
	}

	response := &response.UserCreateResponse{
		ID: ID,
		Fname: newUser.Fname,
		Lname:  newUser.Lname,
	}

	return response,nil
}