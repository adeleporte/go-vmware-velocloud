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

type CloudSecurityServiceSiteDataZscalerGre struct {
	CustomSourceIp string `json:"customSourceIp,omitempty"`
	LinkInternalLogicalId string `json:"linkInternalLogicalId,omitempty"`
	PrimaryAddressing CloudSecurityServiceSiteDataZscalerGrePrimaryAddressing `json:"primaryAddressing,omitempty"`
	SecondaryAddressing CloudSecurityServiceSiteDataZscalerGrePrimaryAddressing `json:"secondaryAddressing,omitempty"`
	UseCustomSourceIp bool `json:"useCustomSourceIp,omitempty"`
}
