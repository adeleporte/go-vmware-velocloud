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

type EdgeListHaDetailsActiveStandby struct {
	HaLastContact time.Time `json:"haLastContact"`
	HaPreviousState string `json:"haPreviousState"`
	HaSerialNumber string `json:"haSerialNumber"`
	HaState string `json:"haState"`
}