// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DecisionsUnauthorized DecisionsUnauthorized DecisionsUnauthorized handles this case with default header values.
//
// The standard error format
//
// swagger:model DecisionsUnauthorized
type DecisionsUnauthorized struct {

	// payload
	Payload *DecisionsUnauthorizedBody `json:"Payload,omitempty"`
}

// Validate validates this decisions unauthorized
func (m *DecisionsUnauthorized) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePayload(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DecisionsUnauthorized) validatePayload(formats strfmt.Registry) error {
	if swag.IsZero(m.Payload) { // not required
		return nil
	}

	if m.Payload != nil {
		if err := m.Payload.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Payload")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this decisions unauthorized based on the context it is used
func (m *DecisionsUnauthorized) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePayload(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DecisionsUnauthorized) contextValidatePayload(ctx context.Context, formats strfmt.Registry) error {

	if m.Payload != nil {
		if err := m.Payload.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Payload")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DecisionsUnauthorized) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DecisionsUnauthorized) UnmarshalBinary(b []byte) error {
	var res DecisionsUnauthorized
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
