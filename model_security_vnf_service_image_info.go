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

type SecurityVnfServiceImageInfo struct {
	DownloadType string `json:"downloadType,omitempty"`
	FileLocation string `json:"fileLocation,omitempty"`
	FileChecksum string `json:"fileChecksum,omitempty"`
	FileChecksumType string `json:"fileChecksumType,omitempty"`
	Https EdgeUpdateEdgeCredentialsByConfigurationCredentials `json:"https,omitempty"`
	S3 SecurityVnfServiceImageInfoS3 `json:"s3,omitempty"`
}
