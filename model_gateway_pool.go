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

type GatewayPool struct {
	Id int32 `json:"id,omitempty"`
	NetworkId int32 `json:"networkId,omitempty"`
	EnterpriseProxyId int32 `json:"enterpriseProxyId,omitempty"`
	Created time.Time `json:"created,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	IsDefault bool `json:"isDefault,omitempty"`
	HandOffType string `json:"handOffType,omitempty"`
	Modified time.Time `json:"modified,omitempty"`
}
