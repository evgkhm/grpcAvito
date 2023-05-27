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

	reservation := entity.UserReservation{Id: revenue.Id,
		IdOrder:   revenue.IdOrder,
		IdService: revenue.IdService,
		Cost:      revenue.Cost,
	}

	//Списать с резервации
	err = u.repo.MinusReservation(ctx, reservation, tx)
	if err != nil {
		u.txService.Rollback(tx)
		return err
	}

	//Начислить в revenue
	err = u.repo.Revenue(ctx, revenue, tx)
	if err != nil {
		u.txService.Rollback(tx)
		return err
	}

	return u.txService.Commit(tx)
}
