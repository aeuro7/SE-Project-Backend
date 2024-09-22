package handlers

import (
	"github.com/B1gdawg0/se-project-backend/internal/transaction/requests"
	"github.com/B1gdawg0/se-project-backend/internal/usecases/user"
	"github.com/gofiber/fiber/v2"
)

type UserRestHandler struct{
	usecase user.UserUseCase
}

func ProvideUserRestHandler(usecase user.UserUseCase) *UserRestHandler{
	return &UserRestHandler{
		usecase: usecase,
	}
}

func (urh *UserRestHandler) FindAll(c *fiber.Ctx) error{
	list, err := urh.usecase.FindAll()

	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"Internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message":"Successful get list of user",
		"payload":fiber.Map{
			"users": list,
		},
	})
}


func (urh *UserRestHandler) CreateUser(c *fiber.Ctx) error{
	rq := new(requests.CreateUserRequest)

	if err:= c.BodyParser(rq); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Invalid data",
		})
	}

	user, err := urh.usecase.CreateUser(rq)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"Internal server error",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":"Successful Created",
		"payload":fiber.Map{
			"user":user,
		},
	})
}