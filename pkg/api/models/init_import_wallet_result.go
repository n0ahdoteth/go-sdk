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

// InitImportWalletResult init import wallet result
//
// swagger:model InitImportWalletResult
type InitImportWalletResult struct {

	// Import bundle containing a public key and signature to use for importing client data.
	// Required: true
	ImportBundle *string `json:"importBundle"`
}

// Validate validates this init import wallet result
func (m *InitImportWalletResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImportBundle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InitImportWalletResult) validateImportBundle(formats strfmt.Registry) error {

	if err := validate.Required("importBundle", "body", m.ImportBundle); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this init import wallet result based on context it is used
func (m *InitImportWalletResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InitImportWalletResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InitImportWalletResult) UnmarshalBinary(b []byte) error {
	var res InitImportWalletResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
