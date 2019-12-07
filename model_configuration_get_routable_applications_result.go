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

type ConfigurationGetRoutableApplicationsResult struct {
	// Maps application class IDs (strings) to application IDs (integers)
	//ApplicationClassToApplication ModelMap `json:"applicationClassToApplication"`
	// Maps application IDs (strings) to class IDs (integers)
	//ApplicationToApplicationClass ModelMap `json:"applicationToApplicationClass"`
	// Map of application class IDs to class names
	//ApplicationClasses ModelMap `json:"applicationClasses"`
	Applications []Application `json:"applications"`
	MetaData ApplicationMetadata `json:"metaData"`
}