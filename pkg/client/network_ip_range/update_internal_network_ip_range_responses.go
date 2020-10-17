// Code generated by go-swagger; DO NOT EDIT.

package network_ip_range

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// UpdateInternalNetworkIPRangeReader is a Reader for the UpdateInternalNetworkIPRange structure.
type UpdateInternalNetworkIPRangeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateInternalNetworkIPRangeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateInternalNetworkIPRangeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateInternalNetworkIPRangeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateInternalNetworkIPRangeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateInternalNetworkIPRangeOK creates a UpdateInternalNetworkIPRangeOK with default headers values
func NewUpdateInternalNetworkIPRangeOK() *UpdateInternalNetworkIPRangeOK {
	return &UpdateInternalNetworkIPRangeOK{}
}

/*UpdateInternalNetworkIPRangeOK handles this case with default header values.

successful operation
*/
type UpdateInternalNetworkIPRangeOK struct {
	Payload *models.NetworkIPRange
}

func (o *UpdateInternalNetworkIPRangeOK) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/network-ip-ranges/{id}][%d] updateInternalNetworkIpRangeOK  %+v", 200, o.Payload)
}

func (o *UpdateInternalNetworkIPRangeOK) GetPayload() *models.NetworkIPRange {
	return o.Payload
}

func (o *UpdateInternalNetworkIPRangeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkIPRange)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInternalNetworkIPRangeForbidden creates a UpdateInternalNetworkIPRangeForbidden with default headers values
func NewUpdateInternalNetworkIPRangeForbidden() *UpdateInternalNetworkIPRangeForbidden {
	return &UpdateInternalNetworkIPRangeForbidden{}
}

/*UpdateInternalNetworkIPRangeForbidden handles this case with default header values.

Forbidden
*/
type UpdateInternalNetworkIPRangeForbidden struct {
}

func (o *UpdateInternalNetworkIPRangeForbidden) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/network-ip-ranges/{id}][%d] updateInternalNetworkIpRangeForbidden ", 403)
}

func (o *UpdateInternalNetworkIPRangeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateInternalNetworkIPRangeNotFound creates a UpdateInternalNetworkIPRangeNotFound with default headers values
func NewUpdateInternalNetworkIPRangeNotFound() *UpdateInternalNetworkIPRangeNotFound {
	return &UpdateInternalNetworkIPRangeNotFound{}
}

/*UpdateInternalNetworkIPRangeNotFound handles this case with default header values.

Not Found
*/
type UpdateInternalNetworkIPRangeNotFound struct {
	Payload *models.Error
}

func (o *UpdateInternalNetworkIPRangeNotFound) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/network-ip-ranges/{id}][%d] updateInternalNetworkIpRangeNotFound  %+v", 404, o.Payload)
}

func (o *UpdateInternalNetworkIPRangeNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateInternalNetworkIPRangeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
