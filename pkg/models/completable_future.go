// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CompletableFuture completable future
//
// swagger:model CompletableFuture
type CompletableFuture struct {

	// cancelled
	Cancelled bool `json:"cancelled,omitempty"`

	// completed exceptionally
	CompletedExceptionally bool `json:"completedExceptionally,omitempty"`

	// done
	Done bool `json:"done,omitempty"`

	// number of dependents
	NumberOfDependents int32 `json:"numberOfDependents,omitempty"`
}

// Validate validates this completable future
func (m *CompletableFuture) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CompletableFuture) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CompletableFuture) UnmarshalBinary(b []byte) error {
	var res CompletableFuture
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
