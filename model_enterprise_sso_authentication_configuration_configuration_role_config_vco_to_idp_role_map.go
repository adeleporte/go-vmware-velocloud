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

type EnterpriseSsoAuthenticationConfigurationConfigurationRoleConfigVcoToIdpRoleMap struct {
	EnterpriseReadOnly []string `json:"Enterprise Read Only,omitempty"`
	EnterpriseStandardAdmin []string `json:"Enterprise Standard Admin,omitempty"`
	EnterpriseSuperuser []string `json:"Enterprise Superuser,omitempty"`
	EnterpriseSupport []string `json:"Enterprise Support,omitempty"`
}
