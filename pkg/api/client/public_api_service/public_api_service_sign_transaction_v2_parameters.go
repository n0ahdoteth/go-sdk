// Code generated by go-swagger; DO NOT EDIT.

package public_api_service

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

// NewPublicAPIServiceSignTransactionV2Params creates a new PublicAPIServiceSignTransactionV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPublicAPIServiceSignTransactionV2Params() *PublicAPIServiceSignTransactionV2Params {
	return &PublicAPIServiceSignTransactionV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewPublicAPIServiceSignTransactionV2ParamsWithTimeout creates a new PublicAPIServiceSignTransactionV2Params object
// with the ability to set a timeout on a request.
func NewPublicAPIServiceSignTransactionV2ParamsWithTimeout(timeout time.Duration) *PublicAPIServiceSignTransactionV2Params {
	return &PublicAPIServiceSignTransactionV2Params{
		timeout: timeout,
	}
}

// NewPublicAPIServiceSignTransactionV2ParamsWithContext creates a new PublicAPIServiceSignTransactionV2Params object
// with the ability to set a context for a request.
func NewPublicAPIServiceSignTransactionV2ParamsWithContext(ctx context.Context) *PublicAPIServiceSignTransactionV2Params {
	return &PublicAPIServiceSignTransactionV2Params{
		Context: ctx,
	}
}

// NewPublicAPIServiceSignTransactionV2ParamsWithHTTPClient creates a new PublicAPIServiceSignTransactionV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewPublicAPIServiceSignTransactionV2ParamsWithHTTPClient(client *http.Client) *PublicAPIServiceSignTransactionV2Params {
	return &PublicAPIServiceSignTransactionV2Params{
		HTTPClient: client,
	}
}

/*
PublicAPIServiceSignTransactionV2Params contains all the parameters to send to the API endpoint

	for the public Api service sign transaction v2 operation.

	Typically these are written to a http.Request.
*/
type PublicAPIServiceSignTransactionV2Params struct {

	// Body.
	Body *models.V1SignTransactionV2Request

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the public Api service sign transaction v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PublicAPIServiceSignTransactionV2Params) WithDefaults() *PublicAPIServiceSignTransactionV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the public Api service sign transaction v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PublicAPIServiceSignTransactionV2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the public Api service sign transaction v2 params
func (o *PublicAPIServiceSignTransactionV2Params) WithTimeout(timeout time.Duration) *PublicAPIServiceSignTransactionV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the public Api service sign transaction v2 params
func (o *PublicAPIServiceSignTransactionV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the public Api service sign transaction v2 params
func (o *PublicAPIServiceSignTransactionV2Params) WithContext(ctx context.Context) *PublicAPIServiceSignTransactionV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the public Api service sign transaction v2 params
func (o *PublicAPIServiceSignTransactionV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the public Api service sign transaction v2 params
func (o *PublicAPIServiceSignTransactionV2Params) WithHTTPClient(client *http.Client) *PublicAPIServiceSignTransactionV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the public Api service sign transaction v2 params
func (o *PublicAPIServiceSignTransactionV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the public Api service sign transaction v2 params
func (o *PublicAPIServiceSignTransactionV2Params) WithBody(body *models.V1SignTransactionV2Request) *PublicAPIServiceSignTransactionV2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the public Api service sign transaction v2 params
func (o *PublicAPIServiceSignTransactionV2Params) SetBody(body *models.V1SignTransactionV2Request) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PublicAPIServiceSignTransactionV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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