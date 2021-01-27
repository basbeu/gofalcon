// Code generated by go-swagger; DO NOT EDIT.

package real_time_response

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

// RTRCheckCommandStatusReader is a Reader for the RTRCheckCommandStatus structure.
type RTRCheckCommandStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RTRCheckCommandStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRTRCheckCommandStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRTRCheckCommandStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRTRCheckCommandStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewRTRCheckCommandStatusTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewRTRCheckCommandStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRTRCheckCommandStatusOK creates a RTRCheckCommandStatusOK with default headers values
func NewRTRCheckCommandStatusOK() *RTRCheckCommandStatusOK {
	return &RTRCheckCommandStatusOK{}
}

/* RTRCheckCommandStatusOK describes a response with status code 200, with default header values.

success
*/
type RTRCheckCommandStatusOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainStatusResponseWrapper
}

func (o *RTRCheckCommandStatusOK) Error() string {
	return fmt.Sprintf("[GET /real-time-response/entities/command/v1][%d] rTRCheckCommandStatusOK  %+v", 200, o.Payload)
}
func (o *RTRCheckCommandStatusOK) GetPayload() *models.DomainStatusResponseWrapper {
	return o.Payload
}

func (o *RTRCheckCommandStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainStatusResponseWrapper)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRCheckCommandStatusUnauthorized creates a RTRCheckCommandStatusUnauthorized with default headers values
func NewRTRCheckCommandStatusUnauthorized() *RTRCheckCommandStatusUnauthorized {
	return &RTRCheckCommandStatusUnauthorized{}
}

/* RTRCheckCommandStatusUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type RTRCheckCommandStatusUnauthorized struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAPIError
}

func (o *RTRCheckCommandStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /real-time-response/entities/command/v1][%d] rTRCheckCommandStatusUnauthorized  %+v", 401, o.Payload)
}
func (o *RTRCheckCommandStatusUnauthorized) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *RTRCheckCommandStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRCheckCommandStatusForbidden creates a RTRCheckCommandStatusForbidden with default headers values
func NewRTRCheckCommandStatusForbidden() *RTRCheckCommandStatusForbidden {
	return &RTRCheckCommandStatusForbidden{}
}

/* RTRCheckCommandStatusForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RTRCheckCommandStatusForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *RTRCheckCommandStatusForbidden) Error() string {
	return fmt.Sprintf("[GET /real-time-response/entities/command/v1][%d] rTRCheckCommandStatusForbidden  %+v", 403, o.Payload)
}
func (o *RTRCheckCommandStatusForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRCheckCommandStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRCheckCommandStatusTooManyRequests creates a RTRCheckCommandStatusTooManyRequests with default headers values
func NewRTRCheckCommandStatusTooManyRequests() *RTRCheckCommandStatusTooManyRequests {
	return &RTRCheckCommandStatusTooManyRequests{}
}

/* RTRCheckCommandStatusTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type RTRCheckCommandStatusTooManyRequests struct {

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

func (o *RTRCheckCommandStatusTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /real-time-response/entities/command/v1][%d] rTRCheckCommandStatusTooManyRequests  %+v", 429, o.Payload)
}
func (o *RTRCheckCommandStatusTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRCheckCommandStatusTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRCheckCommandStatusDefault creates a RTRCheckCommandStatusDefault with default headers values
func NewRTRCheckCommandStatusDefault(code int) *RTRCheckCommandStatusDefault {
	return &RTRCheckCommandStatusDefault{
		_statusCode: code,
	}
}

/* RTRCheckCommandStatusDefault describes a response with status code -1, with default header values.

success
*/
type RTRCheckCommandStatusDefault struct {
	_statusCode int

	Payload *models.DomainStatusResponseWrapper
}

// Code gets the status code for the r t r check command status default response
func (o *RTRCheckCommandStatusDefault) Code() int {
	return o._statusCode
}

func (o *RTRCheckCommandStatusDefault) Error() string {
	return fmt.Sprintf("[GET /real-time-response/entities/command/v1][%d] RTR-CheckCommandStatus default  %+v", o._statusCode, o.Payload)
}
func (o *RTRCheckCommandStatusDefault) GetPayload() *models.DomainStatusResponseWrapper {
	return o.Payload
}

func (o *RTRCheckCommandStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainStatusResponseWrapper)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
