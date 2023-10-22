// Code generated by go-swagger; DO NOT EDIT.

package private_keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new private keys API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for private keys API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreatePrivateKeys(params *CreatePrivateKeysParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreatePrivateKeysOK, error)

	ExportPrivateKey(params *ExportPrivateKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExportPrivateKeyOK, error)

	GetPrivateKey(params *GetPrivateKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPrivateKeyOK, error)

	GetPrivateKeys(params *GetPrivateKeysParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPrivateKeysOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreatePrivateKeys creates private keys

Create new Private Keys
*/
func (a *Client) CreatePrivateKeys(params *CreatePrivateKeysParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreatePrivateKeysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePrivateKeysParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreatePrivateKeys",
		Method:             "POST",
		PathPattern:        "/public/v1/submit/create_private_keys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreatePrivateKeysReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreatePrivateKeysOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreatePrivateKeys: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ExportPrivateKey exports private key

Exports a Private Key
*/
func (a *Client) ExportPrivateKey(params *ExportPrivateKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExportPrivateKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExportPrivateKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ExportPrivateKey",
		Method:             "POST",
		PathPattern:        "/public/v1/submit/export_private_key",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExportPrivateKeyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ExportPrivateKeyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ExportPrivateKey: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPrivateKey gets private key

Get details about a Private Key
*/
func (a *Client) GetPrivateKey(params *GetPrivateKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPrivateKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrivateKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetPrivateKey",
		Method:             "POST",
		PathPattern:        "/public/v1/query/get_private_key",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPrivateKeyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPrivateKeyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetPrivateKey: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPrivateKeys lists private keys

List all Private Keys within an Organization
*/
func (a *Client) GetPrivateKeys(params *GetPrivateKeysParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPrivateKeysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrivateKeysParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetPrivateKeys",
		Method:             "POST",
		PathPattern:        "/public/v1/query/list_private_keys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPrivateKeysReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPrivateKeysOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetPrivateKeys: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
