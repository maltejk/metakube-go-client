// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AzureSecurityGroupsList AzureSecurityGroupsList is the object representing the security groups for vms in azure cloud provider
//
// swagger:model AzureSecurityGroupsList
type AzureSecurityGroupsList struct {

	// security groups
	SecurityGroups []string `json:"securityGroups"`
}

// Validate validates this azure security groups list
func (m *AzureSecurityGroupsList) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this azure security groups list based on context it is used
func (m *AzureSecurityGroupsList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AzureSecurityGroupsList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureSecurityGroupsList) UnmarshalBinary(b []byte) error {
	var res AzureSecurityGroupsList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
