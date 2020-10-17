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

// GetAwsStorageProfileReader is a Reader for the GetAwsStorageProfile structure.
type GetAwsStorageProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAwsStorageProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAwsStorageProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetAwsStorageProfileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAwsStorageProfileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAwsStorageProfileOK creates a GetAwsStorageProfileOK with default headers values
func NewGetAwsStorageProfileOK() *GetAwsStorageProfileOK {
	return &GetAwsStorageProfileOK{}
}

/*GetAwsStorageProfileOK handles this case with default header values.

successful operation
*/
type GetAwsStorageProfileOK struct {
	Payload *models.AwsStorageProfile
}

func (o *GetAwsStorageProfileOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/storage-profiles-aws/{id}][%d] getAwsStorageProfileOK  %+v", 200, o.Payload)
}

func (o *GetAwsStorageProfileOK) GetPayload() *models.AwsStorageProfile {
	return o.Payload
}

func (o *GetAwsStorageProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AwsStorageProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAwsStorageProfileForbidden creates a GetAwsStorageProfileForbidden with default headers values
func NewGetAwsStorageProfileForbidden() *GetAwsStorageProfileForbidden {
	return &GetAwsStorageProfileForbidden{}
}

/*GetAwsStorageProfileForbidden handles this case with default header values.

Forbidden
*/
type GetAwsStorageProfileForbidden struct {
}

func (o *GetAwsStorageProfileForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/storage-profiles-aws/{id}][%d] getAwsStorageProfileForbidden ", 403)
}

func (o *GetAwsStorageProfileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAwsStorageProfileNotFound creates a GetAwsStorageProfileNotFound with default headers values
func NewGetAwsStorageProfileNotFound() *GetAwsStorageProfileNotFound {
	return &GetAwsStorageProfileNotFound{}
}

/*GetAwsStorageProfileNotFound handles this case with default header values.

Not Found
*/
type GetAwsStorageProfileNotFound struct {
	Payload *models.Error
}

func (o *GetAwsStorageProfileNotFound) Error() string {
	return fmt.Sprintf("[GET /iaas/api/storage-profiles-aws/{id}][%d] getAwsStorageProfileNotFound  %+v", 404, o.Payload)
}

func (o *GetAwsStorageProfileNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAwsStorageProfileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
