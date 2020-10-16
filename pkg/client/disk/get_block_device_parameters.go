// Code generated by go-swagger; DO NOT EDIT.

package disk

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

// NewGetBlockDeviceParams creates a new GetBlockDeviceParams object
// with the default values initialized.
func NewGetBlockDeviceParams() *GetBlockDeviceParams {
	var ()
	return &GetBlockDeviceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBlockDeviceParamsWithTimeout creates a new GetBlockDeviceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBlockDeviceParamsWithTimeout(timeout time.Duration) *GetBlockDeviceParams {
	var ()
	return &GetBlockDeviceParams{

		timeout: timeout,
	}
}

// NewGetBlockDeviceParamsWithContext creates a new GetBlockDeviceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBlockDeviceParamsWithContext(ctx context.Context) *GetBlockDeviceParams {
	var ()
	return &GetBlockDeviceParams{

		Context: ctx,
	}
}

// NewGetBlockDeviceParamsWithHTTPClient creates a new GetBlockDeviceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBlockDeviceParamsWithHTTPClient(client *http.Client) *GetBlockDeviceParams {
	var ()
	return &GetBlockDeviceParams{
		HTTPClient: client,
	}
}

/*GetBlockDeviceParams contains all the parameters to send to the API endpoint
for the get block device operation typically these are written to a http.Request
*/
type GetBlockDeviceParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about

	*/
	APIVersion *string
	/*ID
	  The ID of the block device.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get block device params
func (o *GetBlockDeviceParams) WithTimeout(timeout time.Duration) *GetBlockDeviceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get block device params
func (o *GetBlockDeviceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get block device params
func (o *GetBlockDeviceParams) WithContext(ctx context.Context) *GetBlockDeviceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get block device params
func (o *GetBlockDeviceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get block device params
func (o *GetBlockDeviceParams) WithHTTPClient(client *http.Client) *GetBlockDeviceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get block device params
func (o *GetBlockDeviceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get block device params
func (o *GetBlockDeviceParams) WithAPIVersion(aPIVersion *string) *GetBlockDeviceParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get block device params
func (o *GetBlockDeviceParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get block device params
func (o *GetBlockDeviceParams) WithID(id string) *GetBlockDeviceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get block device params
func (o *GetBlockDeviceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetBlockDeviceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
