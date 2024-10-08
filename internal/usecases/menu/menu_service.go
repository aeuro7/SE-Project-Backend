package menu

import (
	"errors"

	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"
	"github.com/B1gdawg0/se-project-backend/internal/transaction/response"
	"github.com/B1gdawg0/se-project-backend/internal/utils"
	"github.com/emicklei/pgtalk/convert"
)

type MenuUseCase interface {
	CreateMenu(*entities.Menu) (*response.CreateMenuResponse, error)
	FindMenuByID(string) (*response.GetMenuResponse, error)
	FindAllMenu() (*response.GetAllMenuResponse, error)
	UpdateMenu(*entities.Menu) (*response.UpdateMenuResponse, error)
	DeleteMenu(string) error
}

type MenuService struct {
	repo MenuRepository
}

func ProvideMenuService(repo MenuRepository) MenuUseCase {
	return &MenuService{
		repo: repo,
	}
}

func (ms *MenuService) CreateMenu(rq *entities.Menu) (*response.CreateMenuResponse, error) {
	menuID := convert.UUIDToString(rq.ID)
	check, err := ms.repo.FindMenuByID(menuID)
	if check != nil && convert.UUIDToString(check.ID) != "" {
		return nil, errors.New("menu already exist")
	}
	if err != nil {
		return nil, err
	}

	menu, err := ms.repo.CreateMenu(rq)
	if err != nil {
		return nil, err
	}

	return &response.CreateMenuResponse{
		ID:          menu.ID,
		Price:       menu.Price,
		Description: menu.Description,
		Url:         menu.Url,
	}, nil
}

func (ms *MenuService) FindMenuByID(id string) (*response.GetMenuResponse, error) {

	if id == "" {
		return nil, errors.New("menu ID is required")
	}

	menu, err := ms.repo.FindMenuByID(id)
	if err != nil {
		return nil, err
	}
	if menu == nil {
		return nil, errors.New("menu not found")
	}

	return &response.GetMenuResponse{
		ID:          menu.ID,
		Price:       menu.Price,
		Description: menu.Description,
		Url:         menu.Url,
	}, nil
}

func (ms *MenuService) FindAllMenu() (*response.GetAllMenuResponse, error) {
	menus, err := ms.repo.FindAllMenu()
	if err != nil {
		return nil, err
	}

	var menuResponses []response.GetMenuResponse
	for _, menu := range menus {
		menuResponses = append(menuResponses, response.GetMenuResponse{
			ID:          menu.ID,
			Price:       menu.Price,
			Description: menu.Description,
			Url:         menu.Url,
		})
	}

	return &response.GetAllMenuResponse{
		Menu: menuResponses,
	}, nil
}

func (ms *MenuService) UpdateMenu(rq *entities.Menu) (*response.UpdateMenuResponse, error) {

	menuID := convert.UUIDToString(rq.ID)
	check := utils.CheckUUID(menuID)
	print("Service")
	if !check {
		print("Invalid")
		return nil, errors.New("invalid UUID")
	}

	_, err := ms.FindMenuByID(menuID)
	if err != nil {
		return nil, errors.New("menu not found")
	}

	updatedMenu, err := ms.repo.UpdateMenu(rq)
	if err != nil {
		return nil, err
	}

	return &response.UpdateMenuResponse{
		ID:          updatedMenu.ID,
		Price:       updatedMenu.Price,
		Description: updatedMenu.Description,
		Url:         updatedMenu.Url,
	}, nil
}

func (ms *MenuService) DeleteMenu(id string) error {
	if id == "" {
		return errors.New("menu ID is required")
	}
	err := ms.repo.DeleteMenu(id)
	if err != nil {
		return err
	}

	return nil
}
