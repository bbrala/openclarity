// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cisco-open/kubei/api/client/models"
)

// GetDashboardPackagesPerLicenseReader is a Reader for the GetDashboardPackagesPerLicense structure.
type GetDashboardPackagesPerLicenseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDashboardPackagesPerLicenseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDashboardPackagesPerLicenseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetDashboardPackagesPerLicenseDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDashboardPackagesPerLicenseOK creates a GetDashboardPackagesPerLicenseOK with default headers values
func NewGetDashboardPackagesPerLicenseOK() *GetDashboardPackagesPerLicenseOK {
	return &GetDashboardPackagesPerLicenseOK{}
}

/* GetDashboardPackagesPerLicenseOK describes a response with status code 200, with default header values.

Success
*/
type GetDashboardPackagesPerLicenseOK struct {
	Payload []*models.PackagesCountPerLicense
}

func (o *GetDashboardPackagesPerLicenseOK) Error() string {
	return fmt.Sprintf("[GET /dashboard/packagesPerLicense][%d] getDashboardPackagesPerLicenseOK  %+v", 200, o.Payload)
}
func (o *GetDashboardPackagesPerLicenseOK) GetPayload() []*models.PackagesCountPerLicense {
	return o.Payload
}

func (o *GetDashboardPackagesPerLicenseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDashboardPackagesPerLicenseDefault creates a GetDashboardPackagesPerLicenseDefault with default headers values
func NewGetDashboardPackagesPerLicenseDefault(code int) *GetDashboardPackagesPerLicenseDefault {
	return &GetDashboardPackagesPerLicenseDefault{
		_statusCode: code,
	}
}

/* GetDashboardPackagesPerLicenseDefault describes a response with status code -1, with default header values.

unknown error
*/
type GetDashboardPackagesPerLicenseDefault struct {
	_statusCode int

	Payload *models.APIResponse
}

// Code gets the status code for the get dashboard packages per license default response
func (o *GetDashboardPackagesPerLicenseDefault) Code() int {
	return o._statusCode
}

func (o *GetDashboardPackagesPerLicenseDefault) Error() string {
	return fmt.Sprintf("[GET /dashboard/packagesPerLicense][%d] GetDashboardPackagesPerLicense default  %+v", o._statusCode, o.Payload)
}
func (o *GetDashboardPackagesPerLicenseDefault) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *GetDashboardPackagesPerLicenseDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
