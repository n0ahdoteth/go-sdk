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

// Effect effect
//
// swagger:model Effect
type Effect string

func NewEffect(value Effect) *Effect {
	return &value
}

// Pointer returns a pointer to a freshly-allocated Effect.
func (m Effect) Pointer() *Effect {
	return &m
}

const (

	// EffectAllow captures enum value "EFFECT_ALLOW"
	EffectAllow Effect = "EFFECT_ALLOW"

	// EffectDeny captures enum value "EFFECT_DENY"
	EffectDeny Effect = "EFFECT_DENY"
)

// for schema
var EffectEnum []Effect

func init() {
	var res []Effect
	if err := json.Unmarshal([]byte(`["EFFECT_ALLOW","EFFECT_DENY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		EffectEnum = append(EffectEnum, v)
	}
}

func (m Effect) validateEffectEnum(path, location string, value Effect) error {
	if err := validate.EnumCase(path, location, value, EffectEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this effect
func (m Effect) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateEffectEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this effect based on context it is used
func (m Effect) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}