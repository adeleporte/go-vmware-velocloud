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

type EdgeSetEdgeHandOffGatewaysGateways struct {
	Primary int32 `json:"primary"`
	PrimaryIpsecDetail GatewayHandoffIpsecGatewayDetail `json:"primaryIpsecDetail,omitempty"`
	Secondary int32 `json:"secondary,omitempty"`
	SecondaryIpsecDetail GatewayHandoffIpsecGatewayDetail `json:"secondaryIpsecDetail,omitempty"`
}
