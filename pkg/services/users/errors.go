package users

import "errors"

var (
	ErrUserNotFound = errors.New("requested user could not be found")

	ErrUserQuery = errors.New("requested users could not be retrieved base on the given criteria")

	ErrUserCreate = errors.New("user could not be created")

	ErrNegativeBalance = errors.New("user could not be created with negative balance")

	ErrUserAlreadyExists = errors.New("user with that index already exists")
)
