package usecase

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"grpcAvito/internal/entity"
	"grpcAvito/internal/repository/postgres"
)

type UserUseCase struct {
	repo      postgres.Repository
	log       *logrus.Logger
	txService *TransactionServiceImpl
}

func NewUserUseCase(repo *postgres.Repo, log *logrus.Logger, txService *TransactionServiceImpl) *UserUseCase {
	return &UserUseCase{
		repo:      repo,
		log:       log,
		txService: txService,
	}
}

func (u UserUseCase) UserBalanceAccrual(ctx context.Context, userDTO *entity.User) error {
	tx, err := u.txService.NewTransaction()
	if err != nil {
		errRollback := u.txService.Rollback(tx)
		if errRollback != nil {
			return fmt.Errorf("usecase - UseCase - UserBalanceAccrual - u.txService.NewTransaction - u.txService.Rollback: %w", err)
		}
		return fmt.Errorf("usecase - UseCase - UserBalanceAccrual - u.txService.NewTransaction: %w", err)
	}

	// Проверка на отрицательный баланс
	if userDTO.Balance < 0 {
		errRollback := u.txService.Rollback(tx)
		if errRollback != nil {
			return fmt.Errorf("usecase - UseCase - UserBalanceAccrual - u.txService.Rollback: %w", err)
		}
		return fmt.Errorf("usecase - UseCase - UserBalanceAccrual: %w", ErrUserAccrualNegativeBalance)
	}

	// Узнать текущий баланс
	currBalance, err := u.repo.GetBalance(ctx, tx, userDTO)
	if err != nil {
		errRollback := u.txService.Rollback(tx)
		if errRollback != nil {
			return fmt.Errorf("usecase - UseCase - UserBalanceAccrual - u.repo.GetBalance - u.txService.Rollback: %w", err)
		}
		return fmt.Errorf("usecase - UseCase - UserBalanceAccrual - u.repo.GetBalance: %w", err)
	}

	// Начислить новый баланс
	newBalance := userDTO.Balance + currBalance
	userDTO.Balance = newBalance

	err = u.repo.UserBalanceAccrual(ctx, tx, userDTO)
	if err != nil {
		errRollback := u.txService.Rollback(tx)
		if errRollback != nil {
			return fmt.Errorf("usecase - UseCase - UserBalanceAccrual - u.repo.UserBalanceAccrual - u.txService.Rollback: %w", err)
		}
		return fmt.Errorf("usecase - UseCase - UserBalanceAccrual - u.repo.UserBalanceAccrual: %w", err)
	}

	return u.txService.Commit(tx)
}

func (u UserUseCase) CreateUser(ctx context.Context, userDTO *entity.User) error {
	tx, err := u.txService.NewTransaction()
	if err != nil {
		errRollback := u.txService.Rollback(tx)
		if errRollback != nil {
			return fmt.Errorf("usecase - UseCase - CreateUser - u.txService.NewTransaction - u.txService.Rollback: %w", err)
		}
		return fmt.Errorf("usecase - UseCase - CreateUser - u.txService.NewTransaction: %w", err)
	}

	err = u.repo.CreateUser(ctx, tx, userDTO)
	if err != nil {
		errRollback := u.txService.Rollback(tx)
		if errRollback != nil {
			return fmt.Errorf("usecase - UseCase - CreateUser - u.repo.CreateUser - u.txService.Rollback: %w", err)
		}
		return fmt.Errorf("usecase - UseCase - CreateUser - u.txService.CreateUser: %w", err)
	}

	return u.txService.Commit(tx)
}

func (u UserUseCase) GetBalance(ctx context.Context, dto *entity.User) error {
	tx, err := u.txService.NewTransaction()
	if err != nil {
		errRollback := u.txService.Rollback(tx)
		if errRollback != nil {
			return fmt.Errorf("usecase - UseCase - GetBalance - u.txService.NewTransaction - u.txService.Rollback: %w", err)
		}
		return fmt.Errorf("usecase - UseCase - GetBalance - u.txService.NewTransaction: %w", err)
	}

	dto.Balance, err = u.repo.GetBalance(ctx, tx, dto)
	if err != nil {
		errRollback := u.txService.Rollback(tx)
		if errRollback != nil {
			return fmt.Errorf("usecase - UseCase - GetBalance - u.repo.GetBalance - u.txService.Rollback: %w", err)
		}
		return fmt.Errorf("usecase - UseCase - GetBalance - u.txService.GetBalance: %w", err)
	}

	return u.txService.Commit(tx)
}
