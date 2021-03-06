// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MeteringConfigurations MeteringConfigurations contains all the configurations for the metering tool.
//
// swagger:model MeteringConfigurations
type MeteringConfigurations struct {

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// StorageClassName is the name of the storage class that the metering tool uses to save processed files before
	// exporting it to s3 bucket. Default value is kubermatic-fast.
	StorageClassName string `json:"storageClassName,omitempty"`

	// StorageSize is the size of the storage class. Default value is 100Gi.
	StorageSize string `json:"storageSize,omitempty"`
}

// Validate validates this metering configurations
func (m *MeteringConfigurations) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this metering configurations based on context it is used
func (m *MeteringConfigurations) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MeteringConfigurations) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MeteringConfigurations) UnmarshalBinary(b []byte) error {
	var res MeteringConfigurations
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
