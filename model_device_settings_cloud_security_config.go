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

type DeviceSettingsCloudSecurityConfig struct {
	TunnelingProtocol       string                                   `json:"tunnelingProtocol,omitempty"`
	AuthenticationAlgorithm string                                   `json:"authenticationAlgorithm,omitempty"`
	EncryptionAlgorithm     string                                   `json:"encryptionAlgorithm,omitempty"`
	Redirect                string                                   `json:"redirect,omitempty"`
	IKEPROP                 DeviceSettingsCloudSecurityConfigIkeprop `json:"IKEPROP,omitempty"`
	GREPROP                 DeviceSettingsCloudSecurityConfigGreprop `json:"GREPROP,omitempty"`
}
