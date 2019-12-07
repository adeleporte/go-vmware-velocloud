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

type OperatorEvent struct {
	Id int32 `json:"id,omitempty"`
	EventTime time.Time `json:"eventTime,omitempty"`
	Event string `json:"event,omitempty"`
	Category string `json:"category,omitempty"`
	Severity string `json:"severity,omitempty"`
	Message string `json:"message,omitempty"`
	Detail string `json:"detail"`
	OperatorUsername string `json:"operatorUsername"`
	NetworkName string `json:"networkName"`
	GatewayName string `json:"gatewayName"`
}