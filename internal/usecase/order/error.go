package order

import "github.com/pkg/errors"

var (
	ErrUserNegativeBalance = errors.New("you cannot reserve with negative balance")
)
