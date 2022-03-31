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

	"github.com/cisco-open/kubei/api/server/models"
)

// GetPackagesIDApplicationResourcesHandlerFunc turns a function with the right signature into a get packages ID application resources handler
type GetPackagesIDApplicationResourcesHandlerFunc func(GetPackagesIDApplicationResourcesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPackagesIDApplicationResourcesHandlerFunc) Handle(params GetPackagesIDApplicationResourcesParams) middleware.Responder {
	return fn(params)
}

// GetPackagesIDApplicationResourcesHandler interface for that can handle valid get packages ID application resources params
type GetPackagesIDApplicationResourcesHandler interface {
	Handle(GetPackagesIDApplicationResourcesParams) middleware.Responder
}

// NewGetPackagesIDApplicationResources creates a new http.Handler for the get packages ID application resources operation
func NewGetPackagesIDApplicationResources(ctx *middleware.Context, handler GetPackagesIDApplicationResourcesHandler) *GetPackagesIDApplicationResources {
	return &GetPackagesIDApplicationResources{Context: ctx, Handler: handler}
}

/* GetPackagesIDApplicationResources swagger:route GET /packages/{id}/applicationResources getPackagesIdApplicationResources

Get Package application resources

*/
type GetPackagesIDApplicationResources struct {
	Context *middleware.Context
	Handler GetPackagesIDApplicationResourcesHandler
}

func (o *GetPackagesIDApplicationResources) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetPackagesIDApplicationResourcesParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetPackagesIDApplicationResourcesOKBody get packages ID application resources o k body
//
// swagger:model GetPackagesIDApplicationResourcesOKBody
type GetPackagesIDApplicationResourcesOKBody struct {

	// List of package application resources in the given filters and page. List length must be lower or equal to pageSize
	Items []*models.PackageApplicationResources `json:"items"`

	// Total package application resources count under the given filters
	// Required: true
	Total *int64 `json:"total"`
}

// Validate validates this get packages ID application resources o k body
func (o *GetPackagesIDApplicationResourcesOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetPackagesIDApplicationResourcesOKBody) validateItems(formats strfmt.Registry) error {
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
					return ve.ValidateName("getPackagesIdApplicationResourcesOK" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetPackagesIDApplicationResourcesOKBody) validateTotal(formats strfmt.Registry) error {

	if err := validate.Required("getPackagesIdApplicationResourcesOK"+"."+"total", "body", o.Total); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get packages ID application resources o k body based on the context it is used
func (o *GetPackagesIDApplicationResourcesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPackagesIDApplicationResourcesOKBody) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Items); i++ {

		if o.Items[i] != nil {
			if err := o.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getPackagesIdApplicationResourcesOK" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPackagesIDApplicationResourcesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPackagesIDApplicationResourcesOKBody) UnmarshalBinary(b []byte) error {
	var res GetPackagesIDApplicationResourcesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
