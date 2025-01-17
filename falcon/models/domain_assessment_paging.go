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

// DomainAssessmentPaging domain assessment paging
//
// swagger:model domain.AssessmentPaging
type DomainAssessmentPaging struct {

	// expires at
	ExpiresAt int64 `json:"expires_at,omitempty"`

	// limit
	Limit int32 `json:"limit,omitempty"`

	// offset
	// Required: true
	Offset *string `json:"offset"`

	// total
	// Required: true
	Total *int64 `json:"total"`
}

// Validate validates this domain assessment paging
func (m *DomainAssessmentPaging) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOffset(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainAssessmentPaging) validateOffset(formats strfmt.Registry) error {

	if err := validate.Required("offset", "body", m.Offset); err != nil {
		return err
	}

	return nil
}

func (m *DomainAssessmentPaging) validateTotal(formats strfmt.Registry) error {

	if err := validate.Required("total", "body", m.Total); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this domain assessment paging based on context it is used
func (m *DomainAssessmentPaging) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainAssessmentPaging) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainAssessmentPaging) UnmarshalBinary(b []byte) error {
	var res DomainAssessmentPaging
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
