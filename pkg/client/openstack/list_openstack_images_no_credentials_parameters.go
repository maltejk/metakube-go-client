// Code generated by go-swagger; DO NOT EDIT.

package openstack

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

// NewListOpenstackImagesNoCredentialsParams creates a new ListOpenstackImagesNoCredentialsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListOpenstackImagesNoCredentialsParams() *ListOpenstackImagesNoCredentialsParams {
	return &ListOpenstackImagesNoCredentialsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListOpenstackImagesNoCredentialsParamsWithTimeout creates a new ListOpenstackImagesNoCredentialsParams object
// with the ability to set a timeout on a request.
func NewListOpenstackImagesNoCredentialsParamsWithTimeout(timeout time.Duration) *ListOpenstackImagesNoCredentialsParams {
	return &ListOpenstackImagesNoCredentialsParams{
		timeout: timeout,
	}
}

// NewListOpenstackImagesNoCredentialsParamsWithContext creates a new ListOpenstackImagesNoCredentialsParams object
// with the ability to set a context for a request.
func NewListOpenstackImagesNoCredentialsParamsWithContext(ctx context.Context) *ListOpenstackImagesNoCredentialsParams {
	return &ListOpenstackImagesNoCredentialsParams{
		Context: ctx,
	}
}

// NewListOpenstackImagesNoCredentialsParamsWithHTTPClient creates a new ListOpenstackImagesNoCredentialsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListOpenstackImagesNoCredentialsParamsWithHTTPClient(client *http.Client) *ListOpenstackImagesNoCredentialsParams {
	return &ListOpenstackImagesNoCredentialsParams{
		HTTPClient: client,
	}
}

/* ListOpenstackImagesNoCredentialsParams contains all the parameters to send to the API endpoint
   for the list openstack images no credentials operation.

   Typically these are written to a http.Request.
*/
type ListOpenstackImagesNoCredentialsParams struct {

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

// WithDefaults hydrates default values in the list openstack images no credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListOpenstackImagesNoCredentialsParams) WithDefaults() *ListOpenstackImagesNoCredentialsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list openstack images no credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListOpenstackImagesNoCredentialsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list openstack images no credentials params
func (o *ListOpenstackImagesNoCredentialsParams) WithTimeout(timeout time.Duration) *ListOpenstackImagesNoCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list openstack images no credentials params
func (o *ListOpenstackImagesNoCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list openstack images no credentials params
func (o *ListOpenstackImagesNoCredentialsParams) WithContext(ctx context.Context) *ListOpenstackImagesNoCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list openstack images no credentials params
func (o *ListOpenstackImagesNoCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list openstack images no credentials params
func (o *ListOpenstackImagesNoCredentialsParams) WithHTTPClient(client *http.Client) *ListOpenstackImagesNoCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list openstack images no credentials params
func (o *ListOpenstackImagesNoCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the list openstack images no credentials params
func (o *ListOpenstackImagesNoCredentialsParams) WithClusterID(clusterID string) *ListOpenstackImagesNoCredentialsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the list openstack images no credentials params
func (o *ListOpenstackImagesNoCredentialsParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithDC adds the dc to the list openstack images no credentials params
func (o *ListOpenstackImagesNoCredentialsParams) WithDC(dc string) *ListOpenstackImagesNoCredentialsParams {
	o.SetDC(dc)
	return o
}

// SetDC adds the dc to the list openstack images no credentials params
func (o *ListOpenstackImagesNoCredentialsParams) SetDC(dc string) {
	o.DC = dc
}

// WithProjectID adds the projectID to the list openstack images no credentials params
func (o *ListOpenstackImagesNoCredentialsParams) WithProjectID(projectID string) *ListOpenstackImagesNoCredentialsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list openstack images no credentials params
func (o *ListOpenstackImagesNoCredentialsParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListOpenstackImagesNoCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
