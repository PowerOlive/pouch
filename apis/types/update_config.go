// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UpdateConfig UpdateConfig holds the mutable attributes of a Container. Those attributes can be updated at runtime.
// swagger:model UpdateConfig

type UpdateConfig struct {
	Resources

	// A list of environment variables to set inside the container in the form `["VAR=value", ...]`. A variable without `=` is removed from the environment, rather than to have an empty value.
	//
	Env []string `json:"Env"`

	// List of labels set to container.
	Labels map[string]string `json:"Labels,omitempty"`

	// restart policy
	RestartPolicy *RestartPolicy `json:"RestartPolicy,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *UpdateConfig) UnmarshalJSON(raw []byte) error {

	var aO0 Resources
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Resources = aO0

	var data struct {
		Env []string `json:"Env,omitempty"`

		Labels map[string]string `json:"Labels,omitempty"`

		RestartPolicy *RestartPolicy `json:"RestartPolicy,omitempty"`
	}
	if err := swag.ReadJSON(raw, &data); err != nil {
		return err
	}

	m.Env = data.Env

	m.Labels = data.Labels

	m.RestartPolicy = data.RestartPolicy

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m UpdateConfig) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	aO0, err := swag.WriteJSON(m.Resources)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var data struct {
		Env []string `json:"Env,omitempty"`

		Labels map[string]string `json:"Labels,omitempty"`

		RestartPolicy *RestartPolicy `json:"RestartPolicy,omitempty"`
	}

	data.Env = m.Env

	data.Labels = m.Labels

	data.RestartPolicy = m.RestartPolicy

	jsonData, err := swag.WriteJSON(data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jsonData)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this update config
func (m *UpdateConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.Resources.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnv(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRestartPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateConfig) validateEnv(formats strfmt.Registry) error {

	if swag.IsZero(m.Env) { // not required
		return nil
	}

	return nil
}

func (m *UpdateConfig) validateRestartPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.RestartPolicy) { // not required
		return nil
	}

	if m.RestartPolicy != nil {

		if err := m.RestartPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RestartPolicy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateConfig) UnmarshalBinary(b []byte) error {
	var res UpdateConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
