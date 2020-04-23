// Code generated by go-swagger; DO NOT EDIT.

package marketplace_downloads

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

// NewDownloadHistoryUsingGETParams creates a new DownloadHistoryUsingGETParams object
// with the default values initialized.
func NewDownloadHistoryUsingGETParams() *DownloadHistoryUsingGETParams {
	var ()
	return &DownloadHistoryUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDownloadHistoryUsingGETParamsWithTimeout creates a new DownloadHistoryUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDownloadHistoryUsingGETParamsWithTimeout(timeout time.Duration) *DownloadHistoryUsingGETParams {
	var ()
	return &DownloadHistoryUsingGETParams{

		timeout: timeout,
	}
}

// NewDownloadHistoryUsingGETParamsWithContext creates a new DownloadHistoryUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewDownloadHistoryUsingGETParamsWithContext(ctx context.Context) *DownloadHistoryUsingGETParams {
	var ()
	return &DownloadHistoryUsingGETParams{

		Context: ctx,
	}
}

// NewDownloadHistoryUsingGETParamsWithHTTPClient creates a new DownloadHistoryUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDownloadHistoryUsingGETParamsWithHTTPClient(client *http.Client) *DownloadHistoryUsingGETParams {
	var ()
	return &DownloadHistoryUsingGETParams{
		HTTPClient: client,
	}
}

/*DownloadHistoryUsingGETParams contains all the parameters to send to the API endpoint
for the download history using g e t operation typically these are written to a http.Request
*/
type DownloadHistoryUsingGETParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information, please refer to /content/api/about

	*/
	APIVersion *string
	/*ContentName
	  Search based on content name

	*/
	ContentName *string
	/*ContentType
	  Search based on content type

	*/
	ContentType *string
	/*Status
	  Search based on download status

	*/
	Status *string
	/*TargetIds
	  Search based on target Ids

	*/
	TargetIds []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the download history using get params
func (o *DownloadHistoryUsingGETParams) WithTimeout(timeout time.Duration) *DownloadHistoryUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the download history using get params
func (o *DownloadHistoryUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the download history using get params
func (o *DownloadHistoryUsingGETParams) WithContext(ctx context.Context) *DownloadHistoryUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the download history using get params
func (o *DownloadHistoryUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the download history using get params
func (o *DownloadHistoryUsingGETParams) WithHTTPClient(client *http.Client) *DownloadHistoryUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the download history using get params
func (o *DownloadHistoryUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the download history using get params
func (o *DownloadHistoryUsingGETParams) WithAPIVersion(aPIVersion *string) *DownloadHistoryUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the download history using get params
func (o *DownloadHistoryUsingGETParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithContentName adds the contentName to the download history using get params
func (o *DownloadHistoryUsingGETParams) WithContentName(contentName *string) *DownloadHistoryUsingGETParams {
	o.SetContentName(contentName)
	return o
}

// SetContentName adds the contentName to the download history using get params
func (o *DownloadHistoryUsingGETParams) SetContentName(contentName *string) {
	o.ContentName = contentName
}

// WithContentType adds the contentType to the download history using get params
func (o *DownloadHistoryUsingGETParams) WithContentType(contentType *string) *DownloadHistoryUsingGETParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the download history using get params
func (o *DownloadHistoryUsingGETParams) SetContentType(contentType *string) {
	o.ContentType = contentType
}

// WithStatus adds the status to the download history using get params
func (o *DownloadHistoryUsingGETParams) WithStatus(status *string) *DownloadHistoryUsingGETParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the download history using get params
func (o *DownloadHistoryUsingGETParams) SetStatus(status *string) {
	o.Status = status
}

// WithTargetIds adds the targetIds to the download history using get params
func (o *DownloadHistoryUsingGETParams) WithTargetIds(targetIds []string) *DownloadHistoryUsingGETParams {
	o.SetTargetIds(targetIds)
	return o
}

// SetTargetIds adds the targetIds to the download history using get params
func (o *DownloadHistoryUsingGETParams) SetTargetIds(targetIds []string) {
	o.TargetIds = targetIds
}

// WriteToRequest writes these params to a swagger request
func (o *DownloadHistoryUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ContentName != nil {

		// query param contentName
		var qrContentName string
		if o.ContentName != nil {
			qrContentName = *o.ContentName
		}
		qContentName := qrContentName
		if qContentName != "" {
			if err := r.SetQueryParam("contentName", qContentName); err != nil {
				return err
			}
		}

	}

	if o.ContentType != nil {

		// query param contentType
		var qrContentType string
		if o.ContentType != nil {
			qrContentType = *o.ContentType
		}
		qContentType := qrContentType
		if qContentType != "" {
			if err := r.SetQueryParam("contentType", qContentType); err != nil {
				return err
			}
		}

	}

	if o.Status != nil {

		// query param status
		var qrStatus string
		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {
			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}

	}

	valuesTargetIds := o.TargetIds

	joinedTargetIds := swag.JoinByFormat(valuesTargetIds, "multi")
	// query array param targetIds
	if err := r.SetQueryParam("targetIds", joinedTargetIds...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}