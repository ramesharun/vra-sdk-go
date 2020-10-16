// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FilterSchema FilterSchema
//
// swagger:model FilterSchema
type FilterSchema struct {

	// filters
	Filters []Filter `json:"filters"`
}

// Validate validates this filter schema
func (m *FilterSchema) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FilterSchema) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FilterSchema) UnmarshalBinary(b []byte) error {
	var res FilterSchema
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
