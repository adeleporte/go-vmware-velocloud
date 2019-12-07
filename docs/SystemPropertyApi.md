# \SystemPropertyApi

All URIs are relative to *http://localhost/portal/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SystemPropertyGetSystemProperties**](SystemPropertyApi.md#SystemPropertyGetSystemProperties) | **Post** /systemProperty/getSystemProperties | Get all system properties
[**SystemPropertyGetSystemProperty**](SystemPropertyApi.md#SystemPropertyGetSystemProperty) | **Post** /systemProperty/getSystemProperty | Get system property
[**SystemPropertyInsertOrUpdateSystemProperty**](SystemPropertyApi.md#SystemPropertyInsertOrUpdateSystemProperty) | **Post** /systemProperty/insertOrUpdateSystemProperty | Create or update system property
[**SystemPropertyInsertSystemProperty**](SystemPropertyApi.md#SystemPropertyInsertSystemProperty) | **Post** /systemProperty/insertSystemProperty | Create system property
[**SystemPropertyUpdateSystemProperty**](SystemPropertyApi.md#SystemPropertyUpdateSystemProperty) | **Post** /systemProperty/updateSystemProperty | Update system property



## SystemPropertyGetSystemProperties

> []SystemPropertyGetSystemPropertiesResultItem SystemPropertyGetSystemProperties(ctx, body)
Get all system properties

Gets all configured system properties.  Privileges required:  `READ` `SYSTEM_PROPERTY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**SystemPropertyGetSystemProperties**](SystemPropertyGetSystemProperties.md)|  | 

### Return type

[**[]SystemPropertyGetSystemPropertiesResultItem**](system_property_get_system_properties_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemPropertyGetSystemProperty

> SystemPropertyGetSystemPropertyResult SystemPropertyGetSystemProperty(ctx, body)
Get system property

Gets the specified system property. Specify by object `id` or `name`.  Privileges required:  `READ` `SYSTEM_PROPERTY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**SystemPropertyGetSystemProperty**](SystemPropertyGetSystemProperty.md)|  | 

### Return type

[**SystemPropertyGetSystemPropertyResult**](system_property_get_system_property_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemPropertyInsertOrUpdateSystemProperty

> SystemPropertyInsertOrUpdateSystemPropertyResult SystemPropertyInsertOrUpdateSystemProperty(ctx, body)
Create or update system property

Creates the specified system property. If the system property with the specified `name` already exists, then the API call updates it.  Privileges required:  `CREATE` `SYSTEM_PROPERTY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**SystemPropertyInsertOrUpdateSystemProperty**](SystemPropertyInsertOrUpdateSystemProperty.md)|  | 

### Return type

[**SystemPropertyInsertOrUpdateSystemPropertyResult**](system_property_insert_or_update_system_property_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemPropertyInsertSystemProperty

> SystemPropertyInsertSystemPropertyResult SystemPropertyInsertSystemProperty(ctx, body)
Create system property

Creates a new system property with the specified `name`, `value`, and optionally other attributes.  Privileges required:  `CREATE` `SYSTEM_PROPERTY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**SystemPropertyInsertSystemProperty**](SystemPropertyInsertSystemProperty.md)|  | 

### Return type

[**SystemPropertyInsertSystemPropertyResult**](system_property_insert_system_property_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemPropertyUpdateSystemProperty

> SystemPropertyUpdateSystemPropertyResult SystemPropertyUpdateSystemProperty(ctx, body)
Update system property

Updates the specified system property (by object `id` or `name`). Expects an `_update` object containing the field(s) to be updated and their corresponding value(s).  Privileges required:  `UPDATE` `SYSTEM_PROPERTY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**SystemPropertyUpdateSystemProperty**](SystemPropertyUpdateSystemProperty.md)|  | 

### Return type

[**SystemPropertyUpdateSystemPropertyResult**](system_property_update_system_property_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

