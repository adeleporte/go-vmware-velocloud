# NetworkGetNetworkGatewaysResultItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivationKey** | **string** |  | [optional] 
**ActivationState** | **string** |  | [optional] 
**ActivationTime** | [**time.Time**](time.Time.md) |  | [optional] 
**BuildNumber** | **string** |  | [optional] 
**Certificates** | [**[]GatewayCertificate**](gateway_certificate.md) |  | [optional] 
**ConnectedEdges** | **int32** |  | [optional] 
**ConnectedEdgeList** | [**[]GatewayPoolGatewayConnectedEdgeList**](gateway_pool_gateway_connectedEdgeList.md) |  | [optional] 
**Created** | [**time.Time**](time.Time.md) |  | [optional] 
**DataCenters** | [**[]map[string]interface{}**](map[string]interface{}.md) |  | [optional] 
**Description** | **string** |  | [optional] 
**DeviceId** | **string** |  | [optional] 
**DnsName** | **string** |  | [optional] 
**EndpointPkiMode** | **string** |  | [optional] 
**EnterpriseAssociations** | [**[]GatewayEnterpriseAssoc**](gateway_enterprise_assoc.md) |  | [optional] 
**EnterpriseAssociationCount** | **map[string]int32** |  | [optional] 
**Enterprises** | [**[]Enterprise**](enterprise.md) |  | [optional] 
**EnterpriseProxyId** | **int32** |  | [optional] 
**GatewayState** | **string** |  | [optional] 
**HandOffDetail** | [**GatewayHandoffDetail**](gateway_handoff_detail.md) |  | [optional] 
**HandOffEdges** | [**[]GatewayHandoffEdge**](gateway_handoff_edge.md) |  | [optional] 
**Id** | **int32** |  | [optional] 
**IpAddress** | **string** |  | [optional] 
**IpsecGatewayDetail** | [**GatewayPoolGatewayIpsecGatewayDetail**](gateway_pool_gateway_ipsecGatewayDetail.md) |  | [optional] 
**IsLoadBalanced** | **bool** |  | [optional] 
**LastContact** | **string** |  | [optional] 
**LogicalId** | **string** |  | [optional] 
**Modified** | [**time.Time**](time.Time.md) |  | [optional] 
**Name** | **string** |  | [optional] 
**NetworkId** | **int32** |  | [optional] 
**Pools** | [**[]GatewayGatewayPool**](gateway_gateway_pool.md) |  | [optional] 
**PrivateIpAddress** | **string** |  | [optional] 
**Roles** | [**[]GatewayRole**](gateway_role.md) |  | [optional] 
**ServiceState** | **string** |  | [optional] 
**ServiceUpSince** | **string** |  | [optional] 
**Site** | [**GatewaySite**](gateway_site.md) |  | [optional] 
**SiteId** | **int32** |  | [optional] 
**SoftwareVersion** | **string** |  | [optional] 
**SystemUpSince** | **string** |  | [optional] 
**Utilization** | **float32** |  | [optional] 
**UtilizationDetail** | [**GatewayPoolGatewayUtilizationDetail**](gateway_pool_gateway_utilizationDetail.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


