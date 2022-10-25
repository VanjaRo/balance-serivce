package errors

import "fmt"

const (
	InternalServerError  = "INTERNAL_SERVER_ERROR"
	BadRequest           = "BAD_REQUEST"
	BadGateway           = "BAD_GATEWAY"
	NotFound             = "NOT_FOUND"
	UnsupportedMediaType = "UNSUPPORTED_MEDIA_TYPE"
)

type ErrorCode string

var Desctiptions = map[ErrorCode]string{
	InternalServerError:  "Internal server error",
	BadRequest:           "Bad request",
	BadGateway:           "Bad gateway",
	NotFound:             "Not found",
	UnsupportedMediaType: "Unsupported media type",
}

type AppError struct {
	Code        ErrorCode `json:"code"`
	Description string    `json:"description"`
	Field       string    `json:"field"`
}

type AppErrors struct {
	Errors []AppError `json:"errors"`
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%s (%s) %s", e.Code, e.Field, e.Description)
}

func NewAppError(code ErrorCode, description, field string) error {
	appErr := &AppError{
		Code:        code,
		Description: description,
		Field:       field,
	}
	return appErr
}
