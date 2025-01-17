// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// APIDeviceCountRespV1 api device count resp v1
//
// swagger:model api.DeviceCountRespV1
type APIDeviceCountRespV1 struct {

	// errors
	// Required: true
	Errors []*MsaAPIError `json:"Errors"`

	// msa meta info
	// Required: true
	MsaMetaInfo *MsaMetaInfo `json:"MsaMetaInfo"`

	// resources
	// Required: true
	Resources []*APIDeviceCountV1 `json:"resources"`
}

// Validate validates this api device count resp v1
func (m *APIDeviceCountRespV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMsaMetaInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIDeviceCountRespV1) validateErrors(formats strfmt.Registry) error {

	if err := validate.Required("Errors", "body", m.Errors); err != nil {
		return err
	}

	for i := 0; i < len(m.Errors); i++ {
		if swag.IsZero(m.Errors[i]) { // not required
			continue
		}

		if m.Errors[i] != nil {
			if err := m.Errors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Errors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *APIDeviceCountRespV1) validateMsaMetaInfo(formats strfmt.Registry) error {

	if err := validate.Required("MsaMetaInfo", "body", m.MsaMetaInfo); err != nil {
		return err
	}

	if m.MsaMetaInfo != nil {
		if err := m.MsaMetaInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MsaMetaInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("MsaMetaInfo")
			}
			return err
		}
	}

	return nil
}

func (m *APIDeviceCountRespV1) validateResources(formats strfmt.Registry) error {

	if err := validate.Required("resources", "body", m.Resources); err != nil {
		return err
	}

	for i := 0; i < len(m.Resources); i++ {
		if swag.IsZero(m.Resources[i]) { // not required
			continue
		}

		if m.Resources[i] != nil {
			if err := m.Resources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("resources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this api device count resp v1 based on the context it is used
func (m *APIDeviceCountRespV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMsaMetaInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIDeviceCountRespV1) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Errors); i++ {

		if m.Errors[i] != nil {

			if swag.IsZero(m.Errors[i]) { // not required
				return nil
			}

			if err := m.Errors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Errors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *APIDeviceCountRespV1) contextValidateMsaMetaInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.MsaMetaInfo != nil {

		if err := m.MsaMetaInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MsaMetaInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("MsaMetaInfo")
			}
			return err
		}
	}

	return nil
}

func (m *APIDeviceCountRespV1) contextValidateResources(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Resources); i++ {

		if m.Resources[i] != nil {

			if swag.IsZero(m.Resources[i]) { // not required
				return nil
			}

			if err := m.Resources[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("resources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIDeviceCountRespV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIDeviceCountRespV1) UnmarshalBinary(b []byte) error {
	var res APIDeviceCountRespV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
