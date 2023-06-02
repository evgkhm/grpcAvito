package usecase

import (
	"context"
	"fmt"
	"grpcAvito/internal/entity"
)

func (u UseCase) Revenue(ctx context.Context, revenue *entity.UserRevenue) error {
	tx, err := u.txService.NewTransaction()
	if err != nil {
		_ = u.txService.Rollback(tx)
		return fmt.Errorf("usecase - UseCase - Revenue - u.txService.NewTransaction: %w", err)
	}

	reservation := &entity.UserReservation{Id: revenue.Id,
		IdOrder:   revenue.IdOrder,
		IdService: revenue.IdService,
		Cost:      revenue.Cost,
	}

	//Списать с резервации
	err = u.repo.DeleteReservation(ctx, tx, reservation)
	if err != nil {
		_ = u.txService.Rollback(tx)
		return fmt.Errorf("usecase - UseCase - Revenue - u.repo.DeleteReservation: %w", err)
	}

	//Начислить в revenue
	err = u.repo.Revenue(ctx, tx, revenue)
	if err != nil {
		_ = u.txService.Rollback(tx)
		return fmt.Errorf("usecase - UseCase - Revenue - u.repo.Revenue: %w", err)
	}

	return u.txService.Commit(tx)
}
