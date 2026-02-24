package domain

import "errors"

var (
	ErrInvalidDSN       = errors.New("invalid DSN provided")
	ErrPostgreSQLOpenDB = errors.New("rrror while opening a postgresql database")
)
