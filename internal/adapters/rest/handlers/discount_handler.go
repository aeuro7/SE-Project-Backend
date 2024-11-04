package handlers

import (
	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"
	"github.com/B1gdawg0/se-project-backend/internal/transaction/requests"
	"github.com/B1gdawg0/se-project-backend/internal/usecases/discount"
	"github.com/B1gdawg0/se-project-backend/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type DiscountRestHandler struct {
	usecase discount.DiscountUseCase
}

func ProvideDiscountRestHandler(usecase discount.DiscountUseCase) *DiscountRestHandler {
	return &DiscountRestHandler{
		usecase: usecase,
	}
}

func (drh *DiscountRestHandler) GetDiscountByID(c *fiber.Ctx) error {
	discountID := c.Params("id")
	uuid, err := utils.StringToUUID(discountID)

	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Bad Request",
			"error":"Invalid ID format",
		})
	}
	
	response, err := drh.usecase.FindDiscountByID(*uuid)

	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"Internal server error",
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message":"Successfully retrieved discount",
		"payload":fiber.Map{
			"order":response,
		},
	})
}

func (drh *DiscountRestHandler) GetAllDiscount(c *fiber.Ctx) error {
	allDiscounts, err := drh.usecase.FindAllDiscount()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error fetching all discounts",
			"error":   err.Error(),
		})
	}

	if allDiscounts == nil || len(allDiscounts.Discounts) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No Discount yet",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully retrieved all discounts",
		"payload": fiber.Map{
			"discounts": allDiscounts.Discounts,
		},
	})
}

func (drh *DiscountRestHandler) CreateDiscount(c *fiber.Ctx) error {
	rq := new(requests.CreateDiscountRequest)
	if err := c.BodyParser(rq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
			"error":   err.Error(),
		})
	}
	discountID := utils.GenerateUUID()

	cID, err := utils.StringToUUID(rq.C_ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
			"error":   "Invalid C_ID format",
		})
	}

	oID, err := utils.StringToUUID(rq.O_ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
			"error":   "Invalid O_ID format",
		})
	}


	payload := &entities.Discount{
		ID: discountID,
		C_ID: *cID,
		O_ID: *oID,
		Percent: rq.Percent,
		Name: rq.Name,
		Code: rq.Code,
		Description: rq.Description,
		StartDate: rq.StartDate,
		ExpDate: rq.ExpDate,
		Status: rq.Status,
	}
	response, err := drh.usecase.CreateDiscount(payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creating discount",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successful Create Discount",
		"payload": fiber.Map{
			"discount": response,
		},
	})
}

