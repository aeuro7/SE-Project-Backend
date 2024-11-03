package handlers

import (
	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"
	"github.com/B1gdawg0/se-project-backend/internal/transaction/requests"
	"github.com/B1gdawg0/se-project-backend/internal/usecases/igline"
	"github.com/B1gdawg0/se-project-backend/internal/utils"
	"github.com/emicklei/pgtalk/convert"
	"github.com/gofiber/fiber/v2"
)

type IgLineHandler struct {
	usecase igline.IgLineUseCase
}

func ProvideIglineHandler(usecase igline.IgLineUseCase) *IgLineHandler {
	return &IgLineHandler{
		usecase: usecase,
	}
}

func (igh *IgLineHandler) CreateIgLine(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	var request requests.IglineRequest

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Can't Parse Request Body",
			"error":   err.Error(),
		})
	}

	payload := &entities.IGLine{
		ID:   utils.GenerateUUID(),
		C_ID: convert.StringToUUID(userID),
		Name: request.IgAccount,
	}

	res, err := igh.usecase.CreateIgLine(payload)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Can't Create IgLine",
			"error":   err,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Succesful create IGLines",
		"payload": fiber.Map{
			"Ig-Line": res,
		},
	})
}

func (igh *IgLineHandler) FindAllIgLine(c *fiber.Ctx) error {
	igLines, err := igh.usecase.FindAllIgLine()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successful get all iglines",
		"payload": fiber.Map{
			"iglines": igLines,
		},
	})
}
