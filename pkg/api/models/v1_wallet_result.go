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

// V1WalletResult v1 wallet result
//
// swagger:model v1WalletResult
type V1WalletResult struct {

	// A list of account addresses.
	// Required: true
	Addresses []string `json:"addresses"`

	// wallet Id
	// Required: true
	WalletID *string `json:"walletId"`
}

// Validate validates this v1 wallet result
func (m *V1WalletResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddresses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWalletID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1WalletResult) validateAddresses(formats strfmt.Registry) error {

	if err := validate.Required("addresses", "body", m.Addresses); err != nil {
		return err
	}

	return nil
}

func (m *V1WalletResult) validateWalletID(formats strfmt.Registry) error {

	if err := validate.Required("walletId", "body", m.WalletID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 wallet result based on context it is used
func (m *V1WalletResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1WalletResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1WalletResult) UnmarshalBinary(b []byte) error {
	var res V1WalletResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
