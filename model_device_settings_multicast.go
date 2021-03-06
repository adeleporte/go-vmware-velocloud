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

type DeviceSettingsMulticast struct {
	Enabled                      bool                                   `json:"enabled"`
	Override                     bool                                   `json:"override"`
	IgmpHostQueryIntervalSeconds int32                                  `json:"igmpHostQueryIntervalSeconds,omitempty"`
	IgmpMaxQueryResponse         int32                                  `json:"igmpMaxQueryResponse,omitempty"`
	PimKeepAliveTimerSeconds     int32                                  `json:"pimKeepAliveTimerSeconds,omitempty"`
	PimOnWanOverlay              DeviceSettingsMulticastPimOnWanOverlay `json:"pimOnWanOverlay,omitempty"`
	PimPruneIntervalSeconds      int32                                  `json:"pimPruneIntervalSeconds,omitempty"`
	Rp                           DeviceSettingsMulticastRp              `json:"rp"`
}
