// Code generated by go-swagger; DO NOT EDIT.

package blueprint_terraform_integrations

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

// NewCreateBlueprintMappingUsingPOST1Params creates a new CreateBlueprintMappingUsingPOST1Params object
// with the default values initialized.
func NewCreateBlueprintMappingUsingPOST1Params() *CreateBlueprintMappingUsingPOST1Params {
	var ()
	return &CreateBlueprintMappingUsingPOST1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateBlueprintMappingUsingPOST1ParamsWithTimeout creates a new CreateBlueprintMappingUsingPOST1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateBlueprintMappingUsingPOST1ParamsWithTimeout(timeout time.Duration) *CreateBlueprintMappingUsingPOST1Params {
	var ()
	return &CreateBlueprintMappingUsingPOST1Params{

		timeout: timeout,
	}
}

// NewCreateBlueprintMappingUsingPOST1ParamsWithContext creates a new CreateBlueprintMappingUsingPOST1Params object
// with the default values initialized, and the ability to set a context for a request
func NewCreateBlueprintMappingUsingPOST1ParamsWithContext(ctx context.Context) *CreateBlueprintMappingUsingPOST1Params {
	var ()
	return &CreateBlueprintMappingUsingPOST1Params{

		Context: ctx,
	}
}

// NewCreateBlueprintMappingUsingPOST1ParamsWithHTTPClient creates a new CreateBlueprintMappingUsingPOST1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateBlueprintMappingUsingPOST1ParamsWithHTTPClient(client *http.Client) *CreateBlueprintMappingUsingPOST1Params {
	var ()
	return &CreateBlueprintMappingUsingPOST1Params{
		HTTPClient: client,
	}
}

/*CreateBlueprintMappingUsingPOST1Params contains all the parameters to send to the API endpoint
for the create blueprint mapping using p o s t 1 operation typically these are written to a http.Request
*/
type CreateBlueprintMappingUsingPOST1Params struct {

	/*ConfigurationSourceReference
	  configurationSourceReference

	*/
	ConfigurationSourceReference *models.TerraformConfigurationSourceReference

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create blueprint mapping using p o s t 1 params
func (o *CreateBlueprintMappingUsingPOST1Params) WithTimeout(timeout time.Duration) *CreateBlueprintMappingUsingPOST1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create blueprint mapping using p o s t 1 params
func (o *CreateBlueprintMappingUsingPOST1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create blueprint mapping using p o s t 1 params
func (o *CreateBlueprintMappingUsingPOST1Params) WithContext(ctx context.Context) *CreateBlueprintMappingUsingPOST1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create blueprint mapping using p o s t 1 params
func (o *CreateBlueprintMappingUsingPOST1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create blueprint mapping using p o s t 1 params
func (o *CreateBlueprintMappingUsingPOST1Params) WithHTTPClient(client *http.Client) *CreateBlueprintMappingUsingPOST1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create blueprint mapping using p o s t 1 params
func (o *CreateBlueprintMappingUsingPOST1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigurationSourceReference adds the configurationSourceReference to the create blueprint mapping using p o s t 1 params
func (o *CreateBlueprintMappingUsingPOST1Params) WithConfigurationSourceReference(configurationSourceReference *models.TerraformConfigurationSourceReference) *CreateBlueprintMappingUsingPOST1Params {
	o.SetConfigurationSourceReference(configurationSourceReference)
	return o
}

// SetConfigurationSourceReference adds the configurationSourceReference to the create blueprint mapping using p o s t 1 params
func (o *CreateBlueprintMappingUsingPOST1Params) SetConfigurationSourceReference(configurationSourceReference *models.TerraformConfigurationSourceReference) {
	o.ConfigurationSourceReference = configurationSourceReference
}

// WriteToRequest writes these params to a swagger request
func (o *CreateBlueprintMappingUsingPOST1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ConfigurationSourceReference != nil {
		if err := r.SetBodyParam(o.ConfigurationSourceReference); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
