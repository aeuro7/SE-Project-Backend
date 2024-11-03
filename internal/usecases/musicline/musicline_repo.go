package musicline

import "github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"

type MusicLineRepository interface {
	CreateMusicLine(rq *entities.MusicLine) (*entities.MusicLine, error)
	FindAllMusicLine() ([]*entities.MusicLine, error)
}
