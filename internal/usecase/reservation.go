package usecase

import (
	"context"
	"errors"
	"grpcAvito/internal/entity"
)

func (u UseCase) Dereservation(ctx context.Context, dereservation entity.UserReservation) error {
	tx, err := u.txService.NewTransaction()
	if err != nil {
		u.txService.Rollback(tx)
		return err
	}
	defer u.txService.Commit(tx)

	//Убрать резерв
	err = u.repo.Dereservation(dereservation, tx)
	if err != nil {
		u.txService.Rollback(tx)
		return err
	}

	//Узнать текущий баланс
	var userDTO *entity.User
	currBalance, err := u.repo.GetBalance(ctx, tx, userDTO)
	if err != nil {
		u.txService.Rollback(tx)
		return err
	}

	//Начисление баланса
	newBalance := currBalance + dereservation.Cost
	userDTO.Balance = newBalance
	err = u.repo.Sum(ctx, tx, userDTO)
	if err != nil {
		u.txService.Rollback(tx)
		return err
	}

	return nil
}

func (u UseCase) Reservation(ctx context.Context, reservation *entity.UserReservation) error {
	tx, err := u.txService.NewTransaction()
	if err != nil {
		u.txService.Rollback(tx)
		return err
	}

	err = u.repo.Reservation(reservation, tx)
	if err != nil {
		u.txService.Rollback(tx)
		return err
	}

	//Узнать текущий баланс
	//var userDTO *entity.User
	userDTO := entity.User{Id: reservation.Id}
	currBalance, err := u.repo.GetBalance(ctx, tx, &userDTO)
	if err != nil {
		u.txService.Rollback(tx)
		return err
	}

	//Проверка на отрицательный баланс
	newBalance := currBalance - reservation.Cost
	if newBalance < 0 {
		u.txService.Rollback(tx)
		errUserNegativeBalance := errors.New("you cannot reserve with negative balance")
		return errUserNegativeBalance
	}

	//Списание баланса
	userDTO.Id = reservation.Id
	userDTO.Balance = newBalance

	err = u.repo.MinusBalance(tx, &userDTO)
	if err != nil {
		u.txService.Rollback(tx)
		return err
	}

	return u.txService.Commit(tx)
}
