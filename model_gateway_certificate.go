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

type GatewayCertificate struct {
	Id int32 `json:"id,omitempty"`
	Created time.Time `json:"created,omitempty"`
	CsrId int32 `json:"csrId,omitempty"`
	GatewayId int32 `json:"gatewayId,omitempty"`
	NetworkId int32 `json:"networkId,omitempty"`
	Certificate string `json:"certificate,omitempty"`
	SerialNumber string `json:"serialNumber,omitempty"`
	SubjectKeyId string `json:"subjectKeyId,omitempty"`
	FingerPrint string `json:"fingerPrint,omitempty"`
	ValidFrom time.Time `json:"validFrom,omitempty"`
	ValidTo time.Time `json:"validTo,omitempty"`
}