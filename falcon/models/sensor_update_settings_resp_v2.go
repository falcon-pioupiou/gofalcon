// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SensorUpdateSettingsRespV2 sensor update settings resp v2
//
// swagger:model sensor_update.SettingsRespV2
type SensorUpdateSettingsRespV2 struct {

	// The target build applied to devices in the policy
	// Required: true
	Build *string `json:"build"`

	// The schedule that disables sensor updates
	// Required: true
	Scheduler *PolicySensorUpdateScheduler `json:"scheduler"`

	// sensor version
	// Required: true
	SensorVersion *string `json:"sensor_version"`

	// If true, early adopter builds will be visible on the sensor update policy page
	// Required: true
	ShowEarlyAdopterBuilds *bool `json:"show_early_adopter_builds"`

	// The release stage this build is in
	// Required: true
	// Enum: [prod early_adopter]
	Stage *string `json:"stage"`

	// The uninstall protection setting to apply to devices in the policy
	// Required: true
	// Enum: [ENABLED DISABLED MAINTENANCE_MODE IGNORE UNKNOWN]
	UninstallProtection *string `json:"uninstall_protection"`

	// variants
	// Required: true
	Variants []*SensorUpdateBuildRespV1 `json:"variants"`
}

// Validate validates this sensor update settings resp v2
func (m *SensorUpdateSettingsRespV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuild(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScheduler(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSensorVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShowEarlyAdopterBuilds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUninstallProtection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVariants(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SensorUpdateSettingsRespV2) validateBuild(formats strfmt.Registry) error {

	if err := validate.Required("build", "body", m.Build); err != nil {
		return err
	}

	return nil
}

func (m *SensorUpdateSettingsRespV2) validateScheduler(formats strfmt.Registry) error {

	if err := validate.Required("scheduler", "body", m.Scheduler); err != nil {
		return err
	}

	if m.Scheduler != nil {
		if err := m.Scheduler.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scheduler")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scheduler")
			}
			return err
		}
	}

	return nil
}

func (m *SensorUpdateSettingsRespV2) validateSensorVersion(formats strfmt.Registry) error {

	if err := validate.Required("sensor_version", "body", m.SensorVersion); err != nil {
		return err
	}

	return nil
}

func (m *SensorUpdateSettingsRespV2) validateShowEarlyAdopterBuilds(formats strfmt.Registry) error {

	if err := validate.Required("show_early_adopter_builds", "body", m.ShowEarlyAdopterBuilds); err != nil {
		return err
	}

	return nil
}

var sensorUpdateSettingsRespV2TypeStagePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["prod","early_adopter"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sensorUpdateSettingsRespV2TypeStagePropEnum = append(sensorUpdateSettingsRespV2TypeStagePropEnum, v)
	}
}

const (

	// SensorUpdateSettingsRespV2StageProd captures enum value "prod"
	SensorUpdateSettingsRespV2StageProd string = "prod"

	// SensorUpdateSettingsRespV2StageEarlyAdopter captures enum value "early_adopter"
	SensorUpdateSettingsRespV2StageEarlyAdopter string = "early_adopter"
)

// prop value enum
func (m *SensorUpdateSettingsRespV2) validateStageEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sensorUpdateSettingsRespV2TypeStagePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SensorUpdateSettingsRespV2) validateStage(formats strfmt.Registry) error {

	if err := validate.Required("stage", "body", m.Stage); err != nil {
		return err
	}

	// value enum
	if err := m.validateStageEnum("stage", "body", *m.Stage); err != nil {
		return err
	}

	return nil
}

var sensorUpdateSettingsRespV2TypeUninstallProtectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ENABLED","DISABLED","MAINTENANCE_MODE","IGNORE","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sensorUpdateSettingsRespV2TypeUninstallProtectionPropEnum = append(sensorUpdateSettingsRespV2TypeUninstallProtectionPropEnum, v)
	}
}

const (

	// SensorUpdateSettingsRespV2UninstallProtectionENABLED captures enum value "ENABLED"
	SensorUpdateSettingsRespV2UninstallProtectionENABLED string = "ENABLED"

	// SensorUpdateSettingsRespV2UninstallProtectionDISABLED captures enum value "DISABLED"
	SensorUpdateSettingsRespV2UninstallProtectionDISABLED string = "DISABLED"

	// SensorUpdateSettingsRespV2UninstallProtectionMAINTENANCEMODE captures enum value "MAINTENANCE_MODE"
	SensorUpdateSettingsRespV2UninstallProtectionMAINTENANCEMODE string = "MAINTENANCE_MODE"

	// SensorUpdateSettingsRespV2UninstallProtectionIGNORE captures enum value "IGNORE"
	SensorUpdateSettingsRespV2UninstallProtectionIGNORE string = "IGNORE"

	// SensorUpdateSettingsRespV2UninstallProtectionUNKNOWN captures enum value "UNKNOWN"
	SensorUpdateSettingsRespV2UninstallProtectionUNKNOWN string = "UNKNOWN"
)

// prop value enum
func (m *SensorUpdateSettingsRespV2) validateUninstallProtectionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sensorUpdateSettingsRespV2TypeUninstallProtectionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SensorUpdateSettingsRespV2) validateUninstallProtection(formats strfmt.Registry) error {

	if err := validate.Required("uninstall_protection", "body", m.UninstallProtection); err != nil {
		return err
	}

	// value enum
	if err := m.validateUninstallProtectionEnum("uninstall_protection", "body", *m.UninstallProtection); err != nil {
		return err
	}

	return nil
}

func (m *SensorUpdateSettingsRespV2) validateVariants(formats strfmt.Registry) error {

	if err := validate.Required("variants", "body", m.Variants); err != nil {
		return err
	}

	for i := 0; i < len(m.Variants); i++ {
		if swag.IsZero(m.Variants[i]) { // not required
			continue
		}

		if m.Variants[i] != nil {
			if err := m.Variants[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("variants" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("variants" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this sensor update settings resp v2 based on the context it is used
func (m *SensorUpdateSettingsRespV2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateScheduler(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVariants(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SensorUpdateSettingsRespV2) contextValidateScheduler(ctx context.Context, formats strfmt.Registry) error {

	if m.Scheduler != nil {

		if err := m.Scheduler.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scheduler")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scheduler")
			}
			return err
		}
	}

	return nil
}

func (m *SensorUpdateSettingsRespV2) contextValidateVariants(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Variants); i++ {

		if m.Variants[i] != nil {

			if swag.IsZero(m.Variants[i]) { // not required
				return nil
			}

			if err := m.Variants[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("variants" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("variants" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SensorUpdateSettingsRespV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SensorUpdateSettingsRespV2) UnmarshalBinary(b []byte) error {
	var res SensorUpdateSettingsRespV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}