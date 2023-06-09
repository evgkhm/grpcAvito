package usecase

import "github.com/pkg/errors"

var (
	ErrWrongYear                  = errors.New("wrong year")
	ErrWrongMonth                 = errors.New("wrong month")
	ErrUserNegativeBalance        = errors.New("you cannot reserve with negative balance")
	ErrUserAccrualNegativeBalance = errors.New("you cannot accrual with negative balance")
)
