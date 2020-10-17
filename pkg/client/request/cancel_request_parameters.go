// Code generated by go-swagger; DO NOT EDIT.

package request

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

// NewCancelRequestParams creates a new CancelRequestParams object
// with the default values initialized.
func NewCancelRequestParams() *CancelRequestParams {
	var ()
	return &CancelRequestParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCancelRequestParamsWithTimeout creates a new CancelRequestParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCancelRequestParamsWithTimeout(timeout time.Duration) *CancelRequestParams {
	var ()
	return &CancelRequestParams{

		timeout: timeout,
	}
}

// NewCancelRequestParamsWithContext creates a new CancelRequestParams object
// with the default values initialized, and the ability to set a context for a request
func NewCancelRequestParamsWithContext(ctx context.Context) *CancelRequestParams {
	var ()
	return &CancelRequestParams{

		Context: ctx,
	}
}

// NewCancelRequestParamsWithHTTPClient creates a new CancelRequestParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCancelRequestParamsWithHTTPClient(client *http.Client) *CancelRequestParams {
	var ()
	return &CancelRequestParams{
		HTTPClient: client,
	}
}

/*CancelRequestParams contains all the parameters to send to the API endpoint
for the cancel request operation typically these are written to a http.Request
*/
type CancelRequestParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about

	*/
	APIVersion *string
	/*ID
	  The ID of the request.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cancel request params
func (o *CancelRequestParams) WithTimeout(timeout time.Duration) *CancelRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cancel request params
func (o *CancelRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cancel request params
func (o *CancelRequestParams) WithContext(ctx context.Context) *CancelRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cancel request params
func (o *CancelRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cancel request params
func (o *CancelRequestParams) WithHTTPClient(client *http.Client) *CancelRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cancel request params
func (o *CancelRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the cancel request params
func (o *CancelRequestParams) WithAPIVersion(aPIVersion *string) *CancelRequestParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the cancel request params
func (o *CancelRequestParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the cancel request params
func (o *CancelRequestParams) WithID(id string) *CancelRequestParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the cancel request params
func (o *CancelRequestParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CancelRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
