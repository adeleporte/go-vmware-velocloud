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

type MonitoringGetAggregateEdgeLinkMetrics struct {
	Enterprises []int32 `json:"enterprises,omitempty"`
	Metrics []EdgeLinkMetric `json:"metrics,omitempty"`
	Interval Interval `json:"interval,omitempty"`
}