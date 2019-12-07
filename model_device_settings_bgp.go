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

type DeviceSettingsBgp struct {
	Enabled                bool                                `json:"enabled"`
	Override               bool                                `json:"override"`
	ASN                    string                              `json:"asn"`
	ConnectedRoutes        bool                                `json:"connectedRoutes"`
	DefaultRoute           DeviceSettingsBgpDefaultRoute       `json:"defaultRoute"`
	DisableASPathCarryOver bool                                `json:"disableASPathCarryOver"`
	Filters                []DeviceSettingsBgpFilter           `json:"filters"`
	Holdtime               string                              `json:"holdtime"`
	IsEdge                 bool                                `json:"isEdge"`
	Keepalive              string                              `json:"keepalive"`
	Neighbors              []DeviceSettingsBgpNeighbor         `json:"neighbors"`
	Networks               []DeviceSettingsBgpNetwork          `json:"networks"`
	Ospf                   DeviceSettingsBgpOspfRedistribution `json:"ospf"`
	OverlayPrefix          bool                                `json:"overlayPrefix"`
	PropagateUplink        bool                                `json:"propagateUplink"`
	RouterId               string                              `json:"routerId"`
	UplinkCommunity        string                              `json:"uplinkCommunity"`
}
