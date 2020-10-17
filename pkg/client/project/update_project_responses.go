// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// UpdateProjectReader is a Reader for the UpdateProject structure.
type UpdateProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateProjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateProjectOK creates a UpdateProjectOK with default headers values
func NewUpdateProjectOK() *UpdateProjectOK {
	return &UpdateProjectOK{}
}

/*UpdateProjectOK handles this case with default header values.

successful operation
*/
type UpdateProjectOK struct {
	Payload *models.Project
}

func (o *UpdateProjectOK) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/projects/{id}][%d] updateProjectOK  %+v", 200, o.Payload)
}

func (o *UpdateProjectOK) GetPayload() *models.Project {
	return o.Payload
}

func (o *UpdateProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Project)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectBadRequest creates a UpdateProjectBadRequest with default headers values
func NewUpdateProjectBadRequest() *UpdateProjectBadRequest {
	return &UpdateProjectBadRequest{}
}

/*UpdateProjectBadRequest handles this case with default header values.

Invalid Request - bad data
*/
type UpdateProjectBadRequest struct {
	Payload *models.Error
}

func (o *UpdateProjectBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/projects/{id}][%d] updateProjectBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateProjectBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateProjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectForbidden creates a UpdateProjectForbidden with default headers values
func NewUpdateProjectForbidden() *UpdateProjectForbidden {
	return &UpdateProjectForbidden{}
}

/*UpdateProjectForbidden handles this case with default header values.

Forbidden
*/
type UpdateProjectForbidden struct {
}

func (o *UpdateProjectForbidden) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/projects/{id}][%d] updateProjectForbidden ", 403)
}

func (o *UpdateProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
