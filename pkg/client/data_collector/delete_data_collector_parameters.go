// Code generated by go-swagger; DO NOT EDIT.

package data_collector

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

// NewDeleteDataCollectorParams creates a new DeleteDataCollectorParams object
// with the default values initialized.
func NewDeleteDataCollectorParams() *DeleteDataCollectorParams {
	var ()
	return &DeleteDataCollectorParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDataCollectorParamsWithTimeout creates a new DeleteDataCollectorParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteDataCollectorParamsWithTimeout(timeout time.Duration) *DeleteDataCollectorParams {
	var ()
	return &DeleteDataCollectorParams{

		timeout: timeout,
	}
}

// NewDeleteDataCollectorParamsWithContext creates a new DeleteDataCollectorParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteDataCollectorParamsWithContext(ctx context.Context) *DeleteDataCollectorParams {
	var ()
	return &DeleteDataCollectorParams{

		Context: ctx,
	}
}

// NewDeleteDataCollectorParamsWithHTTPClient creates a new DeleteDataCollectorParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteDataCollectorParamsWithHTTPClient(client *http.Client) *DeleteDataCollectorParams {
	var ()
	return &DeleteDataCollectorParams{
		HTTPClient: client,
	}
}

/*DeleteDataCollectorParams contains all the parameters to send to the API endpoint
for the delete data collector operation typically these are written to a http.Request
*/
type DeleteDataCollectorParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about

	*/
	APIVersion *string
	/*ID
	  The ID of the Data Collector.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete data collector params
func (o *DeleteDataCollectorParams) WithTimeout(timeout time.Duration) *DeleteDataCollectorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete data collector params
func (o *DeleteDataCollectorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete data collector params
func (o *DeleteDataCollectorParams) WithContext(ctx context.Context) *DeleteDataCollectorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete data collector params
func (o *DeleteDataCollectorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete data collector params
func (o *DeleteDataCollectorParams) WithHTTPClient(client *http.Client) *DeleteDataCollectorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete data collector params
func (o *DeleteDataCollectorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the delete data collector params
func (o *DeleteDataCollectorParams) WithAPIVersion(aPIVersion *string) *DeleteDataCollectorParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the delete data collector params
func (o *DeleteDataCollectorParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the delete data collector params
func (o *DeleteDataCollectorParams) WithID(id string) *DeleteDataCollectorParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete data collector params
func (o *DeleteDataCollectorParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDataCollectorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
