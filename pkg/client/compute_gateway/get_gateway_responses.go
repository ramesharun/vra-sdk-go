// Code generated by go-swagger; DO NOT EDIT.

package compute_gateway

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetGatewayReader is a Reader for the GetGateway structure.
type GetGatewayReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGatewayReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGatewayOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetGatewayForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGatewayNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGatewayOK creates a GetGatewayOK with default headers values
func NewGetGatewayOK() *GetGatewayOK {
	return &GetGatewayOK{}
}

/*GetGatewayOK handles this case with default header values.

successful operation
*/
type GetGatewayOK struct {
	Payload *models.ComputeGateway
}

func (o *GetGatewayOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/compute-gateways/{id}][%d] getGatewayOK  %+v", 200, o.Payload)
}

func (o *GetGatewayOK) GetPayload() *models.ComputeGateway {
	return o.Payload
}

func (o *GetGatewayOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComputeGateway)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGatewayForbidden creates a GetGatewayForbidden with default headers values
func NewGetGatewayForbidden() *GetGatewayForbidden {
	return &GetGatewayForbidden{}
}

/*GetGatewayForbidden handles this case with default header values.

Forbidden
*/
type GetGatewayForbidden struct {
}

func (o *GetGatewayForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/compute-gateways/{id}][%d] getGatewayForbidden ", 403)
}

func (o *GetGatewayForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetGatewayNotFound creates a GetGatewayNotFound with default headers values
func NewGetGatewayNotFound() *GetGatewayNotFound {
	return &GetGatewayNotFound{}
}

/*GetGatewayNotFound handles this case with default header values.

Not Found
*/
type GetGatewayNotFound struct {
	Payload *models.Error
}

func (o *GetGatewayNotFound) Error() string {
	return fmt.Sprintf("[GET /iaas/api/compute-gateways/{id}][%d] getGatewayNotFound  %+v", 404, o.Payload)
}

func (o *GetGatewayNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGatewayNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
