// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TimeBin time bin
//
// swagger:model TimeBin
type TimeBin struct {

	// avg
	Avg float64 `json:"avg,omitempty"`

	// count
	Count float64 `json:"count,omitempty"`

	// latest
	Latest float64 `json:"latest,omitempty"`

	// max
	Max float64 `json:"max,omitempty"`

	// min
	Min float64 `json:"min,omitempty"`

	// sum
	Sum float64 `json:"sum,omitempty"`

	// var
	Var float64 `json:"var,omitempty"`
}

// Validate validates this time bin
func (m *TimeBin) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TimeBin) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TimeBin) UnmarshalBinary(b []byte) error {
	var res TimeBin
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
