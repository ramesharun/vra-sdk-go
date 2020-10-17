// Code generated by go-swagger; DO NOT EDIT.

package storage_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// CreateAwsStorageProfileReader is a Reader for the CreateAwsStorageProfile structure.
type CreateAwsStorageProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAwsStorageProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateAwsStorageProfileCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateAwsStorageProfileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateAwsStorageProfileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateAwsStorageProfileCreated creates a CreateAwsStorageProfileCreated with default headers values
func NewCreateAwsStorageProfileCreated() *CreateAwsStorageProfileCreated {
	return &CreateAwsStorageProfileCreated{}
}

/*CreateAwsStorageProfileCreated handles this case with default header values.

successful operation
*/
type CreateAwsStorageProfileCreated struct {
	Payload *models.AwsStorageProfile
}

func (o *CreateAwsStorageProfileCreated) Error() string {
	return fmt.Sprintf("[POST /iaas/api/storage-profiles-aws][%d] createAwsStorageProfileCreated  %+v", 201, o.Payload)
}

func (o *CreateAwsStorageProfileCreated) GetPayload() *models.AwsStorageProfile {
	return o.Payload
}

func (o *CreateAwsStorageProfileCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AwsStorageProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAwsStorageProfileBadRequest creates a CreateAwsStorageProfileBadRequest with default headers values
func NewCreateAwsStorageProfileBadRequest() *CreateAwsStorageProfileBadRequest {
	return &CreateAwsStorageProfileBadRequest{}
}

/*CreateAwsStorageProfileBadRequest handles this case with default header values.

Invalid Request - bad data
*/
type CreateAwsStorageProfileBadRequest struct {
	Payload *models.Error
}

func (o *CreateAwsStorageProfileBadRequest) Error() string {
	return fmt.Sprintf("[POST /iaas/api/storage-profiles-aws][%d] createAwsStorageProfileBadRequest  %+v", 400, o.Payload)
}

func (o *CreateAwsStorageProfileBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateAwsStorageProfileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAwsStorageProfileForbidden creates a CreateAwsStorageProfileForbidden with default headers values
func NewCreateAwsStorageProfileForbidden() *CreateAwsStorageProfileForbidden {
	return &CreateAwsStorageProfileForbidden{}
}

/*CreateAwsStorageProfileForbidden handles this case with default header values.

Forbidden
*/
type CreateAwsStorageProfileForbidden struct {
}

func (o *CreateAwsStorageProfileForbidden) Error() string {
	return fmt.Sprintf("[POST /iaas/api/storage-profiles-aws][%d] createAwsStorageProfileForbidden ", 403)
}

func (o *CreateAwsStorageProfileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
