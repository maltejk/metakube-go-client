// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
)

// ResourceLabelMap ResourceLabelMap defines list of labels grouped by specific resource types.
//
// swagger:model ResourceLabelMap
type ResourceLabelMap map[string]LabelKeyList

// Validate validates this resource label map
func (m ResourceLabelMap) Validate(formats strfmt.Registry) error {
	var res []error

	for k := range m {

		if err := m[k].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName(k)
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName(k)
			}
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this resource label map based on the context it is used
func (m ResourceLabelMap) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	for k := range m {

		if err := m[k].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName(k)
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName(k)
			}
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
