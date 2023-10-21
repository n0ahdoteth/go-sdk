// Code generated by go-swagger; DO NOT EDIT.

package user_recovery

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

// NewInitUserEmailRecoveryParams creates a new InitUserEmailRecoveryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInitUserEmailRecoveryParams() *InitUserEmailRecoveryParams {
	return &InitUserEmailRecoveryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInitUserEmailRecoveryParamsWithTimeout creates a new InitUserEmailRecoveryParams object
// with the ability to set a timeout on a request.
func NewInitUserEmailRecoveryParamsWithTimeout(timeout time.Duration) *InitUserEmailRecoveryParams {
	return &InitUserEmailRecoveryParams{
		timeout: timeout,
	}
}

// NewInitUserEmailRecoveryParamsWithContext creates a new InitUserEmailRecoveryParams object
// with the ability to set a context for a request.
func NewInitUserEmailRecoveryParamsWithContext(ctx context.Context) *InitUserEmailRecoveryParams {
	return &InitUserEmailRecoveryParams{
		Context: ctx,
	}
}

// NewInitUserEmailRecoveryParamsWithHTTPClient creates a new InitUserEmailRecoveryParams object
// with the ability to set a custom HTTPClient for a request.
func NewInitUserEmailRecoveryParamsWithHTTPClient(client *http.Client) *InitUserEmailRecoveryParams {
	return &InitUserEmailRecoveryParams{
		HTTPClient: client,
	}
}

/*
InitUserEmailRecoveryParams contains all the parameters to send to the API endpoint

	for the init user email recovery operation.

	Typically these are written to a http.Request.
*/
type InitUserEmailRecoveryParams struct {

	// Body.
	Body *models.InitUserEmailRecoveryRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the init user email recovery params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InitUserEmailRecoveryParams) WithDefaults() *InitUserEmailRecoveryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the init user email recovery params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InitUserEmailRecoveryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the init user email recovery params
func (o *InitUserEmailRecoveryParams) WithTimeout(timeout time.Duration) *InitUserEmailRecoveryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the init user email recovery params
func (o *InitUserEmailRecoveryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the init user email recovery params
func (o *InitUserEmailRecoveryParams) WithContext(ctx context.Context) *InitUserEmailRecoveryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the init user email recovery params
func (o *InitUserEmailRecoveryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the init user email recovery params
func (o *InitUserEmailRecoveryParams) WithHTTPClient(client *http.Client) *InitUserEmailRecoveryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the init user email recovery params
func (o *InitUserEmailRecoveryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the init user email recovery params
func (o *InitUserEmailRecoveryParams) WithBody(body *models.InitUserEmailRecoveryRequest) *InitUserEmailRecoveryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the init user email recovery params
func (o *InitUserEmailRecoveryParams) SetBody(body *models.InitUserEmailRecoveryRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *InitUserEmailRecoveryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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