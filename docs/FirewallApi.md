# \FirewallApi

All URIs are relative to *http://localhost/portal/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FirewallGetEnterpriseFirewallLogs**](FirewallApi.md#FirewallGetEnterpriseFirewallLogs) | **Post** /firewall/getEnterpriseFirewallLogs | Get enterprise firewall logs



## FirewallGetEnterpriseFirewallLogs

> FirewallGetEnterpriseFirewallLogsResult FirewallGetEnterpriseFirewallLogs(ctx, body)
Get enterprise firewall logs

Gets firewall logs for the specified enterprise.  Privileges required:  `READ` `EDGE`  `VIEW_FIREWALL_LOGS` `undefined`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**FirewallGetEnterpriseFirewallLogs**](FirewallGetEnterpriseFirewallLogs.md)|  | 

### Return type

[**FirewallGetEnterpriseFirewallLogsResult**](firewall_get_enterprise_firewall_logs_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

