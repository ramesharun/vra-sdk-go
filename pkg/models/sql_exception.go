// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SQLException SQL exception
//
// swagger:model SQLException
type SQLException struct {

	// cause
	Cause *Throwable `json:"cause,omitempty"`

	// error code
	ErrorCode int32 `json:"errorCode,omitempty"`

	// localized message
	LocalizedMessage string `json:"localizedMessage,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// next exception
	NextException *SQLException `json:"nextException,omitempty"`

	// sqlstate
	Sqlstate string `json:"sqlstate,omitempty"`

	// stack trace
	StackTrace []*StackTraceElement `json:"stackTrace"`

	// suppressed
	Suppressed []*Throwable `json:"suppressed"`
}

// Validate validates this SQL exception
func (m *SQLException) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCause(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNextException(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStackTrace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSuppressed(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SQLException) validateCause(formats strfmt.Registry) error {

	if swag.IsZero(m.Cause) { // not required
		return nil
	}

	if m.Cause != nil {
		if err := m.Cause.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cause")
			}
			return err
		}
	}

	return nil
}

func (m *SQLException) validateNextException(formats strfmt.Registry) error {

	if swag.IsZero(m.NextException) { // not required
		return nil
	}

	if m.NextException != nil {
		if err := m.NextException.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nextException")
			}
			return err
		}
	}

	return nil
}

func (m *SQLException) validateStackTrace(formats strfmt.Registry) error {

	if swag.IsZero(m.StackTrace) { // not required
		return nil
	}

	for i := 0; i < len(m.StackTrace); i++ {
		if swag.IsZero(m.StackTrace[i]) { // not required
			continue
		}

		if m.StackTrace[i] != nil {
			if err := m.StackTrace[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("stackTrace" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SQLException) validateSuppressed(formats strfmt.Registry) error {

	if swag.IsZero(m.Suppressed) { // not required
		return nil
	}

	for i := 0; i < len(m.Suppressed); i++ {
		if swag.IsZero(m.Suppressed[i]) { // not required
			continue
		}

		if m.Suppressed[i] != nil {
			if err := m.Suppressed[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("suppressed" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SQLException) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SQLException) UnmarshalBinary(b []byte) error {
	var res SQLException
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
