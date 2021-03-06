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

type DeviceSettingsSyslogCollector struct {
	Host string `json:"host"`
	Port int32 `json:"port"`
	Protocol string `json:"protocol"`
	Roles []string `json:"roles"`
	Severity string `json:"severity"`
	SourceInterface string `json:"sourceInterface"`
}
