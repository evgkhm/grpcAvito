package usecase

import (
	"context"
	"errors"
	"grpcAvito/internal/entity"
)

func (u UseCase) Sum(ctx context.Context, userDTO entity.User) error {
	tx, err := u.txService.NewTransaction()
	if err != nil {
		u.txService.Rollback(tx)
		return err
	}
	defer u.txService.Commit(tx)

	//Проверка на отрицательный баланс
	if userDTO.Balance < 0 {
		u.txService.Rollback(tx)
		errUserNegativeBalance := errors.New("you cannot add a negative balance")
		return errUserNegativeBalance
	}

	//Узнать текущий баланс
	currBalance, err := u.repo.GetBalance(ctx, tx, userDTO)
	if err != nil {
		u.txService.Rollback(tx)
		return err
	}

	//Начислить новый баланс
	newBalance := userDTO.Balance + currBalance
	userDTO.Balance = newBalance

	err = u.repo.Sum(ctx, tx, userDTO)
	if err != nil {
		u.txService.Rollback(tx)
		return err
	}

	return nil
}

func (u UseCase) Create(ctx context.Context, userDTO entity.User) error {
	tx, err := u.txService.NewTransaction()
	if err != nil {
		u.txService.Rollback(tx)
		return err
	}
	defer u.txService.Commit(tx)

	err = u.repo.Create(ctx, tx, userDTO)
	if err != nil {
		u.txService.Rollback(tx)
		return err
	}

	return nil
}

func (u UseCase) Reservation(ctx context.Context, reservation entity.UserReservation) error {
	tx, err := u.txService.NewTransaction()
	if err != nil {
		u.txService.Rollback(tx)
		return err
	}
	defer u.txService.Commit(tx)

	err = u.repo.Reservation(reservation, tx)
	if err != nil {
		u.txService.Rollback(tx)
		return err
	}

	//Узнать текущий баланс
	var userDTO entity.User
	currBalance, err := u.repo.GetBalance(ctx, tx, userDTO)
	if err != nil {
		u.txService.Rollback(tx)
		return err
	}

	//Проверка на отрицательный баланс
	if currBalance-reservation.Cost < 0 {
		u.txService.Rollback(tx)
		errUserNegativeBalance := errors.New("you cannot reserve with negative balance")
		return errUserNegativeBalance
	}

	//Списание баланса
	userDTO.Id = reservation.Id
	userDTO.Balance -= reservation.Cost

	err = u.repo.MinusBalance(tx, userDTO)
	if err != nil {
		u.txService.Rollback(tx)
		return err
	}

	return nil
}
