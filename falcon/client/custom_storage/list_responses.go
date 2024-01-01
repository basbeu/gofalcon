// Code generated by go-swagger; DO NOT EDIT.

package custom_storage

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

// ListReader is a Reader for the List structure.
type ListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewListTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /customobjects/v1/collections/{collection_name}/objects] list", response, response.Code())
	}
}

// NewListOK creates a ListOK with default headers values
func NewListOK() *ListOK {
	return &ListOK{}
}

/*
ListOK describes a response with status code 200, with default header values.

OK
*/
type ListOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.CustomStorageObjectKeys
}

// IsSuccess returns true when this list o k response has a 2xx status code
func (o *ListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list o k response has a 3xx status code
func (o *ListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list o k response has a 4xx status code
func (o *ListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list o k response has a 5xx status code
func (o *ListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list o k response a status code equal to that given
func (o *ListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list o k response
func (o *ListOK) Code() int {
	return 200
}

func (o *ListOK) Error() string {
	return fmt.Sprintf("[GET /customobjects/v1/collections/{collection_name}/objects][%d] listOK  %+v", 200, o.Payload)
}

func (o *ListOK) String() string {
	return fmt.Sprintf("[GET /customobjects/v1/collections/{collection_name}/objects][%d] listOK  %+v", 200, o.Payload)
}

func (o *ListOK) GetPayload() *models.CustomStorageObjectKeys {
	return o.Payload
}

func (o *ListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.CustomStorageObjectKeys)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListForbidden creates a ListForbidden with default headers values
func NewListForbidden() *ListForbidden {
	return &ListForbidden{}
}

/*
ListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ListForbidden struct {

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

// IsSuccess returns true when this list forbidden response has a 2xx status code
func (o *ListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list forbidden response has a 3xx status code
func (o *ListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list forbidden response has a 4xx status code
func (o *ListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list forbidden response has a 5xx status code
func (o *ListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list forbidden response a status code equal to that given
func (o *ListForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list forbidden response
func (o *ListForbidden) Code() int {
	return 403
}

func (o *ListForbidden) Error() string {
	return fmt.Sprintf("[GET /customobjects/v1/collections/{collection_name}/objects][%d] listForbidden  %+v", 403, o.Payload)
}

func (o *ListForbidden) String() string {
	return fmt.Sprintf("[GET /customobjects/v1/collections/{collection_name}/objects][%d] listForbidden  %+v", 403, o.Payload)
}

func (o *ListForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListTooManyRequests creates a ListTooManyRequests with default headers values
func NewListTooManyRequests() *ListTooManyRequests {
	return &ListTooManyRequests{}
}

/*
ListTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type ListTooManyRequests struct {

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

// IsSuccess returns true when this list too many requests response has a 2xx status code
func (o *ListTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list too many requests response has a 3xx status code
func (o *ListTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list too many requests response has a 4xx status code
func (o *ListTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this list too many requests response has a 5xx status code
func (o *ListTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this list too many requests response a status code equal to that given
func (o *ListTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the list too many requests response
func (o *ListTooManyRequests) Code() int {
	return 429
}

func (o *ListTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /customobjects/v1/collections/{collection_name}/objects][%d] listTooManyRequests  %+v", 429, o.Payload)
}

func (o *ListTooManyRequests) String() string {
	return fmt.Sprintf("[GET /customobjects/v1/collections/{collection_name}/objects][%d] listTooManyRequests  %+v", 429, o.Payload)
}

func (o *ListTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ListTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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