// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VSphereNodeSpec VSphereNodeSpec VSphere node settings
//
// swagger:model VSphereNodeSpec
type VSphereNodeSpec struct {

	// c p us
	CPUs int64 `json:"cpus,omitempty"`

	// disk size g b
	DiskSizeGB int64 `json:"diskSizeGB,omitempty"`

	// memory
	Memory int64 `json:"memory,omitempty"`

	// template
	Template string `json:"template,omitempty"`
}

// Validate validates this v sphere node spec
func (m *VSphereNodeSpec) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v sphere node spec based on context it is used
func (m *VSphereNodeSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VSphereNodeSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VSphereNodeSpec) UnmarshalBinary(b []byte) error {
	var res VSphereNodeSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
