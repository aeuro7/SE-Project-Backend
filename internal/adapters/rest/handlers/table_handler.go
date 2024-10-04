package handlers

import (
	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"
	"github.com/B1gdawg0/se-project-backend/internal/transaction/requests"
	"github.com/B1gdawg0/se-project-backend/internal/usecases/table"
	"github.com/B1gdawg0/se-project-backend/internal/utils"
	"github.com/gofiber/fiber/v2"
)

const (
	Available string = "A"
	Reserved    string = "R"
    Occupied    string = "O"
)

type TableRestHandler struct{
	usecase table.TableUseCase
}


func ProvideTableRestHandler(usecase table.TableUseCase) *TableRestHandler{
	return &TableRestHandler{
		usecase: usecase,
	}
}

func (trh *TableRestHandler) GetTables(c *fiber.Ctx) error{
	tables, err :=trh.usecase.FindAllTable()

	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"Internal server error",
			"error":err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message":"Successful get all table",
		"payload":fiber.Map{
			"tables":tables,
		},
	})
}

func (trh *TableRestHandler) GetTableByID(c *fiber.Ctx) error{
	id := c.Params("id")

	if id == ""{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Bad Request",
			"error":"Must pass table's id before query data",
		})
	}

	table, err := trh.usecase.FindTableByID(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"Bad Request",
			"error":err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message":"Succesful find table",
		"payload":fiber.Map{
			"table":table,
		},
	})
}


func (trh *TableRestHandler) CreateTable(c *fiber.Ctx) error{
	rq := new(requests.CreateTableRequest)

	if err := c.BodyParser(rq); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Bad Request",
			"error":err.Error(),
		})
	}

	payload := &entities.Table{
		ID: rq.ID,
		Status: Available,
	}


	response, err := trh.usecase.CreateTable(payload)

	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"Bad Request",
			"error":err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message":"Succesful create table",
		"payload":fiber.Map{
			"table":response,
		},
	})
}

func (trh *TableRestHandler) UpdateTableByID(c *fiber.Ctx) error{
	id := c.Params("id")

	if id == ""{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Bad Request",
			"error":"Must pass id before query data",
		})
	}

	req := new(requests.UpdateTableRequest)

	if err := c.BodyParser(req); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Bad Request",
			"error":err.Error(),
		})
	}

	c_id, err := utils.StringToUUID(req.C_ID)

	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Bad Request",
			"error":"Customer id must be UUID",
		})
	}

	payload := &entities.Table{
		ID: id,
		C_ID: *c_id,
		Status: req.Status,
	}


	response, err := trh.usecase.UpdateTableByID(payload)

	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"Internal server error",
			"error":err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message":"Succesful update table",
		"payload":fiber.Map{
			"table":response,
		},
	})
}


func (trh *TableRestHandler) DeleteTableByID(c *fiber.Ctx) error{
	id := c.Params("id")

	if id == ""{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Bad Request",
			"error":"Must pass id before query data",
		})
	}

	if err := trh.usecase.DeleteTableByID(id);err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"Internal server error",
			"error":err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}