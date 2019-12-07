# \EdgeApi

All URIs are relative to *http://localhost/portal/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EdgeDeleteEdge**](EdgeApi.md#EdgeDeleteEdge) | **Post** /edge/deleteEdge | Delete Edge
[**EdgeDeleteEdgeBgpNeighborRecords**](EdgeApi.md#EdgeDeleteEdgeBgpNeighborRecords) | **Post** /edge/deleteEdgeBgpNeighborRecords | Delete Edge BGP neighbor records
[**EdgeEdgeCancelReactivation**](EdgeApi.md#EdgeEdgeCancelReactivation) | **Post** /edge/edgeCancelReactivation | Cancel pending Edge reactivation request
[**EdgeEdgeProvision**](EdgeApi.md#EdgeEdgeProvision) | **Post** /edge/edgeProvision | Provision Edge
[**EdgeEdgeRequestReactivation**](EdgeApi.md#EdgeEdgeRequestReactivation) | **Post** /edge/edgeRequestReactivation | Prepare Edge for reactivation
[**EdgeGetClientVisibilityMode**](EdgeApi.md#EdgeGetClientVisibilityMode) | **Post** /edge/getClientVisibilityMode | Get an edge&#39;s client visibility mode
[**EdgeGetEdge**](EdgeApi.md#EdgeGetEdge) | **Post** /edge/getEdge | Get Edge
[**EdgeGetEdgeConfigurationModules**](EdgeApi.md#EdgeGetEdgeConfigurationModules) | **Post** /edge/getEdgeConfigurationModules | Get edge configuration modules
[**EdgeGetEdgeConfigurationStack**](EdgeApi.md#EdgeGetEdgeConfigurationStack) | **Post** /edge/getEdgeConfigurationStack | Get Edge configuration stack
[**EdgeSetEdgeEnterpriseConfiguration**](EdgeApi.md#EdgeSetEdgeEnterpriseConfiguration) | **Post** /edge/setEdgeEnterpriseConfiguration | Apply configuration profile to Edge
[**EdgeSetEdgeHandOffGateways**](EdgeApi.md#EdgeSetEdgeHandOffGateways) | **Post** /edge/setEdgeHandOffGateways | Set Edge on premise hand-off gateway(s)
[**EdgeSetEdgeOperatorConfiguration**](EdgeApi.md#EdgeSetEdgeOperatorConfiguration) | **Post** /edge/setEdgeOperatorConfiguration | Apply operator profile to Edge
[**EdgeUpdateEdgeAdminPassword**](EdgeApi.md#EdgeUpdateEdgeAdminPassword) | **Post** /edge/updateEdgeAdminPassword | Update Edge local UI authentication credentials
[**EdgeUpdateEdgeAttributes**](EdgeApi.md#EdgeUpdateEdgeAttributes) | **Post** /edge/updateEdgeAttributes | Update Edge attributes
[**EdgeUpdateEdgeCredentialsByConfiguration**](EdgeApi.md#EdgeUpdateEdgeCredentialsByConfiguration) | **Post** /edge/updateEdgeCredentialsByConfiguration | Update Edge local UI credentials by configuration profile



## EdgeDeleteEdge

> []EdgeDeleteEdgeResultItem EdgeDeleteEdge(ctx, body)
Delete Edge

Deletes the specified Edge(s) (by `id` or an array of `ids`).  Privileges required:  `DELETE` `EDGE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EdgeDeleteEdge**](EdgeDeleteEdge.md)|  | 

### Return type

[**[]EdgeDeleteEdgeResultItem**](edge_delete_edge_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeDeleteEdgeBgpNeighborRecords

> EdgeDeleteEdgeBgpNeighborRecordsResult EdgeDeleteEdgeBgpNeighborRecords(ctx, body)
Delete Edge BGP neighbor records

Deletes BGP record(s) matching the specified record keys (`neighborIp`) on the Edges with the specified `edgeId`s, if they exist.  Privileges required:  `DELETE` `NETWORK_SERVICE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EdgeDeleteEdgeBgpNeighborRecords**](EdgeDeleteEdgeBgpNeighborRecords.md)|  | 

### Return type

[**EdgeDeleteEdgeBgpNeighborRecordsResult**](edge_delete_edge_bgp_neighbor_records_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeEdgeCancelReactivation

> EdgeEdgeCancelReactivationResult EdgeEdgeCancelReactivation(ctx, body)
Cancel pending Edge reactivation request

Cancels a pending Edge reactivation request for the specified Edge (by `edgeId`).  Privileges required:  `CREATE` `EDGE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EdgeEdgeCancelReactivation**](EdgeEdgeCancelReactivation.md)|  | 

### Return type

[**EdgeEdgeCancelReactivationResult**](edge_edge_cancel_reactivation_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeEdgeProvision

> EdgeEdgeProvisionResult EdgeEdgeProvision(ctx, body)
Provision Edge

Provisions an Edge before activation.  Privileges required:  `CREATE` `EDGE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EdgeEdgeProvision**](EdgeEdgeProvision.md)|  | 

### Return type

[**EdgeEdgeProvisionResult**](edge_edge_provision_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeEdgeRequestReactivation

> EdgeEdgeRequestReactivationResult EdgeEdgeRequestReactivation(ctx, body)
Prepare Edge for reactivation

Updates the activation state of the specified Edge to `REACTIVATION_PENDING`.  Privileges required:  `CREATE` `EDGE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EdgeEdgeRequestReactivation**](EdgeEdgeRequestReactivation.md)|  | 

### Return type

[**EdgeEdgeRequestReactivationResult**](edge_edge_request_reactivation_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeGetClientVisibilityMode

> EdgeGetClientVisibilityModeResult EdgeGetClientVisibilityMode(ctx, body)
Get an edge's client visibility mode

Retrieve an edge's client visibility mode.  Privileges required:  `READ` `EDGE`  `VIEW_FLOW_STATS` `undefined`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EdgeGetClientVisibilityMode**](EdgeGetClientVisibilityMode.md)|  | 

### Return type

[**EdgeGetClientVisibilityModeResult**](edge_get_client_visibility_mode_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeGetEdge

> EdgeGetEdgeResult EdgeGetEdge(ctx, body)
Get Edge

Gets the specified Edge with optional link, site, configuration, certificate, orenterprise details. Supports queries by Edge `id`, `deviceId`, `activationKey`, and `logicalId`.  Privileges required:  `READ` `EDGE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EdgeGetEdge**](EdgeGetEdge.md)|  | 

### Return type

[**EdgeGetEdgeResult**](edge_get_edge_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeGetEdgeConfigurationModules

> EdgeGetEdgeConfigurationModulesResult EdgeGetEdgeConfigurationModules(ctx, body)
Get edge configuration modules

Gets edge composite configuration modules for the specified Edge (by `edgeId`).  Privileges required:  `READ` `EDGE`  `READ` `ENTERPRISE_PROFILE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EdgeGetEdgeConfigurationModules**](EdgeGetEdgeConfigurationModules.md)|  | 

### Return type

[**EdgeGetEdgeConfigurationModulesResult**](edge_get_edge_configuration_modules_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeGetEdgeConfigurationStack

> []EdgeGetEdgeConfigurationStackResultItem EdgeGetEdgeConfigurationStack(ctx, body)
Get Edge configuration stack

Gets the complete configuration profile (including all modules) of the specified Edge (by `edgeId`).  Privileges required:  `READ` `EDGE`  `READ` `ENTERPRISE_PROFILE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EdgeGetEdgeConfigurationStack**](EdgeGetEdgeConfigurationStack.md)|  | 

### Return type

[**[]EdgeGetEdgeConfigurationStackResultItem**](edge_get_edge_configuration_stack_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeSetEdgeEnterpriseConfiguration

> EdgeSetEdgeEnterpriseConfigurationResult EdgeSetEdgeEnterpriseConfiguration(ctx, body)
Apply configuration profile to Edge

Sets the enterprise configuration for the specified Edge (by `edgeId`).  Privileges required:  `UPDATE` `EDGE`  `UPDATE` `ENTERPRISE_PROFILE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EdgeSetEdgeEnterpriseConfiguration**](EdgeSetEdgeEnterpriseConfiguration.md)|  | 

### Return type

[**EdgeSetEdgeEnterpriseConfigurationResult**](edge_set_edge_enterprise_configuration_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeSetEdgeHandOffGateways

> EdgeSetEdgeHandOffGatewaysResult EdgeSetEdgeHandOffGateways(ctx, body)
Set Edge on premise hand-off gateway(s)

Sets the on-premise hand off gateways for the specified Edge (by `edgeId`). A primary and secondary gateway are defined. The primary is required, the secondary is optional. Moves all existing Edge-gateway hand-off relationships and replaces them with the specified primary and secondary gateways.  Privileges required:  `UPDATE` `EDGE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EdgeSetEdgeHandOffGateways**](EdgeSetEdgeHandOffGateways.md)|  | 

### Return type

[**EdgeSetEdgeHandOffGatewaysResult**](edge_set_edge_hand_off_gateways_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeSetEdgeOperatorConfiguration

> EdgeSetEdgeOperatorConfigurationResult EdgeSetEdgeOperatorConfiguration(ctx, body)
Apply operator profile to Edge

Sets the operator configuration for the specified Edge (by `edgeId`). This overrides any enterprise-assigned operator configuration and the network default operator configuration.  Privileges required:  `UPDATE` `EDGE`  `READ` `OPERATOR_PROFILE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EdgeSetEdgeOperatorConfiguration**](EdgeSetEdgeOperatorConfiguration.md)|  | 

### Return type

[**EdgeSetEdgeOperatorConfigurationResult**](edge_set_edge_operator_configuration_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeUpdateEdgeAdminPassword

> EdgeUpdateEdgeAdminPasswordResult EdgeUpdateEdgeAdminPassword(ctx, body)
Update Edge local UI authentication credentials

Requests an update to the Edge's local UI authentication credentials. On success, returns a JSON object with the ID of the action queued, the status of which can be queried using the `edgeAction/getEdgeAction` API.  Privileges required:  `UPDATE` `EDGE`  `UPDATE` `ENTERPRISE_KEYS`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EdgeUpdateEdgeAdminPassword**](EdgeUpdateEdgeAdminPassword.md)|  | 

### Return type

[**EdgeUpdateEdgeAdminPasswordResult**](edge_update_edge_admin_password_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeUpdateEdgeAttributes

> EdgeUpdateEdgeAttributesResult EdgeUpdateEdgeAttributes(ctx, body)
Update Edge attributes

Updates basic attributes (including Edge name, description, site information, and serial number) for the specified Edge. Expects an `_update` object containing the field(s) to be updated and their corresponding value(s).  Privileges required:  `UPDATE` `EDGE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EdgeUpdateEdgeAttributes**](EdgeUpdateEdgeAttributes.md)|  | 

### Return type

[**EdgeUpdateEdgeAttributesResult**](edge_update_edge_attributes_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeUpdateEdgeCredentialsByConfiguration

> EdgeUpdateEdgeCredentialsByConfigurationResult EdgeUpdateEdgeCredentialsByConfiguration(ctx, body)
Update Edge local UI credentials by configuration profile

Requests an update to the Edge-local UI authentication credentials for all Edges belonging to the specified configuration profile. On success, returns a JSON object containing a list of each of the action IDs queued.  Privileges required:  `UPDATE` `EDGE`  `UPDATE` `ENTERPRISE_KEYS`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EdgeUpdateEdgeCredentialsByConfiguration**](EdgeUpdateEdgeCredentialsByConfiguration.md)|  | 

### Return type

[**EdgeUpdateEdgeCredentialsByConfigurationResult**](edge_update_edge_credentials_by_configuration_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

