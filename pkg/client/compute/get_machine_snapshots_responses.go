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

// GetMachineSnapshotsReader is a Reader for the GetMachineSnapshots structure.
type GetMachineSnapshotsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMachineSnapshotsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMachineSnapshotsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetMachineSnapshotsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetMachineSnapshotsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetMachineSnapshotsOK creates a GetMachineSnapshotsOK with default headers values
func NewGetMachineSnapshotsOK() *GetMachineSnapshotsOK {
	return &GetMachineSnapshotsOK{}
}

/*GetMachineSnapshotsOK handles this case with default header values.

successful operation
*/
type GetMachineSnapshotsOK struct {
	Payload []*models.Snapshot
}

func (o *GetMachineSnapshotsOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/machines/{id}/snapshots][%d] getMachineSnapshotsOK  %+v", 200, o.Payload)
}

func (o *GetMachineSnapshotsOK) GetPayload() []*models.Snapshot {
	return o.Payload
}

func (o *GetMachineSnapshotsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMachineSnapshotsForbidden creates a GetMachineSnapshotsForbidden with default headers values
func NewGetMachineSnapshotsForbidden() *GetMachineSnapshotsForbidden {
	return &GetMachineSnapshotsForbidden{}
}

/*GetMachineSnapshotsForbidden handles this case with default header values.

Forbidden
*/
type GetMachineSnapshotsForbidden struct {
}

func (o *GetMachineSnapshotsForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/machines/{id}/snapshots][%d] getMachineSnapshotsForbidden ", 403)
}

func (o *GetMachineSnapshotsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetMachineSnapshotsNotFound creates a GetMachineSnapshotsNotFound with default headers values
func NewGetMachineSnapshotsNotFound() *GetMachineSnapshotsNotFound {
	return &GetMachineSnapshotsNotFound{}
}

/*GetMachineSnapshotsNotFound handles this case with default header values.

Not Found
*/
type GetMachineSnapshotsNotFound struct {
	Payload *models.Error
}

func (o *GetMachineSnapshotsNotFound) Error() string {
	return fmt.Sprintf("[GET /iaas/api/machines/{id}/snapshots][%d] getMachineSnapshotsNotFound  %+v", 404, o.Payload)
}

func (o *GetMachineSnapshotsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetMachineSnapshotsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
