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

// NewGetNsxTCloudAccountsParams creates a new GetNsxTCloudAccountsParams object
// with the default values initialized.
func NewGetNsxTCloudAccountsParams() *GetNsxTCloudAccountsParams {
	var ()
	return &GetNsxTCloudAccountsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNsxTCloudAccountsParamsWithTimeout creates a new GetNsxTCloudAccountsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNsxTCloudAccountsParamsWithTimeout(timeout time.Duration) *GetNsxTCloudAccountsParams {
	var ()
	return &GetNsxTCloudAccountsParams{

		timeout: timeout,
	}
}

// NewGetNsxTCloudAccountsParamsWithContext creates a new GetNsxTCloudAccountsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNsxTCloudAccountsParamsWithContext(ctx context.Context) *GetNsxTCloudAccountsParams {
	var ()
	return &GetNsxTCloudAccountsParams{

		Context: ctx,
	}
}

// NewGetNsxTCloudAccountsParamsWithHTTPClient creates a new GetNsxTCloudAccountsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNsxTCloudAccountsParamsWithHTTPClient(client *http.Client) *GetNsxTCloudAccountsParams {
	var ()
	return &GetNsxTCloudAccountsParams{
		HTTPClient: client,
	}
}

/*GetNsxTCloudAccountsParams contains all the parameters to send to the API endpoint
for the get nsx t cloud accounts operation typically these are written to a http.Request
*/
type GetNsxTCloudAccountsParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about

	*/
	APIVersion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get nsx t cloud accounts params
func (o *GetNsxTCloudAccountsParams) WithTimeout(timeout time.Duration) *GetNsxTCloudAccountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get nsx t cloud accounts params
func (o *GetNsxTCloudAccountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get nsx t cloud accounts params
func (o *GetNsxTCloudAccountsParams) WithContext(ctx context.Context) *GetNsxTCloudAccountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get nsx t cloud accounts params
func (o *GetNsxTCloudAccountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get nsx t cloud accounts params
func (o *GetNsxTCloudAccountsParams) WithHTTPClient(client *http.Client) *GetNsxTCloudAccountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get nsx t cloud accounts params
func (o *GetNsxTCloudAccountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get nsx t cloud accounts params
func (o *GetNsxTCloudAccountsParams) WithAPIVersion(aPIVersion *string) *GetNsxTCloudAccountsParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get nsx t cloud accounts params
func (o *GetNsxTCloudAccountsParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *GetNsxTCloudAccountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
