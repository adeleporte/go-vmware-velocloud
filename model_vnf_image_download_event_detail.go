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

type VnfImageDownloadEventDetail struct {
	LogicalId string `json:"logicalId,omitempty"`
	VnfType string `json:"vnfType,omitempty"`
	Vendor string `json:"vendor,omitempty"`
	DownloadType string `json:"downloadType,omitempty"`
	EdgeSerialNumber string `json:"edgeSerialNumber,omitempty"`
	IsEdgeActive bool `json:"isEdgeActive,omitempty"`
	Status string `json:"status,omitempty"`
	Description string `json:"description,omitempty"`
	Service LogicalidReference `json:"service,omitempty"`
}
