// Code generated by go-swagger; DO NOT EDIT.

package blueprint_requests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// ListBlueprintRequestsUsingGET1Reader is a Reader for the ListBlueprintRequestsUsingGET1 structure.
type ListBlueprintRequestsUsingGET1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListBlueprintRequestsUsingGET1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListBlueprintRequestsUsingGET1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListBlueprintRequestsUsingGET1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListBlueprintRequestsUsingGET1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListBlueprintRequestsUsingGET1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListBlueprintRequestsUsingGET1OK creates a ListBlueprintRequestsUsingGET1OK with default headers values
func NewListBlueprintRequestsUsingGET1OK() *ListBlueprintRequestsUsingGET1OK {
	return &ListBlueprintRequestsUsingGET1OK{}
}

/*ListBlueprintRequestsUsingGET1OK handles this case with default header values.

OK
*/
type ListBlueprintRequestsUsingGET1OK struct {
	Payload *models.PageOfBlueprintRequest
}

func (o *ListBlueprintRequestsUsingGET1OK) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-requests][%d] listBlueprintRequestsUsingGET1OK  %+v", 200, o.Payload)
}

func (o *ListBlueprintRequestsUsingGET1OK) GetPayload() *models.PageOfBlueprintRequest {
	return o.Payload
}

func (o *ListBlueprintRequestsUsingGET1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfBlueprintRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListBlueprintRequestsUsingGET1BadRequest creates a ListBlueprintRequestsUsingGET1BadRequest with default headers values
func NewListBlueprintRequestsUsingGET1BadRequest() *ListBlueprintRequestsUsingGET1BadRequest {
	return &ListBlueprintRequestsUsingGET1BadRequest{}
}

/*ListBlueprintRequestsUsingGET1BadRequest handles this case with default header values.

Bad Request
*/
type ListBlueprintRequestsUsingGET1BadRequest struct {
}

func (o *ListBlueprintRequestsUsingGET1BadRequest) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-requests][%d] listBlueprintRequestsUsingGET1BadRequest ", 400)
}

func (o *ListBlueprintRequestsUsingGET1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListBlueprintRequestsUsingGET1Unauthorized creates a ListBlueprintRequestsUsingGET1Unauthorized with default headers values
func NewListBlueprintRequestsUsingGET1Unauthorized() *ListBlueprintRequestsUsingGET1Unauthorized {
	return &ListBlueprintRequestsUsingGET1Unauthorized{}
}

/*ListBlueprintRequestsUsingGET1Unauthorized handles this case with default header values.

Unauthorized
*/
type ListBlueprintRequestsUsingGET1Unauthorized struct {
}

func (o *ListBlueprintRequestsUsingGET1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-requests][%d] listBlueprintRequestsUsingGET1Unauthorized ", 401)
}

func (o *ListBlueprintRequestsUsingGET1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListBlueprintRequestsUsingGET1Forbidden creates a ListBlueprintRequestsUsingGET1Forbidden with default headers values
func NewListBlueprintRequestsUsingGET1Forbidden() *ListBlueprintRequestsUsingGET1Forbidden {
	return &ListBlueprintRequestsUsingGET1Forbidden{}
}

/*ListBlueprintRequestsUsingGET1Forbidden handles this case with default header values.

Forbidden
*/
type ListBlueprintRequestsUsingGET1Forbidden struct {
}

func (o *ListBlueprintRequestsUsingGET1Forbidden) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-requests][%d] listBlueprintRequestsUsingGET1Forbidden ", 403)
}

func (o *ListBlueprintRequestsUsingGET1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}