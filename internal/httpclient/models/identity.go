// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Identity identity
//
// swagger:model Identity
type Identity struct {

	// addresses
	Addresses []*VerifiableAddress `json:"addresses"`

	// id
	// Required: true
	// Format: uuid4
	ID UUID `json:"id"`

	// traits
	// Required: true
	Traits Traits `json:"traits"`

	// TraitsSchemaID is the ID of the JSON Schema to be used for validating the identity's traits.
	// Required: true
	TraitsSchemaID *string `json:"traits_schema_id"`

	// TraitsSchemaURL is the URL of the endpoint where the identity's traits schema can be fetched from.
	//
	// format: url
	TraitsSchemaURL string `json:"traits_schema_url,omitempty"`
}

// Validate validates this identity
func (m *Identity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddresses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTraits(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTraitsSchemaID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Identity) validateAddresses(formats strfmt.Registry) error {

	if swag.IsZero(m.Addresses) { // not required
		return nil
	}

	for i := 0; i < len(m.Addresses); i++ {
		if swag.IsZero(m.Addresses[i]) { // not required
			continue
		}

		if m.Addresses[i] != nil {
			if err := m.Addresses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("addresses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Identity) validateID(formats strfmt.Registry) error {

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *Identity) validateTraits(formats strfmt.Registry) error {

	if err := m.Traits.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("traits")
		}
		return err
	}

	return nil
}

func (m *Identity) validateTraitsSchemaID(formats strfmt.Registry) error {

	if err := validate.Required("traits_schema_id", "body", m.TraitsSchemaID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Identity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Identity) UnmarshalBinary(b []byte) error {
	var res Identity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
