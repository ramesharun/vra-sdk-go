// Code generated by go-swagger; DO NOT EDIT.

package disk

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// CreateFirstClassDiskSnapshotReader is a Reader for the CreateFirstClassDiskSnapshot structure.
type CreateFirstClassDiskSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateFirstClassDiskSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewCreateFirstClassDiskSnapshotAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewCreateFirstClassDiskSnapshotNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateFirstClassDiskSnapshotForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateFirstClassDiskSnapshotAccepted creates a CreateFirstClassDiskSnapshotAccepted with default headers values
func NewCreateFirstClassDiskSnapshotAccepted() *CreateFirstClassDiskSnapshotAccepted {
	return &CreateFirstClassDiskSnapshotAccepted{}
}

/*CreateFirstClassDiskSnapshotAccepted handles this case with default header values.

successful operation
*/
type CreateFirstClassDiskSnapshotAccepted struct {
	Payload *models.RequestTracker
}

func (o *CreateFirstClassDiskSnapshotAccepted) Error() string {
	return fmt.Sprintf("[POST /iaas/api/block-devices/{id}/operations/snapshots][%d] createFirstClassDiskSnapshotAccepted  %+v", 202, o.Payload)
}

func (o *CreateFirstClassDiskSnapshotAccepted) GetPayload() *models.RequestTracker {
	return o.Payload
}

func (o *CreateFirstClassDiskSnapshotAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTracker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFirstClassDiskSnapshotNoContent creates a CreateFirstClassDiskSnapshotNoContent with default headers values
func NewCreateFirstClassDiskSnapshotNoContent() *CreateFirstClassDiskSnapshotNoContent {
	return &CreateFirstClassDiskSnapshotNoContent{}
}

/*CreateFirstClassDiskSnapshotNoContent handles this case with default header values.

No Content
*/
type CreateFirstClassDiskSnapshotNoContent struct {
}

func (o *CreateFirstClassDiskSnapshotNoContent) Error() string {
	return fmt.Sprintf("[POST /iaas/api/block-devices/{id}/operations/snapshots][%d] createFirstClassDiskSnapshotNoContent ", 204)
}

func (o *CreateFirstClassDiskSnapshotNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateFirstClassDiskSnapshotForbidden creates a CreateFirstClassDiskSnapshotForbidden with default headers values
func NewCreateFirstClassDiskSnapshotForbidden() *CreateFirstClassDiskSnapshotForbidden {
	return &CreateFirstClassDiskSnapshotForbidden{}
}

/*CreateFirstClassDiskSnapshotForbidden handles this case with default header values.

Forbidden
*/
type CreateFirstClassDiskSnapshotForbidden struct {
}

func (o *CreateFirstClassDiskSnapshotForbidden) Error() string {
	return fmt.Sprintf("[POST /iaas/api/block-devices/{id}/operations/snapshots][%d] createFirstClassDiskSnapshotForbidden ", 403)
}

func (o *CreateFirstClassDiskSnapshotForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
