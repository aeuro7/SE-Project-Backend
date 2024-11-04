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

func (pg *PGGormDB) FindUserByID(id pgtype.UUID) (*entities.User, error) {
	user := new(entities.User)

	if err := pg.db.Where("id = ?", id).Find(user); err.Error != nil {
		return nil, err.Error
	}

	return user, nil
}

// TODO: change return value to entitles.User
func (pg *PGGormDB) FindUserByEmail(email string) (*entities.User, error) {
	user := new(entities.User)

	if err := pg.db.Where("c_email = ?", email).Find(user); err.Error != nil {
		return nil, err.Error
	}

	return user, nil
}
func (pg *PGGormDB) FindUserByPhone(phone string) (*entities.User, error) {
	user := new(entities.User)

	if err := pg.db.Where("c_phone = ?", phone).Find(user); err.Error != nil {
		return nil, err.Error
	}

	return user, nil
}

// TODO: change return value to entitles.User
func (pg *PGGormDB) FindAll() ([]*entities.User, error) {
	users := new([]*entities.User)

	if err := pg.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return *users, nil
}

// TODO: change return value to entitles.User
func (pg *PGGormDB) CreateUser(rq *entities.User) (*response.CreateUserResponse, error) {

	if err := pg.db.Create(rq).Error; err != nil {
		return nil, err
	}

	response := &response.CreateUserResponse{
		ID:       rq.ID,
		Name:     rq.Name,
		Email:    rq.Email,
		Password: rq.Password,
		Phone:    rq.Phone,
	}
	return response, nil
}

func (pg *PGGormDB) FindAllTable() ([]*entities.Table, error) {
	var tables []*entities.Table

	if err := pg.db.Find(&tables).Error; err != nil {
		return nil, err
	}

	return tables, nil
}

func (pg *PGGormDB) CreateTable(rq *entities.Table) (*entities.Table, error) {
	if err := pg.db.Create(rq).Error; err != nil {
		return nil, err
	}

	return rq, nil
}

func (pg *PGGormDB) FindTableByID(id string) (*entities.Table, error) {
	table := new(entities.Table)
	if err := pg.db.Where("id = ?", id).Find(table).Error; err != nil {
		return nil, err
	}

	return table, nil
}

func (pg *PGGormDB) UpdateTableByID(rq *entities.Table) (*entities.Table, error) {
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

func (pg *PGGormDB) DeleteTableByID(id string) error {
	table := new(entities.Table)

	if err := pg.db.First(table, "id = ?", id).Error; err != nil {
		return err
	}

	if err := pg.db.Delete(table).Error; err != nil {
		return err
	}

	return nil
}

func (pg *PGGormDB) FindAllOrder() ([]*entities.Order, error) {
	orders := new([]*entities.Order)

	if err := pg.db.Find(orders).Error; err != nil {
		return nil, err
	}

	return *orders, nil
}

func (pg *PGGormDB) FindOrderByID(id pgtype.UUID) (*entities.Order, error) {
	order := new(entities.Order)
	if err := pg.db.First(order, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return order, nil
}

func (pg *PGGormDB) CreateOrderByID(rq *entities.Order) (*entities.Order, error) {
	if err := pg.db.Create(rq).Error; err != nil {
		return nil, err
	}

	return rq, nil
}

func (pg *PGGormDB) DeleteOrderByID(id pgtype.UUID) error {
	order := new(entities.Order)

	if err := pg.db.Where("id = ?", id).Find(order).Error; err != nil {
		return err
	}

	if err := pg.db.Delete(order).Error; err != nil {
		return err
	}

	return nil
}
func (pg *PGGormDB) FindAllOrderLine() ([]*entities.OrderLine, error) {
	olines := new([]*entities.OrderLine)

	if err := pg.db.Preload("Menu").Find(olines).Error; err != nil {
		return nil, err
	}

	return *olines, nil
}

func (pg *PGGormDB) FindOrderLineByID(id pgtype.UUID) (*entities.OrderLine, error) {
	oline := new(entities.OrderLine)
	if err := pg.db.Preload("Menu").First(oline, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return oline, nil
}

func (pg *PGGormDB) FindOrderLineByOrderID(id pgtype.UUID) ([]*entities.OrderLine, error) {
	olines := new([]*entities.OrderLine)

	if err := pg.db.Preload("Menu").Where("o_id = ?", id).Find(olines).Error; err != nil {
		return nil, err
	}

	return *olines, nil
}

func (pg *PGGormDB) CreateOrderLine(rq *entities.OrderLine) (*entities.OrderLine, error) {
	if err := pg.db.Create(rq).Error; err != nil {
		return nil, err
	}

	return rq, nil
}

func (pg *PGGormDB) CreateMenu(rq *entities.Menu) (*entities.Menu, error) {
	if err := pg.db.Create(rq).Error; err != nil {
		return nil, err
	}
	return rq, nil
}

func (pg *PGGormDB) FindMenuByID(id string) (*entities.Menu, error) {
	Menu := new(entities.Menu)
	if err := pg.db.Where("id = ?", id).Find(Menu).Error; err != nil {
		return nil, err
	}
	return Menu, nil
}

func (pg *PGGormDB) FindAllMenu() ([]*entities.Menu, error) {
	var menuList []*entities.Menu

	if err := pg.db.Find(&menuList).Error; err != nil {
		return nil, err
	}

	return menuList, nil
}

func (pg *PGGormDB) UpdateMenu(rq *entities.Menu) (*entities.Menu, error) {
	menu := new(entities.Menu)

	if err := pg.db.First(menu, "id = ?", rq.ID).Error; err != nil {
		return nil, err
	}
	menu.ID = rq.ID
	menu.Description = rq.Description
	menu.Price = rq.Price
	menu.Url = rq.Url

	if err := pg.db.Save(menu).Error; err != nil {
		return nil, err
	}
	return menu, nil
}

func (pg *PGGormDB) DeleteMenu(id string) error {
	menu := new(entities.Menu)
	if err := pg.db.First(menu, "id = ?", id).Error; err != nil {
		return err
	}

	if err := pg.db.Delete(menu).Error; err != nil {
		return err
	}

	return nil
}

func (pg *PGGormDB) DeleteOrderLineByID(id pgtype.UUID) error {
    if err := pg.db.Where("id = ?", id).Delete(&entities.OrderLine{}).Error; err != nil {
        return err
    }

    order := new(entities.Order)
    if err := pg.db.Where("id = ?", id).Find(order).Error; err != nil {
        return err
    }

    if err := pg.db.Delete(order).Error; err != nil {
        return err
    }

    return nil
}

func (pg *PGGormDB) CreateIgLine(rq *entities.IGLine) (*entities.IGLine, error) {

	if err := pg.db.Create(rq); err.Error != nil {
		return nil, err.Error
	}

	return rq, nil
}

func (pg *PGGormDB) FindAllIgLine() ([]*entities.IGLine, error) {
	var igLineLs []*entities.IGLine

	if err := pg.db.Find(&igLineLs); err.Error != nil {
		return nil, err.Error
	}

	return igLineLs, nil
}

func (pg *PGGormDB) CreateMusicLine(rq *entities.MusicLine) (*entities.MusicLine, error) {
	if err := pg.db.Create(rq).Error; err != nil {
		return nil, err
	}

	return rq, nil
}

func (pg *PGGormDB) FindAllMusicLine() ([]*entities.MusicLine, error) {
	var musicLines []*entities.MusicLine

	if err := pg.db.Find(&musicLines).Error; err != nil {
		return nil, err
	}

	return musicLines, nil
}


func (pg *PGGormDB) CreateDiscount(rq *entities.Discount) (*entities.Discount, error) {
	if err := pg.db.Create(rq).Error; err != nil {
		return nil, err
	}
	return rq, nil
}

func (pg *PGGormDB) FindDiscountByID(id pgtype.UUID) (*entities.Discount, error) {
	Discount := new(entities.Discount)
	if err := pg.db.Where("id = ?", id).Find(Discount).Error; err != nil {
		return nil, err
	}
	return Discount, nil
}

func (pg *PGGormDB) FindAllDiscount() ([]*entities.Discount, error) {
	var Discounts []*entities.Discount

	if err := pg.db.Find(&Discounts).Error; err != nil {
		return nil, err
	}

	return Discounts, nil
}
