package usecase

import (
	"context"
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
	Create(ctx context.Context, Id, Balance string) error
}

type UseCase struct {
	repo *postgres.RepositoriesPostgres
	log  *logrus.Logger
}

func (u UseCase) Create(ctx context.Context, Id, Balance string) error {
	//TODO: транзакция

	return nil
}

func NewUseCase(repo *postgres.RepositoriesPostgres, log *logrus.Logger) *UseCase {
	return &UseCase{
		repo: repo,
		log:  log,
	}
}
