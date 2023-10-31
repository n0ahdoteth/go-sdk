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

// ActivityV1Operator activity v1 operator
//
// swagger:model activity.v1.Operator
type ActivityV1Operator string

func NewActivityV1Operator(value ActivityV1Operator) *ActivityV1Operator {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ActivityV1Operator.
func (m ActivityV1Operator) Pointer() *ActivityV1Operator {
	return &m
}

const (

	// OperatorEqual captures enum value "OPERATOR_EQUAL"
	OperatorEqual ActivityV1Operator = "OPERATOR_EQUAL"

	// OperatorMoreThan captures enum value "OPERATOR_MORE_THAN"
	OperatorMoreThan ActivityV1Operator = "OPERATOR_MORE_THAN"

	// OperatorMoreThanOrEqual captures enum value "OPERATOR_MORE_THAN_OR_EQUAL"
	OperatorMoreThanOrEqual ActivityV1Operator = "OPERATOR_MORE_THAN_OR_EQUAL"

	// OperatorLessThan captures enum value "OPERATOR_LESS_THAN"
	OperatorLessThan ActivityV1Operator = "OPERATOR_LESS_THAN"

	// OperatorLessThanOrEqual captures enum value "OPERATOR_LESS_THAN_OR_EQUAL"
	OperatorLessThanOrEqual ActivityV1Operator = "OPERATOR_LESS_THAN_OR_EQUAL"

	// OperatorContains captures enum value "OPERATOR_CONTAINS"
	OperatorContains ActivityV1Operator = "OPERATOR_CONTAINS"

	// OperatorNotEqual captures enum value "OPERATOR_NOT_EQUAL"
	OperatorNotEqual ActivityV1Operator = "OPERATOR_NOT_EQUAL"

	// OperatorIn captures enum value "OPERATOR_IN"
	OperatorIn ActivityV1Operator = "OPERATOR_IN"

	// OperatorNotIn captures enum value "OPERATOR_NOT_IN"
	OperatorNotIn ActivityV1Operator = "OPERATOR_NOT_IN"

	// OperatorContainsOne captures enum value "OPERATOR_CONTAINS_ONE"
	OperatorContainsOne ActivityV1Operator = "OPERATOR_CONTAINS_ONE"

	// OperatorContainsAll captures enum value "OPERATOR_CONTAINS_ALL"
	OperatorContainsAll ActivityV1Operator = "OPERATOR_CONTAINS_ALL"
)

// for schema
var ActivityV1OperatorEnum []ActivityV1Operator

func init() {
	var res []ActivityV1Operator
	if err := json.Unmarshal([]byte(`["OPERATOR_EQUAL","OPERATOR_MORE_THAN","OPERATOR_MORE_THAN_OR_EQUAL","OPERATOR_LESS_THAN","OPERATOR_LESS_THAN_OR_EQUAL","OPERATOR_CONTAINS","OPERATOR_NOT_EQUAL","OPERATOR_IN","OPERATOR_NOT_IN","OPERATOR_CONTAINS_ONE","OPERATOR_CONTAINS_ALL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ActivityV1OperatorEnum = append(ActivityV1OperatorEnum, v)
	}
}

func (m ActivityV1Operator) validateActivityV1OperatorEnum(path, location string, value ActivityV1Operator) error {
	if err := validate.EnumCase(path, location, value, ActivityV1OperatorEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this activity v1 operator
func (m ActivityV1Operator) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateActivityV1OperatorEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this activity v1 operator based on context it is used
func (m ActivityV1Operator) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
