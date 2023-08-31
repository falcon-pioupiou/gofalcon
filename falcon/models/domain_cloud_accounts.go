// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DomainCloudAccounts domain cloud accounts
//
// swagger:model domain.CloudAccounts
type DomainCloudAccounts struct {

	// ids
	Ids []string `json:"ids"`

	// provider
	Provider string `json:"provider,omitempty"`
}

// Validate validates this domain cloud accounts
func (m *DomainCloudAccounts) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this domain cloud accounts based on context it is used
func (m *DomainCloudAccounts) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainCloudAccounts) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainCloudAccounts) UnmarshalBinary(b []byte) error {
	var res DomainCloudAccounts
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}