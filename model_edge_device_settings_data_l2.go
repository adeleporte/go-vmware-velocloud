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

type EdgeDeviceSettingsDataL2 struct {
	Autonegotiation bool `json:"autonegotiation,omitempty"`
	Speed string `json:"speed,omitempty"`
	Duplex string `json:"duplex,omitempty"`
	MTU int32 `json:"MTU,omitempty"`
}
