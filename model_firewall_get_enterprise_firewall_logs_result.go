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

type FirewallGetEnterpriseFirewallLogsResult struct {
	MetaData ListMetadata `json:"metaData,omitempty"`
	Data []FirewallLogData `json:"data,omitempty"`
	Rules map[string]FirewallRule `json:"rules,omitempty"`
}
