// Code generated by go-swagger; DO NOT EDIT.

package cspm_registration

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

// NewGetCSPMPolicyParams creates a new GetCSPMPolicyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCSPMPolicyParams() *GetCSPMPolicyParams {
	return &GetCSPMPolicyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCSPMPolicyParamsWithTimeout creates a new GetCSPMPolicyParams object
// with the ability to set a timeout on a request.
func NewGetCSPMPolicyParamsWithTimeout(timeout time.Duration) *GetCSPMPolicyParams {
	return &GetCSPMPolicyParams{
		timeout: timeout,
	}
}

// NewGetCSPMPolicyParamsWithContext creates a new GetCSPMPolicyParams object
// with the ability to set a context for a request.
func NewGetCSPMPolicyParamsWithContext(ctx context.Context) *GetCSPMPolicyParams {
	return &GetCSPMPolicyParams{
		Context: ctx,
	}
}

// NewGetCSPMPolicyParamsWithHTTPClient creates a new GetCSPMPolicyParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCSPMPolicyParamsWithHTTPClient(client *http.Client) *GetCSPMPolicyParams {
	return &GetCSPMPolicyParams{
		HTTPClient: client,
	}
}

/* GetCSPMPolicyParams contains all the parameters to send to the API endpoint
   for the get c s p m policy operation.

   Typically these are written to a http.Request.
*/
type GetCSPMPolicyParams struct {

	/* Ids.

	   Policy ID
	*/
	Ids string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get c s p m policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCSPMPolicyParams) WithDefaults() *GetCSPMPolicyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get c s p m policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCSPMPolicyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get c s p m policy params
func (o *GetCSPMPolicyParams) WithTimeout(timeout time.Duration) *GetCSPMPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get c s p m policy params
func (o *GetCSPMPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get c s p m policy params
func (o *GetCSPMPolicyParams) WithContext(ctx context.Context) *GetCSPMPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get c s p m policy params
func (o *GetCSPMPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get c s p m policy params
func (o *GetCSPMPolicyParams) WithHTTPClient(client *http.Client) *GetCSPMPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get c s p m policy params
func (o *GetCSPMPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the get c s p m policy params
func (o *GetCSPMPolicyParams) WithIds(ids string) *GetCSPMPolicyParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the get c s p m policy params
func (o *GetCSPMPolicyParams) SetIds(ids string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *GetCSPMPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param ids
	qrIds := o.Ids
	qIds := qrIds
	if qIds != "" {

		if err := r.SetQueryParam("ids", qIds); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
