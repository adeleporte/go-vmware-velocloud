# \RoleApi

All URIs are relative to *http://localhost/portal/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RoleCreateRoleCustomization**](RoleApi.md#RoleCreateRoleCustomization) | **Post** /role/createRoleCustomization | Create role customization
[**RoleDeleteRoleCustomization**](RoleApi.md#RoleDeleteRoleCustomization) | **Post** /role/deleteRoleCustomization | Delete role customization
[**RoleGetUserTypeRoles**](RoleApi.md#RoleGetUserTypeRoles) | **Post** /role/getUserTypeRoles | Get roles per user type
[**RoleSetEnterpriseDelegatedToEnterpriseProxy**](RoleApi.md#RoleSetEnterpriseDelegatedToEnterpriseProxy) | **Post** /role/setEnterpriseDelegatedToEnterpriseProxy | Grant enterprise access to partner
[**RoleSetEnterpriseDelegatedToOperator**](RoleApi.md#RoleSetEnterpriseDelegatedToOperator) | **Post** /role/setEnterpriseDelegatedToOperator | Grant enterprise access to network operator
[**RoleSetEnterpriseProxyDelegatedToOperator**](RoleApi.md#RoleSetEnterpriseProxyDelegatedToOperator) | **Post** /role/setEnterpriseProxyDelegatedToOperator | Grant enterprise proxy access to network operator
[**RoleSetEnterpriseUserManagementDelegatedToOperator**](RoleApi.md#RoleSetEnterpriseUserManagementDelegatedToOperator) | **Post** /role/setEnterpriseUserManagementDelegatedToOperator | Grant enterprise user access to network operator



## RoleCreateRoleCustomization

> RoleCreateRoleCustomizationResult RoleCreateRoleCustomization(ctx, body)
Create role customization

Creates a role customization specified by `roleId` and an array of `privilegeIds`.  Privileges required:  `UPDATE` `NETWORK`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**RoleCreateRoleCustomization**](RoleCreateRoleCustomization.md)|  | 

### Return type

[**RoleCreateRoleCustomizationResult**](role_create_role_customization_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RoleDeleteRoleCustomization

> RoleDeleteRoleCustomizationResult RoleDeleteRoleCustomization(ctx, body)
Delete role customization

Deletes the specified role customization (by `name` or `forRoleId`).  Privileges required:  `UPDATE` `NETWORK`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**RoleDeleteRoleCustomization**](RoleDeleteRoleCustomization.md)|  | 

### Return type

[**RoleDeleteRoleCustomizationResult**](role_delete_role_customization_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RoleGetUserTypeRoles

> []RoleGetUserTypeRolesResultItem RoleGetUserTypeRoles(ctx, body)
Get roles per user type

Gets all roles defined for the specified `userType`.  Privileges required:  `READ` `ENTERPRISE_USER`, or  `READ` `PROXY_USER`, or  `READ` `OPERATOR_USER`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**RoleGetUserTypeRoles**](RoleGetUserTypeRoles.md)|  | 

### Return type

[**[]RoleGetUserTypeRolesResultItem**](role_get_user_type_roles_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RoleSetEnterpriseDelegatedToEnterpriseProxy

> RoleSetEnterpriseDelegatedToEnterpriseProxyResult RoleSetEnterpriseDelegatedToEnterpriseProxy(ctx, body)
Grant enterprise access to partner

Grants enterprise access to the specified enterprise proxy (partner). When an enterprise is delegated to a proxy, proxy users are granted access to view, configure, and troubleshoot Edges owned by the enterprise. As a security consideration, proxy Support users cannot view personally identifiable information.  Privileges required:  `UPDATE` `ENTERPRISE_DELEGATION`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**InlineObject3**](InlineObject3.md)|  | 

### Return type

[**RoleSetEnterpriseDelegatedToEnterpriseProxyResult**](role_set_enterprise_delegated_to_enterprise_proxy_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RoleSetEnterpriseDelegatedToOperator

> RoleSetEnterpriseDelegatedToOperatorResult RoleSetEnterpriseDelegatedToOperator(ctx, body)
Grant enterprise access to network operator

Grants enterprise access to the network operator. When an enterprise is delegated to the operator, operator users are permitted to view, configure, and troubleshoot Edges owned by the enterprise. As a security consideration, operator users cannot view personally identifiable information.  Privileges required:  `UPDATE` `ENTERPRISE_DELEGATION`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**InlineObject4**](InlineObject4.md)|  | 

### Return type

[**RoleSetEnterpriseDelegatedToOperatorResult**](role_set_enterprise_delegated_to_operator_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RoleSetEnterpriseProxyDelegatedToOperator

> RoleSetEnterpriseProxyDelegatedToOperatorResult RoleSetEnterpriseProxyDelegatedToOperator(ctx, body)
Grant enterprise proxy access to network operator

Grants enterprise proxy access to the network operator. When an enterprise proxy is delegated to the operator, operator users are granted access to view, configure and troubleshoot objects owned by the proxy.  Privileges required:  `UPDATE` `ENTERPRISE_PROXY_DELEGATION`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**InlineObject5**](InlineObject5.md)|  | 

### Return type

[**RoleSetEnterpriseProxyDelegatedToOperatorResult**](role_set_enterprise_proxy_delegated_to_operator_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RoleSetEnterpriseUserManagementDelegatedToOperator

> RoleSetEnterpriseUserManagementDelegatedToOperatorResult RoleSetEnterpriseUserManagementDelegatedToOperator(ctx, body)
Grant enterprise user access to network operator

Grants enterprise user access to the specified network operator. When enterprise user management is delegated to the operator, operator users are granted enterprise-level user management capabilities (user creation, password resets, etc.).  Privileges required:  `UPDATE` `ENTERPRISE_DELEGATION`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**InlineObject6**](InlineObject6.md)|  | 

### Return type

[**RoleSetEnterpriseUserManagementDelegatedToOperatorResult**](role_set_enterprise_user_management_delegated_to_operator_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

