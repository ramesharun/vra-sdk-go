// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EventLog EventLog
//
// Represents a event log.
//
// swagger:model EventLog
type EventLog struct {

	// Indicates whether this is the last log entry.
	EOF bool `json:"eof,omitempty"`

	// Event Log identifier
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Message of the event log
	// Required: true
	Message *string `json:"message"`

	// Row number of the Event Log.
	Rownum int32 `json:"rownum,omitempty"`

	// Timestamp of the Event log (e.g. date format '2019-07-13T23:16:49.310Z').
	// Required: true
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp"`
}

// Validate validates this event log
func (m *EventLog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EventLog) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EventLog) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

func (m *EventLog) validateTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EventLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventLog) UnmarshalBinary(b []byte) error {
	var res EventLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
