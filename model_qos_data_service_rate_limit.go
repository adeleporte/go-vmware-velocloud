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

type QosDataServiceRateLimit struct {
	Enabled bool `json:"enabled"`
	InputType string `json:"inputType,omitempty"`
	Value int32 `json:"value,omitempty"`
}