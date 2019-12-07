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

type DeviceSettingsDataRadioSettingsRadios struct {
	RadioId   int32  `json:"radioId,omitempty"`
	IsEnabled bool   `json:"isEnabled,omitempty"`
	Name      string `json:"name,omitempty"`
	Band      string `json:"band,omitempty"`
	Channel   string `json:"channel,omitempty"`
	Width     string `json:"width,omitempty"`
	Mode      string `json:"mode,omitempty"`
}
