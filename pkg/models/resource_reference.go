// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ResourceReference ResourceReference
//
// swagger:model ResourceReference
type ResourceReference struct {

	// id
	ID string `json:"id,omitempty"`

	// link
	Link string `json:"link,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this resource reference
func (m *ResourceReference) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResourceReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceReference) UnmarshalBinary(b []byte) error {
	var res ResourceReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
