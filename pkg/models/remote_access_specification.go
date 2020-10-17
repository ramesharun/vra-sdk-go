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

// RemoteAccessSpecification Represents a specification for machine's remote access settings.
//
// swagger:model RemoteAccessSpecification
type RemoteAccessSpecification struct {

	// One of four authentication types.
	// `generatedPublicPrivateKey`: The provisioned machine generates the public/private key pair and enables SSH to use them without user input.
	// `publicPrivateKey`: The user enters the private key in the SSH command. See remoteAccess.sshKey.
	// `usernamePassword`: The user enters a username and password for remote access.
	// `keyPairName`: The user enters an already existing keyPair name. See remoteAccess.keyPair
	// Required: true
	Authentication *string `json:"authentication"`

	// Key Pair Name.
	KeyPair string `json:"keyPair,omitempty"`

	// Remote access password for the Azure machine.
	Password string `json:"password,omitempty"`

	// In key pair authentication, the public key on the provisioned machine. Users are expected to log in with their private key and a default username from the cloud provider. An AWS Ubuntu image comes with default user ubuntu, and Azure comes with default user azureuser. To log in by SSH:
	// `ssh -i <private-key-path> ubuntu@52.90.80.153`
	// `ssh -i <private-key-path> azureuser@40.76.14.255`
	SSHKey string `json:"sshKey,omitempty"`

	// Remote access username for the Azure machine.
	Username string `json:"username,omitempty"`
}

// Validate validates this remote access specification
func (m *RemoteAccessSpecification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthentication(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RemoteAccessSpecification) validateAuthentication(formats strfmt.Registry) error {

	if err := validate.Required("authentication", "body", m.Authentication); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RemoteAccessSpecification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RemoteAccessSpecification) UnmarshalBinary(b []byte) error {
	var res RemoteAccessSpecification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
