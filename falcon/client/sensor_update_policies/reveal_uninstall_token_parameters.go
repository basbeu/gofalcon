// Code generated by go-swagger; DO NOT EDIT.

package sensor_update_policies

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

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// NewRevealUninstallTokenParams creates a new RevealUninstallTokenParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRevealUninstallTokenParams() *RevealUninstallTokenParams {
	return &RevealUninstallTokenParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRevealUninstallTokenParamsWithTimeout creates a new RevealUninstallTokenParams object
// with the ability to set a timeout on a request.
func NewRevealUninstallTokenParamsWithTimeout(timeout time.Duration) *RevealUninstallTokenParams {
	return &RevealUninstallTokenParams{
		timeout: timeout,
	}
}

// NewRevealUninstallTokenParamsWithContext creates a new RevealUninstallTokenParams object
// with the ability to set a context for a request.
func NewRevealUninstallTokenParamsWithContext(ctx context.Context) *RevealUninstallTokenParams {
	return &RevealUninstallTokenParams{
		Context: ctx,
	}
}

// NewRevealUninstallTokenParamsWithHTTPClient creates a new RevealUninstallTokenParams object
// with the ability to set a custom HTTPClient for a request.
func NewRevealUninstallTokenParamsWithHTTPClient(client *http.Client) *RevealUninstallTokenParams {
	return &RevealUninstallTokenParams{
		HTTPClient: client,
	}
}

/* RevealUninstallTokenParams contains all the parameters to send to the API endpoint
   for the reveal uninstall token operation.

   Typically these are written to a http.Request.
*/
type RevealUninstallTokenParams struct {

	// Body.
	Body *models.RequestsRevealUninstallTokenV1

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the reveal uninstall token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RevealUninstallTokenParams) WithDefaults() *RevealUninstallTokenParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the reveal uninstall token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RevealUninstallTokenParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the reveal uninstall token params
func (o *RevealUninstallTokenParams) WithTimeout(timeout time.Duration) *RevealUninstallTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reveal uninstall token params
func (o *RevealUninstallTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reveal uninstall token params
func (o *RevealUninstallTokenParams) WithContext(ctx context.Context) *RevealUninstallTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reveal uninstall token params
func (o *RevealUninstallTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reveal uninstall token params
func (o *RevealUninstallTokenParams) WithHTTPClient(client *http.Client) *RevealUninstallTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reveal uninstall token params
func (o *RevealUninstallTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the reveal uninstall token params
func (o *RevealUninstallTokenParams) WithBody(body *models.RequestsRevealUninstallTokenV1) *RevealUninstallTokenParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the reveal uninstall token params
func (o *RevealUninstallTokenParams) SetBody(body *models.RequestsRevealUninstallTokenV1) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RevealUninstallTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
