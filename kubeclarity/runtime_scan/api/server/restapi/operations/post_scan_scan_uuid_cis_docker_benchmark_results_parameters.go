// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"

	"github.com/cisco-open/kubei/runtime_scan/api/server/models"
)

// NewPostScanScanUUIDCisDockerBenchmarkResultsParams creates a new PostScanScanUUIDCisDockerBenchmarkResultsParams object
//
// There are no default values defined in the spec.
func NewPostScanScanUUIDCisDockerBenchmarkResultsParams() PostScanScanUUIDCisDockerBenchmarkResultsParams {

	return PostScanScanUUIDCisDockerBenchmarkResultsParams{}
}

// PostScanScanUUIDCisDockerBenchmarkResultsParams contains all the bound params for the post scan scan UUID cis docker benchmark results operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostScanScanUUIDCisDockerBenchmarkResults
type PostScanScanUUIDCisDockerBenchmarkResultsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	Body *models.CISDockerBenchmarkScan
	/*
	  Required: true
	  In: path
	*/
	ScanUUID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostScanScanUUIDCisDockerBenchmarkResultsParams() beforehand.
func (o *PostScanScanUUIDCisDockerBenchmarkResultsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.CISDockerBenchmarkScan
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body", ""))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(context.Background())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}
	} else {
		res = append(res, errors.Required("body", "body", ""))
	}

	rScanUUID, rhkScanUUID, _ := route.Params.GetOK("scan-uuid")
	if err := o.bindScanUUID(rScanUUID, rhkScanUUID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindScanUUID binds and validates parameter ScanUUID from path.
func (o *PostScanScanUUIDCisDockerBenchmarkResultsParams) bindScanUUID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("scan-uuid", "path", "strfmt.UUID", raw)
	}
	o.ScanUUID = *(value.(*strfmt.UUID))

	if err := o.validateScanUUID(formats); err != nil {
		return err
	}

	return nil
}

// validateScanUUID carries on validations for parameter ScanUUID
func (o *PostScanScanUUIDCisDockerBenchmarkResultsParams) validateScanUUID(formats strfmt.Registry) error {

	if err := validate.FormatOf("scan-uuid", "path", "uuid", o.ScanUUID.String(), formats); err != nil {
		return err
	}
	return nil
}
