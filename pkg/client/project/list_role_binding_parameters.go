// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListRoleBindingParams creates a new ListRoleBindingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListRoleBindingParams() *ListRoleBindingParams {
	return &ListRoleBindingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListRoleBindingParamsWithTimeout creates a new ListRoleBindingParams object
// with the ability to set a timeout on a request.
func NewListRoleBindingParamsWithTimeout(timeout time.Duration) *ListRoleBindingParams {
	return &ListRoleBindingParams{
		timeout: timeout,
	}
}

// NewListRoleBindingParamsWithContext creates a new ListRoleBindingParams object
// with the ability to set a context for a request.
func NewListRoleBindingParamsWithContext(ctx context.Context) *ListRoleBindingParams {
	return &ListRoleBindingParams{
		Context: ctx,
	}
}

// NewListRoleBindingParamsWithHTTPClient creates a new ListRoleBindingParams object
// with the ability to set a custom HTTPClient for a request.
func NewListRoleBindingParamsWithHTTPClient(client *http.Client) *ListRoleBindingParams {
	return &ListRoleBindingParams{
		HTTPClient: client,
	}
}

/* ListRoleBindingParams contains all the parameters to send to the API endpoint
   for the list role binding operation.

   Typically these are written to a http.Request.
*/
type ListRoleBindingParams struct {

	// ClusterID.
	ClusterID string

	// Dc.
	DC string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list role binding params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListRoleBindingParams) WithDefaults() *ListRoleBindingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list role binding params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListRoleBindingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list role binding params
func (o *ListRoleBindingParams) WithTimeout(timeout time.Duration) *ListRoleBindingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list role binding params
func (o *ListRoleBindingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list role binding params
func (o *ListRoleBindingParams) WithContext(ctx context.Context) *ListRoleBindingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list role binding params
func (o *ListRoleBindingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list role binding params
func (o *ListRoleBindingParams) WithHTTPClient(client *http.Client) *ListRoleBindingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list role binding params
func (o *ListRoleBindingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the list role binding params
func (o *ListRoleBindingParams) WithClusterID(clusterID string) *ListRoleBindingParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the list role binding params
func (o *ListRoleBindingParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithDC adds the dc to the list role binding params
func (o *ListRoleBindingParams) WithDC(dc string) *ListRoleBindingParams {
	o.SetDC(dc)
	return o
}

// SetDC adds the dc to the list role binding params
func (o *ListRoleBindingParams) SetDC(dc string) {
	o.DC = dc
}

// WithProjectID adds the projectID to the list role binding params
func (o *ListRoleBindingParams) WithProjectID(projectID string) *ListRoleBindingParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list role binding params
func (o *ListRoleBindingParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListRoleBindingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param dc
	if err := r.SetPathParam("dc", o.DC); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
