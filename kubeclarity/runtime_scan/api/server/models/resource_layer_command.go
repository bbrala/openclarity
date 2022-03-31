// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ResourceLayerCommand Container image layer with command.
//
// swagger:model ResourceLayerCommand
type ResourceLayerCommand struct {

	// command
	Command string `json:"command,omitempty"`

	// layer
	Layer string `json:"layer,omitempty"`
}

// Validate validates this resource layer command
func (m *ResourceLayerCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this resource layer command based on context it is used
func (m *ResourceLayerCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResourceLayerCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceLayerCommand) UnmarshalBinary(b []byte) error {
	var res ResourceLayerCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
