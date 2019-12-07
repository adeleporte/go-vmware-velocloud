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

type DeviceSettingsStaticRoute struct {
	Destination string `json:"destination,omitempty"`
	Netmask string `json:"netmask,omitempty"`
	SourceIp string `json:"sourceIp,omitempty"`
	Gateway string `json:"gateway,omitempty"`
	Cost int32 `json:"cost,omitempty"`
	Preferred bool `json:"preferred,omitempty"`
	Description string `json:"description,omitempty"`
	CidrPrefix string `json:"cidrPrefix,omitempty"`
	WanInterface string `json:"wanInterface,omitempty"`
	IcmpProbeLogicalId string `json:"icmpProbeLogicalId,omitempty"`
	VlanId int32 `json:"vlanId,omitempty"`
	Advertise bool `json:"advertise,omitempty"`
	SubinterfaceId int32 `json:"subinterfaceId,omitempty"`
}