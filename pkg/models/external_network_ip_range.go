// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ExternalNetworkIPRange State object representing an external IP address range for a Fabric Network.<br>**HATEOAS** links:<br>**region** - Region - Region for the network.<br>**self** - NetworkIPRange - Self link to this IP address range
//
// swagger:model ExternalNetworkIPRange
type ExternalNetworkIPRange struct {

	// HATEOAS of the entity
	// Required: true
	Links map[string]Href `json:"_links"`

	// Address space that the range belongs to
	AddressSpaceID string `json:"addressSpaceId,omitempty"`

	// Date when the entity was created. The date is in ISO 8601 and UTC.
	CreatedAt string `json:"createdAt,omitempty"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// DNS domain search (in order)
	DNSSearchDomains []string `json:"dnsSearchDomains"`

	// DNS IP addresses of the range.
	DNSServerAddresses []string `json:"dnsServerAddresses"`

	// DNS domain of the range.
	Domain string `json:"domain,omitempty"`

	// End IP address of the range.
	// Required: true
	EndIPAddress *string `json:"endIPAddress"`

	// External entity Id on the provider side.
	ExternalID string `json:"externalId,omitempty"`

	// The gateway address of the range
	GatewayAddress string `json:"gatewayAddress,omitempty"`

	// The id of this resource instance
	// Required: true
	ID *string `json:"id"`

	// IP address version: IPv4 or IPv6. Default: IPv4.
	// Enum: [IPv4 IPv6]
	IPVersion string `json:"ipVersion,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	Name string `json:"name,omitempty"`

	// The id of the organization this entity belongs to.
	OrgID string `json:"orgId,omitempty"`

	// This field is deprecated. Use orgId instead. The id of the organization this entity belongs to.
	OrganizationID string `json:"organizationId,omitempty"`

	// Email of the user that owns the entity.
	Owner string `json:"owner,omitempty"`

	// Start IP address of the range.
	// Required: true
	StartIPAddress *string `json:"startIPAddress"`

	// Subnet prefix length (synonymous with "netmask")
	// Required: true
	SubnetPrefixLength *int32 `json:"subnetPrefixLength"`

	// A set of tag keys and optional values that were set on this resource instance.
	Tags []*Tag `json:"tags"`

	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt string `json:"updatedAt,omitempty"`
}

// Validate validates this external network IP range
func (m *ExternalNetworkIPRange) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndIPAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartIPAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnetPrefixLength(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExternalNetworkIPRange) validateLinks(formats strfmt.Registry) error {

	for k := range m.Links {

		if err := validate.Required("_links"+"."+k, "body", m.Links[k]); err != nil {
			return err
		}
		if val, ok := m.Links[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ExternalNetworkIPRange) validateEndIPAddress(formats strfmt.Registry) error {

	if err := validate.Required("endIPAddress", "body", m.EndIPAddress); err != nil {
		return err
	}

	return nil
}

func (m *ExternalNetworkIPRange) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

var externalNetworkIpRangeTypeIPVersionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["IPv4","IPv6"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		externalNetworkIpRangeTypeIPVersionPropEnum = append(externalNetworkIpRangeTypeIPVersionPropEnum, v)
	}
}

const (

	// ExternalNetworkIPRangeIPVersionIPV4 captures enum value "IPv4"
	ExternalNetworkIPRangeIPVersionIPV4 string = "IPv4"

	// ExternalNetworkIPRangeIPVersionIPV6 captures enum value "IPv6"
	ExternalNetworkIPRangeIPVersionIPV6 string = "IPv6"
)

// prop value enum
func (m *ExternalNetworkIPRange) validateIPVersionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, externalNetworkIpRangeTypeIPVersionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ExternalNetworkIPRange) validateIPVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.IPVersion) { // not required
		return nil
	}

	// value enum
	if err := m.validateIPVersionEnum("ipVersion", "body", m.IPVersion); err != nil {
		return err
	}

	return nil
}

func (m *ExternalNetworkIPRange) validateStartIPAddress(formats strfmt.Registry) error {

	if err := validate.Required("startIPAddress", "body", m.StartIPAddress); err != nil {
		return err
	}

	return nil
}

func (m *ExternalNetworkIPRange) validateSubnetPrefixLength(formats strfmt.Registry) error {

	if err := validate.Required("subnetPrefixLength", "body", m.SubnetPrefixLength); err != nil {
		return err
	}

	return nil
}

func (m *ExternalNetworkIPRange) validateTags(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *ExternalNetworkIPRange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExternalNetworkIPRange) UnmarshalBinary(b []byte) error {
	var res ExternalNetworkIPRange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
