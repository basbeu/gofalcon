// Code generated by go-swagger; DO NOT EDIT.

package container_packages

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

// NewReadPackagesCombinedParams creates a new ReadPackagesCombinedParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReadPackagesCombinedParams() *ReadPackagesCombinedParams {
	return &ReadPackagesCombinedParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReadPackagesCombinedParamsWithTimeout creates a new ReadPackagesCombinedParams object
// with the ability to set a timeout on a request.
func NewReadPackagesCombinedParamsWithTimeout(timeout time.Duration) *ReadPackagesCombinedParams {
	return &ReadPackagesCombinedParams{
		timeout: timeout,
	}
}

// NewReadPackagesCombinedParamsWithContext creates a new ReadPackagesCombinedParams object
// with the ability to set a context for a request.
func NewReadPackagesCombinedParamsWithContext(ctx context.Context) *ReadPackagesCombinedParams {
	return &ReadPackagesCombinedParams{
		Context: ctx,
	}
}

// NewReadPackagesCombinedParamsWithHTTPClient creates a new ReadPackagesCombinedParams object
// with the ability to set a custom HTTPClient for a request.
func NewReadPackagesCombinedParamsWithHTTPClient(client *http.Client) *ReadPackagesCombinedParams {
	return &ReadPackagesCombinedParams{
		HTTPClient: client,
	}
}

/*
ReadPackagesCombinedParams contains all the parameters to send to the API endpoint

	for the read packages combined operation.

	Typically these are written to a http.Request.
*/
type ReadPackagesCombinedParams struct {

	/* Filter.

	   Filter packages using a query in Falcon Query Language (FQL). Supported filters:  cid,container_id,cveid,fix_status,image_digest,license,package_name_version,severity,type,vulnerability_count
	*/
	Filter *string

	/* Limit.

	   The upper-bound on the number of records to retrieve.
	*/
	Limit *int64

	/* Offset.

	   The offset from where to begin.
	*/
	Offset *int64

	/* OnlyZeroDayAffected.

	   (true/false) load zero day affected packages, default is false
	*/
	OnlyZeroDayAffected *bool

	/* Sort.

	   The fields to sort the records on. Supported columns:  [license package_name_version type]
	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the read packages combined params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadPackagesCombinedParams) WithDefaults() *ReadPackagesCombinedParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the read packages combined params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadPackagesCombinedParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the read packages combined params
func (o *ReadPackagesCombinedParams) WithTimeout(timeout time.Duration) *ReadPackagesCombinedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read packages combined params
func (o *ReadPackagesCombinedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read packages combined params
func (o *ReadPackagesCombinedParams) WithContext(ctx context.Context) *ReadPackagesCombinedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read packages combined params
func (o *ReadPackagesCombinedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read packages combined params
func (o *ReadPackagesCombinedParams) WithHTTPClient(client *http.Client) *ReadPackagesCombinedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read packages combined params
func (o *ReadPackagesCombinedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the read packages combined params
func (o *ReadPackagesCombinedParams) WithFilter(filter *string) *ReadPackagesCombinedParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the read packages combined params
func (o *ReadPackagesCombinedParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithLimit adds the limit to the read packages combined params
func (o *ReadPackagesCombinedParams) WithLimit(limit *int64) *ReadPackagesCombinedParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the read packages combined params
func (o *ReadPackagesCombinedParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the read packages combined params
func (o *ReadPackagesCombinedParams) WithOffset(offset *int64) *ReadPackagesCombinedParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the read packages combined params
func (o *ReadPackagesCombinedParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOnlyZeroDayAffected adds the onlyZeroDayAffected to the read packages combined params
func (o *ReadPackagesCombinedParams) WithOnlyZeroDayAffected(onlyZeroDayAffected *bool) *ReadPackagesCombinedParams {
	o.SetOnlyZeroDayAffected(onlyZeroDayAffected)
	return o
}

// SetOnlyZeroDayAffected adds the onlyZeroDayAffected to the read packages combined params
func (o *ReadPackagesCombinedParams) SetOnlyZeroDayAffected(onlyZeroDayAffected *bool) {
	o.OnlyZeroDayAffected = onlyZeroDayAffected
}

// WithSort adds the sort to the read packages combined params
func (o *ReadPackagesCombinedParams) WithSort(sort *string) *ReadPackagesCombinedParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the read packages combined params
func (o *ReadPackagesCombinedParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *ReadPackagesCombinedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.OnlyZeroDayAffected != nil {

		// query param only_zero_day_affected
		var qrOnlyZeroDayAffected bool

		if o.OnlyZeroDayAffected != nil {
			qrOnlyZeroDayAffected = *o.OnlyZeroDayAffected
		}
		qOnlyZeroDayAffected := swag.FormatBool(qrOnlyZeroDayAffected)
		if qOnlyZeroDayAffected != "" {

			if err := r.SetQueryParam("only_zero_day_affected", qOnlyZeroDayAffected); err != nil {
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