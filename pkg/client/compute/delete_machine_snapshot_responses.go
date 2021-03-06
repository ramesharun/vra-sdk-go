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

// DeleteMachineSnapshotReader is a Reader for the DeleteMachineSnapshot structure.
type DeleteMachineSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMachineSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteMachineSnapshotAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteMachineSnapshotForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteMachineSnapshotAccepted creates a DeleteMachineSnapshotAccepted with default headers values
func NewDeleteMachineSnapshotAccepted() *DeleteMachineSnapshotAccepted {
	return &DeleteMachineSnapshotAccepted{}
}

/*DeleteMachineSnapshotAccepted handles this case with default header values.

successful operation
*/
type DeleteMachineSnapshotAccepted struct {
	Payload *models.RequestTracker
}

func (o *DeleteMachineSnapshotAccepted) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/machines/{id}/snapshots/{id1}][%d] deleteMachineSnapshotAccepted  %+v", 202, o.Payload)
}

func (o *DeleteMachineSnapshotAccepted) GetPayload() *models.RequestTracker {
	return o.Payload
}

func (o *DeleteMachineSnapshotAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTracker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMachineSnapshotForbidden creates a DeleteMachineSnapshotForbidden with default headers values
func NewDeleteMachineSnapshotForbidden() *DeleteMachineSnapshotForbidden {
	return &DeleteMachineSnapshotForbidden{}
}

/*DeleteMachineSnapshotForbidden handles this case with default header values.

Forbidden
*/
type DeleteMachineSnapshotForbidden struct {
}

func (o *DeleteMachineSnapshotForbidden) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/machines/{id}/snapshots/{id1}][%d] deleteMachineSnapshotForbidden ", 403)
}

func (o *DeleteMachineSnapshotForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
