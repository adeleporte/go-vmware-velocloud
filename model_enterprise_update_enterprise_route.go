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

type EnterpriseUpdateEnterpriseRoute struct {
	EnterpriseId int32 `json:"enterpriseId,omitempty"`
	Original EnterpriseRouteCollection `json:"original,omitempty"`
	Updated EnterpriseRouteCollection `json:"updated,omitempty"`
}
