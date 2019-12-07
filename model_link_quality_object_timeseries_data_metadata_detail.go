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

type LinkQualityObjectTimeseriesDataMetadataDetail struct {
	LatencyMsRx int32 `json:"latencyMsRx,omitempty"`
	LatencyMsTx int32 `json:"latencyMsTx,omitempty"`
	LossPctRx int32 `json:"lossPctRx,omitempty"`
	LossPctTx int32 `json:"lossPctTx,omitempty"`
	JitterMsRx int32 `json:"jitterMsRx,omitempty"`
	JitterMsTx int32 `json:"jitterMsTx,omitempty"`
}
