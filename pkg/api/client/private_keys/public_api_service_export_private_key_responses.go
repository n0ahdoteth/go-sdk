// Code generated by go-swagger; DO NOT EDIT.

package private_keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tkhq/go-sdk/pkg/api/models"
)

// PublicAPIServiceExportPrivateKeyReader is a Reader for the PublicAPIServiceExportPrivateKey structure.
type PublicAPIServiceExportPrivateKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicAPIServiceExportPrivateKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPublicAPIServiceExportPrivateKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPublicAPIServiceExportPrivateKeyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPublicAPIServiceExportPrivateKeyOK creates a PublicAPIServiceExportPrivateKeyOK with default headers values
func NewPublicAPIServiceExportPrivateKeyOK() *PublicAPIServiceExportPrivateKeyOK {
	return &PublicAPIServiceExportPrivateKeyOK{}
}

/*
PublicAPIServiceExportPrivateKeyOK describes a response with status code 200, with default header values.

A successful response.
*/
type PublicAPIServiceExportPrivateKeyOK struct {
	Payload *models.V1ActivityResponse
}

// IsSuccess returns true when this public Api service export private key o k response has a 2xx status code
func (o *PublicAPIServiceExportPrivateKeyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this public Api service export private key o k response has a 3xx status code
func (o *PublicAPIServiceExportPrivateKeyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this public Api service export private key o k response has a 4xx status code
func (o *PublicAPIServiceExportPrivateKeyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this public Api service export private key o k response has a 5xx status code
func (o *PublicAPIServiceExportPrivateKeyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this public Api service export private key o k response a status code equal to that given
func (o *PublicAPIServiceExportPrivateKeyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the public Api service export private key o k response
func (o *PublicAPIServiceExportPrivateKeyOK) Code() int {
	return 200
}

func (o *PublicAPIServiceExportPrivateKeyOK) Error() string {
	return fmt.Sprintf("[POST /public/v1/submit/export_private_key][%d] publicApiServiceExportPrivateKeyOK  %+v", 200, o.Payload)
}

func (o *PublicAPIServiceExportPrivateKeyOK) String() string {
	return fmt.Sprintf("[POST /public/v1/submit/export_private_key][%d] publicApiServiceExportPrivateKeyOK  %+v", 200, o.Payload)
}

func (o *PublicAPIServiceExportPrivateKeyOK) GetPayload() *models.V1ActivityResponse {
	return o.Payload
}

func (o *PublicAPIServiceExportPrivateKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ActivityResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicAPIServiceExportPrivateKeyDefault creates a PublicAPIServiceExportPrivateKeyDefault with default headers values
func NewPublicAPIServiceExportPrivateKeyDefault(code int) *PublicAPIServiceExportPrivateKeyDefault {
	return &PublicAPIServiceExportPrivateKeyDefault{
		_statusCode: code,
	}
}

/*
PublicAPIServiceExportPrivateKeyDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PublicAPIServiceExportPrivateKeyDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this public Api service export private key default response has a 2xx status code
func (o *PublicAPIServiceExportPrivateKeyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this public Api service export private key default response has a 3xx status code
func (o *PublicAPIServiceExportPrivateKeyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this public Api service export private key default response has a 4xx status code
func (o *PublicAPIServiceExportPrivateKeyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this public Api service export private key default response has a 5xx status code
func (o *PublicAPIServiceExportPrivateKeyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this public Api service export private key default response a status code equal to that given
func (o *PublicAPIServiceExportPrivateKeyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the public Api service export private key default response
func (o *PublicAPIServiceExportPrivateKeyDefault) Code() int {
	return o._statusCode
}

func (o *PublicAPIServiceExportPrivateKeyDefault) Error() string {
	return fmt.Sprintf("[POST /public/v1/submit/export_private_key][%d] PublicApiService_ExportPrivateKey default  %+v", o._statusCode, o.Payload)
}

func (o *PublicAPIServiceExportPrivateKeyDefault) String() string {
	return fmt.Sprintf("[POST /public/v1/submit/export_private_key][%d] PublicApiService_ExportPrivateKey default  %+v", o._statusCode, o.Payload)
}

func (o *PublicAPIServiceExportPrivateKeyDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *PublicAPIServiceExportPrivateKeyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}