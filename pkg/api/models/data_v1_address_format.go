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

// DataV1AddressFormat data v1 address format
//
// swagger:model data.v1.AddressFormat
type DataV1AddressFormat string

func NewDataV1AddressFormat(value DataV1AddressFormat) *DataV1AddressFormat {
	return &value
}

// Pointer returns a pointer to a freshly-allocated DataV1AddressFormat.
func (m DataV1AddressFormat) Pointer() *DataV1AddressFormat {
	return &m
}

const (

	// AddressFormatUncompressed captures enum value "ADDRESS_FORMAT_UNCOMPRESSED"
	AddressFormatUncompressed DataV1AddressFormat = "ADDRESS_FORMAT_UNCOMPRESSED"

	// AddressFormatCompressed captures enum value "ADDRESS_FORMAT_COMPRESSED"
	AddressFormatCompressed DataV1AddressFormat = "ADDRESS_FORMAT_COMPRESSED"

	// AddressFormatEthereum captures enum value "ADDRESS_FORMAT_ETHEREUM"
	AddressFormatEthereum DataV1AddressFormat = "ADDRESS_FORMAT_ETHEREUM"
)

// for schema
var DataV1AddressFormatEnum []DataV1AddressFormat

func init() {
	var res []DataV1AddressFormat
	if err := json.Unmarshal([]byte(`["ADDRESS_FORMAT_UNCOMPRESSED","ADDRESS_FORMAT_COMPRESSED","ADDRESS_FORMAT_ETHEREUM"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		DataV1AddressFormatEnum = append(DataV1AddressFormatEnum, v)
	}
}

func (m DataV1AddressFormat) validateDataV1AddressFormatEnum(path, location string, value DataV1AddressFormat) error {
	if err := validate.EnumCase(path, location, value, DataV1AddressFormatEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this data v1 address format
func (m DataV1AddressFormat) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDataV1AddressFormatEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this data v1 address format based on context it is used
func (m DataV1AddressFormat) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
