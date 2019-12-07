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

type FirewallInboundRule struct {
	Name string `json:"name,omitempty"`
	Match FirewallRuleMatch `json:"match"`
	Action FirewallInboundRuleAction `json:"action"`
	RuleLogicalId string `json:"ruleLogicalId,omitempty"`
}
