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

type OperatorUserUpdateOperatorUser struct {
	FirstName string `json:"firstName,omitempty"`
	LastName string `json:"lastName,omitempty"`
	Password string `json:"password,omitempty"`
	Password2 string `json:"password2,omitempty"`
	RoleId int32 `json:"roleId,omitempty"`
	Email string `json:"email,omitempty"`
	MobilePhone string `json:"mobilePhone,omitempty"`
	OfficePhone string `json:"officePhone,omitempty"`
	OperatorId int32 `json:"operatorId,omitempty"`
	UserType string `json:"userType,omitempty"`
	Username string `json:"username,omitempty"`
	Domain string `json:"domain,omitempty"`
	IsNative bool `json:"isNative,omitempty"`
	IsLocked bool `json:"isLocked,omitempty"`
	IsActive bool `json:"isActive,omitempty"`
}
