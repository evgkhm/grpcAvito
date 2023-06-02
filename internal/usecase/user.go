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
		return fmt.Errorf("usecase: Sum: NewTransaction: %w", err)
	}

	//Проверка на отрицательный баланс
	if userDTO.Balance < 0 {
		_ = u.txService.Rollback(tx)
		return fmt.Errorf("usecase: Sum: %w", errUserNegativeBalance)
	}

	//Узнать текущий баланс
	currBalance, err := u.repo.GetBalance(ctx, tx, userDTO)
	if err != nil {
		_ = u.txService.Rollback(tx)
		return fmt.Errorf("usecase: Sum: GetBalance: %w", err)
	}

	//Начислить новый баланс
	newBalance := userDTO.Balance + currBalance
	userDTO.Balance = newBalance

	err = u.repo.Sum(ctx, tx, userDTO)
	if err != nil {
		_ = u.txService.Rollback(tx)
		return fmt.Errorf("usecase: Sum: Sum: %w", err)
	}

	return u.txService.Commit(tx)
}

func (u UseCase) Create(ctx context.Context, userDTO *entity.User) error {
	tx, err := u.txService.NewTransaction()
	if err != nil {
		_ = u.txService.Rollback(tx)
		return fmt.Errorf("usecase: Create: NewTransaction: %w", err)
	}

	err = u.repo.Create(ctx, tx, userDTO)
	if err != nil {
		_ = u.txService.Rollback(tx)
		return fmt.Errorf("usecase: Create: Create: %w", err)
	}

	return u.txService.Commit(tx)
}

func (u UseCase) GetBalance(ctx context.Context, dto *entity.User) error {
	tx, err := u.txService.NewTransaction()
	if err != nil {
		_ = u.txService.Rollback(tx)
		return fmt.Errorf("usecase: GetBalance: NewTransaction: %w", err)
	}

	dto.Balance, err = u.repo.GetBalance(ctx, tx, dto)
	if err != nil {
		_ = u.txService.Rollback(tx)
		return fmt.Errorf("usecase: GetBalance: GetBalance: %w", err)
	}

	return u.txService.Commit(tx)
}
