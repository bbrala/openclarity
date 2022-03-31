// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RuntimeScanCounters runtime scan counters
//
// swagger:model RuntimeScanCounters
type RuntimeScanCounters struct {

	// applications
	Applications uint32 `json:"applications,omitempty"`

	// packages
	Packages uint32 `json:"packages,omitempty"`

	// resources
	Resources uint32 `json:"resources,omitempty"`

	// vulnerabilities
	Vulnerabilities uint32 `json:"vulnerabilities,omitempty"`
}

// Validate validates this runtime scan counters
func (m *RuntimeScanCounters) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this runtime scan counters based on context it is used
func (m *RuntimeScanCounters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RuntimeScanCounters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RuntimeScanCounters) UnmarshalBinary(b []byte) error {
	var res RuntimeScanCounters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
