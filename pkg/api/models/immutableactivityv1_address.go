// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Immutableactivityv1Address immutableactivityv1 address
//
// swagger:model immutableactivityv1Address
type Immutableactivityv1Address struct {

	// address
	Address string `json:"address,omitempty"`

	// format
	Format Immutablecommonv1AddressFormat `json:"format,omitempty"`
}

// Validate validates this immutableactivityv1 address
func (m *Immutableactivityv1Address) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFormat(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Immutableactivityv1Address) validateFormat(formats strfmt.Registry) error {
	if swag.IsZero(m.Format) { // not required
		return nil
	}

	if err := m.Format.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("format")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("format")
		}
		return err
	}

	return nil
}

// ContextValidate validate this immutableactivityv1 address based on the context it is used
func (m *Immutableactivityv1Address) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFormat(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Immutableactivityv1Address) contextValidateFormat(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Format.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("format")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("format")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Immutableactivityv1Address) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Immutableactivityv1Address) UnmarshalBinary(b []byte) error {
	var res Immutableactivityv1Address
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
