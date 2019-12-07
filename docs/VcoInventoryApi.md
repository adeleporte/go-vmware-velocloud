# \VcoInventoryApi

All URIs are relative to *http://localhost/portal/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VcoInventoryAssociateEdge**](VcoInventoryApi.md#VcoInventoryAssociateEdge) | **Post** /vcoInventory/associateEdge | Assign Edge to enterprise
[**VcoInventoryGetInventoryItems**](VcoInventoryApi.md#VcoInventoryGetInventoryItems) | **Post** /vcoInventory/getInventoryItems | Get available VCO inventory items



## VcoInventoryAssociateEdge

> InlineResponse2003 VcoInventoryAssociateEdge(ctx, body)
Assign Edge to enterprise

Assigns an Edge in the inventory to the specified enterprise. To perform the action, the Edge should already be in a `STAGING` state. The assignment can be done at an enterprise level, without selecting a destination Edge profile. In such a case, the inventory Edge is assigned to a staging profile within the enterprise. Optionally, a profile or destination Edge can be assigned to this inventory Edge. The Edge in the inventory can be assigned to any profile. The inventory Edge can be assigned to an enterprise Edge only if it is in a `PENDING` or `REACTIVATION_PENDING` state.  Privileges required:  `CREATE` `ENTERPRISE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**VcoInventoryAssociateEdge**](VcoInventoryAssociateEdge.md)|  | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VcoInventoryGetInventoryItems

> []InventoryItem VcoInventoryGetInventoryItems(ctx, body)
Get available VCO inventory items

Gets all of the inventory information available with this VCO. This method does not have required parameters. The optional parameters are  `enterpriseId` - Returns inventory items belonging to the specified enterprise. If the caller context is an enterprise, then this value will be taken from the token itself. `modifiedSince` - Returns inventory items that have been modified in the last `modifiedSince` hours. `with` - Array containing the string \"edge\" to get details about the provisioned Edge, if any.  Privileges required:  `READ` `INVENTORY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**VcoInventoryGetInventoryItems**](VcoInventoryGetInventoryItems.md)|  | 

### Return type

[**[]InventoryItem**](inventory_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

