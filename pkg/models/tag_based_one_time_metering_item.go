// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TagBasedOneTimeMeteringItem TagBasedOneTimeMeteringItem
// swagger:model TagBasedOneTimeMeteringItem
type TagBasedOneTimeMeteringItem struct {

	// item name
	ItemName string `json:"itemName,omitempty"`

	// one time meterings
	OneTimeMeterings []*TagBasedOneTimeMetering `json:"oneTimeMeterings"`
}

// Validate validates this tag based one time metering item
func (m *TagBasedOneTimeMeteringItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOneTimeMeterings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TagBasedOneTimeMeteringItem) validateOneTimeMeterings(formats strfmt.Registry) error {

	if swag.IsZero(m.OneTimeMeterings) { // not required
		return nil
	}

	for i := 0; i < len(m.OneTimeMeterings); i++ {
		if swag.IsZero(m.OneTimeMeterings[i]) { // not required
			continue
		}

		if m.OneTimeMeterings[i] != nil {
			if err := m.OneTimeMeterings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("oneTimeMeterings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TagBasedOneTimeMeteringItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TagBasedOneTimeMeteringItem) UnmarshalBinary(b []byte) error {
	var res TagBasedOneTimeMeteringItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
