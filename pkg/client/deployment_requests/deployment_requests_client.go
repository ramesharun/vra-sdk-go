// Code generated by go-swagger; DO NOT EDIT.

package deployment_requests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new deployment requests API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for deployment requests API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetDeploymentRequestsUsingGET(params *GetDeploymentRequestsUsingGETParams) (*GetDeploymentRequestsUsingGETOK, error)

	GetEventLogsUsingGET(params *GetEventLogsUsingGETParams) (*GetEventLogsUsingGETOK, error)

	GetRequestEventsUsingGET(params *GetRequestEventsUsingGETParams) (*GetRequestEventsUsingGETOK, error)

	GetRequestUsingGET(params *GetRequestUsingGETParams) (*GetRequestUsingGETOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetDeploymentRequestsUsingGET fetches deployment requests

  Returns the requests for the deployment.
*/
func (a *Client) GetDeploymentRequestsUsingGET(params *GetDeploymentRequestsUsingGETParams) (*GetDeploymentRequestsUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeploymentRequestsUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDeploymentRequestsUsingGET",
		Method:             "GET",
		PathPattern:        "/deployment/api/deployments/{depId}/requests",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeploymentRequestsUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDeploymentRequestsUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDeploymentRequestsUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetEventLogsUsingGET fetches event logs

  Returns the logs for an event.
*/
func (a *Client) GetEventLogsUsingGET(params *GetEventLogsUsingGETParams) (*GetEventLogsUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEventLogsUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getEventLogsUsingGET",
		Method:             "GET",
		PathPattern:        "/deployment/api/deployments/{depId}/requests/{requestId}/events/{eventId}/logs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEventLogsUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetEventLogsUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getEventLogsUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetRequestEventsUsingGET fetches request events

  Returns all the events for a request.
*/
func (a *Client) GetRequestEventsUsingGET(params *GetRequestEventsUsingGETParams) (*GetRequestEventsUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRequestEventsUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRequestEventsUsingGET",
		Method:             "GET",
		PathPattern:        "/deployment/api/deployments/{depId}/requests/{requestId}/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRequestEventsUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRequestEventsUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getRequestEventsUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetRequestUsingGET gets the request

  Returns the request with the given ID.
*/
func (a *Client) GetRequestUsingGET(params *GetRequestUsingGETParams) (*GetRequestUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRequestUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRequestUsingGET",
		Method:             "GET",
		PathPattern:        "/deployment/api/deployments/{depId}/requests/{requestId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRequestUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRequestUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getRequestUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}