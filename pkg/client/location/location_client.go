// Code generated by go-swagger; DO NOT EDIT.

package location

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new location API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for location API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateZone(params *CreateZoneParams) (*CreateZoneCreated, error)

	DeleteZone(params *DeleteZoneParams) (*DeleteZoneNoContent, error)

	GetComputes(params *GetComputesParams) (*GetComputesOK, error)

	GetRegion(params *GetRegionParams) (*GetRegionOK, error)

	GetRegions(params *GetRegionsParams) (*GetRegionsOK, error)

	GetZone(params *GetZoneParams) (*GetZoneOK, error)

	GetZones(params *GetZonesParams) (*GetZonesOK, error)

	UpdateZone(params *UpdateZoneParams) (*UpdateZoneOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateZone creates zone

  Create zone
*/
func (a *Client) CreateZone(params *CreateZoneParams) (*CreateZoneCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateZoneParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createZone",
		Method:             "POST",
		PathPattern:        "/iaas/api/zones",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateZoneReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateZoneCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createZone: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteZone deletes a zone

  Delete a zone
*/
func (a *Client) DeleteZone(params *DeleteZoneParams) (*DeleteZoneNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteZoneParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteZone",
		Method:             "DELETE",
		PathPattern:        "/iaas/api/zones/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteZoneReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteZoneNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteZone: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetComputes gets computes

  Get zone's computes by given zone ID
*/
func (a *Client) GetComputes(params *GetComputesParams) (*GetComputesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetComputesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getComputes",
		Method:             "GET",
		PathPattern:        "/iaas/api/zones/{id}/computes",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetComputesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetComputesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getComputes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetRegion gets region

  Get Region with a given id
*/
func (a *Client) GetRegion(params *GetRegionParams) (*GetRegionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRegionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRegion",
		Method:             "GET",
		PathPattern:        "/iaas/api/regions/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRegionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRegionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getRegion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetRegions gets regions

  Get all regions
*/
func (a *Client) GetRegions(params *GetRegionsParams) (*GetRegionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRegionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRegions",
		Method:             "GET",
		PathPattern:        "/iaas/api/regions",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRegionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRegionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getRegions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetZone gets zone

  Get zone with given id
*/
func (a *Client) GetZone(params *GetZoneParams) (*GetZoneOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetZoneParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getZone",
		Method:             "GET",
		PathPattern:        "/iaas/api/zones/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetZoneReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetZoneOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getZone: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetZones gets zones

  Get all zones
*/
func (a *Client) GetZones(params *GetZonesParams) (*GetZonesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetZonesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getZones",
		Method:             "GET",
		PathPattern:        "/iaas/api/zones",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetZonesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetZonesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getZones: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateZone updates zone

  Update zone
*/
func (a *Client) UpdateZone(params *UpdateZoneParams) (*UpdateZoneOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateZoneParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateZone",
		Method:             "PATCH",
		PathPattern:        "/iaas/api/zones/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateZoneReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateZoneOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateZone: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
