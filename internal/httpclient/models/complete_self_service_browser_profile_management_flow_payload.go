// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CompleteSelfServiceBrowserProfileManagementFlowPayload complete self service browser profile management flow payload
// swagger:model completeSelfServiceBrowserProfileManagementFlowPayload
type CompleteSelfServiceBrowserProfileManagementFlowPayload struct {

	// request
	// Required: true
	// Format: uuid4
	Request UUID `json:"request"`

	// Traits contains all of the identity's traits.
	//
	// type: string
	// format: binary
	// Required: true
	Traits interface{} `json:"traits"`
}

// Validate validates this complete self service browser profile management flow payload
func (m *CompleteSelfServiceBrowserProfileManagementFlowPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRequest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTraits(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CompleteSelfServiceBrowserProfileManagementFlowPayload) validateRequest(formats strfmt.Registry) error {

	if err := m.Request.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("request")
		}
		return err
	}

	return nil
}

func (m *CompleteSelfServiceBrowserProfileManagementFlowPayload) validateTraits(formats strfmt.Registry) error {

	if err := validate.Required("traits", "body", m.Traits); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CompleteSelfServiceBrowserProfileManagementFlowPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CompleteSelfServiceBrowserProfileManagementFlowPayload) UnmarshalBinary(b []byte) error {
	var res CompleteSelfServiceBrowserProfileManagementFlowPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
