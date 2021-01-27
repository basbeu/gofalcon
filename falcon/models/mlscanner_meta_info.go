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

// MlscannerMetaInfo mlscanner meta info
//
// swagger:model mlscanner.MetaInfo
type MlscannerMetaInfo struct {

	// pagination
	Pagination *MsaPaging `json:"pagination,omitempty"`

	// powered by
	PoweredBy string `json:"powered_by,omitempty"`

	// query time
	// Required: true
	QueryTime *float64 `json:"query_time"`

	// quota
	Quota *MlscannerQuota `json:"quota,omitempty"`

	// trace id
	// Required: true
	TraceID *string `json:"trace_id"`

	// writes
	Writes *MsaResources `json:"writes,omitempty"`
}

// Validate validates this mlscanner meta info
func (m *MlscannerMetaInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueryTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuota(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTraceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWrites(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MlscannerMetaInfo) validatePagination(formats strfmt.Registry) error {
	if swag.IsZero(m.Pagination) { // not required
		return nil
	}

	if m.Pagination != nil {
		if err := m.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

func (m *MlscannerMetaInfo) validateQueryTime(formats strfmt.Registry) error {

	if err := validate.Required("query_time", "body", m.QueryTime); err != nil {
		return err
	}

	return nil
}

func (m *MlscannerMetaInfo) validateQuota(formats strfmt.Registry) error {
	if swag.IsZero(m.Quota) { // not required
		return nil
	}

	if m.Quota != nil {
		if err := m.Quota.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quota")
			}
			return err
		}
	}

	return nil
}

func (m *MlscannerMetaInfo) validateTraceID(formats strfmt.Registry) error {

	if err := validate.Required("trace_id", "body", m.TraceID); err != nil {
		return err
	}

	return nil
}

func (m *MlscannerMetaInfo) validateWrites(formats strfmt.Registry) error {
	if swag.IsZero(m.Writes) { // not required
		return nil
	}

	if m.Writes != nil {
		if err := m.Writes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("writes")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this mlscanner meta info based on the context it is used
func (m *MlscannerMetaInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePagination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQuota(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWrites(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MlscannerMetaInfo) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

	if m.Pagination != nil {
		if err := m.Pagination.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

func (m *MlscannerMetaInfo) contextValidateQuota(ctx context.Context, formats strfmt.Registry) error {

	if m.Quota != nil {
		if err := m.Quota.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quota")
			}
			return err
		}
	}

	return nil
}

func (m *MlscannerMetaInfo) contextValidateWrites(ctx context.Context, formats strfmt.Registry) error {

	if m.Writes != nil {
		if err := m.Writes.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("writes")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MlscannerMetaInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MlscannerMetaInfo) UnmarshalBinary(b []byte) error {
	var res MlscannerMetaInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
