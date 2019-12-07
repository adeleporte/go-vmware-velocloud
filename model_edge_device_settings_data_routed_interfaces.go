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

type EdgeDeviceSettingsDataRoutedInterfaces struct {
	Addressing EdgeDeviceSettingsDataAddressing `json:"addressing,omitempty"`
	Advertise bool `json:"advertise,omitempty"`
	PingResponse bool `json:"pingResponse,omitempty"`
	Disabled bool `json:"disabled,omitempty"`
	DhcpServer EdgeDeviceSettingsDataDhcpServer `json:"dhcpServer,omitempty"`
	EncryptOverlay bool `json:"encryptOverlay,omitempty"`
	L2 EdgeDeviceSettingsDataL2 `json:"l2,omitempty"`
	Name string `json:"name,omitempty"`
	NatDirect bool `json:"natDirect,omitempty"`
	Ospf EdgeDeviceSettingsDataOspf `json:"ospf,omitempty"`
	Override bool `json:"override,omitempty"`
	Subinterfaces []EdgeDeviceSettingsDataSubinterfaces `json:"subinterfaces,omitempty"`
	// static only
	VlanId int32 `json:"vlanId,omitempty"`
	WanOverlay string `json:"wanOverlay,omitempty"`
	Trusted bool `json:"trusted,omitempty"`
	Rpf string `json:"rpf,omitempty"`
	UnderlayAccounting bool `json:"underlayAccounting,omitempty"`
}