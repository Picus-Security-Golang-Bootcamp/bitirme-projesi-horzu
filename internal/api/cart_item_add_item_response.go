// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CartItemAddItemResponse cart item add item response
//
// swagger:model cartItem.AddItemResponse
type CartItemAddItemResponse struct {

	// category ID
	CategoryID string `json:"categoryID,omitempty"`

	// desc
	Desc string `json:"desc,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// price
	Price float64 `json:"price,omitempty"`

	// quantity
	Quantity int64 `json:"quantity,omitempty"`

	// sku
	Sku string `json:"sku,omitempty"`
}

// Validate validates this cart item add item response
func (m *CartItemAddItemResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cart item add item response based on context it is used
func (m *CartItemAddItemResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CartItemAddItemResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CartItemAddItemResponse) UnmarshalBinary(b []byte) error {
	var res CartItemAddItemResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
