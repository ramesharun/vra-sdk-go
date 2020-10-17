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

// UpdateTerraformVersionUsingPATCH1Reader is a Reader for the UpdateTerraformVersionUsingPATCH1 structure.
type UpdateTerraformVersionUsingPATCH1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTerraformVersionUsingPATCH1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateTerraformVersionUsingPATCH1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateTerraformVersionUsingPATCH1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateTerraformVersionUsingPATCH1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateTerraformVersionUsingPATCH1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateTerraformVersionUsingPATCH1OK creates a UpdateTerraformVersionUsingPATCH1OK with default headers values
func NewUpdateTerraformVersionUsingPATCH1OK() *UpdateTerraformVersionUsingPATCH1OK {
	return &UpdateTerraformVersionUsingPATCH1OK{}
}

/*UpdateTerraformVersionUsingPATCH1OK handles this case with default header values.

OK
*/
type UpdateTerraformVersionUsingPATCH1OK struct {
	Payload *models.TerraformVersion
}

func (o *UpdateTerraformVersionUsingPATCH1OK) Error() string {
	return fmt.Sprintf("[PATCH /blueprint/api/blueprint-integrations/terraform/versions/{versionId}][%d] updateTerraformVersionUsingPATCH1OK  %+v", 200, o.Payload)
}

func (o *UpdateTerraformVersionUsingPATCH1OK) GetPayload() *models.TerraformVersion {
	return o.Payload
}

func (o *UpdateTerraformVersionUsingPATCH1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TerraformVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTerraformVersionUsingPATCH1Unauthorized creates a UpdateTerraformVersionUsingPATCH1Unauthorized with default headers values
func NewUpdateTerraformVersionUsingPATCH1Unauthorized() *UpdateTerraformVersionUsingPATCH1Unauthorized {
	return &UpdateTerraformVersionUsingPATCH1Unauthorized{}
}

/*UpdateTerraformVersionUsingPATCH1Unauthorized handles this case with default header values.

Unauthorized
*/
type UpdateTerraformVersionUsingPATCH1Unauthorized struct {
}

func (o *UpdateTerraformVersionUsingPATCH1Unauthorized) Error() string {
	return fmt.Sprintf("[PATCH /blueprint/api/blueprint-integrations/terraform/versions/{versionId}][%d] updateTerraformVersionUsingPATCH1Unauthorized ", 401)
}

func (o *UpdateTerraformVersionUsingPATCH1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateTerraformVersionUsingPATCH1Forbidden creates a UpdateTerraformVersionUsingPATCH1Forbidden with default headers values
func NewUpdateTerraformVersionUsingPATCH1Forbidden() *UpdateTerraformVersionUsingPATCH1Forbidden {
	return &UpdateTerraformVersionUsingPATCH1Forbidden{}
}

/*UpdateTerraformVersionUsingPATCH1Forbidden handles this case with default header values.

Forbidden
*/
type UpdateTerraformVersionUsingPATCH1Forbidden struct {
}

func (o *UpdateTerraformVersionUsingPATCH1Forbidden) Error() string {
	return fmt.Sprintf("[PATCH /blueprint/api/blueprint-integrations/terraform/versions/{versionId}][%d] updateTerraformVersionUsingPATCH1Forbidden ", 403)
}

func (o *UpdateTerraformVersionUsingPATCH1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateTerraformVersionUsingPATCH1NotFound creates a UpdateTerraformVersionUsingPATCH1NotFound with default headers values
func NewUpdateTerraformVersionUsingPATCH1NotFound() *UpdateTerraformVersionUsingPATCH1NotFound {
	return &UpdateTerraformVersionUsingPATCH1NotFound{}
}

/*UpdateTerraformVersionUsingPATCH1NotFound handles this case with default header values.

Not Found
*/
type UpdateTerraformVersionUsingPATCH1NotFound struct {
	Payload *models.Error
}

func (o *UpdateTerraformVersionUsingPATCH1NotFound) Error() string {
	return fmt.Sprintf("[PATCH /blueprint/api/blueprint-integrations/terraform/versions/{versionId}][%d] updateTerraformVersionUsingPATCH1NotFound  %+v", 404, o.Payload)
}

func (o *UpdateTerraformVersionUsingPATCH1NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateTerraformVersionUsingPATCH1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
