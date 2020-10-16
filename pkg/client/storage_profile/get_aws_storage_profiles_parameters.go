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

// NewGetAwsStorageProfilesParams creates a new GetAwsStorageProfilesParams object
// with the default values initialized.
func NewGetAwsStorageProfilesParams() *GetAwsStorageProfilesParams {
	var ()
	return &GetAwsStorageProfilesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAwsStorageProfilesParamsWithTimeout creates a new GetAwsStorageProfilesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAwsStorageProfilesParamsWithTimeout(timeout time.Duration) *GetAwsStorageProfilesParams {
	var ()
	return &GetAwsStorageProfilesParams{

		timeout: timeout,
	}
}

// NewGetAwsStorageProfilesParamsWithContext creates a new GetAwsStorageProfilesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAwsStorageProfilesParamsWithContext(ctx context.Context) *GetAwsStorageProfilesParams {
	var ()
	return &GetAwsStorageProfilesParams{

		Context: ctx,
	}
}

// NewGetAwsStorageProfilesParamsWithHTTPClient creates a new GetAwsStorageProfilesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAwsStorageProfilesParamsWithHTTPClient(client *http.Client) *GetAwsStorageProfilesParams {
	var ()
	return &GetAwsStorageProfilesParams{
		HTTPClient: client,
	}
}

/*GetAwsStorageProfilesParams contains all the parameters to send to the API endpoint
for the get aws storage profiles operation typically these are written to a http.Request
*/
type GetAwsStorageProfilesParams struct {

	/*DollarFilter
	  Add a filter to return limited results

	*/
	DollarFilter *string
	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about

	*/
	APIVersion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get aws storage profiles params
func (o *GetAwsStorageProfilesParams) WithTimeout(timeout time.Duration) *GetAwsStorageProfilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get aws storage profiles params
func (o *GetAwsStorageProfilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get aws storage profiles params
func (o *GetAwsStorageProfilesParams) WithContext(ctx context.Context) *GetAwsStorageProfilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get aws storage profiles params
func (o *GetAwsStorageProfilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get aws storage profiles params
func (o *GetAwsStorageProfilesParams) WithHTTPClient(client *http.Client) *GetAwsStorageProfilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get aws storage profiles params
func (o *GetAwsStorageProfilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the get aws storage profiles params
func (o *GetAwsStorageProfilesParams) WithDollarFilter(dollarFilter *string) *GetAwsStorageProfilesParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the get aws storage profiles params
func (o *GetAwsStorageProfilesParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithAPIVersion adds the aPIVersion to the get aws storage profiles params
func (o *GetAwsStorageProfilesParams) WithAPIVersion(aPIVersion *string) *GetAwsStorageProfilesParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get aws storage profiles params
func (o *GetAwsStorageProfilesParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *GetAwsStorageProfilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DollarFilter != nil {

		// query param $filter
		var qrDollarFilter string
		if o.DollarFilter != nil {
			qrDollarFilter = *o.DollarFilter
		}
		qDollarFilter := qrDollarFilter
		if qDollarFilter != "" {
			if err := r.SetQueryParam("$filter", qDollarFilter); err != nil {
				return err
			}
		}

	}

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
