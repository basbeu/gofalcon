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

// MsaEntityActionRequestV2 msa entity action request v2
//
// swagger:model msa.EntityActionRequestV2
type MsaEntityActionRequestV2 struct {

	// action parameters
	ActionParameters []*CsdomainActionParameter `json:"action_parameters"`

	// ids
	// Required: true
	Ids []string `json:"ids"`
}

// Validate validates this msa entity action request v2
func (m *MsaEntityActionRequestV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionParameters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MsaEntityActionRequestV2) validateActionParameters(formats strfmt.Registry) error {
	if swag.IsZero(m.ActionParameters) { // not required
		return nil
	}

	for i := 0; i < len(m.ActionParameters); i++ {
		if swag.IsZero(m.ActionParameters[i]) { // not required
			continue
		}

		if m.ActionParameters[i] != nil {
			if err := m.ActionParameters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("action_parameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MsaEntityActionRequestV2) validateIds(formats strfmt.Registry) error {

	if err := validate.Required("ids", "body", m.Ids); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this msa entity action request v2 based on the context it is used
func (m *MsaEntityActionRequestV2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActionParameters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MsaEntityActionRequestV2) contextValidateActionParameters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ActionParameters); i++ {

		if m.ActionParameters[i] != nil {
			if err := m.ActionParameters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("action_parameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MsaEntityActionRequestV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsaEntityActionRequestV2) UnmarshalBinary(b []byte) error {
	var res MsaEntityActionRequestV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
