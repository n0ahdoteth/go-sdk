// Code generated by go-swagger; DO NOT EDIT.

package private_key_tags

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

// NewDeletePrivateKeyTagsParams creates a new DeletePrivateKeyTagsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeletePrivateKeyTagsParams() *DeletePrivateKeyTagsParams {
	return &DeletePrivateKeyTagsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePrivateKeyTagsParamsWithTimeout creates a new DeletePrivateKeyTagsParams object
// with the ability to set a timeout on a request.
func NewDeletePrivateKeyTagsParamsWithTimeout(timeout time.Duration) *DeletePrivateKeyTagsParams {
	return &DeletePrivateKeyTagsParams{
		timeout: timeout,
	}
}

// NewDeletePrivateKeyTagsParamsWithContext creates a new DeletePrivateKeyTagsParams object
// with the ability to set a context for a request.
func NewDeletePrivateKeyTagsParamsWithContext(ctx context.Context) *DeletePrivateKeyTagsParams {
	return &DeletePrivateKeyTagsParams{
		Context: ctx,
	}
}

// NewDeletePrivateKeyTagsParamsWithHTTPClient creates a new DeletePrivateKeyTagsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeletePrivateKeyTagsParamsWithHTTPClient(client *http.Client) *DeletePrivateKeyTagsParams {
	return &DeletePrivateKeyTagsParams{
		HTTPClient: client,
	}
}

/*
DeletePrivateKeyTagsParams contains all the parameters to send to the API endpoint

	for the delete private key tags operation.

	Typically these are written to a http.Request.
*/
type DeletePrivateKeyTagsParams struct {

	// Body.
	Body *models.DeletePrivateKeyTagsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete private key tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePrivateKeyTagsParams) WithDefaults() *DeletePrivateKeyTagsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete private key tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePrivateKeyTagsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete private key tags params
func (o *DeletePrivateKeyTagsParams) WithTimeout(timeout time.Duration) *DeletePrivateKeyTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete private key tags params
func (o *DeletePrivateKeyTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete private key tags params
func (o *DeletePrivateKeyTagsParams) WithContext(ctx context.Context) *DeletePrivateKeyTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete private key tags params
func (o *DeletePrivateKeyTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete private key tags params
func (o *DeletePrivateKeyTagsParams) WithHTTPClient(client *http.Client) *DeletePrivateKeyTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete private key tags params
func (o *DeletePrivateKeyTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete private key tags params
func (o *DeletePrivateKeyTagsParams) WithBody(body *models.DeletePrivateKeyTagsRequest) *DeletePrivateKeyTagsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete private key tags params
func (o *DeletePrivateKeyTagsParams) SetBody(body *models.DeletePrivateKeyTagsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePrivateKeyTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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