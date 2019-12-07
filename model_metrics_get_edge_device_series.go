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

type MetricsGetEdgeDeviceSeries struct {
	Id int32 `json:"id,omitempty"`
	EdgeId int32 `json:"edgeId,omitempty"`
	EnterpriseId int32 `json:"enterpriseId,omitempty"`
	Interval Interval `json:"interval"`
	Metrics []BasicMetric `json:"metrics,omitempty"`
	Sort BasicMetric `json:"sort,omitempty"`
	Limit int32 `json:"limit,omitempty"`
	MaxSamples int32 `json:"maxSamples,omitempty"`
	Devices []string `json:"devices,omitempty"`
}
