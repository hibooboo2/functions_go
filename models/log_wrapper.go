package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LogWrapper log wrapper
// swagger:model LogWrapper
type LogWrapper struct {

	// Call log entry.
	// Required: true
	Log *Log `json:"log"`
}

// Validate validates this log wrapper
func (m *LogWrapper) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLog(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogWrapper) validateLog(formats strfmt.Registry) error {

	if err := validate.Required("log", "body", m.Log); err != nil {
		return err
	}

	if m.Log != nil {

		if err := m.Log.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LogWrapper) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogWrapper) UnmarshalBinary(b []byte) error {
	var res LogWrapper
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
