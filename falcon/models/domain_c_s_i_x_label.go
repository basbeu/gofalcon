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

// DomainCSIXLabel domain c s i x label
//
// swagger:model domain.CSIXLabel
type DomainCSIXLabel struct {

	// created on
	// Required: true
	CreatedOn *int64 `json:"created_on"`

	// last valid on
	// Required: true
	LastValidOn *int64 `json:"last_valid_on"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this domain c s i x label
func (m *DomainCSIXLabel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastValidOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainCSIXLabel) validateCreatedOn(formats strfmt.Registry) error {

	if err := validate.Required("created_on", "body", m.CreatedOn); err != nil {
		return err
	}

	return nil
}

func (m *DomainCSIXLabel) validateLastValidOn(formats strfmt.Registry) error {

	if err := validate.Required("last_valid_on", "body", m.LastValidOn); err != nil {
		return err
	}

	return nil
}

func (m *DomainCSIXLabel) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this domain c s i x label based on context it is used
func (m *DomainCSIXLabel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainCSIXLabel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainCSIXLabel) UnmarshalBinary(b []byte) error {
	var res DomainCSIXLabel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
