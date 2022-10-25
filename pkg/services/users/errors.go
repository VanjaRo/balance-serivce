package users

import "errors"

var (
	ErrUserNotFound = errors.New("requested user could not be found")

	ErrUserQuery = errors.New("requested users could not be retrieved base on the given criteria")

	ErrUserCreate = errors.New("user could not be created")
)
