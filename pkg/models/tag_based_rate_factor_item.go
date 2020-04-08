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

// TagBasedRateFactorItem TagBasedRateFactorItem
// swagger:model TagBasedRateFactorItem
type TagBasedRateFactorItem struct {

	// item name
	ItemName string `json:"itemName,omitempty"`

	// rate factors
	RateFactors []*RateFactorItem `json:"rateFactors"`
}

// Validate validates this tag based rate factor item
func (m *TagBasedRateFactorItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRateFactors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TagBasedRateFactorItem) validateRateFactors(formats strfmt.Registry) error {

	if swag.IsZero(m.RateFactors) { // not required
		return nil
	}

	for i := 0; i < len(m.RateFactors); i++ {
		if swag.IsZero(m.RateFactors[i]) { // not required
			continue
		}

		if m.RateFactors[i] != nil {
			if err := m.RateFactors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rateFactors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TagBasedRateFactorItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TagBasedRateFactorItem) UnmarshalBinary(b []byte) error {
	var res TagBasedRateFactorItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
