// Code generated by go-swagger; DO NOT EDIT.

package catalog_item_types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new catalog item types API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for catalog item types API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetTypeByIDUsingGET(params *GetTypeByIDUsingGETParams) (*GetTypeByIDUsingGETOK, error)

	GetTypesUsingGET(params *GetTypesUsingGETParams) (*GetTypesUsingGETOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetTypeByIDUsingGET fetches catalog item type associated with the specified ID

  Returns the Catalog Item Type with the specified ID.
*/
func (a *Client) GetTypeByIDUsingGET(params *GetTypeByIDUsingGETParams) (*GetTypeByIDUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTypeByIDUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTypeByIdUsingGET",
		Method:             "GET",
		PathPattern:        "/catalog/api/types/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTypeByIDUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTypeByIDUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTypeByIdUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetTypesUsingGET finds all catalog item types

  Returns a paginated list of all available Catalog Item Types.
*/
func (a *Client) GetTypesUsingGET(params *GetTypesUsingGETParams) (*GetTypesUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTypesUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTypesUsingGET",
		Method:             "GET",
		PathPattern:        "/catalog/api/types",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTypesUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTypesUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTypesUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
