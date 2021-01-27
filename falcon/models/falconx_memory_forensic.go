// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FalconxMemoryForensic falconx memory forensic
//
// swagger:model falconx.MemoryForensic
type FalconxMemoryForensic struct {

	// stream uid
	StreamUID string `json:"stream_uid,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this falconx memory forensic
func (m *FalconxMemoryForensic) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this falconx memory forensic based on context it is used
func (m *FalconxMemoryForensic) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FalconxMemoryForensic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FalconxMemoryForensic) UnmarshalBinary(b []byte) error {
	var res FalconxMemoryForensic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
