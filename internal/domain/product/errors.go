package product

import (
	"errors"
)

var (
	ErrProductNotFound         = errors.New("Product not found")
	ErrProductInsufficientStock = errors.New("Product stock is not enough")
)
