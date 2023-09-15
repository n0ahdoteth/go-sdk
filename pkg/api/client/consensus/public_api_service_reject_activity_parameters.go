// Code generated by go-swagger; DO NOT EDIT.

package consensus

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

// NewPublicAPIServiceRejectActivityParams creates a new PublicAPIServiceRejectActivityParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPublicAPIServiceRejectActivityParams() *PublicAPIServiceRejectActivityParams {
	return &PublicAPIServiceRejectActivityParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPublicAPIServiceRejectActivityParamsWithTimeout creates a new PublicAPIServiceRejectActivityParams object
// with the ability to set a timeout on a request.
func NewPublicAPIServiceRejectActivityParamsWithTimeout(timeout time.Duration) *PublicAPIServiceRejectActivityParams {
	return &PublicAPIServiceRejectActivityParams{
		timeout: timeout,
	}
}

// NewPublicAPIServiceRejectActivityParamsWithContext creates a new PublicAPIServiceRejectActivityParams object
// with the ability to set a context for a request.
func NewPublicAPIServiceRejectActivityParamsWithContext(ctx context.Context) *PublicAPIServiceRejectActivityParams {
	return &PublicAPIServiceRejectActivityParams{
		Context: ctx,
	}
}

// NewPublicAPIServiceRejectActivityParamsWithHTTPClient creates a new PublicAPIServiceRejectActivityParams object
// with the ability to set a custom HTTPClient for a request.
func NewPublicAPIServiceRejectActivityParamsWithHTTPClient(client *http.Client) *PublicAPIServiceRejectActivityParams {
	return &PublicAPIServiceRejectActivityParams{
		HTTPClient: client,
	}
}

/*
PublicAPIServiceRejectActivityParams contains all the parameters to send to the API endpoint

	for the public Api service reject activity operation.

	Typically these are written to a http.Request.
*/
type PublicAPIServiceRejectActivityParams struct {

	// Body.
	Body *models.V1RejectActivityRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the public Api service reject activity params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PublicAPIServiceRejectActivityParams) WithDefaults() *PublicAPIServiceRejectActivityParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the public Api service reject activity params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PublicAPIServiceRejectActivityParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the public Api service reject activity params
func (o *PublicAPIServiceRejectActivityParams) WithTimeout(timeout time.Duration) *PublicAPIServiceRejectActivityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the public Api service reject activity params
func (o *PublicAPIServiceRejectActivityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the public Api service reject activity params
func (o *PublicAPIServiceRejectActivityParams) WithContext(ctx context.Context) *PublicAPIServiceRejectActivityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the public Api service reject activity params
func (o *PublicAPIServiceRejectActivityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the public Api service reject activity params
func (o *PublicAPIServiceRejectActivityParams) WithHTTPClient(client *http.Client) *PublicAPIServiceRejectActivityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the public Api service reject activity params
func (o *PublicAPIServiceRejectActivityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the public Api service reject activity params
func (o *PublicAPIServiceRejectActivityParams) WithBody(body *models.V1RejectActivityRequest) *PublicAPIServiceRejectActivityParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the public Api service reject activity params
func (o *PublicAPIServiceRejectActivityParams) SetBody(body *models.V1RejectActivityRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PublicAPIServiceRejectActivityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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