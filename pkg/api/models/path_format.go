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

// PathFormat path format
//
// swagger:model PathFormat
type PathFormat string

func NewPathFormat(value PathFormat) *PathFormat {
	return &value
}

// Pointer returns a pointer to a freshly-allocated PathFormat.
func (m PathFormat) Pointer() *PathFormat {
	return &m
}

const (

	// PathFormatBip32 captures enum value "PATH_FORMAT_BIP32"
	PathFormatBip32 PathFormat = "PATH_FORMAT_BIP32"
)

// for schema
var PathFormatEnum []PathFormat

func init() {
	var res []PathFormat
	if err := json.Unmarshal([]byte(`["PATH_FORMAT_BIP32"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		PathFormatEnum = append(PathFormatEnum, v)
	}
}

func (m PathFormat) validatePathFormatEnum(path, location string, value PathFormat) error {
	if err := validate.EnumCase(path, location, value, PathFormatEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this path format
func (m PathFormat) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePathFormatEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this path format based on context it is used
func (m PathFormat) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
