// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FindingAttribution finding attribution
//
// swagger:model FindingAttribution
type FindingAttribution struct {

	// alternate identifier
	AlternateIdentifier string `json:"alternateIdentifier,omitempty"`

	// analyzer identity
	// Enum: [INTERNAL_ANALYZER OSSINDEX_ANALYZER NPM_AUDIT_ANALYZER VULNDB_ANALYZER NONE]
	AnalyzerIdentity string `json:"analyzerIdentity,omitempty"`

	// attributed on
	// Required: true
	// Format: date-time
	AttributedOn *strfmt.DateTime `json:"attributedOn"`

	// component
	// Required: true
	Component *Component `json:"component"`

	// reference Url
	ReferenceURL string `json:"referenceUrl,omitempty"`

	// uuid
	// Required: true
	// Format: uuid
	UUID *strfmt.UUID `json:"uuid"`

	// vulnerability
	// Required: true
	Vulnerability *Vulnerability `json:"vulnerability"`
}

// Validate validates this finding attribution
func (m *FindingAttribution) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAnalyzerIdentity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttributedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComponent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVulnerability(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var findingAttributionTypeAnalyzerIdentityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["INTERNAL_ANALYZER","OSSINDEX_ANALYZER","NPM_AUDIT_ANALYZER","VULNDB_ANALYZER","NONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		findingAttributionTypeAnalyzerIdentityPropEnum = append(findingAttributionTypeAnalyzerIdentityPropEnum, v)
	}
}

const (

	// FindingAttributionAnalyzerIdentityINTERNALANALYZER captures enum value "INTERNAL_ANALYZER"
	FindingAttributionAnalyzerIdentityINTERNALANALYZER string = "INTERNAL_ANALYZER"

	// FindingAttributionAnalyzerIdentityOSSINDEXANALYZER captures enum value "OSSINDEX_ANALYZER"
	FindingAttributionAnalyzerIdentityOSSINDEXANALYZER string = "OSSINDEX_ANALYZER"

	// FindingAttributionAnalyzerIdentityNPMAUDITANALYZER captures enum value "NPM_AUDIT_ANALYZER"
	FindingAttributionAnalyzerIdentityNPMAUDITANALYZER string = "NPM_AUDIT_ANALYZER"

	// FindingAttributionAnalyzerIdentityVULNDBANALYZER captures enum value "VULNDB_ANALYZER"
	FindingAttributionAnalyzerIdentityVULNDBANALYZER string = "VULNDB_ANALYZER"

	// FindingAttributionAnalyzerIdentityNONE captures enum value "NONE"
	FindingAttributionAnalyzerIdentityNONE string = "NONE"
)

// prop value enum
func (m *FindingAttribution) validateAnalyzerIdentityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, findingAttributionTypeAnalyzerIdentityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FindingAttribution) validateAnalyzerIdentity(formats strfmt.Registry) error {
	if swag.IsZero(m.AnalyzerIdentity) { // not required
		return nil
	}

	// value enum
	if err := m.validateAnalyzerIdentityEnum("analyzerIdentity", "body", m.AnalyzerIdentity); err != nil {
		return err
	}

	return nil
}

func (m *FindingAttribution) validateAttributedOn(formats strfmt.Registry) error {

	if err := validate.Required("attributedOn", "body", m.AttributedOn); err != nil {
		return err
	}

	if err := validate.FormatOf("attributedOn", "body", "date-time", m.AttributedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *FindingAttribution) validateComponent(formats strfmt.Registry) error {

	if err := validate.Required("component", "body", m.Component); err != nil {
		return err
	}

	if m.Component != nil {
		if err := m.Component.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("component")
			}
			return err
		}
	}

	return nil
}

func (m *FindingAttribution) validateUUID(formats strfmt.Registry) error {

	if err := validate.Required("uuid", "body", m.UUID); err != nil {
		return err
	}

	if err := validate.FormatOf("uuid", "body", "uuid", m.UUID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *FindingAttribution) validateVulnerability(formats strfmt.Registry) error {

	if err := validate.Required("vulnerability", "body", m.Vulnerability); err != nil {
		return err
	}

	if m.Vulnerability != nil {
		if err := m.Vulnerability.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vulnerability")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this finding attribution based on the context it is used
func (m *FindingAttribution) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateComponent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVulnerability(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FindingAttribution) contextValidateComponent(ctx context.Context, formats strfmt.Registry) error {

	if m.Component != nil {
		if err := m.Component.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("component")
			}
			return err
		}
	}

	return nil
}

func (m *FindingAttribution) contextValidateVulnerability(ctx context.Context, formats strfmt.Registry) error {

	if m.Vulnerability != nil {
		if err := m.Vulnerability.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vulnerability")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FindingAttribution) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FindingAttribution) UnmarshalBinary(b []byte) error {
	var res FindingAttribution
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
