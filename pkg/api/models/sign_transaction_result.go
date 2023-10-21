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

// SignTransactionResult sign transaction result
//
// swagger:model SignTransactionResult
type SignTransactionResult struct {

	// signed transaction
	// Required: true
	SignedTransaction *string `json:"signedTransaction"`
}

// Validate validates this sign transaction result
func (m *SignTransactionResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSignedTransaction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SignTransactionResult) validateSignedTransaction(formats strfmt.Registry) error {

	if err := validate.Required("signedTransaction", "body", m.SignedTransaction); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this sign transaction result based on context it is used
func (m *SignTransactionResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SignTransactionResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SignTransactionResult) UnmarshalBinary(b []byte) error {
	var res SignTransactionResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}