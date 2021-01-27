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

// NewDeleteDeviceControlPoliciesParams creates a new DeleteDeviceControlPoliciesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteDeviceControlPoliciesParams() *DeleteDeviceControlPoliciesParams {
	return &DeleteDeviceControlPoliciesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDeviceControlPoliciesParamsWithTimeout creates a new DeleteDeviceControlPoliciesParams object
// with the ability to set a timeout on a request.
func NewDeleteDeviceControlPoliciesParamsWithTimeout(timeout time.Duration) *DeleteDeviceControlPoliciesParams {
	return &DeleteDeviceControlPoliciesParams{
		timeout: timeout,
	}
}

// NewDeleteDeviceControlPoliciesParamsWithContext creates a new DeleteDeviceControlPoliciesParams object
// with the ability to set a context for a request.
func NewDeleteDeviceControlPoliciesParamsWithContext(ctx context.Context) *DeleteDeviceControlPoliciesParams {
	return &DeleteDeviceControlPoliciesParams{
		Context: ctx,
	}
}

// NewDeleteDeviceControlPoliciesParamsWithHTTPClient creates a new DeleteDeviceControlPoliciesParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteDeviceControlPoliciesParamsWithHTTPClient(client *http.Client) *DeleteDeviceControlPoliciesParams {
	return &DeleteDeviceControlPoliciesParams{
		HTTPClient: client,
	}
}

/* DeleteDeviceControlPoliciesParams contains all the parameters to send to the API endpoint
   for the delete device control policies operation.

   Typically these are written to a http.Request.
*/
type DeleteDeviceControlPoliciesParams struct {

	/* Ids.

	   The IDs of the Device Control Policies to delete
	*/
	Ids []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete device control policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDeviceControlPoliciesParams) WithDefaults() *DeleteDeviceControlPoliciesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete device control policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDeviceControlPoliciesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete device control policies params
func (o *DeleteDeviceControlPoliciesParams) WithTimeout(timeout time.Duration) *DeleteDeviceControlPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete device control policies params
func (o *DeleteDeviceControlPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete device control policies params
func (o *DeleteDeviceControlPoliciesParams) WithContext(ctx context.Context) *DeleteDeviceControlPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete device control policies params
func (o *DeleteDeviceControlPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete device control policies params
func (o *DeleteDeviceControlPoliciesParams) WithHTTPClient(client *http.Client) *DeleteDeviceControlPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete device control policies params
func (o *DeleteDeviceControlPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the delete device control policies params
func (o *DeleteDeviceControlPoliciesParams) WithIds(ids []string) *DeleteDeviceControlPoliciesParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the delete device control policies params
func (o *DeleteDeviceControlPoliciesParams) SetIds(ids []string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDeviceControlPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamDeleteDeviceControlPolicies binds the parameter ids
func (o *DeleteDeviceControlPoliciesParams) bindParamIds(formats strfmt.Registry) []string {
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
