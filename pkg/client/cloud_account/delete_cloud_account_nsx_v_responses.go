// Code generated by go-swagger; DO NOT EDIT.

package cloud_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteCloudAccountNsxVReader is a Reader for the DeleteCloudAccountNsxV structure.
type DeleteCloudAccountNsxVReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCloudAccountNsxVReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteCloudAccountNsxVNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteCloudAccountNsxVForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteCloudAccountNsxVNoContent creates a DeleteCloudAccountNsxVNoContent with default headers values
func NewDeleteCloudAccountNsxVNoContent() *DeleteCloudAccountNsxVNoContent {
	return &DeleteCloudAccountNsxVNoContent{}
}

/*DeleteCloudAccountNsxVNoContent handles this case with default header values.

No Content
*/
type DeleteCloudAccountNsxVNoContent struct {
}

func (o *DeleteCloudAccountNsxVNoContent) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/cloud-accounts-nsx-v/{id}][%d] deleteCloudAccountNsxVNoContent ", 204)
}

func (o *DeleteCloudAccountNsxVNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCloudAccountNsxVForbidden creates a DeleteCloudAccountNsxVForbidden with default headers values
func NewDeleteCloudAccountNsxVForbidden() *DeleteCloudAccountNsxVForbidden {
	return &DeleteCloudAccountNsxVForbidden{}
}

/*DeleteCloudAccountNsxVForbidden handles this case with default header values.

Forbidden
*/
type DeleteCloudAccountNsxVForbidden struct {
}

func (o *DeleteCloudAccountNsxVForbidden) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/cloud-accounts-nsx-v/{id}][%d] deleteCloudAccountNsxVForbidden ", 403)
}

func (o *DeleteCloudAccountNsxVForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
