// Code generated by go-swagger; DO NOT EDIT.

package image_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteImageProfileReader is a Reader for the DeleteImageProfile structure.
type DeleteImageProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteImageProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteImageProfileNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteImageProfileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteImageProfileNoContent creates a DeleteImageProfileNoContent with default headers values
func NewDeleteImageProfileNoContent() *DeleteImageProfileNoContent {
	return &DeleteImageProfileNoContent{}
}

/*DeleteImageProfileNoContent handles this case with default header values.

No Content
*/
type DeleteImageProfileNoContent struct {
}

func (o *DeleteImageProfileNoContent) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/image-profiles/{id}][%d] deleteImageProfileNoContent ", 204)
}

func (o *DeleteImageProfileNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteImageProfileForbidden creates a DeleteImageProfileForbidden with default headers values
func NewDeleteImageProfileForbidden() *DeleteImageProfileForbidden {
	return &DeleteImageProfileForbidden{}
}

/*DeleteImageProfileForbidden handles this case with default header values.

Forbidden
*/
type DeleteImageProfileForbidden struct {
}

func (o *DeleteImageProfileForbidden) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/image-profiles/{id}][%d] deleteImageProfileForbidden ", 403)
}

func (o *DeleteImageProfileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
