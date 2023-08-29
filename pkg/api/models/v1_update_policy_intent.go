// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1UpdatePolicyIntent v1 update policy intent
//
// swagger:model v1UpdatePolicyIntent
type V1UpdatePolicyIntent struct {

	// The condition expression that triggers the Effect (optional).
	PolicyCondition string `json:"policyCondition,omitempty"`

	// The consensus expression that triggers the Effect (optional).
	PolicyConsensus string `json:"policyConsensus,omitempty"`

	// The instruction to DENY or ALLOW an activity (optional).
	PolicyEffect Immutableactivityv1Effect `json:"policyEffect,omitempty"`

	// Unique identifier for a given Policy.
	// Required: true
	PolicyID *string `json:"policyId"`

	// Human-readable name for a Policy.
	PolicyName string `json:"policyName,omitempty"`

	// Accompanying notes for a Policy (optional).
	PolicyNotes string `json:"policyNotes,omitempty"`
}

// Validate validates this v1 update policy intent
func (m *V1UpdatePolicyIntent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePolicyEffect(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicyID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1UpdatePolicyIntent) validatePolicyEffect(formats strfmt.Registry) error {
	if swag.IsZero(m.PolicyEffect) { // not required
		return nil
	}

	if err := m.PolicyEffect.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("policyEffect")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("policyEffect")
		}
		return err
	}

	return nil
}

func (m *V1UpdatePolicyIntent) validatePolicyID(formats strfmt.Registry) error {

	if err := validate.Required("policyId", "body", m.PolicyID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 update policy intent based on the context it is used
func (m *V1UpdatePolicyIntent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePolicyEffect(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1UpdatePolicyIntent) contextValidatePolicyEffect(ctx context.Context, formats strfmt.Registry) error {

	if err := m.PolicyEffect.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("policyEffect")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("policyEffect")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1UpdatePolicyIntent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1UpdatePolicyIntent) UnmarshalBinary(b []byte) error {
	var res V1UpdatePolicyIntent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
