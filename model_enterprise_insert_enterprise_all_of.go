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

type EnterpriseInsertEnterpriseAllOf struct {
	ConfigurationId int32 `json:"configurationId"`
	EnableEnterpriseDelegationToOperator bool `json:"enableEnterpriseDelegationToOperator,omitempty"`
	EnableEnterpriseDelegationToProxy bool `json:"enableEnterpriseDelegationToProxy,omitempty"`
	EnableEnterpriseUserManagementDelegationToOperator bool `json:"enableEnterpriseUserManagementDelegationToOperator,omitempty"`
	Licenses EnterpriseInsertEnterpriseAllOfLicenses `json:"licenses,omitempty"`
}
