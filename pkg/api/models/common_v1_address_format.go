// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// CommonV1AddressFormat common v1 address format
//
// swagger:model common.v1.AddressFormat
type CommonV1AddressFormat string

func NewCommonV1AddressFormat(value CommonV1AddressFormat) *CommonV1AddressFormat {
	return &value
}

// Pointer returns a pointer to a freshly-allocated CommonV1AddressFormat.
func (m CommonV1AddressFormat) Pointer() *CommonV1AddressFormat {
	return &m
}

const (

	// AddressFormatUncompressed captures enum value "ADDRESS_FORMAT_UNCOMPRESSED"
	AddressFormatUncompressed CommonV1AddressFormat = "ADDRESS_FORMAT_UNCOMPRESSED"

	// AddressFormatCompressed captures enum value "ADDRESS_FORMAT_COMPRESSED"
	AddressFormatCompressed CommonV1AddressFormat = "ADDRESS_FORMAT_COMPRESSED"

	// AddressFormatEthereum captures enum value "ADDRESS_FORMAT_ETHEREUM"
	AddressFormatEthereum CommonV1AddressFormat = "ADDRESS_FORMAT_ETHEREUM"

	// AddressFormatSolana captures enum value "ADDRESS_FORMAT_SOLANA"
	AddressFormatSolana CommonV1AddressFormat = "ADDRESS_FORMAT_SOLANA"

	// AddressFormatCosmos captures enum value "ADDRESS_FORMAT_COSMOS"
	AddressFormatCosmos CommonV1AddressFormat = "ADDRESS_FORMAT_COSMOS"
)

// for schema
var CommonV1AddressFormatEnum []CommonV1AddressFormat

func init() {
	var res []CommonV1AddressFormat
	if err := json.Unmarshal([]byte(`["ADDRESS_FORMAT_UNCOMPRESSED","ADDRESS_FORMAT_COMPRESSED","ADDRESS_FORMAT_ETHEREUM","ADDRESS_FORMAT_SOLANA","ADDRESS_FORMAT_COSMOS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		CommonV1AddressFormatEnum = append(CommonV1AddressFormatEnum, v)
	}
}

func (m CommonV1AddressFormat) validateCommonV1AddressFormatEnum(path, location string, value CommonV1AddressFormat) error {
	if err := validate.EnumCase(path, location, value, CommonV1AddressFormatEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this common v1 address format
func (m CommonV1AddressFormat) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCommonV1AddressFormatEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this common v1 address format based on context it is used
func (m CommonV1AddressFormat) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
