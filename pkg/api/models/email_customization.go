// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EmailCustomization email customization
//
// swagger:model EmailCustomization
type EmailCustomization struct {

	// body
	Body string `json:"body,omitempty"`

	// styling
	Styling string `json:"styling,omitempty"`

	// subject
	Subject string `json:"subject,omitempty"`

	// url prefix
	URLPrefix string `json:"urlPrefix,omitempty"`
}

// Validate validates this email customization
func (m *EmailCustomization) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this email customization based on context it is used
func (m *EmailCustomization) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EmailCustomization) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmailCustomization) UnmarshalBinary(b []byte) error {
	var res EmailCustomization
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
