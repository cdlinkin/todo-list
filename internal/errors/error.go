package errors

import "errors"

var (
	ErrTitleIsEmpty      = errors.New("error: title is empty")
	ErrTaskAlreadyExists = errors.New("error: task already exists")
	ErrTaskAlreadyDone   = errors.New("error: task already done")
	ErrTaskIsNotFound    = errors.New("error: task is not found")
	ErrInput             = errors.New("error: user input")
	ErrEmptyCommand      = errors.New("error: empty command")
	ErrNotEnoughAruments = errors.New("error: not enough arguments")
	ErrConvertStrconv    = errors.New("error: convert strconv.Atoi to int")
	ErrListIsEmpty       = errors.New("error: list is empty")
)
