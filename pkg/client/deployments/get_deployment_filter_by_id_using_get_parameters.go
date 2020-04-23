// Code generated by go-swagger; DO NOT EDIT.

package deployments

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

// NewGetDeploymentFilterByIDUsingGETParams creates a new GetDeploymentFilterByIDUsingGETParams object
// with the default values initialized.
func NewGetDeploymentFilterByIDUsingGETParams() *GetDeploymentFilterByIDUsingGETParams {
	var ()
	return &GetDeploymentFilterByIDUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeploymentFilterByIDUsingGETParamsWithTimeout creates a new GetDeploymentFilterByIDUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDeploymentFilterByIDUsingGETParamsWithTimeout(timeout time.Duration) *GetDeploymentFilterByIDUsingGETParams {
	var ()
	return &GetDeploymentFilterByIDUsingGETParams{

		timeout: timeout,
	}
}

// NewGetDeploymentFilterByIDUsingGETParamsWithContext creates a new GetDeploymentFilterByIDUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDeploymentFilterByIDUsingGETParamsWithContext(ctx context.Context) *GetDeploymentFilterByIDUsingGETParams {
	var ()
	return &GetDeploymentFilterByIDUsingGETParams{

		Context: ctx,
	}
}

// NewGetDeploymentFilterByIDUsingGETParamsWithHTTPClient creates a new GetDeploymentFilterByIDUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDeploymentFilterByIDUsingGETParamsWithHTTPClient(client *http.Client) *GetDeploymentFilterByIDUsingGETParams {
	var ()
	return &GetDeploymentFilterByIDUsingGETParams{
		HTTPClient: client,
	}
}

/*GetDeploymentFilterByIDUsingGETParams contains all the parameters to send to the API endpoint
for the get deployment filter by Id using g e t operation typically these are written to a http.Request
*/
type GetDeploymentFilterByIDUsingGETParams struct {

	/*DollarOrderby
	  Sorting criteria in the format: property (asc|desc). Default sort order is ascending. Multiple sort criteria are supported.

	*/
	DollarOrderby []string
	/*DollarSkip
	  Number of records you want to skip

	*/
	DollarSkip *int32
	/*DollarTop
	  Number of records you want

	*/
	DollarTop *int32
	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /catalog/api/about

	*/
	APIVersion *string
	/*FilterID
	  Filter Id

	*/
	FilterID string
	/*Search
	  Search string for filters

	*/
	Search *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get deployment filter by Id using get params
func (o *GetDeploymentFilterByIDUsingGETParams) WithTimeout(timeout time.Duration) *GetDeploymentFilterByIDUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get deployment filter by Id using get params
func (o *GetDeploymentFilterByIDUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get deployment filter by Id using get params
func (o *GetDeploymentFilterByIDUsingGETParams) WithContext(ctx context.Context) *GetDeploymentFilterByIDUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get deployment filter by Id using get params
func (o *GetDeploymentFilterByIDUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get deployment filter by Id using get params
func (o *GetDeploymentFilterByIDUsingGETParams) WithHTTPClient(client *http.Client) *GetDeploymentFilterByIDUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get deployment filter by Id using get params
func (o *GetDeploymentFilterByIDUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarOrderby adds the dollarOrderby to the get deployment filter by Id using get params
func (o *GetDeploymentFilterByIDUsingGETParams) WithDollarOrderby(dollarOrderby []string) *GetDeploymentFilterByIDUsingGETParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the get deployment filter by Id using get params
func (o *GetDeploymentFilterByIDUsingGETParams) SetDollarOrderby(dollarOrderby []string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSkip adds the dollarSkip to the get deployment filter by Id using get params
func (o *GetDeploymentFilterByIDUsingGETParams) WithDollarSkip(dollarSkip *int32) *GetDeploymentFilterByIDUsingGETParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the get deployment filter by Id using get params
func (o *GetDeploymentFilterByIDUsingGETParams) SetDollarSkip(dollarSkip *int32) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the get deployment filter by Id using get params
func (o *GetDeploymentFilterByIDUsingGETParams) WithDollarTop(dollarTop *int32) *GetDeploymentFilterByIDUsingGETParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the get deployment filter by Id using get params
func (o *GetDeploymentFilterByIDUsingGETParams) SetDollarTop(dollarTop *int32) {
	o.DollarTop = dollarTop
}

// WithAPIVersion adds the aPIVersion to the get deployment filter by Id using get params
func (o *GetDeploymentFilterByIDUsingGETParams) WithAPIVersion(aPIVersion *string) *GetDeploymentFilterByIDUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get deployment filter by Id using get params
func (o *GetDeploymentFilterByIDUsingGETParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithFilterID adds the filterID to the get deployment filter by Id using get params
func (o *GetDeploymentFilterByIDUsingGETParams) WithFilterID(filterID string) *GetDeploymentFilterByIDUsingGETParams {
	o.SetFilterID(filterID)
	return o
}

// SetFilterID adds the filterId to the get deployment filter by Id using get params
func (o *GetDeploymentFilterByIDUsingGETParams) SetFilterID(filterID string) {
	o.FilterID = filterID
}

// WithSearch adds the search to the get deployment filter by Id using get params
func (o *GetDeploymentFilterByIDUsingGETParams) WithSearch(search *string) *GetDeploymentFilterByIDUsingGETParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the get deployment filter by Id using get params
func (o *GetDeploymentFilterByIDUsingGETParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeploymentFilterByIDUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesDollarOrderby := o.DollarOrderby

	joinedDollarOrderby := swag.JoinByFormat(valuesDollarOrderby, "multi")
	// query array param $orderby
	if err := r.SetQueryParam("$orderby", joinedDollarOrderby...); err != nil {
		return err
	}

	if o.DollarSkip != nil {

		// query param $skip
		var qrDollarSkip int32
		if o.DollarSkip != nil {
			qrDollarSkip = *o.DollarSkip
		}
		qDollarSkip := swag.FormatInt32(qrDollarSkip)
		if qDollarSkip != "" {
			if err := r.SetQueryParam("$skip", qDollarSkip); err != nil {
				return err
			}
		}

	}

	if o.DollarTop != nil {

		// query param $top
		var qrDollarTop int32
		if o.DollarTop != nil {
			qrDollarTop = *o.DollarTop
		}
		qDollarTop := swag.FormatInt32(qrDollarTop)
		if qDollarTop != "" {
			if err := r.SetQueryParam("$top", qDollarTop); err != nil {
				return err
			}
		}

	}

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

	// path param filterId
	if err := r.SetPathParam("filterId", o.FilterID); err != nil {
		return err
	}

	if o.Search != nil {

		// query param search
		var qrSearch string
		if o.Search != nil {
			qrSearch = *o.Search
		}
		qSearch := qrSearch
		if qSearch != "" {
			if err := r.SetQueryParam("search", qSearch); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
