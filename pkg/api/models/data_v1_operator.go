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

// DataV1Operator data v1 operator
//
// swagger:model data.v1.Operator
type DataV1Operator string

func NewDataV1Operator(value DataV1Operator) *DataV1Operator {
	return &value
}

// Pointer returns a pointer to a freshly-allocated DataV1Operator.
func (m DataV1Operator) Pointer() *DataV1Operator {
	return &m
}

const (

	// DataV1OperatorOPERATOREQUAL captures enum value "OPERATOR_EQUAL"
	DataV1OperatorOPERATOREQUAL DataV1Operator = "OPERATOR_EQUAL"

	// DataV1OperatorOPERATORMORETHAN captures enum value "OPERATOR_MORE_THAN"
	DataV1OperatorOPERATORMORETHAN DataV1Operator = "OPERATOR_MORE_THAN"

	// DataV1OperatorOPERATORMORETHANOREQUAL captures enum value "OPERATOR_MORE_THAN_OR_EQUAL"
	DataV1OperatorOPERATORMORETHANOREQUAL DataV1Operator = "OPERATOR_MORE_THAN_OR_EQUAL"

	// DataV1OperatorOPERATORLESSTHAN captures enum value "OPERATOR_LESS_THAN"
	DataV1OperatorOPERATORLESSTHAN DataV1Operator = "OPERATOR_LESS_THAN"

	// DataV1OperatorOPERATORLESSTHANOREQUAL captures enum value "OPERATOR_LESS_THAN_OR_EQUAL"
	DataV1OperatorOPERATORLESSTHANOREQUAL DataV1Operator = "OPERATOR_LESS_THAN_OR_EQUAL"

	// DataV1OperatorOPERATORCONTAINS captures enum value "OPERATOR_CONTAINS"
	DataV1OperatorOPERATORCONTAINS DataV1Operator = "OPERATOR_CONTAINS"

	// DataV1OperatorOPERATORNOTEQUAL captures enum value "OPERATOR_NOT_EQUAL"
	DataV1OperatorOPERATORNOTEQUAL DataV1Operator = "OPERATOR_NOT_EQUAL"

	// DataV1OperatorOPERATORIN captures enum value "OPERATOR_IN"
	DataV1OperatorOPERATORIN DataV1Operator = "OPERATOR_IN"

	// DataV1OperatorOPERATORNOTIN captures enum value "OPERATOR_NOT_IN"
	DataV1OperatorOPERATORNOTIN DataV1Operator = "OPERATOR_NOT_IN"

	// DataV1OperatorOPERATORCONTAINSONE captures enum value "OPERATOR_CONTAINS_ONE"
	DataV1OperatorOPERATORCONTAINSONE DataV1Operator = "OPERATOR_CONTAINS_ONE"

	// DataV1OperatorOPERATORCONTAINSALL captures enum value "OPERATOR_CONTAINS_ALL"
	DataV1OperatorOPERATORCONTAINSALL DataV1Operator = "OPERATOR_CONTAINS_ALL"
)

// for schema
var dataV1OperatorEnum []interface{}

func init() {
	var res []DataV1Operator
	if err := json.Unmarshal([]byte(`["OPERATOR_EQUAL","OPERATOR_MORE_THAN","OPERATOR_MORE_THAN_OR_EQUAL","OPERATOR_LESS_THAN","OPERATOR_LESS_THAN_OR_EQUAL","OPERATOR_CONTAINS","OPERATOR_NOT_EQUAL","OPERATOR_IN","OPERATOR_NOT_IN","OPERATOR_CONTAINS_ONE","OPERATOR_CONTAINS_ALL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dataV1OperatorEnum = append(dataV1OperatorEnum, v)
	}
}

func (m DataV1Operator) validateDataV1OperatorEnum(path, location string, value DataV1Operator) error {
	if err := validate.EnumCase(path, location, value, dataV1OperatorEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this data v1 operator
func (m DataV1Operator) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDataV1OperatorEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this data v1 operator based on context it is used
func (m DataV1Operator) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}