package usecase

import (
	"context"
	"fmt"
	"grpcAvito/internal/entity"
)

func (u UseCase) Sum(ctx context.Context, userDTO *entity.User) error {
	tx, err := u.txService.NewTransaction()
	if err != nil {
		_ = u.txService.Rollback(tx)
		return fmt.Errorf("usecase - UseCase - Sum - u.txService.NewTransaction: %w", err)
	}

	//Проверка на отрицательный баланс
	if userDTO.Balance < 0 {
		_ = u.txService.Rollback(tx)
		return fmt.Errorf("usecase - UseCase - Sum: %w", errUserNegativeBalance)
	}

	//Узнать текущий баланс
	currBalance, err := u.repo.GetBalance(ctx, tx, userDTO)
	if err != nil {
		_ = u.txService.Rollback(tx)
		return fmt.Errorf("usecase - UseCase - Sum - u.repo.GetBalance: %w", err)
	}

	//Начислить новый баланс
	newBalance := userDTO.Balance + currBalance
	userDTO.Balance = newBalance

	err = u.repo.Sum(ctx, tx, userDTO)
	if err != nil {
		_ = u.txService.Rollback(tx)
		return fmt.Errorf("usecase - UseCase - Sum - u.repo.Sum: %w", err)
	}

	return u.txService.Commit(tx)
}

func (u UseCase) Create(ctx context.Context, userDTO *entity.User) error {
	tx, err := u.txService.NewTransaction()
	if err != nil {
		_ = u.txService.Rollback(tx)
		return fmt.Errorf("usecase - UseCase - Create - u.txService.NewTransaction: %w", err)
	}

	err = u.repo.Create(ctx, tx, userDTO)
	if err != nil {
		_ = u.txService.Rollback(tx)
		return fmt.Errorf("usecase - UseCase - Create - u.txService.Create: %w", err)
	}

	return u.txService.Commit(tx)
}

func (u UseCase) GetBalance(ctx context.Context, dto *entity.User) error {
	tx, err := u.txService.NewTransaction()
	if err != nil {
		_ = u.txService.Rollback(tx)
		return fmt.Errorf("usecase - UseCase - GetBalance - u.txService.NewTransaction: %w", err)
	}

	dto.Balance, err = u.repo.GetBalance(ctx, tx, dto)
	if err != nil {
		_ = u.txService.Rollback(tx)
		return fmt.Errorf("usecase - UseCase - GetBalance - u.txService.GetBalance: %w", err)
	}

	return u.txService.Commit(tx)
}
