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

// NewGetMachineSnapshotsParams creates a new GetMachineSnapshotsParams object
// with the default values initialized.
func NewGetMachineSnapshotsParams() *GetMachineSnapshotsParams {
	var ()
	return &GetMachineSnapshotsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMachineSnapshotsParamsWithTimeout creates a new GetMachineSnapshotsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMachineSnapshotsParamsWithTimeout(timeout time.Duration) *GetMachineSnapshotsParams {
	var ()
	return &GetMachineSnapshotsParams{

		timeout: timeout,
	}
}

// NewGetMachineSnapshotsParamsWithContext creates a new GetMachineSnapshotsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMachineSnapshotsParamsWithContext(ctx context.Context) *GetMachineSnapshotsParams {
	var ()
	return &GetMachineSnapshotsParams{

		Context: ctx,
	}
}

// NewGetMachineSnapshotsParamsWithHTTPClient creates a new GetMachineSnapshotsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMachineSnapshotsParamsWithHTTPClient(client *http.Client) *GetMachineSnapshotsParams {
	var ()
	return &GetMachineSnapshotsParams{
		HTTPClient: client,
	}
}

/*GetMachineSnapshotsParams contains all the parameters to send to the API endpoint
for the get machine snapshots operation typically these are written to a http.Request
*/
type GetMachineSnapshotsParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about

	*/
	APIVersion *string
	/*ID
	  The ID of the machine.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get machine snapshots params
func (o *GetMachineSnapshotsParams) WithTimeout(timeout time.Duration) *GetMachineSnapshotsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get machine snapshots params
func (o *GetMachineSnapshotsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get machine snapshots params
func (o *GetMachineSnapshotsParams) WithContext(ctx context.Context) *GetMachineSnapshotsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get machine snapshots params
func (o *GetMachineSnapshotsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get machine snapshots params
func (o *GetMachineSnapshotsParams) WithHTTPClient(client *http.Client) *GetMachineSnapshotsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get machine snapshots params
func (o *GetMachineSnapshotsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get machine snapshots params
func (o *GetMachineSnapshotsParams) WithAPIVersion(aPIVersion *string) *GetMachineSnapshotsParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get machine snapshots params
func (o *GetMachineSnapshotsParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get machine snapshots params
func (o *GetMachineSnapshotsParams) WithID(id string) *GetMachineSnapshotsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get machine snapshots params
func (o *GetMachineSnapshotsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetMachineSnapshotsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
