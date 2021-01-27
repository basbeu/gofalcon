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

// DetectsindexHostInfo detectsindex host info
//
// swagger:model detectsindex.HostInfo
type DetectsindexHostInfo struct {

	// active directory dn display
	ActiveDirectoryDnDisplay []string `json:"active_directory_dn_display"`

	// domain
	// Required: true
	Domain *string `json:"domain"`
}

// Validate validates this detectsindex host info
func (m *DetectsindexHostInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDomain(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DetectsindexHostInfo) validateDomain(formats strfmt.Registry) error {

	if err := validate.Required("domain", "body", m.Domain); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this detectsindex host info based on context it is used
func (m *DetectsindexHostInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DetectsindexHostInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DetectsindexHostInfo) UnmarshalBinary(b []byte) error {
	var res DetectsindexHostInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
