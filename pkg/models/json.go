// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// JSON JSON represents any valid JSON value.
//
// These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil.
//
// swagger:model JSON
type JSON struct {

	// raw
	Raw []uint8 `json:"Raw"`
}

// Validate validates this JSON
func (m *JSON) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this JSON based on context it is used
func (m *JSON) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *JSON) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JSON) UnmarshalBinary(b []byte) error {
	var res JSON
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
