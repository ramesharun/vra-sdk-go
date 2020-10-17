// Code generated by go-swagger; DO NOT EDIT.

package cloud_account

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

// NewDeleteVSphereCloudAccountParams creates a new DeleteVSphereCloudAccountParams object
// with the default values initialized.
func NewDeleteVSphereCloudAccountParams() *DeleteVSphereCloudAccountParams {
	var ()
	return &DeleteVSphereCloudAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteVSphereCloudAccountParamsWithTimeout creates a new DeleteVSphereCloudAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteVSphereCloudAccountParamsWithTimeout(timeout time.Duration) *DeleteVSphereCloudAccountParams {
	var ()
	return &DeleteVSphereCloudAccountParams{

		timeout: timeout,
	}
}

// NewDeleteVSphereCloudAccountParamsWithContext creates a new DeleteVSphereCloudAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteVSphereCloudAccountParamsWithContext(ctx context.Context) *DeleteVSphereCloudAccountParams {
	var ()
	return &DeleteVSphereCloudAccountParams{

		Context: ctx,
	}
}

// NewDeleteVSphereCloudAccountParamsWithHTTPClient creates a new DeleteVSphereCloudAccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteVSphereCloudAccountParamsWithHTTPClient(client *http.Client) *DeleteVSphereCloudAccountParams {
	var ()
	return &DeleteVSphereCloudAccountParams{
		HTTPClient: client,
	}
}

/*DeleteVSphereCloudAccountParams contains all the parameters to send to the API endpoint
for the delete v sphere cloud account operation typically these are written to a http.Request
*/
type DeleteVSphereCloudAccountParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about

	*/
	APIVersion *string
	/*ID
	  The ID of the Cloud Account

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete v sphere cloud account params
func (o *DeleteVSphereCloudAccountParams) WithTimeout(timeout time.Duration) *DeleteVSphereCloudAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v sphere cloud account params
func (o *DeleteVSphereCloudAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v sphere cloud account params
func (o *DeleteVSphereCloudAccountParams) WithContext(ctx context.Context) *DeleteVSphereCloudAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v sphere cloud account params
func (o *DeleteVSphereCloudAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v sphere cloud account params
func (o *DeleteVSphereCloudAccountParams) WithHTTPClient(client *http.Client) *DeleteVSphereCloudAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v sphere cloud account params
func (o *DeleteVSphereCloudAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the delete v sphere cloud account params
func (o *DeleteVSphereCloudAccountParams) WithAPIVersion(aPIVersion *string) *DeleteVSphereCloudAccountParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the delete v sphere cloud account params
func (o *DeleteVSphereCloudAccountParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the delete v sphere cloud account params
func (o *DeleteVSphereCloudAccountParams) WithID(id string) *DeleteVSphereCloudAccountParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete v sphere cloud account params
func (o *DeleteVSphereCloudAccountParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteVSphereCloudAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
