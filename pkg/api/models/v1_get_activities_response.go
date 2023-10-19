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

// V1GetActivitiesResponse v1 get activities response
//
// swagger:model v1GetActivitiesResponse
type V1GetActivitiesResponse struct {

	// A list of Activities.
	// Required: true
	Activities []*V1Activity `json:"activities"`
}

// Validate validates this v1 get activities response
func (m *V1GetActivitiesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActivities(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1GetActivitiesResponse) validateActivities(formats strfmt.Registry) error {

	if err := validate.Required("activities", "body", m.Activities); err != nil {
		return err
	}

	for i := 0; i < len(m.Activities); i++ {
		if swag.IsZero(m.Activities[i]) { // not required
			continue
		}

		if m.Activities[i] != nil {
			if err := m.Activities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("activities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("activities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 get activities response based on the context it is used
func (m *V1GetActivitiesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActivities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1GetActivitiesResponse) contextValidateActivities(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Activities); i++ {

		if m.Activities[i] != nil {
			if err := m.Activities[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("activities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("activities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1GetActivitiesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1GetActivitiesResponse) UnmarshalBinary(b []byte) error {
	var res V1GetActivitiesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
