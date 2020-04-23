// Code generated by go-swagger; DO NOT EDIT.

package fabric_compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new fabric compute API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for fabric compute API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetFabricCompute gets fabric compute

Get fabric compute with a given id
*/
func (a *Client) GetFabricCompute(params *GetFabricComputeParams) (*GetFabricComputeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFabricComputeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFabricCompute",
		Method:             "GET",
		PathPattern:        "/iaas/api/fabric-computes/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFabricComputeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFabricComputeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getFabricCompute: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetFabricComputes gets fabric computes

Get all fabric computes.
*/
func (a *Client) GetFabricComputes(params *GetFabricComputesParams) (*GetFabricComputesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFabricComputesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFabricComputes",
		Method:             "GET",
		PathPattern:        "/iaas/api/fabric-computes",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFabricComputesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFabricComputesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getFabricComputes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateFabricCompute updates fabric compute

Update fabric compute. Only tag updates are supported.
*/
func (a *Client) UpdateFabricCompute(params *UpdateFabricComputeParams) (*UpdateFabricComputeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateFabricComputeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateFabricCompute",
		Method:             "PATCH",
		PathPattern:        "/iaas/api/fabric-computes/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateFabricComputeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateFabricComputeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateFabricCompute: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}