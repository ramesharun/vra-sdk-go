// Code generated by go-swagger; DO NOT EDIT.

package image_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new image profile API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for image profile API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateImageProfile(params *CreateImageProfileParams) (*CreateImageProfileCreated, error)

	DeleteImageProfile(params *DeleteImageProfileParams) (*DeleteImageProfileNoContent, error)

	GetImageProfile(params *GetImageProfileParams) (*GetImageProfileOK, error)

	GetImageProfiles(params *GetImageProfilesParams) (*GetImageProfilesOK, error)

	UpdateImageProfile(params *UpdateImageProfileParams) (*UpdateImageProfileOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateImageProfile creates image profile

  Create image profile
*/
func (a *Client) CreateImageProfile(params *CreateImageProfileParams) (*CreateImageProfileCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateImageProfileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createImageProfile",
		Method:             "POST",
		PathPattern:        "/iaas/api/image-profiles",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateImageProfileReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateImageProfileCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createImageProfile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteImageProfile deletes image profile

  Delete image profile with a given id
*/
func (a *Client) DeleteImageProfile(params *DeleteImageProfileParams) (*DeleteImageProfileNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteImageProfileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteImageProfile",
		Method:             "DELETE",
		PathPattern:        "/iaas/api/image-profiles/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteImageProfileReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteImageProfileNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteImageProfile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetImageProfile gets image profile

  Get image profile with a given id
*/
func (a *Client) GetImageProfile(params *GetImageProfileParams) (*GetImageProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetImageProfileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getImageProfile",
		Method:             "GET",
		PathPattern:        "/iaas/api/image-profiles/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetImageProfileReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetImageProfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getImageProfile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetImageProfiles gets image profile

  Get all image profiles
*/
func (a *Client) GetImageProfiles(params *GetImageProfilesParams) (*GetImageProfilesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetImageProfilesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getImageProfiles",
		Method:             "GET",
		PathPattern:        "/iaas/api/image-profiles",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetImageProfilesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetImageProfilesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getImageProfiles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateImageProfile updates image profile

  Update image profile
*/
func (a *Client) UpdateImageProfile(params *UpdateImageProfileParams) (*UpdateImageProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateImageProfileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateImageProfile",
		Method:             "PATCH",
		PathPattern:        "/iaas/api/image-profiles/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateImageProfileReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateImageProfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateImageProfile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
