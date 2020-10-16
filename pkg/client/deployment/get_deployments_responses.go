// Code generated by go-swagger; DO NOT EDIT.

package deployment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetDeploymentsReader is a Reader for the GetDeployments structure.
type GetDeploymentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeploymentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeploymentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetDeploymentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDeploymentsOK creates a GetDeploymentsOK with default headers values
func NewGetDeploymentsOK() *GetDeploymentsOK {
	return &GetDeploymentsOK{}
}

/*GetDeploymentsOK handles this case with default header values.

successful operation
*/
type GetDeploymentsOK struct {
	Payload []*models.IaaSDeployment
}

func (o *GetDeploymentsOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/deployments][%d] getDeploymentsOK  %+v", 200, o.Payload)
}

func (o *GetDeploymentsOK) GetPayload() []*models.IaaSDeployment {
	return o.Payload
}

func (o *GetDeploymentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentsForbidden creates a GetDeploymentsForbidden with default headers values
func NewGetDeploymentsForbidden() *GetDeploymentsForbidden {
	return &GetDeploymentsForbidden{}
}

/*GetDeploymentsForbidden handles this case with default header values.

Forbidden
*/
type GetDeploymentsForbidden struct {
}

func (o *GetDeploymentsForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/deployments][%d] getDeploymentsForbidden ", 403)
}

func (o *GetDeploymentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
