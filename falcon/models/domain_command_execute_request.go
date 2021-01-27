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

// DomainCommandExecuteRequest domain command execute request
//
// swagger:model domain.CommandExecuteRequest
type DomainCommandExecuteRequest struct {

	// base command
	// Required: true
	BaseCommand *string `json:"base_command"`

	// command string
	// Required: true
	CommandString *string `json:"command_string"`

	// device id
	// Required: true
	DeviceID *string `json:"device_id"`

	// id
	// Required: true
	ID *int32 `json:"id"`

	// persist
	// Required: true
	Persist *bool `json:"persist"`

	// session id
	// Required: true
	SessionID *string `json:"session_id"`
}

// Validate validates this domain command execute request
func (m *DomainCommandExecuteRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBaseCommand(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommandString(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePersist(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSessionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainCommandExecuteRequest) validateBaseCommand(formats strfmt.Registry) error {

	if err := validate.Required("base_command", "body", m.BaseCommand); err != nil {
		return err
	}

	return nil
}

func (m *DomainCommandExecuteRequest) validateCommandString(formats strfmt.Registry) error {

	if err := validate.Required("command_string", "body", m.CommandString); err != nil {
		return err
	}

	return nil
}

func (m *DomainCommandExecuteRequest) validateDeviceID(formats strfmt.Registry) error {

	if err := validate.Required("device_id", "body", m.DeviceID); err != nil {
		return err
	}

	return nil
}

func (m *DomainCommandExecuteRequest) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *DomainCommandExecuteRequest) validatePersist(formats strfmt.Registry) error {

	if err := validate.Required("persist", "body", m.Persist); err != nil {
		return err
	}

	return nil
}

func (m *DomainCommandExecuteRequest) validateSessionID(formats strfmt.Registry) error {

	if err := validate.Required("session_id", "body", m.SessionID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this domain command execute request based on context it is used
func (m *DomainCommandExecuteRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainCommandExecuteRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainCommandExecuteRequest) UnmarshalBinary(b []byte) error {
	var res DomainCommandExecuteRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
