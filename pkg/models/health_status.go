// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// HealthStatus health status
//
// swagger:model HealthStatus
type HealthStatus int64

// Validate validates this health status
func (m HealthStatus) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this health status based on context it is used
func (m HealthStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
