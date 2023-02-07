package error

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiError struct {
	StatusCode int
	BaseError  error
}

func (e ApiError) Error() string {
	return fmt.Sprintf("StatusCode: %v, BaseError: %v", e.StatusCode, e.BaseError)
}

func (e ApiError) ToJson() gin.H {
	return gin.H{"status": e.StatusCode, "error": e.BaseError.Error()}
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
