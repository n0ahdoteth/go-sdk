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
)

// PrivateKeyResult private key result
//
// swagger:model PrivateKeyResult
type PrivateKeyResult struct {

	// addresses
	Addresses []*ActivityV1Address `json:"addresses"`

	// private key Id
	PrivateKeyID string `json:"privateKeyId,omitempty"`
}

// Validate validates this private key result
func (m *PrivateKeyResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddresses(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrivateKeyResult) validateAddresses(formats strfmt.Registry) error {
	if swag.IsZero(m.Addresses) { // not required
		return nil
	}

	for i := 0; i < len(m.Addresses); i++ {
		if swag.IsZero(m.Addresses[i]) { // not required
			continue
		}

		if m.Addresses[i] != nil {
			if err := m.Addresses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("addresses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("addresses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this private key result based on the context it is used
func (m *PrivateKeyResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddresses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrivateKeyResult) contextValidateAddresses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Addresses); i++ {

		if m.Addresses[i] != nil {

			if swag.IsZero(m.Addresses[i]) { // not required
				return nil
			}

			if err := m.Addresses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("addresses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("addresses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PrivateKeyResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrivateKeyResult) UnmarshalBinary(b []byte) error {
	var res PrivateKeyResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
