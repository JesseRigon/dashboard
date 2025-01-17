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

// ApplicationInstallationListItemSpec ApplicationInstallationListItemSpec is the object representing an ApplicationInstallationListItemSpec.
//
// swagger:model ApplicationInstallationListItemSpec
type ApplicationInstallationListItemSpec struct {

	// Labels can contain metadata about the application, such as the owner who manages it.
	Labels map[string]string `json:"labels,omitempty"`

	// application ref
	ApplicationRef *ApplicationRef `json:"applicationRef,omitempty"`

	// namespace
	Namespace *NamespaceSpec `json:"namespace,omitempty"`
}

// Validate validates this application installation list item spec
func (m *ApplicationInstallationListItemSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplicationRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationInstallationListItemSpec) validateApplicationRef(formats strfmt.Registry) error {
	if swag.IsZero(m.ApplicationRef) { // not required
		return nil
	}

	if m.ApplicationRef != nil {
		if err := m.ApplicationRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("applicationRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("applicationRef")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationInstallationListItemSpec) validateNamespace(formats strfmt.Registry) error {
	if swag.IsZero(m.Namespace) { // not required
		return nil
	}

	if m.Namespace != nil {
		if err := m.Namespace.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namespace")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("namespace")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this application installation list item spec based on the context it is used
func (m *ApplicationInstallationListItemSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApplicationRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNamespace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationInstallationListItemSpec) contextValidateApplicationRef(ctx context.Context, formats strfmt.Registry) error {

	if m.ApplicationRef != nil {
		if err := m.ApplicationRef.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("applicationRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("applicationRef")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationInstallationListItemSpec) contextValidateNamespace(ctx context.Context, formats strfmt.Registry) error {

	if m.Namespace != nil {
		if err := m.Namespace.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namespace")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("namespace")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationInstallationListItemSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationInstallationListItemSpec) UnmarshalBinary(b []byte) error {
	var res ApplicationInstallationListItemSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
