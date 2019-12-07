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

type EnterpriseGetEnterpriseGatewayHandoffResult struct {
	EnterpriseId int32 `json:"enterpriseId,omitempty"`
	Value GatewayHandoffValue `json:"value,omitempty"`
	Id int32 `json:"id,omitempty"`
	Created time.Time `json:"created,omitempty"`
	Name string `json:"name,omitempty"`
	IsPassword bool `json:"isPassword,omitempty"`
	DataType string `json:"dataType,omitempty"`
	Description string `json:"description,omitempty"`
	Modified time.Time `json:"modified,omitempty"`
}
