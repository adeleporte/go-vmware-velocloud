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

type DeviceSettingsSubinterface struct {
	Override bool `json:"override,omitempty"`
	Addressing DeviceSettingsRoutedInterfaceAddressing `json:"addressing,omitempty"`
	Advertise bool `json:"advertise,omitempty"`
	PingResponse bool `json:"pingResponse,omitempty"`
	Disabled bool `json:"disabled,omitempty"`
	DhcpServer DeviceSettingsDhcp `json:"dhcpServer,omitempty"`
	EncryptOverlay bool `json:"encryptOverlay,omitempty"`
	L2 DeviceSettingsRoutedInterfaceL2 `json:"l2,omitempty"`
	Multicast DeviceSettingsRoutedInterfaceMulticast `json:"multicast,omitempty"`
	Name string `json:"name,omitempty"`
	NatDirect bool `json:"natDirect,omitempty"`
	Ospf DeviceSettingsRoutedInterfaceOspf `json:"ospf,omitempty"`
	RadiusAuthentication DeviceSettingsRoutedInterfaceRadiusAuthentication `json:"radiusAuthentication,omitempty"`
	SegmentId int32 `json:"segmentId,omitempty"`
	Subinterfaces []DeviceSettingsSubinterface `json:"subinterfaces,omitempty"`
	Rpf string `json:"rpf,omitempty"`
	Trusted bool `json:"trusted,omitempty"`
	UnderlayAccounting bool `json:"underlayAccounting,omitempty"`
	VlanId int32 `json:"vlanId,omitempty"`
	WanOverlay string `json:"wanOverlay,omitempty"`
	SubinterfaceId int32 `json:"subinterfaceId,omitempty"`
	SubinterfaceType string `json:"subinterfaceType,omitempty"`
}
