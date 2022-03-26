// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateClusterSpec CreateClusterSpec is the structure that is used to create cluster with its initial node deployment
//
// swagger:model CreateClusterSpec
type CreateClusterSpec struct {

	// DNS domain
	DNSDomain string `json:"dnsDomain,omitempty"`

	// pods c ID r
	PodsCIDR string `json:"podsCIDR,omitempty"`

	// services c ID r
	ServicesCIDR string `json:"servicesCIDR,omitempty"`

	// cluster
	Cluster *Cluster `json:"cluster,omitempty"`

	// node deployment
	NodeDeployment *NodeDeployment `json:"nodeDeployment,omitempty"`
}

// Validate validates this create cluster spec
func (m *CreateClusterSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeDeployment(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateClusterSpec) validateCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.Cluster) { // not required
		return nil
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *CreateClusterSpec) validateNodeDeployment(formats strfmt.Registry) error {
	if swag.IsZero(m.NodeDeployment) { // not required
		return nil
	}

	if m.NodeDeployment != nil {
		if err := m.NodeDeployment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodeDeployment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nodeDeployment")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create cluster spec based on the context it is used
func (m *CreateClusterSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNodeDeployment(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateClusterSpec) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.Cluster != nil {
		if err := m.Cluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *CreateClusterSpec) contextValidateNodeDeployment(ctx context.Context, formats strfmt.Registry) error {

	if m.NodeDeployment != nil {
		if err := m.NodeDeployment.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodeDeployment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nodeDeployment")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateClusterSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateClusterSpec) UnmarshalBinary(b []byte) error {
	var res CreateClusterSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}