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

type DataCenterDataIaasProvider struct {
	SubscriptionObjectId int32 `json:"subscriptionObjectId,omitempty"`
	Vendor string `json:"vendor,omitempty"`
	VendorSpecificData map[string]interface{} `json:"vendorSpecificData,omitempty"`
	TunnelsEnabledOnSync bool `json:"tunnelsEnabledOnSync,omitempty"`
	SyncStatus DataCenterDataIaasProviderSyncStatus `json:"syncStatus,omitempty"`
}
