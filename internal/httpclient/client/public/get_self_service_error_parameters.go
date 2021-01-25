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

// NewGetSelfServiceErrorParams creates a new GetSelfServiceErrorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSelfServiceErrorParams() *GetSelfServiceErrorParams {
	return &GetSelfServiceErrorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSelfServiceErrorParamsWithTimeout creates a new GetSelfServiceErrorParams object
// with the ability to set a timeout on a request.
func NewGetSelfServiceErrorParamsWithTimeout(timeout time.Duration) *GetSelfServiceErrorParams {
	return &GetSelfServiceErrorParams{
		timeout: timeout,
	}
}

// NewGetSelfServiceErrorParamsWithContext creates a new GetSelfServiceErrorParams object
// with the ability to set a context for a request.
func NewGetSelfServiceErrorParamsWithContext(ctx context.Context) *GetSelfServiceErrorParams {
	return &GetSelfServiceErrorParams{
		Context: ctx,
	}
}

// NewGetSelfServiceErrorParamsWithHTTPClient creates a new GetSelfServiceErrorParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSelfServiceErrorParamsWithHTTPClient(client *http.Client) *GetSelfServiceErrorParams {
	return &GetSelfServiceErrorParams{
		HTTPClient: client,
	}
}

/* GetSelfServiceErrorParams contains all the parameters to send to the API endpoint
   for the get self service error operation.

   Typically these are written to a http.Request.
*/
type GetSelfServiceErrorParams struct {

	/* Error.

	   Error is the container's ID
	*/
	Error string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get self service error params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSelfServiceErrorParams) WithDefaults() *GetSelfServiceErrorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get self service error params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSelfServiceErrorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get self service error params
func (o *GetSelfServiceErrorParams) WithTimeout(timeout time.Duration) *GetSelfServiceErrorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get self service error params
func (o *GetSelfServiceErrorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get self service error params
func (o *GetSelfServiceErrorParams) WithContext(ctx context.Context) *GetSelfServiceErrorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get self service error params
func (o *GetSelfServiceErrorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get self service error params
func (o *GetSelfServiceErrorParams) WithHTTPClient(client *http.Client) *GetSelfServiceErrorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get self service error params
func (o *GetSelfServiceErrorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithError adds the error to the get self service error params
func (o *GetSelfServiceErrorParams) WithError(error string) *GetSelfServiceErrorParams {
	o.SetError(error)
	return o
}

// SetError adds the error to the get self service error params
func (o *GetSelfServiceErrorParams) SetError(error string) {
	o.Error = error
}

// WriteToRequest writes these params to a swagger request
func (o *GetSelfServiceErrorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param error
	qrError := o.Error
	qError := qrError
	if qError != "" {

		if err := r.SetQueryParam("error", qError); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
