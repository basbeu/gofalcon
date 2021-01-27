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

// FwmgrMsaRangeSpec fwmgr msa range spec
//
// swagger:model fwmgr.msa.RangeSpec
type FwmgrMsaRangeSpec struct {

	// from
	// Required: true
	From *float64 `json:"From"`

	// to
	// Required: true
	To *float64 `json:"To"`
}

// Validate validates this fwmgr msa range spec
func (m *FwmgrMsaRangeSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FwmgrMsaRangeSpec) validateFrom(formats strfmt.Registry) error {

	if err := validate.Required("From", "body", m.From); err != nil {
		return err
	}

	return nil
}

func (m *FwmgrMsaRangeSpec) validateTo(formats strfmt.Registry) error {

	if err := validate.Required("To", "body", m.To); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this fwmgr msa range spec based on context it is used
func (m *FwmgrMsaRangeSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FwmgrMsaRangeSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FwmgrMsaRangeSpec) UnmarshalBinary(b []byte) error {
	var res FwmgrMsaRangeSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
