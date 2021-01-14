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

// LoginViaAPIResponse LoginViaAPIResponse The Response for Login Flows via API
//
// swagger:model loginViaApiResponse
type LoginViaAPIResponse struct {

	// session
	// Required: true
	Session *Session `json:"session"`

	// The Session Token
	//
	// A session token is equivalent to a session cookie, but it can be sent in the HTTP Authorization
	// Header:
	//
	// Authorization: bearer ${session-token}
	//
	// The session token is only issued for API flows, not for Browser flows!
	// Required: true
	SessionToken *string `json:"session_token"`
}

// Validate validates this login via Api response
func (m *LoginViaAPIResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSession(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSessionToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoginViaAPIResponse) validateSession(formats strfmt.Registry) error {

	if err := validate.Required("session", "body", m.Session); err != nil {
		return err
	}

	if m.Session != nil {
		if err := m.Session.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("session")
			}
			return err
		}
	}

	return nil
}

func (m *LoginViaAPIResponse) validateSessionToken(formats strfmt.Registry) error {

	if err := validate.Required("session_token", "body", m.SessionToken); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LoginViaAPIResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoginViaAPIResponse) UnmarshalBinary(b []byte) error {
	var res LoginViaAPIResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
