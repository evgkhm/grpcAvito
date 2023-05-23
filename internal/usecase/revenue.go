package usecase

import (
	"context"
	"grpcAvito/internal/entity"
)

func (u UseCase) Revenue(ctx context.Context, revenue entity.UserRevenue) error {
	tx, err := u.txService.NewTransaction()
	if err != nil {
		u.txService.Rollback(tx)
		return err
	}
	defer u.txService.Commit(tx)

	return nil
}
