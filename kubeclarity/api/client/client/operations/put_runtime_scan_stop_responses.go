// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openclarity/kubeclarity/api/client/models"
)

// PutRuntimeScanStopReader is a Reader for the PutRuntimeScanStop structure.
type PutRuntimeScanStopReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutRuntimeScanStopReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPutRuntimeScanStopCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutRuntimeScanStopBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPutRuntimeScanStopDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutRuntimeScanStopCreated creates a PutRuntimeScanStopCreated with default headers values
func NewPutRuntimeScanStopCreated() *PutRuntimeScanStopCreated {
	return &PutRuntimeScanStopCreated{}
}

/* PutRuntimeScanStopCreated describes a response with status code 201, with default header values.

Success
*/
type PutRuntimeScanStopCreated struct {
	Payload interface{}
}

func (o *PutRuntimeScanStopCreated) Error() string {
	return fmt.Sprintf("[PUT /runtime/scan/stop][%d] putRuntimeScanStopCreated  %+v", 201, o.Payload)
}
func (o *PutRuntimeScanStopCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *PutRuntimeScanStopCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRuntimeScanStopBadRequest creates a PutRuntimeScanStopBadRequest with default headers values
func NewPutRuntimeScanStopBadRequest() *PutRuntimeScanStopBadRequest {
	return &PutRuntimeScanStopBadRequest{}
}

/* PutRuntimeScanStopBadRequest describes a response with status code 400, with default header values.

Scan failed to stop
*/
type PutRuntimeScanStopBadRequest struct {
	Payload string
}

func (o *PutRuntimeScanStopBadRequest) Error() string {
	return fmt.Sprintf("[PUT /runtime/scan/stop][%d] putRuntimeScanStopBadRequest  %+v", 400, o.Payload)
}
func (o *PutRuntimeScanStopBadRequest) GetPayload() string {
	return o.Payload
}

func (o *PutRuntimeScanStopBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRuntimeScanStopDefault creates a PutRuntimeScanStopDefault with default headers values
func NewPutRuntimeScanStopDefault(code int) *PutRuntimeScanStopDefault {
	return &PutRuntimeScanStopDefault{
		_statusCode: code,
	}
}

/* PutRuntimeScanStopDefault describes a response with status code -1, with default header values.

unknown error
*/
type PutRuntimeScanStopDefault struct {
	_statusCode int

	Payload *models.APIResponse
}

// Code gets the status code for the put runtime scan stop default response
func (o *PutRuntimeScanStopDefault) Code() int {
	return o._statusCode
}

func (o *PutRuntimeScanStopDefault) Error() string {
	return fmt.Sprintf("[PUT /runtime/scan/stop][%d] PutRuntimeScanStop default  %+v", o._statusCode, o.Payload)
}
func (o *PutRuntimeScanStopDefault) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *PutRuntimeScanStopDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
