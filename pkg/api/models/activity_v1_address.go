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

// ActivityV1Address activity v1 address
//
// swagger:model activity.v1.Address
type ActivityV1Address struct {

	// address
	Address string `json:"address,omitempty"`

	// format
	Format CommonV1AddressFormat `json:"format,omitempty"`
}

// Validate validates this activity v1 address
func (m *ActivityV1Address) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFormat(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActivityV1Address) validateFormat(formats strfmt.Registry) error {
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

// ContextValidate validate this activity v1 address based on the context it is used
func (m *ActivityV1Address) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFormat(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActivityV1Address) contextValidateFormat(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Format) { // not required
		return nil
	}

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
func (m *ActivityV1Address) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActivityV1Address) UnmarshalBinary(b []byte) error {
	var res ActivityV1Address
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
