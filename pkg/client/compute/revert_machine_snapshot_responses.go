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

// RevertMachineSnapshotReader is a Reader for the RevertMachineSnapshot structure.
type RevertMachineSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RevertMachineSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewRevertMachineSnapshotAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewRevertMachineSnapshotForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRevertMachineSnapshotNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRevertMachineSnapshotAccepted creates a RevertMachineSnapshotAccepted with default headers values
func NewRevertMachineSnapshotAccepted() *RevertMachineSnapshotAccepted {
	return &RevertMachineSnapshotAccepted{}
}

/*RevertMachineSnapshotAccepted handles this case with default header values.

successful operation
*/
type RevertMachineSnapshotAccepted struct {
	Payload *models.RequestTracker
}

func (o *RevertMachineSnapshotAccepted) Error() string {
	return fmt.Sprintf("[POST /iaas/api/machines/{id}/operations/revert][%d] revertMachineSnapshotAccepted  %+v", 202, o.Payload)
}

func (o *RevertMachineSnapshotAccepted) GetPayload() *models.RequestTracker {
	return o.Payload
}

func (o *RevertMachineSnapshotAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTracker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRevertMachineSnapshotForbidden creates a RevertMachineSnapshotForbidden with default headers values
func NewRevertMachineSnapshotForbidden() *RevertMachineSnapshotForbidden {
	return &RevertMachineSnapshotForbidden{}
}

/*RevertMachineSnapshotForbidden handles this case with default header values.

Forbidden
*/
type RevertMachineSnapshotForbidden struct {
}

func (o *RevertMachineSnapshotForbidden) Error() string {
	return fmt.Sprintf("[POST /iaas/api/machines/{id}/operations/revert][%d] revertMachineSnapshotForbidden ", 403)
}

func (o *RevertMachineSnapshotForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRevertMachineSnapshotNotFound creates a RevertMachineSnapshotNotFound with default headers values
func NewRevertMachineSnapshotNotFound() *RevertMachineSnapshotNotFound {
	return &RevertMachineSnapshotNotFound{}
}

/*RevertMachineSnapshotNotFound handles this case with default header values.

Not Found
*/
type RevertMachineSnapshotNotFound struct {
	Payload *models.Error
}

func (o *RevertMachineSnapshotNotFound) Error() string {
	return fmt.Sprintf("[POST /iaas/api/machines/{id}/operations/revert][%d] revertMachineSnapshotNotFound  %+v", 404, o.Payload)
}

func (o *RevertMachineSnapshotNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *RevertMachineSnapshotNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
