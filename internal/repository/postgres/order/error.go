package order

import "github.com/pkg/errors"

var (
	ErrOrderAlreadyExist   = errors.New("such order already exists")
	ErrRevenueAlreadyExist = errors.New("such revenue already exists")
)
