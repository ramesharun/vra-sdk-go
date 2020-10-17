// Code generated by go-swagger; DO NOT EDIT.

package deployment_actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// SubmitResourceActionRequestUsingPOSTReader is a Reader for the SubmitResourceActionRequestUsingPOST structure.
type SubmitResourceActionRequestUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubmitResourceActionRequestUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSubmitResourceActionRequestUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSubmitResourceActionRequestUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSubmitResourceActionRequestUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSubmitResourceActionRequestUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSubmitResourceActionRequestUsingPOSTOK creates a SubmitResourceActionRequestUsingPOSTOK with default headers values
func NewSubmitResourceActionRequestUsingPOSTOK() *SubmitResourceActionRequestUsingPOSTOK {
	return &SubmitResourceActionRequestUsingPOSTOK{}
}

/*SubmitResourceActionRequestUsingPOSTOK handles this case with default header values.

OK
*/
type SubmitResourceActionRequestUsingPOSTOK struct {
	Payload *models.DeploymentRequest
}

func (o *SubmitResourceActionRequestUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /deployment/api/deployments/{depId}/resources/{resourceId}/requests][%d] submitResourceActionRequestUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *SubmitResourceActionRequestUsingPOSTOK) GetPayload() *models.DeploymentRequest {
	return o.Payload
}

func (o *SubmitResourceActionRequestUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitResourceActionRequestUsingPOSTUnauthorized creates a SubmitResourceActionRequestUsingPOSTUnauthorized with default headers values
func NewSubmitResourceActionRequestUsingPOSTUnauthorized() *SubmitResourceActionRequestUsingPOSTUnauthorized {
	return &SubmitResourceActionRequestUsingPOSTUnauthorized{}
}

/*SubmitResourceActionRequestUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type SubmitResourceActionRequestUsingPOSTUnauthorized struct {
}

func (o *SubmitResourceActionRequestUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /deployment/api/deployments/{depId}/resources/{resourceId}/requests][%d] submitResourceActionRequestUsingPOSTUnauthorized ", 401)
}

func (o *SubmitResourceActionRequestUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSubmitResourceActionRequestUsingPOSTForbidden creates a SubmitResourceActionRequestUsingPOSTForbidden with default headers values
func NewSubmitResourceActionRequestUsingPOSTForbidden() *SubmitResourceActionRequestUsingPOSTForbidden {
	return &SubmitResourceActionRequestUsingPOSTForbidden{}
}

/*SubmitResourceActionRequestUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type SubmitResourceActionRequestUsingPOSTForbidden struct {
}

func (o *SubmitResourceActionRequestUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /deployment/api/deployments/{depId}/resources/{resourceId}/requests][%d] submitResourceActionRequestUsingPOSTForbidden ", 403)
}

func (o *SubmitResourceActionRequestUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSubmitResourceActionRequestUsingPOSTNotFound creates a SubmitResourceActionRequestUsingPOSTNotFound with default headers values
func NewSubmitResourceActionRequestUsingPOSTNotFound() *SubmitResourceActionRequestUsingPOSTNotFound {
	return &SubmitResourceActionRequestUsingPOSTNotFound{}
}

/*SubmitResourceActionRequestUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type SubmitResourceActionRequestUsingPOSTNotFound struct {
	Payload *models.Error
}

func (o *SubmitResourceActionRequestUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /deployment/api/deployments/{depId}/resources/{resourceId}/requests][%d] submitResourceActionRequestUsingPOSTNotFound  %+v", 404, o.Payload)
}

func (o *SubmitResourceActionRequestUsingPOSTNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *SubmitResourceActionRequestUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
