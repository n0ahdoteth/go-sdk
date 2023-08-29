// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1GetAuthenticatorResponse v1 get authenticator response
//
// swagger:model v1GetAuthenticatorResponse
type V1GetAuthenticatorResponse struct {

	// An authenticator.
	// Required: true
	Authenticator *V1Authenticator `json:"authenticator"`
}

// Validate validates this v1 get authenticator response
func (m *V1GetAuthenticatorResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthenticator(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1GetAuthenticatorResponse) validateAuthenticator(formats strfmt.Registry) error {

	if err := validate.Required("authenticator", "body", m.Authenticator); err != nil {
		return err
	}

	if m.Authenticator != nil {
		if err := m.Authenticator.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authenticator")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("authenticator")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 get authenticator response based on the context it is used
func (m *V1GetAuthenticatorResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuthenticator(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1GetAuthenticatorResponse) contextValidateAuthenticator(ctx context.Context, formats strfmt.Registry) error {

	if m.Authenticator != nil {
		if err := m.Authenticator.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authenticator")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("authenticator")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1GetAuthenticatorResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1GetAuthenticatorResponse) UnmarshalBinary(b []byte) error {
	var res V1GetAuthenticatorResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
