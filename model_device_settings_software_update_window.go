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

type DeviceSettingsSoftwareUpdateWindow struct {
	Day int32 `json:"day,omitempty"`
	BeginHour int32 `json:"beginHour,omitempty"`
	EndHour int32 `json:"endHour,omitempty"`
}
