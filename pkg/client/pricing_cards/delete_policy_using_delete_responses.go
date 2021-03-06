// Code generated by go-swagger; DO NOT EDIT.

package pricing_cards

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// DeletePolicyUsingDELETEReader is a Reader for the DeletePolicyUsingDELETE structure.
type DeletePolicyUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePolicyUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeletePolicyUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeletePolicyUsingDELETENoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeletePolicyUsingDELETEUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeletePolicyUsingDELETENotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeletePolicyUsingDELETEOK creates a DeletePolicyUsingDELETEOK with default headers values
func NewDeletePolicyUsingDELETEOK() *DeletePolicyUsingDELETEOK {
	return &DeletePolicyUsingDELETEOK{}
}

/*DeletePolicyUsingDELETEOK handles this case with default header values.

OK
*/
type DeletePolicyUsingDELETEOK struct {
}

func (o *DeletePolicyUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /price/api/private/pricing-cards/{id}][%d] deletePolicyUsingDELETEOK ", 200)
}

func (o *DeletePolicyUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePolicyUsingDELETENoContent creates a DeletePolicyUsingDELETENoContent with default headers values
func NewDeletePolicyUsingDELETENoContent() *DeletePolicyUsingDELETENoContent {
	return &DeletePolicyUsingDELETENoContent{}
}

/*DeletePolicyUsingDELETENoContent handles this case with default header values.

No Content
*/
type DeletePolicyUsingDELETENoContent struct {
}

func (o *DeletePolicyUsingDELETENoContent) Error() string {
	return fmt.Sprintf("[DELETE /price/api/private/pricing-cards/{id}][%d] deletePolicyUsingDELETENoContent ", 204)
}

func (o *DeletePolicyUsingDELETENoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePolicyUsingDELETEUnauthorized creates a DeletePolicyUsingDELETEUnauthorized with default headers values
func NewDeletePolicyUsingDELETEUnauthorized() *DeletePolicyUsingDELETEUnauthorized {
	return &DeletePolicyUsingDELETEUnauthorized{}
}

/*DeletePolicyUsingDELETEUnauthorized handles this case with default header values.

Unauthorized
*/
type DeletePolicyUsingDELETEUnauthorized struct {
}

func (o *DeletePolicyUsingDELETEUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /price/api/private/pricing-cards/{id}][%d] deletePolicyUsingDELETEUnauthorized ", 401)
}

func (o *DeletePolicyUsingDELETEUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePolicyUsingDELETENotFound creates a DeletePolicyUsingDELETENotFound with default headers values
func NewDeletePolicyUsingDELETENotFound() *DeletePolicyUsingDELETENotFound {
	return &DeletePolicyUsingDELETENotFound{}
}

/*DeletePolicyUsingDELETENotFound handles this case with default header values.

Not Found
*/
type DeletePolicyUsingDELETENotFound struct {
	Payload *models.Error
}

func (o *DeletePolicyUsingDELETENotFound) Error() string {
	return fmt.Sprintf("[DELETE /price/api/private/pricing-cards/{id}][%d] deletePolicyUsingDELETENotFound  %+v", 404, o.Payload)
}

func (o *DeletePolicyUsingDELETENotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeletePolicyUsingDELETENotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
