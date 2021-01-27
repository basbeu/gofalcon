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

// CloudontologyPatternDisposition cloudontology pattern disposition
//
// swagger:model cloudontology.PatternDisposition
type CloudontologyPatternDisposition struct {

	// bootup safeguard enabled
	// Required: true
	BootupSafeguardEnabled *bool `json:"bootup_safeguard_enabled"`

	// critical process disabled
	// Required: true
	CriticalProcessDisabled *bool `json:"critical_process_disabled"`

	// detect
	// Required: true
	Detect *bool `json:"detect"`

	// fs operation blocked
	// Required: true
	FsOperationBlocked *bool `json:"fs_operation_blocked"`

	// inddet mask
	// Required: true
	InddetMask *bool `json:"inddet_mask"`

	// indicator
	// Required: true
	Indicator *bool `json:"indicator"`

	// kill parent
	// Required: true
	KillParent *bool `json:"kill_parent"`

	// kill process
	// Required: true
	KillProcess *bool `json:"kill_process"`

	// kill subprocess
	// Required: true
	KillSubprocess *bool `json:"kill_subprocess"`

	// operation blocked
	// Required: true
	OperationBlocked *bool `json:"operation_blocked"`

	// policy disabled
	// Required: true
	PolicyDisabled *bool `json:"policy_disabled"`

	// process blocked
	// Required: true
	ProcessBlocked *bool `json:"process_blocked"`

	// quarantine file
	// Required: true
	QuarantineFile *bool `json:"quarantine_file"`

	// quarantine machine
	// Required: true
	QuarantineMachine *bool `json:"quarantine_machine"`

	// registry operation blocked
	// Required: true
	RegistryOperationBlocked *bool `json:"registry_operation_blocked"`

	// rooting
	// Required: true
	Rooting *bool `json:"rooting"`

	// sensor only
	// Required: true
	SensorOnly *bool `json:"sensor_only"`
}

// Validate validates this cloudontology pattern disposition
func (m *CloudontologyPatternDisposition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBootupSafeguardEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCriticalProcessDisabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDetect(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFsOperationBlocked(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInddetMask(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIndicator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKillParent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKillProcess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKillSubprocess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperationBlocked(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicyDisabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcessBlocked(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuarantineFile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuarantineMachine(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegistryOperationBlocked(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRooting(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSensorOnly(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudontologyPatternDisposition) validateBootupSafeguardEnabled(formats strfmt.Registry) error {

	if err := validate.Required("bootup_safeguard_enabled", "body", m.BootupSafeguardEnabled); err != nil {
		return err
	}

	return nil
}

func (m *CloudontologyPatternDisposition) validateCriticalProcessDisabled(formats strfmt.Registry) error {

	if err := validate.Required("critical_process_disabled", "body", m.CriticalProcessDisabled); err != nil {
		return err
	}

	return nil
}

func (m *CloudontologyPatternDisposition) validateDetect(formats strfmt.Registry) error {

	if err := validate.Required("detect", "body", m.Detect); err != nil {
		return err
	}

	return nil
}

func (m *CloudontologyPatternDisposition) validateFsOperationBlocked(formats strfmt.Registry) error {

	if err := validate.Required("fs_operation_blocked", "body", m.FsOperationBlocked); err != nil {
		return err
	}

	return nil
}

func (m *CloudontologyPatternDisposition) validateInddetMask(formats strfmt.Registry) error {

	if err := validate.Required("inddet_mask", "body", m.InddetMask); err != nil {
		return err
	}

	return nil
}

func (m *CloudontologyPatternDisposition) validateIndicator(formats strfmt.Registry) error {

	if err := validate.Required("indicator", "body", m.Indicator); err != nil {
		return err
	}

	return nil
}

func (m *CloudontologyPatternDisposition) validateKillParent(formats strfmt.Registry) error {

	if err := validate.Required("kill_parent", "body", m.KillParent); err != nil {
		return err
	}

	return nil
}

func (m *CloudontologyPatternDisposition) validateKillProcess(formats strfmt.Registry) error {

	if err := validate.Required("kill_process", "body", m.KillProcess); err != nil {
		return err
	}

	return nil
}

func (m *CloudontologyPatternDisposition) validateKillSubprocess(formats strfmt.Registry) error {

	if err := validate.Required("kill_subprocess", "body", m.KillSubprocess); err != nil {
		return err
	}

	return nil
}

func (m *CloudontologyPatternDisposition) validateOperationBlocked(formats strfmt.Registry) error {

	if err := validate.Required("operation_blocked", "body", m.OperationBlocked); err != nil {
		return err
	}

	return nil
}

func (m *CloudontologyPatternDisposition) validatePolicyDisabled(formats strfmt.Registry) error {

	if err := validate.Required("policy_disabled", "body", m.PolicyDisabled); err != nil {
		return err
	}

	return nil
}

func (m *CloudontologyPatternDisposition) validateProcessBlocked(formats strfmt.Registry) error {

	if err := validate.Required("process_blocked", "body", m.ProcessBlocked); err != nil {
		return err
	}

	return nil
}

func (m *CloudontologyPatternDisposition) validateQuarantineFile(formats strfmt.Registry) error {

	if err := validate.Required("quarantine_file", "body", m.QuarantineFile); err != nil {
		return err
	}

	return nil
}

func (m *CloudontologyPatternDisposition) validateQuarantineMachine(formats strfmt.Registry) error {

	if err := validate.Required("quarantine_machine", "body", m.QuarantineMachine); err != nil {
		return err
	}

	return nil
}

func (m *CloudontologyPatternDisposition) validateRegistryOperationBlocked(formats strfmt.Registry) error {

	if err := validate.Required("registry_operation_blocked", "body", m.RegistryOperationBlocked); err != nil {
		return err
	}

	return nil
}

func (m *CloudontologyPatternDisposition) validateRooting(formats strfmt.Registry) error {

	if err := validate.Required("rooting", "body", m.Rooting); err != nil {
		return err
	}

	return nil
}

func (m *CloudontologyPatternDisposition) validateSensorOnly(formats strfmt.Registry) error {

	if err := validate.Required("sensor_only", "body", m.SensorOnly); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cloudontology pattern disposition based on context it is used
func (m *CloudontologyPatternDisposition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CloudontologyPatternDisposition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudontologyPatternDisposition) UnmarshalBinary(b []byte) error {
	var res CloudontologyPatternDisposition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
