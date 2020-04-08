// Code generated by go-swagger; DO NOT EDIT.

package disk

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

// NewAttachMachineDiskParams creates a new AttachMachineDiskParams object
// with the default values initialized.
func NewAttachMachineDiskParams() *AttachMachineDiskParams {
	var ()
	return &AttachMachineDiskParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAttachMachineDiskParamsWithTimeout creates a new AttachMachineDiskParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAttachMachineDiskParamsWithTimeout(timeout time.Duration) *AttachMachineDiskParams {
	var ()
	return &AttachMachineDiskParams{

		timeout: timeout,
	}
}

// NewAttachMachineDiskParamsWithContext creates a new AttachMachineDiskParams object
// with the default values initialized, and the ability to set a context for a request
func NewAttachMachineDiskParamsWithContext(ctx context.Context) *AttachMachineDiskParams {
	var ()
	return &AttachMachineDiskParams{

		Context: ctx,
	}
}

// NewAttachMachineDiskParamsWithHTTPClient creates a new AttachMachineDiskParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAttachMachineDiskParamsWithHTTPClient(client *http.Client) *AttachMachineDiskParams {
	var ()
	return &AttachMachineDiskParams{
		HTTPClient: client,
	}
}

/*AttachMachineDiskParams contains all the parameters to send to the API endpoint
for the attach machine disk operation typically these are written to a http.Request
*/
type AttachMachineDiskParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about

	*/
	APIVersion *string
	/*Body
	  Disk Specification instance

	*/
	Body *models.DiskAttachmentSpecification
	/*ID
	  The ID of the machine.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the attach machine disk params
func (o *AttachMachineDiskParams) WithTimeout(timeout time.Duration) *AttachMachineDiskParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the attach machine disk params
func (o *AttachMachineDiskParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the attach machine disk params
func (o *AttachMachineDiskParams) WithContext(ctx context.Context) *AttachMachineDiskParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the attach machine disk params
func (o *AttachMachineDiskParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the attach machine disk params
func (o *AttachMachineDiskParams) WithHTTPClient(client *http.Client) *AttachMachineDiskParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the attach machine disk params
func (o *AttachMachineDiskParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the attach machine disk params
func (o *AttachMachineDiskParams) WithAPIVersion(aPIVersion *string) *AttachMachineDiskParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the attach machine disk params
func (o *AttachMachineDiskParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the attach machine disk params
func (o *AttachMachineDiskParams) WithBody(body *models.DiskAttachmentSpecification) *AttachMachineDiskParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the attach machine disk params
func (o *AttachMachineDiskParams) SetBody(body *models.DiskAttachmentSpecification) {
	o.Body = body
}

// WithID adds the id to the attach machine disk params
func (o *AttachMachineDiskParams) WithID(id string) *AttachMachineDiskParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the attach machine disk params
func (o *AttachMachineDiskParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AttachMachineDiskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
