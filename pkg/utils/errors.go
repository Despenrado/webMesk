package utils

import (
	"errors"
	"strconv"
)

var (
	// ErrRecordNotFound ...
	ErrRecordNotFound = errors.New("record not found")
	// ErrRecordAlreadyExists ...
	ErrRecordAlreadyExists = errors.New("record already exists")
	// ErrIncorrectEmailOrPassword ...
	ErrIncorrectEmailOrPassword = errors.New("Incorrect Email or Password")
	// ErrWrongRequest ...
	ErrWrongRequest = errors.New("wrong request")

	// ErrRowsNumberAffected ...
)

func ErrRowsNumberAffected(rowNumber int) error {
	return errors.New("Rows Affected " + strconv.Itoa(rowNumber))
}
