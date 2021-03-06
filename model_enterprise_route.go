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

type EnterpriseRoute struct {
	Type string `json:"type,omitempty"`
	ExitType string `json:"exitType,omitempty"`
	EdgeId int32 `json:"edgeId,omitempty"`
	EdgeName string `json:"edgeName,omitempty"`
	ProfileId int32 `json:"profileId,omitempty"`
	CidrIp string `json:"cidrIp,omitempty"`
	Cost int32 `json:"cost,omitempty"`
	Advertise bool `json:"advertise,omitempty"`
}
