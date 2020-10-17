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

// MarketplaceFilter MarketplaceFilter
//
// swagger:model MarketplaceFilter
type MarketplaceFilter struct {

	// entries
	Entries *MarketplaceFilterEntries `json:"entries,omitempty"`

	// Filter Id
	ID string `json:"id,omitempty"`

	// Filter Display Name
	Name string `json:"name,omitempty"`

	// Filter Type
	// Enum: [COLLECTION DATE NUMBER]
	Type string `json:"type,omitempty"`

	// Filter Value Type
	// Enum: [RANGE SINGLE MULTIPLE]
	ValueType string `json:"valueType,omitempty"`
}

// Validate validates this marketplace filter
func (m *MarketplaceFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntries(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValueType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MarketplaceFilter) validateEntries(formats strfmt.Registry) error {

	if swag.IsZero(m.Entries) { // not required
		return nil
	}

	if m.Entries != nil {
		if err := m.Entries.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entries")
			}
			return err
		}
	}

	return nil
}

var marketplaceFilterTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["COLLECTION","DATE","NUMBER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		marketplaceFilterTypeTypePropEnum = append(marketplaceFilterTypeTypePropEnum, v)
	}
}

const (

	// MarketplaceFilterTypeCOLLECTION captures enum value "COLLECTION"
	MarketplaceFilterTypeCOLLECTION string = "COLLECTION"

	// MarketplaceFilterTypeDATE captures enum value "DATE"
	MarketplaceFilterTypeDATE string = "DATE"

	// MarketplaceFilterTypeNUMBER captures enum value "NUMBER"
	MarketplaceFilterTypeNUMBER string = "NUMBER"
)

// prop value enum
func (m *MarketplaceFilter) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, marketplaceFilterTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MarketplaceFilter) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

var marketplaceFilterTypeValueTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["RANGE","SINGLE","MULTIPLE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		marketplaceFilterTypeValueTypePropEnum = append(marketplaceFilterTypeValueTypePropEnum, v)
	}
}

const (

	// MarketplaceFilterValueTypeRANGE captures enum value "RANGE"
	MarketplaceFilterValueTypeRANGE string = "RANGE"

	// MarketplaceFilterValueTypeSINGLE captures enum value "SINGLE"
	MarketplaceFilterValueTypeSINGLE string = "SINGLE"

	// MarketplaceFilterValueTypeMULTIPLE captures enum value "MULTIPLE"
	MarketplaceFilterValueTypeMULTIPLE string = "MULTIPLE"
)

// prop value enum
func (m *MarketplaceFilter) validateValueTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, marketplaceFilterTypeValueTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MarketplaceFilter) validateValueType(formats strfmt.Registry) error {

	if swag.IsZero(m.ValueType) { // not required
		return nil
	}

	// value enum
	if err := m.validateValueTypeEnum("valueType", "body", m.ValueType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MarketplaceFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MarketplaceFilter) UnmarshalBinary(b []byte) error {
	var res MarketplaceFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
