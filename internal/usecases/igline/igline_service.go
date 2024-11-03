package igline

import (
	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"
)

type IgLineUseCase interface {
	CreateIgLine(rq *entities.IGLine) (*entities.IGLine, error)
	FindAllIgLine() ([]*entities.IGLine, error)
}

type IgLineService struct {
	igLineRepo IgLineRepository
}

func ProvideIgLineService(repo IgLineRepository) IgLineUseCase {
	return &IgLineService{
		igLineRepo: repo,
	}
}

func (igs *IgLineService) CreateIgLine(rq *entities.IGLine) (*entities.IGLine, error) {
	igLine, err := igs.igLineRepo.CreateIgLine(rq)

	if err != nil {
		return nil, err
	}

	return igLine, nil
}

func (igs *IgLineService) FindAllIgLine() ([]*entities.IGLine, error) {
	var igLineLs []*entities.IGLine
	igLineLs, err := igs.igLineRepo.FindAllIgLine()

	if err != nil {
		return nil, err
	}

	return igLineLs, nil
}

