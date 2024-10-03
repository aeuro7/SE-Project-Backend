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

func (pg *PGGormDB) FindAllTable() ([]*entities.Table, error){
	var tables []*entities.Table

	if err := pg.db.Find(&tables).Error; err != nil {
		return nil, err
	}

	return tables,nil
}

// TODO: change return value to entitles.User
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

// TODO: change return value to entitles.User
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

// TODO: change return value to entitles.User
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


func (pg *PGGormDB) CreateTable(rq *entities.Table) (*entities.Table, error){
	if err := pg.db.Create(rq).Error; err != nil{
		return nil, err
	}

	return rq, nil
}


func (pg *PGGormDB) FindTableByID(id string) (*entities.Table, error){
	table := new(entities.Table)
	if err := pg.db.Where("id = ?", id).Find(table).Error; err != nil{
		return nil, err
	}


	return table, nil
}

func (pg *PGGormDB) UpdateTableByID(rq *entities.Table) (*entities.Table, error){
	table := new(entities.Table)
    if err := pg.db.First(table, "id = ?", rq.ID).Error; err != nil {
        return nil, err
    }

	table.C_ID = rq.C_ID
	table.Status = rq.Status

    if err := pg.db.Save(table).Error; err != nil {
        return nil, err
    }

    return table, nil
}


func (pg *PGGormDB) DeleteTableByID(id string) (error){
	table:= new(entities.Table)

	if err := pg.db.First(table, "id = ?", id).Error; err != nil{
		return err
	}


	if err := pg.db.Delete(table).Error; err != nil{
		return err
	}


	return nil
}