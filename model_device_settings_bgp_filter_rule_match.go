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

type DeviceSettingsBgpFilterRuleMatch struct {
	ExactMatch bool   `json:"exactMatch,omitempty"`
	Type       string `json:"type,omitempty"`
	Value      string `json:"value,omitempty"`
}
