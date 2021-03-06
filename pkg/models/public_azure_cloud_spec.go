// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PublicAzureCloudSpec PublicAzureCloudSpec is a public counterpart of apiv1.AzureCloudSpec.
//
// swagger:model PublicAzureCloudSpec
type PublicAzureCloudSpec struct {

	// assign availability set
	AssignAvailabilitySet bool `json:"assignAvailabilitySet,omitempty"`
}

// Validate validates this public azure cloud spec
func (m *PublicAzureCloudSpec) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this public azure cloud spec based on context it is used
func (m *PublicAzureCloudSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PublicAzureCloudSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicAzureCloudSpec) UnmarshalBinary(b []byte) error {
	var res PublicAzureCloudSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
