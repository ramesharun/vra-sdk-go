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
	"github.com/go-openapi/swag"
)

// NewGetDataCollectorsParams creates a new GetDataCollectorsParams object
// with the default values initialized.
func NewGetDataCollectorsParams() *GetDataCollectorsParams {
	var ()
	return &GetDataCollectorsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDataCollectorsParamsWithTimeout creates a new GetDataCollectorsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDataCollectorsParamsWithTimeout(timeout time.Duration) *GetDataCollectorsParams {
	var ()
	return &GetDataCollectorsParams{

		timeout: timeout,
	}
}

// NewGetDataCollectorsParamsWithContext creates a new GetDataCollectorsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDataCollectorsParamsWithContext(ctx context.Context) *GetDataCollectorsParams {
	var ()
	return &GetDataCollectorsParams{

		Context: ctx,
	}
}

// NewGetDataCollectorsParamsWithHTTPClient creates a new GetDataCollectorsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDataCollectorsParamsWithHTTPClient(client *http.Client) *GetDataCollectorsParams {
	var ()
	return &GetDataCollectorsParams{
		HTTPClient: client,
	}
}

/*GetDataCollectorsParams contains all the parameters to send to the API endpoint
for the get data collectors operation typically these are written to a http.Request
*/
type GetDataCollectorsParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about

	*/
	APIVersion *string
	/*Disabled
	  If query param is provided with value equals to true, only disabled data collectors will be retrieved.

	*/
	Disabled *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get data collectors params
func (o *GetDataCollectorsParams) WithTimeout(timeout time.Duration) *GetDataCollectorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get data collectors params
func (o *GetDataCollectorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get data collectors params
func (o *GetDataCollectorsParams) WithContext(ctx context.Context) *GetDataCollectorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get data collectors params
func (o *GetDataCollectorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get data collectors params
func (o *GetDataCollectorsParams) WithHTTPClient(client *http.Client) *GetDataCollectorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get data collectors params
func (o *GetDataCollectorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get data collectors params
func (o *GetDataCollectorsParams) WithAPIVersion(aPIVersion *string) *GetDataCollectorsParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get data collectors params
func (o *GetDataCollectorsParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithDisabled adds the disabled to the get data collectors params
func (o *GetDataCollectorsParams) WithDisabled(disabled *bool) *GetDataCollectorsParams {
	o.SetDisabled(disabled)
	return o
}

// SetDisabled adds the disabled to the get data collectors params
func (o *GetDataCollectorsParams) SetDisabled(disabled *bool) {
	o.Disabled = disabled
}

// WriteToRequest writes these params to a swagger request
func (o *GetDataCollectorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Disabled != nil {

		// query param disabled
		var qrDisabled bool
		if o.Disabled != nil {
			qrDisabled = *o.Disabled
		}
		qDisabled := swag.FormatBool(qrDisabled)
		if qDisabled != "" {
			if err := r.SetQueryParam("disabled", qDisabled); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
