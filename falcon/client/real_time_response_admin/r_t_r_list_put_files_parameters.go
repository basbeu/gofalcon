// Code generated by go-swagger; DO NOT EDIT.

package real_time_response_admin

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

// NewRTRListPutFilesParams creates a new RTRListPutFilesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRTRListPutFilesParams() *RTRListPutFilesParams {
	return &RTRListPutFilesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRTRListPutFilesParamsWithTimeout creates a new RTRListPutFilesParams object
// with the ability to set a timeout on a request.
func NewRTRListPutFilesParamsWithTimeout(timeout time.Duration) *RTRListPutFilesParams {
	return &RTRListPutFilesParams{
		timeout: timeout,
	}
}

// NewRTRListPutFilesParamsWithContext creates a new RTRListPutFilesParams object
// with the ability to set a context for a request.
func NewRTRListPutFilesParamsWithContext(ctx context.Context) *RTRListPutFilesParams {
	return &RTRListPutFilesParams{
		Context: ctx,
	}
}

// NewRTRListPutFilesParamsWithHTTPClient creates a new RTRListPutFilesParams object
// with the ability to set a custom HTTPClient for a request.
func NewRTRListPutFilesParamsWithHTTPClient(client *http.Client) *RTRListPutFilesParams {
	return &RTRListPutFilesParams{
		HTTPClient: client,
	}
}

/* RTRListPutFilesParams contains all the parameters to send to the API endpoint
   for the r t r list put files operation.

   Typically these are written to a http.Request.
*/
type RTRListPutFilesParams struct {

	/* Filter.

	   Optional filter criteria in the form of an FQL query. For more information about FQL queries, see our [FQL documentation in Falcon](https://falcon.crowdstrike.com/support/documentation/45/falcon-query-language-feature-guide).
	*/
	Filter *string

	/* Limit.

	   Number of ids to return.
	*/
	Limit *int64

	/* Offset.

	   Starting index of overall result set from which to return ids.
	*/
	Offset *string

	/* Sort.

	   Sort by spec. Ex: 'created_at|asc'.
	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the r t r list put files params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RTRListPutFilesParams) WithDefaults() *RTRListPutFilesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the r t r list put files params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RTRListPutFilesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the r t r list put files params
func (o *RTRListPutFilesParams) WithTimeout(timeout time.Duration) *RTRListPutFilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the r t r list put files params
func (o *RTRListPutFilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the r t r list put files params
func (o *RTRListPutFilesParams) WithContext(ctx context.Context) *RTRListPutFilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the r t r list put files params
func (o *RTRListPutFilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the r t r list put files params
func (o *RTRListPutFilesParams) WithHTTPClient(client *http.Client) *RTRListPutFilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the r t r list put files params
func (o *RTRListPutFilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the r t r list put files params
func (o *RTRListPutFilesParams) WithFilter(filter *string) *RTRListPutFilesParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the r t r list put files params
func (o *RTRListPutFilesParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithLimit adds the limit to the r t r list put files params
func (o *RTRListPutFilesParams) WithLimit(limit *int64) *RTRListPutFilesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the r t r list put files params
func (o *RTRListPutFilesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the r t r list put files params
func (o *RTRListPutFilesParams) WithOffset(offset *string) *RTRListPutFilesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the r t r list put files params
func (o *RTRListPutFilesParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithSort adds the sort to the r t r list put files params
func (o *RTRListPutFilesParams) WithSort(sort *string) *RTRListPutFilesParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the r t r list put files params
func (o *RTRListPutFilesParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *RTRListPutFilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filter != nil {

		// query param filter
		var qrFilter string

		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {

			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset string

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := qrOffset
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
