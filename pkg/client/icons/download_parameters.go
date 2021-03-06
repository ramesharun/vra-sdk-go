// Code generated by go-swagger; DO NOT EDIT.

package icons

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

// NewDownloadParams creates a new DownloadParams object
// with the default values initialized.
func NewDownloadParams() *DownloadParams {
	var ()
	return &DownloadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDownloadParamsWithTimeout creates a new DownloadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDownloadParamsWithTimeout(timeout time.Duration) *DownloadParams {
	var ()
	return &DownloadParams{

		timeout: timeout,
	}
}

// NewDownloadParamsWithContext creates a new DownloadParams object
// with the default values initialized, and the ability to set a context for a request
func NewDownloadParamsWithContext(ctx context.Context) *DownloadParams {
	var ()
	return &DownloadParams{

		Context: ctx,
	}
}

// NewDownloadParamsWithHTTPClient creates a new DownloadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDownloadParamsWithHTTPClient(client *http.Client) *DownloadParams {
	var ()
	return &DownloadParams{
		HTTPClient: client,
	}
}

/*DownloadParams contains all the parameters to send to the API endpoint
for the download operation typically these are written to a http.Request
*/
type DownloadParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). If you do not specify explicitly an exact version, you will be calling the latest supported API version.

	*/
	APIVersion *string
	/*ID
	  Icon id

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the download params
func (o *DownloadParams) WithTimeout(timeout time.Duration) *DownloadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the download params
func (o *DownloadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the download params
func (o *DownloadParams) WithContext(ctx context.Context) *DownloadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the download params
func (o *DownloadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the download params
func (o *DownloadParams) WithHTTPClient(client *http.Client) *DownloadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the download params
func (o *DownloadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the download params
func (o *DownloadParams) WithAPIVersion(aPIVersion *string) *DownloadParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the download params
func (o *DownloadParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the download params
func (o *DownloadParams) WithID(id strfmt.UUID) *DownloadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the download params
func (o *DownloadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DownloadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
