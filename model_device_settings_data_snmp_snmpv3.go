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

type DeviceSettingsDataSnmpSnmpv3 struct {
	Enabled bool                                    `json:"enabled,omitempty"`
	Users   []EdgeDeviceSettingsDataSnmpSnmpv3Users `json:"users,omitempty"`
}
