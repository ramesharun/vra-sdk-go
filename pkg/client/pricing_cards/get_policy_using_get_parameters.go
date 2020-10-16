// Code generated by go-swagger; DO NOT EDIT.

package pricing_cards

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

// NewGetPolicyUsingGETParams creates a new GetPolicyUsingGETParams object
// with the default values initialized.
func NewGetPolicyUsingGETParams() *GetPolicyUsingGETParams {
	var ()
	return &GetPolicyUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPolicyUsingGETParamsWithTimeout creates a new GetPolicyUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPolicyUsingGETParamsWithTimeout(timeout time.Duration) *GetPolicyUsingGETParams {
	var ()
	return &GetPolicyUsingGETParams{

		timeout: timeout,
	}
}

// NewGetPolicyUsingGETParamsWithContext creates a new GetPolicyUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPolicyUsingGETParamsWithContext(ctx context.Context) *GetPolicyUsingGETParams {
	var ()
	return &GetPolicyUsingGETParams{

		Context: ctx,
	}
}

// NewGetPolicyUsingGETParamsWithHTTPClient creates a new GetPolicyUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPolicyUsingGETParamsWithHTTPClient(client *http.Client) *GetPolicyUsingGETParams {
	var ()
	return &GetPolicyUsingGETParams{
		HTTPClient: client,
	}
}

/*GetPolicyUsingGETParams contains all the parameters to send to the API endpoint
for the get policy using g e t operation typically these are written to a http.Request
*/
type GetPolicyUsingGETParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /catalog/api/about

	*/
	APIVersion *string
	/*ID
	  pricing card Id

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get policy using get params
func (o *GetPolicyUsingGETParams) WithTimeout(timeout time.Duration) *GetPolicyUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get policy using get params
func (o *GetPolicyUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get policy using get params
func (o *GetPolicyUsingGETParams) WithContext(ctx context.Context) *GetPolicyUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get policy using get params
func (o *GetPolicyUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get policy using get params
func (o *GetPolicyUsingGETParams) WithHTTPClient(client *http.Client) *GetPolicyUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get policy using get params
func (o *GetPolicyUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get policy using get params
func (o *GetPolicyUsingGETParams) WithAPIVersion(aPIVersion *string) *GetPolicyUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get policy using get params
func (o *GetPolicyUsingGETParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get policy using get params
func (o *GetPolicyUsingGETParams) WithID(id strfmt.UUID) *GetPolicyUsingGETParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get policy using get params
func (o *GetPolicyUsingGETParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetPolicyUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
