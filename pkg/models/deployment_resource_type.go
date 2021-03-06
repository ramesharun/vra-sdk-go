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

// DeploymentResourceType ResourceType
//
// swagger:model DeploymentResourceType
type DeploymentResourceType struct {

	// Optional. Account type to which the resource type belongs to. Example: AWS, Azure etc
	// Enum: [AWS Azure GCP vSphere vSphere-cloud Azure-EA NSX-V NSX-T NSX-P NSX-P-cloud vCloud Director VMC Puppet Ansible]
	AccountType string `json:"accountType,omitempty"`

	// Composable with other types or not
	Composable bool `json:"composable,omitempty"`

	// Time at which the resource type was created.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Name of the user who created the resource type.
	CreatedBy string `json:"createdBy,omitempty"`

	// Resource type description
	Description string `json:"description,omitempty"`

	// Resource display name
	DisplayName string `json:"displayName,omitempty"`

	// Resource type id
	ID string `json:"id,omitempty"`

	// Resource type name
	Name string `json:"name,omitempty"`

	// Org ID where resource type belongs
	OrgID string `json:"orgId,omitempty"`

	// Project ID where resource type belongs
	ProjectID string `json:"projectId,omitempty"`

	// Provider Id
	ProviderID string `json:"providerId,omitempty"`

	// Json schema that represents resource type, a simplified version of http://json-schema.org/latest/json-schema-validation.html#rfc.section.5
	Schema interface{} `json:"schema,omitempty"`

	// Time at which the resource type was updated.
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`

	// Name of the user who updated the resource type.
	UpdatedBy string `json:"updatedBy,omitempty"`
}

// Validate validates this deployment resource type
func (m *DeploymentResourceType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var deploymentResourceTypeTypeAccountTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AWS","Azure","GCP","vSphere","vSphere-cloud","Azure-EA","NSX-V","NSX-T","NSX-P","NSX-P-cloud","vCloud Director","VMC","Puppet","Ansible"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deploymentResourceTypeTypeAccountTypePropEnum = append(deploymentResourceTypeTypeAccountTypePropEnum, v)
	}
}

const (

	// DeploymentResourceTypeAccountTypeAWS captures enum value "AWS"
	DeploymentResourceTypeAccountTypeAWS string = "AWS"

	// DeploymentResourceTypeAccountTypeAzure captures enum value "Azure"
	DeploymentResourceTypeAccountTypeAzure string = "Azure"

	// DeploymentResourceTypeAccountTypeGCP captures enum value "GCP"
	DeploymentResourceTypeAccountTypeGCP string = "GCP"

	// DeploymentResourceTypeAccountTypeVSphere captures enum value "vSphere"
	DeploymentResourceTypeAccountTypeVSphere string = "vSphere"

	// DeploymentResourceTypeAccountTypeVSphereCloud captures enum value "vSphere-cloud"
	DeploymentResourceTypeAccountTypeVSphereCloud string = "vSphere-cloud"

	// DeploymentResourceTypeAccountTypeAzureEA captures enum value "Azure-EA"
	DeploymentResourceTypeAccountTypeAzureEA string = "Azure-EA"

	// DeploymentResourceTypeAccountTypeNSXV captures enum value "NSX-V"
	DeploymentResourceTypeAccountTypeNSXV string = "NSX-V"

	// DeploymentResourceTypeAccountTypeNSXT captures enum value "NSX-T"
	DeploymentResourceTypeAccountTypeNSXT string = "NSX-T"

	// DeploymentResourceTypeAccountTypeNSXP captures enum value "NSX-P"
	DeploymentResourceTypeAccountTypeNSXP string = "NSX-P"

	// DeploymentResourceTypeAccountTypeNSXPCloud captures enum value "NSX-P-cloud"
	DeploymentResourceTypeAccountTypeNSXPCloud string = "NSX-P-cloud"

	// DeploymentResourceTypeAccountTypeVCloudDirector captures enum value "vCloud Director"
	DeploymentResourceTypeAccountTypeVCloudDirector string = "vCloud Director"

	// DeploymentResourceTypeAccountTypeVMC captures enum value "VMC"
	DeploymentResourceTypeAccountTypeVMC string = "VMC"

	// DeploymentResourceTypeAccountTypePuppet captures enum value "Puppet"
	DeploymentResourceTypeAccountTypePuppet string = "Puppet"

	// DeploymentResourceTypeAccountTypeAnsible captures enum value "Ansible"
	DeploymentResourceTypeAccountTypeAnsible string = "Ansible"
)

// prop value enum
func (m *DeploymentResourceType) validateAccountTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, deploymentResourceTypeTypeAccountTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DeploymentResourceType) validateAccountType(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountType) { // not required
		return nil
	}

	// value enum
	if err := m.validateAccountTypeEnum("accountType", "body", m.AccountType); err != nil {
		return err
	}

	return nil
}

func (m *DeploymentResourceType) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeploymentResourceType) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeploymentResourceType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeploymentResourceType) UnmarshalBinary(b []byte) error {
	var res DeploymentResourceType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
