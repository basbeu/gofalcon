// Code generated by go-swagger; DO NOT EDIT.

package firewall_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// UpdateRuleGroupValidationReader is a Reader for the UpdateRuleGroupValidation structure.
type UpdateRuleGroupValidationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRuleGroupValidationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRuleGroupValidationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateRuleGroupValidationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRuleGroupValidationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateRuleGroupValidationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateRuleGroupValidationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateRuleGroupValidationOK creates a UpdateRuleGroupValidationOK with default headers values
func NewUpdateRuleGroupValidationOK() *UpdateRuleGroupValidationOK {
	return &UpdateRuleGroupValidationOK{}
}

/*
UpdateRuleGroupValidationOK describes a response with status code 200, with default header values.

OK
*/
type UpdateRuleGroupValidationOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.FwmgrMsaQueryResponse
}

// IsSuccess returns true when this update rule group validation o k response has a 2xx status code
func (o *UpdateRuleGroupValidationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update rule group validation o k response has a 3xx status code
func (o *UpdateRuleGroupValidationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update rule group validation o k response has a 4xx status code
func (o *UpdateRuleGroupValidationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update rule group validation o k response has a 5xx status code
func (o *UpdateRuleGroupValidationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update rule group validation o k response a status code equal to that given
func (o *UpdateRuleGroupValidationOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateRuleGroupValidationOK) Error() string {
	return fmt.Sprintf("[PATCH /fwmgr/entities/rule-groups/validation/v1][%d] updateRuleGroupValidationOK  %+v", 200, o.Payload)
}

func (o *UpdateRuleGroupValidationOK) String() string {
	return fmt.Sprintf("[PATCH /fwmgr/entities/rule-groups/validation/v1][%d] updateRuleGroupValidationOK  %+v", 200, o.Payload)
}

func (o *UpdateRuleGroupValidationOK) GetPayload() *models.FwmgrMsaQueryResponse {
	return o.Payload
}

func (o *UpdateRuleGroupValidationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.FwmgrMsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRuleGroupValidationBadRequest creates a UpdateRuleGroupValidationBadRequest with default headers values
func NewUpdateRuleGroupValidationBadRequest() *UpdateRuleGroupValidationBadRequest {
	return &UpdateRuleGroupValidationBadRequest{}
}

/*
UpdateRuleGroupValidationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateRuleGroupValidationBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.FwmgrMsaReplyMetaOnly
}

// IsSuccess returns true when this update rule group validation bad request response has a 2xx status code
func (o *UpdateRuleGroupValidationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update rule group validation bad request response has a 3xx status code
func (o *UpdateRuleGroupValidationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update rule group validation bad request response has a 4xx status code
func (o *UpdateRuleGroupValidationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update rule group validation bad request response has a 5xx status code
func (o *UpdateRuleGroupValidationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update rule group validation bad request response a status code equal to that given
func (o *UpdateRuleGroupValidationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UpdateRuleGroupValidationBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /fwmgr/entities/rule-groups/validation/v1][%d] updateRuleGroupValidationBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateRuleGroupValidationBadRequest) String() string {
	return fmt.Sprintf("[PATCH /fwmgr/entities/rule-groups/validation/v1][%d] updateRuleGroupValidationBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateRuleGroupValidationBadRequest) GetPayload() *models.FwmgrMsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateRuleGroupValidationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.FwmgrMsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRuleGroupValidationForbidden creates a UpdateRuleGroupValidationForbidden with default headers values
func NewUpdateRuleGroupValidationForbidden() *UpdateRuleGroupValidationForbidden {
	return &UpdateRuleGroupValidationForbidden{}
}

/*
UpdateRuleGroupValidationForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateRuleGroupValidationForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this update rule group validation forbidden response has a 2xx status code
func (o *UpdateRuleGroupValidationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update rule group validation forbidden response has a 3xx status code
func (o *UpdateRuleGroupValidationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update rule group validation forbidden response has a 4xx status code
func (o *UpdateRuleGroupValidationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update rule group validation forbidden response has a 5xx status code
func (o *UpdateRuleGroupValidationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update rule group validation forbidden response a status code equal to that given
func (o *UpdateRuleGroupValidationForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateRuleGroupValidationForbidden) Error() string {
	return fmt.Sprintf("[PATCH /fwmgr/entities/rule-groups/validation/v1][%d] updateRuleGroupValidationForbidden  %+v", 403, o.Payload)
}

func (o *UpdateRuleGroupValidationForbidden) String() string {
	return fmt.Sprintf("[PATCH /fwmgr/entities/rule-groups/validation/v1][%d] updateRuleGroupValidationForbidden  %+v", 403, o.Payload)
}

func (o *UpdateRuleGroupValidationForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateRuleGroupValidationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRuleGroupValidationTooManyRequests creates a UpdateRuleGroupValidationTooManyRequests with default headers values
func NewUpdateRuleGroupValidationTooManyRequests() *UpdateRuleGroupValidationTooManyRequests {
	return &UpdateRuleGroupValidationTooManyRequests{}
}

/*
UpdateRuleGroupValidationTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateRuleGroupValidationTooManyRequests struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	/* Too many requests, retry after this time (as milliseconds since epoch)
	 */
	XRateLimitRetryAfter int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this update rule group validation too many requests response has a 2xx status code
func (o *UpdateRuleGroupValidationTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update rule group validation too many requests response has a 3xx status code
func (o *UpdateRuleGroupValidationTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update rule group validation too many requests response has a 4xx status code
func (o *UpdateRuleGroupValidationTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this update rule group validation too many requests response has a 5xx status code
func (o *UpdateRuleGroupValidationTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this update rule group validation too many requests response a status code equal to that given
func (o *UpdateRuleGroupValidationTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *UpdateRuleGroupValidationTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /fwmgr/entities/rule-groups/validation/v1][%d] updateRuleGroupValidationTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateRuleGroupValidationTooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /fwmgr/entities/rule-groups/validation/v1][%d] updateRuleGroupValidationTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateRuleGroupValidationTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateRuleGroupValidationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	// hydrates response header X-RateLimit-RetryAfter
	hdrXRateLimitRetryAfter := response.GetHeader("X-RateLimit-RetryAfter")

	if hdrXRateLimitRetryAfter != "" {
		valxRateLimitRetryAfter, err := swag.ConvertInt64(hdrXRateLimitRetryAfter)
		if err != nil {
			return errors.InvalidType("X-RateLimit-RetryAfter", "header", "int64", hdrXRateLimitRetryAfter)
		}
		o.XRateLimitRetryAfter = valxRateLimitRetryAfter
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRuleGroupValidationDefault creates a UpdateRuleGroupValidationDefault with default headers values
func NewUpdateRuleGroupValidationDefault(code int) *UpdateRuleGroupValidationDefault {
	return &UpdateRuleGroupValidationDefault{
		_statusCode: code,
	}
}

/*
UpdateRuleGroupValidationDefault describes a response with status code -1, with default header values.

OK
*/
type UpdateRuleGroupValidationDefault struct {
	_statusCode int

	Payload *models.FwmgrMsaQueryResponse
}

// Code gets the status code for the update rule group validation default response
func (o *UpdateRuleGroupValidationDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this update rule group validation default response has a 2xx status code
func (o *UpdateRuleGroupValidationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update rule group validation default response has a 3xx status code
func (o *UpdateRuleGroupValidationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update rule group validation default response has a 4xx status code
func (o *UpdateRuleGroupValidationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update rule group validation default response has a 5xx status code
func (o *UpdateRuleGroupValidationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update rule group validation default response a status code equal to that given
func (o *UpdateRuleGroupValidationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UpdateRuleGroupValidationDefault) Error() string {
	return fmt.Sprintf("[PATCH /fwmgr/entities/rule-groups/validation/v1][%d] update-rule-group-validation default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateRuleGroupValidationDefault) String() string {
	return fmt.Sprintf("[PATCH /fwmgr/entities/rule-groups/validation/v1][%d] update-rule-group-validation default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateRuleGroupValidationDefault) GetPayload() *models.FwmgrMsaQueryResponse {
	return o.Payload
}

func (o *UpdateRuleGroupValidationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FwmgrMsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
