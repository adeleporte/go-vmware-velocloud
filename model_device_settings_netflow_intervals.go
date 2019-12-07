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

type DeviceSettingsNetflowIntervals struct {
	FlowStats int32 `json:"flowStats,omitempty"`
	FlowLinkStats int32 `json:"flowLinkStats,omitempty"`
	VrfTable int32 `json:"vrfTable,omitempty"`
	ApplicationTable int32 `json:"applicationTable,omitempty"`
	InterfaceTable int32 `json:"interfaceTable,omitempty"`
	LinkTable int32 `json:"linkTable,omitempty"`
}