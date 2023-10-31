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

// ActivityV1PayloadEncoding activity v1 payload encoding
//
// swagger:model activity.v1.PayloadEncoding
type ActivityV1PayloadEncoding string

func NewActivityV1PayloadEncoding(value ActivityV1PayloadEncoding) *ActivityV1PayloadEncoding {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ActivityV1PayloadEncoding.
func (m ActivityV1PayloadEncoding) Pointer() *ActivityV1PayloadEncoding {
	return &m
}

const (

	// PayloadEncodingHexadecimal captures enum value "PAYLOAD_ENCODING_HEXADECIMAL"
	PayloadEncodingHexadecimal ActivityV1PayloadEncoding = "PAYLOAD_ENCODING_HEXADECIMAL"

	// PayloadEncodingTextUTF8 captures enum value "PAYLOAD_ENCODING_TEXT_UTF8"
	PayloadEncodingTextUTF8 ActivityV1PayloadEncoding = "PAYLOAD_ENCODING_TEXT_UTF8"
)

// for schema
var ActivityV1PayloadEncodingEnum []ActivityV1PayloadEncoding

func init() {
	var res []ActivityV1PayloadEncoding
	if err := json.Unmarshal([]byte(`["PAYLOAD_ENCODING_HEXADECIMAL","PAYLOAD_ENCODING_TEXT_UTF8"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ActivityV1PayloadEncodingEnum = append(ActivityV1PayloadEncodingEnum, v)
	}
}

func (m ActivityV1PayloadEncoding) validateActivityV1PayloadEncodingEnum(path, location string, value ActivityV1PayloadEncoding) error {
	if err := validate.EnumCase(path, location, value, ActivityV1PayloadEncodingEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this activity v1 payload encoding
func (m ActivityV1PayloadEncoding) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateActivityV1PayloadEncodingEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this activity v1 payload encoding based on context it is used
func (m ActivityV1PayloadEncoding) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
