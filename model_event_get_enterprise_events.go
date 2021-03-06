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

type EventGetEnterpriseEvents struct {
	EnterpriseId int32 `json:"enterpriseId,omitempty"`
	Interval Interval `json:"interval,omitempty"`
	Filter Filter `json:"filter,omitempty"`
	EdgeId []int32 `json:"edgeId,omitempty"`
}
