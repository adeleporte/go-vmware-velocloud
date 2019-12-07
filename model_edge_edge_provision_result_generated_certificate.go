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

type EdgeEdgeProvisionResultGeneratedCertificate struct {
	Certificate string `json:"certificate,omitempty"`
	CaCertificate string `json:"ca-certificate,omitempty"`
	PrivateKey string `json:"privateKey,omitempty"`
	PrivateKeyPassword string `json:"privateKeyPassword,omitempty"`
	Csr string `json:"csr,omitempty"`
}