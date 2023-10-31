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

// SignTransactionIntent sign transaction intent
//
// swagger:model SignTransactionIntent
type SignTransactionIntent struct {

	// Unique identifier for a given Private Key.
	// Required: true
	PrivateKeyID *string `json:"privateKeyId"`

	// type
	// Required: true
	Type *TransactionType `json:"type"`

	// Raw unsigned transaction to be signed by a particular Private Key.
	// Required: true
	UnsignedTransaction *string `json:"unsignedTransaction"`
}

// Validate validates this sign transaction intent
func (m *SignTransactionIntent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrivateKeyID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnsignedTransaction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SignTransactionIntent) validatePrivateKeyID(formats strfmt.Registry) error {

	if err := validate.Required("privateKeyId", "body", m.PrivateKeyID); err != nil {
		return err
	}

	return nil
}

func (m *SignTransactionIntent) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *SignTransactionIntent) validateUnsignedTransaction(formats strfmt.Registry) error {

	if err := validate.Required("unsignedTransaction", "body", m.UnsignedTransaction); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this sign transaction intent based on the context it is used
func (m *SignTransactionIntent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SignTransactionIntent) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {

		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SignTransactionIntent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SignTransactionIntent) UnmarshalBinary(b []byte) error {
	var res SignTransactionIntent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
