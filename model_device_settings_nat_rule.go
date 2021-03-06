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

type DeviceSettingsNatRule struct {
	Type              string `json:"type,omitempty"`
	Description       string `json:"description,omitempty"`
	InsideCidrIp      string `json:"insideCidrIp"`
	InsideCidrPrefix  int32  `json:"insideCidrPrefix"`
	InsideNetmask     string `json:"insideNetmask"`
	InsidePort        int32  `json:"insidePort,omitempty"`
	OutsideCidrIp     string `json:"outsideCidrIp"`
	OutsideCidrPrefix int32  `json:"outsideCidrPrefix"`
	OutsideNetmask    string `json:"outsideNetmask"`
	OutsidePort       int32  `json:"outsidePort,omitempty"`
}
