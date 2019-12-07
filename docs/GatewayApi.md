# \GatewayApi

All URIs are relative to *http://localhost/portal/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GatewayDeleteGateway**](GatewayApi.md#GatewayDeleteGateway) | **Post** /gateway/deleteGateway | Delete a gateway
[**GatewayGatewayProvision**](GatewayApi.md#GatewayGatewayProvision) | **Post** /gateway/gatewayProvision | Provision gateway
[**GatewayGetGatewayEdgeAssignments**](GatewayApi.md#GatewayGetGatewayEdgeAssignments) | **Post** /gateway/getGatewayEdgeAssignments | Get edge assignments for a gateway
[**GatewayUpdateGatewayAttributes**](GatewayApi.md#GatewayUpdateGatewayAttributes) | **Post** /gateway/updateGatewayAttributes | Update gateway attributes



## GatewayDeleteGateway

> GatewayDeleteGatewayResult GatewayDeleteGateway(ctx, body)
Delete a gateway

Delete a gateway by id.  Privileges required:  `DELETE` `GATEWAY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**GatewayDeleteGateway**](GatewayDeleteGateway.md)|  | 

### Return type

[**GatewayDeleteGatewayResult**](gateway_delete_gateway_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayGatewayProvision

> GatewayGatewayProvisionResult GatewayGatewayProvision(ctx, body)
Provision gateway

Provisions a gateway into the specified operator network.  Privileges required:  `CREATE` `GATEWAY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**GatewayGatewayProvision**](GatewayGatewayProvision.md)|  | 

### Return type

[**GatewayGatewayProvisionResult**](gateway_gateway_provision_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayGetGatewayEdgeAssignments

> []GatewayGetGatewayEdgeAssignmentsResultItem GatewayGetGatewayEdgeAssignments(ctx, body)
Get edge assignments for a gateway

Get edge assignments for a gateway  Privileges required:  `READ` `GATEWAY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**GatewayGetGatewayEdgeAssignments**](GatewayGetGatewayEdgeAssignments.md)|  | 

### Return type

[**[]GatewayGetGatewayEdgeAssignmentsResultItem**](gateway_get_gateway_edge_assignments_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayUpdateGatewayAttributes

> GatewayUpdateGatewayAttributesResult GatewayUpdateGatewayAttributes(ctx, body)
Update gateway attributes

Updates gateway attributes (`name`, `ipAddress`, on-premise parametrization, and `description`) and associated site attributes for the specified gateway (by `id`).  Privileges required:  `UPDATE` `GATEWAY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**GatewayUpdateGatewayAttributes**](GatewayUpdateGatewayAttributes.md)|  | 

### Return type

[**GatewayUpdateGatewayAttributesResult**](gateway_update_gateway_attributes_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

