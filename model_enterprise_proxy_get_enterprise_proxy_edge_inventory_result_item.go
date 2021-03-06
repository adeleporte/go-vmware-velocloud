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

type EnterpriseProxyGetEnterpriseProxyEdgeInventoryResultItem struct {
	EnterpriseName string `json:"enterpriseName"`
	EnterpriseId int32 `json:"enterpriseId"`
	EdgeName string `json:"edgeName"`
	EdgeId int32 `json:"edgeId"`
	Created time.Time `json:"created"`
	EdgeState string `json:"edgeState"`
	SerialNumber string `json:"serialNumber"`
	HaSerialNumber string `json:"haSerialNumber"`
	ActivationState string `json:"activationState"`
	ActivationTime time.Time `json:"activationTime"`
	ModelNumber string `json:"modelNumber"`
	SoftwareVersion string `json:"softwareVersion"`
	SoftwareUpdated time.Time `json:"softwareUpdated"`
	LastContact time.Time `json:"lastContact"`
}
