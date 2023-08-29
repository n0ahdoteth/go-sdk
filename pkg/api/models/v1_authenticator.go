// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1Authenticator v1 authenticator
//
// swagger:model v1Authenticator
type V1Authenticator struct {

	// Identifier indicating the type of the Security Key.
	// Required: true
	Aaguid *string `json:"aaguid"`

	// attestation type
	// Required: true
	AttestationType *string `json:"attestationType"`

	// Unique identifier for a given Authenticator.
	// Required: true
	AuthenticatorID *string `json:"authenticatorId"`

	// Human-readable name for an Authenticator.
	// Required: true
	AuthenticatorName *string `json:"authenticatorName"`

	// created at
	// Required: true
	CreatedAt *V1Timestamp `json:"createdAt"`

	// A User credential that can be used to authenticate to Turnkey.
	// Required: true
	Credential *V1Credential `json:"credential"`

	// Unique identifier for a WebAuthn credential.
	// Required: true
	CredentialID *string `json:"credentialId"`

	// The type of Authenticator device.
	// Required: true
	Model *string `json:"model"`

	// Types of transports that may be used by an Authenticator (e.g., USB, NFC, BLE).
	// Required: true
	Transports []Externaldatav1AuthenticatorTransport `json:"transports"`

	// updated at
	// Required: true
	UpdatedAt *V1Timestamp `json:"updatedAt"`
}

// Validate validates this v1 authenticator
func (m *V1Authenticator) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAaguid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttestationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthenticatorID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthenticatorName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredential(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentialID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransports(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Authenticator) validateAaguid(formats strfmt.Registry) error {

	if err := validate.Required("aaguid", "body", m.Aaguid); err != nil {
		return err
	}

	return nil
}

func (m *V1Authenticator) validateAttestationType(formats strfmt.Registry) error {

	if err := validate.Required("attestationType", "body", m.AttestationType); err != nil {
		return err
	}

	return nil
}

func (m *V1Authenticator) validateAuthenticatorID(formats strfmt.Registry) error {

	if err := validate.Required("authenticatorId", "body", m.AuthenticatorID); err != nil {
		return err
	}

	return nil
}

func (m *V1Authenticator) validateAuthenticatorName(formats strfmt.Registry) error {

	if err := validate.Required("authenticatorName", "body", m.AuthenticatorName); err != nil {
		return err
	}

	return nil
}

func (m *V1Authenticator) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createdAt", "body", m.CreatedAt); err != nil {
		return err
	}

	if m.CreatedAt != nil {
		if err := m.CreatedAt.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createdAt")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createdAt")
			}
			return err
		}
	}

	return nil
}

func (m *V1Authenticator) validateCredential(formats strfmt.Registry) error {

	if err := validate.Required("credential", "body", m.Credential); err != nil {
		return err
	}

	if m.Credential != nil {
		if err := m.Credential.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credential")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("credential")
			}
			return err
		}
	}

	return nil
}

func (m *V1Authenticator) validateCredentialID(formats strfmt.Registry) error {

	if err := validate.Required("credentialId", "body", m.CredentialID); err != nil {
		return err
	}

	return nil
}

func (m *V1Authenticator) validateModel(formats strfmt.Registry) error {

	if err := validate.Required("model", "body", m.Model); err != nil {
		return err
	}

	return nil
}

func (m *V1Authenticator) validateTransports(formats strfmt.Registry) error {

	if err := validate.Required("transports", "body", m.Transports); err != nil {
		return err
	}

	for i := 0; i < len(m.Transports); i++ {

		if err := m.Transports[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("transports" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("transports" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *V1Authenticator) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updatedAt", "body", m.UpdatedAt); err != nil {
		return err
	}

	if m.UpdatedAt != nil {
		if err := m.UpdatedAt.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updatedAt")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updatedAt")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 authenticator based on the context it is used
func (m *V1Authenticator) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCredential(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTransports(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Authenticator) contextValidateCreatedAt(ctx context.Context, formats strfmt.Registry) error {

	if m.CreatedAt != nil {
		if err := m.CreatedAt.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createdAt")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createdAt")
			}
			return err
		}
	}

	return nil
}

func (m *V1Authenticator) contextValidateCredential(ctx context.Context, formats strfmt.Registry) error {

	if m.Credential != nil {
		if err := m.Credential.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credential")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("credential")
			}
			return err
		}
	}

	return nil
}

func (m *V1Authenticator) contextValidateTransports(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Transports); i++ {

		if err := m.Transports[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("transports" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("transports" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *V1Authenticator) contextValidateUpdatedAt(ctx context.Context, formats strfmt.Registry) error {

	if m.UpdatedAt != nil {
		if err := m.UpdatedAt.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updatedAt")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updatedAt")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Authenticator) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Authenticator) UnmarshalBinary(b []byte) error {
	var res V1Authenticator
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
