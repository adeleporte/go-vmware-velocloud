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

type EnterpriseGetEnterpriseEdgesResultItemAllOf struct {
	Certificates []EdgeCertificate `json:"certificates,omitempty"`
	Configuration EnterpriseGetEnterpriseEdgesResultItemAllOfConfiguration `json:"configuration,omitempty"`
	Ha EnterpriseGetEnterpriseEdgesResultItemAllOfHa `json:"ha,omitempty"`
	Licenses []EdgeLicense `json:"licenses,omitempty"`
	Links []Link `json:"links,omitempty"`
	RecentLinks []Link `json:"recentLinks,omitempty"`
	Site Site `json:"site,omitempty"`
	IsHub bool `json:"isHub,omitempty"`
	IsSoftwareVersionSupportedByVco bool `json:"isSoftwareVersionSupportedByVco,omitempty"`
}
