package errors

import "errors"

var ErrTitleIsEmpty = errors.New("error: title is empty")
var ErrTaskAlreadyExists = errors.New("error: task already exists")
var ErrTaskAlreadyDone = errors.New("error: task already done")
var ErrTaskIsNotFound = errors.New("error: task is not found")
var ErrInput = errors.New("error: user input")
var ErrEmptyCommand = errors.New("error: empty command")
var ErrNotEnoughAruments = errors.New("error: not enough arguments")
var ErrConvertStrconv = errors.New("error: convert strconv.Atoi to int")
var ErrListIsEmpty = errors.New("error: list is empty")
