// Code generated by go-swagger; DO NOT EDIT.

package private_keys

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

// NewImportPrivateKeyParams creates a new ImportPrivateKeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewImportPrivateKeyParams() *ImportPrivateKeyParams {
	return &ImportPrivateKeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewImportPrivateKeyParamsWithTimeout creates a new ImportPrivateKeyParams object
// with the ability to set a timeout on a request.
func NewImportPrivateKeyParamsWithTimeout(timeout time.Duration) *ImportPrivateKeyParams {
	return &ImportPrivateKeyParams{
		timeout: timeout,
	}
}

// NewImportPrivateKeyParamsWithContext creates a new ImportPrivateKeyParams object
// with the ability to set a context for a request.
func NewImportPrivateKeyParamsWithContext(ctx context.Context) *ImportPrivateKeyParams {
	return &ImportPrivateKeyParams{
		Context: ctx,
	}
}

// NewImportPrivateKeyParamsWithHTTPClient creates a new ImportPrivateKeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewImportPrivateKeyParamsWithHTTPClient(client *http.Client) *ImportPrivateKeyParams {
	return &ImportPrivateKeyParams{
		HTTPClient: client,
	}
}

/*
ImportPrivateKeyParams contains all the parameters to send to the API endpoint

	for the import private key operation.

	Typically these are written to a http.Request.
*/
type ImportPrivateKeyParams struct {

	// Body.
	Body *models.ImportPrivateKeyRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the import private key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImportPrivateKeyParams) WithDefaults() *ImportPrivateKeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the import private key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImportPrivateKeyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the import private key params
func (o *ImportPrivateKeyParams) WithTimeout(timeout time.Duration) *ImportPrivateKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the import private key params
func (o *ImportPrivateKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the import private key params
func (o *ImportPrivateKeyParams) WithContext(ctx context.Context) *ImportPrivateKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the import private key params
func (o *ImportPrivateKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the import private key params
func (o *ImportPrivateKeyParams) WithHTTPClient(client *http.Client) *ImportPrivateKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the import private key params
func (o *ImportPrivateKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the import private key params
func (o *ImportPrivateKeyParams) WithBody(body *models.ImportPrivateKeyRequest) *ImportPrivateKeyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the import private key params
func (o *ImportPrivateKeyParams) SetBody(body *models.ImportPrivateKeyRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ImportPrivateKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
