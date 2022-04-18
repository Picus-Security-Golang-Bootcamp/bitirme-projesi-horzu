package httpErrors

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/horzu/golang/cart-api/internal/api"
	"gorm.io/gorm"
)

var (
	InternalServerError   = errors.New("Internal Server Error")
	NotFound              = errors.New("Not Found")
	RequestTimeoutError   = errors.New("Request Timeout")
	CannotBindGivenData   = errors.New("Could not bind given data")
	ValidationError       = errors.New("Validation failed for given payload")
	UniqueError           = errors.New("Item should be unique on database")
	LoginError            = errors.New("Wrong username or password")
	ErrItemAlreadyInCart  = errors.New("Item is already in cart")
	ErrInvalidOrder       = errors.New("Quantity should be positive integer")
	ErrInvalidUpdateQty   = errors.New("Update quantity should be positive integer")
	ErrItemNotExistInCart = errors.New("Item does not exist in the cart")
	ErrProductNotFound         = errors.New("Product not found")
	ErrProductInsufficientStock = errors.New("Product stock is not enough")
)

type RestError api.SuccessfulAPIResponse

type RestErr interface {
	Status() int
	Error() string
	Causes() interface{}
}

// Error  Error() interface method
func (e RestError) Error() string {
	return fmt.Sprintf("status: %d - errors: %s - causes: %v", e.Code, e.Message, e.Data)
}

func (e RestError) Status() int {
	return int(e.Code)
}

func (e RestError) Causes() interface{} {
	return e.Data
}

func NewRestError(status int, err string, causes interface{}) RestErr {
	return RestError{
		Code:    int64(status),
		Message: err,
		Data:    causes,
	}
}

func NewInternalServerError(causes interface{}) RestErr {
	result := RestError{
		Code:    http.StatusInternalServerError,
		Message: InternalServerError.Error(),
		Data:    causes,
	}
	return result
}

// ParseErrors Parser of error string messages returns RestError
func ParseErrors(err error) RestErr {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return NewRestError(http.StatusNotFound, NotFound.Error(), err)
	case errors.Is(err, context.DeadlineExceeded):
		return NewRestError(http.StatusRequestTimeout, RequestTimeoutError.Error(), err)
	case errors.Is(err, CannotBindGivenData):
		return NewRestError(http.StatusBadRequest, CannotBindGivenData.Error(), err)
	case errors.Is(err, gorm.ErrRecordNotFound):
		return NewRestError(http.StatusNotFound, gorm.ErrRecordNotFound.Error(), err)
	case strings.Contains(err.Error(), "validation"):
		return NewRestError(http.StatusBadRequest, ValidationError.Error(), err)
	case strings.Contains(err.Error(), "23505"):
		return NewRestError(http.StatusBadRequest, UniqueError.Error(), err)
	case strings.Contains(err.Error(), "crypto/bcrypt"):
		return NewRestError(http.StatusBadRequest, LoginError.Error(), err)

	default:
		if restErr, ok := err.(RestErr); ok {
			return restErr
		}
		return NewInternalServerError(err)
	}
}

func ErrorResponse(err error) (int, interface{}) {
	return ParseErrors(err).Status(), ParseErrors(err)
}
