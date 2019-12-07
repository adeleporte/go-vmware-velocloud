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

type EdgeDeviceSettingsDataMultiSourceQos struct {
	Enable bool `json:"enable,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	HighRatio int32 `json:"highRatio,omitempty"`
	NormalRatio int32 `json:"normalRatio,omitempty"`
	LowRatio int32 `json:"lowRatio,omitempty"`
	MaxCapThreshold int32 `json:"maxCapThreshold,omitempty"`
	MinCapThreshold int32 `json:"minCapThreshold,omitempty"`
}
