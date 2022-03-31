// Code generated by go-swagger; DO NOT EDIT.

package finding

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/cisco-open/kubei/shared/pkg/scanner/dependency_track/api/client/models"
)

// GetFindingsByProjectReader is a Reader for the GetFindingsByProject structure.
type GetFindingsByProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFindingsByProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFindingsByProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetFindingsByProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFindingsByProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFindingsByProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFindingsByProjectOK creates a GetFindingsByProjectOK with default headers values
func NewGetFindingsByProjectOK() *GetFindingsByProjectOK {
	return &GetFindingsByProjectOK{}
}

/* GetFindingsByProjectOK describes a response with status code 200, with default header values.

successful operation
*/
type GetFindingsByProjectOK struct {

	/* The total number of findings

	   Format: int64
	*/
	XTotalCount int64

	Payload []*models.Finding
}

func (o *GetFindingsByProjectOK) Error() string {
	return fmt.Sprintf("[GET /v1/finding/project/{uuid}][%d] getFindingsByProjectOK  %+v", 200, o.Payload)
}
func (o *GetFindingsByProjectOK) GetPayload() []*models.Finding {
	return o.Payload
}

func (o *GetFindingsByProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Total-Count
	hdrXTotalCount := response.GetHeader("X-Total-Count")

	if hdrXTotalCount != "" {
		valxTotalCount, err := swag.ConvertInt64(hdrXTotalCount)
		if err != nil {
			return errors.InvalidType("X-Total-Count", "header", "int64", hdrXTotalCount)
		}
		o.XTotalCount = valxTotalCount
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFindingsByProjectUnauthorized creates a GetFindingsByProjectUnauthorized with default headers values
func NewGetFindingsByProjectUnauthorized() *GetFindingsByProjectUnauthorized {
	return &GetFindingsByProjectUnauthorized{}
}

/* GetFindingsByProjectUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetFindingsByProjectUnauthorized struct {
}

func (o *GetFindingsByProjectUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/finding/project/{uuid}][%d] getFindingsByProjectUnauthorized ", 401)
}

func (o *GetFindingsByProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFindingsByProjectForbidden creates a GetFindingsByProjectForbidden with default headers values
func NewGetFindingsByProjectForbidden() *GetFindingsByProjectForbidden {
	return &GetFindingsByProjectForbidden{}
}

/* GetFindingsByProjectForbidden describes a response with status code 403, with default header values.

Access to the specified project is forbidden
*/
type GetFindingsByProjectForbidden struct {
}

func (o *GetFindingsByProjectForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/finding/project/{uuid}][%d] getFindingsByProjectForbidden ", 403)
}

func (o *GetFindingsByProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFindingsByProjectNotFound creates a GetFindingsByProjectNotFound with default headers values
func NewGetFindingsByProjectNotFound() *GetFindingsByProjectNotFound {
	return &GetFindingsByProjectNotFound{}
}

/* GetFindingsByProjectNotFound describes a response with status code 404, with default header values.

The project could not be found
*/
type GetFindingsByProjectNotFound struct {
}

func (o *GetFindingsByProjectNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/finding/project/{uuid}][%d] getFindingsByProjectNotFound ", 404)
}

func (o *GetFindingsByProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
