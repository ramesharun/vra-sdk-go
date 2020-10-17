// Code generated by go-swagger; DO NOT EDIT.

package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetDeploymentFiltersUsingGETReader is a Reader for the GetDeploymentFiltersUsingGET structure.
type GetDeploymentFiltersUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeploymentFiltersUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeploymentFiltersUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetDeploymentFiltersUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDeploymentFiltersUsingGETOK creates a GetDeploymentFiltersUsingGETOK with default headers values
func NewGetDeploymentFiltersUsingGETOK() *GetDeploymentFiltersUsingGETOK {
	return &GetDeploymentFiltersUsingGETOK{}
}

/*GetDeploymentFiltersUsingGETOK handles this case with default header values.

OK
*/
type GetDeploymentFiltersUsingGETOK struct {
	Payload *models.FilterSchema
}

func (o *GetDeploymentFiltersUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/filters][%d] getDeploymentFiltersUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetDeploymentFiltersUsingGETOK) GetPayload() *models.FilterSchema {
	return o.Payload
}

func (o *GetDeploymentFiltersUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FilterSchema)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentFiltersUsingGETUnauthorized creates a GetDeploymentFiltersUsingGETUnauthorized with default headers values
func NewGetDeploymentFiltersUsingGETUnauthorized() *GetDeploymentFiltersUsingGETUnauthorized {
	return &GetDeploymentFiltersUsingGETUnauthorized{}
}

/*GetDeploymentFiltersUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetDeploymentFiltersUsingGETUnauthorized struct {
}

func (o *GetDeploymentFiltersUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/filters][%d] getDeploymentFiltersUsingGETUnauthorized ", 401)
}

func (o *GetDeploymentFiltersUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
