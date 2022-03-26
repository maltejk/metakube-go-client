// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClusterTemplate ClusterTemplate represents a ClusterTemplate object
//
// swagger:model ClusterTemplate
type ClusterTemplate struct {

	// ID
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// project ID
	ProjectID string `json:"projectID,omitempty"`

	// scope
	Scope string `json:"scope,omitempty"`

	// user
	User string `json:"user,omitempty"`

	// user SSH keys
	UserSSHKeys []*ClusterTemplateSSHKey `json:"userSshKeys"`

	// cluster
	Cluster *Cluster `json:"cluster,omitempty"`

	// node deployment
	NodeDeployment *NodeDeployment `json:"nodeDeployment,omitempty"`
}

// Validate validates this cluster template
func (m *ClusterTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUserSSHKeys(formats); err != nil {
		res = append(res, err)
	}

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

func (m *ClusterTemplate) validateUserSSHKeys(formats strfmt.Registry) error {
	if swag.IsZero(m.UserSSHKeys) { // not required
		return nil
	}

	for i := 0; i < len(m.UserSSHKeys); i++ {
		if swag.IsZero(m.UserSSHKeys[i]) { // not required
			continue
		}

		if m.UserSSHKeys[i] != nil {
			if err := m.UserSSHKeys[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("userSshKeys" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("userSshKeys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterTemplate) validateCluster(formats strfmt.Registry) error {
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

func (m *ClusterTemplate) validateNodeDeployment(formats strfmt.Registry) error {
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

// ContextValidate validate this cluster template based on the context it is used
func (m *ClusterTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUserSSHKeys(ctx, formats); err != nil {
		res = append(res, err)
	}

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

func (m *ClusterTemplate) contextValidateUserSSHKeys(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.UserSSHKeys); i++ {

		if m.UserSSHKeys[i] != nil {
			if err := m.UserSSHKeys[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("userSshKeys" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("userSshKeys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterTemplate) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ClusterTemplate) contextValidateNodeDeployment(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ClusterTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterTemplate) UnmarshalBinary(b []byte) error {
	var res ClusterTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
