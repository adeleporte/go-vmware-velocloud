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

type DeviceSettingsDataSnmpSnmpv2c struct {
	Enabled   bool     `json:"enabled,omitempty"`
	Community string   `json:"community,omitempty"`
	AllowedIp []string `json:"allowedIp,omitempty"`
}