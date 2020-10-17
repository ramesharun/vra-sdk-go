// Code generated by go-swagger; DO NOT EDIT.

package blueprint_requests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// DeleteBlueprintRequestUsingDELETE1Reader is a Reader for the DeleteBlueprintRequestUsingDELETE1 structure.
type DeleteBlueprintRequestUsingDELETE1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteBlueprintRequestUsingDELETE1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteBlueprintRequestUsingDELETE1NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteBlueprintRequestUsingDELETE1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteBlueprintRequestUsingDELETE1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteBlueprintRequestUsingDELETE1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteBlueprintRequestUsingDELETE1NoContent creates a DeleteBlueprintRequestUsingDELETE1NoContent with default headers values
func NewDeleteBlueprintRequestUsingDELETE1NoContent() *DeleteBlueprintRequestUsingDELETE1NoContent {
	return &DeleteBlueprintRequestUsingDELETE1NoContent{}
}

/*DeleteBlueprintRequestUsingDELETE1NoContent handles this case with default header values.

No Content
*/
type DeleteBlueprintRequestUsingDELETE1NoContent struct {
}

func (o *DeleteBlueprintRequestUsingDELETE1NoContent) Error() string {
	return fmt.Sprintf("[DELETE /blueprint/api/blueprint-requests/{requestId}][%d] deleteBlueprintRequestUsingDELETE1NoContent ", 204)
}

func (o *DeleteBlueprintRequestUsingDELETE1NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteBlueprintRequestUsingDELETE1Unauthorized creates a DeleteBlueprintRequestUsingDELETE1Unauthorized with default headers values
func NewDeleteBlueprintRequestUsingDELETE1Unauthorized() *DeleteBlueprintRequestUsingDELETE1Unauthorized {
	return &DeleteBlueprintRequestUsingDELETE1Unauthorized{}
}

/*DeleteBlueprintRequestUsingDELETE1Unauthorized handles this case with default header values.

Unauthorized
*/
type DeleteBlueprintRequestUsingDELETE1Unauthorized struct {
}

func (o *DeleteBlueprintRequestUsingDELETE1Unauthorized) Error() string {
	return fmt.Sprintf("[DELETE /blueprint/api/blueprint-requests/{requestId}][%d] deleteBlueprintRequestUsingDELETE1Unauthorized ", 401)
}

func (o *DeleteBlueprintRequestUsingDELETE1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteBlueprintRequestUsingDELETE1Forbidden creates a DeleteBlueprintRequestUsingDELETE1Forbidden with default headers values
func NewDeleteBlueprintRequestUsingDELETE1Forbidden() *DeleteBlueprintRequestUsingDELETE1Forbidden {
	return &DeleteBlueprintRequestUsingDELETE1Forbidden{}
}

/*DeleteBlueprintRequestUsingDELETE1Forbidden handles this case with default header values.

Forbidden
*/
type DeleteBlueprintRequestUsingDELETE1Forbidden struct {
}

func (o *DeleteBlueprintRequestUsingDELETE1Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /blueprint/api/blueprint-requests/{requestId}][%d] deleteBlueprintRequestUsingDELETE1Forbidden ", 403)
}

func (o *DeleteBlueprintRequestUsingDELETE1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteBlueprintRequestUsingDELETE1NotFound creates a DeleteBlueprintRequestUsingDELETE1NotFound with default headers values
func NewDeleteBlueprintRequestUsingDELETE1NotFound() *DeleteBlueprintRequestUsingDELETE1NotFound {
	return &DeleteBlueprintRequestUsingDELETE1NotFound{}
}

/*DeleteBlueprintRequestUsingDELETE1NotFound handles this case with default header values.

Not Found
*/
type DeleteBlueprintRequestUsingDELETE1NotFound struct {
	Payload *models.Error
}

func (o *DeleteBlueprintRequestUsingDELETE1NotFound) Error() string {
	return fmt.Sprintf("[DELETE /blueprint/api/blueprint-requests/{requestId}][%d] deleteBlueprintRequestUsingDELETE1NotFound  %+v", 404, o.Payload)
}

func (o *DeleteBlueprintRequestUsingDELETE1NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteBlueprintRequestUsingDELETE1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
