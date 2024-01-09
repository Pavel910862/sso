package storage

import "errors"

var (
	ErrUserExicts   = errors.New("iser already exists")
	ErrUserNotFound = errors.New("user not found")
	ErrAppNotFound  = errors.New("app not found")
)
