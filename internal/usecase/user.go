package usecase

import (
	"context"
	"github.com/pkg/errors"
	"grpcAvito/internal/entity"
)

func (u UseCase) Sum(ctx context.Context, userDTO *entity.User) error {
	tx, err := u.txService.NewTransaction()
	if err != nil {
		u.txService.Rollback(tx)
		return err
	}

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

	return u.txService.Commit(tx)
}

func (u UseCase) Create(ctx context.Context, userDTO entity.User) error {
	tx, err := u.txService.NewTransaction()
	if err != nil {
		u.txService.Rollback(tx)
		return err
	}

	err = u.repo.Create(ctx, tx, userDTO)
	if err != nil {
		u.txService.Rollback(tx)
		return err
	}

	return u.txService.Commit(tx)
}

func (u UseCase) GetBalance(ctx context.Context, dto *entity.User) error {
	tx, err := u.txService.NewTransaction()
	if err != nil {
		u.txService.Rollback(tx)
		return err
	}

	dto.Balance, err = u.repo.GetBalance(ctx, tx, dto)
	if err != nil {
		u.txService.Rollback(tx)
		return err
	}

	return u.txService.Commit(tx)
}
