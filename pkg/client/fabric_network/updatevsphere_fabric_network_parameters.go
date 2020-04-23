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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// NewUpdatevSphereFabricNetworkParams creates a new UpdatevSphereFabricNetworkParams object
// with the default values initialized.
func NewUpdatevSphereFabricNetworkParams() *UpdatevSphereFabricNetworkParams {
	var ()
	return &UpdatevSphereFabricNetworkParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatevSphereFabricNetworkParamsWithTimeout creates a new UpdatevSphereFabricNetworkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdatevSphereFabricNetworkParamsWithTimeout(timeout time.Duration) *UpdatevSphereFabricNetworkParams {
	var ()
	return &UpdatevSphereFabricNetworkParams{

		timeout: timeout,
	}
}

// NewUpdatevSphereFabricNetworkParamsWithContext creates a new UpdatevSphereFabricNetworkParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdatevSphereFabricNetworkParamsWithContext(ctx context.Context) *UpdatevSphereFabricNetworkParams {
	var ()
	return &UpdatevSphereFabricNetworkParams{

		Context: ctx,
	}
}

// NewUpdatevSphereFabricNetworkParamsWithHTTPClient creates a new UpdatevSphereFabricNetworkParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdatevSphereFabricNetworkParamsWithHTTPClient(client *http.Client) *UpdatevSphereFabricNetworkParams {
	var ()
	return &UpdatevSphereFabricNetworkParams{
		HTTPClient: client,
	}
}

/*UpdatevSphereFabricNetworkParams contains all the parameters to send to the API endpoint
for the updatev sphere fabric network operation typically these are written to a http.Request
*/
type UpdatevSphereFabricNetworkParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about

	*/
	APIVersion *string
	/*Body
	  Vsphere Fabric Network Specification

	*/
	Body *models.FabricNetworkVsphereSpecification
	/*ID
	  The ID of the vSphere Fabric Network.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the updatev sphere fabric network params
func (o *UpdatevSphereFabricNetworkParams) WithTimeout(timeout time.Duration) *UpdatevSphereFabricNetworkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the updatev sphere fabric network params
func (o *UpdatevSphereFabricNetworkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the updatev sphere fabric network params
func (o *UpdatevSphereFabricNetworkParams) WithContext(ctx context.Context) *UpdatevSphereFabricNetworkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the updatev sphere fabric network params
func (o *UpdatevSphereFabricNetworkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the updatev sphere fabric network params
func (o *UpdatevSphereFabricNetworkParams) WithHTTPClient(client *http.Client) *UpdatevSphereFabricNetworkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the updatev sphere fabric network params
func (o *UpdatevSphereFabricNetworkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the updatev sphere fabric network params
func (o *UpdatevSphereFabricNetworkParams) WithAPIVersion(aPIVersion *string) *UpdatevSphereFabricNetworkParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the updatev sphere fabric network params
func (o *UpdatevSphereFabricNetworkParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the updatev sphere fabric network params
func (o *UpdatevSphereFabricNetworkParams) WithBody(body *models.FabricNetworkVsphereSpecification) *UpdatevSphereFabricNetworkParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the updatev sphere fabric network params
func (o *UpdatevSphereFabricNetworkParams) SetBody(body *models.FabricNetworkVsphereSpecification) {
	o.Body = body
}

// WithID adds the id to the updatev sphere fabric network params
func (o *UpdatevSphereFabricNetworkParams) WithID(id string) *UpdatevSphereFabricNetworkParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the updatev sphere fabric network params
func (o *UpdatevSphereFabricNetworkParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatevSphereFabricNetworkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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