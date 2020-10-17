// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MeteringItem MeteringItem
//
// swagger:model MeteringItem
type MeteringItem struct {

	// item name
	ItemName string `json:"itemName,omitempty"`

	// metering
	Metering *Metering `json:"metering,omitempty"`
}

// Validate validates this metering item
func (m *MeteringItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetering(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MeteringItem) validateMetering(formats strfmt.Registry) error {

	if swag.IsZero(m.Metering) { // not required
		return nil
	}

	if m.Metering != nil {
		if err := m.Metering.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metering")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MeteringItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MeteringItem) UnmarshalBinary(b []byte) error {
	var res MeteringItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
