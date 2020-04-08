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

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetNetworkDomainParams creates a new GetNetworkDomainParams object
// with the default values initialized.
func NewGetNetworkDomainParams() *GetNetworkDomainParams {
	var ()
	return &GetNetworkDomainParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkDomainParamsWithTimeout creates a new GetNetworkDomainParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworkDomainParamsWithTimeout(timeout time.Duration) *GetNetworkDomainParams {
	var ()
	return &GetNetworkDomainParams{

		timeout: timeout,
	}
}

// NewGetNetworkDomainParamsWithContext creates a new GetNetworkDomainParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworkDomainParamsWithContext(ctx context.Context) *GetNetworkDomainParams {
	var ()
	return &GetNetworkDomainParams{

		Context: ctx,
	}
}

// NewGetNetworkDomainParamsWithHTTPClient creates a new GetNetworkDomainParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworkDomainParamsWithHTTPClient(client *http.Client) *GetNetworkDomainParams {
	var ()
	return &GetNetworkDomainParams{
		HTTPClient: client,
	}
}

/*GetNetworkDomainParams contains all the parameters to send to the API endpoint
for the get network domain operation typically these are written to a http.Request
*/
type GetNetworkDomainParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about

	*/
	APIVersion *string
	/*ID
	  The ID of the network domain.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get network domain params
func (o *GetNetworkDomainParams) WithTimeout(timeout time.Duration) *GetNetworkDomainParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network domain params
func (o *GetNetworkDomainParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network domain params
func (o *GetNetworkDomainParams) WithContext(ctx context.Context) *GetNetworkDomainParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network domain params
func (o *GetNetworkDomainParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network domain params
func (o *GetNetworkDomainParams) WithHTTPClient(client *http.Client) *GetNetworkDomainParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network domain params
func (o *GetNetworkDomainParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get network domain params
func (o *GetNetworkDomainParams) WithAPIVersion(aPIVersion *string) *GetNetworkDomainParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get network domain params
func (o *GetNetworkDomainParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get network domain params
func (o *GetNetworkDomainParams) WithID(id string) *GetNetworkDomainParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get network domain params
func (o *GetNetworkDomainParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkDomainParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
