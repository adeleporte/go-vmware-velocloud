# \MonitoringApi

All URIs are relative to *http://localhost/portal/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MonitoringGetAggregateEdgeLinkMetrics**](MonitoringApi.md#MonitoringGetAggregateEdgeLinkMetrics) | **Post** /monitoring/getAggregateEdgeLinkMetrics | Get aggregate Edge link metrics across enterprises
[**MonitoringGetAggregateEnterpriseEvents**](MonitoringApi.md#MonitoringGetAggregateEnterpriseEvents) | **Post** /monitoring/getAggregateEnterpriseEvents | Get events across all enterprises
[**MonitoringGetAggregates**](MonitoringApi.md#MonitoringGetAggregates) | **Post** /monitoring/getAggregates | Get aggregate enterprise and Edge information
[**MonitoringGetEnterpriseBgpPeerStatus**](MonitoringApi.md#MonitoringGetEnterpriseBgpPeerStatus) | **Post** /monitoring/getEnterpriseBgpPeerStatus | Get gateway BGP peer status for all enterprise gateways
[**MonitoringGetEnterpriseEdgeBgpPeerStatus**](MonitoringApi.md#MonitoringGetEnterpriseEdgeBgpPeerStatus) | **Post** /monitoring/getEnterpriseEdgeBgpPeerStatus | Get Edge BGP peer status for all enterprise Edges
[**MonitoringGetEnterpriseEdgeClusterStatus**](MonitoringApi.md#MonitoringGetEnterpriseEdgeClusterStatus) | **Post** /monitoring/getEnterpriseEdgeClusterStatus | Get Edge Cluster status
[**MonitoringGetEnterpriseEdgeLinkStatus**](MonitoringApi.md#MonitoringGetEnterpriseEdgeLinkStatus) | **Post** /monitoring/getEnterpriseEdgeLinkStatus | Get Edge and link status data
[**MonitoringGetEnterpriseEdgeStatus**](MonitoringApi.md#MonitoringGetEnterpriseEdgeStatus) | **Post** /monitoring/getEnterpriseEdgeStatus | Get Enterprise Edge monitoring status
[**MonitoringGetNetworkGatewayStatus**](MonitoringApi.md#MonitoringGetNetworkGatewayStatus) | **Post** /monitoring/getNetworkGatewayStatus | Get Network Gateway monitoring status



## MonitoringGetAggregateEdgeLinkMetrics

> []MonitoringGetAggregateEdgeLinkMetricsResultItem MonitoringGetAggregateEdgeLinkMetrics(ctx, body)
Get aggregate Edge link metrics across enterprises

Gets aggregate link metrics for the request interval for all active links across all enterprises, where a link is considered to be active if an Edge has reported any activity for it in the last 24 hours. On success, returns an array of network utilization metrics, one per link.  Privileges required:  `READ` `ENTERPRISE`  `READ` `EDGE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MonitoringGetAggregateEdgeLinkMetrics**](MonitoringGetAggregateEdgeLinkMetrics.md)|  | 

### Return type

[**[]MonitoringGetAggregateEdgeLinkMetricsResultItem**](monitoring_get_aggregate_edge_link_metrics_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitoringGetAggregateEnterpriseEvents

> MonitoringGetAggregateEnterpriseEventsResult MonitoringGetAggregateEnterpriseEvents(ctx, body)
Get events across all enterprises

Gets events across all enterprises in a paginated list. When called in the MSP/Partner context, queries only enterprises managed by the MSP.  Privileges required:  `READ` `ENTERPRISE`  `READ` `EDGE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MonitoringGetAggregateEnterpriseEvents**](MonitoringGetAggregateEnterpriseEvents.md)|  | 

### Return type

[**MonitoringGetAggregateEnterpriseEventsResult**](monitoring_get_aggregate_enterprise_events_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitoringGetAggregates

> MonitoringGetAggregatesResult MonitoringGetAggregates(ctx, body)
Get aggregate enterprise and Edge information

Gets a comprehensive listing of all enterprises and Edges on a network. Returns an object containing an aggregate `edgeCount`, a list (`enterprises`) containing enterprise objects, and a map (`edges`) that provides Edge counts per enterprise.  Privileges required:  `READ` `ENTERPRISE`  `READ` `EDGE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MonitoringGetAggregates**](MonitoringGetAggregates.md)|  | 

### Return type

[**MonitoringGetAggregatesResult**](monitoring_get_aggregates_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitoringGetEnterpriseBgpPeerStatus

> []MonitoringGetEnterpriseBgpPeerStatusResultItem MonitoringGetEnterpriseBgpPeerStatus(ctx, body)
Get gateway BGP peer status for all enterprise gateways

Gets the gateway BGP peer status for all enterprise gateways. Returns an array in which each entry corresponds to a gateway and contains an associated set of BGP peers with state records.  Privileges required:  `READ` `NETWORK_SERVICE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**InlineObject**](InlineObject.md)|  | 

### Return type

[**[]MonitoringGetEnterpriseBgpPeerStatusResultItem**](monitoring_get_enterprise_bgp_peer_status_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitoringGetEnterpriseEdgeBgpPeerStatus

> []MonitoringGetEnterpriseEdgeBgpPeerStatusResultItem MonitoringGetEnterpriseEdgeBgpPeerStatus(ctx, body)
Get Edge BGP peer status for all enterprise Edges

Gets the Edge BGP peer status for all enterprise Edges. Returns an array in which each entry corresponds to an Edge and contains an associated set of BGP peers and state records.  Privileges required:  `READ` `EDGE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**InlineObject1**](InlineObject1.md)|  | 

### Return type

[**[]MonitoringGetEnterpriseEdgeBgpPeerStatusResultItem**](monitoring_get_enterprise_edge_bgp_peer_status_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitoringGetEnterpriseEdgeClusterStatus

> []InlineResponse2001 MonitoringGetEnterpriseEdgeClusterStatus(ctx, body)
Get Edge Cluster status

Get edge`s utilization data by clusters  Privileges required:  `READ` `EDGE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MonitoringGetEnterpriseEdgeClusterStatus**](MonitoringGetEnterpriseEdgeClusterStatus.md)|  | 

### Return type

[**[]InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitoringGetEnterpriseEdgeLinkStatus

> []MonitoringGetEnterpriseEdgeLinkStatusResultItem MonitoringGetEnterpriseEdgeLinkStatus(ctx, body)
Get Edge and link status data

Gets the current Edge and Edge-link status for all enterprises. edges.  Privileges required:  `READ` `ENTERPRISE`  `READ` `EDGE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**InlineObject2**](InlineObject2.md)|  | 

### Return type

[**[]MonitoringGetEnterpriseEdgeLinkStatusResultItem**](monitoring_get_enterprise_edge_link_status_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitoringGetEnterpriseEdgeStatus

> MonitoringGetEnterpriseEdgeStatusResult MonitoringGetEnterpriseEdgeStatus(ctx, body)
Get Enterprise Edge monitoring status

Fetch Enterprise Edge monitoring status time series for the given time interval, all for the enterprise or list of edgess, and list of metrics. On success, this method returns anarray of edge monitoring statuses.  Privileges required:  `READ` `EDGE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MonitoringGetEnterpriseEdgeStatus**](MonitoringGetEnterpriseEdgeStatus.md)|  | 

### Return type

[**MonitoringGetEnterpriseEdgeStatusResult**](monitoring_get_enterprise_edge_status_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitoringGetNetworkGatewayStatus

> MonitoringGetNetworkGatewayStatusResult MonitoringGetNetworkGatewayStatus(ctx, body)
Get Network Gateway monitoring status

Fetch Network Gateway monitoring status time series for the given time interval, all for the network or list of gateways, and list of metrics. On success, this method returns anarray of gateway monitoring statuses.  Privileges required:  `READ` `GATEWAY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MonitoringGetNetworkGatewayStatus**](MonitoringGetNetworkGatewayStatus.md)|  | 

### Return type

[**MonitoringGetNetworkGatewayStatusResult**](monitoring_get_network_gateway_status_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

