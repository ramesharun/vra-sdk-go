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

// Metering Metering
//
// swagger:model Metering
type Metering struct {

	// base rate
	BaseRate float64 `json:"baseRate,omitempty"`

	// charge based on
	// Enum: [USAGE]
	ChargeBasedOn string `json:"chargeBasedOn,omitempty"`

	// charge on power state
	// Enum: [ALWAYS ONLY_WHEN_POWERED_ON POWERED_ON_AT_LEAST_ONCE]
	ChargeOnPowerState string `json:"chargeOnPowerState,omitempty"`

	// charge period
	// Enum: [HOURLY DAILY WEEKLY MONTHLY]
	ChargePeriod string `json:"chargePeriod,omitempty"`

	// fixed price
	FixedPrice float64 `json:"fixedPrice,omitempty"`

	// unit
	Unit string `json:"unit,omitempty"`
}

// Validate validates this metering
func (m *Metering) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChargeBasedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChargeOnPowerState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChargePeriod(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var meteringTypeChargeBasedOnPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["USAGE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		meteringTypeChargeBasedOnPropEnum = append(meteringTypeChargeBasedOnPropEnum, v)
	}
}

const (

	// MeteringChargeBasedOnUSAGE captures enum value "USAGE"
	MeteringChargeBasedOnUSAGE string = "USAGE"
)

// prop value enum
func (m *Metering) validateChargeBasedOnEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, meteringTypeChargeBasedOnPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Metering) validateChargeBasedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ChargeBasedOn) { // not required
		return nil
	}

	// value enum
	if err := m.validateChargeBasedOnEnum("chargeBasedOn", "body", m.ChargeBasedOn); err != nil {
		return err
	}

	return nil
}

var meteringTypeChargeOnPowerStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ALWAYS","ONLY_WHEN_POWERED_ON","POWERED_ON_AT_LEAST_ONCE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		meteringTypeChargeOnPowerStatePropEnum = append(meteringTypeChargeOnPowerStatePropEnum, v)
	}
}

const (

	// MeteringChargeOnPowerStateALWAYS captures enum value "ALWAYS"
	MeteringChargeOnPowerStateALWAYS string = "ALWAYS"

	// MeteringChargeOnPowerStateONLYWHENPOWEREDON captures enum value "ONLY_WHEN_POWERED_ON"
	MeteringChargeOnPowerStateONLYWHENPOWEREDON string = "ONLY_WHEN_POWERED_ON"

	// MeteringChargeOnPowerStatePOWEREDONATLEASTONCE captures enum value "POWERED_ON_AT_LEAST_ONCE"
	MeteringChargeOnPowerStatePOWEREDONATLEASTONCE string = "POWERED_ON_AT_LEAST_ONCE"
)

// prop value enum
func (m *Metering) validateChargeOnPowerStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, meteringTypeChargeOnPowerStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Metering) validateChargeOnPowerState(formats strfmt.Registry) error {

	if swag.IsZero(m.ChargeOnPowerState) { // not required
		return nil
	}

	// value enum
	if err := m.validateChargeOnPowerStateEnum("chargeOnPowerState", "body", m.ChargeOnPowerState); err != nil {
		return err
	}

	return nil
}

var meteringTypeChargePeriodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["HOURLY","DAILY","WEEKLY","MONTHLY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		meteringTypeChargePeriodPropEnum = append(meteringTypeChargePeriodPropEnum, v)
	}
}

const (

	// MeteringChargePeriodHOURLY captures enum value "HOURLY"
	MeteringChargePeriodHOURLY string = "HOURLY"

	// MeteringChargePeriodDAILY captures enum value "DAILY"
	MeteringChargePeriodDAILY string = "DAILY"

	// MeteringChargePeriodWEEKLY captures enum value "WEEKLY"
	MeteringChargePeriodWEEKLY string = "WEEKLY"

	// MeteringChargePeriodMONTHLY captures enum value "MONTHLY"
	MeteringChargePeriodMONTHLY string = "MONTHLY"
)

// prop value enum
func (m *Metering) validateChargePeriodEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, meteringTypeChargePeriodPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Metering) validateChargePeriod(formats strfmt.Registry) error {

	if swag.IsZero(m.ChargePeriod) { // not required
		return nil
	}

	// value enum
	if err := m.validateChargePeriodEnum("chargePeriod", "body", m.ChargePeriod); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Metering) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Metering) UnmarshalBinary(b []byte) error {
	var res Metering
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
