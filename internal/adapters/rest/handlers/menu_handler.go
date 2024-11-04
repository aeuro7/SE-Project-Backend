package handlers

import (
	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"
	"github.com/B1gdawg0/se-project-backend/internal/transaction/requests"
	"github.com/B1gdawg0/se-project-backend/internal/usecases/menu"
	"github.com/B1gdawg0/se-project-backend/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type MenuRestHandler struct {
	usecase menu.MenuUseCase
}

func ProvideMenuRestHandler(usecase menu.MenuUseCase) *MenuRestHandler {
	return &MenuRestHandler{
		usecase: usecase,
	}
}

func (mrh *MenuRestHandler) GetMenuByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
			"error":   "Must pass id to get the menu",
		})
	}
	
	if (!utils.CheckUUID(id)) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
			"error":   "Invalid UUID",
		})
	}

	menu, err := mrh.usecase.FindMenuByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error fetching menu",
			"error":   err.Error(),
		})
	}

	if menu == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Menu not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successful get menu",
		"payload": fiber.Map{
			"menu": menu,
		},
	})
}

func (mrh *MenuRestHandler) GetAllMenu(c *fiber.Ctx) error {

	allMenus, err := mrh.usecase.FindAllMenu()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error fetching all menus",
			"error":   err.Error(),
		})
	}

	if allMenus == nil || len(allMenus.Menu) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No Menu yet",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully retrieved all menus",
		"payload": fiber.Map{
			"menus": allMenus.Menu,
		},
	})
}

func (mrh *MenuRestHandler) CreateMenu(c *fiber.Ctx) error {
	rq := new(requests.CreateMenuRequest)
	if err := c.BodyParser(rq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
			"error":   err.Error(),
		})
	}
	menuID := utils.GenerateUUID()

	payload := &entities.Menu{
		ID:          menuID,
		Price:       rq.Price,
		Description: rq.Description,
		Url:         rq.Url,
	}
	response, err := mrh.usecase.CreateMenu(payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creating menu",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successful Create Menu",
		"payload": fiber.Map{
			"menu": response,
		},
	})
}

func (mrh *MenuRestHandler) UpdateMenuByID(c *fiber.Ctx) error {
	id := c.Params("id")
	check := utils.CheckUUID(id);
	if (!check) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
			"error":   "ID must be UUID",
		})
	}
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
			"error":   "Must pass id before updating data",
		})
	}

	req := new(requests.UpdateMenuRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
			"error":   err.Error(),
		})
	}

	menuID, err := utils.StringToUUID(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
			"error":   "Invalid UUID format",
		})
	}

	payload := &entities.Menu{
		ID:          *menuID,
		Price:       req.Price,
		Description: req.Description,
		Url:         req.Url,
	}

	response, err := mrh.usecase.UpdateMenu(payload)
	if err != nil {
		print("Handler")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error updating menu",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successful update menu",
		"payload": fiber.Map{
			"menu": response,
		},
	})
}

func (mrh *MenuRestHandler) DeleteMenuByID(c *fiber.Ctx) error {
	id := c.Params("id")

	check := utils.CheckUUID(id);

	if (!check) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
			"error":   "Invalid UUID",
		})
	}

	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
			"error":   "Must pass id before updating data",
		})
	}
	if err := mrh.usecase.DeleteMenu(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Deletion Successful",
	})
}
