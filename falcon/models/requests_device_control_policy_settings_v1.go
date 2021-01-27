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

// RequestsDeviceControlPolicySettingsV1 A specific setting to update
//
// swagger:model requests.DeviceControlPolicySettingsV1
type RequestsDeviceControlPolicySettingsV1 struct {

	// Settings that apply to a USB Class
	// Required: true
	Classes []*RequestsDeviceControlPolicyClassSettingsV1 `json:"classes"`

	// Does the end user receives a notification when the policy is violated
	// Required: true
	// Enum: [TRUE FALSE]
	EndUserNotification *string `json:"end_user_notification"`

	// How is this policy enforced
	// Required: true
	EnforcementMode *string `json:"enforcement_mode"`

	// The id of the setting to update
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this requests device control policy settings v1
func (m *RequestsDeviceControlPolicySettingsV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClasses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndUserNotification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnforcementMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RequestsDeviceControlPolicySettingsV1) validateClasses(formats strfmt.Registry) error {

	if err := validate.Required("classes", "body", m.Classes); err != nil {
		return err
	}

	for i := 0; i < len(m.Classes); i++ {
		if swag.IsZero(m.Classes[i]) { // not required
			continue
		}

		if m.Classes[i] != nil {
			if err := m.Classes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("classes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var requestsDeviceControlPolicySettingsV1TypeEndUserNotificationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["TRUE","FALSE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		requestsDeviceControlPolicySettingsV1TypeEndUserNotificationPropEnum = append(requestsDeviceControlPolicySettingsV1TypeEndUserNotificationPropEnum, v)
	}
}

const (

	// RequestsDeviceControlPolicySettingsV1EndUserNotificationTRUE captures enum value "TRUE"
	RequestsDeviceControlPolicySettingsV1EndUserNotificationTRUE string = "TRUE"

	// RequestsDeviceControlPolicySettingsV1EndUserNotificationFALSE captures enum value "FALSE"
	RequestsDeviceControlPolicySettingsV1EndUserNotificationFALSE string = "FALSE"
)

// prop value enum
func (m *RequestsDeviceControlPolicySettingsV1) validateEndUserNotificationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, requestsDeviceControlPolicySettingsV1TypeEndUserNotificationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RequestsDeviceControlPolicySettingsV1) validateEndUserNotification(formats strfmt.Registry) error {

	if err := validate.Required("end_user_notification", "body", m.EndUserNotification); err != nil {
		return err
	}

	// value enum
	if err := m.validateEndUserNotificationEnum("end_user_notification", "body", *m.EndUserNotification); err != nil {
		return err
	}

	return nil
}

func (m *RequestsDeviceControlPolicySettingsV1) validateEnforcementMode(formats strfmt.Registry) error {

	if err := validate.Required("enforcement_mode", "body", m.EnforcementMode); err != nil {
		return err
	}

	return nil
}

func (m *RequestsDeviceControlPolicySettingsV1) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this requests device control policy settings v1 based on the context it is used
func (m *RequestsDeviceControlPolicySettingsV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClasses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RequestsDeviceControlPolicySettingsV1) contextValidateClasses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Classes); i++ {

		if m.Classes[i] != nil {
			if err := m.Classes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("classes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RequestsDeviceControlPolicySettingsV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RequestsDeviceControlPolicySettingsV1) UnmarshalBinary(b []byte) error {
	var res RequestsDeviceControlPolicySettingsV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
