# \ClientDeviceApi

All URIs are relative to *http://localhost/portal/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SetClientDeviceHostName**](ClientDeviceApi.md#SetClientDeviceHostName) | **Post** /clientDevice/setClientDeviceHostName | Set hostname for client device



## SetClientDeviceHostName

> SetClientDeviceHostNameResult SetClientDeviceHostName(ctx, body)
Set hostname for client device

Sets the `hostName` for client device  Privileges required:  `UPDATE` `CLIENT_DEVICE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**SetClientDeviceHostName**](SetClientDeviceHostName.md)|  | 

### Return type

[**SetClientDeviceHostNameResult**](set_client_device_host_name_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

