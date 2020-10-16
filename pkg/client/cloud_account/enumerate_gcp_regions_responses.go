// Code generated by go-swagger; DO NOT EDIT.

package cloud_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// EnumerateGcpRegionsReader is a Reader for the EnumerateGcpRegions structure.
type EnumerateGcpRegionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnumerateGcpRegionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEnumerateGcpRegionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEnumerateGcpRegionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEnumerateGcpRegionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEnumerateGcpRegionsOK creates a EnumerateGcpRegionsOK with default headers values
func NewEnumerateGcpRegionsOK() *EnumerateGcpRegionsOK {
	return &EnumerateGcpRegionsOK{}
}

/*EnumerateGcpRegionsOK handles this case with default header values.

successful operation
*/
type EnumerateGcpRegionsOK struct {
	Payload *models.CloudAccountRegions
}

func (o *EnumerateGcpRegionsOK) Error() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-gcp/region-enumeration][%d] enumerateGcpRegionsOK  %+v", 200, o.Payload)
}

func (o *EnumerateGcpRegionsOK) GetPayload() *models.CloudAccountRegions {
	return o.Payload
}

func (o *EnumerateGcpRegionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudAccountRegions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnumerateGcpRegionsBadRequest creates a EnumerateGcpRegionsBadRequest with default headers values
func NewEnumerateGcpRegionsBadRequest() *EnumerateGcpRegionsBadRequest {
	return &EnumerateGcpRegionsBadRequest{}
}

/*EnumerateGcpRegionsBadRequest handles this case with default header values.

Invalid Request - bad data
*/
type EnumerateGcpRegionsBadRequest struct {
	Payload *models.Error
}

func (o *EnumerateGcpRegionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-gcp/region-enumeration][%d] enumerateGcpRegionsBadRequest  %+v", 400, o.Payload)
}

func (o *EnumerateGcpRegionsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *EnumerateGcpRegionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnumerateGcpRegionsForbidden creates a EnumerateGcpRegionsForbidden with default headers values
func NewEnumerateGcpRegionsForbidden() *EnumerateGcpRegionsForbidden {
	return &EnumerateGcpRegionsForbidden{}
}

/*EnumerateGcpRegionsForbidden handles this case with default header values.

Forbidden
*/
type EnumerateGcpRegionsForbidden struct {
}

func (o *EnumerateGcpRegionsForbidden) Error() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-gcp/region-enumeration][%d] enumerateGcpRegionsForbidden ", 403)
}

func (o *EnumerateGcpRegionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
