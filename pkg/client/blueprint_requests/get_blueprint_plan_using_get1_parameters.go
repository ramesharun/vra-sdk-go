// Code generated by go-swagger; DO NOT EDIT.

package blueprint_requests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetBlueprintPlanUsingGET1Params creates a new GetBlueprintPlanUsingGET1Params object
// with the default values initialized.
func NewGetBlueprintPlanUsingGET1Params() *GetBlueprintPlanUsingGET1Params {
	var ()
	return &GetBlueprintPlanUsingGET1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBlueprintPlanUsingGET1ParamsWithTimeout creates a new GetBlueprintPlanUsingGET1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBlueprintPlanUsingGET1ParamsWithTimeout(timeout time.Duration) *GetBlueprintPlanUsingGET1Params {
	var ()
	return &GetBlueprintPlanUsingGET1Params{

		timeout: timeout,
	}
}

// NewGetBlueprintPlanUsingGET1ParamsWithContext creates a new GetBlueprintPlanUsingGET1Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetBlueprintPlanUsingGET1ParamsWithContext(ctx context.Context) *GetBlueprintPlanUsingGET1Params {
	var ()
	return &GetBlueprintPlanUsingGET1Params{

		Context: ctx,
	}
}

// NewGetBlueprintPlanUsingGET1ParamsWithHTTPClient creates a new GetBlueprintPlanUsingGET1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBlueprintPlanUsingGET1ParamsWithHTTPClient(client *http.Client) *GetBlueprintPlanUsingGET1Params {
	var ()
	return &GetBlueprintPlanUsingGET1Params{
		HTTPClient: client,
	}
}

/*GetBlueprintPlanUsingGET1Params contains all the parameters to send to the API endpoint
for the get blueprint plan using get1 operation typically these are written to a http.Request
*/
type GetBlueprintPlanUsingGET1Params struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /blueprint/api/about

	*/
	APIVersion *string
	/*Expand
	  Detailed plan

	*/
	Expand *bool
	/*RequestID
	  requestId

	*/
	RequestID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get blueprint plan using get1 params
func (o *GetBlueprintPlanUsingGET1Params) WithTimeout(timeout time.Duration) *GetBlueprintPlanUsingGET1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get blueprint plan using get1 params
func (o *GetBlueprintPlanUsingGET1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get blueprint plan using get1 params
func (o *GetBlueprintPlanUsingGET1Params) WithContext(ctx context.Context) *GetBlueprintPlanUsingGET1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get blueprint plan using get1 params
func (o *GetBlueprintPlanUsingGET1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get blueprint plan using get1 params
func (o *GetBlueprintPlanUsingGET1Params) WithHTTPClient(client *http.Client) *GetBlueprintPlanUsingGET1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get blueprint plan using get1 params
func (o *GetBlueprintPlanUsingGET1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get blueprint plan using get1 params
func (o *GetBlueprintPlanUsingGET1Params) WithAPIVersion(aPIVersion *string) *GetBlueprintPlanUsingGET1Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get blueprint plan using get1 params
func (o *GetBlueprintPlanUsingGET1Params) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithExpand adds the expand to the get blueprint plan using get1 params
func (o *GetBlueprintPlanUsingGET1Params) WithExpand(expand *bool) *GetBlueprintPlanUsingGET1Params {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get blueprint plan using get1 params
func (o *GetBlueprintPlanUsingGET1Params) SetExpand(expand *bool) {
	o.Expand = expand
}

// WithRequestID adds the requestID to the get blueprint plan using get1 params
func (o *GetBlueprintPlanUsingGET1Params) WithRequestID(requestID strfmt.UUID) *GetBlueprintPlanUsingGET1Params {
	o.SetRequestID(requestID)
	return o
}

// SetRequestID adds the requestId to the get blueprint plan using get1 params
func (o *GetBlueprintPlanUsingGET1Params) SetRequestID(requestID strfmt.UUID) {
	o.RequestID = requestID
}

// WriteToRequest writes these params to a swagger request
func (o *GetBlueprintPlanUsingGET1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Expand != nil {

		// query param expand
		var qrExpand bool
		if o.Expand != nil {
			qrExpand = *o.Expand
		}
		qExpand := swag.FormatBool(qrExpand)
		if qExpand != "" {
			if err := r.SetQueryParam("expand", qExpand); err != nil {
				return err
			}
		}

	}

	// path param requestId
	if err := r.SetPathParam("requestId", o.RequestID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}