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

type EnterpriseAlertNotificationUserData struct {
	Email string `json:"email"`
	EmailEnabled int32 `json:"emailEnabled"`
	Enabled int32 `json:"enabled"`
	EnterpriseUserId int32 `json:"enterpriseUserId"`
	MobileEnabled int32 `json:"mobileEnabled"`
	MobilePhone string `json:"mobilePhone"`
	SmsEnabled int32 `json:"smsEnabled"`
	Username string `json:"username"`
}
