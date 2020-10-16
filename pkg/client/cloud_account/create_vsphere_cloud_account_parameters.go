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

// NewCreateVSphereCloudAccountParams creates a new CreateVSphereCloudAccountParams object
// with the default values initialized.
func NewCreateVSphereCloudAccountParams() *CreateVSphereCloudAccountParams {
	var ()
	return &CreateVSphereCloudAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateVSphereCloudAccountParamsWithTimeout creates a new CreateVSphereCloudAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateVSphereCloudAccountParamsWithTimeout(timeout time.Duration) *CreateVSphereCloudAccountParams {
	var ()
	return &CreateVSphereCloudAccountParams{

		timeout: timeout,
	}
}

// NewCreateVSphereCloudAccountParamsWithContext creates a new CreateVSphereCloudAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateVSphereCloudAccountParamsWithContext(ctx context.Context) *CreateVSphereCloudAccountParams {
	var ()
	return &CreateVSphereCloudAccountParams{

		Context: ctx,
	}
}

// NewCreateVSphereCloudAccountParamsWithHTTPClient creates a new CreateVSphereCloudAccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateVSphereCloudAccountParamsWithHTTPClient(client *http.Client) *CreateVSphereCloudAccountParams {
	var ()
	return &CreateVSphereCloudAccountParams{
		HTTPClient: client,
	}
}

/*CreateVSphereCloudAccountParams contains all the parameters to send to the API endpoint
for the create v sphere cloud account operation typically these are written to a http.Request
*/
type CreateVSphereCloudAccountParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about

	*/
	APIVersion *string
	/*Body
	  CloudAccountVsphere specification

	*/
	Body *models.CloudAccountVsphereSpecification

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create v sphere cloud account params
func (o *CreateVSphereCloudAccountParams) WithTimeout(timeout time.Duration) *CreateVSphereCloudAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create v sphere cloud account params
func (o *CreateVSphereCloudAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create v sphere cloud account params
func (o *CreateVSphereCloudAccountParams) WithContext(ctx context.Context) *CreateVSphereCloudAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create v sphere cloud account params
func (o *CreateVSphereCloudAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create v sphere cloud account params
func (o *CreateVSphereCloudAccountParams) WithHTTPClient(client *http.Client) *CreateVSphereCloudAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create v sphere cloud account params
func (o *CreateVSphereCloudAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create v sphere cloud account params
func (o *CreateVSphereCloudAccountParams) WithAPIVersion(aPIVersion *string) *CreateVSphereCloudAccountParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create v sphere cloud account params
func (o *CreateVSphereCloudAccountParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the create v sphere cloud account params
func (o *CreateVSphereCloudAccountParams) WithBody(body *models.CloudAccountVsphereSpecification) *CreateVSphereCloudAccountParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create v sphere cloud account params
func (o *CreateVSphereCloudAccountParams) SetBody(body *models.CloudAccountVsphereSpecification) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateVSphereCloudAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
