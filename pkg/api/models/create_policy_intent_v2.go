// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreatePolicyIntentV2 create policy intent v2
//
// swagger:model CreatePolicyIntentV2
type CreatePolicyIntentV2 struct {

	// Whether to ALLOW or DENY requests that match the condition and consensus requirements.
	// Required: true
	Effect *ActivityV1Effect `json:"effect"`

	// notes
	Notes string `json:"notes,omitempty"`

	// Human-readable name for a Policy.
	// Required: true
	PolicyName *string `json:"policyName"`

	// A list of simple functions each including a subject, target and boolean. See Policy Engine Language section for additional details.
	// Required: true
	Selectors []*SelectorV2 `json:"selectors"`
}

// Validate validates this create policy intent v2
func (m *CreatePolicyIntentV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEffect(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicyName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelectors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreatePolicyIntentV2) validateEffect(formats strfmt.Registry) error {

	if err := validate.Required("effect", "body", m.Effect); err != nil {
		return err
	}

	if err := validate.Required("effect", "body", m.Effect); err != nil {
		return err
	}

	if m.Effect != nil {
		if err := m.Effect.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("effect")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("effect")
			}
			return err
		}
	}

	return nil
}

func (m *CreatePolicyIntentV2) validatePolicyName(formats strfmt.Registry) error {

	if err := validate.Required("policyName", "body", m.PolicyName); err != nil {
		return err
	}

	return nil
}

func (m *CreatePolicyIntentV2) validateSelectors(formats strfmt.Registry) error {

	if err := validate.Required("selectors", "body", m.Selectors); err != nil {
		return err
	}

	for i := 0; i < len(m.Selectors); i++ {
		if swag.IsZero(m.Selectors[i]) { // not required
			continue
		}

		if m.Selectors[i] != nil {
			if err := m.Selectors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("selectors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("selectors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this create policy intent v2 based on the context it is used
func (m *CreatePolicyIntentV2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEffect(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelectors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreatePolicyIntentV2) contextValidateEffect(ctx context.Context, formats strfmt.Registry) error {

	if m.Effect != nil {
		if err := m.Effect.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("effect")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("effect")
			}
			return err
		}
	}

	return nil
}

func (m *CreatePolicyIntentV2) contextValidateSelectors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Selectors); i++ {

		if m.Selectors[i] != nil {
			if err := m.Selectors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("selectors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("selectors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreatePolicyIntentV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreatePolicyIntentV2) UnmarshalBinary(b []byte) error {
	var res CreatePolicyIntentV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
