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

type EnterpriseGetEnterpriseEdgesResultItemAllOfConfigurationEnterprise struct {
	Id int32 `json:"id,omitempty"`
	Modules []EnterpriseGetEnterpriseEdgesResultItemAllOfConfigurationEnterpriseModules `json:"modules,omitempty"`
	Name string `json:"name,omitempty"`
}
