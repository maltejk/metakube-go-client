// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OpenstackSecurityGroup OpenstackSecurityGroup is the object representing a openstack security group.
//
// swagger:model OpenstackSecurityGroup
type OpenstackSecurityGroup struct {

	// Id uniquely identifies the current security group
	ID string `json:"id,omitempty"`

	// Name is the name of the security group
	Name string `json:"name,omitempty"`
}

// Validate validates this openstack security group
func (m *OpenstackSecurityGroup) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this openstack security group based on context it is used
func (m *OpenstackSecurityGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenstackSecurityGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenstackSecurityGroup) UnmarshalBinary(b []byte) error {
	var res OpenstackSecurityGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
