// Code generated by go-swagger; DO NOT EDIT.

package real_time_response

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

// NewRTRDeleteQueuedSessionParams creates a new RTRDeleteQueuedSessionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRTRDeleteQueuedSessionParams() *RTRDeleteQueuedSessionParams {
	return &RTRDeleteQueuedSessionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRTRDeleteQueuedSessionParamsWithTimeout creates a new RTRDeleteQueuedSessionParams object
// with the ability to set a timeout on a request.
func NewRTRDeleteQueuedSessionParamsWithTimeout(timeout time.Duration) *RTRDeleteQueuedSessionParams {
	return &RTRDeleteQueuedSessionParams{
		timeout: timeout,
	}
}

// NewRTRDeleteQueuedSessionParamsWithContext creates a new RTRDeleteQueuedSessionParams object
// with the ability to set a context for a request.
func NewRTRDeleteQueuedSessionParamsWithContext(ctx context.Context) *RTRDeleteQueuedSessionParams {
	return &RTRDeleteQueuedSessionParams{
		Context: ctx,
	}
}

// NewRTRDeleteQueuedSessionParamsWithHTTPClient creates a new RTRDeleteQueuedSessionParams object
// with the ability to set a custom HTTPClient for a request.
func NewRTRDeleteQueuedSessionParamsWithHTTPClient(client *http.Client) *RTRDeleteQueuedSessionParams {
	return &RTRDeleteQueuedSessionParams{
		HTTPClient: client,
	}
}

/* RTRDeleteQueuedSessionParams contains all the parameters to send to the API endpoint
   for the r t r delete queued session operation.

   Typically these are written to a http.Request.
*/
type RTRDeleteQueuedSessionParams struct {

	/* CloudRequestID.

	   Cloud Request ID of the executed command to query
	*/
	CloudRequestID string

	/* SessionID.

	   RTR Session id
	*/
	SessionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the r t r delete queued session params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RTRDeleteQueuedSessionParams) WithDefaults() *RTRDeleteQueuedSessionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the r t r delete queued session params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RTRDeleteQueuedSessionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the r t r delete queued session params
func (o *RTRDeleteQueuedSessionParams) WithTimeout(timeout time.Duration) *RTRDeleteQueuedSessionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the r t r delete queued session params
func (o *RTRDeleteQueuedSessionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the r t r delete queued session params
func (o *RTRDeleteQueuedSessionParams) WithContext(ctx context.Context) *RTRDeleteQueuedSessionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the r t r delete queued session params
func (o *RTRDeleteQueuedSessionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the r t r delete queued session params
func (o *RTRDeleteQueuedSessionParams) WithHTTPClient(client *http.Client) *RTRDeleteQueuedSessionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the r t r delete queued session params
func (o *RTRDeleteQueuedSessionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudRequestID adds the cloudRequestID to the r t r delete queued session params
func (o *RTRDeleteQueuedSessionParams) WithCloudRequestID(cloudRequestID string) *RTRDeleteQueuedSessionParams {
	o.SetCloudRequestID(cloudRequestID)
	return o
}

// SetCloudRequestID adds the cloudRequestId to the r t r delete queued session params
func (o *RTRDeleteQueuedSessionParams) SetCloudRequestID(cloudRequestID string) {
	o.CloudRequestID = cloudRequestID
}

// WithSessionID adds the sessionID to the r t r delete queued session params
func (o *RTRDeleteQueuedSessionParams) WithSessionID(sessionID string) *RTRDeleteQueuedSessionParams {
	o.SetSessionID(sessionID)
	return o
}

// SetSessionID adds the sessionId to the r t r delete queued session params
func (o *RTRDeleteQueuedSessionParams) SetSessionID(sessionID string) {
	o.SessionID = sessionID
}

// WriteToRequest writes these params to a swagger request
func (o *RTRDeleteQueuedSessionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param cloud_request_id
	qrCloudRequestID := o.CloudRequestID
	qCloudRequestID := qrCloudRequestID
	if qCloudRequestID != "" {

		if err := r.SetQueryParam("cloud_request_id", qCloudRequestID); err != nil {
			return err
		}
	}

	// query param session_id
	qrSessionID := o.SessionID
	qSessionID := qrSessionID
	if qSessionID != "" {

		if err := r.SetQueryParam("session_id", qSessionID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
