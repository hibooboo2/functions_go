package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Route route
// swagger:model Route
type Route struct {

	// Route configuration - overrides application configuration
	Config map[string]string `json:"config,omitempty"`

	// Payload format sent into function.
	Format interface{} `json:"format,omitempty"`

	// Map of http headers that will be sent with the response
	Headers map[string][]string `json:"headers,omitempty"`

	// Hot functions idle timeout before termination. Value in Seconds
	IDLETimeout *int64 `json:"idle_timeout,omitempty"`

	// Name of Docker image to use in this route. You should include the image tag, which should be a version number, to be more accurate. Can be overridden on a per route basis with route.image.
	Image string `json:"image,omitempty"`

	// Maximum number of hot functions concurrency
	MaxConcurrency int32 `json:"max_concurrency,omitempty"`

	// Max usable memory for this route (MiB).
	Memory int64 `json:"memory,omitempty"`

	// URL path that will be matched to this route
	// Read Only: true
	Path string `json:"path,omitempty"`

	// Timeout for executions of this route. Value in Seconds
	Timeout *int64 `json:"timeout,omitempty"`

	// Route type
	Type interface{} `json:"type,omitempty"`
}

// Validate validates this route
func (m *Route) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFormat(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHeaders(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var routeTypeFormatPropEnum []interface{}

func init() {
	var res []interface{}
	if err := json.Unmarshal([]byte(`["default","http","json"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		routeTypeFormatPropEnum = append(routeTypeFormatPropEnum, v)
	}
}

// prop value enum
func (m *Route) validateFormatEnum(path, location string, value interface{}) error {
	if err := validate.Enum(path, location, value, routeTypeFormatPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Route) validateFormat(formats strfmt.Registry) error {

	if swag.IsZero(m.Format) { // not required
		return nil
	}

	return nil
}

func (m *Route) validateHeaders(formats strfmt.Registry) error {

	if swag.IsZero(m.Headers) { // not required
		return nil
	}

	if swag.IsZero(m.Headers) { // not required
		return nil
	}

	return nil
}

var routeTypeTypePropEnum []interface{}

func init() {
	var res []interface{}
	if err := json.Unmarshal([]byte(`["sync","async"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		routeTypeTypePropEnum = append(routeTypeTypePropEnum, v)
	}
}

// prop value enum
func (m *Route) validateTypeEnum(path, location string, value interface{}) error {
	if err := validate.Enum(path, location, value, routeTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Route) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Route) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Route) UnmarshalBinary(b []byte) error {
	var res Route
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}