package domain

import "errors"

var (
	ErrInvalidDSN       = errors.New("invalid DSN provided")
	ErrPostgreSQLOpenDB = errors.New("error while opening a postgresql database")

	ErrInvalidEmailFormat = errors.New("invalid email format")
)
