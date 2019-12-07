# \LinkQualityEventApi

All URIs are relative to *http://localhost/portal/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LinkQualityEventGetLinkQualityEvents**](LinkQualityEventApi.md#LinkQualityEventGetLinkQualityEvents) | **Post** /linkQualityEvent/getLinkQualityEvents | Get link quality events



## LinkQualityEventGetLinkQualityEvents

> LinkQualityEventGetLinkQualityEventsResult LinkQualityEventGetLinkQualityEvents(ctx, body)
Get link quality events

Gets all link quality scores per link for the specified Edge within the specified time interval. Rolls up link quality events to provide an aggregate score for the Edge. Returns an empty array if no link quality events are available in the specified timeframe.  Privileges required:  `READ` `EDGE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**LinkQualityEventGetLinkQualityEvents**](LinkQualityEventGetLinkQualityEvents.md)|  | 

### Return type

[**LinkQualityEventGetLinkQualityEventsResult**](link_quality_event_get_link_quality_events_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

