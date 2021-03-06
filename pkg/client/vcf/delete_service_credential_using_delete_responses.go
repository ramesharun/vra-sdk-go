// Code generated by go-swagger; DO NOT EDIT.

package vcf

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteServiceCredentialUsingDELETEReader is a Reader for the DeleteServiceCredentialUsingDELETE structure.
type DeleteServiceCredentialUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteServiceCredentialUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteServiceCredentialUsingDELETENoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteServiceCredentialUsingDELETENoContent creates a DeleteServiceCredentialUsingDELETENoContent with default headers values
func NewDeleteServiceCredentialUsingDELETENoContent() *DeleteServiceCredentialUsingDELETENoContent {
	return &DeleteServiceCredentialUsingDELETENoContent{}
}

/*DeleteServiceCredentialUsingDELETENoContent handles this case with default header values.

No Content
*/
type DeleteServiceCredentialUsingDELETENoContent struct {
}

func (o *DeleteServiceCredentialUsingDELETENoContent) Error() string {
	return fmt.Sprintf("[DELETE /content/api/vcf/{integrationId}/domain/{domainId}/service-accounts/{id}][%d] deleteServiceCredentialUsingDELETENoContent ", 204)
}

func (o *DeleteServiceCredentialUsingDELETENoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
