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

// ExportWalletAccountResult export wallet account result
//
// swagger:model ExportWalletAccountResult
type ExportWalletAccountResult struct {

	// Address to identify Wallet Account.
	// Required: true
	Address *string `json:"address"`

	// Export bundle containing a private key encrypted by the client's target public key.
	// Required: true
	ExportBundle *string `json:"exportBundle"`
}

// Validate validates this export wallet account result
func (m *ExportWalletAccountResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExportBundle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExportWalletAccountResult) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("address", "body", m.Address); err != nil {
		return err
	}

	return nil
}

func (m *ExportWalletAccountResult) validateExportBundle(formats strfmt.Registry) error {

	if err := validate.Required("exportBundle", "body", m.ExportBundle); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this export wallet account result based on context it is used
func (m *ExportWalletAccountResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExportWalletAccountResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExportWalletAccountResult) UnmarshalBinary(b []byte) error {
	var res ExportWalletAccountResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}