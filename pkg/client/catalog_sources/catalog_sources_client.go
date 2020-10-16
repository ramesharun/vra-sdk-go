// Code generated by go-swagger; DO NOT EDIT.

package catalog_sources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new catalog sources API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for catalog sources API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteUsingDELETE(params *DeleteUsingDELETEParams) (*DeleteUsingDELETENoContent, error)

	GetPageUsingGET(params *GetPageUsingGETParams) (*GetPageUsingGETOK, error)

	GetUsingGET(params *GetUsingGETParams) (*GetUsingGETOK, error)

	PostUsingPOST(params *PostUsingPOSTParams) (*PostUsingPOSTOK, *PostUsingPOSTCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteUsingDELETE deletes catalog source

  Deletes the catalog source with the supplied ID.
*/
func (a *Client) DeleteUsingDELETE(params *DeleteUsingDELETEParams) (*DeleteUsingDELETENoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUsingDELETEParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteUsingDELETE",
		Method:             "DELETE",
		PathPattern:        "/catalog/api/admin/sources/{sourceId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteUsingDELETEReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteUsingDELETENoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteUsingDELETE: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetPageUsingGET fetches catalog sources

  Returns a paginated list of catalog sources.
*/
func (a *Client) GetPageUsingGET(params *GetPageUsingGETParams) (*GetPageUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPageUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPageUsingGET",
		Method:             "GET",
		PathPattern:        "/catalog/api/admin/sources",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPageUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPageUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPageUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetUsingGET fetches a specific catalog source for the given ID

  Returns the catalog source with the supplied ID.
*/
func (a *Client) GetUsingGET(params *GetUsingGETParams) (*GetUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUsingGET",
		Method:             "GET",
		PathPattern:        "/catalog/api/admin/sources/{sourceId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostUsingPOST creates a new catalog source

  Creates a catalog source and imports catalog items from it.
*/
func (a *Client) PostUsingPOST(params *PostUsingPOSTParams) (*PostUsingPOSTOK, *PostUsingPOSTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostUsingPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postUsingPOST",
		Method:             "POST",
		PathPattern:        "/catalog/api/admin/sources",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostUsingPOSTReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PostUsingPOSTOK:
		return value, nil, nil
	case *PostUsingPOSTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for catalog_sources: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
