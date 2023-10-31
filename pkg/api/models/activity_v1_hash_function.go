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

// ActivityV1HashFunction activity v1 hash function
//
// swagger:model activity.v1.HashFunction
type ActivityV1HashFunction string

func NewActivityV1HashFunction(value ActivityV1HashFunction) *ActivityV1HashFunction {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ActivityV1HashFunction.
func (m ActivityV1HashFunction) Pointer() *ActivityV1HashFunction {
	return &m
}

const (

	// HashFunctionNoOp captures enum value "HASH_FUNCTION_NO_OP"
	HashFunctionNoOp ActivityV1HashFunction = "HASH_FUNCTION_NO_OP"

	// HashFunctionSha256 captures enum value "HASH_FUNCTION_SHA256"
	HashFunctionSha256 ActivityV1HashFunction = "HASH_FUNCTION_SHA256"

	// HashFunctionKeccak256 captures enum value "HASH_FUNCTION_KECCAK256"
	HashFunctionKeccak256 ActivityV1HashFunction = "HASH_FUNCTION_KECCAK256"

	// HashFunctionNotApplicable captures enum value "HASH_FUNCTION_NOT_APPLICABLE"
	HashFunctionNotApplicable ActivityV1HashFunction = "HASH_FUNCTION_NOT_APPLICABLE"
)

// for schema
var ActivityV1HashFunctionEnum []ActivityV1HashFunction

func init() {
	var res []ActivityV1HashFunction
	if err := json.Unmarshal([]byte(`["HASH_FUNCTION_NO_OP","HASH_FUNCTION_SHA256","HASH_FUNCTION_KECCAK256","HASH_FUNCTION_NOT_APPLICABLE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ActivityV1HashFunctionEnum = append(ActivityV1HashFunctionEnum, v)
	}
}

func (m ActivityV1HashFunction) validateActivityV1HashFunctionEnum(path, location string, value ActivityV1HashFunction) error {
	if err := validate.EnumCase(path, location, value, ActivityV1HashFunctionEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this activity v1 hash function
func (m ActivityV1HashFunction) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateActivityV1HashFunctionEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this activity v1 hash function based on context it is used
func (m ActivityV1HashFunction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
