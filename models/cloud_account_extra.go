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

// CloudAccountExtra cloud account extra
//
// swagger:model CloudAccountExtra
type CloudAccountExtra struct {

	// cloud account information (currently only supports AWS)
	// Required: true
	Account *CloudAccount `json:"account"`

	// default cloud service monitoring settings
	// Required: true
	Default *CloudServiceSettings `json:"default"`

	// Cloud devices to monitor in the group
	// Read Only: true
	Devices *CloudDevice `json:"devices,omitempty"`

	// Cloud account services to monitor
	// Required: true
	Services *CloudServices `json:"services"`
}

// Validate validates this cloud account extra
func (m *CloudAccountExtra) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefault(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudAccountExtra) validateAccount(formats strfmt.Registry) error {

	if err := validate.Required("account", "body", m.Account); err != nil {
		return err
	}

	if m.Account != nil {
		if err := m.Account.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("account")
			}
			return err
		}
	}

	return nil
}

func (m *CloudAccountExtra) validateDefault(formats strfmt.Registry) error {

	if err := validate.Required("default", "body", m.Default); err != nil {
		return err
	}

	if m.Default != nil {
		if err := m.Default.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("default")
			}
			return err
		}
	}

	return nil
}

func (m *CloudAccountExtra) validateDevices(formats strfmt.Registry) error {
	if swag.IsZero(m.Devices) { // not required
		return nil
	}

	if m.Devices != nil {
		if err := m.Devices.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("devices")
			}
			return err
		}
	}

	return nil
}

func (m *CloudAccountExtra) validateServices(formats strfmt.Registry) error {

	if err := validate.Required("services", "body", m.Services); err != nil {
		return err
	}

	if m.Services != nil {
		if err := m.Services.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("services")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cloud account extra based on the context it is used
func (m *CloudAccountExtra) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDefault(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDevices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudAccountExtra) contextValidateAccount(ctx context.Context, formats strfmt.Registry) error {

	if m.Account != nil {
		if err := m.Account.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("account")
			}
			return err
		}
	}

	return nil
}

func (m *CloudAccountExtra) contextValidateDefault(ctx context.Context, formats strfmt.Registry) error {

	if m.Default != nil {
		if err := m.Default.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("default")
			}
			return err
		}
	}

	return nil
}

func (m *CloudAccountExtra) contextValidateDevices(ctx context.Context, formats strfmt.Registry) error {

	if m.Devices != nil {
		if err := m.Devices.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("devices")
			}
			return err
		}
	}

	return nil
}

func (m *CloudAccountExtra) contextValidateServices(ctx context.Context, formats strfmt.Registry) error {

	if m.Services != nil {
		if err := m.Services.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("services")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudAccountExtra) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudAccountExtra) UnmarshalBinary(b []byte) error {
	var res CloudAccountExtra
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
