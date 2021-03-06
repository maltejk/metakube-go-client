// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// AdmissionPluginList AdmissionPluginList represents a list of admission plugins
//
// swagger:model AdmissionPluginList
type AdmissionPluginList []string

// Validate validates this admission plugin list
func (m AdmissionPluginList) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this admission plugin list based on context it is used
func (m AdmissionPluginList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
