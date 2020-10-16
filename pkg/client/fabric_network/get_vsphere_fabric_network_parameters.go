// Code generated by go-swagger; DO NOT EDIT.

package fabric_network

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

// NewGetVsphereFabricNetworkParams creates a new GetVsphereFabricNetworkParams object
// with the default values initialized.
func NewGetVsphereFabricNetworkParams() *GetVsphereFabricNetworkParams {
	var ()
	return &GetVsphereFabricNetworkParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVsphereFabricNetworkParamsWithTimeout creates a new GetVsphereFabricNetworkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVsphereFabricNetworkParamsWithTimeout(timeout time.Duration) *GetVsphereFabricNetworkParams {
	var ()
	return &GetVsphereFabricNetworkParams{

		timeout: timeout,
	}
}

// NewGetVsphereFabricNetworkParamsWithContext creates a new GetVsphereFabricNetworkParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVsphereFabricNetworkParamsWithContext(ctx context.Context) *GetVsphereFabricNetworkParams {
	var ()
	return &GetVsphereFabricNetworkParams{

		Context: ctx,
	}
}

// NewGetVsphereFabricNetworkParamsWithHTTPClient creates a new GetVsphereFabricNetworkParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVsphereFabricNetworkParamsWithHTTPClient(client *http.Client) *GetVsphereFabricNetworkParams {
	var ()
	return &GetVsphereFabricNetworkParams{
		HTTPClient: client,
	}
}

/*GetVsphereFabricNetworkParams contains all the parameters to send to the API endpoint
for the get vsphere fabric network operation typically these are written to a http.Request
*/
type GetVsphereFabricNetworkParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about

	*/
	APIVersion *string
	/*ID
	  The ID of the fabric network.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get vsphere fabric network params
func (o *GetVsphereFabricNetworkParams) WithTimeout(timeout time.Duration) *GetVsphereFabricNetworkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vsphere fabric network params
func (o *GetVsphereFabricNetworkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vsphere fabric network params
func (o *GetVsphereFabricNetworkParams) WithContext(ctx context.Context) *GetVsphereFabricNetworkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vsphere fabric network params
func (o *GetVsphereFabricNetworkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vsphere fabric network params
func (o *GetVsphereFabricNetworkParams) WithHTTPClient(client *http.Client) *GetVsphereFabricNetworkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vsphere fabric network params
func (o *GetVsphereFabricNetworkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get vsphere fabric network params
func (o *GetVsphereFabricNetworkParams) WithAPIVersion(aPIVersion *string) *GetVsphereFabricNetworkParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get vsphere fabric network params
func (o *GetVsphereFabricNetworkParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get vsphere fabric network params
func (o *GetVsphereFabricNetworkParams) WithID(id string) *GetVsphereFabricNetworkParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get vsphere fabric network params
func (o *GetVsphereFabricNetworkParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetVsphereFabricNetworkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
