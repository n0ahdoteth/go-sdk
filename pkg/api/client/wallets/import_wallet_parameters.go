// Code generated by go-swagger; DO NOT EDIT.

package wallets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/tkhq/go-sdk/pkg/api/models"
)

// NewImportWalletParams creates a new ImportWalletParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewImportWalletParams() *ImportWalletParams {
	return &ImportWalletParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewImportWalletParamsWithTimeout creates a new ImportWalletParams object
// with the ability to set a timeout on a request.
func NewImportWalletParamsWithTimeout(timeout time.Duration) *ImportWalletParams {
	return &ImportWalletParams{
		timeout: timeout,
	}
}

// NewImportWalletParamsWithContext creates a new ImportWalletParams object
// with the ability to set a context for a request.
func NewImportWalletParamsWithContext(ctx context.Context) *ImportWalletParams {
	return &ImportWalletParams{
		Context: ctx,
	}
}

// NewImportWalletParamsWithHTTPClient creates a new ImportWalletParams object
// with the ability to set a custom HTTPClient for a request.
func NewImportWalletParamsWithHTTPClient(client *http.Client) *ImportWalletParams {
	return &ImportWalletParams{
		HTTPClient: client,
	}
}

/*
ImportWalletParams contains all the parameters to send to the API endpoint

	for the import wallet operation.

	Typically these are written to a http.Request.
*/
type ImportWalletParams struct {

	// Body.
	Body *models.ImportWalletRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the import wallet params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImportWalletParams) WithDefaults() *ImportWalletParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the import wallet params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImportWalletParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the import wallet params
func (o *ImportWalletParams) WithTimeout(timeout time.Duration) *ImportWalletParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the import wallet params
func (o *ImportWalletParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the import wallet params
func (o *ImportWalletParams) WithContext(ctx context.Context) *ImportWalletParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the import wallet params
func (o *ImportWalletParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the import wallet params
func (o *ImportWalletParams) WithHTTPClient(client *http.Client) *ImportWalletParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the import wallet params
func (o *ImportWalletParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the import wallet params
func (o *ImportWalletParams) WithBody(body *models.ImportWalletRequest) *ImportWalletParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the import wallet params
func (o *ImportWalletParams) SetBody(body *models.ImportWalletRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ImportWalletParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
