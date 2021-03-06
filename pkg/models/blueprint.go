// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Blueprint Blueprint
//
// swagger:model Blueprint
type Blueprint struct {

	// Blueprint YAML content
	Content string `json:"content,omitempty"`

	// Content source id
	// Read Only: true
	ContentSourceID string `json:"contentSourceId,omitempty"`

	// Content source path
	// Read Only: true
	ContentSourcePath string `json:"contentSourcePath,omitempty"`

	// Content source last sync time
	// Read Only: true
	// Format: date-time
	ContentSourceSyncAt strfmt.DateTime `json:"contentSourceSyncAt,omitempty"`

	// Content source last sync messages
	// Read Only: true
	ContentSourceSyncMessages []string `json:"contentSourceSyncMessages"`

	// Content source last sync status
	// Read Only: true
	// Enum: [SUCCESSFUL FAILED]
	ContentSourceSyncStatus string `json:"contentSourceSyncStatus,omitempty"`

	// Content source type
	// Read Only: true
	ContentSourceType string `json:"contentSourceType,omitempty"`

	// Created time
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Created by
	// Read Only: true
	CreatedBy string `json:"createdBy,omitempty"`

	// Blueprint description
	Description string `json:"description,omitempty"`

	// Object ID
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Blueprint name
	Name string `json:"name,omitempty"`

	// Org ID
	// Read Only: true
	OrgID string `json:"orgId,omitempty"`

	// Project ID
	ProjectID string `json:"projectId,omitempty"`

	// Project Name
	// Read Only: true
	ProjectName string `json:"projectName,omitempty"`

	// Flag to indicate blueprint can be requested from any project in org
	RequestScopeOrg bool `json:"requestScopeOrg,omitempty"`

	// Blueprint self link
	// Read Only: true
	SelfLink string `json:"selfLink,omitempty"`

	// Blueprint status
	// Read Only: true
	// Enum: [DRAFT VERSIONED RELEASED]
	Status string `json:"status,omitempty"`

	// Total released versions
	// Read Only: true
	TotalReleasedVersions int32 `json:"totalReleasedVersions,omitempty"`

	// Total versions
	// Read Only: true
	TotalVersions int32 `json:"totalVersions,omitempty"`

	// Updated time
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`

	// Updated by
	// Read Only: true
	UpdatedBy string `json:"updatedBy,omitempty"`

	// Validation result on update
	// Read Only: true
	Valid *bool `json:"valid,omitempty"`

	// Validation messages
	// Read Only: true
	ValidationMessages []*BlueprintValidationMessage `json:"validationMessages"`
}

// Validate validates this blueprint
func (m *Blueprint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContentSourceSyncAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContentSourceSyncStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidationMessages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Blueprint) validateContentSourceSyncAt(formats strfmt.Registry) error {

	if swag.IsZero(m.ContentSourceSyncAt) { // not required
		return nil
	}

	if err := validate.FormatOf("contentSourceSyncAt", "body", "date-time", m.ContentSourceSyncAt.String(), formats); err != nil {
		return err
	}

	return nil
}

var blueprintTypeContentSourceSyncStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SUCCESSFUL","FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		blueprintTypeContentSourceSyncStatusPropEnum = append(blueprintTypeContentSourceSyncStatusPropEnum, v)
	}
}

const (

	// BlueprintContentSourceSyncStatusSUCCESSFUL captures enum value "SUCCESSFUL"
	BlueprintContentSourceSyncStatusSUCCESSFUL string = "SUCCESSFUL"

	// BlueprintContentSourceSyncStatusFAILED captures enum value "FAILED"
	BlueprintContentSourceSyncStatusFAILED string = "FAILED"
)

// prop value enum
func (m *Blueprint) validateContentSourceSyncStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, blueprintTypeContentSourceSyncStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Blueprint) validateContentSourceSyncStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.ContentSourceSyncStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateContentSourceSyncStatusEnum("contentSourceSyncStatus", "body", m.ContentSourceSyncStatus); err != nil {
		return err
	}

	return nil
}

func (m *Blueprint) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

var blueprintTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DRAFT","VERSIONED","RELEASED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		blueprintTypeStatusPropEnum = append(blueprintTypeStatusPropEnum, v)
	}
}

const (

	// BlueprintStatusDRAFT captures enum value "DRAFT"
	BlueprintStatusDRAFT string = "DRAFT"

	// BlueprintStatusVERSIONED captures enum value "VERSIONED"
	BlueprintStatusVERSIONED string = "VERSIONED"

	// BlueprintStatusRELEASED captures enum value "RELEASED"
	BlueprintStatusRELEASED string = "RELEASED"
)

// prop value enum
func (m *Blueprint) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, blueprintTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Blueprint) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *Blueprint) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Blueprint) validateValidationMessages(formats strfmt.Registry) error {

	if swag.IsZero(m.ValidationMessages) { // not required
		return nil
	}

	for i := 0; i < len(m.ValidationMessages); i++ {
		if swag.IsZero(m.ValidationMessages[i]) { // not required
			continue
		}

		if m.ValidationMessages[i] != nil {
			if err := m.ValidationMessages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("validationMessages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Blueprint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Blueprint) UnmarshalBinary(b []byte) error {
	var res Blueprint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
