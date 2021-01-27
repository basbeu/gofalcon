// Code generated by go-swagger; DO NOT EDIT.

package cloud_connect_aws

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

// VerifyAWSAccountAccessReader is a Reader for the VerifyAWSAccountAccess structure.
type VerifyAWSAccountAccessReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VerifyAWSAccountAccessReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVerifyAWSAccountAccessOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewVerifyAWSAccountAccessBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewVerifyAWSAccountAccessForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewVerifyAWSAccountAccessTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewVerifyAWSAccountAccessInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewVerifyAWSAccountAccessDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVerifyAWSAccountAccessOK creates a VerifyAWSAccountAccessOK with default headers values
func NewVerifyAWSAccountAccessOK() *VerifyAWSAccountAccessOK {
	return &VerifyAWSAccountAccessOK{}
}

/* VerifyAWSAccountAccessOK describes a response with status code 200, with default header values.

OK
*/
type VerifyAWSAccountAccessOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ModelsVerifyAccessResponseV1
}

func (o *VerifyAWSAccountAccessOK) Error() string {
	return fmt.Sprintf("[POST /cloud-connect-aws/entities/verify-account-access/v1][%d] verifyAWSAccountAccessOK  %+v", 200, o.Payload)
}
func (o *VerifyAWSAccountAccessOK) GetPayload() *models.ModelsVerifyAccessResponseV1 {
	return o.Payload
}

func (o *VerifyAWSAccountAccessOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ModelsVerifyAccessResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVerifyAWSAccountAccessBadRequest creates a VerifyAWSAccountAccessBadRequest with default headers values
func NewVerifyAWSAccountAccessBadRequest() *VerifyAWSAccountAccessBadRequest {
	return &VerifyAWSAccountAccessBadRequest{}
}

/* VerifyAWSAccountAccessBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type VerifyAWSAccountAccessBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ModelsVerifyAccessResponseV1
}

func (o *VerifyAWSAccountAccessBadRequest) Error() string {
	return fmt.Sprintf("[POST /cloud-connect-aws/entities/verify-account-access/v1][%d] verifyAWSAccountAccessBadRequest  %+v", 400, o.Payload)
}
func (o *VerifyAWSAccountAccessBadRequest) GetPayload() *models.ModelsVerifyAccessResponseV1 {
	return o.Payload
}

func (o *VerifyAWSAccountAccessBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ModelsVerifyAccessResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVerifyAWSAccountAccessForbidden creates a VerifyAWSAccountAccessForbidden with default headers values
func NewVerifyAWSAccountAccessForbidden() *VerifyAWSAccountAccessForbidden {
	return &VerifyAWSAccountAccessForbidden{}
}

/* VerifyAWSAccountAccessForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type VerifyAWSAccountAccessForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *VerifyAWSAccountAccessForbidden) Error() string {
	return fmt.Sprintf("[POST /cloud-connect-aws/entities/verify-account-access/v1][%d] verifyAWSAccountAccessForbidden  %+v", 403, o.Payload)
}
func (o *VerifyAWSAccountAccessForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *VerifyAWSAccountAccessForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewVerifyAWSAccountAccessTooManyRequests creates a VerifyAWSAccountAccessTooManyRequests with default headers values
func NewVerifyAWSAccountAccessTooManyRequests() *VerifyAWSAccountAccessTooManyRequests {
	return &VerifyAWSAccountAccessTooManyRequests{}
}

/* VerifyAWSAccountAccessTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type VerifyAWSAccountAccessTooManyRequests struct {

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

func (o *VerifyAWSAccountAccessTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /cloud-connect-aws/entities/verify-account-access/v1][%d] verifyAWSAccountAccessTooManyRequests  %+v", 429, o.Payload)
}
func (o *VerifyAWSAccountAccessTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *VerifyAWSAccountAccessTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewVerifyAWSAccountAccessInternalServerError creates a VerifyAWSAccountAccessInternalServerError with default headers values
func NewVerifyAWSAccountAccessInternalServerError() *VerifyAWSAccountAccessInternalServerError {
	return &VerifyAWSAccountAccessInternalServerError{}
}

/* VerifyAWSAccountAccessInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type VerifyAWSAccountAccessInternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ModelsVerifyAccessResponseV1
}

func (o *VerifyAWSAccountAccessInternalServerError) Error() string {
	return fmt.Sprintf("[POST /cloud-connect-aws/entities/verify-account-access/v1][%d] verifyAWSAccountAccessInternalServerError  %+v", 500, o.Payload)
}
func (o *VerifyAWSAccountAccessInternalServerError) GetPayload() *models.ModelsVerifyAccessResponseV1 {
	return o.Payload
}

func (o *VerifyAWSAccountAccessInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ModelsVerifyAccessResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVerifyAWSAccountAccessDefault creates a VerifyAWSAccountAccessDefault with default headers values
func NewVerifyAWSAccountAccessDefault(code int) *VerifyAWSAccountAccessDefault {
	return &VerifyAWSAccountAccessDefault{
		_statusCode: code,
	}
}

/* VerifyAWSAccountAccessDefault describes a response with status code -1, with default header values.

OK
*/
type VerifyAWSAccountAccessDefault struct {
	_statusCode int

	Payload *models.ModelsVerifyAccessResponseV1
}

// Code gets the status code for the verify a w s account access default response
func (o *VerifyAWSAccountAccessDefault) Code() int {
	return o._statusCode
}

func (o *VerifyAWSAccountAccessDefault) Error() string {
	return fmt.Sprintf("[POST /cloud-connect-aws/entities/verify-account-access/v1][%d] VerifyAWSAccountAccess default  %+v", o._statusCode, o.Payload)
}
func (o *VerifyAWSAccountAccessDefault) GetPayload() *models.ModelsVerifyAccessResponseV1 {
	return o.Payload
}

func (o *VerifyAWSAccountAccessDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelsVerifyAccessResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
