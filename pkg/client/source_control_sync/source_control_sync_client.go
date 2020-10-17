// Code generated by go-swagger; DO NOT EDIT.

package source_control_sync

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new source control sync API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for source control sync API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetSyncRequestUsingGET(params *GetSyncRequestUsingGETParams) (*GetSyncRequestUsingGETOK, error)

	ScheduleSyncAllUsingPOST(params *ScheduleSyncAllUsingPOSTParams) (*ScheduleSyncAllUsingPOSTAccepted, error)

	ScheduleSyncUsingPOST(params *ScheduleSyncUsingPOSTParams) (*ScheduleSyncUsingPOSTAccepted, error)

	SyncHistoryUsingGET(params *SyncHistoryUsingGETParams) (*SyncHistoryUsingGETOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetSyncRequestUsingGET gets a sync request by id
*/
func (a *Client) GetSyncRequestUsingGET(params *GetSyncRequestUsingGETParams) (*GetSyncRequestUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSyncRequestUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSyncRequestUsingGET",
		Method:             "GET",
		PathPattern:        "/content/api/sourcecontrol/sync-requests/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSyncRequestUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSyncRequestUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSyncRequestUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ScheduleSyncAllUsingPOST submits sync request for the projects

  Request sync for the sources belonging to the user's projects or specified projects
*/
func (a *Client) ScheduleSyncAllUsingPOST(params *ScheduleSyncAllUsingPOSTParams) (*ScheduleSyncAllUsingPOSTAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewScheduleSyncAllUsingPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "scheduleSyncAllUsingPOST",
		Method:             "POST",
		PathPattern:        "/content/api/sourcecontrol/sync-all-requests",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ScheduleSyncAllUsingPOSTReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ScheduleSyncAllUsingPOSTAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for scheduleSyncAllUsingPOST: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ScheduleSyncUsingPOST creates a sync request

  Request sync for a content source
*/
func (a *Client) ScheduleSyncUsingPOST(params *ScheduleSyncUsingPOSTParams) (*ScheduleSyncUsingPOSTAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewScheduleSyncUsingPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "scheduleSyncUsingPOST",
		Method:             "POST",
		PathPattern:        "/content/api/sourcecontrol/sync-requests",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ScheduleSyncUsingPOSTReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ScheduleSyncUsingPOSTAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for scheduleSyncUsingPOST: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SyncHistoryUsingGET gets sync history

  Get history for files synced via source control sync
*/
func (a *Client) SyncHistoryUsingGET(params *SyncHistoryUsingGETParams) (*SyncHistoryUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSyncHistoryUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "syncHistoryUsingGET",
		Method:             "GET",
		PathPattern:        "/content/api/sourcecontrol/sync-history",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SyncHistoryUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SyncHistoryUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for syncHistoryUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
