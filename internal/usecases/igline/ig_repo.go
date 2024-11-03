package igline

import "github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"

type IgLineRepository interface {
	CreateIgLine(rq *entities.IGLine) (*entities.IGLine, error)
	FindAllIgLine() ([]*entities.IGLine, error)
}