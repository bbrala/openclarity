// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openclarity/kubeclarity/sbom_db/api/server/models"
)

// GetSbomDBResourceHashOKCode is the HTTP code returned for type GetSbomDBResourceHashOK
const GetSbomDBResourceHashOKCode int = 200

/*GetSbomDBResourceHashOK Success

swagger:response getSbomDBResourceHashOK
*/
type GetSbomDBResourceHashOK struct {

	/*
	  In: Body
	*/
	Payload *models.SBOM `json:"body,omitempty"`
}

// NewGetSbomDBResourceHashOK creates GetSbomDBResourceHashOK with default headers values
func NewGetSbomDBResourceHashOK() *GetSbomDBResourceHashOK {

	return &GetSbomDBResourceHashOK{}
}

// WithPayload adds the payload to the get sbom d b resource hash o k response
func (o *GetSbomDBResourceHashOK) WithPayload(payload *models.SBOM) *GetSbomDBResourceHashOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get sbom d b resource hash o k response
func (o *GetSbomDBResourceHashOK) SetPayload(payload *models.SBOM) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSbomDBResourceHashOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSbomDBResourceHashNotFoundCode is the HTTP code returned for type GetSbomDBResourceHashNotFound
const GetSbomDBResourceHashNotFoundCode int = 404

/*GetSbomDBResourceHashNotFound SBOM not found.

swagger:response getSbomDBResourceHashNotFound
*/
type GetSbomDBResourceHashNotFound struct {
}

// NewGetSbomDBResourceHashNotFound creates GetSbomDBResourceHashNotFound with default headers values
func NewGetSbomDBResourceHashNotFound() *GetSbomDBResourceHashNotFound {

	return &GetSbomDBResourceHashNotFound{}
}

// WriteResponse to the client
func (o *GetSbomDBResourceHashNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

/*GetSbomDBResourceHashDefault Unknown error

swagger:response getSbomDBResourceHashDefault
*/
type GetSbomDBResourceHashDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetSbomDBResourceHashDefault creates GetSbomDBResourceHashDefault with default headers values
func NewGetSbomDBResourceHashDefault(code int) *GetSbomDBResourceHashDefault {
	if code <= 0 {
		code = 500
	}

	return &GetSbomDBResourceHashDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get sbom d b resource hash default response
func (o *GetSbomDBResourceHashDefault) WithStatusCode(code int) *GetSbomDBResourceHashDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get sbom d b resource hash default response
func (o *GetSbomDBResourceHashDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get sbom d b resource hash default response
func (o *GetSbomDBResourceHashDefault) WithPayload(payload *models.APIResponse) *GetSbomDBResourceHashDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get sbom d b resource hash default response
func (o *GetSbomDBResourceHashDefault) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSbomDBResourceHashDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
