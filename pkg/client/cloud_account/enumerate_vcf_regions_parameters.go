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

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// NewEnumerateVcfRegionsParams creates a new EnumerateVcfRegionsParams object
// with the default values initialized.
func NewEnumerateVcfRegionsParams() *EnumerateVcfRegionsParams {
	var ()
	return &EnumerateVcfRegionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEnumerateVcfRegionsParamsWithTimeout creates a new EnumerateVcfRegionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEnumerateVcfRegionsParamsWithTimeout(timeout time.Duration) *EnumerateVcfRegionsParams {
	var ()
	return &EnumerateVcfRegionsParams{

		timeout: timeout,
	}
}

// NewEnumerateVcfRegionsParamsWithContext creates a new EnumerateVcfRegionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewEnumerateVcfRegionsParamsWithContext(ctx context.Context) *EnumerateVcfRegionsParams {
	var ()
	return &EnumerateVcfRegionsParams{

		Context: ctx,
	}
}

// NewEnumerateVcfRegionsParamsWithHTTPClient creates a new EnumerateVcfRegionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEnumerateVcfRegionsParamsWithHTTPClient(client *http.Client) *EnumerateVcfRegionsParams {
	var ()
	return &EnumerateVcfRegionsParams{
		HTTPClient: client,
	}
}

/*EnumerateVcfRegionsParams contains all the parameters to send to the API endpoint
for the enumerate vcf regions operation typically these are written to a http.Request
*/
type EnumerateVcfRegionsParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about

	*/
	APIVersion *string
	/*Body
	  CloudAccount specification

	*/
	Body *models.CloudAccountVcfSpecification

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the enumerate vcf regions params
func (o *EnumerateVcfRegionsParams) WithTimeout(timeout time.Duration) *EnumerateVcfRegionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the enumerate vcf regions params
func (o *EnumerateVcfRegionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the enumerate vcf regions params
func (o *EnumerateVcfRegionsParams) WithContext(ctx context.Context) *EnumerateVcfRegionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the enumerate vcf regions params
func (o *EnumerateVcfRegionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the enumerate vcf regions params
func (o *EnumerateVcfRegionsParams) WithHTTPClient(client *http.Client) *EnumerateVcfRegionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the enumerate vcf regions params
func (o *EnumerateVcfRegionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the enumerate vcf regions params
func (o *EnumerateVcfRegionsParams) WithAPIVersion(aPIVersion *string) *EnumerateVcfRegionsParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the enumerate vcf regions params
func (o *EnumerateVcfRegionsParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the enumerate vcf regions params
func (o *EnumerateVcfRegionsParams) WithBody(body *models.CloudAccountVcfSpecification) *EnumerateVcfRegionsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the enumerate vcf regions params
func (o *EnumerateVcfRegionsParams) SetBody(body *models.CloudAccountVcfSpecification) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *EnumerateVcfRegionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
