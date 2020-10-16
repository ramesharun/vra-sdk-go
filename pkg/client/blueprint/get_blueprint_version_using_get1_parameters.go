// Code generated by go-swagger; DO NOT EDIT.

package blueprint

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
	"github.com/go-openapi/swag"
)

// NewGetBlueprintVersionUsingGET1Params creates a new GetBlueprintVersionUsingGET1Params object
// with the default values initialized.
func NewGetBlueprintVersionUsingGET1Params() *GetBlueprintVersionUsingGET1Params {
	var ()
	return &GetBlueprintVersionUsingGET1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBlueprintVersionUsingGET1ParamsWithTimeout creates a new GetBlueprintVersionUsingGET1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBlueprintVersionUsingGET1ParamsWithTimeout(timeout time.Duration) *GetBlueprintVersionUsingGET1Params {
	var ()
	return &GetBlueprintVersionUsingGET1Params{

		timeout: timeout,
	}
}

// NewGetBlueprintVersionUsingGET1ParamsWithContext creates a new GetBlueprintVersionUsingGET1Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetBlueprintVersionUsingGET1ParamsWithContext(ctx context.Context) *GetBlueprintVersionUsingGET1Params {
	var ()
	return &GetBlueprintVersionUsingGET1Params{

		Context: ctx,
	}
}

// NewGetBlueprintVersionUsingGET1ParamsWithHTTPClient creates a new GetBlueprintVersionUsingGET1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBlueprintVersionUsingGET1ParamsWithHTTPClient(client *http.Client) *GetBlueprintVersionUsingGET1Params {
	var ()
	return &GetBlueprintVersionUsingGET1Params{
		HTTPClient: client,
	}
}

/*GetBlueprintVersionUsingGET1Params contains all the parameters to send to the API endpoint
for the get blueprint version using get1 operation typically these are written to a http.Request
*/
type GetBlueprintVersionUsingGET1Params struct {

	/*DollarSelect
	  Fields to include in content. Use * to get all fields.

	*/
	DollarSelect []string
	/*BlueprintID
	  blueprintId

	*/
	BlueprintID strfmt.UUID
	/*Version
	  version

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get blueprint version using get1 params
func (o *GetBlueprintVersionUsingGET1Params) WithTimeout(timeout time.Duration) *GetBlueprintVersionUsingGET1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get blueprint version using get1 params
func (o *GetBlueprintVersionUsingGET1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get blueprint version using get1 params
func (o *GetBlueprintVersionUsingGET1Params) WithContext(ctx context.Context) *GetBlueprintVersionUsingGET1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get blueprint version using get1 params
func (o *GetBlueprintVersionUsingGET1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get blueprint version using get1 params
func (o *GetBlueprintVersionUsingGET1Params) WithHTTPClient(client *http.Client) *GetBlueprintVersionUsingGET1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get blueprint version using get1 params
func (o *GetBlueprintVersionUsingGET1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarSelect adds the dollarSelect to the get blueprint version using get1 params
func (o *GetBlueprintVersionUsingGET1Params) WithDollarSelect(dollarSelect []string) *GetBlueprintVersionUsingGET1Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the get blueprint version using get1 params
func (o *GetBlueprintVersionUsingGET1Params) SetDollarSelect(dollarSelect []string) {
	o.DollarSelect = dollarSelect
}

// WithBlueprintID adds the blueprintID to the get blueprint version using get1 params
func (o *GetBlueprintVersionUsingGET1Params) WithBlueprintID(blueprintID strfmt.UUID) *GetBlueprintVersionUsingGET1Params {
	o.SetBlueprintID(blueprintID)
	return o
}

// SetBlueprintID adds the blueprintId to the get blueprint version using get1 params
func (o *GetBlueprintVersionUsingGET1Params) SetBlueprintID(blueprintID strfmt.UUID) {
	o.BlueprintID = blueprintID
}

// WithVersion adds the version to the get blueprint version using get1 params
func (o *GetBlueprintVersionUsingGET1Params) WithVersion(version string) *GetBlueprintVersionUsingGET1Params {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get blueprint version using get1 params
func (o *GetBlueprintVersionUsingGET1Params) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetBlueprintVersionUsingGET1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesDollarSelect := o.DollarSelect

	joinedDollarSelect := swag.JoinByFormat(valuesDollarSelect, "multi")
	// query array param $select
	if err := r.SetQueryParam("$select", joinedDollarSelect...); err != nil {
		return err
	}

	// path param blueprintId
	if err := r.SetPathParam("blueprintId", o.BlueprintID.String()); err != nil {
		return err
	}

	// path param version
	if err := r.SetPathParam("version", o.Version); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
