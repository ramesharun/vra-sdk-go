// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MachineBootConfigSettings Machine boot config settings that will define how the provisioning will handle the boot config script execution.
//
// swagger:model MachineBootConfigSettings
type MachineBootConfigSettings struct {

	// In case a timeout occurs whether the provisioning process should fail or continue.
	PhoneHomeFailOnTimeout bool `json:"phoneHomeFailOnTimeout,omitempty"`

	// A phone_home module will be added to the Cloud Config and the provisioning will wait on a callback prior proceeding
	PhoneHomeShouldWait bool `json:"phoneHomeShouldWait,omitempty"`

	// The period of time to wait for the phone_home module callback to occur
	PhoneHomeTimeoutSeconds int32 `json:"phoneHomeTimeoutSeconds,omitempty"`
}

// Validate validates this machine boot config settings
func (m *MachineBootConfigSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MachineBootConfigSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MachineBootConfigSettings) UnmarshalBinary(b []byte) error {
	var res MachineBootConfigSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
