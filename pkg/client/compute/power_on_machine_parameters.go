// Code generated by go-swagger; DO NOT EDIT.

package compute

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

// NewPowerOnMachineParams creates a new PowerOnMachineParams object
// with the default values initialized.
func NewPowerOnMachineParams() *PowerOnMachineParams {
	var ()
	return &PowerOnMachineParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPowerOnMachineParamsWithTimeout creates a new PowerOnMachineParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPowerOnMachineParamsWithTimeout(timeout time.Duration) *PowerOnMachineParams {
	var ()
	return &PowerOnMachineParams{

		timeout: timeout,
	}
}

// NewPowerOnMachineParamsWithContext creates a new PowerOnMachineParams object
// with the default values initialized, and the ability to set a context for a request
func NewPowerOnMachineParamsWithContext(ctx context.Context) *PowerOnMachineParams {
	var ()
	return &PowerOnMachineParams{

		Context: ctx,
	}
}

// NewPowerOnMachineParamsWithHTTPClient creates a new PowerOnMachineParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPowerOnMachineParamsWithHTTPClient(client *http.Client) *PowerOnMachineParams {
	var ()
	return &PowerOnMachineParams{
		HTTPClient: client,
	}
}

/*PowerOnMachineParams contains all the parameters to send to the API endpoint
for the power on machine operation typically these are written to a http.Request
*/
type PowerOnMachineParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about

	*/
	APIVersion *string
	/*ID
	  The id of the Machine.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the power on machine params
func (o *PowerOnMachineParams) WithTimeout(timeout time.Duration) *PowerOnMachineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the power on machine params
func (o *PowerOnMachineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the power on machine params
func (o *PowerOnMachineParams) WithContext(ctx context.Context) *PowerOnMachineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the power on machine params
func (o *PowerOnMachineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the power on machine params
func (o *PowerOnMachineParams) WithHTTPClient(client *http.Client) *PowerOnMachineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the power on machine params
func (o *PowerOnMachineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the power on machine params
func (o *PowerOnMachineParams) WithAPIVersion(aPIVersion *string) *PowerOnMachineParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the power on machine params
func (o *PowerOnMachineParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the power on machine params
func (o *PowerOnMachineParams) WithID(id string) *PowerOnMachineParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the power on machine params
func (o *PowerOnMachineParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PowerOnMachineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
