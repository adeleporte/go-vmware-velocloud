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

type EnterpriseUpdateEnterpriseRouteConfigurationDataEdge struct {
	Ospf EnterpriseUpdateEnterpriseRouteConfigurationDataEdgeOspf `json:"ospf"`
	Bgp EnterpriseUpdateEnterpriseRouteConfigurationDataEdgeBgp `json:"bgp"`
	Assigned EnterpriseUpdateEnterpriseRouteConfigurationDataEdgeAssigned `json:"assigned"`
}
