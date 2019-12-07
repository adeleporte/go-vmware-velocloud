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

type DeviceSettingsVpn struct {
	Enabled          bool                              `json:"enabled"`
	EdgeToDataCenter bool                              `json:"edgeToDataCenter"`
	Ref              string                            `json:"ref,omitempty"`
	IsolationGroupId string                            `json:"isolationGroupId,omitempty"`
	EdgeToEdgeHub    DeviceSettingsVpnEdgeToEdgeHub    `json:"edgeToEdgeHub,omitempty"`
	EdgeToEdge       bool                              `json:"edgeToEdge"`
	EdgeToEdgeDetail DeviceSettingsVpnEdgeToEdgeDetail `json:"edgeToEdgeDetail,omitempty"`
}