package repositories

import (
	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"
	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/queries"
	"github.com/B1gdawg0/se-project-backend/internal/usecases/discount"
	"github.com/jackc/pgx/v5/pgtype"
	"gorm.io/gorm"
)

type DiscountRepositoryIPML struct {
	Queries queries.Database
}

func ProvideDiscountRepository(db *gorm.DB) discount.DiscountRepository {
	return &DiscountRepositoryIPML{
		Queries: queries.New(db),
	}
}

func (dri *DiscountRepositoryIPML) FindAllDiscount() ([]*entities.Discount, error) {
	return dri.Queries.FindAllDiscount()
}

func (dri *DiscountRepositoryIPML) FindDiscountByID(id pgtype.UUID) (*entities.Discount, error) {
	return dri.Queries.FindDiscountByID(id)
}

func (dri *DiscountRepositoryIPML) CreateDiscount(rq *entities.Discount) (*entities.Discount, error) {
	return dri.Queries.CreateDiscount(rq)
}

