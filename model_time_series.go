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
import (
	"time"
)

type TimeSeries struct {
	Data []int64 `json:"data"`
	Max int64 `json:"max"`
	Metric BasicMetric `json:"metric"`
	Min int64 `json:"min"`
	StartTime time.Time `json:"startTime"`
	TickInterval int32 `json:"tickInterval"`
	Total int64 `json:"total"`
}
