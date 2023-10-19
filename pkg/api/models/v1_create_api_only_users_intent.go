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

// V1CreateAPIOnlyUsersIntent v1 create Api only users intent
//
// swagger:model v1CreateApiOnlyUsersIntent
type V1CreateAPIOnlyUsersIntent struct {

	// A list of API-only Users to create.
	// Required: true
	APIOnlyUsers []*V1APIOnlyUserParams `json:"apiOnlyUsers"`
}

// Validate validates this v1 create Api only users intent
func (m *V1CreateAPIOnlyUsersIntent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIOnlyUsers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CreateAPIOnlyUsersIntent) validateAPIOnlyUsers(formats strfmt.Registry) error {

	if err := validate.Required("apiOnlyUsers", "body", m.APIOnlyUsers); err != nil {
		return err
	}

	for i := 0; i < len(m.APIOnlyUsers); i++ {
		if swag.IsZero(m.APIOnlyUsers[i]) { // not required
			continue
		}

		if m.APIOnlyUsers[i] != nil {
			if err := m.APIOnlyUsers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("apiOnlyUsers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("apiOnlyUsers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 create Api only users intent based on the context it is used
func (m *V1CreateAPIOnlyUsersIntent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAPIOnlyUsers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CreateAPIOnlyUsersIntent) contextValidateAPIOnlyUsers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.APIOnlyUsers); i++ {

		if m.APIOnlyUsers[i] != nil {
			if err := m.APIOnlyUsers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("apiOnlyUsers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("apiOnlyUsers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1CreateAPIOnlyUsersIntent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CreateAPIOnlyUsersIntent) UnmarshalBinary(b []byte) error {
	var res V1CreateAPIOnlyUsersIntent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
