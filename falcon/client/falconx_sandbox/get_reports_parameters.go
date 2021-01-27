// Code generated by go-swagger; DO NOT EDIT.

package falconx_sandbox

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

// NewGetReportsParams creates a new GetReportsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetReportsParams() *GetReportsParams {
	return &GetReportsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetReportsParamsWithTimeout creates a new GetReportsParams object
// with the ability to set a timeout on a request.
func NewGetReportsParamsWithTimeout(timeout time.Duration) *GetReportsParams {
	return &GetReportsParams{
		timeout: timeout,
	}
}

// NewGetReportsParamsWithContext creates a new GetReportsParams object
// with the ability to set a context for a request.
func NewGetReportsParamsWithContext(ctx context.Context) *GetReportsParams {
	return &GetReportsParams{
		Context: ctx,
	}
}

// NewGetReportsParamsWithHTTPClient creates a new GetReportsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetReportsParamsWithHTTPClient(client *http.Client) *GetReportsParams {
	return &GetReportsParams{
		HTTPClient: client,
	}
}

/* GetReportsParams contains all the parameters to send to the API endpoint
   for the get reports operation.

   Typically these are written to a http.Request.
*/
type GetReportsParams struct {

	/* Ids.

	   ID of a report. Find a report ID from the response when submitting a malware sample or search with `/falconx/queries/reports/v1`.
	*/
	Ids []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get reports params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetReportsParams) WithDefaults() *GetReportsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get reports params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetReportsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get reports params
func (o *GetReportsParams) WithTimeout(timeout time.Duration) *GetReportsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get reports params
func (o *GetReportsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get reports params
func (o *GetReportsParams) WithContext(ctx context.Context) *GetReportsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get reports params
func (o *GetReportsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get reports params
func (o *GetReportsParams) WithHTTPClient(client *http.Client) *GetReportsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get reports params
func (o *GetReportsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the get reports params
func (o *GetReportsParams) WithIds(ids []string) *GetReportsParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the get reports params
func (o *GetReportsParams) SetIds(ids []string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *GetReportsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamGetReports binds the parameter ids
func (o *GetReportsParams) bindParamIds(formats strfmt.Registry) []string {
	idsIR := o.Ids

	var idsIC []string
	for _, idsIIR := range idsIR { // explode []string

		idsIIV := idsIIR // string as string
		idsIC = append(idsIC, idsIIV)
	}

	// items.CollectionFormat: "csv"
	idsIS := swag.JoinByFormat(idsIC, "csv")

	return idsIS
}
