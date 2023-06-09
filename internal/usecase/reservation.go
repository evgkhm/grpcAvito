package usecase

import (
	"context"
	"fmt"
	"grpcAvito/internal/entity"
)

func (u UseCase) UserOrderDeleteReservation(ctx context.Context, reservation *entity.UserReservation) error {
	tx, err := u.txService.NewTransaction()
	if err != nil {
		errRollback := u.txService.Rollback(tx)
		if errRollback != nil {
			return fmt.Errorf("usecase - UseCase - UserOrderDeleteReservation - u.txService.NewTransaction - u.txService.Rollback: %w", err)
		}
		return fmt.Errorf("usecase - UseCase - UserOrderDeleteReservation - u.txService.NewTransaction: %w", err)
	}

	// Убрать резерв
	err = u.repo.UserOrderDeleteReservation(ctx, tx, reservation)
	if err != nil {
		errRollback := u.txService.Rollback(tx)
		if errRollback != nil {
			return fmt.Errorf("usecase - UseCase - UserOrderDeleteReservation - u.repo.UserOrderDeleteReservation - u.txService.Rollback: %w", err)
		}
		return fmt.Errorf("usecase - UseCase - UserOrderDeleteReservation - u.repo.UserOrderDeleteReservation: %w", err)
	}

	// Узнать текущий баланс
	userDTO := entity.User{ID: reservation.ID}
	currBalance, err := u.repo.GetBalance(ctx, tx, &userDTO)
	if err != nil {
		errRollback := u.txService.Rollback(tx)
		if errRollback != nil {
			return fmt.Errorf("usecase - UseCase - UserOrderDeleteReservation - u.repo.GetBalance - u.txService.Rollback: %w", err)
		}
		return fmt.Errorf("usecase - UseCase - UserOrderDeleteReservation - u.repo.GetBalance: %w", err)
	}

	// Начисление баланса
	newBalance := currBalance + reservation.Cost
	userDTO.Balance = newBalance
	err = u.repo.UserBalanceAccrual(ctx, tx, &userDTO)
	if err != nil {
		errRollback := u.txService.Rollback(tx)
		if errRollback != nil {
			return fmt.Errorf("usecase - UseCase - UserOrderDeleteReservation - u.repo.UserBalanceAccrual - u.txService.Rollback: %w", err)
		}
		return fmt.Errorf("usecase - UseCase - UserOrderDeleteReservation - u.repo.UserBalanceAccrual: %w", err)
	}

	return u.txService.Commit(tx)
}

func (u UseCase) UserOrderReservation(ctx context.Context, reservation *entity.UserReservation) error {
	tx, err := u.txService.NewTransaction()
	if err != nil {
		errRollback := u.txService.Rollback(tx)
		if errRollback != nil {
			return fmt.Errorf("usecase - UseCase - UserOrderReservation - u.txService.NewTransaction - u.txService.Rollback: %w", err)
		}
		return fmt.Errorf("usecase - UseCase - UserOrderReservation - u.txService.NewTransaction: %w", err)
	}

	err = u.repo.UserOrderReservation(ctx, tx, reservation)
	if err != nil {
		errRollback := u.txService.Rollback(tx)
		if errRollback != nil {
			return fmt.Errorf("usecase - UseCase - UserOrderReservation - u.repo.UserOrderReservation - u.txService.Rollback: %w", err)
		}
		return fmt.Errorf("usecase - UseCase - UserOrderReservation - u.repo.UserOrderReservation: %w", err)
	}

	// Узнать текущий баланс
	userDTO := entity.User{ID: reservation.ID}
	currBalance, err := u.repo.GetBalance(ctx, tx, &userDTO)
	if err != nil {
		errRollback := u.txService.Rollback(tx)
		if errRollback != nil {
			return fmt.Errorf("usecase - UseCase - UserOrderReservation - u.repo.GetBalance - u.txService.Rollback: %w", err)
		}
		return fmt.Errorf("usecase - UseCase - UserOrderReservation - u.repo.GetBalance: %w", err)
	}

	// Проверка на отрицательный баланс
	newBalance := currBalance - reservation.Cost
	if newBalance < 0 {
		errRollback := u.txService.Rollback(tx)
		if errRollback != nil {
			return fmt.Errorf("usecase - UseCase - UserOrderReservation - newBalance - u.txService.Rollback: %w", err)
		}
		return fmt.Errorf("usecase - UseCase - UserOrderReservation: %w", ErrUserNegativeBalance)
	}

	// Списание баланса
	userDTO.ID = reservation.ID
	userDTO.Balance = newBalance

	err = u.repo.MinusBalance(ctx, tx, &userDTO)
	if err != nil {
		errRollback := u.txService.Rollback(tx)
		if errRollback != nil {
			return fmt.Errorf("usecase - UseCase - UserOrderReservation - u.repo.MinusBalance - u.txService.Rollback: %w", err)
		}
		return fmt.Errorf("usecase - UseCase - UserOrderReservation - u.repo.MinusBalance: %w", err)
	}

	return u.txService.Commit(tx)
}
