package handlers

import (
	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"
	"github.com/B1gdawg0/se-project-backend/internal/transaction/requests"
	"github.com/B1gdawg0/se-project-backend/internal/usecases/musicline"
	"github.com/B1gdawg0/se-project-backend/internal/utils"
	"github.com/emicklei/pgtalk/convert"
	"github.com/gofiber/fiber/v2"
)

type MusicLineHandler struct {
	usecase musicline.MusicLineUseCase
}

func ProvideMusicLineHandler(usecase musicline.MusicLineUseCase) *MusicLineHandler {
	return &MusicLineHandler{usecase: usecase}
}

func (mlh *MusicLineHandler) CreateMusicLine(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	var rq requests.MusicLineRequest

	if err := c.BodyParser(&rq); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	payload := &entities.MusicLine{
		ID:   utils.GenerateUUID(),
		C_ID: convert.StringToUUID(userID),
		Name: rq.MusicName,
	}

	res, err := mlh.usecase.CreateMusicLine(payload)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Succesful create MusicLine",
		"payload": fiber.Map{
			"MusicLine": res,
		},
	})

}

func (mlh *MusicLineHandler) FindAllMusicLine(c *fiber.Ctx) error {
	musicLines, err := mlh.usecase.FindAllMusicLine()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successful get all MusicLine",
		"payload": fiber.Map{
			"MusicLines": musicLines,
		},
	})
}
