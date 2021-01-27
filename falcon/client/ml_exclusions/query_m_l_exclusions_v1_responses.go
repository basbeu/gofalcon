// Code generated by go-swagger; DO NOT EDIT.

package ml_exclusions

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

// QueryMLExclusionsV1Reader is a Reader for the QueryMLExclusionsV1 structure.
type QueryMLExclusionsV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryMLExclusionsV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryMLExclusionsV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryMLExclusionsV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryMLExclusionsV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryMLExclusionsV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryMLExclusionsV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewQueryMLExclusionsV1Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQueryMLExclusionsV1OK creates a QueryMLExclusionsV1OK with default headers values
func NewQueryMLExclusionsV1OK() *QueryMLExclusionsV1OK {
	return &QueryMLExclusionsV1OK{}
}

/* QueryMLExclusionsV1OK describes a response with status code 200, with default header values.

OK
*/
type QueryMLExclusionsV1OK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

func (o *QueryMLExclusionsV1OK) Error() string {
	return fmt.Sprintf("[GET /policy/queries/ml-exclusions/v1][%d] queryMLExclusionsV1OK  %+v", 200, o.Payload)
}
func (o *QueryMLExclusionsV1OK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryMLExclusionsV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryMLExclusionsV1BadRequest creates a QueryMLExclusionsV1BadRequest with default headers values
func NewQueryMLExclusionsV1BadRequest() *QueryMLExclusionsV1BadRequest {
	return &QueryMLExclusionsV1BadRequest{}
}

/* QueryMLExclusionsV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type QueryMLExclusionsV1BadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

func (o *QueryMLExclusionsV1BadRequest) Error() string {
	return fmt.Sprintf("[GET /policy/queries/ml-exclusions/v1][%d] queryMLExclusionsV1BadRequest  %+v", 400, o.Payload)
}
func (o *QueryMLExclusionsV1BadRequest) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryMLExclusionsV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryMLExclusionsV1Forbidden creates a QueryMLExclusionsV1Forbidden with default headers values
func NewQueryMLExclusionsV1Forbidden() *QueryMLExclusionsV1Forbidden {
	return &QueryMLExclusionsV1Forbidden{}
}

/* QueryMLExclusionsV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryMLExclusionsV1Forbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *QueryMLExclusionsV1Forbidden) Error() string {
	return fmt.Sprintf("[GET /policy/queries/ml-exclusions/v1][%d] queryMLExclusionsV1Forbidden  %+v", 403, o.Payload)
}
func (o *QueryMLExclusionsV1Forbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *QueryMLExclusionsV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryMLExclusionsV1TooManyRequests creates a QueryMLExclusionsV1TooManyRequests with default headers values
func NewQueryMLExclusionsV1TooManyRequests() *QueryMLExclusionsV1TooManyRequests {
	return &QueryMLExclusionsV1TooManyRequests{}
}

/* QueryMLExclusionsV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryMLExclusionsV1TooManyRequests struct {

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

func (o *QueryMLExclusionsV1TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /policy/queries/ml-exclusions/v1][%d] queryMLExclusionsV1TooManyRequests  %+v", 429, o.Payload)
}
func (o *QueryMLExclusionsV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryMLExclusionsV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryMLExclusionsV1InternalServerError creates a QueryMLExclusionsV1InternalServerError with default headers values
func NewQueryMLExclusionsV1InternalServerError() *QueryMLExclusionsV1InternalServerError {
	return &QueryMLExclusionsV1InternalServerError{}
}

/* QueryMLExclusionsV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type QueryMLExclusionsV1InternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

func (o *QueryMLExclusionsV1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /policy/queries/ml-exclusions/v1][%d] queryMLExclusionsV1InternalServerError  %+v", 500, o.Payload)
}
func (o *QueryMLExclusionsV1InternalServerError) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryMLExclusionsV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryMLExclusionsV1Default creates a QueryMLExclusionsV1Default with default headers values
func NewQueryMLExclusionsV1Default(code int) *QueryMLExclusionsV1Default {
	return &QueryMLExclusionsV1Default{
		_statusCode: code,
	}
}

/* QueryMLExclusionsV1Default describes a response with status code -1, with default header values.

OK
*/
type QueryMLExclusionsV1Default struct {
	_statusCode int

	Payload *models.MsaQueryResponse
}

// Code gets the status code for the query m l exclusions v1 default response
func (o *QueryMLExclusionsV1Default) Code() int {
	return o._statusCode
}

func (o *QueryMLExclusionsV1Default) Error() string {
	return fmt.Sprintf("[GET /policy/queries/ml-exclusions/v1][%d] queryMLExclusionsV1 default  %+v", o._statusCode, o.Payload)
}
func (o *QueryMLExclusionsV1Default) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryMLExclusionsV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
