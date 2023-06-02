package postgres

import "github.com/pkg/errors"

var (
	errUserAlreadyExist = errors.New("such user already exists")
	errUserNotExist     = errors.New("such user does not exist")
	errGetYearMonth     = errors.New("getting year or month")
)
