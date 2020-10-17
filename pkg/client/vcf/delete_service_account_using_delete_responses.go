// Code generated by go-swagger; DO NOT EDIT.

package vcf

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteServiceAccountUsingDELETEReader is a Reader for the DeleteServiceAccountUsingDELETE structure.
type DeleteServiceAccountUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteServiceAccountUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteServiceAccountUsingDELETENoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteServiceAccountUsingDELETENoContent creates a DeleteServiceAccountUsingDELETENoContent with default headers values
func NewDeleteServiceAccountUsingDELETENoContent() *DeleteServiceAccountUsingDELETENoContent {
	return &DeleteServiceAccountUsingDELETENoContent{}
}

/*DeleteServiceAccountUsingDELETENoContent handles this case with default header values.

No Content
*/
type DeleteServiceAccountUsingDELETENoContent struct {
}

func (o *DeleteServiceAccountUsingDELETENoContent) Error() string {
	return fmt.Sprintf("[DELETE /content/api/vcf/{integrationId}/domain/{domainId}/service-accounts][%d] deleteServiceAccountUsingDELETENoContent ", 204)
}

func (o *DeleteServiceAccountUsingDELETENoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
