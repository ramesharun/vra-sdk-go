// Code generated by go-swagger; DO NOT EDIT.

package pricing_card_assignments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// CreateMeteringPolicyAssignmentUsingPOSTReader is a Reader for the CreateMeteringPolicyAssignmentUsingPOST structure.
type CreateMeteringPolicyAssignmentUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMeteringPolicyAssignmentUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateMeteringPolicyAssignmentUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateMeteringPolicyAssignmentUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateMeteringPolicyAssignmentUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateMeteringPolicyAssignmentUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateMeteringPolicyAssignmentUsingPOSTOK creates a CreateMeteringPolicyAssignmentUsingPOSTOK with default headers values
func NewCreateMeteringPolicyAssignmentUsingPOSTOK() *CreateMeteringPolicyAssignmentUsingPOSTOK {
	return &CreateMeteringPolicyAssignmentUsingPOSTOK{}
}

/*CreateMeteringPolicyAssignmentUsingPOSTOK handles this case with default header values.

OK
*/
type CreateMeteringPolicyAssignmentUsingPOSTOK struct {
	Payload *models.MeteringPolicyAssignment
}

func (o *CreateMeteringPolicyAssignmentUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /price/api/private/pricing-card-assignments][%d] createMeteringPolicyAssignmentUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *CreateMeteringPolicyAssignmentUsingPOSTOK) GetPayload() *models.MeteringPolicyAssignment {
	return o.Payload
}

func (o *CreateMeteringPolicyAssignmentUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MeteringPolicyAssignment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMeteringPolicyAssignmentUsingPOSTCreated creates a CreateMeteringPolicyAssignmentUsingPOSTCreated with default headers values
func NewCreateMeteringPolicyAssignmentUsingPOSTCreated() *CreateMeteringPolicyAssignmentUsingPOSTCreated {
	return &CreateMeteringPolicyAssignmentUsingPOSTCreated{}
}

/*CreateMeteringPolicyAssignmentUsingPOSTCreated handles this case with default header values.

Created
*/
type CreateMeteringPolicyAssignmentUsingPOSTCreated struct {
	Payload *models.MeteringPolicyAssignment
}

func (o *CreateMeteringPolicyAssignmentUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /price/api/private/pricing-card-assignments][%d] createMeteringPolicyAssignmentUsingPOSTCreated  %+v", 201, o.Payload)
}

func (o *CreateMeteringPolicyAssignmentUsingPOSTCreated) GetPayload() *models.MeteringPolicyAssignment {
	return o.Payload
}

func (o *CreateMeteringPolicyAssignmentUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MeteringPolicyAssignment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMeteringPolicyAssignmentUsingPOSTUnauthorized creates a CreateMeteringPolicyAssignmentUsingPOSTUnauthorized with default headers values
func NewCreateMeteringPolicyAssignmentUsingPOSTUnauthorized() *CreateMeteringPolicyAssignmentUsingPOSTUnauthorized {
	return &CreateMeteringPolicyAssignmentUsingPOSTUnauthorized{}
}

/*CreateMeteringPolicyAssignmentUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateMeteringPolicyAssignmentUsingPOSTUnauthorized struct {
}

func (o *CreateMeteringPolicyAssignmentUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /price/api/private/pricing-card-assignments][%d] createMeteringPolicyAssignmentUsingPOSTUnauthorized ", 401)
}

func (o *CreateMeteringPolicyAssignmentUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateMeteringPolicyAssignmentUsingPOSTForbidden creates a CreateMeteringPolicyAssignmentUsingPOSTForbidden with default headers values
func NewCreateMeteringPolicyAssignmentUsingPOSTForbidden() *CreateMeteringPolicyAssignmentUsingPOSTForbidden {
	return &CreateMeteringPolicyAssignmentUsingPOSTForbidden{}
}

/*CreateMeteringPolicyAssignmentUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type CreateMeteringPolicyAssignmentUsingPOSTForbidden struct {
}

func (o *CreateMeteringPolicyAssignmentUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /price/api/private/pricing-card-assignments][%d] createMeteringPolicyAssignmentUsingPOSTForbidden ", 403)
}

func (o *CreateMeteringPolicyAssignmentUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
