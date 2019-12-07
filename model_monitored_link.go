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

type MonitoredLink struct {
	DisplayName string `json:"displayName"`
	EdgeHASerialNumber string `json:"edgeHASerialNumber"`
	EdgeId int32 `json:"edgeId"`
	EdgeLastContact time.Time `json:"edgeLastContact"`
	EdgeLatitude float64 `json:"edgeLatitude"`
	EdgeLongitude float64 `json:"edgeLongitude"`
	EdgeModelNumber string `json:"edgeModelNumber"`
	EdgeName string `json:"edgeName"`
	EdgeSerialNumber string `json:"edgeSerialNumber"`
	EdgeServiceUpSince time.Time `json:"edgeServiceUpSince"`
	EdgeState string `json:"edgeState"`
	EdgeSystemUpSince time.Time `json:"edgeSystemUpSince"`
	EnterpriseId int32 `json:"enterpriseId"`
	EnterpriseName string `json:"enterpriseName"`
	EnterpriseProxyId int32 `json:"enterpriseProxyId"`
	EnterpriseProxyName string `json:"enterpriseProxyName"`
	Interface string `json:"interface"`
	InternalId string `json:"internalId"`
	Isp string `json:"isp"`
	LinkId int32 `json:"linkId"`
	LinkIpAddress string `json:"linkIpAddress"`
	LinkLastActive time.Time `json:"linkLastActive"`
	LinkState string `json:"linkState"`
	LinkVpnState string `json:"linkVpnState"`
}
