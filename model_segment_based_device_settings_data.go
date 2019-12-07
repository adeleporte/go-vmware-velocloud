/*
 * VeloCloud Orchestrator API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 3.3.2
 * Contact: support@velocloud.net
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type SegmentBasedDeviceSettingsData struct {
	Ha               *DeviceSettingsHa               `json:"ha,omitempty"` //optional
	Lan              DeviceSettingsLan               `json:"lan,omitempty"`
	Models           *map[string]DeviceSettingsModel `json:"models,omitempty"`
	MultiSourceQos   *DeviceSettingsMultiSourceQos   `json:"multiSourceQos,omitempty"`
	Ntp              *DeviceSettingsNtp              `json:"ntp,omitempty"`
	RadioSettings    *DeviceSettingsRadioSettings    `json:"radioSettings,omitempty"`
	RoutedInterfaces []DeviceSettingsRoutedInterface `json:"routedInterfaces,omitempty"`
	Segments         []DeviceSettingsSegment         `json:"segments,omitempty"`
	Snmp             *DeviceSettingsSnmp             `json:"snmp,omitempty"`
	SoftwareUpdate   *DeviceSettingsSoftwareUpdate   `json:"softwareUpdate,omitempty"`
	Tacacs           *DeviceSettingsTacacs           `json:"tacacs,omitempty"`
	Vnfs             *DeviceSettingsVnfs             `json:"vnfs,omitempty"`
}
