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

type DeviceSettingsLanNetwork struct {
	Override        bool                               `json:"override,omitempty"`
	Advertise       bool                               `json:"advertise,omitempty"`
	PingResponse    bool                               `json:"pingResponse,omitempty"`
	BindEdgeAddress bool                               `json:"bindEdgeAddress,omitempty"`
	BaseDhcpAddr    int32                              `json:"baseDhcpAddr,omitempty"`
	CidrIp          string                             `json:"cidrIp,omitempty"`
	CidrPrefix      int32                              `json:"cidrPrefix,omitempty"`
	Cost            int32                              `json:"cost,omitempty"`
	Dhcp            *DeviceSettingsDhcp                `json:"dhcp,omitempty"`
	Disabled        bool                               `json:"disabled,omitempty"`
	Interfaces      []string                           `json:"interfaces,omitempty"`
	Multicast       *DeviceSettingsLanNetworkMulticast `json:"multicast,omitempty"`
	Name            string                             `json:"name,omitempty"`
	Netmask         string                             `json:"netmask,omitempty"`
	NumDhcpAddr     int32                              `json:"numDhcpAddr,omitempty"`
	Ospf            *DeviceSettingsLanNetworkOspf      `json:"ospf,omitempty"`
	SegmentId       int32                              `json:"segmentId"`
	StaticReserved  int32                              `json:"staticReserved,omitempty"`
	VlanId          int32                              `json:"vlanId,omitempty"`
}
