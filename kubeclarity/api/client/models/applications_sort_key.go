// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ApplicationsSortKey applications sort key
//
// swagger:model ApplicationsSortKey
type ApplicationsSortKey string

func NewApplicationsSortKey(value ApplicationsSortKey) *ApplicationsSortKey {
	v := value
	return &v
}

const (

	// ApplicationsSortKeyApplicationName captures enum value "applicationName"
	ApplicationsSortKeyApplicationName ApplicationsSortKey = "applicationName"

	// ApplicationsSortKeyApplicationType captures enum value "applicationType"
	ApplicationsSortKeyApplicationType ApplicationsSortKey = "applicationType"

	// ApplicationsSortKeyVulnerabilities captures enum value "vulnerabilities"
	ApplicationsSortKeyVulnerabilities ApplicationsSortKey = "vulnerabilities"

	// ApplicationsSortKeyApplicationResources captures enum value "applicationResources"
	ApplicationsSortKeyApplicationResources ApplicationsSortKey = "applicationResources"

	// ApplicationsSortKeyPackages captures enum value "packages"
	ApplicationsSortKeyPackages ApplicationsSortKey = "packages"
)

// for schema
var applicationsSortKeyEnum []interface{}

func init() {
	var res []ApplicationsSortKey
	if err := json.Unmarshal([]byte(`["applicationName","applicationType","vulnerabilities","applicationResources","packages"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		applicationsSortKeyEnum = append(applicationsSortKeyEnum, v)
	}
}

func (m ApplicationsSortKey) validateApplicationsSortKeyEnum(path, location string, value ApplicationsSortKey) error {
	if err := validate.EnumCase(path, location, value, applicationsSortKeyEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this applications sort key
func (m ApplicationsSortKey) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateApplicationsSortKeyEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this applications sort key based on context it is used
func (m ApplicationsSortKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
