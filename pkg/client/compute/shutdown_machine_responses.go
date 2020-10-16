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

// ShutdownMachineReader is a Reader for the ShutdownMachine structure.
type ShutdownMachineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShutdownMachineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewShutdownMachineAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewShutdownMachineForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewShutdownMachineNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewShutdownMachineAccepted creates a ShutdownMachineAccepted with default headers values
func NewShutdownMachineAccepted() *ShutdownMachineAccepted {
	return &ShutdownMachineAccepted{}
}

/*ShutdownMachineAccepted handles this case with default header values.

successful operation
*/
type ShutdownMachineAccepted struct {
	Payload *models.RequestTracker
}

func (o *ShutdownMachineAccepted) Error() string {
	return fmt.Sprintf("[POST /iaas/api/machines/{id}/operations/shutdown][%d] shutdownMachineAccepted  %+v", 202, o.Payload)
}

func (o *ShutdownMachineAccepted) GetPayload() *models.RequestTracker {
	return o.Payload
}

func (o *ShutdownMachineAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTracker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShutdownMachineForbidden creates a ShutdownMachineForbidden with default headers values
func NewShutdownMachineForbidden() *ShutdownMachineForbidden {
	return &ShutdownMachineForbidden{}
}

/*ShutdownMachineForbidden handles this case with default header values.

Forbidden
*/
type ShutdownMachineForbidden struct {
}

func (o *ShutdownMachineForbidden) Error() string {
	return fmt.Sprintf("[POST /iaas/api/machines/{id}/operations/shutdown][%d] shutdownMachineForbidden ", 403)
}

func (o *ShutdownMachineForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewShutdownMachineNotFound creates a ShutdownMachineNotFound with default headers values
func NewShutdownMachineNotFound() *ShutdownMachineNotFound {
	return &ShutdownMachineNotFound{}
}

/*ShutdownMachineNotFound handles this case with default header values.

Not Found
*/
type ShutdownMachineNotFound struct {
	Payload *models.Error
}

func (o *ShutdownMachineNotFound) Error() string {
	return fmt.Sprintf("[POST /iaas/api/machines/{id}/operations/shutdown][%d] shutdownMachineNotFound  %+v", 404, o.Payload)
}

func (o *ShutdownMachineNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ShutdownMachineNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
