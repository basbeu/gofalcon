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

// DomainMultiStatusSensorResponse domain multi status sensor response
//
// swagger:model domain.MultiStatusSensorResponse
type DomainMultiStatusSensorResponse struct {

	// aid
	// Required: true
	Aid *string `json:"aid"`

	// base command
	BaseCommand string `json:"base_command,omitempty"`

	// complete
	// Required: true
	Complete *bool `json:"complete"`

	// errors
	// Required: true
	Errors []*MsaAPIError `json:"errors"`

	// offline queued
	// Required: true
	OfflineQueued *bool `json:"offline_queued"`

	// query time
	// Required: true
	QueryTime *float64 `json:"query_time"`

	// sequence id
	SequenceID int64 `json:"sequence_id,omitempty"`

	// session id
	// Required: true
	SessionID *string `json:"session_id"`

	// stderr
	// Required: true
	Stderr *string `json:"stderr"`

	// stdout
	// Required: true
	Stdout *string `json:"stdout"`

	// task id
	TaskID string `json:"task_id,omitempty"`
}

// Validate validates this domain multi status sensor response
func (m *DomainMultiStatusSensorResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComplete(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOfflineQueued(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueryTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSessionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStderr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStdout(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainMultiStatusSensorResponse) validateAid(formats strfmt.Registry) error {

	if err := validate.Required("aid", "body", m.Aid); err != nil {
		return err
	}

	return nil
}

func (m *DomainMultiStatusSensorResponse) validateComplete(formats strfmt.Registry) error {

	if err := validate.Required("complete", "body", m.Complete); err != nil {
		return err
	}

	return nil
}

func (m *DomainMultiStatusSensorResponse) validateErrors(formats strfmt.Registry) error {

	if err := validate.Required("errors", "body", m.Errors); err != nil {
		return err
	}

	for i := 0; i < len(m.Errors); i++ {
		if swag.IsZero(m.Errors[i]) { // not required
			continue
		}

		if m.Errors[i] != nil {
			if err := m.Errors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainMultiStatusSensorResponse) validateOfflineQueued(formats strfmt.Registry) error {

	if err := validate.Required("offline_queued", "body", m.OfflineQueued); err != nil {
		return err
	}

	return nil
}

func (m *DomainMultiStatusSensorResponse) validateQueryTime(formats strfmt.Registry) error {

	if err := validate.Required("query_time", "body", m.QueryTime); err != nil {
		return err
	}

	return nil
}

func (m *DomainMultiStatusSensorResponse) validateSessionID(formats strfmt.Registry) error {

	if err := validate.Required("session_id", "body", m.SessionID); err != nil {
		return err
	}

	return nil
}

func (m *DomainMultiStatusSensorResponse) validateStderr(formats strfmt.Registry) error {

	if err := validate.Required("stderr", "body", m.Stderr); err != nil {
		return err
	}

	return nil
}

func (m *DomainMultiStatusSensorResponse) validateStdout(formats strfmt.Registry) error {

	if err := validate.Required("stdout", "body", m.Stdout); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this domain multi status sensor response based on the context it is used
func (m *DomainMultiStatusSensorResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainMultiStatusSensorResponse) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Errors); i++ {

		if m.Errors[i] != nil {
			if err := m.Errors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainMultiStatusSensorResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainMultiStatusSensorResponse) UnmarshalBinary(b []byte) error {
	var res DomainMultiStatusSensorResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
