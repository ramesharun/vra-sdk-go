// Code generated by go-swagger; DO NOT EDIT.

package policy_decisions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetDecisionsUsingGETReader is a Reader for the GetDecisionsUsingGET structure.
type GetDecisionsUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDecisionsUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDecisionsUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDecisionsUsingGETOK creates a GetDecisionsUsingGETOK with default headers values
func NewGetDecisionsUsingGETOK() *GetDecisionsUsingGETOK {
	return &GetDecisionsUsingGETOK{}
}

/*GetDecisionsUsingGETOK handles this case with default header values.

OK
*/
type GetDecisionsUsingGETOK struct {
	Payload *models.PageOfPolicyDecisionOfObjectNode
}

func (o *GetDecisionsUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /policy/api/policyDecisions][%d] getDecisionsUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetDecisionsUsingGETOK) GetPayload() *models.PageOfPolicyDecisionOfObjectNode {
	return o.Payload
}

func (o *GetDecisionsUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfPolicyDecisionOfObjectNode)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
