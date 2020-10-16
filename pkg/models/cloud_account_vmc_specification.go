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

// CloudAccountVmcSpecification Specification for a VMC cloud account.<br><br>A cloud account identifies a cloud account type and an account-specific deployment region or data center where the associated cloud account resources are hosted.
//
// swagger:model CloudAccountVmcSpecification
type CloudAccountVmcSpecification struct {

	// Accept self signed certificate when connecting to vSphere
	AcceptSelfSignedCertificate bool `json:"acceptSelfSignedCertificate,omitempty"`

	// VMC API access key
	APIKey string `json:"apiKey,omitempty"`

	// Create default cloud zones for the enabled regions.
	CreateDefaultZones bool `json:"createDefaultZones,omitempty"`

	// The host name of the CSP service.
	CspHostName string `json:"cspHostName,omitempty"`

	// Identifier of a data collector vm deployed in the on premise infrastructure. Refer to the data-collector API to create or list data collectors
	DcID string `json:"dcId,omitempty"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// Enter the IP address or FQDN of the vCenter Server in the specified SDDC. The cloud proxy belongs on this vCenter.
	// Required: true
	HostName *string `json:"hostName"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Required: true
	Name *string `json:"name"`

	// The IP address of the NSX Manager server in the specified SDDC / FQDN.
	// Required: true
	NsxHostName *string `json:"nsxHostName"`

	// Password for the user used to authenticate with the cloud Account
	// Required: true
	Password *string `json:"password"`

	// A set of Region names to enable provisioning on.Refer to /iaas/cloud-accounts/region-enumeration.
	// Required: true
	RegionIds []string `json:"regionIds"`

	// Identifier of the on-premise SDDC to be used by this cloud account. Note that NSX-V SDDCs are not supported.
	SddcID string `json:"sddcId,omitempty"`

	// A set of tag keys and optional values to set on the Cloud Account.Cloud account capability tags may enable different features.
	Tags []*Tag `json:"tags"`

	// vCenter user name for the specified SDDC.The specified user requires CloudAdmin credentials. The user does not require CloudGlobalAdmin credentials.
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this cloud account vmc specification
func (m *CloudAccountVmcSpecification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNsxHostName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegionIds(formats); err != nil {
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

func (m *CloudAccountVmcSpecification) validateHostName(formats strfmt.Registry) error {

	if err := validate.Required("hostName", "body", m.HostName); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountVmcSpecification) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountVmcSpecification) validateNsxHostName(formats strfmt.Registry) error {

	if err := validate.Required("nsxHostName", "body", m.NsxHostName); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountVmcSpecification) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountVmcSpecification) validateRegionIds(formats strfmt.Registry) error {

	if err := validate.Required("regionIds", "body", m.RegionIds); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountVmcSpecification) validateTags(formats strfmt.Registry) error {

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

func (m *CloudAccountVmcSpecification) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudAccountVmcSpecification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudAccountVmcSpecification) UnmarshalBinary(b []byte) error {
	var res CloudAccountVmcSpecification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
