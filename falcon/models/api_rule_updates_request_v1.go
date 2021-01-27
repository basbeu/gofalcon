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

// APIRuleUpdatesRequestV1 api rule updates request v1
//
// swagger:model api.RuleUpdatesRequestV1
type APIRuleUpdatesRequestV1 struct {

	// comment
	// Required: true
	Comment *string `json:"comment"`

	// rule updates
	// Required: true
	RuleUpdates []*APIRuleUpdateV1 `json:"rule_updates"`

	// rulegroup id
	// Required: true
	RulegroupID *string `json:"rulegroup_id"`

	// rulegroup version
	// Required: true
	RulegroupVersion *int64 `json:"rulegroup_version"`
}

// Validate validates this api rule updates request v1
func (m *APIRuleUpdatesRequestV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuleUpdates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRulegroupID(formats); err != nil {
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

func (m *APIRuleUpdatesRequestV1) validateComment(formats strfmt.Registry) error {

	if err := validate.Required("comment", "body", m.Comment); err != nil {
		return err
	}

	return nil
}

func (m *APIRuleUpdatesRequestV1) validateRuleUpdates(formats strfmt.Registry) error {

	if err := validate.Required("rule_updates", "body", m.RuleUpdates); err != nil {
		return err
	}

	for i := 0; i < len(m.RuleUpdates); i++ {
		if swag.IsZero(m.RuleUpdates[i]) { // not required
			continue
		}

		if m.RuleUpdates[i] != nil {
			if err := m.RuleUpdates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rule_updates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *APIRuleUpdatesRequestV1) validateRulegroupID(formats strfmt.Registry) error {

	if err := validate.Required("rulegroup_id", "body", m.RulegroupID); err != nil {
		return err
	}

	return nil
}

func (m *APIRuleUpdatesRequestV1) validateRulegroupVersion(formats strfmt.Registry) error {

	if err := validate.Required("rulegroup_version", "body", m.RulegroupVersion); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this api rule updates request v1 based on the context it is used
func (m *APIRuleUpdatesRequestV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRuleUpdates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIRuleUpdatesRequestV1) contextValidateRuleUpdates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RuleUpdates); i++ {

		if m.RuleUpdates[i] != nil {
			if err := m.RuleUpdates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rule_updates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIRuleUpdatesRequestV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIRuleUpdatesRequestV1) UnmarshalBinary(b []byte) error {
	var res APIRuleUpdatesRequestV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
