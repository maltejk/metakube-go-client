// Code generated by go-swagger; DO NOT EDIT.

package constraint

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

// NewPatchDefaultConstraintParams creates a new PatchDefaultConstraintParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchDefaultConstraintParams() *PatchDefaultConstraintParams {
	return &PatchDefaultConstraintParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchDefaultConstraintParamsWithTimeout creates a new PatchDefaultConstraintParams object
// with the ability to set a timeout on a request.
func NewPatchDefaultConstraintParamsWithTimeout(timeout time.Duration) *PatchDefaultConstraintParams {
	return &PatchDefaultConstraintParams{
		timeout: timeout,
	}
}

// NewPatchDefaultConstraintParamsWithContext creates a new PatchDefaultConstraintParams object
// with the ability to set a context for a request.
func NewPatchDefaultConstraintParamsWithContext(ctx context.Context) *PatchDefaultConstraintParams {
	return &PatchDefaultConstraintParams{
		Context: ctx,
	}
}

// NewPatchDefaultConstraintParamsWithHTTPClient creates a new PatchDefaultConstraintParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchDefaultConstraintParamsWithHTTPClient(client *http.Client) *PatchDefaultConstraintParams {
	return &PatchDefaultConstraintParams{
		HTTPClient: client,
	}
}

/* PatchDefaultConstraintParams contains all the parameters to send to the API endpoint
   for the patch default constraint operation.

   Typically these are written to a http.Request.
*/
type PatchDefaultConstraintParams struct {

	// Patch.
	Patch interface{}

	// ConstraintName.
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch default constraint params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchDefaultConstraintParams) WithDefaults() *PatchDefaultConstraintParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch default constraint params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchDefaultConstraintParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch default constraint params
func (o *PatchDefaultConstraintParams) WithTimeout(timeout time.Duration) *PatchDefaultConstraintParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch default constraint params
func (o *PatchDefaultConstraintParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch default constraint params
func (o *PatchDefaultConstraintParams) WithContext(ctx context.Context) *PatchDefaultConstraintParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch default constraint params
func (o *PatchDefaultConstraintParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch default constraint params
func (o *PatchDefaultConstraintParams) WithHTTPClient(client *http.Client) *PatchDefaultConstraintParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch default constraint params
func (o *PatchDefaultConstraintParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPatch adds the patch to the patch default constraint params
func (o *PatchDefaultConstraintParams) WithPatch(patch interface{}) *PatchDefaultConstraintParams {
	o.SetPatch(patch)
	return o
}

// SetPatch adds the patch to the patch default constraint params
func (o *PatchDefaultConstraintParams) SetPatch(patch interface{}) {
	o.Patch = patch
}

// WithName adds the constraintName to the patch default constraint params
func (o *PatchDefaultConstraintParams) WithName(constraintName string) *PatchDefaultConstraintParams {
	o.SetName(constraintName)
	return o
}

// SetName adds the constraintName to the patch default constraint params
func (o *PatchDefaultConstraintParams) SetName(constraintName string) {
	o.Name = constraintName
}

// WriteToRequest writes these params to a swagger request
func (o *PatchDefaultConstraintParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Patch != nil {
		if err := r.SetBodyParam(o.Patch); err != nil {
			return err
		}
	}

	// path param constraint_name
	if err := r.SetPathParam("constraint_name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
