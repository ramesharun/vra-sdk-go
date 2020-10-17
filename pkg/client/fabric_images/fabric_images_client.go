// Code generated by go-swagger; DO NOT EDIT.

package fabric_images

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new fabric images API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for fabric images API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetFabricImage(params *GetFabricImageParams) (*GetFabricImageOK, error)

	GetFabricImages(params *GetFabricImagesParams) (*GetFabricImagesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetFabricImage gets fabric image

  Get fabric image with a given id
*/
func (a *Client) GetFabricImage(params *GetFabricImageParams) (*GetFabricImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFabricImageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFabricImage",
		Method:             "GET",
		PathPattern:        "/iaas/api/fabric-images/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFabricImageReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFabricImageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getFabricImage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetFabricImages gets fabric images

  Get all fabric images
*/
func (a *Client) GetFabricImages(params *GetFabricImagesParams) (*GetFabricImagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFabricImagesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFabricImages",
		Method:             "GET",
		PathPattern:        "/iaas/api/fabric-images",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFabricImagesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFabricImagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getFabricImages: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
