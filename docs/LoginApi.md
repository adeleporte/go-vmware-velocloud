# \LoginApi

All URIs are relative to *http://localhost/portal/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LoginEnterpriseLogin**](LoginApi.md#LoginEnterpriseLogin) | **Post** /login/enterpriseLogin | Authenticate enterprise or partner (MSP) user
[**LoginOperatorLogin**](LoginApi.md#LoginOperatorLogin) | **Post** /login/operatorLogin | Authenticate operator user
[**Logout**](LoginApi.md#Logout) | **Post** /logout | Logout and invalidate authorization session cookie



## LoginEnterpriseLogin

> LoginEnterpriseLogin(ctx, authorization)
Authenticate enterprise or partner (MSP) user

Authenticates an enterprise or partner (MSP) user and, upon successful login, returns a velocloud.session cookie. Pass this session cookie in the authentication header in subsequent VCO calls.  If you are using an HTTP client (e.g. Postman) that is configured to automatically follow HTTP redirects, a successful authentication request will cause your client to follow an HTTP 302 redirect to the portal 'Home' web page. Your session cookie can then be used to make VCO API calls.  Note that session cookies expire after a period of time specified in the VCO configuration (default is 24 hours).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | [**AuthObject**](AuthObject.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoginOperatorLogin

> LoginOperatorLogin(ctx, authorization)
Authenticate operator user

Authenticates an operator user and, upon successful login, returns a velocloud.session cookie. Pass this session cookie in the authentication header in subsequent VCO calls.  If you are using an HTTP client (e.g. Postman) that is configured to automatically follow HTTP redirects, a successful authentication request will cause your client to follow an HTTP 302 redirect to the portal 'Home' web page. Your session cookie can then be used to make VCO API calls.   Note that session cookies expire after a period of time specified in the VCO configuration (default is 24 hours).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | [**AuthObject**](AuthObject.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Logout

> Logout(ctx, )
Logout and invalidate authorization session cookie

Logs out the VCO API user and invalidates the session cookie.

### Required Parameters

This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

