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

type OpenidWellKnownConfig struct {
	Issuer string `json:"issuer,omitempty"`
	AuthorizationEndpoint string `json:"authorization_endpoint,omitempty"`
	TokenEndpoint string `json:"token_endpoint,omitempty"`
	RevocationEndpoint string `json:"revocation_endpoint,omitempty"`
	UserinfoEndpoint string `json:"userinfo_endpoint,omitempty"`
	JwksUri string `json:"jwks_uri,omitempty"`
	ScopesSupported []string `json:"scopes_supported,omitempty"`
	ResponseTypesSupported []string `json:"response_types_supported,omitempty"`
	ResponseModesSupported []string `json:"response_modes_supported,omitempty"`
	ClaimsSupported []string `json:"claims_supported,omitempty"`
}
