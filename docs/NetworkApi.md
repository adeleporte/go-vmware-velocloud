# \NetworkApi

All URIs are relative to *http://localhost/portal/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NetworkDeleteNetworkGatewayPool**](NetworkApi.md#NetworkDeleteNetworkGatewayPool) | **Post** /network/deleteNetworkGatewayPool | Delete gateway pool
[**NetworkGetNetworkConfigurations**](NetworkApi.md#NetworkGetNetworkConfigurations) | **Post** /network/getNetworkConfigurations | Get operator configuration profiles
[**NetworkGetNetworkEnterprises**](NetworkApi.md#NetworkGetNetworkEnterprises) | **Post** /network/getNetworkEnterprises | Get enterprises on network
[**NetworkGetNetworkGatewayPools**](NetworkApi.md#NetworkGetNetworkGatewayPools) | **Post** /network/getNetworkGatewayPools | Get gateway pools on network
[**NetworkGetNetworkGateways**](NetworkApi.md#NetworkGetNetworkGateways) | **Post** /network/getNetworkGateways | Get gateways on network
[**NetworkGetNetworkOperatorUsers**](NetworkApi.md#NetworkGetNetworkOperatorUsers) | **Post** /network/getNetworkOperatorUsers | Get operator users for network
[**NetworkInsertNetworkGatewayPool**](NetworkApi.md#NetworkInsertNetworkGatewayPool) | **Post** /network/insertNetworkGatewayPool | Create gateway pool
[**NetworkUpdateNetworkGatewayPoolAttributes**](NetworkApi.md#NetworkUpdateNetworkGatewayPoolAttributes) | **Post** /network/updateNetworkGatwayPoolAttributes | Update gateway pool attributes



## NetworkDeleteNetworkGatewayPool

> NetworkDeleteNetworkGatewayPoolResult NetworkDeleteNetworkGatewayPool(ctx, body)
Delete gateway pool

Deletes the specified gateway pool (by `id`).  Privileges required:  `DELETE` `GATEWAY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**NetworkDeleteNetworkGatewayPool**](NetworkDeleteNetworkGatewayPool.md)|  | 

### Return type

[**NetworkDeleteNetworkGatewayPoolResult**](network_delete_network_gateway_pool_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkGetNetworkConfigurations

> []NetworkGetNetworkConfigurationsResultItem NetworkGetNetworkConfigurations(ctx, body)
Get operator configuration profiles

Gets all operator configuration profiles associated with an operator's network. Optionally includes the modules associated with each profile. This call does not return templates.  Privileges required:  `READ` `OPERATOR_PROFILE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**NetworkGetNetworkConfigurations**](NetworkGetNetworkConfigurations.md)|  | 

### Return type

[**[]NetworkGetNetworkConfigurationsResultItem**](network_get_network_configurations_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkGetNetworkEnterprises

> []NetworkGetNetworkEnterprisesResultItem NetworkGetNetworkEnterprises(ctx, body)
Get enterprises on network

Gets all enterprises existing on a network, optionally including all Edges or Edge counts. The `edgeConfigUpdate` \"with\" option may also be passed to check whether the application of configuration updates to Edges is enabled for each enterprise.  Privileges required:  `READ` `ENTERPRISE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**NetworkGetNetworkEnterprises**](NetworkGetNetworkEnterprises.md)|  | 

### Return type

[**[]NetworkGetNetworkEnterprisesResultItem**](network_get_network_enterprises_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkGetNetworkGatewayPools

> []NetworkGetNetworkGatewayPoolsResultItem NetworkGetNetworkGatewayPools(ctx, body)
Get gateway pools on network

Gets all gateway pools associated with the specified network. Optionally includes the gateways or enterprises belonging to each pool.  Privileges required:  `READ` `GATEWAY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**NetworkGetNetworkGatewayPools**](NetworkGetNetworkGatewayPools.md)|  | 

### Return type

[**[]NetworkGetNetworkGatewayPoolsResultItem**](network_get_network_gateway_pools_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkGetNetworkGateways

> []NetworkGetNetworkGatewaysResultItem NetworkGetNetworkGateways(ctx, body)
Get gateways on network

Gets all gateways associated with the specified network.  Privileges required:  `READ` `GATEWAY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**NetworkGetNetworkGateways**](NetworkGetNetworkGateways.md)|  | 

### Return type

[**[]NetworkGetNetworkGatewaysResultItem**](network_get_network_gateways_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkGetNetworkOperatorUsers

> []NetworkGetNetworkOperatorUsersResultItem NetworkGetNetworkOperatorUsers(ctx, body)
Get operator users for network

Gets all operator users associated with the specified network.  Privileges required:  `READ` `OPERATOR_USER`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**NetworkGetNetworkOperatorUsers**](NetworkGetNetworkOperatorUsers.md)|  | 

### Return type

[**[]NetworkGetNetworkOperatorUsersResultItem**](network_get_network_operator_users_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkInsertNetworkGatewayPool

> NetworkInsertNetworkGatewayPoolResult NetworkInsertNetworkGatewayPool(ctx, body)
Create gateway pool

Creates a gateway pool for the specified network.  Privileges required:  `CREATE` `GATEWAY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**NetworkInsertNetworkGatewayPool**](NetworkInsertNetworkGatewayPool.md)|  | 

### Return type

[**NetworkInsertNetworkGatewayPoolResult**](network_insert_network_gateway_pool_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkUpdateNetworkGatewayPoolAttributes

> NetworkUpdateNetworkGatewayPoolAttributesResult NetworkUpdateNetworkGatewayPoolAttributes(ctx, body)
Update gateway pool attributes

Updates the configurable attributes (`name`, `description`, and `handOffType`) of the specified gateway pool (by `id` or `gatewayPoolId`).  Privileges required:  `UPDATE` `GATEWAY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**NetworkUpdateNetworkGatewayPoolAttributes**](NetworkUpdateNetworkGatewayPoolAttributes.md)|  | 

### Return type

[**NetworkUpdateNetworkGatewayPoolAttributesResult**](network_update_network_gateway_pool_attributes_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

