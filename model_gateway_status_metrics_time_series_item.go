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

type GatewayStatusMetricsTimeSeriesItem struct {
	TunnelCount int32 `json:"tunnelCount,omitempty"`
	MemoryPct float32 `json:"memoryPct,omitempty"`
	FlowCount float32 `json:"flowCount,omitempty"`
	CpuPct float32 `json:"cpuPct,omitempty"`
	HandoffQueueDrops int32 `json:"handoffQueueDrops,omitempty"`
	ConnectedEdges int32 `json:"connectedEdges,omitempty"`
	StartTime time.Time `json:"startTime"`
}