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

type VcoDiagnosticBundle struct {
	Id int32 `json:"id,omitempty"`
	RequestId int64 `json:"requestId,omitempty"`
	UserId string `json:"userId,omitempty"`
	AgeOutTime time.Time `json:"ageOutTime,omitempty"`
	Reason string `json:"reason,omitempty"`
	Created time.Time `json:"created,omitempty"`
	BlobId int32 `json:"blobId,omitempty"`
	FileName string `json:"fileName,omitempty"`
	FileStore string `json:"fileStore,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
	JobId string `json:"jobId,omitempty"`
	Status string `json:"status,omitempty"`
}
