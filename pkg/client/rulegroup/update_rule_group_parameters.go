// Code generated by go-swagger; DO NOT EDIT.

package rulegroup

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

// NewUpdateRuleGroupParams creates a new UpdateRuleGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateRuleGroupParams() *UpdateRuleGroupParams {
	return &UpdateRuleGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRuleGroupParamsWithTimeout creates a new UpdateRuleGroupParams object
// with the ability to set a timeout on a request.
func NewUpdateRuleGroupParamsWithTimeout(timeout time.Duration) *UpdateRuleGroupParams {
	return &UpdateRuleGroupParams{
		timeout: timeout,
	}
}

// NewUpdateRuleGroupParamsWithContext creates a new UpdateRuleGroupParams object
// with the ability to set a context for a request.
func NewUpdateRuleGroupParamsWithContext(ctx context.Context) *UpdateRuleGroupParams {
	return &UpdateRuleGroupParams{
		Context: ctx,
	}
}

// NewUpdateRuleGroupParamsWithHTTPClient creates a new UpdateRuleGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateRuleGroupParamsWithHTTPClient(client *http.Client) *UpdateRuleGroupParams {
	return &UpdateRuleGroupParams{
		HTTPClient: client,
	}
}

/* UpdateRuleGroupParams contains all the parameters to send to the API endpoint
   for the update rule group operation.

   Typically these are written to a http.Request.
*/
type UpdateRuleGroupParams struct {

	// Body.
	Body *models.RuleGroup

	// ClusterID.
	ClusterID string

	// ProjectID.
	ProjectID string

	// RulegroupID.
	RuleGroupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update rule group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRuleGroupParams) WithDefaults() *UpdateRuleGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update rule group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRuleGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update rule group params
func (o *UpdateRuleGroupParams) WithTimeout(timeout time.Duration) *UpdateRuleGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update rule group params
func (o *UpdateRuleGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update rule group params
func (o *UpdateRuleGroupParams) WithContext(ctx context.Context) *UpdateRuleGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update rule group params
func (o *UpdateRuleGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update rule group params
func (o *UpdateRuleGroupParams) WithHTTPClient(client *http.Client) *UpdateRuleGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update rule group params
func (o *UpdateRuleGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update rule group params
func (o *UpdateRuleGroupParams) WithBody(body *models.RuleGroup) *UpdateRuleGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update rule group params
func (o *UpdateRuleGroupParams) SetBody(body *models.RuleGroup) {
	o.Body = body
}

// WithClusterID adds the clusterID to the update rule group params
func (o *UpdateRuleGroupParams) WithClusterID(clusterID string) *UpdateRuleGroupParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the update rule group params
func (o *UpdateRuleGroupParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithProjectID adds the projectID to the update rule group params
func (o *UpdateRuleGroupParams) WithProjectID(projectID string) *UpdateRuleGroupParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the update rule group params
func (o *UpdateRuleGroupParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WithRuleGroupID adds the rulegroupID to the update rule group params
func (o *UpdateRuleGroupParams) WithRuleGroupID(rulegroupID string) *UpdateRuleGroupParams {
	o.SetRuleGroupID(rulegroupID)
	return o
}

// SetRuleGroupID adds the rulegroupId to the update rule group params
func (o *UpdateRuleGroupParams) SetRuleGroupID(rulegroupID string) {
	o.RuleGroupID = rulegroupID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRuleGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	// path param rulegroup_id
	if err := r.SetPathParam("rulegroup_id", o.RuleGroupID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
