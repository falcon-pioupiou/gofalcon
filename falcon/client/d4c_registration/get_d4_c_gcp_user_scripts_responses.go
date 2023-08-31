// Code generated by go-swagger; DO NOT EDIT.

package d4c_registration

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

// GetD4CGcpUserScriptsReader is a Reader for the GetD4CGcpUserScripts structure.
type GetD4CGcpUserScriptsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetD4CGcpUserScriptsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetD4CGcpUserScriptsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewGetD4CGcpUserScriptsMultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetD4CGcpUserScriptsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetD4CGcpUserScriptsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetD4CGcpUserScriptsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetD4CGcpUserScriptsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cloud-connect-gcp/entities/user-scripts/v1] GetD4CGcpUserScripts", response, response.Code())
	}
}

// NewGetD4CGcpUserScriptsOK creates a GetD4CGcpUserScriptsOK with default headers values
func NewGetD4CGcpUserScriptsOK() *GetD4CGcpUserScriptsOK {
	return &GetD4CGcpUserScriptsOK{}
}

/*
GetD4CGcpUserScriptsOK describes a response with status code 200, with default header values.

OK
*/
type GetD4CGcpUserScriptsOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationGCPProvisionGetUserScriptResponseV1
}

// IsSuccess returns true when this get d4 c gcp user scripts o k response has a 2xx status code
func (o *GetD4CGcpUserScriptsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get d4 c gcp user scripts o k response has a 3xx status code
func (o *GetD4CGcpUserScriptsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get d4 c gcp user scripts o k response has a 4xx status code
func (o *GetD4CGcpUserScriptsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get d4 c gcp user scripts o k response has a 5xx status code
func (o *GetD4CGcpUserScriptsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get d4 c gcp user scripts o k response a status code equal to that given
func (o *GetD4CGcpUserScriptsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get d4 c gcp user scripts o k response
func (o *GetD4CGcpUserScriptsOK) Code() int {
	return 200
}

func (o *GetD4CGcpUserScriptsOK) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/user-scripts/v1][%d] getD4CGcpUserScriptsOK  %+v", 200, o.Payload)
}

func (o *GetD4CGcpUserScriptsOK) String() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/user-scripts/v1][%d] getD4CGcpUserScriptsOK  %+v", 200, o.Payload)
}

func (o *GetD4CGcpUserScriptsOK) GetPayload() *models.RegistrationGCPProvisionGetUserScriptResponseV1 {
	return o.Payload
}

func (o *GetD4CGcpUserScriptsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationGCPProvisionGetUserScriptResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetD4CGcpUserScriptsMultiStatus creates a GetD4CGcpUserScriptsMultiStatus with default headers values
func NewGetD4CGcpUserScriptsMultiStatus() *GetD4CGcpUserScriptsMultiStatus {
	return &GetD4CGcpUserScriptsMultiStatus{}
}

/*
GetD4CGcpUserScriptsMultiStatus describes a response with status code 207, with default header values.

Multi-Status
*/
type GetD4CGcpUserScriptsMultiStatus struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationGCPProvisionGetUserScriptResponseV1
}

// IsSuccess returns true when this get d4 c gcp user scripts multi status response has a 2xx status code
func (o *GetD4CGcpUserScriptsMultiStatus) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get d4 c gcp user scripts multi status response has a 3xx status code
func (o *GetD4CGcpUserScriptsMultiStatus) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get d4 c gcp user scripts multi status response has a 4xx status code
func (o *GetD4CGcpUserScriptsMultiStatus) IsClientError() bool {
	return false
}

// IsServerError returns true when this get d4 c gcp user scripts multi status response has a 5xx status code
func (o *GetD4CGcpUserScriptsMultiStatus) IsServerError() bool {
	return false
}

// IsCode returns true when this get d4 c gcp user scripts multi status response a status code equal to that given
func (o *GetD4CGcpUserScriptsMultiStatus) IsCode(code int) bool {
	return code == 207
}

// Code gets the status code for the get d4 c gcp user scripts multi status response
func (o *GetD4CGcpUserScriptsMultiStatus) Code() int {
	return 207
}

func (o *GetD4CGcpUserScriptsMultiStatus) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/user-scripts/v1][%d] getD4CGcpUserScriptsMultiStatus  %+v", 207, o.Payload)
}

func (o *GetD4CGcpUserScriptsMultiStatus) String() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/user-scripts/v1][%d] getD4CGcpUserScriptsMultiStatus  %+v", 207, o.Payload)
}

func (o *GetD4CGcpUserScriptsMultiStatus) GetPayload() *models.RegistrationGCPProvisionGetUserScriptResponseV1 {
	return o.Payload
}

func (o *GetD4CGcpUserScriptsMultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationGCPProvisionGetUserScriptResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetD4CGcpUserScriptsBadRequest creates a GetD4CGcpUserScriptsBadRequest with default headers values
func NewGetD4CGcpUserScriptsBadRequest() *GetD4CGcpUserScriptsBadRequest {
	return &GetD4CGcpUserScriptsBadRequest{}
}

/*
GetD4CGcpUserScriptsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetD4CGcpUserScriptsBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationGCPProvisionGetUserScriptResponseV1
}

// IsSuccess returns true when this get d4 c gcp user scripts bad request response has a 2xx status code
func (o *GetD4CGcpUserScriptsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get d4 c gcp user scripts bad request response has a 3xx status code
func (o *GetD4CGcpUserScriptsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get d4 c gcp user scripts bad request response has a 4xx status code
func (o *GetD4CGcpUserScriptsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get d4 c gcp user scripts bad request response has a 5xx status code
func (o *GetD4CGcpUserScriptsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get d4 c gcp user scripts bad request response a status code equal to that given
func (o *GetD4CGcpUserScriptsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get d4 c gcp user scripts bad request response
func (o *GetD4CGcpUserScriptsBadRequest) Code() int {
	return 400
}

func (o *GetD4CGcpUserScriptsBadRequest) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/user-scripts/v1][%d] getD4CGcpUserScriptsBadRequest  %+v", 400, o.Payload)
}

func (o *GetD4CGcpUserScriptsBadRequest) String() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/user-scripts/v1][%d] getD4CGcpUserScriptsBadRequest  %+v", 400, o.Payload)
}

func (o *GetD4CGcpUserScriptsBadRequest) GetPayload() *models.RegistrationGCPProvisionGetUserScriptResponseV1 {
	return o.Payload
}

func (o *GetD4CGcpUserScriptsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationGCPProvisionGetUserScriptResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetD4CGcpUserScriptsForbidden creates a GetD4CGcpUserScriptsForbidden with default headers values
func NewGetD4CGcpUserScriptsForbidden() *GetD4CGcpUserScriptsForbidden {
	return &GetD4CGcpUserScriptsForbidden{}
}

/*
GetD4CGcpUserScriptsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetD4CGcpUserScriptsForbidden struct {

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

// IsSuccess returns true when this get d4 c gcp user scripts forbidden response has a 2xx status code
func (o *GetD4CGcpUserScriptsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get d4 c gcp user scripts forbidden response has a 3xx status code
func (o *GetD4CGcpUserScriptsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get d4 c gcp user scripts forbidden response has a 4xx status code
func (o *GetD4CGcpUserScriptsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get d4 c gcp user scripts forbidden response has a 5xx status code
func (o *GetD4CGcpUserScriptsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get d4 c gcp user scripts forbidden response a status code equal to that given
func (o *GetD4CGcpUserScriptsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get d4 c gcp user scripts forbidden response
func (o *GetD4CGcpUserScriptsForbidden) Code() int {
	return 403
}

func (o *GetD4CGcpUserScriptsForbidden) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/user-scripts/v1][%d] getD4CGcpUserScriptsForbidden  %+v", 403, o.Payload)
}

func (o *GetD4CGcpUserScriptsForbidden) String() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/user-scripts/v1][%d] getD4CGcpUserScriptsForbidden  %+v", 403, o.Payload)
}

func (o *GetD4CGcpUserScriptsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetD4CGcpUserScriptsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetD4CGcpUserScriptsTooManyRequests creates a GetD4CGcpUserScriptsTooManyRequests with default headers values
func NewGetD4CGcpUserScriptsTooManyRequests() *GetD4CGcpUserScriptsTooManyRequests {
	return &GetD4CGcpUserScriptsTooManyRequests{}
}

/*
GetD4CGcpUserScriptsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetD4CGcpUserScriptsTooManyRequests struct {

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

// IsSuccess returns true when this get d4 c gcp user scripts too many requests response has a 2xx status code
func (o *GetD4CGcpUserScriptsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get d4 c gcp user scripts too many requests response has a 3xx status code
func (o *GetD4CGcpUserScriptsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get d4 c gcp user scripts too many requests response has a 4xx status code
func (o *GetD4CGcpUserScriptsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get d4 c gcp user scripts too many requests response has a 5xx status code
func (o *GetD4CGcpUserScriptsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get d4 c gcp user scripts too many requests response a status code equal to that given
func (o *GetD4CGcpUserScriptsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get d4 c gcp user scripts too many requests response
func (o *GetD4CGcpUserScriptsTooManyRequests) Code() int {
	return 429
}

func (o *GetD4CGcpUserScriptsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/user-scripts/v1][%d] getD4CGcpUserScriptsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetD4CGcpUserScriptsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/user-scripts/v1][%d] getD4CGcpUserScriptsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetD4CGcpUserScriptsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetD4CGcpUserScriptsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetD4CGcpUserScriptsInternalServerError creates a GetD4CGcpUserScriptsInternalServerError with default headers values
func NewGetD4CGcpUserScriptsInternalServerError() *GetD4CGcpUserScriptsInternalServerError {
	return &GetD4CGcpUserScriptsInternalServerError{}
}

/*
GetD4CGcpUserScriptsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetD4CGcpUserScriptsInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationGCPProvisionGetUserScriptResponseV1
}

// IsSuccess returns true when this get d4 c gcp user scripts internal server error response has a 2xx status code
func (o *GetD4CGcpUserScriptsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get d4 c gcp user scripts internal server error response has a 3xx status code
func (o *GetD4CGcpUserScriptsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get d4 c gcp user scripts internal server error response has a 4xx status code
func (o *GetD4CGcpUserScriptsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get d4 c gcp user scripts internal server error response has a 5xx status code
func (o *GetD4CGcpUserScriptsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get d4 c gcp user scripts internal server error response a status code equal to that given
func (o *GetD4CGcpUserScriptsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get d4 c gcp user scripts internal server error response
func (o *GetD4CGcpUserScriptsInternalServerError) Code() int {
	return 500
}

func (o *GetD4CGcpUserScriptsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/user-scripts/v1][%d] getD4CGcpUserScriptsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetD4CGcpUserScriptsInternalServerError) String() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/user-scripts/v1][%d] getD4CGcpUserScriptsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetD4CGcpUserScriptsInternalServerError) GetPayload() *models.RegistrationGCPProvisionGetUserScriptResponseV1 {
	return o.Payload
}

func (o *GetD4CGcpUserScriptsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationGCPProvisionGetUserScriptResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}