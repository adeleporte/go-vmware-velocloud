# \EnterpriseProxyApi

All URIs are relative to *http://localhost/portal/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnterpriseProxyGetEnterpriseProxyEdgeInventory**](EnterpriseProxyApi.md#EnterpriseProxyGetEnterpriseProxyEdgeInventory) | **Post** /enterpriseProxy/getEnterpriseProxyEdgeInventory | Get partner enterprises and associated Edge inventory
[**EnterpriseProxyGetEnterpriseProxyEnterprises**](EnterpriseProxyApi.md#EnterpriseProxyGetEnterpriseProxyEnterprises) | **Post** /enterpriseProxy/getEnterpriseProxyEnterprises | Get partner enterprises
[**EnterpriseProxyGetEnterpriseProxyGatewayPools**](EnterpriseProxyApi.md#EnterpriseProxyGetEnterpriseProxyGatewayPools) | **Post** /enterpriseProxy/getEnterpriseProxyGatewayPools | Get gateway pools
[**EnterpriseProxyGetEnterpriseProxyGateways**](EnterpriseProxyApi.md#EnterpriseProxyGetEnterpriseProxyGateways) | **Post** /enterpriseProxy/getEnterpriseProxyGateways | Get gateways for enterprise proxy
[**EnterpriseProxyGetEnterpriseProxyOperatorProfiles**](EnterpriseProxyApi.md#EnterpriseProxyGetEnterpriseProxyOperatorProfiles) | **Post** /enterpriseProxy/getEnterpriseProxyOperatorProfiles | Get partner operator profiles
[**EnterpriseProxyInsertEnterpriseProxyEnterprise**](EnterpriseProxyApi.md#EnterpriseProxyInsertEnterpriseProxyEnterprise) | **Post** /enterpriseProxy/insertEnterpriseProxyEnterprise | Create enterprise owned by MSP
[**EnterpriseProxyInsertOrUpdateEnterpriseProxyProperty**](EnterpriseProxyApi.md#EnterpriseProxyInsertOrUpdateEnterpriseProxyProperty) | **Post** /enterpriseProxy/insertOrUpdateEnterpriseProxyProperty | Insert or update an enterprise proxy (MSP) property



## EnterpriseProxyGetEnterpriseProxyEdgeInventory

> []EnterpriseProxyGetEnterpriseProxyEdgeInventoryResultItem EnterpriseProxyGetEnterpriseProxyEdgeInventory(ctx, body)
Get partner enterprises and associated Edge inventory

Gets all partner enterprises and their associated Edge inventory.  Privileges required:  `READ` `ENTERPRISE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseProxyGetEnterpriseProxyEdgeInventory**](EnterpriseProxyGetEnterpriseProxyEdgeInventory.md)|  | 

### Return type

[**[]EnterpriseProxyGetEnterpriseProxyEdgeInventoryResultItem**](enterprise_proxy_get_enterprise_proxy_edge_inventory_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseProxyGetEnterpriseProxyEnterprises

> []EnterpriseProxyGetEnterpriseProxyEnterprisesResultItem EnterpriseProxyGetEnterpriseProxyEnterprises(ctx, body)
Get partner enterprises

Gets all partner enterprises. Optionally includes all Edges or Edge counts.  Privileges required:  `READ` `ENTERPRISE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseProxyGetEnterpriseProxyEnterprises**](EnterpriseProxyGetEnterpriseProxyEnterprises.md)|  | 

### Return type

[**[]EnterpriseProxyGetEnterpriseProxyEnterprisesResultItem**](enterprise_proxy_get_enterprise_proxy_enterprises_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseProxyGetEnterpriseProxyGatewayPools

> []EnterpriseProxyGetEnterpriseProxyGatewayPoolsResultItem EnterpriseProxyGetEnterpriseProxyGatewayPools(ctx, body)
Get gateway pools

Gets all gateway pools associated with the specified enterprise proxy. Optionally includes all gateways or enterprises belonging to each pool.  Privileges required:  `READ` `GATEWAY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseProxyGetEnterpriseProxyGatewayPools**](EnterpriseProxyGetEnterpriseProxyGatewayPools.md)|  | 

### Return type

[**[]EnterpriseProxyGetEnterpriseProxyGatewayPoolsResultItem**](enterprise_proxy_get_enterprise_proxy_gateway_pools_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseProxyGetEnterpriseProxyGateways

> []EnterpriseProxyGetEnterpriseProxyGatewaysResultItem EnterpriseProxyGetEnterpriseProxyGateways(ctx, body)
Get gateways for enterprise proxy

Gets all gateways associated with the specified enterprise proxy.  Privileges required:  `READ` `GATEWAY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseProxyGetEnterpriseProxyGateways**](EnterpriseProxyGetEnterpriseProxyGateways.md)|  | 

### Return type

[**[]EnterpriseProxyGetEnterpriseProxyGatewaysResultItem**](enterprise_proxy_get_enterprise_proxy_gateways_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseProxyGetEnterpriseProxyOperatorProfiles

> []EnterpriseProxyGetEnterpriseProxyOperatorProfilesResultItem EnterpriseProxyGetEnterpriseProxyOperatorProfiles(ctx, body)
Get partner operator profiles

Gets all operator profiles associated with the specified partner (MSP), as assigned by the operator.  Privileges required:  `READ` `OPERATOR_PROFILE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseProxyGetEnterpriseProxyOperatorProfiles**](EnterpriseProxyGetEnterpriseProxyOperatorProfiles.md)|  | 

### Return type

[**[]EnterpriseProxyGetEnterpriseProxyOperatorProfilesResultItem**](enterprise_proxy_get_enterprise_proxy_operator_profiles_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseProxyInsertEnterpriseProxyEnterprise

> EnterpriseProxyInsertEnterpriseProxyEnterpriseResult EnterpriseProxyInsertEnterpriseProxyEnterprise(ctx, body)
Create enterprise owned by MSP

Creates a new enterprise owned by an MSP. Whereas the `insertEnterprise` method creates an enterprise in the global or network context with no MSP association, this method creates one owned by an MSP, as determined by the credentials of the caller.  Privileges required:  `CREATE` `ENTERPRISE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseProxyInsertEnterpriseProxyEnterprise**](EnterpriseProxyInsertEnterpriseProxyEnterprise.md)|  | 

### Return type

[**EnterpriseProxyInsertEnterpriseProxyEnterpriseResult**](enterprise_proxy_insert_enterprise_proxy_enterprise_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseProxyInsertOrUpdateEnterpriseProxyProperty

> EnterpriseProxyInsertOrUpdateEnterpriseProxyPropertyResult EnterpriseProxyInsertOrUpdateEnterpriseProxyProperty(ctx, body)
Insert or update an enterprise proxy (MSP) property

Insert a enterprise proxy property. If property with the given name already exists, the property will be updated.  Privileges required:  `UPDATE` `PROXY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseProxyInsertOrUpdateEnterpriseProxyProperty**](EnterpriseProxyInsertOrUpdateEnterpriseProxyProperty.md)|  | 

### Return type

[**EnterpriseProxyInsertOrUpdateEnterpriseProxyPropertyResult**](enterprise_proxy_insert_or_update_enterprise_proxy_property_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

