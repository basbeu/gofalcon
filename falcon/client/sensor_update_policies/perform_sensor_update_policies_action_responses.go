// Code generated by go-swagger; DO NOT EDIT.

package sensor_update_policies

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

// PerformSensorUpdatePoliciesActionReader is a Reader for the PerformSensorUpdatePoliciesAction structure.
type PerformSensorUpdatePoliciesActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PerformSensorUpdatePoliciesActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPerformSensorUpdatePoliciesActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPerformSensorUpdatePoliciesActionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPerformSensorUpdatePoliciesActionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPerformSensorUpdatePoliciesActionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPerformSensorUpdatePoliciesActionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPerformSensorUpdatePoliciesActionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPerformSensorUpdatePoliciesActionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPerformSensorUpdatePoliciesActionOK creates a PerformSensorUpdatePoliciesActionOK with default headers values
func NewPerformSensorUpdatePoliciesActionOK() *PerformSensorUpdatePoliciesActionOK {
	return &PerformSensorUpdatePoliciesActionOK{}
}

/* PerformSensorUpdatePoliciesActionOK describes a response with status code 200, with default header values.

OK
*/
type PerformSensorUpdatePoliciesActionOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesSensorUpdatePoliciesV1
}

func (o *PerformSensorUpdatePoliciesActionOK) Error() string {
	return fmt.Sprintf("[POST /policy/entities/sensor-update-actions/v1][%d] performSensorUpdatePoliciesActionOK  %+v", 200, o.Payload)
}
func (o *PerformSensorUpdatePoliciesActionOK) GetPayload() *models.ResponsesSensorUpdatePoliciesV1 {
	return o.Payload
}

func (o *PerformSensorUpdatePoliciesActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesSensorUpdatePoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformSensorUpdatePoliciesActionBadRequest creates a PerformSensorUpdatePoliciesActionBadRequest with default headers values
func NewPerformSensorUpdatePoliciesActionBadRequest() *PerformSensorUpdatePoliciesActionBadRequest {
	return &PerformSensorUpdatePoliciesActionBadRequest{}
}

/* PerformSensorUpdatePoliciesActionBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PerformSensorUpdatePoliciesActionBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesSensorUpdatePoliciesV1
}

func (o *PerformSensorUpdatePoliciesActionBadRequest) Error() string {
	return fmt.Sprintf("[POST /policy/entities/sensor-update-actions/v1][%d] performSensorUpdatePoliciesActionBadRequest  %+v", 400, o.Payload)
}
func (o *PerformSensorUpdatePoliciesActionBadRequest) GetPayload() *models.ResponsesSensorUpdatePoliciesV1 {
	return o.Payload
}

func (o *PerformSensorUpdatePoliciesActionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesSensorUpdatePoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformSensorUpdatePoliciesActionForbidden creates a PerformSensorUpdatePoliciesActionForbidden with default headers values
func NewPerformSensorUpdatePoliciesActionForbidden() *PerformSensorUpdatePoliciesActionForbidden {
	return &PerformSensorUpdatePoliciesActionForbidden{}
}

/* PerformSensorUpdatePoliciesActionForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PerformSensorUpdatePoliciesActionForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *PerformSensorUpdatePoliciesActionForbidden) Error() string {
	return fmt.Sprintf("[POST /policy/entities/sensor-update-actions/v1][%d] performSensorUpdatePoliciesActionForbidden  %+v", 403, o.Payload)
}
func (o *PerformSensorUpdatePoliciesActionForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *PerformSensorUpdatePoliciesActionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPerformSensorUpdatePoliciesActionNotFound creates a PerformSensorUpdatePoliciesActionNotFound with default headers values
func NewPerformSensorUpdatePoliciesActionNotFound() *PerformSensorUpdatePoliciesActionNotFound {
	return &PerformSensorUpdatePoliciesActionNotFound{}
}

/* PerformSensorUpdatePoliciesActionNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PerformSensorUpdatePoliciesActionNotFound struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesSensorUpdatePoliciesV1
}

func (o *PerformSensorUpdatePoliciesActionNotFound) Error() string {
	return fmt.Sprintf("[POST /policy/entities/sensor-update-actions/v1][%d] performSensorUpdatePoliciesActionNotFound  %+v", 404, o.Payload)
}
func (o *PerformSensorUpdatePoliciesActionNotFound) GetPayload() *models.ResponsesSensorUpdatePoliciesV1 {
	return o.Payload
}

func (o *PerformSensorUpdatePoliciesActionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesSensorUpdatePoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformSensorUpdatePoliciesActionTooManyRequests creates a PerformSensorUpdatePoliciesActionTooManyRequests with default headers values
func NewPerformSensorUpdatePoliciesActionTooManyRequests() *PerformSensorUpdatePoliciesActionTooManyRequests {
	return &PerformSensorUpdatePoliciesActionTooManyRequests{}
}

/* PerformSensorUpdatePoliciesActionTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type PerformSensorUpdatePoliciesActionTooManyRequests struct {

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

func (o *PerformSensorUpdatePoliciesActionTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /policy/entities/sensor-update-actions/v1][%d] performSensorUpdatePoliciesActionTooManyRequests  %+v", 429, o.Payload)
}
func (o *PerformSensorUpdatePoliciesActionTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *PerformSensorUpdatePoliciesActionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPerformSensorUpdatePoliciesActionInternalServerError creates a PerformSensorUpdatePoliciesActionInternalServerError with default headers values
func NewPerformSensorUpdatePoliciesActionInternalServerError() *PerformSensorUpdatePoliciesActionInternalServerError {
	return &PerformSensorUpdatePoliciesActionInternalServerError{}
}

/* PerformSensorUpdatePoliciesActionInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PerformSensorUpdatePoliciesActionInternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesSensorUpdatePoliciesV1
}

func (o *PerformSensorUpdatePoliciesActionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /policy/entities/sensor-update-actions/v1][%d] performSensorUpdatePoliciesActionInternalServerError  %+v", 500, o.Payload)
}
func (o *PerformSensorUpdatePoliciesActionInternalServerError) GetPayload() *models.ResponsesSensorUpdatePoliciesV1 {
	return o.Payload
}

func (o *PerformSensorUpdatePoliciesActionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesSensorUpdatePoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformSensorUpdatePoliciesActionDefault creates a PerformSensorUpdatePoliciesActionDefault with default headers values
func NewPerformSensorUpdatePoliciesActionDefault(code int) *PerformSensorUpdatePoliciesActionDefault {
	return &PerformSensorUpdatePoliciesActionDefault{
		_statusCode: code,
	}
}

/* PerformSensorUpdatePoliciesActionDefault describes a response with status code -1, with default header values.

OK
*/
type PerformSensorUpdatePoliciesActionDefault struct {
	_statusCode int

	Payload *models.ResponsesSensorUpdatePoliciesV1
}

// Code gets the status code for the perform sensor update policies action default response
func (o *PerformSensorUpdatePoliciesActionDefault) Code() int {
	return o._statusCode
}

func (o *PerformSensorUpdatePoliciesActionDefault) Error() string {
	return fmt.Sprintf("[POST /policy/entities/sensor-update-actions/v1][%d] performSensorUpdatePoliciesAction default  %+v", o._statusCode, o.Payload)
}
func (o *PerformSensorUpdatePoliciesActionDefault) GetPayload() *models.ResponsesSensorUpdatePoliciesV1 {
	return o.Payload
}

func (o *PerformSensorUpdatePoliciesActionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponsesSensorUpdatePoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
