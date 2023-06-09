package postgres

import "github.com/pkg/errors"

var (
	ErrUserAlreadyExist    = errors.New("such user already exists")
	ErrOrderAlreadyExist   = errors.New("such order already exists")
	ErrRevenueAlreadyExist = errors.New("such revenue already exists")
	ErrUserNotExist        = errors.New("such user does not exist")
	ErrGetYearMonth        = errors.New("getting year or month")
)
