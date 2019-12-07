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

type MetricsGetEdgeDeviceMetricsResultItem struct {
	BytesRx int64 `json:"bytesRx,omitempty"`
	BytesTx int64 `json:"bytesTx,omitempty"`
	FlowCount int32 `json:"flowCount,omitempty"`
	PacketsRx int64 `json:"packetsRx,omitempty"`
	PacketsTx int64 `json:"packetsTx,omitempty"`
	TotalBytes int64 `json:"totalBytes,omitempty"`
	TotalPackets int64 `json:"totalPackets,omitempty"`
	EdgeInfo MetricsGetEdgeDeviceMetricsDeviceEdgeInfo `json:"edgeInfo"`
	Info ClientDevice `json:"info"`
	Name string `json:"name"`
	SourceMac string `json:"sourceMac"`
}
