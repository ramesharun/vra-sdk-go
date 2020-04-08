// Code generated by go-swagger; DO NOT EDIT.

package deployment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSingleDeploymentParams creates a new GetSingleDeploymentParams object
// with the default values initialized.
func NewGetSingleDeploymentParams() *GetSingleDeploymentParams {
	var ()
	return &GetSingleDeploymentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSingleDeploymentParamsWithTimeout creates a new GetSingleDeploymentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSingleDeploymentParamsWithTimeout(timeout time.Duration) *GetSingleDeploymentParams {
	var ()
	return &GetSingleDeploymentParams{

		timeout: timeout,
	}
}

// NewGetSingleDeploymentParamsWithContext creates a new GetSingleDeploymentParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSingleDeploymentParamsWithContext(ctx context.Context) *GetSingleDeploymentParams {
	var ()
	return &GetSingleDeploymentParams{

		Context: ctx,
	}
}

// NewGetSingleDeploymentParamsWithHTTPClient creates a new GetSingleDeploymentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSingleDeploymentParamsWithHTTPClient(client *http.Client) *GetSingleDeploymentParams {
	var ()
	return &GetSingleDeploymentParams{
		HTTPClient: client,
	}
}

/*GetSingleDeploymentParams contains all the parameters to send to the API endpoint
for the get single deployment operation typically these are written to a http.Request
*/
type GetSingleDeploymentParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about

	*/
	APIVersion *string
	/*ID
	  The id of the deployment.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get single deployment params
func (o *GetSingleDeploymentParams) WithTimeout(timeout time.Duration) *GetSingleDeploymentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get single deployment params
func (o *GetSingleDeploymentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get single deployment params
func (o *GetSingleDeploymentParams) WithContext(ctx context.Context) *GetSingleDeploymentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get single deployment params
func (o *GetSingleDeploymentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get single deployment params
func (o *GetSingleDeploymentParams) WithHTTPClient(client *http.Client) *GetSingleDeploymentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get single deployment params
func (o *GetSingleDeploymentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get single deployment params
func (o *GetSingleDeploymentParams) WithAPIVersion(aPIVersion *string) *GetSingleDeploymentParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get single deployment params
func (o *GetSingleDeploymentParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get single deployment params
func (o *GetSingleDeploymentParams) WithID(id string) *GetSingleDeploymentParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get single deployment params
func (o *GetSingleDeploymentParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetSingleDeploymentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
