// Code generated by go-swagger; DO NOT EDIT.

package devices

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
	"github.com/go-openapi/swag"
)

// NewGetDeviceDetailsV2Params creates a new GetDeviceDetailsV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeviceDetailsV2Params() *GetDeviceDetailsV2Params {
	return &GetDeviceDetailsV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceDetailsV2ParamsWithTimeout creates a new GetDeviceDetailsV2Params object
// with the ability to set a timeout on a request.
func NewGetDeviceDetailsV2ParamsWithTimeout(timeout time.Duration) *GetDeviceDetailsV2Params {
	return &GetDeviceDetailsV2Params{
		timeout: timeout,
	}
}

// NewGetDeviceDetailsV2ParamsWithContext creates a new GetDeviceDetailsV2Params object
// with the ability to set a context for a request.
func NewGetDeviceDetailsV2ParamsWithContext(ctx context.Context) *GetDeviceDetailsV2Params {
	return &GetDeviceDetailsV2Params{
		Context: ctx,
	}
}

// NewGetDeviceDetailsV2ParamsWithHTTPClient creates a new GetDeviceDetailsV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeviceDetailsV2ParamsWithHTTPClient(client *http.Client) *GetDeviceDetailsV2Params {
	return &GetDeviceDetailsV2Params{
		HTTPClient: client,
	}
}

/*
GetDeviceDetailsV2Params contains all the parameters to send to the API endpoint

	for the get device details v2 operation.

	Typically these are written to a http.Request.
*/
type GetDeviceDetailsV2Params struct {

	/* Ids.

	   The host agentIDs used to get details on
	*/
	Ids []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get device details v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceDetailsV2Params) WithDefaults() *GetDeviceDetailsV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get device details v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceDetailsV2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get device details v2 params
func (o *GetDeviceDetailsV2Params) WithTimeout(timeout time.Duration) *GetDeviceDetailsV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device details v2 params
func (o *GetDeviceDetailsV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device details v2 params
func (o *GetDeviceDetailsV2Params) WithContext(ctx context.Context) *GetDeviceDetailsV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device details v2 params
func (o *GetDeviceDetailsV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device details v2 params
func (o *GetDeviceDetailsV2Params) WithHTTPClient(client *http.Client) *GetDeviceDetailsV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device details v2 params
func (o *GetDeviceDetailsV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the get device details v2 params
func (o *GetDeviceDetailsV2Params) WithIds(ids []string) *GetDeviceDetailsV2Params {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the get device details v2 params
func (o *GetDeviceDetailsV2Params) SetIds(ids []string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceDetailsV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Ids != nil {

		// binding items for ids
		joinedIds := o.bindParamIds(reg)

		// query array param ids
		if err := r.SetQueryParam("ids", joinedIds...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetDeviceDetailsV2 binds the parameter ids
func (o *GetDeviceDetailsV2Params) bindParamIds(formats strfmt.Registry) []string {
	idsIR := o.Ids

	var idsIC []string
	for _, idsIIR := range idsIR { // explode []string

		idsIIV := idsIIR // string as string
		idsIC = append(idsIC, idsIIV)
	}

	// items.CollectionFormat: "multi"
	idsIS := swag.JoinByFormat(idsIC, "multi")

	return idsIS
}