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

type FirewallDataServicesSnmp struct {
	Enabled bool `json:"enabled"`
	// List of IP addresses allowed SNMP access
	AllowSelectedIp []string `json:"allowSelectedIp,omitempty"`
	RuleLogicalId string `json:"ruleLogicalId,omitempty"`
}
