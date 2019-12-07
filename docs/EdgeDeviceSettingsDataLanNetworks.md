# EdgeDeviceSettingsDataLanNetworks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Space** | **string** |  | [optional] 
**Guest** | **bool** |  | [optional] 
**Secure** | **bool** |  | [optional] 
**Advertise** | **bool** |  | [optional] 
**PingResponse** | **bool** |  | [optional] [default to true]
**Cost** | **int32** |  | [optional] 
**Dhcp** | [**EdgeDeviceSettingsDataLanDhcp**](edgeDeviceSettingsData_lan_dhcp.md) |  | [optional] 
**StaticReserved** | **int32** |  | [optional] 
**Netmask** | **string** |  | [optional] 
**CidrPrefix** | **int32** |  | [optional] 
**CidrIp** | **string** |  | [optional] 
**BaseDhcpAddr** | **int32** | An offset from the cidrIp including staticReserved (if any) | [optional] 
**NumDhcpAddr** | **int32** |  | [optional] 
**Name** | **string** |  | [optional] 
**Interfaces** | **[]string** |  | [optional] 
**VlanId** | **int32** |  | [optional] 
**ManagementIp** | **string** |  | [optional] 
**Disabled** | **bool** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


