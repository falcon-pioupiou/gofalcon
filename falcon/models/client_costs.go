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

// ClientCosts client costs
//
// swagger:model client.Costs
type ClientCosts struct {

	// live cost
	// Required: true
	LiveCost *float64 `json:"liveCost"`

	// live cost rate
	// Required: true
	LiveCostRate *float64 `json:"liveCostRate"`

	// static cost
	// Required: true
	StaticCost *float64 `json:"staticCost"`

	// static cost rate
	// Required: true
	StaticCostRate *float64 `json:"staticCostRate"`
}

// Validate validates this client costs
func (m *ClientCosts) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLiveCost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLiveCostRate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStaticCost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStaticCostRate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClientCosts) validateLiveCost(formats strfmt.Registry) error {

	if err := validate.Required("liveCost", "body", m.LiveCost); err != nil {
		return err
	}

	return nil
}

func (m *ClientCosts) validateLiveCostRate(formats strfmt.Registry) error {

	if err := validate.Required("liveCostRate", "body", m.LiveCostRate); err != nil {
		return err
	}

	return nil
}

func (m *ClientCosts) validateStaticCost(formats strfmt.Registry) error {

	if err := validate.Required("staticCost", "body", m.StaticCost); err != nil {
		return err
	}

	return nil
}

func (m *ClientCosts) validateStaticCostRate(formats strfmt.Registry) error {

	if err := validate.Required("staticCostRate", "body", m.StaticCostRate); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this client costs based on context it is used
func (m *ClientCosts) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClientCosts) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClientCosts) UnmarshalBinary(b []byte) error {
	var res ClientCosts
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}