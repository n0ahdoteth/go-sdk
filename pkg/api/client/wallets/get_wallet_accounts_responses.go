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

// GetWalletAccountsReader is a Reader for the GetWalletAccounts structure.
type GetWalletAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWalletAccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWalletAccountsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /public/v1/query/list_wallet_accounts] GetWalletAccounts", response, response.Code())
	}
}

// NewGetWalletAccountsOK creates a GetWalletAccountsOK with default headers values
func NewGetWalletAccountsOK() *GetWalletAccountsOK {
	return &GetWalletAccountsOK{}
}

/*
GetWalletAccountsOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetWalletAccountsOK struct {
	Payload *models.GetWalletAccountsResponse
}

// IsSuccess returns true when this get wallet accounts o k response has a 2xx status code
func (o *GetWalletAccountsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get wallet accounts o k response has a 3xx status code
func (o *GetWalletAccountsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get wallet accounts o k response has a 4xx status code
func (o *GetWalletAccountsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get wallet accounts o k response has a 5xx status code
func (o *GetWalletAccountsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get wallet accounts o k response a status code equal to that given
func (o *GetWalletAccountsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get wallet accounts o k response
func (o *GetWalletAccountsOK) Code() int {
	return 200
}

func (o *GetWalletAccountsOK) Error() string {
	return fmt.Sprintf("[POST /public/v1/query/list_wallet_accounts][%d] getWalletAccountsOK  %+v", 200, o.Payload)
}

func (o *GetWalletAccountsOK) String() string {
	return fmt.Sprintf("[POST /public/v1/query/list_wallet_accounts][%d] getWalletAccountsOK  %+v", 200, o.Payload)
}

func (o *GetWalletAccountsOK) GetPayload() *models.GetWalletAccountsResponse {
	return o.Payload
}

func (o *GetWalletAccountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetWalletAccountsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
