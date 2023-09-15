// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tkhq/go-sdk/pkg/api/models"
)

// PublicAPIServiceCreateSubOrganizationReader is a Reader for the PublicAPIServiceCreateSubOrganization structure.
type PublicAPIServiceCreateSubOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicAPIServiceCreateSubOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPublicAPIServiceCreateSubOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPublicAPIServiceCreateSubOrganizationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPublicAPIServiceCreateSubOrganizationOK creates a PublicAPIServiceCreateSubOrganizationOK with default headers values
func NewPublicAPIServiceCreateSubOrganizationOK() *PublicAPIServiceCreateSubOrganizationOK {
	return &PublicAPIServiceCreateSubOrganizationOK{}
}

/*
PublicAPIServiceCreateSubOrganizationOK describes a response with status code 200, with default header values.

A successful response.
*/
type PublicAPIServiceCreateSubOrganizationOK struct {
	Payload *models.V1ActivityResponse
}

// IsSuccess returns true when this public Api service create sub organization o k response has a 2xx status code
func (o *PublicAPIServiceCreateSubOrganizationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this public Api service create sub organization o k response has a 3xx status code
func (o *PublicAPIServiceCreateSubOrganizationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this public Api service create sub organization o k response has a 4xx status code
func (o *PublicAPIServiceCreateSubOrganizationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this public Api service create sub organization o k response has a 5xx status code
func (o *PublicAPIServiceCreateSubOrganizationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this public Api service create sub organization o k response a status code equal to that given
func (o *PublicAPIServiceCreateSubOrganizationOK) IsCode(code int) bool {
	return code == 200
}

func (o *PublicAPIServiceCreateSubOrganizationOK) Error() string {
	return fmt.Sprintf("[POST /public/v1/submit/create_sub_organization][%d] publicApiServiceCreateSubOrganizationOK  %+v", 200, o.Payload)
}

func (o *PublicAPIServiceCreateSubOrganizationOK) String() string {
	return fmt.Sprintf("[POST /public/v1/submit/create_sub_organization][%d] publicApiServiceCreateSubOrganizationOK  %+v", 200, o.Payload)
}

func (o *PublicAPIServiceCreateSubOrganizationOK) GetPayload() *models.V1ActivityResponse {
	return o.Payload
}

func (o *PublicAPIServiceCreateSubOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ActivityResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicAPIServiceCreateSubOrganizationDefault creates a PublicAPIServiceCreateSubOrganizationDefault with default headers values
func NewPublicAPIServiceCreateSubOrganizationDefault(code int) *PublicAPIServiceCreateSubOrganizationDefault {
	return &PublicAPIServiceCreateSubOrganizationDefault{
		_statusCode: code,
	}
}

/*
PublicAPIServiceCreateSubOrganizationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PublicAPIServiceCreateSubOrganizationDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the public Api service create sub organization default response
func (o *PublicAPIServiceCreateSubOrganizationDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this public Api service create sub organization default response has a 2xx status code
func (o *PublicAPIServiceCreateSubOrganizationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this public Api service create sub organization default response has a 3xx status code
func (o *PublicAPIServiceCreateSubOrganizationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this public Api service create sub organization default response has a 4xx status code
func (o *PublicAPIServiceCreateSubOrganizationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this public Api service create sub organization default response has a 5xx status code
func (o *PublicAPIServiceCreateSubOrganizationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this public Api service create sub organization default response a status code equal to that given
func (o *PublicAPIServiceCreateSubOrganizationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PublicAPIServiceCreateSubOrganizationDefault) Error() string {
	return fmt.Sprintf("[POST /public/v1/submit/create_sub_organization][%d] PublicApiService_CreateSubOrganization default  %+v", o._statusCode, o.Payload)
}

func (o *PublicAPIServiceCreateSubOrganizationDefault) String() string {
	return fmt.Sprintf("[POST /public/v1/submit/create_sub_organization][%d] PublicApiService_CreateSubOrganization default  %+v", o._statusCode, o.Payload)
}

func (o *PublicAPIServiceCreateSubOrganizationDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *PublicAPIServiceCreateSubOrganizationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}