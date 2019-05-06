// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GoogleMapWidget google map widget
// swagger:model GoogleMapWidget
type GoogleMapWidget struct {
	dashboardIdField *int32

	descriptionField string

	idField int32

	intervalField int32

	lastUpdatedByField string

	lastUpdatedOnField int64

	nameField *string

	themeField string

	timescaleField string

	userPermissionField string

	// Whether ACKed alerts should be displayed
	AckChecked bool `json:"ackChecked,omitempty"`

	// Whether critical alerts should be displayed
	DisplayCriticalAlert bool `json:"displayCriticalAlert,omitempty"`

	// Whether error alerts should be displayed
	DisplayErrorAlert bool `json:"displayErrorAlert,omitempty"`

	// Whether warning alerts should be displayed
	DisplayWarnAlert bool `json:"displayWarnAlert,omitempty"`

	// The points info
	// Required: true
	MapPoints []*PointSource `json:"mapPoints"`

	// Whether alerts occuring during an SDT period should be displayed
	SDTChecked bool `json:"sdtChecked,omitempty"`
}

// DashboardID gets the dashboard Id of this subtype
func (m *GoogleMapWidget) DashboardID() *int32 {
	return m.dashboardIdField
}

// SetDashboardID sets the dashboard Id of this subtype
func (m *GoogleMapWidget) SetDashboardID(val *int32) {
	m.dashboardIdField = val
}

// Description gets the description of this subtype
func (m *GoogleMapWidget) Description() string {
	return m.descriptionField
}

// SetDescription sets the description of this subtype
func (m *GoogleMapWidget) SetDescription(val string) {
	m.descriptionField = val
}

// ID gets the id of this subtype
func (m *GoogleMapWidget) ID() int32 {
	return m.idField
}

// SetID sets the id of this subtype
func (m *GoogleMapWidget) SetID(val int32) {
	m.idField = val
}

// Interval gets the interval of this subtype
func (m *GoogleMapWidget) Interval() int32 {
	return m.intervalField
}

// SetInterval sets the interval of this subtype
func (m *GoogleMapWidget) SetInterval(val int32) {
	m.intervalField = val
}

// LastUpdatedBy gets the last updated by of this subtype
func (m *GoogleMapWidget) LastUpdatedBy() string {
	return m.lastUpdatedByField
}

// SetLastUpdatedBy sets the last updated by of this subtype
func (m *GoogleMapWidget) SetLastUpdatedBy(val string) {
	m.lastUpdatedByField = val
}

// LastUpdatedOn gets the last updated on of this subtype
func (m *GoogleMapWidget) LastUpdatedOn() int64 {
	return m.lastUpdatedOnField
}

// SetLastUpdatedOn sets the last updated on of this subtype
func (m *GoogleMapWidget) SetLastUpdatedOn(val int64) {
	m.lastUpdatedOnField = val
}

// Name gets the name of this subtype
func (m *GoogleMapWidget) Name() *string {
	return m.nameField
}

// SetName sets the name of this subtype
func (m *GoogleMapWidget) SetName(val *string) {
	m.nameField = val
}

// Theme gets the theme of this subtype
func (m *GoogleMapWidget) Theme() string {
	return m.themeField
}

// SetTheme sets the theme of this subtype
func (m *GoogleMapWidget) SetTheme(val string) {
	m.themeField = val
}

// Timescale gets the timescale of this subtype
func (m *GoogleMapWidget) Timescale() string {
	return m.timescaleField
}

// SetTimescale sets the timescale of this subtype
func (m *GoogleMapWidget) SetTimescale(val string) {
	m.timescaleField = val
}

// Type gets the type of this subtype
func (m *GoogleMapWidget) Type() string {
	return "gmap"
}

// SetType sets the type of this subtype
func (m *GoogleMapWidget) SetType(val string) {

}

// UserPermission gets the user permission of this subtype
func (m *GoogleMapWidget) UserPermission() string {
	return m.userPermissionField
}

// SetUserPermission sets the user permission of this subtype
func (m *GoogleMapWidget) SetUserPermission(val string) {
	m.userPermissionField = val
}

// AckChecked gets the ack checked of this subtype

// DisplayCriticalAlert gets the display critical alert of this subtype

// DisplayErrorAlert gets the display error alert of this subtype

// DisplayWarnAlert gets the display warn alert of this subtype

// MapPoints gets the map points of this subtype

// SDTChecked gets the sdt checked of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *GoogleMapWidget) UnmarshalJSON(raw []byte) error {
	var data struct {

		// Whether ACKed alerts should be displayed
		AckChecked bool `json:"ackChecked,omitempty"`

		// Whether critical alerts should be displayed
		DisplayCriticalAlert bool `json:"displayCriticalAlert,omitempty"`

		// Whether error alerts should be displayed
		DisplayErrorAlert bool `json:"displayErrorAlert,omitempty"`

		// Whether warning alerts should be displayed
		DisplayWarnAlert bool `json:"displayWarnAlert,omitempty"`

		// The points info
		// Required: true
		MapPoints []*PointSource `json:"mapPoints"`

		// Whether alerts occuring during an SDT period should be displayed
		SDTChecked bool `json:"sdtChecked,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		DashboardID *int32 `json:"dashboardId"`

		Description string `json:"description,omitempty"`

		ID int32 `json:"id,omitempty"`

		Interval int32 `json:"interval,omitempty"`

		LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`

		LastUpdatedOn int64 `json:"lastUpdatedOn,omitempty"`

		Name *string `json:"name"`

		Theme string `json:"theme,omitempty"`

		Timescale string `json:"timescale,omitempty"`

		Type string `json:"type"`

		UserPermission string `json:"userPermission,omitempty"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result GoogleMapWidget

	result.dashboardIdField = base.DashboardID

	result.descriptionField = base.Description

	result.idField = base.ID

	result.intervalField = base.Interval

	result.lastUpdatedByField = base.LastUpdatedBy

	result.lastUpdatedOnField = base.LastUpdatedOn

	result.nameField = base.Name

	result.themeField = base.Theme

	result.timescaleField = base.Timescale

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.userPermissionField = base.UserPermission

	result.AckChecked = data.AckChecked

	result.DisplayCriticalAlert = data.DisplayCriticalAlert

	result.DisplayErrorAlert = data.DisplayErrorAlert

	result.DisplayWarnAlert = data.DisplayWarnAlert

	result.MapPoints = data.MapPoints

	result.SDTChecked = data.SDTChecked

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m GoogleMapWidget) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// Whether ACKed alerts should be displayed
		AckChecked bool `json:"ackChecked,omitempty"`

		// Whether critical alerts should be displayed
		DisplayCriticalAlert bool `json:"displayCriticalAlert,omitempty"`

		// Whether error alerts should be displayed
		DisplayErrorAlert bool `json:"displayErrorAlert,omitempty"`

		// Whether warning alerts should be displayed
		DisplayWarnAlert bool `json:"displayWarnAlert,omitempty"`

		// The points info
		// Required: true
		MapPoints []*PointSource `json:"mapPoints"`

		// Whether alerts occuring during an SDT period should be displayed
		SDTChecked bool `json:"sdtChecked,omitempty"`
	}{

		AckChecked: m.AckChecked,

		DisplayCriticalAlert: m.DisplayCriticalAlert,

		DisplayErrorAlert: m.DisplayErrorAlert,

		DisplayWarnAlert: m.DisplayWarnAlert,

		MapPoints: m.MapPoints,

		SDTChecked: m.SDTChecked,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		DashboardID *int32 `json:"dashboardId"`

		Description string `json:"description,omitempty"`

		ID int32 `json:"id,omitempty"`

		Interval int32 `json:"interval,omitempty"`

		LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`

		LastUpdatedOn int64 `json:"lastUpdatedOn,omitempty"`

		Name *string `json:"name"`

		Theme string `json:"theme,omitempty"`

		Timescale string `json:"timescale,omitempty"`

		Type string `json:"type"`

		UserPermission string `json:"userPermission,omitempty"`
	}{

		DashboardID: m.DashboardID(),

		Description: m.Description(),

		ID: m.ID(),

		Interval: m.Interval(),

		LastUpdatedBy: m.LastUpdatedBy(),

		LastUpdatedOn: m.LastUpdatedOn(),

		Name: m.Name(),

		Theme: m.Theme(),

		Timescale: m.Timescale(),

		Type: m.Type(),

		UserPermission: m.UserPermission(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this google map widget
func (m *GoogleMapWidget) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDashboardID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMapPoints(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GoogleMapWidget) validateDashboardID(formats strfmt.Registry) error {

	if err := validate.Required("dashboardId", "body", m.DashboardID()); err != nil {
		return err
	}

	return nil
}

func (m *GoogleMapWidget) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	return nil
}

func (m *GoogleMapWidget) validateMapPoints(formats strfmt.Registry) error {

	if err := validate.Required("mapPoints", "body", m.MapPoints); err != nil {
		return err
	}

	for i := 0; i < len(m.MapPoints); i++ {
		if swag.IsZero(m.MapPoints[i]) { // not required
			continue
		}

		if m.MapPoints[i] != nil {
			if err := m.MapPoints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mapPoints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GoogleMapWidget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GoogleMapWidget) UnmarshalBinary(b []byte) error {
	var res GoogleMapWidget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
