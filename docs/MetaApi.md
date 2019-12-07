# \MetaApi

All URIs are relative to *http://localhost/portal/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Meta**](MetaApi.md#Meta) | **Post** /meta/{apiPath} | Get Swagger specification for any VCO API call



## Meta

> InlineResponse200 Meta(ctx, apiPath)
Get Swagger specification for any VCO API call

Gets the Swagger specification for any VCO API call.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiPath** | **string**| the path to another api method, starting after /rest/ | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

