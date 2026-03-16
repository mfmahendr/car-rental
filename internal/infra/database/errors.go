package database

import "errors"

var (
	ErrDatabaseParsing  = errors.New("failed to parse database configuration")
	ErrDatabaseCreation = errors.New("failed to create database pool connection")
	ErrUnreachable      = errors.New("database is unreachable")

	ErrDBOperation = errors.New("database operation error")
	ErrDBIterating = errors.New("error iterating rows")
	ErrDBScan = errors.New("failed to scan car")
)
