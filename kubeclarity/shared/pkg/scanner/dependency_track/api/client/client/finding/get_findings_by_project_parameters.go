// Code generated by go-swagger; DO NOT EDIT.

package finding

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetFindingsByProjectParams creates a new GetFindingsByProjectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFindingsByProjectParams() *GetFindingsByProjectParams {
	return &GetFindingsByProjectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFindingsByProjectParamsWithTimeout creates a new GetFindingsByProjectParams object
// with the ability to set a timeout on a request.
func NewGetFindingsByProjectParamsWithTimeout(timeout time.Duration) *GetFindingsByProjectParams {
	return &GetFindingsByProjectParams{
		timeout: timeout,
	}
}

// NewGetFindingsByProjectParamsWithContext creates a new GetFindingsByProjectParams object
// with the ability to set a context for a request.
func NewGetFindingsByProjectParamsWithContext(ctx context.Context) *GetFindingsByProjectParams {
	return &GetFindingsByProjectParams{
		Context: ctx,
	}
}

// NewGetFindingsByProjectParamsWithHTTPClient creates a new GetFindingsByProjectParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFindingsByProjectParamsWithHTTPClient(client *http.Client) *GetFindingsByProjectParams {
	return &GetFindingsByProjectParams{
		HTTPClient: client,
	}
}

/* GetFindingsByProjectParams contains all the parameters to send to the API endpoint
   for the get findings by project operation.

   Typically these are written to a http.Request.
*/
type GetFindingsByProjectParams struct {

	/* Suppressed.

	   Optionally includes suppressed findings
	*/
	Suppressed *bool

	// UUID.
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get findings by project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFindingsByProjectParams) WithDefaults() *GetFindingsByProjectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get findings by project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFindingsByProjectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get findings by project params
func (o *GetFindingsByProjectParams) WithTimeout(timeout time.Duration) *GetFindingsByProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get findings by project params
func (o *GetFindingsByProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get findings by project params
func (o *GetFindingsByProjectParams) WithContext(ctx context.Context) *GetFindingsByProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get findings by project params
func (o *GetFindingsByProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get findings by project params
func (o *GetFindingsByProjectParams) WithHTTPClient(client *http.Client) *GetFindingsByProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get findings by project params
func (o *GetFindingsByProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSuppressed adds the suppressed to the get findings by project params
func (o *GetFindingsByProjectParams) WithSuppressed(suppressed *bool) *GetFindingsByProjectParams {
	o.SetSuppressed(suppressed)
	return o
}

// SetSuppressed adds the suppressed to the get findings by project params
func (o *GetFindingsByProjectParams) SetSuppressed(suppressed *bool) {
	o.Suppressed = suppressed
}

// WithUUID adds the uuid to the get findings by project params
func (o *GetFindingsByProjectParams) WithUUID(uuid string) *GetFindingsByProjectParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the get findings by project params
func (o *GetFindingsByProjectParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *GetFindingsByProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Suppressed != nil {

		// query param suppressed
		var qrSuppressed bool

		if o.Suppressed != nil {
			qrSuppressed = *o.Suppressed
		}
		qSuppressed := swag.FormatBool(qrSuppressed)
		if qSuppressed != "" {

			if err := r.SetQueryParam("suppressed", qSuppressed); err != nil {
				return err
			}
		}
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
