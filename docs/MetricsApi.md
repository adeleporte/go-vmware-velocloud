# \MetricsApi

All URIs are relative to *http://localhost/portal/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MetricsGetEdgeAppLinkMetrics**](MetricsApi.md#MetricsGetEdgeAppLinkMetrics) | **Post** /metrics/getEdgeAppLinkMetrics | Get flow metric aggregate data
[**MetricsGetEdgeAppLinkSeries**](MetricsApi.md#MetricsGetEdgeAppLinkSeries) | **Post** /metrics/getEdgeAppLinkSeries | Get flow metric time series data
[**MetricsGetEdgeAppMetrics**](MetricsApi.md#MetricsGetEdgeAppMetrics) | **Post** /metrics/getEdgeAppMetrics | Get flow metric aggregate data by application
[**MetricsGetEdgeAppSeries**](MetricsApi.md#MetricsGetEdgeAppSeries) | **Post** /metrics/getEdgeAppSeries | Get flow metric time series data by application
[**MetricsGetEdgeCategoryMetrics**](MetricsApi.md#MetricsGetEdgeCategoryMetrics) | **Post** /metrics/getEdgeCategoryMetrics | Get flow metric aggregate data by application category
[**MetricsGetEdgeCategorySeries**](MetricsApi.md#MetricsGetEdgeCategorySeries) | **Post** /metrics/getEdgeCategorySeries | Get flow metric time series data by application category
[**MetricsGetEdgeDestMetrics**](MetricsApi.md#MetricsGetEdgeDestMetrics) | **Post** /metrics/getEdgeDestMetrics | Get flow metric aggregate data by destination
[**MetricsGetEdgeDestSeries**](MetricsApi.md#MetricsGetEdgeDestSeries) | **Post** /metrics/getEdgeDestSeries | Get flow metric time series data by destination
[**MetricsGetEdgeDeviceMetrics**](MetricsApi.md#MetricsGetEdgeDeviceMetrics) | **Post** /metrics/getEdgeDeviceMetrics | Get flow metric aggregate data by client device
[**MetricsGetEdgeDeviceSeries**](MetricsApi.md#MetricsGetEdgeDeviceSeries) | **Post** /metrics/getEdgeDeviceSeries | Get flow metric time series data by client device
[**MetricsGetEdgeLinkMetrics**](MetricsApi.md#MetricsGetEdgeLinkMetrics) | **Post** /metrics/getEdgeLinkMetrics | Get advanced flow metric aggregate data by link
[**MetricsGetEdgeLinkSeries**](MetricsApi.md#MetricsGetEdgeLinkSeries) | **Post** /metrics/getEdgeLinkSeries | Get advanced flow metric time series data by link
[**MetricsGetEdgeOsMetrics**](MetricsApi.md#MetricsGetEdgeOsMetrics) | **Post** /metrics/getEdgeOsMetrics | Get flow metric aggregate data by client OS
[**MetricsGetEdgeOsSeries**](MetricsApi.md#MetricsGetEdgeOsSeries) | **Post** /metrics/getEdgeOsSeries | Get flow metric time series data by client OS
[**MetricsGetEdgeSegmentMetrics**](MetricsApi.md#MetricsGetEdgeSegmentMetrics) | **Post** /metrics/getEdgeSegmentMetrics | Get flow metric aggregate data by segment Id
[**MetricsGetEdgeSegmentSeries**](MetricsApi.md#MetricsGetEdgeSegmentSeries) | **Post** /metrics/getEdgeSegmentSeries | Get flow metric time series data by segment id
[**MetricsGetEdgeStatusMetrics**](MetricsApi.md#MetricsGetEdgeStatusMetrics) | **Post** /metrics/getEdgeStatusMetrics | Get Edge healthStats metrics for an interval
[**MetricsGetEdgeStatusSeries**](MetricsApi.md#MetricsGetEdgeStatusSeries) | **Post** /metrics/getEdgeStatusSeries | Get Edge healthStats time series for an interval 
[**MetricsGetGatewayStatusMetrics**](MetricsApi.md#MetricsGetGatewayStatusMetrics) | **Post** /metrics/getGatewayStatusMetrics | Get Gateway health metric summaries for an interval
[**MetricsGetGatewayStatusSeries**](MetricsApi.md#MetricsGetGatewayStatusSeries) | **Post** /metrics/getGatewayStatusSeries | Get Gateway health metric time series for an interval



## MetricsGetEdgeAppLinkMetrics

> []MetricsGetEdgeAppLinkMetricsResultItem MetricsGetEdgeAppLinkMetrics(ctx, body)
Get flow metric aggregate data

Gets flow metric summaries for the specified time interval by link. On success, this method returns an array of flow data in which each entry corresponds to a link on the specified Edge. In the request body, the `id` and `edgeId` property names are interchangeable. The `enterpriseId` property is required when this method is invoked in the operator context.  Privileges required:  `READ` `EDGE`  `VIEW_FLOW_STATS` `undefined`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MetricsGetEdgeAppLinkMetrics**](MetricsGetEdgeAppLinkMetrics.md)|  | 

### Return type

[**[]MetricsGetEdgeAppLinkMetricsResultItem**](metrics_get_edge_app_link_metrics_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsGetEdgeAppLinkSeries

> []MetricsGetEdgeAppLinkSeriesResultItem MetricsGetEdgeAppLinkSeries(ctx, body)
Get flow metric time series data

Gets flow metric time series data for the specified time interval by link. On success, this method returns an array of flow data in which each entry corresponds to a link on the specified Edge. In the request body, the `id` and `edgeId` property names are interchangeable. The `enterpriseId` property is required when this method is invoked in the operator context.  Privileges required:  `READ` `EDGE`  `VIEW_FLOW_STATS` `undefined`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MetricsGetEdgeAppLinkSeries**](MetricsGetEdgeAppLinkSeries.md)|  | 

### Return type

[**[]MetricsGetEdgeAppLinkSeriesResultItem**](metrics_get_edge_app_link_series_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsGetEdgeAppMetrics

> []MetricsGetEdgeAppMetricsResultItem MetricsGetEdgeAppMetrics(ctx, body)
Get flow metric aggregate data by application

Gets flow metric summaries for the specified time interval by application. On success, this method returns an array of flow data in which each entry corresponds to a single application. In the request body, the `id` and `edgeId` property names are interchangeable. The `enterpriseId` property is required when this method is invoked in the operator context.  Privileges required:  `READ` `EDGE`  `VIEW_FLOW_STATS` `undefined`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MetricsGetEdgeAppMetrics**](MetricsGetEdgeAppMetrics.md)|  | 

### Return type

[**[]MetricsGetEdgeAppMetricsResultItem**](metrics_get_edge_app_metrics_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsGetEdgeAppSeries

> []MetricsGetEdgeAppSeriesResultItem MetricsGetEdgeAppSeries(ctx, body)
Get flow metric time series data by application

Gets flow metric time series data for the specified time interval by application. On success, this method returns an array of flow data in which each entry corresponds to a single application. In the request body, the `id` and `edgeId` property names are interchangeable. The `enterpriseId` property is required when this method is invoked in the operator context.  Privileges required:  `READ` `EDGE`  `VIEW_FLOW_STATS` `undefined`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MetricsGetEdgeAppSeries**](MetricsGetEdgeAppSeries.md)|  | 

### Return type

[**[]MetricsGetEdgeAppSeriesResultItem**](metrics_get_edge_app_series_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsGetEdgeCategoryMetrics

> []MetricsGetEdgeCategoryMetricsResultItem MetricsGetEdgeCategoryMetrics(ctx, body)
Get flow metric aggregate data by application category

Gets flow metric summaries for the specified time interval by application category. On success, this method returns an array of flow data in which each entry corresponds to a category of application traffic that has traversed the specified Edge. In the request body, the `id` and `edgeId` property names are interchangeable. The `enterpriseId` property is required when this method is invoked in the operator context.  Privileges required:  `READ` `EDGE`  `VIEW_FLOW_STATS` `undefined`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MetricsGetEdgeCategoryMetrics**](MetricsGetEdgeCategoryMetrics.md)|  | 

### Return type

[**[]MetricsGetEdgeCategoryMetricsResultItem**](metrics_get_edge_category_metrics_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsGetEdgeCategorySeries

> []MetricsGetEdgeCategorySeriesResultItem MetricsGetEdgeCategorySeries(ctx, body)
Get flow metric time series data by application category

Gets flow metric time series data for the specified time interval by application category. On success, this method returns an array of flow data in which each entry corresponds to a category of application traffic that has traversed the specified Edge. In the request body, the `id` and `edgeId` property names are interchangeable. The `enterpriseId` property is required when this method is invoked in the operator context.  Privileges required:  `READ` `EDGE`  `VIEW_FLOW_STATS` `undefined`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MetricsGetEdgeCategorySeries**](MetricsGetEdgeCategorySeries.md)|  | 

### Return type

[**[]MetricsGetEdgeCategorySeriesResultItem**](metrics_get_edge_category_series_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsGetEdgeDestMetrics

> []MetricsGetEdgeDestMetricsResultItem MetricsGetEdgeDestMetrics(ctx, body)
Get flow metric aggregate data by destination

Gets flow metric summaries for the specified time interval by destination. On success, this method returns an array of flow data in which each entry corresponds to a distinct traffic destination. In the request body, the `id` and `edgeId` property names are interchangeable. The `enterpriseId` property is required when this method is invoked in the operator context.  Privileges required:  `READ` `EDGE`  `VIEW_FLOW_STATS` `undefined`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MetricsGetEdgeDestMetrics**](MetricsGetEdgeDestMetrics.md)|  | 

### Return type

[**[]MetricsGetEdgeDestMetricsResultItem**](metrics_get_edge_dest_metrics_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsGetEdgeDestSeries

> []MetricsGetEdgeDestSeriesResultItem MetricsGetEdgeDestSeries(ctx, body)
Get flow metric time series data by destination

Gets flow metric time series data for the specified time interval by destination. On success, this method returns an array of flow data in which each entry corresponds to a distinct traffic destination. In the request body, the `id` and `edgeId` property names are interchangeable. The `enterpriseId` property is required when this method is invoked in the operator context.  Privileges required:  `READ` `EDGE`  `VIEW_FLOW_STATS` `undefined`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MetricsGetEdgeDestSeries**](MetricsGetEdgeDestSeries.md)|  | 

### Return type

[**[]MetricsGetEdgeDestSeriesResultItem**](metrics_get_edge_dest_series_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsGetEdgeDeviceMetrics

> []MetricsGetEdgeDeviceMetricsResultItem MetricsGetEdgeDeviceMetrics(ctx, body)
Get flow metric aggregate data by client device

Gets flow metric summaries for the specified time interval by client device. On success, this method returns an array of flow data in which each entry corresponds to a distinct device. In the request body, the `id` and `edgeId` property names are interchangeable. The `enterpriseId` property is required when this method is invoked in the operator context.  Privileges required:  `READ` `EDGE`  `VIEW_USER_IDENTIFIABLE_FLOW_STATS` `undefined`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MetricsGetEdgeDeviceMetrics**](MetricsGetEdgeDeviceMetrics.md)|  | 

### Return type

[**[]MetricsGetEdgeDeviceMetricsResultItem**](metrics_get_edge_device_metrics_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsGetEdgeDeviceSeries

> []MetricsGetEdgeDeviceSeriesResultItem MetricsGetEdgeDeviceSeries(ctx, body)
Get flow metric time series data by client device

Gets flow metric time series data for the specified time interval by client device. On success, this method returns an array of flow data in which each entry corresponds to a distinct device. In the request body, the `id` and `edgeId` property names are interchangeable. The `enterpriseId` property is required when this method is invoked in the operator context.  Privileges required:  `READ` `EDGE`  `VIEW_USER_IDENTIFIABLE_FLOW_STATS` `undefined`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MetricsGetEdgeDeviceSeries**](MetricsGetEdgeDeviceSeries.md)|  | 

### Return type

[**[]MetricsGetEdgeDeviceSeriesResultItem**](metrics_get_edge_device_series_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsGetEdgeLinkMetrics

> []MetricsGetEdgeLinkMetricsResultItem MetricsGetEdgeLinkMetrics(ctx, body)
Get advanced flow metric aggregate data by link

Gets advanced flow metric summaries for the specified time interval by link. On success, this method returns an array of flow data in which each entry corresponds to a link on the specified Edge. In the request body, the `id` and `edgeId` property names are interchangeable. The `enterpriseId` property is required when this method is invoked in the operator context.  Privileges required:  `READ` `EDGE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MetricsGetEdgeLinkMetrics**](MetricsGetEdgeLinkMetrics.md)|  | 

### Return type

[**[]MetricsGetEdgeLinkMetricsResultItem**](metrics_get_edge_link_metrics_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsGetEdgeLinkSeries

> []MetricsGetEdgeLinkSeriesResultItem MetricsGetEdgeLinkSeries(ctx, body)
Get advanced flow metric time series data by link

Gets advanced flow metric time series data for the specified time interval by link. On success, this method returns an array of flow data in which each entry corresponds to a link on the specified Edge. In the request body, the `id` and `edgeId` property names are interchangeable. The `enterpriseId` property is required when this method is invoked in the operator context.  Privileges required:  `READ` `EDGE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MetricsGetEdgeLinkSeries**](MetricsGetEdgeLinkSeries.md)|  | 

### Return type

[**[]MetricsGetEdgeLinkSeriesResultItem**](metrics_get_edge_link_series_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsGetEdgeOsMetrics

> []MetricsGetEdgeOsMetricsResultItem MetricsGetEdgeOsMetrics(ctx, body)
Get flow metric aggregate data by client OS

Gets flow metric summaries for the specified time interval by client OS. On success, this method returns an array of flow data in which each entry corresponds to a distinct OS on a client device. In the request body, the `id` and `edgeId` property names are interchangeable. The `enterpriseId` property is required when this method is invoked in the operator context.  Privileges required:  `READ` `EDGE`  `VIEW_FLOW_STATS` `undefined`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MetricsGetEdgeOsMetrics**](MetricsGetEdgeOsMetrics.md)|  | 

### Return type

[**[]MetricsGetEdgeOsMetricsResultItem**](metrics_get_edge_os_metrics_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsGetEdgeOsSeries

> []MetricsGetEdgeOsSeriesResultItem MetricsGetEdgeOsSeries(ctx, body)
Get flow metric time series data by client OS

Gets flow metric time series data for the specified time interval by client OS. On success, this method returns an array of flow data in which each entry corresponds to a distinct OS on a client device. In the request body, the `id` and `edgeId` property names are interchangeable. The `enterpriseId` property is required when this method is invoked in the operator context.  Privileges required:  `READ` `EDGE`  `VIEW_FLOW_STATS` `undefined`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MetricsGetEdgeOsSeries**](MetricsGetEdgeOsSeries.md)|  | 

### Return type

[**[]MetricsGetEdgeOsSeriesResultItem**](metrics_get_edge_os_series_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsGetEdgeSegmentMetrics

> []MetricsGetEdgeSegmentMetricsResultItem MetricsGetEdgeSegmentMetrics(ctx, body)
Get flow metric aggregate data by segment Id

Gets flow metric summaries for the specified time interval by segment id. On success, this method returns an array of flow data where each entry corresponds to a segment id traffic that has traversed the specified Edge. In the request body, the `id` and `edgeId` property names are interchangeable. The `enterpriseId` property is required when this method is invoked in the operator context.  Privileges required:  `READ` `EDGE`  `VIEW_FLOW_STATS` `undefined`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MetricsGetEdgeSegmentMetrics**](MetricsGetEdgeSegmentMetrics.md)|  | 

### Return type

[**[]MetricsGetEdgeSegmentMetricsResultItem**](metrics_get_edge_segment_metrics_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsGetEdgeSegmentSeries

> []MetricsGetEdgeSegmentSeriesResultItem MetricsGetEdgeSegmentSeries(ctx, body)
Get flow metric time series data by segment id

Gets flow metric time series data for the specified time interval by segment id. On success, this method returns an array of flow data in which each entry corresponds to a segment id of traffic that has traversed the specified Edge. In the request body, the `id` and `edgeId` property names are interchangeable. The `enterpriseId` property is required when this method is invoked in the operator context.  Privileges required:  `READ` `EDGE`  `VIEW_FLOW_STATS` `undefined`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MetricsGetEdgeSegmentSeries**](MetricsGetEdgeSegmentSeries.md)|  | 

### Return type

[**[]MetricsGetEdgeSegmentSeriesResultItem**](metrics_get_edge_segment_series_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsGetEdgeStatusMetrics

> EdgeStatusMetricsSummary MetricsGetEdgeStatusMetrics(ctx, body)
Get Edge healthStats metrics for an interval

Fetch healthStats metric summaries for the given edge, time interval and list of metrics. On success, this method returns anhealthsStats object with min, max and average values of requested metrics  Privileges required:  `READ` `EDGE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MetricsGetEdgeStatusMetrics**](MetricsGetEdgeStatusMetrics.md)|  | 

### Return type

[**EdgeStatusMetricsSummary**](edge_status_metrics_summary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsGetEdgeStatusSeries

> GetEdgeStatusMetricsTimeSeriesResult MetricsGetEdgeStatusSeries(ctx, body)
Get Edge healthStats time series for an interval 

Fetch healthStats time series for the given edge, time interval and list of metrics. On success, this method returns anarray of healthsStats series of requested metrics  Privileges required:  `READ` `EDGE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MetricsGetEdgeStatusSeries**](MetricsGetEdgeStatusSeries.md)|  | 

### Return type

[**GetEdgeStatusMetricsTimeSeriesResult**](get_edge_status_metrics_time_series_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsGetGatewayStatusMetrics

> GatewayStatusMetricsSummary MetricsGetGatewayStatusMetrics(ctx, body)
Get Gateway health metric summaries for an interval

Fetch health metric summaries given a target gateway, time interval and list of metrics. On success, this method returns an object containing min, max and average values of requested metrics  Privileges required:  `READ` `GATEWAY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MetricsGetGatewayStatusMetrics**](MetricsGetGatewayStatusMetrics.md)|  | 

### Return type

[**GatewayStatusMetricsSummary**](gateway_status_metrics_summary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsGetGatewayStatusSeries

> GetGatewayStatusMetricsTimeSeriesResult MetricsGetGatewayStatusSeries(ctx, body)
Get Gateway health metric time series for an interval

Fetch health metric time series given a target gateway, time interval and list of metrics. On success, this method returns an array of time series, one per metric.  Privileges required:  `READ` `GATEWAY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MetricsGetGatewayStatusSeries**](MetricsGetGatewayStatusSeries.md)|  | 

### Return type

[**GetGatewayStatusMetricsTimeSeriesResult**](get_gateway_status_metrics_time_series_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

