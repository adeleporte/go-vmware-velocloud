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

type DeviceSettingsDhcp struct {
	Enabled          bool                       `json:"enabled,omitempty"`
	DhcpRelay        DeviceSettingsDhcpRelay    `json:"dhcpRelay,omitempty"`
	LeaseTimeSeconds int32                      `json:"leaseTimeSeconds,omitempty"`
	BaseDhcpAddr     int32                      `json:"baseDhcpAddr,omitempty"`
	NumDhcpAddr      int32                      `json:"numDhcpAddr,omitempty"`
	Options          []DeviceSettingsDhcpOption `json:"options,omitempty"`
	Override         bool                       `json:"override,omitempty"`
}
