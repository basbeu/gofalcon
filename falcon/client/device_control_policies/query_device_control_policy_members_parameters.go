// Code generated by go-swagger; DO NOT EDIT.

package device_control_policies

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

// NewQueryDeviceControlPolicyMembersParams creates a new QueryDeviceControlPolicyMembersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQueryDeviceControlPolicyMembersParams() *QueryDeviceControlPolicyMembersParams {
	return &QueryDeviceControlPolicyMembersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQueryDeviceControlPolicyMembersParamsWithTimeout creates a new QueryDeviceControlPolicyMembersParams object
// with the ability to set a timeout on a request.
func NewQueryDeviceControlPolicyMembersParamsWithTimeout(timeout time.Duration) *QueryDeviceControlPolicyMembersParams {
	return &QueryDeviceControlPolicyMembersParams{
		timeout: timeout,
	}
}

// NewQueryDeviceControlPolicyMembersParamsWithContext creates a new QueryDeviceControlPolicyMembersParams object
// with the ability to set a context for a request.
func NewQueryDeviceControlPolicyMembersParamsWithContext(ctx context.Context) *QueryDeviceControlPolicyMembersParams {
	return &QueryDeviceControlPolicyMembersParams{
		Context: ctx,
	}
}

// NewQueryDeviceControlPolicyMembersParamsWithHTTPClient creates a new QueryDeviceControlPolicyMembersParams object
// with the ability to set a custom HTTPClient for a request.
func NewQueryDeviceControlPolicyMembersParamsWithHTTPClient(client *http.Client) *QueryDeviceControlPolicyMembersParams {
	return &QueryDeviceControlPolicyMembersParams{
		HTTPClient: client,
	}
}

/* QueryDeviceControlPolicyMembersParams contains all the parameters to send to the API endpoint
   for the query device control policy members operation.

   Typically these are written to a http.Request.
*/
type QueryDeviceControlPolicyMembersParams struct {

	/* Filter.

	   The filter expression that should be used to limit the results
	*/
	Filter *string

	/* ID.

	   The ID of the Device Control Policy to search for members of
	*/
	ID *string

	/* Limit.

	   The maximum records to return. [1-5000]
	*/
	Limit *int64

	/* Offset.

	   The offset to start retrieving records from
	*/
	Offset *int64

	/* Sort.

	   The property to sort by
	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the query device control policy members params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryDeviceControlPolicyMembersParams) WithDefaults() *QueryDeviceControlPolicyMembersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the query device control policy members params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryDeviceControlPolicyMembersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the query device control policy members params
func (o *QueryDeviceControlPolicyMembersParams) WithTimeout(timeout time.Duration) *QueryDeviceControlPolicyMembersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query device control policy members params
func (o *QueryDeviceControlPolicyMembersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query device control policy members params
func (o *QueryDeviceControlPolicyMembersParams) WithContext(ctx context.Context) *QueryDeviceControlPolicyMembersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query device control policy members params
func (o *QueryDeviceControlPolicyMembersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query device control policy members params
func (o *QueryDeviceControlPolicyMembersParams) WithHTTPClient(client *http.Client) *QueryDeviceControlPolicyMembersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query device control policy members params
func (o *QueryDeviceControlPolicyMembersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the query device control policy members params
func (o *QueryDeviceControlPolicyMembersParams) WithFilter(filter *string) *QueryDeviceControlPolicyMembersParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the query device control policy members params
func (o *QueryDeviceControlPolicyMembersParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithID adds the id to the query device control policy members params
func (o *QueryDeviceControlPolicyMembersParams) WithID(id *string) *QueryDeviceControlPolicyMembersParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the query device control policy members params
func (o *QueryDeviceControlPolicyMembersParams) SetID(id *string) {
	o.ID = id
}

// WithLimit adds the limit to the query device control policy members params
func (o *QueryDeviceControlPolicyMembersParams) WithLimit(limit *int64) *QueryDeviceControlPolicyMembersParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the query device control policy members params
func (o *QueryDeviceControlPolicyMembersParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the query device control policy members params
func (o *QueryDeviceControlPolicyMembersParams) WithOffset(offset *int64) *QueryDeviceControlPolicyMembersParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the query device control policy members params
func (o *QueryDeviceControlPolicyMembersParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithSort adds the sort to the query device control policy members params
func (o *QueryDeviceControlPolicyMembersParams) WithSort(sort *string) *QueryDeviceControlPolicyMembersParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the query device control policy members params
func (o *QueryDeviceControlPolicyMembersParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *QueryDeviceControlPolicyMembersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ID != nil {

		// query param id
		var qrID string

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
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
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
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
