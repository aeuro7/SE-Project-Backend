package order

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"
	"github.com/B1gdawg0/se-project-backend/internal/transaction/requests"
	"github.com/B1gdawg0/se-project-backend/internal/transaction/response"
	"github.com/B1gdawg0/se-project-backend/internal/usecases/orderline"
	"github.com/B1gdawg0/se-project-backend/internal/utils"
	"github.com/emicklei/pgtalk/convert"
	"github.com/jackc/pgx/v5/pgtype"
	"gorm.io/gorm"
)

type OrderUseCase interface {
	FindAllOrder() (*response.GetOrdersResponse, error)
	FindOrderByID(id pgtype.UUID) (*response.GetOrderResponse, error)
	FindOrderByTableID(id string) (*response.GetOrdersResponse, error)
	CreateOrderByID(*entities.Order) (*response.CreateOrderResponse, error)
	CreateOrderWithOrderLines(*requests.CreateOrderWithOrderLinesRequest) (*response.CreateOrderResponse, error)
}

type OrderService struct {
	repo      OrderReposity
	olineRepo orderline.OrderLineRepository
}

func ProvideOrderService(repo OrderReposity, olineRepo orderline.OrderLineRepository) OrderUseCase {
	return &OrderService{
		repo:      repo,
		olineRepo: olineRepo,
	}
}

func (os *OrderService) FindAllOrder() (*response.GetOrdersResponse, error) {
	orders, err := os.repo.FindAllOrder()

	if err != nil {
		return nil, err
	}

	list := &response.GetOrdersResponse{
		Orders: make([]response.GetOrderResponse, 0),
	}

	for _, obj := range orders {
		list2 := make([]response.CreateOrderLineResponse, 0)

		for _, temp := range obj.OrderLines {
			list2 = append(list2, response.CreateOrderLineResponse{
				ID:       convert.UUIDToString(temp.ID),
				Time:     temp.Time.Format("2006-01-02 15:04:05"),
				O_ID:     convert.UUIDToString(temp.O_ID),
				M_ID:     convert.UUIDToString(temp.M_ID),
				Quantity: fmt.Sprintf("%d", temp.Quantity),
				Price:    fmt.Sprintf("%.2f", temp.Price),
				Menu: response.GetMenuResponse{
					ID:          temp.Menu.ID,
					Description: temp.Menu.Description,
					Price:       temp.Menu.Price,
					Url:         temp.Menu.Url,
				},
			})
		}

		list.Orders = append(list.Orders, response.GetOrderResponse{
			ID:        convert.UUIDToString(obj.ID),
			T_ID:      obj.T_ID,
			Time:      obj.Time.Format("2006-01-02 15:04:05"),
			Url:       obj.Url,
			Orderline: list2,
		})
	}

	return list, nil
}

func (os *OrderService) FindOrderByID(id pgtype.UUID) (*response.GetOrderResponse, error) {
	selected, err := os.repo.FindOrderByID(id)

	if err != nil {
		return nil, err
	}

	list := make([]response.CreateOrderLineResponse, 0)

	for _, obj := range selected.OrderLines {
		list = append(list, response.CreateOrderLineResponse{
			ID:       convert.UUIDToString(obj.ID),
			Time:     obj.Time.Format("2006-01-02 15:04:05"),
			O_ID:     convert.UUIDToString(obj.O_ID),
			M_ID:     convert.UUIDToString(obj.M_ID),
			Quantity: fmt.Sprintf("%d", obj.Quantity),
			Price:    fmt.Sprintf("%.2f", obj.Price),
			Menu: response.GetMenuResponse{
				ID:          obj.Menu.ID,
				Description: obj.Menu.Description,
				Price:       obj.Menu.Price,
				Url:         obj.Menu.Url,
			},
		})
	}

	return &response.GetOrderResponse{
		ID:        convert.UUIDToString(selected.ID),
		T_ID:      selected.T_ID,
		Time:      selected.Time.Format("2006-01-02 15:04:05"),
		Url:       selected.Url,
		Orderline: list,
	}, nil
}

func (os *OrderService) FindOrderByTableID(id string) (*response.GetOrdersResponse, error) {
	selected, err := os.repo.FindOrderByTableID(id)
	if err != nil {
		return nil, err
	}

	ordersResponse := &response.GetOrdersResponse{
		Orders: make([]response.GetOrderResponse, 0),
	}

	for _, obj := range selected {
		orderLines := make([]response.CreateOrderLineResponse, 0)

		for _, line := range obj.OrderLines {
			orderLines = append(orderLines, response.CreateOrderLineResponse{
				ID:       convert.UUIDToString(line.ID),
				Time:     line.Time.Format("2006-01-02 15:04:05"),
				O_ID:     convert.UUIDToString(line.O_ID),
				M_ID:     convert.UUIDToString(line.M_ID),
				Quantity: fmt.Sprintf("%d", line.Quantity),
				Price:    fmt.Sprintf("%.2f", line.Price),
				Menu: response.GetMenuResponse{
					ID:          line.Menu.ID,
					Description: line.Menu.Description,
					Price:       line.Menu.Price,
					Url:         line.Menu.Url,
				},
			})
		}

		ordersResponse.Orders = append(ordersResponse.Orders, response.GetOrderResponse{
			ID:        convert.UUIDToString(obj.ID),
			T_ID:      obj.T_ID,
			Time:      obj.Time.Format("2006-01-02 15:04:05"),
			Url:       obj.Url,
			Orderline: orderLines,
		})
	}

	return ordersResponse, nil
}



func (os *OrderService) CreateOrderByID(rq *entities.Order) (*response.CreateOrderResponse, error) {
	rq.ID = utils.GenerateUUID()

	selected, err := os.repo.FindOrderByID(rq.ID)

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if selected != nil && selected.ID == rq.ID {
		return nil, errors.New("order already exists")
	}

	rq.Time = time.Now()
	order, err := os.repo.CreateOrderByID(rq)

	if err != nil {
		return nil, err
	}

	return &response.CreateOrderResponse{
		ID:   convert.UUIDToString(order.ID),
		T_ID: order.T_ID,
		Time: order.Time.Format("2006-01-02 15:04:05"),
		Url:  selected.Url,
	}, nil
}

func (os *OrderService) CreateOrderWithOrderLines(rq *requests.CreateOrderWithOrderLinesRequest) (*response.CreateOrderResponse, error) {
	order := &entities.Order{
		ID:   utils.GenerateUUID(),
		T_ID: rq.T_ID,
		Time: time.Now(),
		Url:  rq.Url,
	}

	created, err := os.repo.CreateOrderByID(order)

	if err != nil {
		return nil, err
	}

	for _, obj := range rq.OrderLines {
		m_id, err := utils.StringToUUID(obj.M_ID)
		if err != nil {
			return nil, err
		}

		quantity, err := strconv.Atoi(obj.Quantity)
		if err != nil {
			return nil, err
		}

		price, err := strconv.ParseFloat(obj.Price, 32)
		if err != nil {
			return nil, err
		}

		temp := &entities.OrderLine{
			ID:       utils.GenerateUUID(),
			Time:     time.Now(),
			O_ID:     created.ID,
			M_ID:     *m_id,
			Quantity: quantity,
			Price:    float32(price),
		}

		_, err = os.olineRepo.CreateOrderLine(temp)
		if err != nil {
			return nil, err
		}
	}

	return &response.CreateOrderResponse{
		ID:   convert.UUIDToString(created.ID),
		T_ID: created.T_ID,
		Time: created.Time.Format("2006-01-02 15:04:05"),
		Url:  created.Url,
	}, nil
}
