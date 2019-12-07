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

type DeviceSettingsRoutedInterfaceMulticast struct {
	Igmp DeviceSettingsLanNetworkMulticastIgmp `json:"igmp,omitempty"`
	IgmpHostQueryIntervalSeconds int32 `json:"igmpHostQueryIntervalSeconds,omitempty"`
	IgmpMaxQueryResponse int32 `json:"igmpMaxQueryResponse,omitempty"`
	Pim DeviceSettingsLanNetworkMulticastPim `json:"pim,omitempty"`
	PimKeepAliveTimerSeconds int32 `json:"pimKeepAliveTimerSeconds,omitempty"`
	PimPruneIntervalSeconds int32 `json:"pimPruneIntervalSeconds,omitempty"`
	PimHelloTimerSeconds int32 `json:"pimHelloTimerSeconds,omitempty"`
}
