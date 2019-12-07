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

type EdgeDeviceSettingsDataOspf struct {
	Area int32 `json:"area,omitempty"`
	Authentication bool `json:"authentication,omitempty"`
	AuthId int32 `json:"authId,omitempty"`
	AuthPassphrase string `json:"authPassphrase,omitempty"`
	Cost int32 `json:"cost,omitempty"`
	DeadTimer int32 `json:"deadTimer,omitempty"`
	Mode string `json:"mode,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	HelloTimer int32 `json:"helloTimer,omitempty"`
	InboundRouteLearning EdgeDeviceSettingsDataOspfInboundRouteLearning `json:"inboundRouteLearning,omitempty"`
	Md5Authentication bool `json:"md5Authentication,omitempty"`
	MTU int32 `json:"MTU,omitempty"`
	OutboundRouteAdvertisement EdgeDeviceSettingsDataOspfInboundRouteLearning `json:"outboundRouteAdvertisement,omitempty"`
	Passive bool `json:"passive,omitempty"`
	VlanId int32 `json:"vlanId,omitempty"`
}
