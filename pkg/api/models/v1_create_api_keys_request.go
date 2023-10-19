// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1CreateAPIKeysRequest v1 create Api keys request
//
// swagger:model v1CreateApiKeysRequest
type V1CreateAPIKeysRequest struct {

	// Unique identifier for a given Organization.
	// Required: true
	OrganizationID *string `json:"organizationId"`

	// parameters
	// Required: true
	Parameters *V1CreateAPIKeysIntent `json:"parameters"`

	// Timestamp (in milliseconds) of the request, used to verify liveness of user requests.
	// Required: true
	TimestampMs *string `json:"timestampMs"`

	// type
	// Required: true
	// Enum: [ACTIVITY_TYPE_CREATE_API_KEYS]
	Type *string `json:"type"`
}

// Validate validates this v1 create Api keys request
func (m *V1CreateAPIKeysRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrganizationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParameters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestampMs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CreateAPIKeysRequest) validateOrganizationID(formats strfmt.Registry) error {

	if err := validate.Required("organizationId", "body", m.OrganizationID); err != nil {
		return err
	}

	return nil
}

func (m *V1CreateAPIKeysRequest) validateParameters(formats strfmt.Registry) error {

	if err := validate.Required("parameters", "body", m.Parameters); err != nil {
		return err
	}

	if m.Parameters != nil {
		if err := m.Parameters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parameters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parameters")
			}
			return err
		}
	}

	return nil
}

func (m *V1CreateAPIKeysRequest) validateTimestampMs(formats strfmt.Registry) error {

	if err := validate.Required("timestampMs", "body", m.TimestampMs); err != nil {
		return err
	}

	return nil
}

var v1CreateApiKeysRequestTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACTIVITY_TYPE_CREATE_API_KEYS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1CreateApiKeysRequestTypeTypePropEnum = append(v1CreateApiKeysRequestTypeTypePropEnum, v)
	}
}

const (

	// V1CreateAPIKeysRequestTypeACTIVITYTYPECREATEAPIKEYS captures enum value "ACTIVITY_TYPE_CREATE_API_KEYS"
	V1CreateAPIKeysRequestTypeACTIVITYTYPECREATEAPIKEYS string = "ACTIVITY_TYPE_CREATE_API_KEYS"
)

// prop value enum
func (m *V1CreateAPIKeysRequest) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1CreateApiKeysRequestTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1CreateAPIKeysRequest) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 create Api keys request based on the context it is used
func (m *V1CreateAPIKeysRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateParameters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CreateAPIKeysRequest) contextValidateParameters(ctx context.Context, formats strfmt.Registry) error {

	if m.Parameters != nil {
		if err := m.Parameters.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parameters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parameters")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1CreateAPIKeysRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CreateAPIKeysRequest) UnmarshalBinary(b []byte) error {
	var res V1CreateAPIKeysRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
