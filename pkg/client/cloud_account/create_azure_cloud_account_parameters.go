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

// NewCreateAzureCloudAccountParams creates a new CreateAzureCloudAccountParams object
// with the default values initialized.
func NewCreateAzureCloudAccountParams() *CreateAzureCloudAccountParams {
	var ()
	return &CreateAzureCloudAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAzureCloudAccountParamsWithTimeout creates a new CreateAzureCloudAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateAzureCloudAccountParamsWithTimeout(timeout time.Duration) *CreateAzureCloudAccountParams {
	var ()
	return &CreateAzureCloudAccountParams{

		timeout: timeout,
	}
}

// NewCreateAzureCloudAccountParamsWithContext creates a new CreateAzureCloudAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateAzureCloudAccountParamsWithContext(ctx context.Context) *CreateAzureCloudAccountParams {
	var ()
	return &CreateAzureCloudAccountParams{

		Context: ctx,
	}
}

// NewCreateAzureCloudAccountParamsWithHTTPClient creates a new CreateAzureCloudAccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateAzureCloudAccountParamsWithHTTPClient(client *http.Client) *CreateAzureCloudAccountParams {
	var ()
	return &CreateAzureCloudAccountParams{
		HTTPClient: client,
	}
}

/*CreateAzureCloudAccountParams contains all the parameters to send to the API endpoint
for the create azure cloud account operation typically these are written to a http.Request
*/
type CreateAzureCloudAccountParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about

	*/
	APIVersion *string
	/*Body
	  CloudAccountAzure specification

	*/
	Body *models.CloudAccountAzureSpecification

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create azure cloud account params
func (o *CreateAzureCloudAccountParams) WithTimeout(timeout time.Duration) *CreateAzureCloudAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create azure cloud account params
func (o *CreateAzureCloudAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create azure cloud account params
func (o *CreateAzureCloudAccountParams) WithContext(ctx context.Context) *CreateAzureCloudAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create azure cloud account params
func (o *CreateAzureCloudAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create azure cloud account params
func (o *CreateAzureCloudAccountParams) WithHTTPClient(client *http.Client) *CreateAzureCloudAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create azure cloud account params
func (o *CreateAzureCloudAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create azure cloud account params
func (o *CreateAzureCloudAccountParams) WithAPIVersion(aPIVersion *string) *CreateAzureCloudAccountParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create azure cloud account params
func (o *CreateAzureCloudAccountParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the create azure cloud account params
func (o *CreateAzureCloudAccountParams) WithBody(body *models.CloudAccountAzureSpecification) *CreateAzureCloudAccountParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create azure cloud account params
func (o *CreateAzureCloudAccountParams) SetBody(body *models.CloudAccountAzureSpecification) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAzureCloudAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
