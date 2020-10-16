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

// NewGetBlockDevicesParams creates a new GetBlockDevicesParams object
// with the default values initialized.
func NewGetBlockDevicesParams() *GetBlockDevicesParams {
	var ()
	return &GetBlockDevicesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBlockDevicesParamsWithTimeout creates a new GetBlockDevicesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBlockDevicesParamsWithTimeout(timeout time.Duration) *GetBlockDevicesParams {
	var ()
	return &GetBlockDevicesParams{

		timeout: timeout,
	}
}

// NewGetBlockDevicesParamsWithContext creates a new GetBlockDevicesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBlockDevicesParamsWithContext(ctx context.Context) *GetBlockDevicesParams {
	var ()
	return &GetBlockDevicesParams{

		Context: ctx,
	}
}

// NewGetBlockDevicesParamsWithHTTPClient creates a new GetBlockDevicesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBlockDevicesParamsWithHTTPClient(client *http.Client) *GetBlockDevicesParams {
	var ()
	return &GetBlockDevicesParams{
		HTTPClient: client,
	}
}

/*GetBlockDevicesParams contains all the parameters to send to the API endpoint
for the get block devices operation typically these are written to a http.Request
*/
type GetBlockDevicesParams struct {

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

// WithTimeout adds the timeout to the get block devices params
func (o *GetBlockDevicesParams) WithTimeout(timeout time.Duration) *GetBlockDevicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get block devices params
func (o *GetBlockDevicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get block devices params
func (o *GetBlockDevicesParams) WithContext(ctx context.Context) *GetBlockDevicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get block devices params
func (o *GetBlockDevicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get block devices params
func (o *GetBlockDevicesParams) WithHTTPClient(client *http.Client) *GetBlockDevicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get block devices params
func (o *GetBlockDevicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the get block devices params
func (o *GetBlockDevicesParams) WithDollarFilter(dollarFilter *string) *GetBlockDevicesParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the get block devices params
func (o *GetBlockDevicesParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithAPIVersion adds the aPIVersion to the get block devices params
func (o *GetBlockDevicesParams) WithAPIVersion(aPIVersion *string) *GetBlockDevicesParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get block devices params
func (o *GetBlockDevicesParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *GetBlockDevicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
