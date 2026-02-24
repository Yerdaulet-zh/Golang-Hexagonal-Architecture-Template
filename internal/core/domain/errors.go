package domain

import "errors"

var (
	ErrInvalidDSN       = errors.New("Invalid DSN provided")
	ErrPostgreSQLOpenDB = errors.New("Error while opening a postgresql database")
)
