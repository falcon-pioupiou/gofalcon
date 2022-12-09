// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// K8sregAzureTenantConfig k8sreg azure tenant config
//
// swagger:model k8sreg.AzureTenantConfig
type K8sregAzureTenantConfig struct {

	// client id
	// Required: true
	ClientID *string `json:"client_id"`

	// public certificate
	PublicCertificate string `json:"public_certificate,omitempty"`

	// tenant id
	// Required: true
	TenantID *string `json:"tenant_id"`
}

// Validate validates this k8sreg azure tenant config
func (m *K8sregAzureTenantConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenantID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8sregAzureTenantConfig) validateClientID(formats strfmt.Registry) error {

	if err := validate.Required("client_id", "body", m.ClientID); err != nil {
		return err
	}

	return nil
}

func (m *K8sregAzureTenantConfig) validateTenantID(formats strfmt.Registry) error {

	if err := validate.Required("tenant_id", "body", m.TenantID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this k8sreg azure tenant config based on context it is used
func (m *K8sregAzureTenantConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *K8sregAzureTenantConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *K8sregAzureTenantConfig) UnmarshalBinary(b []byte) error {
	var res K8sregAzureTenantConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
