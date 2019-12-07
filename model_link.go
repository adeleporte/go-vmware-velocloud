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

type Link struct {
	Id int32 `json:"id"`
	Created time.Time `json:"created"`
	EdgeId int32 `json:"edgeId"`
	LogicalId string `json:"logicalId"`
	InternalId string `json:"internalId"`
	Interface string `json:"interface"`
	MacAddress string `json:"macAddress"`
	IpAddress string `json:"ipAddress"`
	Netmask string `json:"netmask"`
	NetworkSide string `json:"networkSide"`
	NetworkType string `json:"networkType"`
	DisplayName string `json:"displayName"`
	Isp string `json:"isp"`
	Org string `json:"org"`
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
	LastActive time.Time `json:"lastActive"`
	State string `json:"state"`
	BackupState string `json:"backupState"`
	VpnState string `json:"vpnState"`
	LastEvent time.Time `json:"lastEvent"`
	LastEventState string `json:"lastEventState"`
	AlertsEnabled Tinyint `json:"alertsEnabled"`
	OperatorAlertsEnabled Tinyint `json:"operatorAlertsEnabled"`
	ServiceState string `json:"serviceState"`
	Modified time.Time `json:"modified"`
	ServiceGroups []string `json:"serviceGroups,omitempty"`
}
