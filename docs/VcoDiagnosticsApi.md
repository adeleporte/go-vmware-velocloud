# \VcoDiagnosticsApi

All URIs are relative to *http://localhost/portal/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VcoDiagnosticsGetVcoDbDiagnostics**](VcoDiagnosticsApi.md#VcoDiagnosticsGetVcoDbDiagnostics) | **Post** /vcoDiagnostics/getVcoDbDiagnostics | Get VCO Database Diagnostics



## VcoDiagnosticsGetVcoDbDiagnostics

> []InlineResponse2002 VcoDiagnosticsGetVcoDbDiagnostics(ctx, body)
Get VCO Database Diagnostics

Gets the diagnostic information of the VCO database.  Privileges required:  `READ` `VCO_DIAGNOSTICS`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | **map[string]interface{}**|  | 

### Return type

[**[]InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

