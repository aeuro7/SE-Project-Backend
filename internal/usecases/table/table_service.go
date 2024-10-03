package table

import (
	"errors"

	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"
	"github.com/B1gdawg0/se-project-backend/internal/transaction/response"
	"github.com/emicklei/pgtalk/convert"
)

type TableUseCase interface {
	FindAllTable() (*response.GetTablesResponse, error)
	FindTableByID(string) (*response.GetTableResponse, error)
	CreateTable(*entities.Table) (*response.CreateTableResponse, error)
	UpdateTableByID(*entities.Table) (*response.UpdateTableResponse, error)
	DeleteTableByID(string) (error)
}

type TableService struct {
	repo TableRepository
}

func ProvideTableService(repo TableRepository) TableUseCase {
	return &TableService{
		repo: repo,
	}
}

func (ts *TableService) FindTableByID(id string) (*response.GetTableResponse, error){
	table, err := ts.repo.FindTableByID(id)

	if err != nil {
        return nil, errors.New("failed to retrieve table: " + err.Error())
    }

    if table.ID == "" {
        return nil, errors.New("table not found")
    }

	C_IDStr := convert.UUIDToString(table.C_ID)
	if C_IDStr == ""{
		C_IDStr = "Nobody booked this table"
	}


	return &response.GetTableResponse{
		ID: table.ID,
		C_ID: C_IDStr,
		Status: table.Status,
	}, nil
}

func (ts *TableService) FindAllTable() (*response.GetTablesResponse, error){
	list, err := ts.repo.FindAllTable()

	if err != nil{
		return nil, err
	}

	res := &response.GetTablesResponse{
		Tables: make([]response.GetTableResponse, 0),
	}

	for _, obj := range(list) {
		C_IDStr := convert.UUIDToString(obj.C_ID)

		if C_IDStr == ""{
		C_IDStr = "Nobody book this table"
		}

		res.Tables = append(res.Tables, response.GetTableResponse{
			ID: obj.ID,
			C_ID: C_IDStr,
			Status: obj.Status,
		})
	}

	return res, nil
}

func (ts *TableService) CreateTable(rq *entities.Table) (*response.CreateTableResponse, error) {
	selected, err := ts.repo.FindTableByID(rq.ID)

	if err != nil{
		return nil, err
	}

	if selected != nil && selected.ID != ""{
		return nil, errors.New("table with this ID already exists")
	}

	table, err := ts.repo.CreateTable(rq)


	if err != nil{
		return nil ,err
	}

	C_IDStr := convert.UUIDToString(table.C_ID)
	if C_IDStr == ""{
		C_IDStr = "Nobody book this table"
	}

	return &response.CreateTableResponse{
		ID: table.ID,
		C_ID: C_IDStr,
		Status: table.Status,
	}, nil
}

func (ts *TableService) UpdateTableByID(rq *entities.Table) (*response.UpdateTableResponse, error){
	exist, err := ts.repo.FindTableByID(rq.ID)

	if err != nil{
		return nil, err
	}

	if exist == nil{
		return nil, errors.New("table not found")
	}

	table, err := ts.repo.UpdateTableByID(rq)

	if err != nil{
		return nil, err
	}


	C_IDStr := convert.UUIDToString(table.C_ID)
	if C_IDStr == ""{
		C_IDStr = "Nobody book this table"
	}

	return &response.UpdateTableResponse{
		ID: table.ID,
		C_ID: C_IDStr,
		Status: table.Status,
	}, nil
}


func (ts *TableService) DeleteTableByID(id string) (error){
	 exist, err := ts.repo.FindTableByID(id)

	 if err != nil{
		return err
	 }

	 if exist == nil {
		return errors.New("table not found")
	 }

	 return ts.repo.DeleteTableByID(id)
}
