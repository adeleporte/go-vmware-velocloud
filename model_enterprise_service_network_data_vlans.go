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

type EnterpriseServiceNetworkDataVlans struct {
	Name string `json:"name,omitempty"`
	VlanId int32 `json:"vlanId,omitempty"`
	Advertise bool `json:"advertise,omitempty"`
	Cost int32 `json:"cost,omitempty"`
	StaticReserved int32 `json:"staticReserved,omitempty"`
	Dhcp EnterpriseServiceNetworkDataDhcp `json:"dhcp,omitempty"`
}
