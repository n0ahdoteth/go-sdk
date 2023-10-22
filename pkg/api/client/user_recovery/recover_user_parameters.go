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

// NewRecoverUserParams creates a new RecoverUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRecoverUserParams() *RecoverUserParams {
	return &RecoverUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRecoverUserParamsWithTimeout creates a new RecoverUserParams object
// with the ability to set a timeout on a request.
func NewRecoverUserParamsWithTimeout(timeout time.Duration) *RecoverUserParams {
	return &RecoverUserParams{
		timeout: timeout,
	}
}

// NewRecoverUserParamsWithContext creates a new RecoverUserParams object
// with the ability to set a context for a request.
func NewRecoverUserParamsWithContext(ctx context.Context) *RecoverUserParams {
	return &RecoverUserParams{
		Context: ctx,
	}
}

// NewRecoverUserParamsWithHTTPClient creates a new RecoverUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewRecoverUserParamsWithHTTPClient(client *http.Client) *RecoverUserParams {
	return &RecoverUserParams{
		HTTPClient: client,
	}
}

/*
RecoverUserParams contains all the parameters to send to the API endpoint

	for the recover user operation.

	Typically these are written to a http.Request.
*/
type RecoverUserParams struct {

	// Body.
	Body *models.RecoverUserRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the recover user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RecoverUserParams) WithDefaults() *RecoverUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the recover user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RecoverUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the recover user params
func (o *RecoverUserParams) WithTimeout(timeout time.Duration) *RecoverUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the recover user params
func (o *RecoverUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the recover user params
func (o *RecoverUserParams) WithContext(ctx context.Context) *RecoverUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the recover user params
func (o *RecoverUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the recover user params
func (o *RecoverUserParams) WithHTTPClient(client *http.Client) *RecoverUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the recover user params
func (o *RecoverUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the recover user params
func (o *RecoverUserParams) WithBody(body *models.RecoverUserRequest) *RecoverUserParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the recover user params
func (o *RecoverUserParams) SetBody(body *models.RecoverUserRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RecoverUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
