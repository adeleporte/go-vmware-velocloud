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

type EnterpriseUserUpdateEnterpriseUser1 struct {
	Update EnterpriseUserUpdateEnterpriseUser `json:"_update"`
	Id int32 `json:"id,omitempty"`
	EnterpriseId int32 `json:"enterpriseId,omitempty"`
	Username string `json:"username,omitempty"`
}
