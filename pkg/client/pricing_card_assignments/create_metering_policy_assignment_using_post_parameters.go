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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// NewCreateMeteringPolicyAssignmentUsingPOSTParams creates a new CreateMeteringPolicyAssignmentUsingPOSTParams object
// with the default values initialized.
func NewCreateMeteringPolicyAssignmentUsingPOSTParams() *CreateMeteringPolicyAssignmentUsingPOSTParams {
	var ()
	return &CreateMeteringPolicyAssignmentUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateMeteringPolicyAssignmentUsingPOSTParamsWithTimeout creates a new CreateMeteringPolicyAssignmentUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateMeteringPolicyAssignmentUsingPOSTParamsWithTimeout(timeout time.Duration) *CreateMeteringPolicyAssignmentUsingPOSTParams {
	var ()
	return &CreateMeteringPolicyAssignmentUsingPOSTParams{

		timeout: timeout,
	}
}

// NewCreateMeteringPolicyAssignmentUsingPOSTParamsWithContext creates a new CreateMeteringPolicyAssignmentUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateMeteringPolicyAssignmentUsingPOSTParamsWithContext(ctx context.Context) *CreateMeteringPolicyAssignmentUsingPOSTParams {
	var ()
	return &CreateMeteringPolicyAssignmentUsingPOSTParams{

		Context: ctx,
	}
}

// NewCreateMeteringPolicyAssignmentUsingPOSTParamsWithHTTPClient creates a new CreateMeteringPolicyAssignmentUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateMeteringPolicyAssignmentUsingPOSTParamsWithHTTPClient(client *http.Client) *CreateMeteringPolicyAssignmentUsingPOSTParams {
	var ()
	return &CreateMeteringPolicyAssignmentUsingPOSTParams{
		HTTPClient: client,
	}
}

/*CreateMeteringPolicyAssignmentUsingPOSTParams contains all the parameters to send to the API endpoint
for the create metering policy assignment using p o s t operation typically these are written to a http.Request
*/
type CreateMeteringPolicyAssignmentUsingPOSTParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /catalog/api/about

	*/
	APIVersion *string
	/*MeteringPolicyAssignment
	  The pricing card assignment to be created

	*/
	MeteringPolicyAssignment *models.MeteringPolicyAssignment

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create metering policy assignment using p o s t params
func (o *CreateMeteringPolicyAssignmentUsingPOSTParams) WithTimeout(timeout time.Duration) *CreateMeteringPolicyAssignmentUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create metering policy assignment using p o s t params
func (o *CreateMeteringPolicyAssignmentUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create metering policy assignment using p o s t params
func (o *CreateMeteringPolicyAssignmentUsingPOSTParams) WithContext(ctx context.Context) *CreateMeteringPolicyAssignmentUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create metering policy assignment using p o s t params
func (o *CreateMeteringPolicyAssignmentUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create metering policy assignment using p o s t params
func (o *CreateMeteringPolicyAssignmentUsingPOSTParams) WithHTTPClient(client *http.Client) *CreateMeteringPolicyAssignmentUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create metering policy assignment using p o s t params
func (o *CreateMeteringPolicyAssignmentUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create metering policy assignment using p o s t params
func (o *CreateMeteringPolicyAssignmentUsingPOSTParams) WithAPIVersion(aPIVersion *string) *CreateMeteringPolicyAssignmentUsingPOSTParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create metering policy assignment using p o s t params
func (o *CreateMeteringPolicyAssignmentUsingPOSTParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithMeteringPolicyAssignment adds the meteringPolicyAssignment to the create metering policy assignment using p o s t params
func (o *CreateMeteringPolicyAssignmentUsingPOSTParams) WithMeteringPolicyAssignment(meteringPolicyAssignment *models.MeteringPolicyAssignment) *CreateMeteringPolicyAssignmentUsingPOSTParams {
	o.SetMeteringPolicyAssignment(meteringPolicyAssignment)
	return o
}

// SetMeteringPolicyAssignment adds the meteringPolicyAssignment to the create metering policy assignment using p o s t params
func (o *CreateMeteringPolicyAssignmentUsingPOSTParams) SetMeteringPolicyAssignment(meteringPolicyAssignment *models.MeteringPolicyAssignment) {
	o.MeteringPolicyAssignment = meteringPolicyAssignment
}

// WriteToRequest writes these params to a swagger request
func (o *CreateMeteringPolicyAssignmentUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.MeteringPolicyAssignment != nil {
		if err := r.SetBodyParam(o.MeteringPolicyAssignment); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}