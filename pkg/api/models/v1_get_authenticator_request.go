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

// V1GetAuthenticatorRequest v1 get authenticator request
//
// swagger:model v1GetAuthenticatorRequest
type V1GetAuthenticatorRequest struct {

	// Unique identifier for a given Authenticator.
	// Required: true
	AuthenticatorID *string `json:"authenticatorId"`

	// Unique identifier for a given Organization.
	// Required: true
	OrganizationID *string `json:"organizationId"`
}

// Validate validates this v1 get authenticator request
func (m *V1GetAuthenticatorRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthenticatorID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganizationID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1GetAuthenticatorRequest) validateAuthenticatorID(formats strfmt.Registry) error {

	if err := validate.Required("authenticatorId", "body", m.AuthenticatorID); err != nil {
		return err
	}

	return nil
}

func (m *V1GetAuthenticatorRequest) validateOrganizationID(formats strfmt.Registry) error {

	if err := validate.Required("organizationId", "body", m.OrganizationID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 get authenticator request based on context it is used
func (m *V1GetAuthenticatorRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1GetAuthenticatorRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1GetAuthenticatorRequest) UnmarshalBinary(b []byte) error {
	var res V1GetAuthenticatorRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
