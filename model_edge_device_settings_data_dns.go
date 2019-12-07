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

type EdgeDeviceSettingsDataDns struct {
	PrimaryProvider EdgeDeviceSettingsDataDnsPrimaryProvider `json:"primaryProvider,omitempty"`
	BackupProvider EdgeDeviceSettingsDataDnsPrimaryProvider `json:"backupProvider,omitempty"`
	PrivateProviders EdgeDeviceSettingsDataDnsPrimaryProvider `json:"privateProviders,omitempty"`
}
