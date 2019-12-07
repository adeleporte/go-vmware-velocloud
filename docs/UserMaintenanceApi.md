# \UserMaintenanceApi

All URIs are relative to *http://localhost/portal/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnterpriseGetEnterpriseUsers**](UserMaintenanceApi.md#EnterpriseGetEnterpriseUsers) | **Post** /enterprise/getEnterpriseUsers | Get enterprise users
[**EnterpriseInsertEnterpriseUser**](UserMaintenanceApi.md#EnterpriseInsertEnterpriseUser) | **Post** /enterprise/insertEnterpriseUser | Create enterprise user
[**EnterpriseProxyDeleteEnterpriseProxyUser**](UserMaintenanceApi.md#EnterpriseProxyDeleteEnterpriseProxyUser) | **Post** /enterpriseProxy/deleteEnterpriseProxyUser | Delete partner admin user
[**EnterpriseProxyGetEnterpriseProxyUser**](UserMaintenanceApi.md#EnterpriseProxyGetEnterpriseProxyUser) | **Post** /enterpriseProxy/getEnterpriseProxyUser | Get enterprise proxy user
[**EnterpriseProxyGetEnterpriseProxyUsers**](UserMaintenanceApi.md#EnterpriseProxyGetEnterpriseProxyUsers) | **Post** /enterpriseProxy/getEnterpriseProxyUsers | Get enterprise proxy admin users
[**EnterpriseProxyInsertEnterpriseProxyUser**](UserMaintenanceApi.md#EnterpriseProxyInsertEnterpriseProxyUser) | **Post** /enterpriseProxy/insertEnterpriseProxyUser | Create partner admin user
[**EnterpriseProxyUpdateEnterpriseProxyUser**](UserMaintenanceApi.md#EnterpriseProxyUpdateEnterpriseProxyUser) | **Post** /enterpriseProxy/updateEnterpriseProxyUser | Update enterprise proxy admin user
[**EnterpriseUserDeleteEnterpriseUser**](UserMaintenanceApi.md#EnterpriseUserDeleteEnterpriseUser) | **Post** /enterpriseUser/deleteEnterpriseUser | Delete enterprise user
[**EnterpriseUserGetEnterpriseUser**](UserMaintenanceApi.md#EnterpriseUserGetEnterpriseUser) | **Post** /enterpriseUser/getEnterpriseUser | Get enterprise user
[**EnterpriseUserUpdateEnterpriseUser**](UserMaintenanceApi.md#EnterpriseUserUpdateEnterpriseUser) | **Post** /enterpriseUser/updateEnterpriseUser | Update enterprise user
[**OperatorUserDeleteOperatorUser**](UserMaintenanceApi.md#OperatorUserDeleteOperatorUser) | **Post** /operatorUser/deleteOperatorUser | Delete operator user
[**OperatorUserGetOperatorUser**](UserMaintenanceApi.md#OperatorUserGetOperatorUser) | **Post** /operatorUser/getOperatorUser | Get operator user
[**OperatorUserInsertOperatorUser**](UserMaintenanceApi.md#OperatorUserInsertOperatorUser) | **Post** /operatorUser/insertOperatorUser | Create operator user
[**OperatorUserUpdateOperatorUser**](UserMaintenanceApi.md#OperatorUserUpdateOperatorUser) | **Post** /operatorUser/updateOperatorUser | Update operator user



## EnterpriseGetEnterpriseUsers

> []EnterpriseGetEnterpriseUsersResultItem EnterpriseGetEnterpriseUsers(ctx, body)
Get enterprise users

Gets all enterprise users for the specified enterprise (by `enterpriseId`).  Privileges required:  `READ` `ENTERPRISE`  `READ` `ENTERPRISE_USER`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseGetEnterpriseUsers**](EnterpriseGetEnterpriseUsers.md)|  | 

### Return type

[**[]EnterpriseGetEnterpriseUsersResultItem**](enterprise_get_enterprise_users_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseInsertEnterpriseUser

> EnterpriseInsertEnterpriseUserResult EnterpriseInsertEnterpriseUser(ctx, body)
Create enterprise user

Creates a new enterprise user.  Privileges required:  `CREATE` `ENTERPRISE_USER`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**NewEnterpriseUser**](NewEnterpriseUser.md)|  | 

### Return type

[**EnterpriseInsertEnterpriseUserResult**](enterprise_insert_enterprise_user_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseProxyDeleteEnterpriseProxyUser

> EnterpriseProxyDeleteEnterpriseProxyUserResult EnterpriseProxyDeleteEnterpriseProxyUser(ctx, body)
Delete partner admin user

Deletes the specified enterprise proxy user (by `id` or `username`). Note: `enterpriseProxyId` is a required parameter when invoking this method as an operator or as a partner user.  Privileges required:  `DELETE` `PROXY_USER`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseProxyDeleteEnterpriseProxyUser**](EnterpriseProxyDeleteEnterpriseProxyUser.md)|  | 

### Return type

[**EnterpriseProxyDeleteEnterpriseProxyUserResult**](enterprise_proxy_delete_enterprise_proxy_user_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseProxyGetEnterpriseProxyUser

> EnterpriseProxyGetEnterpriseProxyUserResult EnterpriseProxyGetEnterpriseProxyUser(ctx, body)
Get enterprise proxy user

Gets the specified enterprise proxy user (by `id` or `username`).  Privileges required:  `READ` `PROXY_USER`  `READ` `PROXY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseProxyGetEnterpriseProxyUser**](EnterpriseProxyGetEnterpriseProxyUser.md)|  | 

### Return type

[**EnterpriseProxyGetEnterpriseProxyUserResult**](enterprise_proxy_get_enterprise_proxy_user_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseProxyGetEnterpriseProxyUsers

> []EnterpriseProxyGetEnterpriseProxyUsersResultItem EnterpriseProxyGetEnterpriseProxyUsers(ctx, body)
Get enterprise proxy admin users

Gets all proxy admin users for the specified enterprise.  Privileges required:  `READ` `ENTERPRISE`  `READ` `PROXY_USER`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseProxyGetEnterpriseProxyUsers**](EnterpriseProxyGetEnterpriseProxyUsers.md)|  | 

### Return type

[**[]EnterpriseProxyGetEnterpriseProxyUsersResultItem**](enterprise_proxy_get_enterprise_proxy_users_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseProxyInsertEnterpriseProxyUser

> EnterpriseProxyInsertEnterpriseProxyUserResult EnterpriseProxyInsertEnterpriseProxyUser(ctx, body)
Create partner admin user

Creates a new partner admin user.  Privileges required:  `CREATE` `PROXY_USER`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**NewEnterpriseProxyUser**](NewEnterpriseProxyUser.md)|  | 

### Return type

[**EnterpriseProxyInsertEnterpriseProxyUserResult**](enterprise_proxy_insert_enterprise_proxy_user_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseProxyUpdateEnterpriseProxyUser

> EnterpriseProxyUpdateEnterpriseProxyUserResult EnterpriseProxyUpdateEnterpriseProxyUser(ctx, body)
Update enterprise proxy admin user

Updates the specified enterprise proxy admin user (by object `id` or other identifying attribute). Expects an `_update` object containing the field(s) to be updated and their corresponding value(s).  Privileges required:  `UPDATE` `PROXY_USER`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseProxyUpdateEnterpriseProxyUser**](EnterpriseProxyUpdateEnterpriseProxyUser.md)|  | 

### Return type

[**EnterpriseProxyUpdateEnterpriseProxyUserResult**](enterprise_proxy_update_enterprise_proxy_user_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseUserDeleteEnterpriseUser

> EnterpriseUserDeleteEnterpriseUserResult EnterpriseUserDeleteEnterpriseUser(ctx, body)
Delete enterprise user

Deletes the specified enterprise user (by `id` or `username`). Note: `enterpriseId` is a required parameter when invoking this method as an operator or as a partner user.  Privileges required:  `DELETE` `ENTERPRISE_USER`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseUserDeleteEnterpriseUser**](EnterpriseUserDeleteEnterpriseUser.md)|  | 

### Return type

[**EnterpriseUserDeleteEnterpriseUserResult**](enterprise_user_delete_enterprise_user_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseUserGetEnterpriseUser

> EnterpriseUserGetEnterpriseUserResult EnterpriseUserGetEnterpriseUser(ctx, body)
Get enterprise user

Gets the specified enterprise user (by `id` or `username`).  Privileges required:  `READ` `ENTERPRISE_USER`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseUserGetEnterpriseUser**](EnterpriseUserGetEnterpriseUser.md)|  | 

### Return type

[**EnterpriseUserGetEnterpriseUserResult**](enterprise_user_get_enterprise_user_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseUserUpdateEnterpriseUser

> EnterpriseUserUpdateEnterpriseUserResult EnterpriseUserUpdateEnterpriseUser(ctx, body)
Update enterprise user

Updates the specified enterprise user (by object `id` or other identifying attribute). Expects an `_update` object containing the field(s) to be updated and their corresponding value(s).  Privileges required:  `UPDATE` `ENTERPRISE_USER`, or  `UPDATE` `OPERATOR_USER`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseUserUpdateEnterpriseUser1**](EnterpriseUserUpdateEnterpriseUser1.md)|  | 

### Return type

[**EnterpriseUserUpdateEnterpriseUserResult**](enterprise_user_update_enterprise_user_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperatorUserDeleteOperatorUser

> OperatorUserDeleteOperatorUserResult OperatorUserDeleteOperatorUser(ctx, body)
Delete operator user

Deletes the specified operator user (by `id` or `username`).  Privileges required:  `DELETE` `OPERATOR_USER`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**OperatorUserDeleteOperatorUser**](OperatorUserDeleteOperatorUser.md)|  | 

### Return type

[**OperatorUserDeleteOperatorUserResult**](operator_user_delete_operator_user_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperatorUserGetOperatorUser

> OperatorUserGetOperatorUserResult OperatorUserGetOperatorUser(ctx, body)
Get operator user

Gets the specified operator user object (by `id` or `username`).  Privileges required:  `READ` `OPERATOR_USER`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**OperatorUserGetOperatorUser**](OperatorUserGetOperatorUser.md)|  | 

### Return type

[**OperatorUserGetOperatorUserResult**](operator_user_get_operator_user_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperatorUserInsertOperatorUser

> OperatorUserGetOperatorUserResult OperatorUserInsertOperatorUser(ctx, body)
Create operator user

Creates an operator user and associates it with the operator's network.  Privileges required:  `CREATE` `OPERATOR_USER`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**OperatorUserInsertOperatorUser**](OperatorUserInsertOperatorUser.md)|  | 

### Return type

[**OperatorUserGetOperatorUserResult**](operator_user_get_operator_user_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperatorUserUpdateOperatorUser

> OperatorUserUpdateOperatorUserResult OperatorUserUpdateOperatorUser(ctx, body)
Update operator user

Updates the specified operator user (by object `id` or `username`). Expects an `_update` object containing the field(s) to be updated and their corresponding value(s).  Privileges required:  `UPDATE` `OPERATOR_USER`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**OperatorUserUpdateOperatorUser1**](OperatorUserUpdateOperatorUser1.md)|  | 

### Return type

[**OperatorUserUpdateOperatorUserResult**](operator_user_update_operator_user_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

