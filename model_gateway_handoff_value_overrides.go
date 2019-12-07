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

type GatewayHandoffValueOverrides struct {
	VLAN map[string]GatewayHandoffValueOverridesVlan `json:"VLAN,omitempty"`
	Bgp map[string]GatewayHandoffValueBgp `json:"bgp,omitempty"`
	BgpInboundMap map[string]GatewayHandoffBgpRulesMap `json:"bgpInboundMap,omitempty"`
	BgpOutboundMap map[string]GatewayHandoffBgpRulesMap `json:"bgpOutboundMap,omitempty"`
	LocalAddress map[string]GatewayHandoffValueOverridesLocalAddress `json:"localAddress,omitempty"`
	Subnets map[string][]map[string]interface{} `json:"subnets,omitempty"`
}