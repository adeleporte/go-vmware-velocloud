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

type ImageUpdateScheduledTime struct {
	DayOfWeek int32 `json:"dayOfWeek"`
	Specified bool `json:"specified"`
	TimeOfDayMins int32 `json:"timeOfDayMins"`
	UseEdgeTimeZone bool `json:"useEdgeTimeZone"`
}
