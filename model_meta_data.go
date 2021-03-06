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

type MetaData struct {
	Created time.Time `json:"created,omitempty"`
	Effective time.Time `json:"effective,omitempty"`
	Id int32 `json:"id,omitempty"`
	IsSmallData Tinyint `json:"isSmallData,omitempty"`
	Modified time.Time `json:"modified,omitempty"`
	Name string `json:"name"`
	Type string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
	ConfigurationId int32 `json:"configurationId,omitempty"`
	Data MetadataData `json:"data"`
	Refs map[string]interface{} `json:"refs,omitempty"`
	SchemaVersion string `json:"schemaVersion,omitempty"`
}
