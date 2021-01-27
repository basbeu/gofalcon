// Code generated by go-swagger; DO NOT EDIT.

package device_control_policies

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

// QueryCombinedDeviceControlPolicyMembersReader is a Reader for the QueryCombinedDeviceControlPolicyMembers structure.
type QueryCombinedDeviceControlPolicyMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryCombinedDeviceControlPolicyMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryCombinedDeviceControlPolicyMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryCombinedDeviceControlPolicyMembersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryCombinedDeviceControlPolicyMembersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewQueryCombinedDeviceControlPolicyMembersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryCombinedDeviceControlPolicyMembersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryCombinedDeviceControlPolicyMembersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewQueryCombinedDeviceControlPolicyMembersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQueryCombinedDeviceControlPolicyMembersOK creates a QueryCombinedDeviceControlPolicyMembersOK with default headers values
func NewQueryCombinedDeviceControlPolicyMembersOK() *QueryCombinedDeviceControlPolicyMembersOK {
	return &QueryCombinedDeviceControlPolicyMembersOK{}
}

/* QueryCombinedDeviceControlPolicyMembersOK describes a response with status code 200, with default header values.

OK
*/
type QueryCombinedDeviceControlPolicyMembersOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesPolicyMembersRespV1
}

func (o *QueryCombinedDeviceControlPolicyMembersOK) Error() string {
	return fmt.Sprintf("[GET /policy/combined/device-control-members/v1][%d] queryCombinedDeviceControlPolicyMembersOK  %+v", 200, o.Payload)
}
func (o *QueryCombinedDeviceControlPolicyMembersOK) GetPayload() *models.ResponsesPolicyMembersRespV1 {
	return o.Payload
}

func (o *QueryCombinedDeviceControlPolicyMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesPolicyMembersRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCombinedDeviceControlPolicyMembersBadRequest creates a QueryCombinedDeviceControlPolicyMembersBadRequest with default headers values
func NewQueryCombinedDeviceControlPolicyMembersBadRequest() *QueryCombinedDeviceControlPolicyMembersBadRequest {
	return &QueryCombinedDeviceControlPolicyMembersBadRequest{}
}

/* QueryCombinedDeviceControlPolicyMembersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type QueryCombinedDeviceControlPolicyMembersBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesPolicyMembersRespV1
}

func (o *QueryCombinedDeviceControlPolicyMembersBadRequest) Error() string {
	return fmt.Sprintf("[GET /policy/combined/device-control-members/v1][%d] queryCombinedDeviceControlPolicyMembersBadRequest  %+v", 400, o.Payload)
}
func (o *QueryCombinedDeviceControlPolicyMembersBadRequest) GetPayload() *models.ResponsesPolicyMembersRespV1 {
	return o.Payload
}

func (o *QueryCombinedDeviceControlPolicyMembersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesPolicyMembersRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCombinedDeviceControlPolicyMembersForbidden creates a QueryCombinedDeviceControlPolicyMembersForbidden with default headers values
func NewQueryCombinedDeviceControlPolicyMembersForbidden() *QueryCombinedDeviceControlPolicyMembersForbidden {
	return &QueryCombinedDeviceControlPolicyMembersForbidden{}
}

/* QueryCombinedDeviceControlPolicyMembersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryCombinedDeviceControlPolicyMembersForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *QueryCombinedDeviceControlPolicyMembersForbidden) Error() string {
	return fmt.Sprintf("[GET /policy/combined/device-control-members/v1][%d] queryCombinedDeviceControlPolicyMembersForbidden  %+v", 403, o.Payload)
}
func (o *QueryCombinedDeviceControlPolicyMembersForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *QueryCombinedDeviceControlPolicyMembersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCombinedDeviceControlPolicyMembersNotFound creates a QueryCombinedDeviceControlPolicyMembersNotFound with default headers values
func NewQueryCombinedDeviceControlPolicyMembersNotFound() *QueryCombinedDeviceControlPolicyMembersNotFound {
	return &QueryCombinedDeviceControlPolicyMembersNotFound{}
}

/* QueryCombinedDeviceControlPolicyMembersNotFound describes a response with status code 404, with default header values.

Not Found
*/
type QueryCombinedDeviceControlPolicyMembersNotFound struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesPolicyMembersRespV1
}

func (o *QueryCombinedDeviceControlPolicyMembersNotFound) Error() string {
	return fmt.Sprintf("[GET /policy/combined/device-control-members/v1][%d] queryCombinedDeviceControlPolicyMembersNotFound  %+v", 404, o.Payload)
}
func (o *QueryCombinedDeviceControlPolicyMembersNotFound) GetPayload() *models.ResponsesPolicyMembersRespV1 {
	return o.Payload
}

func (o *QueryCombinedDeviceControlPolicyMembersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesPolicyMembersRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCombinedDeviceControlPolicyMembersTooManyRequests creates a QueryCombinedDeviceControlPolicyMembersTooManyRequests with default headers values
func NewQueryCombinedDeviceControlPolicyMembersTooManyRequests() *QueryCombinedDeviceControlPolicyMembersTooManyRequests {
	return &QueryCombinedDeviceControlPolicyMembersTooManyRequests{}
}

/* QueryCombinedDeviceControlPolicyMembersTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryCombinedDeviceControlPolicyMembersTooManyRequests struct {

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

func (o *QueryCombinedDeviceControlPolicyMembersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /policy/combined/device-control-members/v1][%d] queryCombinedDeviceControlPolicyMembersTooManyRequests  %+v", 429, o.Payload)
}
func (o *QueryCombinedDeviceControlPolicyMembersTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryCombinedDeviceControlPolicyMembersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryCombinedDeviceControlPolicyMembersInternalServerError creates a QueryCombinedDeviceControlPolicyMembersInternalServerError with default headers values
func NewQueryCombinedDeviceControlPolicyMembersInternalServerError() *QueryCombinedDeviceControlPolicyMembersInternalServerError {
	return &QueryCombinedDeviceControlPolicyMembersInternalServerError{}
}

/* QueryCombinedDeviceControlPolicyMembersInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type QueryCombinedDeviceControlPolicyMembersInternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesPolicyMembersRespV1
}

func (o *QueryCombinedDeviceControlPolicyMembersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /policy/combined/device-control-members/v1][%d] queryCombinedDeviceControlPolicyMembersInternalServerError  %+v", 500, o.Payload)
}
func (o *QueryCombinedDeviceControlPolicyMembersInternalServerError) GetPayload() *models.ResponsesPolicyMembersRespV1 {
	return o.Payload
}

func (o *QueryCombinedDeviceControlPolicyMembersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesPolicyMembersRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCombinedDeviceControlPolicyMembersDefault creates a QueryCombinedDeviceControlPolicyMembersDefault with default headers values
func NewQueryCombinedDeviceControlPolicyMembersDefault(code int) *QueryCombinedDeviceControlPolicyMembersDefault {
	return &QueryCombinedDeviceControlPolicyMembersDefault{
		_statusCode: code,
	}
}

/* QueryCombinedDeviceControlPolicyMembersDefault describes a response with status code -1, with default header values.

OK
*/
type QueryCombinedDeviceControlPolicyMembersDefault struct {
	_statusCode int

	Payload *models.ResponsesPolicyMembersRespV1
}

// Code gets the status code for the query combined device control policy members default response
func (o *QueryCombinedDeviceControlPolicyMembersDefault) Code() int {
	return o._statusCode
}

func (o *QueryCombinedDeviceControlPolicyMembersDefault) Error() string {
	return fmt.Sprintf("[GET /policy/combined/device-control-members/v1][%d] queryCombinedDeviceControlPolicyMembers default  %+v", o._statusCode, o.Payload)
}
func (o *QueryCombinedDeviceControlPolicyMembersDefault) GetPayload() *models.ResponsesPolicyMembersRespV1 {
	return o.Payload
}

func (o *QueryCombinedDeviceControlPolicyMembersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponsesPolicyMembersRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
