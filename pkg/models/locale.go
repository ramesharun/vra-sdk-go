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

// Locale locale
//
// swagger:model Locale
type Locale struct {

	// country
	Country string `json:"country,omitempty"`

	// display country
	DisplayCountry string `json:"displayCountry,omitempty"`

	// display language
	DisplayLanguage string `json:"displayLanguage,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// display script
	DisplayScript string `json:"displayScript,omitempty"`

	// display variant
	DisplayVariant string `json:"displayVariant,omitempty"`

	// extension keys
	// Unique: true
	ExtensionKeys []string `json:"extensionKeys"`

	// iso3 country
	Iso3Country string `json:"iso3Country,omitempty"`

	// iso3 language
	Iso3Language string `json:"iso3Language,omitempty"`

	// language
	Language string `json:"language,omitempty"`

	// script
	Script string `json:"script,omitempty"`

	// unicode locale attributes
	// Unique: true
	UnicodeLocaleAttributes []string `json:"unicodeLocaleAttributes"`

	// unicode locale keys
	// Unique: true
	UnicodeLocaleKeys []string `json:"unicodeLocaleKeys"`

	// variant
	Variant string `json:"variant,omitempty"`
}

// Validate validates this locale
func (m *Locale) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExtensionKeys(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnicodeLocaleAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnicodeLocaleKeys(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Locale) validateExtensionKeys(formats strfmt.Registry) error {

	if swag.IsZero(m.ExtensionKeys) { // not required
		return nil
	}

	if err := validate.UniqueItems("extensionKeys", "body", m.ExtensionKeys); err != nil {
		return err
	}

	return nil
}

func (m *Locale) validateUnicodeLocaleAttributes(formats strfmt.Registry) error {

	if swag.IsZero(m.UnicodeLocaleAttributes) { // not required
		return nil
	}

	if err := validate.UniqueItems("unicodeLocaleAttributes", "body", m.UnicodeLocaleAttributes); err != nil {
		return err
	}

	return nil
}

func (m *Locale) validateUnicodeLocaleKeys(formats strfmt.Registry) error {

	if swag.IsZero(m.UnicodeLocaleKeys) { // not required
		return nil
	}

	if err := validate.UniqueItems("unicodeLocaleKeys", "body", m.UnicodeLocaleKeys); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Locale) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Locale) UnmarshalBinary(b []byte) error {
	var res Locale
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
