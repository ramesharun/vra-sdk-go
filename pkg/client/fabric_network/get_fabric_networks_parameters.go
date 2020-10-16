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
	"github.com/go-openapi/strfmt"
)

// NewGetFabricNetworksParams creates a new GetFabricNetworksParams object
// with the default values initialized.
func NewGetFabricNetworksParams() *GetFabricNetworksParams {
	var ()
	return &GetFabricNetworksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFabricNetworksParamsWithTimeout creates a new GetFabricNetworksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFabricNetworksParamsWithTimeout(timeout time.Duration) *GetFabricNetworksParams {
	var ()
	return &GetFabricNetworksParams{

		timeout: timeout,
	}
}

// NewGetFabricNetworksParamsWithContext creates a new GetFabricNetworksParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFabricNetworksParamsWithContext(ctx context.Context) *GetFabricNetworksParams {
	var ()
	return &GetFabricNetworksParams{

		Context: ctx,
	}
}

// NewGetFabricNetworksParamsWithHTTPClient creates a new GetFabricNetworksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFabricNetworksParamsWithHTTPClient(client *http.Client) *GetFabricNetworksParams {
	var ()
	return &GetFabricNetworksParams{
		HTTPClient: client,
	}
}

/*GetFabricNetworksParams contains all the parameters to send to the API endpoint
for the get fabric networks operation typically these are written to a http.Request
*/
type GetFabricNetworksParams struct {

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

// WithTimeout adds the timeout to the get fabric networks params
func (o *GetFabricNetworksParams) WithTimeout(timeout time.Duration) *GetFabricNetworksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get fabric networks params
func (o *GetFabricNetworksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get fabric networks params
func (o *GetFabricNetworksParams) WithContext(ctx context.Context) *GetFabricNetworksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get fabric networks params
func (o *GetFabricNetworksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get fabric networks params
func (o *GetFabricNetworksParams) WithHTTPClient(client *http.Client) *GetFabricNetworksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get fabric networks params
func (o *GetFabricNetworksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the get fabric networks params
func (o *GetFabricNetworksParams) WithDollarFilter(dollarFilter *string) *GetFabricNetworksParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the get fabric networks params
func (o *GetFabricNetworksParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithAPIVersion adds the aPIVersion to the get fabric networks params
func (o *GetFabricNetworksParams) WithAPIVersion(aPIVersion *string) *GetFabricNetworksParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get fabric networks params
func (o *GetFabricNetworksParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *GetFabricNetworksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
