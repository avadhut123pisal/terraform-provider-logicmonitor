// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CloudServiceSettings cloud service settings
//
// swagger:model CloudServiceSettings
type CloudServiceSettings struct {

	// The NetScan schedule for how frequently the cloud collector should scan/discover new resources in the cloud account. It's format is similar to Linux crontab but doesn't support some complex representations ('-', '/', ',') supported in standard linux crontabs.
	// Format: '*(minute) *(hour) *(day) *(week of month) *(weekday)'
	// Examples: '50 * * * *' means scheduling at 50th minute per hour
	// '50 10 20 * *' means scheduling at 10:50 of the 20th day per month
	// '50 10 * 1 3' means scheduling at wednesday of the first week per month
	CustomNSPSchedule string `json:"customNSPSchedule,omitempty"`

	// How to handle dead cloud devices (deletion from monitoring). Valid options include MANUALLY, IMMEDIATELY, KEEP_1_DAY, KEEP_3_DAYS, KEEP_7_DAYS, KEEP_14_DAYS, and KEEP_30_DAYS
	DeadOperation string `json:"deadOperation,omitempty"`

	// Template used to name devices found within the cloud account within the LM portal
	DeviceDisplayNameTemplate string `json:"deviceDisplayNameTemplate,omitempty"`

	// If monitoring of stopped/terminated hosts is disabled
	DisableStopTerminateHostMonitor bool `json:"disableStopTerminateHostMonitor,omitempty"`

	// If alerting should be disabled when a cloud device is terminated
	DisableTerminatedHostAlerting bool `json:"disableTerminatedHostAlerting,omitempty"`

	// The regions this group will monitor
	// Unique: true
	MonitoringRegionInfos []string `json:"monitoringRegionInfos"`

	// The regions this group will monitor
	// Read Only: true
	// Unique: true
	MonitoringRegions []string `json:"monitoringRegions"`

	// Filter to determine which cloud devices should be monitored based on the device name
	NameFilter []string `json:"nameFilter"`

	// Configuration of 'normal' collector installed to the cloud (ie an EC2 instance)
	NormalCollectorConfig *CloudNormalCollectorConfig `json:"normalCollectorConfig,omitempty"`

	// Whether or not to use all regions (can be used instead of monitoringRegions and monitoringRegionInfos)
	SelectAll bool `json:"selectAll,omitempty"`

	// Tags used to filter whether or not a cloud device is included in this group
	Tags []*CloudTagFilter `json:"tags"`

	// Whether or not to use the default settings for this service. When set to true in a Service's settings, all other fields will be ignored
	// Required: true
	UseDefault *bool `json:"useDefault"`
}

// Validate validates this cloud service settings
func (m *CloudServiceSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMonitoringRegionInfos(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMonitoringRegions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNormalCollectorConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUseDefault(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudServiceSettings) validateMonitoringRegionInfos(formats strfmt.Registry) error {
	if swag.IsZero(m.MonitoringRegionInfos) { // not required
		return nil
	}

	if err := validate.UniqueItems("monitoringRegionInfos", "body", m.MonitoringRegionInfos); err != nil {
		return err
	}

	return nil
}

func (m *CloudServiceSettings) validateMonitoringRegions(formats strfmt.Registry) error {
	if swag.IsZero(m.MonitoringRegions) { // not required
		return nil
	}

	if err := validate.UniqueItems("monitoringRegions", "body", m.MonitoringRegions); err != nil {
		return err
	}

	return nil
}

func (m *CloudServiceSettings) validateNormalCollectorConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.NormalCollectorConfig) { // not required
		return nil
	}

	if m.NormalCollectorConfig != nil {
		if err := m.NormalCollectorConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("normalCollectorConfig")
			}
			return err
		}
	}

	return nil
}

func (m *CloudServiceSettings) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CloudServiceSettings) validateUseDefault(formats strfmt.Registry) error {

	if err := validate.Required("useDefault", "body", m.UseDefault); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this cloud service settings based on the context it is used
func (m *CloudServiceSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMonitoringRegions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNormalCollectorConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudServiceSettings) contextValidateMonitoringRegions(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "monitoringRegions", "body", []string(m.MonitoringRegions)); err != nil {
		return err
	}

	return nil
}

func (m *CloudServiceSettings) contextValidateNormalCollectorConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.NormalCollectorConfig != nil {
		if err := m.NormalCollectorConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("normalCollectorConfig")
			}
			return err
		}
	}

	return nil
}

func (m *CloudServiceSettings) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {
			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudServiceSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudServiceSettings) UnmarshalBinary(b []byte) error {
	var res CloudServiceSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
