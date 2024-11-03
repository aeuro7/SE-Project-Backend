package repositories

import (
	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"
	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/queries"
	"github.com/B1gdawg0/se-project-backend/internal/usecases/musicline"
	"gorm.io/gorm"
)

type MusicLineRepositoryImpl struct {
	Queries queries.Database
}

func ProvideMusicLineRepository(db *gorm.DB) musicline.MusicLineRepository {
	return &MusicLineRepositoryImpl{
		Queries: queries.New(db),
	}
}

func (msr *MusicLineRepositoryImpl) CreateMusicLine(rq *entities.MusicLine) (*entities.MusicLine, error) {
	return msr.Queries.CreateMusicLine(rq)
}

func (msr *MusicLineRepositoryImpl) FindAllMusicLine() ([]*entities.MusicLine, error) {
	return msr.Queries.FindAllMusicLine()
}
