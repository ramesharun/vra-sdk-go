// Code generated by go-swagger; DO NOT EDIT.

package cloud_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// GetVmcCloudAccountReader is a Reader for the GetVmcCloudAccount structure.
type GetVmcCloudAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVmcCloudAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVmcCloudAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetVmcCloudAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetVmcCloudAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetVmcCloudAccountOK creates a GetVmcCloudAccountOK with default headers values
func NewGetVmcCloudAccountOK() *GetVmcCloudAccountOK {
	return &GetVmcCloudAccountOK{}
}

/*GetVmcCloudAccountOK handles this case with default header values.

successful operation
*/
type GetVmcCloudAccountOK struct {
	Payload *models.CloudAccountVmc
}

func (o *GetVmcCloudAccountOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/cloud-accounts-vmc/{id}][%d] getVmcCloudAccountOK  %+v", 200, o.Payload)
}

func (o *GetVmcCloudAccountOK) GetPayload() *models.CloudAccountVmc {
	return o.Payload
}

func (o *GetVmcCloudAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudAccountVmc)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVmcCloudAccountForbidden creates a GetVmcCloudAccountForbidden with default headers values
func NewGetVmcCloudAccountForbidden() *GetVmcCloudAccountForbidden {
	return &GetVmcCloudAccountForbidden{}
}

/*GetVmcCloudAccountForbidden handles this case with default header values.

Forbidden
*/
type GetVmcCloudAccountForbidden struct {
}

func (o *GetVmcCloudAccountForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/cloud-accounts-vmc/{id}][%d] getVmcCloudAccountForbidden ", 403)
}

func (o *GetVmcCloudAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetVmcCloudAccountNotFound creates a GetVmcCloudAccountNotFound with default headers values
func NewGetVmcCloudAccountNotFound() *GetVmcCloudAccountNotFound {
	return &GetVmcCloudAccountNotFound{}
}

/*GetVmcCloudAccountNotFound handles this case with default header values.

Not Found
*/
type GetVmcCloudAccountNotFound struct {
}

func (o *GetVmcCloudAccountNotFound) Error() string {
	return fmt.Sprintf("[GET /iaas/api/cloud-accounts-vmc/{id}][%d] getVmcCloudAccountNotFound ", 404)
}

func (o *GetVmcCloudAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
