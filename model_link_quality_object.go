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

type LinkQualityObject struct {
	Distribution map[string]map[string]float32 `json:"distribution"`
	SampleCount int32 `json:"sampleCount"`
	SampleLength int32 `json:"sampleLength"`
	Score map[string]float32 `json:"score"`
	Timeseries []LinkQualityObjectTimeseriesData `json:"timeseries"`
	TotalScore float32 `json:"totalScore"`
}
