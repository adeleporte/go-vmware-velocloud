# \EventApi

All URIs are relative to *http://localhost/portal/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EventGetEnterpriseEvents**](EventApi.md#EventGetEnterpriseEvents) | **Post** /event/getEnterpriseEvents | Get Edge events
[**EventGetOperatorEvents**](EventApi.md#EventGetOperatorEvents) | **Post** /event/getOperatorEvents | Get operator events
[**EventGetProxyEvents**](EventApi.md#EventGetProxyEvents) | **Post** /event/getProxyEvents | Fetch Partner events



## EventGetEnterpriseEvents

> EventGetEnterpriseEventsResult EventGetEnterpriseEvents(ctx, body)
Get Edge events

Gets Edge events in an enterprise or Edge context. Returns an array of Edge events sorted by `eventTime`..  Privileges required:  `READ` `ENTERPRISE_EVENT`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EventGetEnterpriseEvents**](EventGetEnterpriseEvents.md)|  | 

### Return type

[**EventGetEnterpriseEventsResult**](event_get_enterprise_events_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventGetOperatorEvents

> EventGetOperatorEventsResult EventGetOperatorEvents(ctx, body)
Get operator events

Gets operator events by network ID (optional). If not specified, will be taken for the caller's security context. Optionally, use a filter object to limit the number of events returned. Additionally, specify a time interval with an interval object. If no end date is specified, then the default is the current date. Specify a `gatewayId` to filter events for the specified gateway.  Privileges required:  `READ` `OPERATOR_EVENT`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EventGetOperatorEvents**](EventGetOperatorEvents.md)|  | 

### Return type

[**EventGetOperatorEventsResult**](event_get_operator_events_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventGetProxyEvents

> EventGetProxyEventsResult EventGetProxyEvents(ctx, body)
Fetch Partner events

Fetch Partner Events by enterprise ID. Optionally, a filter object to specify search criteria and limit on number of records. A time interval could be specified with an interval object. If no end date is given, it will default to the current date. Enterprise ID can be supplied to filter events to those of a specific enterprises.  Privileges required:  `READ` `PROXY_EVENT`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EventGetProxyEvents**](EventGetProxyEvents.md)|  | 

### Return type

[**EventGetProxyEventsResult**](event_get_proxy_events_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

