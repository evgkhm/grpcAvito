package usecase

import (
	"context"
	"fmt"
	"grpcAvito/internal/entity"
)

func (u UseCase) UserOrderRevenue(ctx context.Context, revenue *entity.UserRevenue) error {
	tx, err := u.txService.NewTransaction()
	if err != nil {
		errRollback := u.txService.Rollback(tx)
		if errRollback != nil {
			return fmt.Errorf("usecase - UseCase - UserOrderRevenue - u.txService.NewTransaction - u.txService.Rollback: %w", err)
		}
		return fmt.Errorf("usecase - UseCase - UserOrderRevenue - u.txService.NewTransaction: %w", err)
	}

	reservation := &entity.UserReservation{ID: revenue.ID,
		IDOrder:   revenue.IDOrder,
		IDService: revenue.IDService,
		Cost:      revenue.Cost,
	}

	// Списать с резервации
	err = u.repo.UserOrderDeleteReservation(ctx, tx, reservation)
	if err != nil {
		errRollback := u.txService.Rollback(tx)
		if errRollback != nil {
			return fmt.Errorf("usecase - UseCase - UserOrderRevenue - u.repo.UserOrderDeleteReservation - u.txService.Rollback: %w", err)
		}
		return fmt.Errorf("usecase - UseCase - UserOrderRevenue - u.repo.UserOrderDeleteReservation: %w", err)
	}

	// Начислить в revenue
	err = u.repo.UserOrderRevenue(ctx, tx, revenue)
	if err != nil {
		errRollback := u.txService.Rollback(tx)
		if errRollback != nil {
			return fmt.Errorf("usecase - UseCase - UserOrderRevenue - u.repo.UserOrderRevenue - u.txService.Rollback: %w", err)
		}
		return fmt.Errorf("usecase - UseCase - UserOrderRevenue - u.repo.UserOrderRevenue: %w", err)
	}

	return u.txService.Commit(tx)
}
