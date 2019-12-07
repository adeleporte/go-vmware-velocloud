# NetflowFilterRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Appid** | **int32** | Integer ID corresponding to an application in the network-level application map | [optional] 
**Classid** | **int32** | Integer ID corresponding to an application class in the network-level application map | [optional] 
**Dscp** | **int32** | Integer ID indicating DSCP classification, where mappings are as follows: [EF:46,VA:44,AF11:10,AF12:12,AF13:14,AF21:18,AF22:20,AF23:22,AF31:26,AF32:28,AF33:30,AF41:34,AF42:36,AF43:38,CS0:0,CS1:8,CS2:16,CS3:24,CS4:32,CS5:40,CS6:48,CS7:56] | [optional] 
**Sip** | **string** | Source IP address | [optional] 
**SportHigh** | **int32** | Upper bound of a source port range | [optional] 
**SportLow** | **int32** | Lower bound of a source port range | [optional] 
**SAddressGroup** | **string** | Source address group reference | [optional] 
**SPortGroup** | **string** | Source port group reference | [optional] 
**Ssm** | **string** | Source subnet mask, e.g. 255.255.255.0 | [optional] 
**Smac** | **string** | Source MAC address | [optional] 
**Svlan** | **int32** | Integer ID for the source VLAN | [optional] 
**OsVersion** | **int32** | Index corresponding to the OS in the array: [OTHER,WINDOWS,LINUX,MACOS,IOS,ANDROID,EDGE] | [optional] 
**Hostname** | **string** |  | [optional] 
**Dip** | **string** | Destination IP address | [optional] 
**DportLow** | **int32** | Lower bound of a destination port range | [optional] 
**DportHigh** | **int32** | Upper bound of a destination port range | [optional] 
**DAddressGroup** | **string** | Destination address group reference | [optional] 
**DPortGroup** | **string** | Destination port group reference | [optional] 
**Dsm** | **string** | Destination subnet mask e.g. 255.255.255.0 | [optional] 
**Dmac** | **string** | Destination MAC address | [optional] 
**Dvlan** | **int32** | Integer ID for the destination VLAN | [optional] 
**Proto** | **int32** | Integer ID corresponding to a protocol | [optional] 
**SRuleType** | **string** | Source rule type | [optional] 
**DRuleType** | **string** | Destination rule type | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


