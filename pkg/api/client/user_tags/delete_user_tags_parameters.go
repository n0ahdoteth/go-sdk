// Code generated by go-swagger; DO NOT EDIT.

package user_tags

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

// NewDeleteUserTagsParams creates a new DeleteUserTagsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteUserTagsParams() *DeleteUserTagsParams {
	return &DeleteUserTagsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUserTagsParamsWithTimeout creates a new DeleteUserTagsParams object
// with the ability to set a timeout on a request.
func NewDeleteUserTagsParamsWithTimeout(timeout time.Duration) *DeleteUserTagsParams {
	return &DeleteUserTagsParams{
		timeout: timeout,
	}
}

// NewDeleteUserTagsParamsWithContext creates a new DeleteUserTagsParams object
// with the ability to set a context for a request.
func NewDeleteUserTagsParamsWithContext(ctx context.Context) *DeleteUserTagsParams {
	return &DeleteUserTagsParams{
		Context: ctx,
	}
}

// NewDeleteUserTagsParamsWithHTTPClient creates a new DeleteUserTagsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteUserTagsParamsWithHTTPClient(client *http.Client) *DeleteUserTagsParams {
	return &DeleteUserTagsParams{
		HTTPClient: client,
	}
}

/*
DeleteUserTagsParams contains all the parameters to send to the API endpoint

	for the delete user tags operation.

	Typically these are written to a http.Request.
*/
type DeleteUserTagsParams struct {

	// Body.
	Body *models.DeleteUserTagsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete user tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUserTagsParams) WithDefaults() *DeleteUserTagsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete user tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUserTagsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete user tags params
func (o *DeleteUserTagsParams) WithTimeout(timeout time.Duration) *DeleteUserTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete user tags params
func (o *DeleteUserTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete user tags params
func (o *DeleteUserTagsParams) WithContext(ctx context.Context) *DeleteUserTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete user tags params
func (o *DeleteUserTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete user tags params
func (o *DeleteUserTagsParams) WithHTTPClient(client *http.Client) *DeleteUserTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete user tags params
func (o *DeleteUserTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete user tags params
func (o *DeleteUserTagsParams) WithBody(body *models.DeleteUserTagsRequest) *DeleteUserTagsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete user tags params
func (o *DeleteUserTagsParams) SetBody(body *models.DeleteUserTagsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUserTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
