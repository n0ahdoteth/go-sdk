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

// ActivityV1Effect activity v1 effect
//
// swagger:model activity.v1.Effect
type ActivityV1Effect string

func NewActivityV1Effect(value ActivityV1Effect) *ActivityV1Effect {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ActivityV1Effect.
func (m ActivityV1Effect) Pointer() *ActivityV1Effect {
	return &m
}

const (

	// ActivityV1EffectEFFECTALLOW captures enum value "EFFECT_ALLOW"
	ActivityV1EffectEFFECTALLOW ActivityV1Effect = "EFFECT_ALLOW"

	// ActivityV1EffectEFFECTDENY captures enum value "EFFECT_DENY"
	ActivityV1EffectEFFECTDENY ActivityV1Effect = "EFFECT_DENY"
)

// for schema
var activityV1EffectEnum []interface{}

func init() {
	var res []ActivityV1Effect
	if err := json.Unmarshal([]byte(`["EFFECT_ALLOW","EFFECT_DENY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		activityV1EffectEnum = append(activityV1EffectEnum, v)
	}
}

func (m ActivityV1Effect) validateActivityV1EffectEnum(path, location string, value ActivityV1Effect) error {
	if err := validate.EnumCase(path, location, value, activityV1EffectEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this activity v1 effect
func (m ActivityV1Effect) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateActivityV1EffectEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this activity v1 effect based on context it is used
func (m ActivityV1Effect) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}