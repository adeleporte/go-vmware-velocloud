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

type FirewallSegment struct {
	FirewallLoggingEnabled bool `json:"firewall_logging_enabled"`
	StatefulFirewallEnabled bool `json:"stateful_firewall_enabled,omitempty"`
	Outbound []FirewallOutboundRule `json:"outbound"`
	Segment ConfigurationModuleSegmentMetadata `json:"segment"`
}