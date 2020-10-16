// Code generated by go-swagger; DO NOT EDIT.

package marketplace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new marketplace API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for marketplace API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	FilterItemsUsingGET(params *FilterItemsUsingGETParams) (*FilterItemsUsingGETOK, error)

	FiltersUsingGET(params *FiltersUsingGETParams) (*FiltersUsingGETOK, error)

	GetDetails(params *GetDetailsParams) (*GetDetailsOK, error)

	GetLink(params *GetLinkParams) (*GetLinkOK, error)

	GetReviews(params *GetReviewsParams) (*GetReviewsOK, error)

	Search(params *SearchParams) (*SearchOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  FilterItemsUsingGET lists filter entries for a filter
*/
func (a *Client) FilterItemsUsingGET(params *FilterItemsUsingGETParams) (*FilterItemsUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFilterItemsUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "filterItemsUsingGET",
		Method:             "GET",
		PathPattern:        "/content/api/marketplace/sources/{sourceId}/filters/{filterId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FilterItemsUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FilterItemsUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for filterItemsUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  FiltersUsingGET lists filters by endpoint
*/
func (a *Client) FiltersUsingGET(params *FiltersUsingGETParams) (*FiltersUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFiltersUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "filtersUsingGET",
		Method:             "GET",
		PathPattern:        "/content/api/marketplace/sources/{sourceId}/filters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FiltersUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FiltersUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for filtersUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDetails gets content details
*/
func (a *Client) GetDetails(params *GetDetailsParams) (*GetDetailsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDetailsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDetails",
		Method:             "GET",
		PathPattern:        "/content/api/marketplace/sources/{sourceId}/contents/{contentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDetailsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDetailsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDetails: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetLink gets link to the content
*/
func (a *Client) GetLink(params *GetLinkParams) (*GetLinkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLinkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLink",
		Method:             "GET",
		PathPattern:        "/content/api/marketplace/sources/{sourceId}/contents/{contentId}/link",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLinkReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLinkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getLink: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetReviews gets reviews for the content
*/
func (a *Client) GetReviews(params *GetReviewsParams) (*GetReviewsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetReviewsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getReviews",
		Method:             "GET",
		PathPattern:        "/content/api/marketplace/sources/{sourceId}/contents/{contentId}/reviews",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetReviewsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetReviewsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getReviews: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Search searches for content based on the search term
*/
func (a *Client) Search(params *SearchParams) (*SearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "search",
		Method:             "GET",
		PathPattern:        "/content/api/marketplace/sources/{sourceId}/contents",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for search: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
