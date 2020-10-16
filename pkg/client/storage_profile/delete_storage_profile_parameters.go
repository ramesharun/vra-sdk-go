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

// NewDeleteStorageProfileParams creates a new DeleteStorageProfileParams object
// with the default values initialized.
func NewDeleteStorageProfileParams() *DeleteStorageProfileParams {
	var ()
	return &DeleteStorageProfileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteStorageProfileParamsWithTimeout creates a new DeleteStorageProfileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteStorageProfileParamsWithTimeout(timeout time.Duration) *DeleteStorageProfileParams {
	var ()
	return &DeleteStorageProfileParams{

		timeout: timeout,
	}
}

// NewDeleteStorageProfileParamsWithContext creates a new DeleteStorageProfileParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteStorageProfileParamsWithContext(ctx context.Context) *DeleteStorageProfileParams {
	var ()
	return &DeleteStorageProfileParams{

		Context: ctx,
	}
}

// NewDeleteStorageProfileParamsWithHTTPClient creates a new DeleteStorageProfileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteStorageProfileParamsWithHTTPClient(client *http.Client) *DeleteStorageProfileParams {
	var ()
	return &DeleteStorageProfileParams{
		HTTPClient: client,
	}
}

/*DeleteStorageProfileParams contains all the parameters to send to the API endpoint
for the delete storage profile operation typically these are written to a http.Request
*/
type DeleteStorageProfileParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about

	*/
	APIVersion *string
	/*ID
	  The ID of the storage profile.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete storage profile params
func (o *DeleteStorageProfileParams) WithTimeout(timeout time.Duration) *DeleteStorageProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete storage profile params
func (o *DeleteStorageProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete storage profile params
func (o *DeleteStorageProfileParams) WithContext(ctx context.Context) *DeleteStorageProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete storage profile params
func (o *DeleteStorageProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete storage profile params
func (o *DeleteStorageProfileParams) WithHTTPClient(client *http.Client) *DeleteStorageProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete storage profile params
func (o *DeleteStorageProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the delete storage profile params
func (o *DeleteStorageProfileParams) WithAPIVersion(aPIVersion *string) *DeleteStorageProfileParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the delete storage profile params
func (o *DeleteStorageProfileParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the delete storage profile params
func (o *DeleteStorageProfileParams) WithID(id string) *DeleteStorageProfileParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete storage profile params
func (o *DeleteStorageProfileParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteStorageProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
