// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CloudAccountNsxVSpecification Specification for an NSX-v cloud account.<br><br>A cloud account identifies a cloud account type and an account-specific deployment region or data center where the associated cloud account resources are hosted.
//
// swagger:model CloudAccountNsxVSpecification
type CloudAccountNsxVSpecification struct {

	// Accept self signed certificate when connecting.
	AcceptSelfSignedCertificate bool `json:"acceptSelfSignedCertificate,omitempty"`

	// vSphere cloud account associated with this NSX-V cloud account. NSX-V cloud account can be associated with a single vSphere cloud account.
	AssociatedCloudAccountIds []string `json:"associatedCloudAccountIds"`

	// Identifier of a data collector vm deployed in the on premise infrastructure. Refer to the data-collector API to create or list data collectors
	// Required: true
	Dcid *string `json:"dcid"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// Host name for the NSX-v endpoint
	// Required: true
	HostName *string `json:"hostName"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Required: true
	Name *string `json:"name"`

	// Password for the user used to authenticate with the cloud Account
	// Required: true
	Password *string `json:"password"`

	// A set of tag keys and optional values to set on the Cloud Account
	Tags []*Tag `json:"tags"`

	// Username to authenticate with the cloud account
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this cloud account nsx v specification
func (m *CloudAccountNsxVSpecification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDcid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudAccountNsxVSpecification) validateDcid(formats strfmt.Registry) error {

	if err := validate.Required("dcid", "body", m.Dcid); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountNsxVSpecification) validateHostName(formats strfmt.Registry) error {

	if err := validate.Required("hostName", "body", m.HostName); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountNsxVSpecification) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountNsxVSpecification) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountNsxVSpecification) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CloudAccountNsxVSpecification) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudAccountNsxVSpecification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudAccountNsxVSpecification) UnmarshalBinary(b []byte) error {
	var res CloudAccountNsxVSpecification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
