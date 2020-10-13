// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CompleteSelfServiceSettingsFlowWithPasswordMethod complete self service settings flow with password method
//
// swagger:model CompleteSelfServiceSettingsFlowWithPasswordMethod
type CompleteSelfServiceSettingsFlowWithPasswordMethod struct {

	// CSRFToken is the anti-CSRF token
	//
	// type: string
	CsrfToken string `json:"csrf_token,omitempty"`

	// Password is the updated password
	//
	// type: string
	// Required: true
	Password *string `json:"password"`
}

// Validate validates this complete self service settings flow with password method
func (m *CompleteSelfServiceSettingsFlowWithPasswordMethod) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CompleteSelfServiceSettingsFlowWithPasswordMethod) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CompleteSelfServiceSettingsFlowWithPasswordMethod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CompleteSelfServiceSettingsFlowWithPasswordMethod) UnmarshalBinary(b []byte) error {
	var res CompleteSelfServiceSettingsFlowWithPasswordMethod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
