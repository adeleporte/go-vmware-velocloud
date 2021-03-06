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

type EnterpriseServiceNetworkDataSpaces struct {
	CidrIp string `json:"cidrIp,omitempty"`
	CidrPrefix int32 `json:"cidrPrefix,omitempty"`
	MaxVlans int32 `json:"maxVlans,omitempty"`
	Mode string `json:"mode,omitempty"`
	Name string `json:"name,omitempty"`
	BranchCidrPrefix int32 `json:"branchCidrPrefix,omitempty"`
	Guest bool `json:"guest,omitempty"`
	Vlans []EnterpriseServiceNetworkDataVlans `json:"vlans,omitempty"`
}
