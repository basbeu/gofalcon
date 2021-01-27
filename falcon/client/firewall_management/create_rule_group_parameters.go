// Code generated by go-swagger; DO NOT EDIT.

package firewall_management

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

// NewCreateRuleGroupParams creates a new CreateRuleGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRuleGroupParams() *CreateRuleGroupParams {
	return &CreateRuleGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRuleGroupParamsWithTimeout creates a new CreateRuleGroupParams object
// with the ability to set a timeout on a request.
func NewCreateRuleGroupParamsWithTimeout(timeout time.Duration) *CreateRuleGroupParams {
	return &CreateRuleGroupParams{
		timeout: timeout,
	}
}

// NewCreateRuleGroupParamsWithContext creates a new CreateRuleGroupParams object
// with the ability to set a context for a request.
func NewCreateRuleGroupParamsWithContext(ctx context.Context) *CreateRuleGroupParams {
	return &CreateRuleGroupParams{
		Context: ctx,
	}
}

// NewCreateRuleGroupParamsWithHTTPClient creates a new CreateRuleGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRuleGroupParamsWithHTTPClient(client *http.Client) *CreateRuleGroupParams {
	return &CreateRuleGroupParams{
		HTTPClient: client,
	}
}

/* CreateRuleGroupParams contains all the parameters to send to the API endpoint
   for the create rule group operation.

   Typically these are written to a http.Request.
*/
type CreateRuleGroupParams struct {

	/* XCSUSERNAME.

	   The user id
	*/
	XCSUSERNAME string

	// Body.
	Body *models.FwmgrAPIRuleGroupCreateRequestV1

	/* CloneID.

	   A rule group ID from which to copy rules. If this is provided then the 'rules' property of the body is ignored.
	*/
	CloneID *string

	/* Comment.

	   Audit log comment for this action
	*/
	Comment *string

	/* Library.

	   If this flag is set to true then the rules will be cloned from the clone_id from the CrowdStrike Firewal Rule Groups Library.
	*/
	Library *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create rule group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRuleGroupParams) WithDefaults() *CreateRuleGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create rule group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRuleGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create rule group params
func (o *CreateRuleGroupParams) WithTimeout(timeout time.Duration) *CreateRuleGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create rule group params
func (o *CreateRuleGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create rule group params
func (o *CreateRuleGroupParams) WithContext(ctx context.Context) *CreateRuleGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create rule group params
func (o *CreateRuleGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create rule group params
func (o *CreateRuleGroupParams) WithHTTPClient(client *http.Client) *CreateRuleGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create rule group params
func (o *CreateRuleGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXCSUSERNAME adds the xCSUSERNAME to the create rule group params
func (o *CreateRuleGroupParams) WithXCSUSERNAME(xCSUSERNAME string) *CreateRuleGroupParams {
	o.SetXCSUSERNAME(xCSUSERNAME)
	return o
}

// SetXCSUSERNAME adds the xCSUSERNAME to the create rule group params
func (o *CreateRuleGroupParams) SetXCSUSERNAME(xCSUSERNAME string) {
	o.XCSUSERNAME = xCSUSERNAME
}

// WithBody adds the body to the create rule group params
func (o *CreateRuleGroupParams) WithBody(body *models.FwmgrAPIRuleGroupCreateRequestV1) *CreateRuleGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create rule group params
func (o *CreateRuleGroupParams) SetBody(body *models.FwmgrAPIRuleGroupCreateRequestV1) {
	o.Body = body
}

// WithCloneID adds the cloneID to the create rule group params
func (o *CreateRuleGroupParams) WithCloneID(cloneID *string) *CreateRuleGroupParams {
	o.SetCloneID(cloneID)
	return o
}

// SetCloneID adds the cloneId to the create rule group params
func (o *CreateRuleGroupParams) SetCloneID(cloneID *string) {
	o.CloneID = cloneID
}

// WithComment adds the comment to the create rule group params
func (o *CreateRuleGroupParams) WithComment(comment *string) *CreateRuleGroupParams {
	o.SetComment(comment)
	return o
}

// SetComment adds the comment to the create rule group params
func (o *CreateRuleGroupParams) SetComment(comment *string) {
	o.Comment = comment
}

// WithLibrary adds the library to the create rule group params
func (o *CreateRuleGroupParams) WithLibrary(library *string) *CreateRuleGroupParams {
	o.SetLibrary(library)
	return o
}

// SetLibrary adds the library to the create rule group params
func (o *CreateRuleGroupParams) SetLibrary(library *string) {
	o.Library = library
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRuleGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-CS-USERNAME
	if err := r.SetHeaderParam("X-CS-USERNAME", o.XCSUSERNAME); err != nil {
		return err
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.CloneID != nil {

		// query param clone_id
		var qrCloneID string

		if o.CloneID != nil {
			qrCloneID = *o.CloneID
		}
		qCloneID := qrCloneID
		if qCloneID != "" {

			if err := r.SetQueryParam("clone_id", qCloneID); err != nil {
				return err
			}
		}
	}

	if o.Comment != nil {

		// query param comment
		var qrComment string

		if o.Comment != nil {
			qrComment = *o.Comment
		}
		qComment := qrComment
		if qComment != "" {

			if err := r.SetQueryParam("comment", qComment); err != nil {
				return err
			}
		}
	}

	if o.Library != nil {

		// query param library
		var qrLibrary string

		if o.Library != nil {
			qrLibrary = *o.Library
		}
		qLibrary := qrLibrary
		if qLibrary != "" {

			if err := r.SetQueryParam("library", qLibrary); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
