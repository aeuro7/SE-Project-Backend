package repositories

import (
	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"
	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/queries"
	"github.com/B1gdawg0/se-project-backend/internal/usecases/igline"
	"gorm.io/gorm"
)

type IgLineRepositoryImpl struct {
	Queries queries.Database
}

func ProvideIgLineRepository(db *gorm.DB) igline.IgLineRepository {
	return &IgLineRepositoryImpl{
		Queries: queries.New(db),
	}
}

func (igr *IgLineRepositoryImpl) CreateIgLine(rq *entities.IGLine) (*entities.IGLine, error) {
	return igr.Queries.CreateIgLine(rq)
}


func (igr *IgLineRepositoryImpl) FindAllIgLine() ([]*entities.IGLine, error) {
	return igr.Queries.FindAllIgLine()
}
