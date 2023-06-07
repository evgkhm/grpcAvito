package usecase

import (
	"context"
	"fmt"
	"grpcAvito/internal/entity"
)

func (u UseCase) DeleteReservation(ctx context.Context, reservation *entity.UserReservation) error {
	tx, err := u.txService.NewTransaction()
	if err != nil {
		_ = u.txService.Rollback(tx)
		return fmt.Errorf("usecase - UseCase - UserOrderDeleteReservation - u.txService.NewTransaction: %w", err)
	}

	// Убрать резерв
	err = u.repo.DeleteReservation(ctx, tx, reservation)
	if err != nil {
		_ = u.txService.Rollback(tx)
		return fmt.Errorf("usecase - UseCase - UserOrderDeleteReservation - u.repo.UserOrderDeleteReservation: %w", err)
	}

	// Узнать текущий баланс
	userDTO := entity.User{ID: reservation.ID}
	currBalance, err := u.repo.GetBalance(ctx, tx, &userDTO)
	if err != nil {
		_ = u.txService.Rollback(tx)
		return fmt.Errorf("usecase - UseCase - UserOrderDeleteReservation - u.repo.GetBalance: %w", err)
	}

	// Начисление баланса
	newBalance := currBalance + reservation.Cost
	userDTO.Balance = newBalance
	err = u.repo.Sum(ctx, tx, &userDTO)
	if err != nil {
		_ = u.txService.Rollback(tx)
		return fmt.Errorf("usecase - UseCase - UserOrderDeleteReservation - u.repo.Sum: %w", err)
	}

	return u.txService.Commit(tx)
}

func (u UseCase) Reservation(ctx context.Context, reservation *entity.UserReservation) error {
	tx, err := u.txService.NewTransaction()
	if err != nil {
		_ = u.txService.Rollback(tx)
		return fmt.Errorf("usecase - UseCase - UserOrderReservation - u.txService.NewTransaction: %w", err)
	}

	err = u.repo.Reservation(ctx, tx, reservation)
	if err != nil {
		_ = u.txService.Rollback(tx)
		return fmt.Errorf("usecase - UseCase - UserOrderReservation - u.repo.UserOrderReservation: %w", err)
	}

	// Узнать текущий баланс
	userDTO := entity.User{ID: reservation.ID}
	currBalance, err := u.repo.GetBalance(ctx, tx, &userDTO)
	if err != nil {
		_ = u.txService.Rollback(tx)
		return fmt.Errorf("usecase - UseCase - UserOrderReservation - u.repo.GetBalance: %w", err)
	}

	// Проверка на отрицательный баланс
	newBalance := currBalance - reservation.Cost
	if newBalance < 0 {
		_ = u.txService.Rollback(tx)
		return fmt.Errorf("usecase - UseCase - UserOrderReservation: %w", errUserNegativeBalance)
	}

	// Списание баланса
	userDTO.ID = reservation.ID
	userDTO.Balance = newBalance

	err = u.repo.MinusBalance(ctx, tx, &userDTO)
	if err != nil {
		_ = u.txService.Rollback(tx)
		return fmt.Errorf("usecase - UseCase - UserOrderReservation - u.repo.MinusBalance: %w", err)
	}

	return u.txService.Commit(tx)
}
