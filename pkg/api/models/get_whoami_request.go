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

// GetWhoamiRequest get whoami request
//
// swagger:model GetWhoamiRequest
type GetWhoamiRequest struct {

	// Unique identifier for a given Organization. If the request is being made by a WebAuthN user and their Sub-Organization ID is unknown, this can be the Parent Organization ID; using the Sub-Organization ID when possible is preferred due to performance reasons.
	// Required: true
	OrganizationID *string `json:"organizationId"`
}

// Validate validates this get whoami request
func (m *GetWhoamiRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrganizationID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetWhoamiRequest) validateOrganizationID(formats strfmt.Registry) error {

	if err := validate.Required("organizationId", "body", m.OrganizationID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get whoami request based on context it is used
func (m *GetWhoamiRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetWhoamiRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetWhoamiRequest) UnmarshalBinary(b []byte) error {
	var res GetWhoamiRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}