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

// HumioViewConnection humio view connection
//
// swagger:model humio.ViewConnection
type HumioViewConnection struct {

	// filter
	// Required: true
	Filter *string `json:"filter"`

	// repo name
	// Required: true
	RepoName *string `json:"repo_name"`
}

// Validate validates this humio view connection
func (m *HumioViewConnection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepoName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HumioViewConnection) validateFilter(formats strfmt.Registry) error {

	if err := validate.Required("filter", "body", m.Filter); err != nil {
		return err
	}

	return nil
}

func (m *HumioViewConnection) validateRepoName(formats strfmt.Registry) error {

	if err := validate.Required("repo_name", "body", m.RepoName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this humio view connection based on context it is used
func (m *HumioViewConnection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HumioViewConnection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HumioViewConnection) UnmarshalBinary(b []byte) error {
	var res HumioViewConnection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}