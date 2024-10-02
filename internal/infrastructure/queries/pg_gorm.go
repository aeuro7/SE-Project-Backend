package queries

import (
	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"
	"github.com/B1gdawg0/se-project-backend/internal/transaction/response"
	"github.com/jackc/pgx/v5/pgtype"
	"gorm.io/gorm"
)

type PGGormDB struct {
	db *gorm.DB
}

func New(db *gorm.DB) Database {
	return &PGGormDB{
		db: db,
	}
}

func (pg *PGGormDB) FindUserByID(id pgtype.UUID) (*response.FindUserResponse, error) {
	user := new(entities.User)

	if err := pg.db.Where("id = ?", id).Find(user); err.Error != nil {
		return nil, err.Error
	}

	response := &response.FindUserResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Password: user.Password,
		Phone: user.Phone,
	}

	return response, nil
}

func (pg *PGGormDB) FindUserByEmail(email string) (*response.FindUserResponse, error) {
	user := new(entities.User)

	if err := pg.db.Where("c_email = ?", email).Find(user); err.Error != nil{
		return nil, err.Error
	}

	response := &response.FindUserResponse{
		ID: user.ID,
		Email: user.Email,
		Name: user.Name,
		Password: user.Password,
		Phone: user.Phone,
	}

	return response, nil
}

func (pg *PGGormDB) FindAll() (*response.FindUsersResponse, error) {
	var users = new([]entities.User)

	if err := pg.db.Find(users).Error; err != nil {
		return nil, err
	}

	list := make([]response.FindUserResponse, len(*users))
	for i, ctx := range *users {
		list[i] = response.FindUserResponse{
			ID: ctx.ID,
			Name: ctx.Name,
			Email: ctx.Email,
			Password: ctx.Password,
			Phone: ctx.Phone,
		}
	}

	usersResponse := &response.FindUsersResponse{
		Users: list,
	}

	return usersResponse, nil
}

func (pg *PGGormDB) CreateUser(rq *entities.User) (*response.CreateUserResponse, error) {

	if err := pg.db.Create(rq).Error; err != nil {
		return nil, err
	}

	response := &response.CreateUserResponse{
		ID: rq.ID,
		Name: rq.Name,
		Email: rq.Email,
		Password: rq.Password,
		Phone: rq.Phone,
	}

	return response, nil
}
