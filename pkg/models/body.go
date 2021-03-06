// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Body body
//
// swagger:model body
type Body struct {

	// Kubeconfig Base64 encoded kubeconfig
	Kubeconfig string `json:"kubeconfig,omitempty"`

	// Name is human readable name for the external cluster
	Name string `json:"name,omitempty"`
}

// Validate validates this body
func (m *Body) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this body based on context it is used
func (m *Body) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Body) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Body) UnmarshalBinary(b []byte) error {
	var res Body
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
