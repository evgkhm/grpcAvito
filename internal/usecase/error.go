package usecase

import "github.com/pkg/errors"

var (
	errWrongYear           = errors.New("wrong year")
	errWrongMonth          = errors.New("wrong month")
	errUserNegativeBalance = errors.New("you cannot reserve with negative balance")
)
