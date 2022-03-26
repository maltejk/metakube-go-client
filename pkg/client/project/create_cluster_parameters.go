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

	"github.com/maltejk/metakube-go-client/pkg/models"
)

// NewCreateClusterParams creates a new CreateClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateClusterParams() *CreateClusterParams {
	return &CreateClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateClusterParamsWithTimeout creates a new CreateClusterParams object
// with the ability to set a timeout on a request.
func NewCreateClusterParamsWithTimeout(timeout time.Duration) *CreateClusterParams {
	return &CreateClusterParams{
		timeout: timeout,
	}
}

// NewCreateClusterParamsWithContext creates a new CreateClusterParams object
// with the ability to set a context for a request.
func NewCreateClusterParamsWithContext(ctx context.Context) *CreateClusterParams {
	return &CreateClusterParams{
		Context: ctx,
	}
}

// NewCreateClusterParamsWithHTTPClient creates a new CreateClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateClusterParamsWithHTTPClient(client *http.Client) *CreateClusterParams {
	return &CreateClusterParams{
		HTTPClient: client,
	}
}

/* CreateClusterParams contains all the parameters to send to the API endpoint
   for the create cluster operation.

   Typically these are written to a http.Request.
*/
type CreateClusterParams struct {

	// Body.
	Body *models.CreateClusterSpec

	// Dc.
	DC string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateClusterParams) WithDefaults() *CreateClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create cluster params
func (o *CreateClusterParams) WithTimeout(timeout time.Duration) *CreateClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create cluster params
func (o *CreateClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create cluster params
func (o *CreateClusterParams) WithContext(ctx context.Context) *CreateClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create cluster params
func (o *CreateClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create cluster params
func (o *CreateClusterParams) WithHTTPClient(client *http.Client) *CreateClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create cluster params
func (o *CreateClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create cluster params
func (o *CreateClusterParams) WithBody(body *models.CreateClusterSpec) *CreateClusterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create cluster params
func (o *CreateClusterParams) SetBody(body *models.CreateClusterSpec) {
	o.Body = body
}

// WithDC adds the dc to the create cluster params
func (o *CreateClusterParams) WithDC(dc string) *CreateClusterParams {
	o.SetDC(dc)
	return o
}

// SetDC adds the dc to the create cluster params
func (o *CreateClusterParams) SetDC(dc string) {
	o.DC = dc
}

// WithProjectID adds the projectID to the create cluster params
func (o *CreateClusterParams) WithProjectID(projectID string) *CreateClusterParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the create cluster params
func (o *CreateClusterParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
