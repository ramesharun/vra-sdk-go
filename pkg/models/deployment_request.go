// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DeploymentRequest DeploymentRequest
//
// Represents deployment requests.
//
// swagger:model DeploymentRequest
type DeploymentRequest struct {

	// Identifier of the requested action
	ActionID string `json:"actionId,omitempty"`

	// Time at which the request was approved.
	// Format: date-time
	ApprovedAt strfmt.DateTime `json:"approvedAt,omitempty"`

	// Identifier of the requested blueprint in the form 'UUID:version'
	BlueprintID string `json:"blueprintId,omitempty"`

	// Indicates whether request can be canceled or not.
	Cancelable bool `json:"cancelable,omitempty"`

	// Identifier of the requested catalog item in the form 'UUID:version'
	CatalogItemID string `json:"catalogItemId,omitempty"`

	// Time at which the request completed.
	// Format: date-time
	CompletedAt strfmt.DateTime `json:"completedAt,omitempty"`

	// The number of tasks completed while fulfilling this request.
	CompletedTasks int32 `json:"completedTasks,omitempty"`

	// Creation time (e.g. date format '2019-07-13T23:16:49.310Z').
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Deployment id to which the event applies to
	// Format: uuid
	DeploymentID strfmt.UUID `json:"deploymentId,omitempty"`

	// Longer user-friendly details of the event.
	Details string `json:"details,omitempty"`

	// Indicates whether request is in dismissed state.
	Dismissed bool `json:"dismissed,omitempty"`

	// Event identifier
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Time at which the request was initialized.
	// Format: date-time
	InitializedAt strfmt.DateTime `json:"initializedAt,omitempty"`

	// Request inputs
	Inputs interface{} `json:"inputs,omitempty"`

	// Short user-friendly label of the event (e.g. 'shuting down myVM')
	Name string `json:"name,omitempty"`

	// Parent event/request identifier
	// Format: uuid
	ParentID strfmt.UUID `json:"parentId,omitempty"`

	// User that initiated the request
	RequestedBy string `json:"requestedBy,omitempty"`

	// Optional resource name to which the event applies to
	ResourceName string `json:"resourceName,omitempty"`

	// Optional resource type to which the event applies to
	ResourceType string `json:"resourceType,omitempty"`

	// Request overall execution status.
	// Enum: [CREATED PENDING INITIALIZATION APPROVAL_PENDING INPROGRESS COMPLETION APPROVAL_REJECTED ABORTED SUCCESSFUL FAILED]
	Status string `json:"status,omitempty"`

	// The total number of tasks need to be completed to fulfil this request.
	TotalTasks int32 `json:"totalTasks,omitempty"`

	// Last update time (e.g. date format '2019-07-13T23:16:49.310Z').
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`
}

// Validate validates this deployment request
func (m *DeploymentRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApprovedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompletedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeploymentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitializedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeploymentRequest) validateApprovedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.ApprovedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("approvedAt", "body", "date-time", m.ApprovedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeploymentRequest) validateCompletedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CompletedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("completedAt", "body", "date-time", m.CompletedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeploymentRequest) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeploymentRequest) validateDeploymentID(formats strfmt.Registry) error {

	if swag.IsZero(m.DeploymentID) { // not required
		return nil
	}

	if err := validate.FormatOf("deploymentId", "body", "uuid", m.DeploymentID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeploymentRequest) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeploymentRequest) validateInitializedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.InitializedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("initializedAt", "body", "date-time", m.InitializedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeploymentRequest) validateParentID(formats strfmt.Registry) error {

	if swag.IsZero(m.ParentID) { // not required
		return nil
	}

	if err := validate.FormatOf("parentId", "body", "uuid", m.ParentID.String(), formats); err != nil {
		return err
	}

	return nil
}

var deploymentRequestTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CREATED","PENDING","INITIALIZATION","APPROVAL_PENDING","INPROGRESS","COMPLETION","APPROVAL_REJECTED","ABORTED","SUCCESSFUL","FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deploymentRequestTypeStatusPropEnum = append(deploymentRequestTypeStatusPropEnum, v)
	}
}

const (

	// DeploymentRequestStatusCREATED captures enum value "CREATED"
	DeploymentRequestStatusCREATED string = "CREATED"

	// DeploymentRequestStatusPENDING captures enum value "PENDING"
	DeploymentRequestStatusPENDING string = "PENDING"

	// DeploymentRequestStatusINITIALIZATION captures enum value "INITIALIZATION"
	DeploymentRequestStatusINITIALIZATION string = "INITIALIZATION"

	// DeploymentRequestStatusAPPROVALPENDING captures enum value "APPROVAL_PENDING"
	DeploymentRequestStatusAPPROVALPENDING string = "APPROVAL_PENDING"

	// DeploymentRequestStatusINPROGRESS captures enum value "INPROGRESS"
	DeploymentRequestStatusINPROGRESS string = "INPROGRESS"

	// DeploymentRequestStatusCOMPLETION captures enum value "COMPLETION"
	DeploymentRequestStatusCOMPLETION string = "COMPLETION"

	// DeploymentRequestStatusAPPROVALREJECTED captures enum value "APPROVAL_REJECTED"
	DeploymentRequestStatusAPPROVALREJECTED string = "APPROVAL_REJECTED"

	// DeploymentRequestStatusABORTED captures enum value "ABORTED"
	DeploymentRequestStatusABORTED string = "ABORTED"

	// DeploymentRequestStatusSUCCESSFUL captures enum value "SUCCESSFUL"
	DeploymentRequestStatusSUCCESSFUL string = "SUCCESSFUL"

	// DeploymentRequestStatusFAILED captures enum value "FAILED"
	DeploymentRequestStatusFAILED string = "FAILED"
)

// prop value enum
func (m *DeploymentRequest) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, deploymentRequestTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DeploymentRequest) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *DeploymentRequest) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeploymentRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeploymentRequest) UnmarshalBinary(b []byte) error {
	var res DeploymentRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
