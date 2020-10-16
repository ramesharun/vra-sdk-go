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
)

// NewGetUpfrontPriceResponseForCatalogItemUsingGETParams creates a new GetUpfrontPriceResponseForCatalogItemUsingGETParams object
// with the default values initialized.
func NewGetUpfrontPriceResponseForCatalogItemUsingGETParams() *GetUpfrontPriceResponseForCatalogItemUsingGETParams {
	var ()
	return &GetUpfrontPriceResponseForCatalogItemUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUpfrontPriceResponseForCatalogItemUsingGETParamsWithTimeout creates a new GetUpfrontPriceResponseForCatalogItemUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUpfrontPriceResponseForCatalogItemUsingGETParamsWithTimeout(timeout time.Duration) *GetUpfrontPriceResponseForCatalogItemUsingGETParams {
	var ()
	return &GetUpfrontPriceResponseForCatalogItemUsingGETParams{

		timeout: timeout,
	}
}

// NewGetUpfrontPriceResponseForCatalogItemUsingGETParamsWithContext creates a new GetUpfrontPriceResponseForCatalogItemUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUpfrontPriceResponseForCatalogItemUsingGETParamsWithContext(ctx context.Context) *GetUpfrontPriceResponseForCatalogItemUsingGETParams {
	var ()
	return &GetUpfrontPriceResponseForCatalogItemUsingGETParams{

		Context: ctx,
	}
}

// NewGetUpfrontPriceResponseForCatalogItemUsingGETParamsWithHTTPClient creates a new GetUpfrontPriceResponseForCatalogItemUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUpfrontPriceResponseForCatalogItemUsingGETParamsWithHTTPClient(client *http.Client) *GetUpfrontPriceResponseForCatalogItemUsingGETParams {
	var ()
	return &GetUpfrontPriceResponseForCatalogItemUsingGETParams{
		HTTPClient: client,
	}
}

/*GetUpfrontPriceResponseForCatalogItemUsingGETParams contains all the parameters to send to the API endpoint
for the get upfront price response for catalog item using g e t operation typically these are written to a http.Request
*/
type GetUpfrontPriceResponseForCatalogItemUsingGETParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /catalog/api/about

	*/
	APIVersion *string
	/*ID
	  Catalog Item ID

	*/
	ID strfmt.UUID
	/*UpfrontPriceID
	  Upfront Price Request ID

	*/
	UpfrontPriceID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get upfront price response for catalog item using get params
func (o *GetUpfrontPriceResponseForCatalogItemUsingGETParams) WithTimeout(timeout time.Duration) *GetUpfrontPriceResponseForCatalogItemUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get upfront price response for catalog item using get params
func (o *GetUpfrontPriceResponseForCatalogItemUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get upfront price response for catalog item using get params
func (o *GetUpfrontPriceResponseForCatalogItemUsingGETParams) WithContext(ctx context.Context) *GetUpfrontPriceResponseForCatalogItemUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get upfront price response for catalog item using get params
func (o *GetUpfrontPriceResponseForCatalogItemUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get upfront price response for catalog item using get params
func (o *GetUpfrontPriceResponseForCatalogItemUsingGETParams) WithHTTPClient(client *http.Client) *GetUpfrontPriceResponseForCatalogItemUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get upfront price response for catalog item using get params
func (o *GetUpfrontPriceResponseForCatalogItemUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get upfront price response for catalog item using get params
func (o *GetUpfrontPriceResponseForCatalogItemUsingGETParams) WithAPIVersion(aPIVersion *string) *GetUpfrontPriceResponseForCatalogItemUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get upfront price response for catalog item using get params
func (o *GetUpfrontPriceResponseForCatalogItemUsingGETParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get upfront price response for catalog item using get params
func (o *GetUpfrontPriceResponseForCatalogItemUsingGETParams) WithID(id strfmt.UUID) *GetUpfrontPriceResponseForCatalogItemUsingGETParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get upfront price response for catalog item using get params
func (o *GetUpfrontPriceResponseForCatalogItemUsingGETParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithUpfrontPriceID adds the upfrontPriceID to the get upfront price response for catalog item using get params
func (o *GetUpfrontPriceResponseForCatalogItemUsingGETParams) WithUpfrontPriceID(upfrontPriceID strfmt.UUID) *GetUpfrontPriceResponseForCatalogItemUsingGETParams {
	o.SetUpfrontPriceID(upfrontPriceID)
	return o
}

// SetUpfrontPriceID adds the upfrontPriceId to the get upfront price response for catalog item using get params
func (o *GetUpfrontPriceResponseForCatalogItemUsingGETParams) SetUpfrontPriceID(upfrontPriceID strfmt.UUID) {
	o.UpfrontPriceID = upfrontPriceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUpfrontPriceResponseForCatalogItemUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param upfrontPriceId
	if err := r.SetPathParam("upfrontPriceId", o.UpfrontPriceID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
