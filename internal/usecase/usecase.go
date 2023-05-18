package usecase

import (
	"github.com/sirupsen/logrus"
	"grpcAvito/internal/repository/postgres"
)

type ServerUseCase interface {
	/*AddTrack(ctx context.Context, playlistId, trackId string) (bool, error)
	AddAlbum(ctx context.Context, playlistId, albumId string) (bool, error)
	DeleteTrack(ctx context.Context, playlistId, trackId string) (bool, error)
	GetTrack(ctx context.Context, playlistId, trackId string) (domain.Track, error)
	GetAllTracks(ctx context.Context, playlistId string) ([]domain.Track, error)*/
	//TODO: методы реализовать
}

type UseCase struct {
	ServerUseCase
}

func NewUseCase(repo *postgres.RepositoriesPostgres, log *logrus.Logger) *UseCase {
	return &UseCase{
		ServerUseCase: NewServerUseCase(repo, log),
	}
}
