// Code generated by go-swagger; DO NOT EDIT.

package storage_profile

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

// NewGetVSphereStorageProfileParams creates a new GetVSphereStorageProfileParams object
// with the default values initialized.
func NewGetVSphereStorageProfileParams() *GetVSphereStorageProfileParams {
	var ()
	return &GetVSphereStorageProfileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVSphereStorageProfileParamsWithTimeout creates a new GetVSphereStorageProfileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVSphereStorageProfileParamsWithTimeout(timeout time.Duration) *GetVSphereStorageProfileParams {
	var ()
	return &GetVSphereStorageProfileParams{

		timeout: timeout,
	}
}

// NewGetVSphereStorageProfileParamsWithContext creates a new GetVSphereStorageProfileParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVSphereStorageProfileParamsWithContext(ctx context.Context) *GetVSphereStorageProfileParams {
	var ()
	return &GetVSphereStorageProfileParams{

		Context: ctx,
	}
}

// NewGetVSphereStorageProfileParamsWithHTTPClient creates a new GetVSphereStorageProfileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVSphereStorageProfileParamsWithHTTPClient(client *http.Client) *GetVSphereStorageProfileParams {
	var ()
	return &GetVSphereStorageProfileParams{
		HTTPClient: client,
	}
}

/*GetVSphereStorageProfileParams contains all the parameters to send to the API endpoint
for the get v sphere storage profile operation typically these are written to a http.Request
*/
type GetVSphereStorageProfileParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about

	*/
	APIVersion *string
	/*ID
	  The ID of storage profile.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get v sphere storage profile params
func (o *GetVSphereStorageProfileParams) WithTimeout(timeout time.Duration) *GetVSphereStorageProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v sphere storage profile params
func (o *GetVSphereStorageProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v sphere storage profile params
func (o *GetVSphereStorageProfileParams) WithContext(ctx context.Context) *GetVSphereStorageProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v sphere storage profile params
func (o *GetVSphereStorageProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v sphere storage profile params
func (o *GetVSphereStorageProfileParams) WithHTTPClient(client *http.Client) *GetVSphereStorageProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v sphere storage profile params
func (o *GetVSphereStorageProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get v sphere storage profile params
func (o *GetVSphereStorageProfileParams) WithAPIVersion(aPIVersion *string) *GetVSphereStorageProfileParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get v sphere storage profile params
func (o *GetVSphereStorageProfileParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get v sphere storage profile params
func (o *GetVSphereStorageProfileParams) WithID(id string) *GetVSphereStorageProfileParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get v sphere storage profile params
func (o *GetVSphereStorageProfileParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetVSphereStorageProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
