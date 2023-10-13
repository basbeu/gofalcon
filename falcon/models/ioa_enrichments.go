// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoaEnrichments ioa enrichments
//
// swagger:model ioa.Enrichments
type IoaEnrichments struct {

	// inventory
	Inventory *DetectionInventoryEnrichment `json:"inventory,omitempty"`

	// sensor events
	SensorEvents *DetectionEnrichment `json:"sensor_events,omitempty"`
}

// Validate validates this ioa enrichments
func (m *IoaEnrichments) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInventory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSensorEvents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoaEnrichments) validateInventory(formats strfmt.Registry) error {
	if swag.IsZero(m.Inventory) { // not required
		return nil
	}

	if m.Inventory != nil {
		if err := m.Inventory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inventory")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("inventory")
			}
			return err
		}
	}

	return nil
}

func (m *IoaEnrichments) validateSensorEvents(formats strfmt.Registry) error {
	if swag.IsZero(m.SensorEvents) { // not required
		return nil
	}

	if m.SensorEvents != nil {
		if err := m.SensorEvents.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sensor_events")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sensor_events")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ioa enrichments based on the context it is used
func (m *IoaEnrichments) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInventory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSensorEvents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoaEnrichments) contextValidateInventory(ctx context.Context, formats strfmt.Registry) error {

	if m.Inventory != nil {

		if swag.IsZero(m.Inventory) { // not required
			return nil
		}

		if err := m.Inventory.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inventory")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("inventory")
			}
			return err
		}
	}

	return nil
}

func (m *IoaEnrichments) contextValidateSensorEvents(ctx context.Context, formats strfmt.Registry) error {

	if m.SensorEvents != nil {

		if swag.IsZero(m.SensorEvents) { // not required
			return nil
		}

		if err := m.SensorEvents.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sensor_events")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sensor_events")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoaEnrichments) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoaEnrichments) UnmarshalBinary(b []byte) error {
	var res IoaEnrichments
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}