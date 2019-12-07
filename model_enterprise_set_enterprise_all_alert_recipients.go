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

type EnterpriseSetEnterpriseAllAlertRecipients struct {
	EnterpriseId int32 `json:"enterpriseId,omitempty"`
	EnterpriseUsers []EnterpriseSetEnterpriseAllAlertsRecipientsEnterpriseUsers `json:"enterpriseUsers,omitempty"`
	SmsEnabled bool `json:"smsEnabled,omitempty"`
	SmsList []EnterpriseSetEnterpriseAllAlertsRecipientsSmsList `json:"smsList,omitempty"`
	EmailEnabled bool `json:"emailEnabled,omitempty"`
	EmailList []string `json:"emailList,omitempty"`
	MobileEnabled bool `json:"mobileEnabled,omitempty"`
	MobileList []string `json:"mobileList,omitempty"`
}