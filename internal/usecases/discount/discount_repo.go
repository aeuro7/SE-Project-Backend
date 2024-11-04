package discount

import (
	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"
	"github.com/jackc/pgx/v5/pgtype"
)


type DiscountRepository  interface{
	FindAllDiscount()([]*entities.Discount, error)
	FindDiscountByID(id pgtype.UUID) (*entities.Discount, error)
	CreateDiscount(rq *entities.Discount) (*entities.Discount, error)
}