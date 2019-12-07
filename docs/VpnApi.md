# \VpnApi

All URIs are relative to *http://localhost/portal/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VpnGenerateVpnGatewayConfiguration**](VpnApi.md#VpnGenerateVpnGatewayConfiguration) | **Post** /vpn/generateVpnGatewayConfiguration | Provision a non-SD-WAN VPN site



## VpnGenerateVpnGatewayConfiguration

> VpnGenerateVpnGatewayConfigurationResult VpnGenerateVpnGatewayConfiguration(ctx, body)
Provision a non-SD-WAN VPN site

Provision a non-SD-WAN site (e.g. a data center or cloud service PoP) and generate VPN configuration.  Privileges required:  `CREATE` `NETWORK_SERVICE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**VpnGenerateVpnGatewayConfiguration**](VpnGenerateVpnGatewayConfiguration.md)|  | 

### Return type

[**VpnGenerateVpnGatewayConfigurationResult**](vpn_generate_vpn_gateway_configuration_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

