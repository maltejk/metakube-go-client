// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PacketNodeSpec PacketNodeSpec specifies packet specific node settings
//
// swagger:model PacketNodeSpec
type PacketNodeSpec struct {

	// InstanceType denotes the plan to which the device will be provisioned.
	// Required: true
	InstanceType *string `json:"instanceType"`

	// additional instance tags
	Tags []string `json:"tags"`
}

// Validate validates this packet node spec
func (m *PacketNodeSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstanceType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PacketNodeSpec) validateInstanceType(formats strfmt.Registry) error {

	if err := validate.Required("instanceType", "body", m.InstanceType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this packet node spec based on context it is used
func (m *PacketNodeSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PacketNodeSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PacketNodeSpec) UnmarshalBinary(b []byte) error {
	var res PacketNodeSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
