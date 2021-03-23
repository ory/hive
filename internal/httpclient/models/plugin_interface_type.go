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

// PluginInterfaceType PluginInterfaceType PluginInterfaceType PluginInterfaceType plugin interface type
//
// swagger:model PluginInterfaceType
type PluginInterfaceType struct {

	// capability
	// Required: true
	Capability *string `json:"Capability"`

	// prefix
	// Required: true
	Prefix *string `json:"Prefix"`

	// version
	// Required: true
	Version *string `json:"Version"`
}

// Validate validates this plugin interface type
func (m *PluginInterfaceType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCapability(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrefix(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PluginInterfaceType) validateCapability(formats strfmt.Registry) error {

	if err := validate.Required("Capability", "body", m.Capability); err != nil {
		return err
	}

	return nil
}

func (m *PluginInterfaceType) validatePrefix(formats strfmt.Registry) error {

	if err := validate.Required("Prefix", "body", m.Prefix); err != nil {
		return err
	}

	return nil
}

func (m *PluginInterfaceType) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("Version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this plugin interface type based on context it is used
func (m *PluginInterfaceType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PluginInterfaceType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PluginInterfaceType) UnmarshalBinary(b []byte) error {
	var res PluginInterfaceType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
