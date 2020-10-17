// Code generated by go-swagger; DO NOT EDIT.

package catalog_items

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

// NewGetCatalogItemUsingGET1Params creates a new GetCatalogItemUsingGET1Params object
// with the default values initialized.
func NewGetCatalogItemUsingGET1Params() *GetCatalogItemUsingGET1Params {
	var ()
	return &GetCatalogItemUsingGET1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCatalogItemUsingGET1ParamsWithTimeout creates a new GetCatalogItemUsingGET1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCatalogItemUsingGET1ParamsWithTimeout(timeout time.Duration) *GetCatalogItemUsingGET1Params {
	var ()
	return &GetCatalogItemUsingGET1Params{

		timeout: timeout,
	}
}

// NewGetCatalogItemUsingGET1ParamsWithContext creates a new GetCatalogItemUsingGET1Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetCatalogItemUsingGET1ParamsWithContext(ctx context.Context) *GetCatalogItemUsingGET1Params {
	var ()
	return &GetCatalogItemUsingGET1Params{

		Context: ctx,
	}
}

// NewGetCatalogItemUsingGET1ParamsWithHTTPClient creates a new GetCatalogItemUsingGET1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCatalogItemUsingGET1ParamsWithHTTPClient(client *http.Client) *GetCatalogItemUsingGET1Params {
	var ()
	return &GetCatalogItemUsingGET1Params{
		HTTPClient: client,
	}
}

/*GetCatalogItemUsingGET1Params contains all the parameters to send to the API endpoint
for the get catalog item using get1 operation typically these are written to a http.Request
*/
type GetCatalogItemUsingGET1Params struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). If you do not specify explicitly an exact version, you will be calling the latest supported API version.

	*/
	APIVersion *string
	/*ExpandProjects
	  Retrieves the 'projects' field of the catalog item

	*/
	ExpandProjects *bool
	/*ID
	  Catalog item ID

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get catalog item using get1 params
func (o *GetCatalogItemUsingGET1Params) WithTimeout(timeout time.Duration) *GetCatalogItemUsingGET1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get catalog item using get1 params
func (o *GetCatalogItemUsingGET1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get catalog item using get1 params
func (o *GetCatalogItemUsingGET1Params) WithContext(ctx context.Context) *GetCatalogItemUsingGET1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get catalog item using get1 params
func (o *GetCatalogItemUsingGET1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get catalog item using get1 params
func (o *GetCatalogItemUsingGET1Params) WithHTTPClient(client *http.Client) *GetCatalogItemUsingGET1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get catalog item using get1 params
func (o *GetCatalogItemUsingGET1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get catalog item using get1 params
func (o *GetCatalogItemUsingGET1Params) WithAPIVersion(aPIVersion *string) *GetCatalogItemUsingGET1Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get catalog item using get1 params
func (o *GetCatalogItemUsingGET1Params) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithExpandProjects adds the expandProjects to the get catalog item using get1 params
func (o *GetCatalogItemUsingGET1Params) WithExpandProjects(expandProjects *bool) *GetCatalogItemUsingGET1Params {
	o.SetExpandProjects(expandProjects)
	return o
}

// SetExpandProjects adds the expandProjects to the get catalog item using get1 params
func (o *GetCatalogItemUsingGET1Params) SetExpandProjects(expandProjects *bool) {
	o.ExpandProjects = expandProjects
}

// WithID adds the id to the get catalog item using get1 params
func (o *GetCatalogItemUsingGET1Params) WithID(id strfmt.UUID) *GetCatalogItemUsingGET1Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the get catalog item using get1 params
func (o *GetCatalogItemUsingGET1Params) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetCatalogItemUsingGET1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ExpandProjects != nil {

		// query param expandProjects
		var qrExpandProjects bool
		if o.ExpandProjects != nil {
			qrExpandProjects = *o.ExpandProjects
		}
		qExpandProjects := swag.FormatBool(qrExpandProjects)
		if qExpandProjects != "" {
			if err := r.SetQueryParam("expandProjects", qExpandProjects); err != nil {
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
