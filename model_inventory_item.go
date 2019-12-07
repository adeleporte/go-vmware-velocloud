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

type InventoryItem struct {
	Id int32 `json:"id,omitempty"`
	DeviceSerialNumber string `json:"deviceSerialNumber,omitempty"`
	DeviceUuid string `json:"deviceUuid,omitempty"`
	ModelNumber string `json:"modelNumber,omitempty"`
	SiteId int32 `json:"siteId,omitempty"`
	Description string `json:"description,omitempty"`
	Acknowledged int32 `json:"acknowledged,omitempty"`
	EdgeId int32 `json:"edgeId,omitempty"`
	Edge InventoryItemEdge `json:"edge,omitempty"`
	InventoryState string `json:"inventoryState,omitempty"`
	InventoryEdgeState string `json:"inventoryEdgeState,omitempty"`
	InventoryAction string `json:"inventoryAction,omitempty"`
	VcoOwnerId int32 `json:"vcoOwnerId,omitempty"`
	VcoOwner InventoryItemVcoOwner `json:"vcoOwner,omitempty"`
	Modified time.Time `json:"modified,omitempty"`
}