// Code generated by go-swagger; DO NOT EDIT.

package blueprint_terraform_integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// CreateBlueprintFromMappingUsingPOST1Reader is a Reader for the CreateBlueprintFromMappingUsingPOST1 structure.
type CreateBlueprintFromMappingUsingPOST1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateBlueprintFromMappingUsingPOST1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateBlueprintFromMappingUsingPOST1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateBlueprintFromMappingUsingPOST1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateBlueprintFromMappingUsingPOST1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateBlueprintFromMappingUsingPOST1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateBlueprintFromMappingUsingPOST1OK creates a CreateBlueprintFromMappingUsingPOST1OK with default headers values
func NewCreateBlueprintFromMappingUsingPOST1OK() *CreateBlueprintFromMappingUsingPOST1OK {
	return &CreateBlueprintFromMappingUsingPOST1OK{}
}

/*CreateBlueprintFromMappingUsingPOST1OK handles this case with default header values.

OK
*/
type CreateBlueprintFromMappingUsingPOST1OK struct {
	Payload *models.Blueprint
}

func (o *CreateBlueprintFromMappingUsingPOST1OK) Error() string {
	return fmt.Sprintf("[POST /blueprint/api/blueprint-integrations/terraform/create-blueprint-from-mapping][%d] createBlueprintFromMappingUsingPOST1OK  %+v", 200, o.Payload)
}

func (o *CreateBlueprintFromMappingUsingPOST1OK) GetPayload() *models.Blueprint {
	return o.Payload
}

func (o *CreateBlueprintFromMappingUsingPOST1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Blueprint)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBlueprintFromMappingUsingPOST1BadRequest creates a CreateBlueprintFromMappingUsingPOST1BadRequest with default headers values
func NewCreateBlueprintFromMappingUsingPOST1BadRequest() *CreateBlueprintFromMappingUsingPOST1BadRequest {
	return &CreateBlueprintFromMappingUsingPOST1BadRequest{}
}

/*CreateBlueprintFromMappingUsingPOST1BadRequest handles this case with default header values.

Bad Request
*/
type CreateBlueprintFromMappingUsingPOST1BadRequest struct {
	Payload *models.Error
}

func (o *CreateBlueprintFromMappingUsingPOST1BadRequest) Error() string {
	return fmt.Sprintf("[POST /blueprint/api/blueprint-integrations/terraform/create-blueprint-from-mapping][%d] createBlueprintFromMappingUsingPOST1BadRequest  %+v", 400, o.Payload)
}

func (o *CreateBlueprintFromMappingUsingPOST1BadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateBlueprintFromMappingUsingPOST1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBlueprintFromMappingUsingPOST1Unauthorized creates a CreateBlueprintFromMappingUsingPOST1Unauthorized with default headers values
func NewCreateBlueprintFromMappingUsingPOST1Unauthorized() *CreateBlueprintFromMappingUsingPOST1Unauthorized {
	return &CreateBlueprintFromMappingUsingPOST1Unauthorized{}
}

/*CreateBlueprintFromMappingUsingPOST1Unauthorized handles this case with default header values.

Unauthorized
*/
type CreateBlueprintFromMappingUsingPOST1Unauthorized struct {
}

func (o *CreateBlueprintFromMappingUsingPOST1Unauthorized) Error() string {
	return fmt.Sprintf("[POST /blueprint/api/blueprint-integrations/terraform/create-blueprint-from-mapping][%d] createBlueprintFromMappingUsingPOST1Unauthorized ", 401)
}

func (o *CreateBlueprintFromMappingUsingPOST1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateBlueprintFromMappingUsingPOST1Forbidden creates a CreateBlueprintFromMappingUsingPOST1Forbidden with default headers values
func NewCreateBlueprintFromMappingUsingPOST1Forbidden() *CreateBlueprintFromMappingUsingPOST1Forbidden {
	return &CreateBlueprintFromMappingUsingPOST1Forbidden{}
}

/*CreateBlueprintFromMappingUsingPOST1Forbidden handles this case with default header values.

Forbidden
*/
type CreateBlueprintFromMappingUsingPOST1Forbidden struct {
}

func (o *CreateBlueprintFromMappingUsingPOST1Forbidden) Error() string {
	return fmt.Sprintf("[POST /blueprint/api/blueprint-integrations/terraform/create-blueprint-from-mapping][%d] createBlueprintFromMappingUsingPOST1Forbidden ", 403)
}

func (o *CreateBlueprintFromMappingUsingPOST1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
