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

// GetAzureStorageProfileReader is a Reader for the GetAzureStorageProfile structure.
type GetAzureStorageProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAzureStorageProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAzureStorageProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetAzureStorageProfileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAzureStorageProfileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAzureStorageProfileOK creates a GetAzureStorageProfileOK with default headers values
func NewGetAzureStorageProfileOK() *GetAzureStorageProfileOK {
	return &GetAzureStorageProfileOK{}
}

/*GetAzureStorageProfileOK handles this case with default header values.

successful operation
*/
type GetAzureStorageProfileOK struct {
	Payload *models.AzureStorageProfile
}

func (o *GetAzureStorageProfileOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/storage-profiles-azure/{id}][%d] getAzureStorageProfileOK  %+v", 200, o.Payload)
}

func (o *GetAzureStorageProfileOK) GetPayload() *models.AzureStorageProfile {
	return o.Payload
}

func (o *GetAzureStorageProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AzureStorageProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAzureStorageProfileForbidden creates a GetAzureStorageProfileForbidden with default headers values
func NewGetAzureStorageProfileForbidden() *GetAzureStorageProfileForbidden {
	return &GetAzureStorageProfileForbidden{}
}

/*GetAzureStorageProfileForbidden handles this case with default header values.

Forbidden
*/
type GetAzureStorageProfileForbidden struct {
}

func (o *GetAzureStorageProfileForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/storage-profiles-azure/{id}][%d] getAzureStorageProfileForbidden ", 403)
}

func (o *GetAzureStorageProfileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAzureStorageProfileNotFound creates a GetAzureStorageProfileNotFound with default headers values
func NewGetAzureStorageProfileNotFound() *GetAzureStorageProfileNotFound {
	return &GetAzureStorageProfileNotFound{}
}

/*GetAzureStorageProfileNotFound handles this case with default header values.

Not Found
*/
type GetAzureStorageProfileNotFound struct {
	Payload *models.Error
}

func (o *GetAzureStorageProfileNotFound) Error() string {
	return fmt.Sprintf("[GET /iaas/api/storage-profiles-azure/{id}][%d] getAzureStorageProfileNotFound  %+v", 404, o.Payload)
}

func (o *GetAzureStorageProfileNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAzureStorageProfileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
