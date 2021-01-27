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

// FwmgrDomainICMP fwmgr domain i c m p
//
// swagger:model fwmgr.domain.ICMP
type FwmgrDomainICMP struct {

	// icmp code
	// Required: true
	IcmpCode *string `json:"icmp_code"`

	// icmp type
	// Required: true
	IcmpType *string `json:"icmp_type"`
}

// Validate validates this fwmgr domain i c m p
func (m *FwmgrDomainICMP) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIcmpCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIcmpType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FwmgrDomainICMP) validateIcmpCode(formats strfmt.Registry) error {

	if err := validate.Required("icmp_code", "body", m.IcmpCode); err != nil {
		return err
	}

	return nil
}

func (m *FwmgrDomainICMP) validateIcmpType(formats strfmt.Registry) error {

	if err := validate.Required("icmp_type", "body", m.IcmpType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this fwmgr domain i c m p based on context it is used
func (m *FwmgrDomainICMP) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FwmgrDomainICMP) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FwmgrDomainICMP) UnmarshalBinary(b []byte) error {
	var res FwmgrDomainICMP
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
