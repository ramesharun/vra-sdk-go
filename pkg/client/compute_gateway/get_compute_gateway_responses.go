// Code generated by go-swagger; DO NOT EDIT.

package compute_gateway

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// GetComputeGatewayReader is a Reader for the GetComputeGateway structure.
type GetComputeGatewayReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetComputeGatewayReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetComputeGatewayOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetComputeGatewayForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetComputeGatewayOK creates a GetComputeGatewayOK with default headers values
func NewGetComputeGatewayOK() *GetComputeGatewayOK {
	return &GetComputeGatewayOK{}
}

/*GetComputeGatewayOK handles this case with default header values.

successful operation
*/
type GetComputeGatewayOK struct {
	Payload *models.ComputeGatewayResult
}

func (o *GetComputeGatewayOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/compute-gateways][%d] getComputeGatewayOK  %+v", 200, o.Payload)
}

func (o *GetComputeGatewayOK) GetPayload() *models.ComputeGatewayResult {
	return o.Payload
}

func (o *GetComputeGatewayOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComputeGatewayResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetComputeGatewayForbidden creates a GetComputeGatewayForbidden with default headers values
func NewGetComputeGatewayForbidden() *GetComputeGatewayForbidden {
	return &GetComputeGatewayForbidden{}
}

/*GetComputeGatewayForbidden handles this case with default header values.

Forbidden
*/
type GetComputeGatewayForbidden struct {
}

func (o *GetComputeGatewayForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/compute-gateways][%d] getComputeGatewayForbidden ", 403)
}

func (o *GetComputeGatewayForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}