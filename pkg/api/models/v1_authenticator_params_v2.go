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

// V1AuthenticatorParamsV2 v1 authenticator params v2
//
// swagger:model v1AuthenticatorParamsV2
type V1AuthenticatorParamsV2 struct {

	// The attestation that proves custody of the authenticator and provides metadata about it.
	// Required: true
	Attestation *V1Attestation `json:"attestation"`

	// Human-readable name for an Authenticator.
	// Required: true
	AuthenticatorName *string `json:"authenticatorName"`

	// Challenge presented for authentication purposes.
	// Required: true
	Challenge *string `json:"challenge"`
}

// Validate validates this v1 authenticator params v2
func (m *V1AuthenticatorParamsV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttestation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthenticatorName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChallenge(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AuthenticatorParamsV2) validateAttestation(formats strfmt.Registry) error {

	if err := validate.Required("attestation", "body", m.Attestation); err != nil {
		return err
	}

	if m.Attestation != nil {
		if err := m.Attestation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attestation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attestation")
			}
			return err
		}
	}

	return nil
}

func (m *V1AuthenticatorParamsV2) validateAuthenticatorName(formats strfmt.Registry) error {

	if err := validate.Required("authenticatorName", "body", m.AuthenticatorName); err != nil {
		return err
	}

	return nil
}

func (m *V1AuthenticatorParamsV2) validateChallenge(formats strfmt.Registry) error {

	if err := validate.Required("challenge", "body", m.Challenge); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 authenticator params v2 based on the context it is used
func (m *V1AuthenticatorParamsV2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttestation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AuthenticatorParamsV2) contextValidateAttestation(ctx context.Context, formats strfmt.Registry) error {

	if m.Attestation != nil {
		if err := m.Attestation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attestation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attestation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1AuthenticatorParamsV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AuthenticatorParamsV2) UnmarshalBinary(b []byte) error {
	var res V1AuthenticatorParamsV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
