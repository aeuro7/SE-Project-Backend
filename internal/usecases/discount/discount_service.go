package discount

import (
	"errors"
	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"
	"github.com/B1gdawg0/se-project-backend/internal/transaction/response"
	"github.com/B1gdawg0/se-project-backend/internal/utils"
	"github.com/emicklei/pgtalk/convert"
	"github.com/jackc/pgx/v5/pgtype"
	"gorm.io/gorm"
)


type DiscountUseCase interface{
	FindAllDiscount()(*response.GetDiscountsResponse, error)
	FindDiscountByID(id pgtype.UUID) (*response.GetDiscountResponse, error)
	CreateDiscount( *entities.Discount) (*response.CreateDiscountResponse, error)
}

type DiscountService struct{
	repo DiscountRepository
}


func ProvideDiscountService(repo DiscountRepository) DiscountUseCase{
	return &DiscountService{
		repo: repo,
	}
}


func (ds *DiscountService) FindAllDiscount()(*response.GetDiscountsResponse, error){
	all, err := ds.repo.FindAllDiscount()

	if err != nil {
		return nil, err
	}

	list := &response.GetDiscountsResponse{
		Discounts: make([]response.GetDiscountResponse, 0),
	}

	for _, obj := range(all) {
		list.Discounts = append(list.Discounts, response.GetDiscountResponse{
			ID: convert.UUIDToString(obj.ID),
			C_ID: convert.UUIDToString(obj.C_ID),
			O_ID: convert.UUIDToString(obj.O_ID),
			Percent: obj.Percent,
			Name: obj.Name,
			Code: obj.Code,
			Description : obj.Description,
			StartDate: obj.StartDate.Format("2006-01-02 15:04:05"),
			ExpDate: obj.ExpDate.Format("2006-01-02 15:04:05"),
			Status: obj.Status,
		})
	}

	return list, nil
}


func (ds *DiscountService) FindDiscountByID(id pgtype.UUID) (*response.GetDiscountResponse, error){
	selected, err := ds.repo.FindDiscountByID(id)

	if err != nil{
		return nil, err
	}

	return &response.GetDiscountResponse{
		ID: convert.UUIDToString(selected.ID),
		C_ID: convert.UUIDToString(selected.C_ID),
		O_ID: convert.UUIDToString(selected.O_ID),
		Percent: selected.Percent,
		Name: selected.Name,
		Code: selected.Code,
		Description : selected.Description,
		StartDate: selected.StartDate.Format("2006-01-02 15:04:05"),
		ExpDate: selected.ExpDate.Format("2006-01-02 15:04:05"),
		Status: selected.Status,
	},nil
}

func (ds *DiscountService) CreateDiscount(rq *entities.Discount) (*response.CreateDiscountResponse, error){
	rq.ID = utils.GenerateUUID()

	selected, err := ds.repo.FindDiscountByID(rq.ID)

	if err != nil && err != gorm.ErrRecordNotFound{
		return nil, err
	}

	if selected != nil && selected.ID == rq.ID{
		return nil, errors.New("discount already exists")
	}
	
	discount, err := ds.repo.CreateDiscount(rq)

	if err != nil{
		return nil, err
	}

	return &response.CreateDiscountResponse{
		ID: convert.UUIDToString(discount.ID),
		C_ID: convert.UUIDToString(discount.C_ID),
		O_ID: convert.UUIDToString(discount.O_ID),
		Percent: discount.Percent,
		Name: discount.Name,
		Code: discount.Code,
		Description : discount.Description,
		StartDate: discount.StartDate.Format("2006-01-02 15:04:05"),
		ExpDate: discount.ExpDate.Format("2006-01-02 15:04:05"),
		Status: discount.Status,
	}, nil
}
