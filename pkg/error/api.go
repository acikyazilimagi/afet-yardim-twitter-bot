package error

import (
	"fmt"
	"net/http"
)

type ApiError struct {
	StatusCode int
	BaseError  error
}

func (e ApiError) Error() string {
	return fmt.Sprintf("StatusCode: %v, BaseError: %v", e.StatusCode, e.BaseError)
}

func NewBadRequestError(baseError error) ApiError {
	return ApiError{
		StatusCode: http.StatusBadRequest,
		BaseError:  baseError,
	}
}

func NewInternalError(baseError error) ApiError {
	return ApiError{
		StatusCode: http.StatusInternalServerError,
		BaseError:  baseError,
	}
}
