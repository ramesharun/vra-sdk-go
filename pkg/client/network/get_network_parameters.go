// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetNetworkParams creates a new GetNetworkParams object
// with the default values initialized.
func NewGetNetworkParams() *GetNetworkParams {
	var ()
	return &GetNetworkParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkParamsWithTimeout creates a new GetNetworkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworkParamsWithTimeout(timeout time.Duration) *GetNetworkParams {
	var ()
	return &GetNetworkParams{

		timeout: timeout,
	}
}

// NewGetNetworkParamsWithContext creates a new GetNetworkParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworkParamsWithContext(ctx context.Context) *GetNetworkParams {
	var ()
	return &GetNetworkParams{

		Context: ctx,
	}
}

// NewGetNetworkParamsWithHTTPClient creates a new GetNetworkParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworkParamsWithHTTPClient(client *http.Client) *GetNetworkParams {
	var ()
	return &GetNetworkParams{
		HTTPClient: client,
	}
}

/*GetNetworkParams contains all the parameters to send to the API endpoint
for the get network operation typically these are written to a http.Request
*/
type GetNetworkParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about

	*/
	APIVersion *string
	/*ID
	  The ID of the network.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get network params
func (o *GetNetworkParams) WithTimeout(timeout time.Duration) *GetNetworkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network params
func (o *GetNetworkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network params
func (o *GetNetworkParams) WithContext(ctx context.Context) *GetNetworkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network params
func (o *GetNetworkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network params
func (o *GetNetworkParams) WithHTTPClient(client *http.Client) *GetNetworkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network params
func (o *GetNetworkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get network params
func (o *GetNetworkParams) WithAPIVersion(aPIVersion *string) *GetNetworkParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get network params
func (o *GetNetworkParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get network params
func (o *GetNetworkParams) WithID(id string) *GetNetworkParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get network params
func (o *GetNetworkParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.APIVersion != nil {

		// query param apiVersion
		var qrAPIVersion string
		if o.APIVersion != nil {
			qrAPIVersion = *o.APIVersion
		}
		qAPIVersion := qrAPIVersion
		if qAPIVersion != "" {
			if err := r.SetQueryParam("apiVersion", qAPIVersion); err != nil {
				return err
			}
		}

	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
