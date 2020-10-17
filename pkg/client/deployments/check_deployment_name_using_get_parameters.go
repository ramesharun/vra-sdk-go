// Code generated by go-swagger; DO NOT EDIT.

package deployments

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

// NewCheckDeploymentNameUsingGETParams creates a new CheckDeploymentNameUsingGETParams object
// with the default values initialized.
func NewCheckDeploymentNameUsingGETParams() *CheckDeploymentNameUsingGETParams {
	var ()
	return &CheckDeploymentNameUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCheckDeploymentNameUsingGETParamsWithTimeout creates a new CheckDeploymentNameUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCheckDeploymentNameUsingGETParamsWithTimeout(timeout time.Duration) *CheckDeploymentNameUsingGETParams {
	var ()
	return &CheckDeploymentNameUsingGETParams{

		timeout: timeout,
	}
}

// NewCheckDeploymentNameUsingGETParamsWithContext creates a new CheckDeploymentNameUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewCheckDeploymentNameUsingGETParamsWithContext(ctx context.Context) *CheckDeploymentNameUsingGETParams {
	var ()
	return &CheckDeploymentNameUsingGETParams{

		Context: ctx,
	}
}

// NewCheckDeploymentNameUsingGETParamsWithHTTPClient creates a new CheckDeploymentNameUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCheckDeploymentNameUsingGETParamsWithHTTPClient(client *http.Client) *CheckDeploymentNameUsingGETParams {
	var ()
	return &CheckDeploymentNameUsingGETParams{
		HTTPClient: client,
	}
}

/*CheckDeploymentNameUsingGETParams contains all the parameters to send to the API endpoint
for the check deployment name using g e t operation typically these are written to a http.Request
*/
type CheckDeploymentNameUsingGETParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /catalog/api/about

	*/
	APIVersion *string
	/*Name
	  Deployment name

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the check deployment name using get params
func (o *CheckDeploymentNameUsingGETParams) WithTimeout(timeout time.Duration) *CheckDeploymentNameUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the check deployment name using get params
func (o *CheckDeploymentNameUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the check deployment name using get params
func (o *CheckDeploymentNameUsingGETParams) WithContext(ctx context.Context) *CheckDeploymentNameUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the check deployment name using get params
func (o *CheckDeploymentNameUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the check deployment name using get params
func (o *CheckDeploymentNameUsingGETParams) WithHTTPClient(client *http.Client) *CheckDeploymentNameUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the check deployment name using get params
func (o *CheckDeploymentNameUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the check deployment name using get params
func (o *CheckDeploymentNameUsingGETParams) WithAPIVersion(aPIVersion *string) *CheckDeploymentNameUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the check deployment name using get params
func (o *CheckDeploymentNameUsingGETParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithName adds the name to the check deployment name using get params
func (o *CheckDeploymentNameUsingGETParams) WithName(name string) *CheckDeploymentNameUsingGETParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the check deployment name using get params
func (o *CheckDeploymentNameUsingGETParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *CheckDeploymentNameUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
