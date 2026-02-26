package domain

import "errors"

var (
	ErrInvalidDSN       = errors.New("invalid DSN provided")
	ErrPostgreSQLOpenDB = errors.New("error while opening a postgresql database")

	ErrInvalidEmailFormat   = errors.New("invalid email format")
	ErrInvalidEmailHost     = errors.New("invalid email host")
	ErrInvalidMessageLenght = errors.New("invalid message length")

	// Repository
	ErrDatabaseInternalError = errors.New("database internal error")
)
