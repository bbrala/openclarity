// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/openclarity/kubeclarity/api/server/models"
)

// GetApplicationResourcesHandlerFunc turns a function with the right signature into a get application resources handler
type GetApplicationResourcesHandlerFunc func(GetApplicationResourcesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetApplicationResourcesHandlerFunc) Handle(params GetApplicationResourcesParams) middleware.Responder {
	return fn(params)
}

// GetApplicationResourcesHandler interface for that can handle valid get application resources params
type GetApplicationResourcesHandler interface {
	Handle(GetApplicationResourcesParams) middleware.Responder
}

// NewGetApplicationResources creates a new http.Handler for the get application resources operation
func NewGetApplicationResources(ctx *middleware.Context, handler GetApplicationResourcesHandler) *GetApplicationResources {
	return &GetApplicationResources{Context: ctx, Handler: handler}
}

/* GetApplicationResources swagger:route GET /applicationResources getApplicationResources

Get application resources

*/
type GetApplicationResources struct {
	Context *middleware.Context
	Handler GetApplicationResourcesHandler
}

func (o *GetApplicationResources) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetApplicationResourcesParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetApplicationResourcesOKBody get application resources o k body
//
// swagger:model GetApplicationResourcesOKBody
type GetApplicationResourcesOKBody struct {

	// List of application resources in the given filters and page. List length must be lower or equal to pageSize
	Items []*models.ApplicationResource `json:"items"`

	// Total application resources count under the given filters
	// Required: true
	Total *int64 `json:"total"`
}

// Validate validates this get application resources o k body
func (o *GetApplicationResourcesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTotal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetApplicationResourcesOKBody) validateItems(formats strfmt.Registry) error {
	if swag.IsZero(o.Items) { // not required
		return nil
	}

	for i := 0; i < len(o.Items); i++ {
		if swag.IsZero(o.Items[i]) { // not required
			continue
		}

		if o.Items[i] != nil {
			if err := o.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getApplicationResourcesOK" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetApplicationResourcesOKBody) validateTotal(formats strfmt.Registry) error {

	if err := validate.Required("getApplicationResourcesOK"+"."+"total", "body", o.Total); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get application resources o k body based on the context it is used
func (o *GetApplicationResourcesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetApplicationResourcesOKBody) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Items); i++ {

		if o.Items[i] != nil {
			if err := o.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getApplicationResourcesOK" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetApplicationResourcesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetApplicationResourcesOKBody) UnmarshalBinary(b []byte) error {
	var res GetApplicationResourcesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
