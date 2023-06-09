package user

import "github.com/pkg/errors"

var (
	ErrUserAccrualNegativeBalance = errors.New("you cannot accrual with negative balance")
)
