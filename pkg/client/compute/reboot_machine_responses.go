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

// RebootMachineReader is a Reader for the RebootMachine structure.
type RebootMachineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RebootMachineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewRebootMachineAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewRebootMachineForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRebootMachineNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRebootMachineAccepted creates a RebootMachineAccepted with default headers values
func NewRebootMachineAccepted() *RebootMachineAccepted {
	return &RebootMachineAccepted{}
}

/*RebootMachineAccepted handles this case with default header values.

successful operation
*/
type RebootMachineAccepted struct {
	Payload *models.RequestTracker
}

func (o *RebootMachineAccepted) Error() string {
	return fmt.Sprintf("[POST /iaas/api/machines/{id}/operations/reboot][%d] rebootMachineAccepted  %+v", 202, o.Payload)
}

func (o *RebootMachineAccepted) GetPayload() *models.RequestTracker {
	return o.Payload
}

func (o *RebootMachineAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTracker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRebootMachineForbidden creates a RebootMachineForbidden with default headers values
func NewRebootMachineForbidden() *RebootMachineForbidden {
	return &RebootMachineForbidden{}
}

/*RebootMachineForbidden handles this case with default header values.

Forbidden
*/
type RebootMachineForbidden struct {
}

func (o *RebootMachineForbidden) Error() string {
	return fmt.Sprintf("[POST /iaas/api/machines/{id}/operations/reboot][%d] rebootMachineForbidden ", 403)
}

func (o *RebootMachineForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRebootMachineNotFound creates a RebootMachineNotFound with default headers values
func NewRebootMachineNotFound() *RebootMachineNotFound {
	return &RebootMachineNotFound{}
}

/*RebootMachineNotFound handles this case with default header values.

Not Found
*/
type RebootMachineNotFound struct {
	Payload *models.Error
}

func (o *RebootMachineNotFound) Error() string {
	return fmt.Sprintf("[POST /iaas/api/machines/{id}/operations/reboot][%d] rebootMachineNotFound  %+v", 404, o.Payload)
}

func (o *RebootMachineNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *RebootMachineNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
