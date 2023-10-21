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

// DeleteInvitationIntent delete invitation intent
//
// swagger:model DeleteInvitationIntent
type DeleteInvitationIntent struct {

	// Unique identifier for a given Invitation object.
	// Required: true
	InvitationID *string `json:"invitationId"`
}

// Validate validates this delete invitation intent
func (m *DeleteInvitationIntent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInvitationID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeleteInvitationIntent) validateInvitationID(formats strfmt.Registry) error {

	if err := validate.Required("invitationId", "body", m.InvitationID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this delete invitation intent based on context it is used
func (m *DeleteInvitationIntent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteInvitationIntent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteInvitationIntent) UnmarshalBinary(b []byte) error {
	var res DeleteInvitationIntent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}