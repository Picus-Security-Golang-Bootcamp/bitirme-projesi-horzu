package cart

import (
	"errors"
)

var (
	ErrItemAlreadyInCart  = errors.New("Item is already in cart")
	ErrInvalidOrder       = errors.New("Quantity should be positive integer")
	ErrInvalidUpdateQty   = errors.New("Update quantity should be positive integer")
	ErrItemNotExistInCart = errors.New("Item does not exist in the cart")
)
