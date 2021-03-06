// Code generated by go-swagger; DO NOT EDIT.

package marketplace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// SearchReader is a Reader for the Search structure.
type SearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewSearchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchOK creates a SearchOK with default headers values
func NewSearchOK() *SearchOK {
	return &SearchOK{}
}

/*SearchOK handles this case with default header values.

Search result
*/
type SearchOK struct {
	Payload *models.MarketplaceContents
}

func (o *SearchOK) Error() string {
	return fmt.Sprintf("[GET /content/api/marketplace/sources/{sourceId}/contents][%d] searchOK  %+v", 200, o.Payload)
}

func (o *SearchOK) GetPayload() *models.MarketplaceContents {
	return o.Payload
}

func (o *SearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MarketplaceContents)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchNotFound creates a SearchNotFound with default headers values
func NewSearchNotFound() *SearchNotFound {
	return &SearchNotFound{}
}

/*SearchNotFound handles this case with default header values.

Source not found
*/
type SearchNotFound struct {
	Payload *models.Error
}

func (o *SearchNotFound) Error() string {
	return fmt.Sprintf("[GET /content/api/marketplace/sources/{sourceId}/contents][%d] searchNotFound  %+v", 404, o.Payload)
}

func (o *SearchNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *SearchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
