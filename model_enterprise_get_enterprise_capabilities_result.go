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

type EnterpriseGetEnterpriseCapabilitiesResult struct {
	EdgeVnfsEnable bool `json:"edgeVnfs.enable,omitempty"`
	EdgeVnfsSecurityVnfPaloAlto bool `json:"edgeVnfs.securityVnf.paloAlto,omitempty"`
	EdgeVnfsSecurityVnfCheckpoint bool `json:"edgeVnfs.securityVnf.checkpoint,omitempty"`
	EdgeVnfsSecurityVnfFortinet bool `json:"edgeVnfs.securityVnf.fortinet,omitempty"`
	EnableBGP bool `json:"enableBGP,omitempty"`
	EnableCosMapping bool `json:"enableCosMapping,omitempty"`
	EnableEnterpriseAuth bool `json:"enableEnterpriseAuth,omitempty"`
	EnableFwLogs bool `json:"enableFwLogs,omitempty"`
	EnableStatefulFirewall bool `json:"enableStatefulFirewall,omitempty"`
	EnableNetworks bool `json:"enableNetworks,omitempty"`
	EnableOSPF bool `json:"enableOSPF,omitempty"`
	EnablePKI bool `json:"enablePKI,omitempty"`
	EnablePremium bool `json:"enablePremium,omitempty"`
	EnableSegmentation bool `json:"enableSegmentation,omitempty"`
	EnableServiceRateLimiting bool `json:"enableServiceRateLimiting,omitempty"`
	EnableVQM bool `json:"enableVQM,omitempty"`
}
