// Code generated by go-swagger; DO NOT EDIT.

package image_profile

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

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// NewCreateImageProfileParams creates a new CreateImageProfileParams object
// with the default values initialized.
func NewCreateImageProfileParams() *CreateImageProfileParams {
	var ()
	return &CreateImageProfileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateImageProfileParamsWithTimeout creates a new CreateImageProfileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateImageProfileParamsWithTimeout(timeout time.Duration) *CreateImageProfileParams {
	var ()
	return &CreateImageProfileParams{

		timeout: timeout,
	}
}

// NewCreateImageProfileParamsWithContext creates a new CreateImageProfileParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateImageProfileParamsWithContext(ctx context.Context) *CreateImageProfileParams {
	var ()
	return &CreateImageProfileParams{

		Context: ctx,
	}
}

// NewCreateImageProfileParamsWithHTTPClient creates a new CreateImageProfileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateImageProfileParamsWithHTTPClient(client *http.Client) *CreateImageProfileParams {
	var ()
	return &CreateImageProfileParams{
		HTTPClient: client,
	}
}

/*CreateImageProfileParams contains all the parameters to send to the API endpoint
for the create image profile operation typically these are written to a http.Request
*/
type CreateImageProfileParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about

	*/
	APIVersion *string
	/*Body
	  ImageProfile instance

	*/
	Body *models.ImageProfileSpecification

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create image profile params
func (o *CreateImageProfileParams) WithTimeout(timeout time.Duration) *CreateImageProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create image profile params
func (o *CreateImageProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create image profile params
func (o *CreateImageProfileParams) WithContext(ctx context.Context) *CreateImageProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create image profile params
func (o *CreateImageProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create image profile params
func (o *CreateImageProfileParams) WithHTTPClient(client *http.Client) *CreateImageProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create image profile params
func (o *CreateImageProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create image profile params
func (o *CreateImageProfileParams) WithAPIVersion(aPIVersion *string) *CreateImageProfileParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create image profile params
func (o *CreateImageProfileParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the create image profile params
func (o *CreateImageProfileParams) WithBody(body *models.ImageProfileSpecification) *CreateImageProfileParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create image profile params
func (o *CreateImageProfileParams) SetBody(body *models.ImageProfileSpecification) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateImageProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
