// Code generated by go-swagger; DO NOT EDIT.

package private_key_tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tkhq/go-sdk/pkg/api/models"
)

// PublicAPIServiceUpdatePrivateKeyTagReader is a Reader for the PublicAPIServiceUpdatePrivateKeyTag structure.
type PublicAPIServiceUpdatePrivateKeyTagReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicAPIServiceUpdatePrivateKeyTagReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPublicAPIServiceUpdatePrivateKeyTagOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPublicAPIServiceUpdatePrivateKeyTagDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPublicAPIServiceUpdatePrivateKeyTagOK creates a PublicAPIServiceUpdatePrivateKeyTagOK with default headers values
func NewPublicAPIServiceUpdatePrivateKeyTagOK() *PublicAPIServiceUpdatePrivateKeyTagOK {
	return &PublicAPIServiceUpdatePrivateKeyTagOK{}
}

/*
PublicAPIServiceUpdatePrivateKeyTagOK describes a response with status code 200, with default header values.

A successful response.
*/
type PublicAPIServiceUpdatePrivateKeyTagOK struct {
	Payload *models.V1ActivityResponse
}

// IsSuccess returns true when this public Api service update private key tag o k response has a 2xx status code
func (o *PublicAPIServiceUpdatePrivateKeyTagOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this public Api service update private key tag o k response has a 3xx status code
func (o *PublicAPIServiceUpdatePrivateKeyTagOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this public Api service update private key tag o k response has a 4xx status code
func (o *PublicAPIServiceUpdatePrivateKeyTagOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this public Api service update private key tag o k response has a 5xx status code
func (o *PublicAPIServiceUpdatePrivateKeyTagOK) IsServerError() bool {
	return false
}

// IsCode returns true when this public Api service update private key tag o k response a status code equal to that given
func (o *PublicAPIServiceUpdatePrivateKeyTagOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the public Api service update private key tag o k response
func (o *PublicAPIServiceUpdatePrivateKeyTagOK) Code() int {
	return 200
}

func (o *PublicAPIServiceUpdatePrivateKeyTagOK) Error() string {
	return fmt.Sprintf("[POST /public/v1/submit/update_private_key_tag][%d] publicApiServiceUpdatePrivateKeyTagOK  %+v", 200, o.Payload)
}

func (o *PublicAPIServiceUpdatePrivateKeyTagOK) String() string {
	return fmt.Sprintf("[POST /public/v1/submit/update_private_key_tag][%d] publicApiServiceUpdatePrivateKeyTagOK  %+v", 200, o.Payload)
}

func (o *PublicAPIServiceUpdatePrivateKeyTagOK) GetPayload() *models.V1ActivityResponse {
	return o.Payload
}

func (o *PublicAPIServiceUpdatePrivateKeyTagOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ActivityResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicAPIServiceUpdatePrivateKeyTagDefault creates a PublicAPIServiceUpdatePrivateKeyTagDefault with default headers values
func NewPublicAPIServiceUpdatePrivateKeyTagDefault(code int) *PublicAPIServiceUpdatePrivateKeyTagDefault {
	return &PublicAPIServiceUpdatePrivateKeyTagDefault{
		_statusCode: code,
	}
}

/*
PublicAPIServiceUpdatePrivateKeyTagDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PublicAPIServiceUpdatePrivateKeyTagDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this public Api service update private key tag default response has a 2xx status code
func (o *PublicAPIServiceUpdatePrivateKeyTagDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this public Api service update private key tag default response has a 3xx status code
func (o *PublicAPIServiceUpdatePrivateKeyTagDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this public Api service update private key tag default response has a 4xx status code
func (o *PublicAPIServiceUpdatePrivateKeyTagDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this public Api service update private key tag default response has a 5xx status code
func (o *PublicAPIServiceUpdatePrivateKeyTagDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this public Api service update private key tag default response a status code equal to that given
func (o *PublicAPIServiceUpdatePrivateKeyTagDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the public Api service update private key tag default response
func (o *PublicAPIServiceUpdatePrivateKeyTagDefault) Code() int {
	return o._statusCode
}

func (o *PublicAPIServiceUpdatePrivateKeyTagDefault) Error() string {
	return fmt.Sprintf("[POST /public/v1/submit/update_private_key_tag][%d] PublicApiService_UpdatePrivateKeyTag default  %+v", o._statusCode, o.Payload)
}

func (o *PublicAPIServiceUpdatePrivateKeyTagDefault) String() string {
	return fmt.Sprintf("[POST /public/v1/submit/update_private_key_tag][%d] PublicApiService_UpdatePrivateKeyTag default  %+v", o._statusCode, o.Payload)
}

func (o *PublicAPIServiceUpdatePrivateKeyTagDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *PublicAPIServiceUpdatePrivateKeyTagDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
