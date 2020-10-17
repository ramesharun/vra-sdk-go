// Code generated by go-swagger; DO NOT EDIT.

package security_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// CreateOnDemandSecurityGroupReader is a Reader for the CreateOnDemandSecurityGroup structure.
type CreateOnDemandSecurityGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOnDemandSecurityGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewCreateOnDemandSecurityGroupAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateOnDemandSecurityGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateOnDemandSecurityGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateOnDemandSecurityGroupAccepted creates a CreateOnDemandSecurityGroupAccepted with default headers values
func NewCreateOnDemandSecurityGroupAccepted() *CreateOnDemandSecurityGroupAccepted {
	return &CreateOnDemandSecurityGroupAccepted{}
}

/*CreateOnDemandSecurityGroupAccepted handles this case with default header values.

successful operation
*/
type CreateOnDemandSecurityGroupAccepted struct {
	Payload *models.RequestTracker
}

func (o *CreateOnDemandSecurityGroupAccepted) Error() string {
	return fmt.Sprintf("[POST /iaas/api/security-groups][%d] createOnDemandSecurityGroupAccepted  %+v", 202, o.Payload)
}

func (o *CreateOnDemandSecurityGroupAccepted) GetPayload() *models.RequestTracker {
	return o.Payload
}

func (o *CreateOnDemandSecurityGroupAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTracker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOnDemandSecurityGroupBadRequest creates a CreateOnDemandSecurityGroupBadRequest with default headers values
func NewCreateOnDemandSecurityGroupBadRequest() *CreateOnDemandSecurityGroupBadRequest {
	return &CreateOnDemandSecurityGroupBadRequest{}
}

/*CreateOnDemandSecurityGroupBadRequest handles this case with default header values.

Invalid Request - bad data
*/
type CreateOnDemandSecurityGroupBadRequest struct {
	Payload *models.Error
}

func (o *CreateOnDemandSecurityGroupBadRequest) Error() string {
	return fmt.Sprintf("[POST /iaas/api/security-groups][%d] createOnDemandSecurityGroupBadRequest  %+v", 400, o.Payload)
}

func (o *CreateOnDemandSecurityGroupBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateOnDemandSecurityGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOnDemandSecurityGroupForbidden creates a CreateOnDemandSecurityGroupForbidden with default headers values
func NewCreateOnDemandSecurityGroupForbidden() *CreateOnDemandSecurityGroupForbidden {
	return &CreateOnDemandSecurityGroupForbidden{}
}

/*CreateOnDemandSecurityGroupForbidden handles this case with default header values.

Forbidden
*/
type CreateOnDemandSecurityGroupForbidden struct {
}

func (o *CreateOnDemandSecurityGroupForbidden) Error() string {
	return fmt.Sprintf("[POST /iaas/api/security-groups][%d] createOnDemandSecurityGroupForbidden ", 403)
}

func (o *CreateOnDemandSecurityGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
