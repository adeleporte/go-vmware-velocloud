# \ConfigurationApi

All URIs are relative to *http://localhost/portal/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfigurationCloneAndConvertConfiguration**](ConfigurationApi.md#ConfigurationCloneAndConvertConfiguration) | **Post** /configuration/cloneAndConvertConfiguration | Create new segment-based profile from existing network-based profile
[**ConfigurationCloneConfiguration**](ConfigurationApi.md#ConfigurationCloneConfiguration) | **Post** /configuration/cloneConfiguration | Clone configuration profile
[**ConfigurationCloneEnterpriseTemplate**](ConfigurationApi.md#ConfigurationCloneEnterpriseTemplate) | **Post** /configuration/cloneEnterpriseTemplate | Clone default enterprise configuration profile
[**ConfigurationDeleteConfiguration**](ConfigurationApi.md#ConfigurationDeleteConfiguration) | **Post** /configuration/deleteConfiguration | Delete configuration profile
[**ConfigurationGetConfiguration**](ConfigurationApi.md#ConfigurationGetConfiguration) | **Post** /configuration/getConfiguration | Get configuration profile
[**ConfigurationGetConfigurationModules**](ConfigurationApi.md#ConfigurationGetConfigurationModules) | **Post** /configuration/getConfigurationModules | Get configuration profile modules
[**ConfigurationGetIdentifiableApplications**](ConfigurationApi.md#ConfigurationGetIdentifiableApplications) | **Post** /configuration/getIdentifiableApplications | Get the applications that are identifiable through DPI.
[**ConfigurationGetRoutableApplications**](ConfigurationApi.md#ConfigurationGetRoutableApplications) | **Post** /configuration/getRoutableApplications | Get first packet routable applications
[**ConfigurationInsertConfigurationModule**](ConfigurationApi.md#ConfigurationInsertConfigurationModule) | **Post** /configuration/insertConfigurationModule | Create configuration module
[**ConfigurationUpdateConfiguration**](ConfigurationApi.md#ConfigurationUpdateConfiguration) | **Post** /configuration/updateConfiguration | Update configuration profile
[**ConfigurationUpdateConfigurationModule**](ConfigurationApi.md#ConfigurationUpdateConfigurationModule) | **Post** /configuration/updateConfigurationModule | Update configuration module



## ConfigurationCloneAndConvertConfiguration

> ConfigurationCloneAndConvertConfigurationResult ConfigurationCloneAndConvertConfiguration(ctx, body)
Create new segment-based profile from existing network-based profile

Clones and converts the specified network configuration (by `configurationId`). Accepts an `enterpriseId` or `networkId` to associate the new configuration with an enterprise or network. On success, returns the ID of the newly created configuration object.  Privileges required:  `CREATE` `ENTERPRISE_PROFILE`, or  `CREATE` `OPERATOR_PROFILE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**ConfigurationCloneAndConvertConfiguration**](ConfigurationCloneAndConvertConfiguration.md)|  | 

### Return type

[**ConfigurationCloneAndConvertConfigurationResult**](configuration_clone_and_convert_configuration_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigurationCloneConfiguration

> ConfigurationCloneConfigurationResult ConfigurationCloneConfiguration(ctx, body)
Clone configuration profile

Clones the specified configuration (by `configurationId`) and all associated configuration modules. Accepts an `enterpriseId` or `networkId` to associate the new configuration with an enterprise or network. Select modules may also be specified. On success, returns the `id` of the newly created configuration object.  Privileges required:  `CREATE` `ENTERPRISE_PROFILE`, or  `CREATE` `OPERATOR_PROFILE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**ConfigurationCloneConfiguration**](ConfigurationCloneConfiguration.md)|  | 

### Return type

[**ConfigurationCloneConfigurationResult**](configuration_clone_configuration_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigurationCloneEnterpriseTemplate

> ConfigurationCloneEnterpriseTemplateResult ConfigurationCloneEnterpriseTemplate(ctx, body)
Clone default enterprise configuration profile

Creates a new enterprise configuration from the enterprise default configuration. On success, returns the `id` of the newly created configuration object.  Privileges required:  `CREATE` `ENTERPRISE_PROFILE`, or  `CREATE` `OPERATOR_PROFILE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**ConfigurationCloneEnterpriseTemplate**](ConfigurationCloneEnterpriseTemplate.md)|  | 

### Return type

[**ConfigurationCloneEnterpriseTemplateResult**](configuration_clone_enterprise_template_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigurationDeleteConfiguration

> ConfigurationDeleteConfigurationResult ConfigurationDeleteConfiguration(ctx, body)
Delete configuration profile

Deletes the specified configuration profile (by `id`). On success, returns an object indicating the number of objects (rows) deleted (1 or 0).  Privileges required:  `DELETE` `ENTERPRISE_PROFILE`, or  `DELETE` `OPERATOR_PROFILE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**ConfigurationDeleteConfiguration**](ConfigurationDeleteConfiguration.md)|  | 

### Return type

[**ConfigurationDeleteConfigurationResult**](configuration_delete_configuration_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigurationGetConfiguration

> ConfigurationGetConfigurationResult ConfigurationGetConfiguration(ctx, body)
Get configuration profile

Gets the specified configuration profile, optionally with module details.  Privileges required:  `READ` `ENTERPRISE_PROFILE`, or  `READ` `OPERATOR_PROFILE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**ConfigurationGetConfiguration**](ConfigurationGetConfiguration.md)|  | 

### Return type

[**ConfigurationGetConfigurationResult**](configuration_get_configuration_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigurationGetConfigurationModules

> []ConfigurationGetConfigurationModulesResultItem ConfigurationGetConfigurationModules(ctx, body)
Get configuration profile modules

Gets all configuration modules for the specified configuration profile.  Privileges required:  `READ` `ENTERPRISE_PROFILE`, or  `READ` `OPERATOR_PROFILE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**ConfigurationGetConfigurationModules**](ConfigurationGetConfigurationModules.md)|  | 

### Return type

[**[]ConfigurationGetConfigurationModulesResultItem**](configuration_get_configuration_modules_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigurationGetIdentifiableApplications

> ConfigurationGetIdentifiableApplicationsResult ConfigurationGetIdentifiableApplications(ctx, body)
Get the applications that are identifiable through DPI.

Gets all applications that are identifiable through DPI. If called from an operator or MSP context, then `enterpriseId` is required.  Privileges required:  `READ` `ENTERPRISE_PROFILE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**ConfigurationGetIdentifiableApplications**](ConfigurationGetIdentifiableApplications.md)|  | 

### Return type

[**ConfigurationGetIdentifiableApplicationsResult**](configuration_get_identifiable_applications_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigurationGetRoutableApplications

> ConfigurationGetRoutableApplicationsResult ConfigurationGetRoutableApplications(ctx, body)
Get first packet routable applications

Gets all applications that are first packet routable. If called from an operator or MSP context, then `enterpriseId` is required.Optionally, specify `edgeId` to get the map for a specific Edge.  Privileges required:  `VIEW_FLOW_STATS` `undefined`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**ConfigurationGetRoutableApplications**](ConfigurationGetRoutableApplications.md)|  | 

### Return type

[**ConfigurationGetRoutableApplicationsResult**](configuration_get_routable_applications_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigurationInsertConfigurationModule

> ConfigurationInsertConfigurationModuleResult ConfigurationInsertConfigurationModule(ctx, body)
Create configuration module

Creates a new configuration module for the specified configuration profile.  Privileges required:  `UPDATE` `ENTERPRISE_PROFILE`, or  `UPDATE` `OPERATOR_PROFILE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**ConfigurationInsertConfigurationModule**](ConfigurationInsertConfigurationModule.md)|  | 

### Return type

[**ConfigurationInsertConfigurationModuleResult**](configuration_insert_configuration_module_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigurationUpdateConfiguration

> RowsModifiedConfirmation ConfigurationUpdateConfiguration(ctx, body)
Update configuration profile

Updates the specified configuration profile record according to the attribute:value mappings specified in the `_update` object.  Privileges required:  `UPDATE` `ENTERPRISE_PROFILE`, or  `UPDATE` `OPERATOR_PROFILE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**ConfigurationUpdateConfiguration**](ConfigurationUpdateConfiguration.md)|  | 

### Return type

[**RowsModifiedConfirmation**](rows_modified_confirmation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigurationUpdateConfigurationModule

> ConfigurationUpdateConfigurationModuleResult ConfigurationUpdateConfigurationModule(ctx, body)
Update configuration module

Updates the specified configuration module. Expects an `_update` object containing the field(s) to be updated and their corresponding value(s).  Privileges required:  `UPDATE` `ENTERPRISE_PROFILE`, or  `UPDATE` `OPERATOR_PROFILE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**ConfigurationUpdateConfigurationModule**](ConfigurationUpdateConfigurationModule.md)|  | 

### Return type

[**ConfigurationUpdateConfigurationModuleResult**](configuration_update_configuration_module_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

