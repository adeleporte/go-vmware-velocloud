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

type BgpPeerStatusRecord struct {
	Timestamp time.Time `json:"timestamp,omitempty"`
	State string `json:"state,omitempty"`
	MsgRcvd int32 `json:"msgRcvd,omitempty"`
	PfxRcvd int32 `json:"pfxRcvd,omitempty"`
	MsgSent int32 `json:"msgSent,omitempty"`
	UpDownTime int32 `json:"upDownTime,omitempty"`
}
