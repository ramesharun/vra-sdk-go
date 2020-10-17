// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// CreateMachineReader is a Reader for the CreateMachine structure.
type CreateMachineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMachineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewCreateMachineAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateMachineBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateMachineForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateMachineAccepted creates a CreateMachineAccepted with default headers values
func NewCreateMachineAccepted() *CreateMachineAccepted {
	return &CreateMachineAccepted{}
}

/*CreateMachineAccepted handles this case with default header values.

successful operation
*/
type CreateMachineAccepted struct {
	Payload *models.RequestTracker
}

func (o *CreateMachineAccepted) Error() string {
	return fmt.Sprintf("[POST /iaas/api/machines][%d] createMachineAccepted  %+v", 202, o.Payload)
}

func (o *CreateMachineAccepted) GetPayload() *models.RequestTracker {
	return o.Payload
}

func (o *CreateMachineAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTracker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMachineBadRequest creates a CreateMachineBadRequest with default headers values
func NewCreateMachineBadRequest() *CreateMachineBadRequest {
	return &CreateMachineBadRequest{}
}

/*CreateMachineBadRequest handles this case with default header values.

Invalid Request - bad data
*/
type CreateMachineBadRequest struct {
	Payload *models.Error
}

func (o *CreateMachineBadRequest) Error() string {
	return fmt.Sprintf("[POST /iaas/api/machines][%d] createMachineBadRequest  %+v", 400, o.Payload)
}

func (o *CreateMachineBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateMachineBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMachineForbidden creates a CreateMachineForbidden with default headers values
func NewCreateMachineForbidden() *CreateMachineForbidden {
	return &CreateMachineForbidden{}
}

/*CreateMachineForbidden handles this case with default header values.

Forbidden
*/
type CreateMachineForbidden struct {
}

func (o *CreateMachineForbidden) Error() string {
	return fmt.Sprintf("[POST /iaas/api/machines][%d] createMachineForbidden ", 403)
}

func (o *CreateMachineForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
