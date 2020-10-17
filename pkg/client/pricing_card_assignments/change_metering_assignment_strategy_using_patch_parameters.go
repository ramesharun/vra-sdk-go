// Code generated by go-swagger; DO NOT EDIT.

package pricing_card_assignments

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

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// NewChangeMeteringAssignmentStrategyUsingPATCHParams creates a new ChangeMeteringAssignmentStrategyUsingPATCHParams object
// with the default values initialized.
func NewChangeMeteringAssignmentStrategyUsingPATCHParams() *ChangeMeteringAssignmentStrategyUsingPATCHParams {
	var ()
	return &ChangeMeteringAssignmentStrategyUsingPATCHParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChangeMeteringAssignmentStrategyUsingPATCHParamsWithTimeout creates a new ChangeMeteringAssignmentStrategyUsingPATCHParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChangeMeteringAssignmentStrategyUsingPATCHParamsWithTimeout(timeout time.Duration) *ChangeMeteringAssignmentStrategyUsingPATCHParams {
	var ()
	return &ChangeMeteringAssignmentStrategyUsingPATCHParams{

		timeout: timeout,
	}
}

// NewChangeMeteringAssignmentStrategyUsingPATCHParamsWithContext creates a new ChangeMeteringAssignmentStrategyUsingPATCHParams object
// with the default values initialized, and the ability to set a context for a request
func NewChangeMeteringAssignmentStrategyUsingPATCHParamsWithContext(ctx context.Context) *ChangeMeteringAssignmentStrategyUsingPATCHParams {
	var ()
	return &ChangeMeteringAssignmentStrategyUsingPATCHParams{

		Context: ctx,
	}
}

// NewChangeMeteringAssignmentStrategyUsingPATCHParamsWithHTTPClient creates a new ChangeMeteringAssignmentStrategyUsingPATCHParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChangeMeteringAssignmentStrategyUsingPATCHParamsWithHTTPClient(client *http.Client) *ChangeMeteringAssignmentStrategyUsingPATCHParams {
	var ()
	return &ChangeMeteringAssignmentStrategyUsingPATCHParams{
		HTTPClient: client,
	}
}

/*ChangeMeteringAssignmentStrategyUsingPATCHParams contains all the parameters to send to the API endpoint
for the change metering assignment strategy using p a t c h operation typically these are written to a http.Request
*/
type ChangeMeteringAssignmentStrategyUsingPATCHParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). If you do not specify explicitly an exact version, you will be calling the latest supported API version.

	*/
	APIVersion *string
	/*MeteringAssignmentStrategy
	  Pricing card assignment strategy with 'EntityType' to override the strategy for the org

	*/
	MeteringAssignmentStrategy *models.MeteringAssignmentStrategy

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the change metering assignment strategy using p a t c h params
func (o *ChangeMeteringAssignmentStrategyUsingPATCHParams) WithTimeout(timeout time.Duration) *ChangeMeteringAssignmentStrategyUsingPATCHParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change metering assignment strategy using p a t c h params
func (o *ChangeMeteringAssignmentStrategyUsingPATCHParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change metering assignment strategy using p a t c h params
func (o *ChangeMeteringAssignmentStrategyUsingPATCHParams) WithContext(ctx context.Context) *ChangeMeteringAssignmentStrategyUsingPATCHParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change metering assignment strategy using p a t c h params
func (o *ChangeMeteringAssignmentStrategyUsingPATCHParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change metering assignment strategy using p a t c h params
func (o *ChangeMeteringAssignmentStrategyUsingPATCHParams) WithHTTPClient(client *http.Client) *ChangeMeteringAssignmentStrategyUsingPATCHParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change metering assignment strategy using p a t c h params
func (o *ChangeMeteringAssignmentStrategyUsingPATCHParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the change metering assignment strategy using p a t c h params
func (o *ChangeMeteringAssignmentStrategyUsingPATCHParams) WithAPIVersion(aPIVersion *string) *ChangeMeteringAssignmentStrategyUsingPATCHParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the change metering assignment strategy using p a t c h params
func (o *ChangeMeteringAssignmentStrategyUsingPATCHParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithMeteringAssignmentStrategy adds the meteringAssignmentStrategy to the change metering assignment strategy using p a t c h params
func (o *ChangeMeteringAssignmentStrategyUsingPATCHParams) WithMeteringAssignmentStrategy(meteringAssignmentStrategy *models.MeteringAssignmentStrategy) *ChangeMeteringAssignmentStrategyUsingPATCHParams {
	o.SetMeteringAssignmentStrategy(meteringAssignmentStrategy)
	return o
}

// SetMeteringAssignmentStrategy adds the meteringAssignmentStrategy to the change metering assignment strategy using p a t c h params
func (o *ChangeMeteringAssignmentStrategyUsingPATCHParams) SetMeteringAssignmentStrategy(meteringAssignmentStrategy *models.MeteringAssignmentStrategy) {
	o.MeteringAssignmentStrategy = meteringAssignmentStrategy
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeMeteringAssignmentStrategyUsingPATCHParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.MeteringAssignmentStrategy != nil {
		if err := r.SetBodyParam(o.MeteringAssignmentStrategy); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
