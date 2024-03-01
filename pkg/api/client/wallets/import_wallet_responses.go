// Code generated by go-swagger; DO NOT EDIT.

package wallets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tkhq/go-sdk/pkg/api/models"
)

// ImportWalletReader is a Reader for the ImportWallet structure.
type ImportWalletReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImportWalletReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImportWalletOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /public/v1/submit/import_wallet] ImportWallet", response, response.Code())
	}
}

// NewImportWalletOK creates a ImportWalletOK with default headers values
func NewImportWalletOK() *ImportWalletOK {
	return &ImportWalletOK{}
}

/*
ImportWalletOK describes a response with status code 200, with default header values.

A successful response.
*/
type ImportWalletOK struct {
	Payload *models.ActivityResponse
}

// IsSuccess returns true when this import wallet o k response has a 2xx status code
func (o *ImportWalletOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this import wallet o k response has a 3xx status code
func (o *ImportWalletOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this import wallet o k response has a 4xx status code
func (o *ImportWalletOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this import wallet o k response has a 5xx status code
func (o *ImportWalletOK) IsServerError() bool {
	return false
}

// IsCode returns true when this import wallet o k response a status code equal to that given
func (o *ImportWalletOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the import wallet o k response
func (o *ImportWalletOK) Code() int {
	return 200
}

func (o *ImportWalletOK) Error() string {
	return fmt.Sprintf("[POST /public/v1/submit/import_wallet][%d] importWalletOK  %+v", 200, o.Payload)
}

func (o *ImportWalletOK) String() string {
	return fmt.Sprintf("[POST /public/v1/submit/import_wallet][%d] importWalletOK  %+v", 200, o.Payload)
}

func (o *ImportWalletOK) GetPayload() *models.ActivityResponse {
	return o.Payload
}

func (o *ImportWalletOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActivityResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
