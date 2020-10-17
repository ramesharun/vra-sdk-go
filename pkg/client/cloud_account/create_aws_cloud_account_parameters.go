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

// NewCreateAwsCloudAccountParams creates a new CreateAwsCloudAccountParams object
// with the default values initialized.
func NewCreateAwsCloudAccountParams() *CreateAwsCloudAccountParams {
	var ()
	return &CreateAwsCloudAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAwsCloudAccountParamsWithTimeout creates a new CreateAwsCloudAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateAwsCloudAccountParamsWithTimeout(timeout time.Duration) *CreateAwsCloudAccountParams {
	var ()
	return &CreateAwsCloudAccountParams{

		timeout: timeout,
	}
}

// NewCreateAwsCloudAccountParamsWithContext creates a new CreateAwsCloudAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateAwsCloudAccountParamsWithContext(ctx context.Context) *CreateAwsCloudAccountParams {
	var ()
	return &CreateAwsCloudAccountParams{

		Context: ctx,
	}
}

// NewCreateAwsCloudAccountParamsWithHTTPClient creates a new CreateAwsCloudAccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateAwsCloudAccountParamsWithHTTPClient(client *http.Client) *CreateAwsCloudAccountParams {
	var ()
	return &CreateAwsCloudAccountParams{
		HTTPClient: client,
	}
}

/*CreateAwsCloudAccountParams contains all the parameters to send to the API endpoint
for the create aws cloud account operation typically these are written to a http.Request
*/
type CreateAwsCloudAccountParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about

	*/
	APIVersion *string
	/*Body
	  CloudAccountAws specification

	*/
	Body *models.CloudAccountAwsSpecification

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create aws cloud account params
func (o *CreateAwsCloudAccountParams) WithTimeout(timeout time.Duration) *CreateAwsCloudAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create aws cloud account params
func (o *CreateAwsCloudAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create aws cloud account params
func (o *CreateAwsCloudAccountParams) WithContext(ctx context.Context) *CreateAwsCloudAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create aws cloud account params
func (o *CreateAwsCloudAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create aws cloud account params
func (o *CreateAwsCloudAccountParams) WithHTTPClient(client *http.Client) *CreateAwsCloudAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create aws cloud account params
func (o *CreateAwsCloudAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create aws cloud account params
func (o *CreateAwsCloudAccountParams) WithAPIVersion(aPIVersion *string) *CreateAwsCloudAccountParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create aws cloud account params
func (o *CreateAwsCloudAccountParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the create aws cloud account params
func (o *CreateAwsCloudAccountParams) WithBody(body *models.CloudAccountAwsSpecification) *CreateAwsCloudAccountParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create aws cloud account params
func (o *CreateAwsCloudAccountParams) SetBody(body *models.CloudAccountAwsSpecification) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAwsCloudAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
