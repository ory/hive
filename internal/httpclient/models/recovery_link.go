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

// RecoveryLink recovery link
//
// swagger:model recoveryLink
type RecoveryLink struct {

	// Recovery Link Expires At
	//
	// The timestamp when the recovery link expires.
	// Format: date-time
	ExpiresAt strfmt.DateTime `json:"expires_at,omitempty"`

	// Recovery Link
	//
	// This link can be used to recover the account.
	// Required: true
	RecoveryLink *string `json:"recovery_link"`
}

// Validate validates this recovery link
func (m *RecoveryLink) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpiresAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecoveryLink(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoveryLink) validateExpiresAt(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpiresAt) { // not required
		return nil
	}

	if err := validate.FormatOf("expires_at", "body", "date-time", m.ExpiresAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecoveryLink) validateRecoveryLink(formats strfmt.Registry) error {

	if err := validate.Required("recovery_link", "body", m.RecoveryLink); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this recovery link based on context it is used
func (m *RecoveryLink) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RecoveryLink) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoveryLink) UnmarshalBinary(b []byte) error {
	var res RecoveryLink
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
