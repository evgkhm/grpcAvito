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
		return fmt.Errorf("usecase: DeleteReservation: NewTransaction: %w", err)
	}

	//Убрать резерв
	err = u.repo.DeleteReservation(ctx, tx, reservation)
	if err != nil {
		_ = u.txService.Rollback(tx)
		return fmt.Errorf("usecase: DeleteReservation: DeleteReservation: %w", err)
	}

	//Узнать текущий баланс
	//var userDTO *entity.User
	userDTO := entity.User{Id: reservation.Id}
	currBalance, err := u.repo.GetBalance(ctx, tx, &userDTO)
	if err != nil {
		_ = u.txService.Rollback(tx)
		return fmt.Errorf("usecase: DeleteReservation: GetBalance: %w", err)
	}

	//Начисление баланса
	newBalance := currBalance + reservation.Cost
	userDTO.Balance = newBalance
	err = u.repo.Sum(ctx, tx, &userDTO)
	if err != nil {
		_ = u.txService.Rollback(tx)
		return fmt.Errorf("usecase: DeleteReservation: Sum: %w", err)
	}

	return u.txService.Commit(tx)
}

func (u UseCase) Reservation(ctx context.Context, reservation *entity.UserReservation) error {
	tx, err := u.txService.NewTransaction()
	if err != nil {
		_ = u.txService.Rollback(tx)
		return fmt.Errorf("usecase: Reservation: NewTransaction: %w", err)
	}

	err = u.repo.Reservation(ctx, tx, reservation)
	if err != nil {
		_ = u.txService.Rollback(tx)
		return fmt.Errorf("usecase: Reservation: Reservation: %w", err)
	}

	//Узнать текущий баланс
	//var userDTO *entity.User
	userDTO := entity.User{Id: reservation.Id}
	currBalance, err := u.repo.GetBalance(ctx, tx, &userDTO)
	if err != nil {
		_ = u.txService.Rollback(tx)
		return fmt.Errorf("usecase: Reservation: GetBalance: %w", err)
	}

	//Проверка на отрицательный баланс
	newBalance := currBalance - reservation.Cost
	if newBalance < 0 {
		_ = u.txService.Rollback(tx)
		return fmt.Errorf("usecase: Reservation: %w", errUserNegativeBalance)
	}

	//Списание баланса
	userDTO.Id = reservation.Id
	userDTO.Balance = newBalance

	err = u.repo.MinusBalance(ctx, tx, &userDTO)
	if err != nil {
		_ = u.txService.Rollback(tx)
		return fmt.Errorf("usecase: Reservation: MinusBalance: %w", err)
	}

	return u.txService.Commit(tx)
}
