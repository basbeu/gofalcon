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

// APIRuleGroupModifyRequestV1 api rule group modify request v1
//
// swagger:model api.RuleGroupModifyRequestV1
type APIRuleGroupModifyRequestV1 struct {

	// comment
	// Required: true
	Comment *string `json:"comment"`

	// description
	// Required: true
	Description *string `json:"description"`

	// enabled
	// Required: true
	Enabled *bool `json:"enabled"`

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`

	// rulegroup version
	// Required: true
	RulegroupVersion *int64 `json:"rulegroup_version"`
}

// Validate validates this api rule group modify request v1
func (m *APIRuleGroupModifyRequestV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRulegroupVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIRuleGroupModifyRequestV1) validateComment(formats strfmt.Registry) error {

	if err := validate.Required("comment", "body", m.Comment); err != nil {
		return err
	}

	return nil
}

func (m *APIRuleGroupModifyRequestV1) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *APIRuleGroupModifyRequestV1) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *APIRuleGroupModifyRequestV1) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *APIRuleGroupModifyRequestV1) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *APIRuleGroupModifyRequestV1) validateRulegroupVersion(formats strfmt.Registry) error {

	if err := validate.Required("rulegroup_version", "body", m.RulegroupVersion); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this api rule group modify request v1 based on context it is used
func (m *APIRuleGroupModifyRequestV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIRuleGroupModifyRequestV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIRuleGroupModifyRequestV1) UnmarshalBinary(b []byte) error {
	var res APIRuleGroupModifyRequestV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
