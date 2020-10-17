// Code generated by go-swagger; DO NOT EDIT.

package catalog_items

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetCatalogItemsUsingGET1Reader is a Reader for the GetCatalogItemsUsingGET1 structure.
type GetCatalogItemsUsingGET1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCatalogItemsUsingGET1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCatalogItemsUsingGET1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetCatalogItemsUsingGET1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCatalogItemsUsingGET1OK creates a GetCatalogItemsUsingGET1OK with default headers values
func NewGetCatalogItemsUsingGET1OK() *GetCatalogItemsUsingGET1OK {
	return &GetCatalogItemsUsingGET1OK{}
}

/*GetCatalogItemsUsingGET1OK handles this case with default header values.

OK
*/
type GetCatalogItemsUsingGET1OK struct {
	Payload *models.PageOfCatalogItem
}

func (o *GetCatalogItemsUsingGET1OK) Error() string {
	return fmt.Sprintf("[GET /catalog/api/items][%d] getCatalogItemsUsingGET1OK  %+v", 200, o.Payload)
}

func (o *GetCatalogItemsUsingGET1OK) GetPayload() *models.PageOfCatalogItem {
	return o.Payload
}

func (o *GetCatalogItemsUsingGET1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfCatalogItem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCatalogItemsUsingGET1Unauthorized creates a GetCatalogItemsUsingGET1Unauthorized with default headers values
func NewGetCatalogItemsUsingGET1Unauthorized() *GetCatalogItemsUsingGET1Unauthorized {
	return &GetCatalogItemsUsingGET1Unauthorized{}
}

/*GetCatalogItemsUsingGET1Unauthorized handles this case with default header values.

Unauthorized
*/
type GetCatalogItemsUsingGET1Unauthorized struct {
}

func (o *GetCatalogItemsUsingGET1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /catalog/api/items][%d] getCatalogItemsUsingGET1Unauthorized ", 401)
}

func (o *GetCatalogItemsUsingGET1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
