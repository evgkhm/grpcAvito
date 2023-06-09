package report

import "github.com/pkg/errors"

var (
	ErrWrongYear  = errors.New("wrong year")
	ErrWrongMonth = errors.New("wrong month")
)
