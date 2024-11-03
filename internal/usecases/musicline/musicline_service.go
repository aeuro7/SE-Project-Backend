package musicline

import "github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"

type MusicLineUseCase interface {
	CreateMusicLine(rq *entities.MusicLine) (*entities.MusicLine, error)
	FindAllMusicLine() ([]*entities.MusicLine, error)
}

type MusicService struct {
	musicRepo MusicLineRepository
}

func ProvideMusicService(repo MusicLineRepository) MusicLineUseCase {
	return &MusicService{
		musicRepo: repo,
	}
}

func (ms *MusicService) CreateMusicLine(rq *entities.MusicLine) (*entities.MusicLine, error) {
	musicLine, err := ms.musicRepo.CreateMusicLine(rq)
	if err != nil {
		return nil, err
	}

	return musicLine, nil
}

func (ms *MusicService) FindAllMusicLine() ([]*entities.MusicLine, error) {
	var musicLineLs []*entities.MusicLine

	musicLineLs, err := ms.musicRepo.FindAllMusicLine()
	if err != nil {
		return nil, err
	}

	return musicLineLs, nil
}
