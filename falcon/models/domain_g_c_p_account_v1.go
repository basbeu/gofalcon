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

// DomainGCPAccountV1 domain g c p account v1
//
// swagger:model domain.GCPAccountV1
type DomainGCPAccountV1 struct {

	// cid
	// Required: true
	Cid *string `json:"cid"`

	// GCP ParentID.
	// Required: true
	ParentID *string `json:"parent_id"`

	// Account registration status.
	Status string `json:"status,omitempty"`
}

// Validate validates this domain g c p account v1
func (m *DomainGCPAccountV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainGCPAccountV1) validateCid(formats strfmt.Registry) error {

	if err := validate.Required("cid", "body", m.Cid); err != nil {
		return err
	}

	return nil
}

func (m *DomainGCPAccountV1) validateParentID(formats strfmt.Registry) error {

	if err := validate.Required("parent_id", "body", m.ParentID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this domain g c p account v1 based on context it is used
func (m *DomainGCPAccountV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainGCPAccountV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainGCPAccountV1) UnmarshalBinary(b []byte) error {
	var res DomainGCPAccountV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
