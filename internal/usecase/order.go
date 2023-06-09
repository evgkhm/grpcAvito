package usecase

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"grpcAvito/internal/entity"
	"grpcAvito/internal/repository/postgres"
)

type OrderUseCase struct {
	userRepo   postgres.UserRepository
	orderRepo  postgres.OrderRepository
	reportRepo postgres.ReportRepository
	log        *logrus.Logger
	txService  *TransactionServiceImpl
}

func NewOrderUseCase(userRepo postgres.UserRepository, orderRepo postgres.OrderRepository, reportRepo postgres.ReportRepository,
	log *logrus.Logger, txService *TransactionServiceImpl) *OrderUseCase {
	return &OrderUseCase{
		userRepo:   userRepo,
		orderRepo:  orderRepo,
		reportRepo: reportRepo,
		log:        log,
		txService:  txService,
	}
}

func (u OrderUseCase) UserOrderRevenue(ctx context.Context, revenue *entity.UserRevenue) error {
	tx, err := u.txService.NewTransaction()
	if err != nil {
		errRollback := u.txService.Rollback(tx)
		if errRollback != nil {
			return fmt.Errorf("usecase - UseCase - UserOrderRevenue - u.txService.NewTransaction - u.txService.Rollback: %w", err)
		}
		return fmt.Errorf("usecase - UseCase - UserOrderRevenue - u.txService.NewTransaction: %w", err)
	}

	reservation := &entity.UserReservation{ID: revenue.ID,
		OrderID:   revenue.OrderID,
		ServiceID: revenue.ServiceID,
		Cost:      revenue.Cost,
	}

	// Списать с резервации
	err = u.orderRepo.UserOrderDeleteReservation(ctx, tx, reservation)
	if err != nil {
		errRollback := u.txService.Rollback(tx)
		if errRollback != nil {
			return fmt.Errorf("usecase - UseCase - UserOrderRevenue - u.repo.UserOrderDeleteReservation - u.txService.Rollback: %w", err)
		}
		return fmt.Errorf("usecase - UseCase - UserOrderRevenue - u.repo.UserOrderDeleteReservation: %w", err)
	}

	// Начислить в revenue
	err = u.orderRepo.UserOrderRevenue(ctx, tx, revenue)
	if err != nil {
		errRollback := u.txService.Rollback(tx)
		if errRollback != nil {
			return fmt.Errorf("usecase - UseCase - UserOrderRevenue - u.repo.UserOrderRevenue - u.txService.Rollback: %w", err)
		}
		return fmt.Errorf("usecase - UseCase - UserOrderRevenue - u.repo.UserOrderRevenue: %w", err)
	}

	return u.txService.Commit(tx)
}

func (u OrderUseCase) UserOrderDeleteReservation(ctx context.Context, reservation *entity.UserReservation) error {
	tx, err := u.txService.NewTransaction()
	if err != nil {
		errRollback := u.txService.Rollback(tx)
		if errRollback != nil {
			return fmt.Errorf("usecase - UseCase - UserOrderDeleteReservation - u.txService.NewTransaction - u.txService.Rollback: %w", err)
		}
		return fmt.Errorf("usecase - UseCase - UserOrderDeleteReservation - u.txService.NewTransaction: %w", err)
	}

	// Убрать резерв
	err = u.orderRepo.UserOrderDeleteReservation(ctx, tx, reservation)
	if err != nil {
		errRollback := u.txService.Rollback(tx)
		if errRollback != nil {
			return fmt.Errorf("usecase - UseCase - UserOrderDeleteReservation - u.repo.UserOrderDeleteReservation - u.txService.Rollback: %w", err)
		}
		return fmt.Errorf("usecase - UseCase - UserOrderDeleteReservation - u.repo.UserOrderDeleteReservation: %w", err)
	}

	// Узнать текущий баланс
	userDTO := entity.User{ID: reservation.ID}
	currBalance, err := u.userRepo.GetBalance(ctx, tx, &userDTO)
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
	err = u.userRepo.UserBalanceAccrual(ctx, tx, &userDTO)
	if err != nil {
		errRollback := u.txService.Rollback(tx)
		if errRollback != nil {
			return fmt.Errorf("usecase - UseCase - UserOrderDeleteReservation - u.repo.UserBalanceAccrual - u.txService.Rollback: %w", err)
		}
		return fmt.Errorf("usecase - UseCase - UserOrderDeleteReservation - u.repo.UserBalanceAccrual: %w", err)
	}

	return u.txService.Commit(tx)
}

func (u OrderUseCase) UserOrderReservation(ctx context.Context, reservation *entity.UserReservation) error {
	tx, err := u.txService.NewTransaction()
	if err != nil {
		errRollback := u.txService.Rollback(tx)
		if errRollback != nil {
			return fmt.Errorf("usecase - UseCase - UserOrderReservation - u.txService.NewTransaction - u.txService.Rollback: %w", err)
		}
		return fmt.Errorf("usecase - UseCase - UserOrderReservation - u.txService.NewTransaction: %w", err)
	}

	err = u.orderRepo.UserOrderReservation(ctx, tx, reservation)
	if err != nil {
		errRollback := u.txService.Rollback(tx)
		if errRollback != nil {
			return fmt.Errorf("usecase - UseCase - UserOrderReservation - u.repo.UserOrderReservation - u.txService.Rollback: %w", err)
		}
		return fmt.Errorf("usecase - UseCase - UserOrderReservation - u.repo.UserOrderReservation: %w", err)
	}

	// Узнать текущий баланс
	userDTO := entity.User{ID: reservation.ID}
	currBalance, err := u.userRepo.GetBalance(ctx, tx, &userDTO)
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

	err = u.orderRepo.MinusBalance(ctx, tx, &userDTO)
	if err != nil {
		errRollback := u.txService.Rollback(tx)
		if errRollback != nil {
			return fmt.Errorf("usecase - UseCase - UserOrderReservation - u.repo.MinusBalance - u.txService.Rollback: %w", err)
		}
		return fmt.Errorf("usecase - UseCase - UserOrderReservation - u.repo.MinusBalance: %w", err)
	}

	return u.txService.Commit(tx)
}
