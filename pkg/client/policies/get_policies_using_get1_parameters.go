// Code generated by go-swagger; DO NOT EDIT.

package policies

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

// NewGetPoliciesUsingGET1Params creates a new GetPoliciesUsingGET1Params object
// with the default values initialized.
func NewGetPoliciesUsingGET1Params() *GetPoliciesUsingGET1Params {
	var ()
	return &GetPoliciesUsingGET1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPoliciesUsingGET1ParamsWithTimeout creates a new GetPoliciesUsingGET1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPoliciesUsingGET1ParamsWithTimeout(timeout time.Duration) *GetPoliciesUsingGET1Params {
	var ()
	return &GetPoliciesUsingGET1Params{

		timeout: timeout,
	}
}

// NewGetPoliciesUsingGET1ParamsWithContext creates a new GetPoliciesUsingGET1Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetPoliciesUsingGET1ParamsWithContext(ctx context.Context) *GetPoliciesUsingGET1Params {
	var ()
	return &GetPoliciesUsingGET1Params{

		Context: ctx,
	}
}

// NewGetPoliciesUsingGET1ParamsWithHTTPClient creates a new GetPoliciesUsingGET1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPoliciesUsingGET1ParamsWithHTTPClient(client *http.Client) *GetPoliciesUsingGET1Params {
	var ()
	return &GetPoliciesUsingGET1Params{
		HTTPClient: client,
	}
}

/*GetPoliciesUsingGET1Params contains all the parameters to send to the API endpoint
for the get policies using get1 operation typically these are written to a http.Request
*/
type GetPoliciesUsingGET1Params struct {

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
	  The version of the API in yyyy-MM-dd format (UTC). If you do not specify explicitly an exact version, you will be calling the latest supported API version.

	*/
	APIVersion *string
	/*ComputeStats
	  computeStats

	*/
	ComputeStats *bool
	/*ExpandDefinition
	  Retrieves policy definition information for each returned policy.

	*/
	ExpandDefinition *bool
	/*Search
	  Matches will start with this string in their name or have this string somewhere in their description.

	*/
	Search *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get policies using get1 params
func (o *GetPoliciesUsingGET1Params) WithTimeout(timeout time.Duration) *GetPoliciesUsingGET1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get policies using get1 params
func (o *GetPoliciesUsingGET1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get policies using get1 params
func (o *GetPoliciesUsingGET1Params) WithContext(ctx context.Context) *GetPoliciesUsingGET1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get policies using get1 params
func (o *GetPoliciesUsingGET1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get policies using get1 params
func (o *GetPoliciesUsingGET1Params) WithHTTPClient(client *http.Client) *GetPoliciesUsingGET1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get policies using get1 params
func (o *GetPoliciesUsingGET1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarOrderby adds the dollarOrderby to the get policies using get1 params
func (o *GetPoliciesUsingGET1Params) WithDollarOrderby(dollarOrderby []string) *GetPoliciesUsingGET1Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the get policies using get1 params
func (o *GetPoliciesUsingGET1Params) SetDollarOrderby(dollarOrderby []string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSkip adds the dollarSkip to the get policies using get1 params
func (o *GetPoliciesUsingGET1Params) WithDollarSkip(dollarSkip *int32) *GetPoliciesUsingGET1Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the get policies using get1 params
func (o *GetPoliciesUsingGET1Params) SetDollarSkip(dollarSkip *int32) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the get policies using get1 params
func (o *GetPoliciesUsingGET1Params) WithDollarTop(dollarTop *int32) *GetPoliciesUsingGET1Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the get policies using get1 params
func (o *GetPoliciesUsingGET1Params) SetDollarTop(dollarTop *int32) {
	o.DollarTop = dollarTop
}

// WithAPIVersion adds the aPIVersion to the get policies using get1 params
func (o *GetPoliciesUsingGET1Params) WithAPIVersion(aPIVersion *string) *GetPoliciesUsingGET1Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get policies using get1 params
func (o *GetPoliciesUsingGET1Params) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithComputeStats adds the computeStats to the get policies using get1 params
func (o *GetPoliciesUsingGET1Params) WithComputeStats(computeStats *bool) *GetPoliciesUsingGET1Params {
	o.SetComputeStats(computeStats)
	return o
}

// SetComputeStats adds the computeStats to the get policies using get1 params
func (o *GetPoliciesUsingGET1Params) SetComputeStats(computeStats *bool) {
	o.ComputeStats = computeStats
}

// WithExpandDefinition adds the expandDefinition to the get policies using get1 params
func (o *GetPoliciesUsingGET1Params) WithExpandDefinition(expandDefinition *bool) *GetPoliciesUsingGET1Params {
	o.SetExpandDefinition(expandDefinition)
	return o
}

// SetExpandDefinition adds the expandDefinition to the get policies using get1 params
func (o *GetPoliciesUsingGET1Params) SetExpandDefinition(expandDefinition *bool) {
	o.ExpandDefinition = expandDefinition
}

// WithSearch adds the search to the get policies using get1 params
func (o *GetPoliciesUsingGET1Params) WithSearch(search *string) *GetPoliciesUsingGET1Params {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the get policies using get1 params
func (o *GetPoliciesUsingGET1Params) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *GetPoliciesUsingGET1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ComputeStats != nil {

		// query param computeStats
		var qrComputeStats bool
		if o.ComputeStats != nil {
			qrComputeStats = *o.ComputeStats
		}
		qComputeStats := swag.FormatBool(qrComputeStats)
		if qComputeStats != "" {
			if err := r.SetQueryParam("computeStats", qComputeStats); err != nil {
				return err
			}
		}

	}

	if o.ExpandDefinition != nil {

		// query param expandDefinition
		var qrExpandDefinition bool
		if o.ExpandDefinition != nil {
			qrExpandDefinition = *o.ExpandDefinition
		}
		qExpandDefinition := swag.FormatBool(qrExpandDefinition)
		if qExpandDefinition != "" {
			if err := r.SetQueryParam("expandDefinition", qExpandDefinition); err != nil {
				return err
			}
		}

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
