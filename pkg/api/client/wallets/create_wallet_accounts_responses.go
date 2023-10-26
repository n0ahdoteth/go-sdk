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

// CreateWalletAccountsReader is a Reader for the CreateWalletAccounts structure.
type CreateWalletAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateWalletAccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateWalletAccountsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateWalletAccountsOK creates a CreateWalletAccountsOK with default headers values
func NewCreateWalletAccountsOK() *CreateWalletAccountsOK {
	return &CreateWalletAccountsOK{}
}

/*
CreateWalletAccountsOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateWalletAccountsOK struct {
	Payload *models.ActivityResponse
}

// IsSuccess returns true when this create wallet accounts o k response has a 2xx status code
func (o *CreateWalletAccountsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create wallet accounts o k response has a 3xx status code
func (o *CreateWalletAccountsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create wallet accounts o k response has a 4xx status code
func (o *CreateWalletAccountsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create wallet accounts o k response has a 5xx status code
func (o *CreateWalletAccountsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create wallet accounts o k response a status code equal to that given
func (o *CreateWalletAccountsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create wallet accounts o k response
func (o *CreateWalletAccountsOK) Code() int {
	return 200
}

func (o *CreateWalletAccountsOK) Error() string {
	return fmt.Sprintf("[POST /public/v1/submit/create_wallet_accounts][%d] createWalletAccountsOK  %+v", 200, o.Payload)
}

func (o *CreateWalletAccountsOK) String() string {
	return fmt.Sprintf("[POST /public/v1/submit/create_wallet_accounts][%d] createWalletAccountsOK  %+v", 200, o.Payload)
}

func (o *CreateWalletAccountsOK) GetPayload() *models.ActivityResponse {
	return o.Payload
}

func (o *CreateWalletAccountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActivityResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
