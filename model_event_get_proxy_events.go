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

type EventGetProxyEvents struct {
	Interval Interval `json:"interval,omitempty"`
	Filter Filter1 `json:"filter,omitempty"`
	EnterpriseId []int32 `json:"enterpriseId,omitempty"`
}
