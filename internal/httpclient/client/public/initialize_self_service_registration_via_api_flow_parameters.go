// Code generated by go-swagger; DO NOT EDIT.

package public

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

// NewInitializeSelfServiceRegistrationViaAPIFlowParams creates a new InitializeSelfServiceRegistrationViaAPIFlowParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInitializeSelfServiceRegistrationViaAPIFlowParams() *InitializeSelfServiceRegistrationViaAPIFlowParams {
	return &InitializeSelfServiceRegistrationViaAPIFlowParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInitializeSelfServiceRegistrationViaAPIFlowParamsWithTimeout creates a new InitializeSelfServiceRegistrationViaAPIFlowParams object
// with the ability to set a timeout on a request.
func NewInitializeSelfServiceRegistrationViaAPIFlowParamsWithTimeout(timeout time.Duration) *InitializeSelfServiceRegistrationViaAPIFlowParams {
	return &InitializeSelfServiceRegistrationViaAPIFlowParams{
		timeout: timeout,
	}
}

// NewInitializeSelfServiceRegistrationViaAPIFlowParamsWithContext creates a new InitializeSelfServiceRegistrationViaAPIFlowParams object
// with the ability to set a context for a request.
func NewInitializeSelfServiceRegistrationViaAPIFlowParamsWithContext(ctx context.Context) *InitializeSelfServiceRegistrationViaAPIFlowParams {
	return &InitializeSelfServiceRegistrationViaAPIFlowParams{
		Context: ctx,
	}
}

// NewInitializeSelfServiceRegistrationViaAPIFlowParamsWithHTTPClient creates a new InitializeSelfServiceRegistrationViaAPIFlowParams object
// with the ability to set a custom HTTPClient for a request.
func NewInitializeSelfServiceRegistrationViaAPIFlowParamsWithHTTPClient(client *http.Client) *InitializeSelfServiceRegistrationViaAPIFlowParams {
	return &InitializeSelfServiceRegistrationViaAPIFlowParams{
		HTTPClient: client,
	}
}

/* InitializeSelfServiceRegistrationViaAPIFlowParams contains all the parameters to send to the API endpoint
   for the initialize self service registration via API flow operation.

   Typically these are written to a http.Request.
*/
type InitializeSelfServiceRegistrationViaAPIFlowParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the initialize self service registration via API flow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InitializeSelfServiceRegistrationViaAPIFlowParams) WithDefaults() *InitializeSelfServiceRegistrationViaAPIFlowParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the initialize self service registration via API flow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InitializeSelfServiceRegistrationViaAPIFlowParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the initialize self service registration via API flow params
func (o *InitializeSelfServiceRegistrationViaAPIFlowParams) WithTimeout(timeout time.Duration) *InitializeSelfServiceRegistrationViaAPIFlowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the initialize self service registration via API flow params
func (o *InitializeSelfServiceRegistrationViaAPIFlowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the initialize self service registration via API flow params
func (o *InitializeSelfServiceRegistrationViaAPIFlowParams) WithContext(ctx context.Context) *InitializeSelfServiceRegistrationViaAPIFlowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the initialize self service registration via API flow params
func (o *InitializeSelfServiceRegistrationViaAPIFlowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the initialize self service registration via API flow params
func (o *InitializeSelfServiceRegistrationViaAPIFlowParams) WithHTTPClient(client *http.Client) *InitializeSelfServiceRegistrationViaAPIFlowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the initialize self service registration via API flow params
func (o *InitializeSelfServiceRegistrationViaAPIFlowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *InitializeSelfServiceRegistrationViaAPIFlowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
