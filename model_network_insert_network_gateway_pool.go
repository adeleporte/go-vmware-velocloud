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

type NetworkInsertNetworkGatewayPool struct {
	NetworkId int32 `json:"networkId"`
	EnterpriseProxyId int32 `json:"enterpriseProxyId,omitempty"`
	Name string `json:"name"`
	Description string `json:"description,omitempty"`
	HandOffType string `json:"handOffType,omitempty"`
}