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

// UpdateUserTagIntent update user tag intent
//
// swagger:model UpdateUserTagIntent
type UpdateUserTagIntent struct {

	// A list of User IDs to add this tag to.
	// Required: true
	AddUserIds []string `json:"addUserIds"`

	// The new, human-readable name for the tag with the given ID.
	NewUserTagName string `json:"newUserTagName,omitempty"`

	// A list of User IDs to remove this tag from.
	// Required: true
	RemoveUserIds []string `json:"removeUserIds"`

	// Unique identifier for a given User Tag.
	// Required: true
	UserTagID *string `json:"userTagId"`
}

// Validate validates this update user tag intent
func (m *UpdateUserTagIntent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddUserIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemoveUserIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserTagID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateUserTagIntent) validateAddUserIds(formats strfmt.Registry) error {

	if err := validate.Required("addUserIds", "body", m.AddUserIds); err != nil {
		return err
	}

	return nil
}

func (m *UpdateUserTagIntent) validateRemoveUserIds(formats strfmt.Registry) error {

	if err := validate.Required("removeUserIds", "body", m.RemoveUserIds); err != nil {
		return err
	}

	return nil
}

func (m *UpdateUserTagIntent) validateUserTagID(formats strfmt.Registry) error {

	if err := validate.Required("userTagId", "body", m.UserTagID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update user tag intent based on context it is used
func (m *UpdateUserTagIntent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateUserTagIntent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateUserTagIntent) UnmarshalBinary(b []byte) error {
	var res UpdateUserTagIntent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}