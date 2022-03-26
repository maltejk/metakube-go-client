// Code generated by go-swagger; DO NOT EDIT.

package azure

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

// NewListAzureSizesParams creates a new ListAzureSizesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListAzureSizesParams() *ListAzureSizesParams {
	return &ListAzureSizesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListAzureSizesParamsWithTimeout creates a new ListAzureSizesParams object
// with the ability to set a timeout on a request.
func NewListAzureSizesParamsWithTimeout(timeout time.Duration) *ListAzureSizesParams {
	return &ListAzureSizesParams{
		timeout: timeout,
	}
}

// NewListAzureSizesParamsWithContext creates a new ListAzureSizesParams object
// with the ability to set a context for a request.
func NewListAzureSizesParamsWithContext(ctx context.Context) *ListAzureSizesParams {
	return &ListAzureSizesParams{
		Context: ctx,
	}
}

// NewListAzureSizesParamsWithHTTPClient creates a new ListAzureSizesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListAzureSizesParamsWithHTTPClient(client *http.Client) *ListAzureSizesParams {
	return &ListAzureSizesParams{
		HTTPClient: client,
	}
}

/* ListAzureSizesParams contains all the parameters to send to the API endpoint
   for the list azure sizes operation.

   Typically these are written to a http.Request.
*/
type ListAzureSizesParams struct {

	// ClientID.
	ClientID *string

	// ClientSecret.
	ClientSecret *string

	// Credential.
	Credential *string

	// Location.
	Location *string

	// SubscriptionID.
	SubscriptionID *string

	// TenantID.
	TenantID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list azure sizes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListAzureSizesParams) WithDefaults() *ListAzureSizesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list azure sizes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListAzureSizesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list azure sizes params
func (o *ListAzureSizesParams) WithTimeout(timeout time.Duration) *ListAzureSizesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list azure sizes params
func (o *ListAzureSizesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list azure sizes params
func (o *ListAzureSizesParams) WithContext(ctx context.Context) *ListAzureSizesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list azure sizes params
func (o *ListAzureSizesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list azure sizes params
func (o *ListAzureSizesParams) WithHTTPClient(client *http.Client) *ListAzureSizesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list azure sizes params
func (o *ListAzureSizesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientID adds the clientID to the list azure sizes params
func (o *ListAzureSizesParams) WithClientID(clientID *string) *ListAzureSizesParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the list azure sizes params
func (o *ListAzureSizesParams) SetClientID(clientID *string) {
	o.ClientID = clientID
}

// WithClientSecret adds the clientSecret to the list azure sizes params
func (o *ListAzureSizesParams) WithClientSecret(clientSecret *string) *ListAzureSizesParams {
	o.SetClientSecret(clientSecret)
	return o
}

// SetClientSecret adds the clientSecret to the list azure sizes params
func (o *ListAzureSizesParams) SetClientSecret(clientSecret *string) {
	o.ClientSecret = clientSecret
}

// WithCredential adds the credential to the list azure sizes params
func (o *ListAzureSizesParams) WithCredential(credential *string) *ListAzureSizesParams {
	o.SetCredential(credential)
	return o
}

// SetCredential adds the credential to the list azure sizes params
func (o *ListAzureSizesParams) SetCredential(credential *string) {
	o.Credential = credential
}

// WithLocation adds the location to the list azure sizes params
func (o *ListAzureSizesParams) WithLocation(location *string) *ListAzureSizesParams {
	o.SetLocation(location)
	return o
}

// SetLocation adds the location to the list azure sizes params
func (o *ListAzureSizesParams) SetLocation(location *string) {
	o.Location = location
}

// WithSubscriptionID adds the subscriptionID to the list azure sizes params
func (o *ListAzureSizesParams) WithSubscriptionID(subscriptionID *string) *ListAzureSizesParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the list azure sizes params
func (o *ListAzureSizesParams) SetSubscriptionID(subscriptionID *string) {
	o.SubscriptionID = subscriptionID
}

// WithTenantID adds the tenantID to the list azure sizes params
func (o *ListAzureSizesParams) WithTenantID(tenantID *string) *ListAzureSizesParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the list azure sizes params
func (o *ListAzureSizesParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *ListAzureSizesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClientID != nil {

		// header param ClientID
		if err := r.SetHeaderParam("ClientID", *o.ClientID); err != nil {
			return err
		}
	}

	if o.ClientSecret != nil {

		// header param ClientSecret
		if err := r.SetHeaderParam("ClientSecret", *o.ClientSecret); err != nil {
			return err
		}
	}

	if o.Credential != nil {

		// header param Credential
		if err := r.SetHeaderParam("Credential", *o.Credential); err != nil {
			return err
		}
	}

	if o.Location != nil {

		// header param Location
		if err := r.SetHeaderParam("Location", *o.Location); err != nil {
			return err
		}
	}

	if o.SubscriptionID != nil {

		// header param SubscriptionID
		if err := r.SetHeaderParam("SubscriptionID", *o.SubscriptionID); err != nil {
			return err
		}
	}

	if o.TenantID != nil {

		// header param TenantID
		if err := r.SetHeaderParam("TenantID", *o.TenantID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
