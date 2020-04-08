// Code generated by go-swagger; DO NOT EDIT.

package project

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

// NewUpdateProjectResourceMetadataParams creates a new UpdateProjectResourceMetadataParams object
// with the default values initialized.
func NewUpdateProjectResourceMetadataParams() *UpdateProjectResourceMetadataParams {
	var ()
	return &UpdateProjectResourceMetadataParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateProjectResourceMetadataParamsWithTimeout creates a new UpdateProjectResourceMetadataParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateProjectResourceMetadataParamsWithTimeout(timeout time.Duration) *UpdateProjectResourceMetadataParams {
	var ()
	return &UpdateProjectResourceMetadataParams{

		timeout: timeout,
	}
}

// NewUpdateProjectResourceMetadataParamsWithContext creates a new UpdateProjectResourceMetadataParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateProjectResourceMetadataParamsWithContext(ctx context.Context) *UpdateProjectResourceMetadataParams {
	var ()
	return &UpdateProjectResourceMetadataParams{

		Context: ctx,
	}
}

// NewUpdateProjectResourceMetadataParamsWithHTTPClient creates a new UpdateProjectResourceMetadataParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateProjectResourceMetadataParamsWithHTTPClient(client *http.Client) *UpdateProjectResourceMetadataParams {
	var ()
	return &UpdateProjectResourceMetadataParams{
		HTTPClient: client,
	}
}

/*UpdateProjectResourceMetadataParams contains all the parameters to send to the API endpoint
for the update project resource metadata operation typically these are written to a http.Request
*/
type UpdateProjectResourceMetadataParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about

	*/
	APIVersion *string
	/*Body
	  Project specification

	*/
	Body *models.ProjectResourceMetadataSpecification
	/*ID
	  The ID of the project.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update project resource metadata params
func (o *UpdateProjectResourceMetadataParams) WithTimeout(timeout time.Duration) *UpdateProjectResourceMetadataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update project resource metadata params
func (o *UpdateProjectResourceMetadataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update project resource metadata params
func (o *UpdateProjectResourceMetadataParams) WithContext(ctx context.Context) *UpdateProjectResourceMetadataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update project resource metadata params
func (o *UpdateProjectResourceMetadataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update project resource metadata params
func (o *UpdateProjectResourceMetadataParams) WithHTTPClient(client *http.Client) *UpdateProjectResourceMetadataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update project resource metadata params
func (o *UpdateProjectResourceMetadataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the update project resource metadata params
func (o *UpdateProjectResourceMetadataParams) WithAPIVersion(aPIVersion *string) *UpdateProjectResourceMetadataParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the update project resource metadata params
func (o *UpdateProjectResourceMetadataParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the update project resource metadata params
func (o *UpdateProjectResourceMetadataParams) WithBody(body *models.ProjectResourceMetadataSpecification) *UpdateProjectResourceMetadataParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update project resource metadata params
func (o *UpdateProjectResourceMetadataParams) SetBody(body *models.ProjectResourceMetadataSpecification) {
	o.Body = body
}

// WithID adds the id to the update project resource metadata params
func (o *UpdateProjectResourceMetadataParams) WithID(id string) *UpdateProjectResourceMetadataParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update project resource metadata params
func (o *UpdateProjectResourceMetadataParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateProjectResourceMetadataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
