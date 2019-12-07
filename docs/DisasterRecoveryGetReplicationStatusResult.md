# DisasterRecoveryGetReplicationStatusResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveAddress** | **string** |  | 
**ActiveReplicationAddress** | **string** |  | [optional] 
**ClientContact** | [**[]DisasterRecoveryClientContact**](disaster_recovery_client_contact.md) |  | [optional] 
**ClientCount** | [**DisasterRecoveryGetReplicationStatusResultClientCount**](disaster_recovery_get_replication_status_result_clientCount.md) |  | [optional] 
**DrState** | **string** |  | 
**DrVCOUser** | **string** |  | 
**FailureDescription** | **string** |  | 
**LastDrProtectedTime** | [**time.Time**](time.Time.md) |  | [optional] 
**Role** | **string** |  | 
**RoleTimestamp** | [**time.Time**](time.Time.md) |  | 
**StandbyList** | [**[]DisasterRecoveryGetReplicationStatusResultStandbyList**](disaster_recovery_get_replication_status_result_standbyList.md) |  | 
**StateHistory** | [**[]DisasterRecoveryGetReplicationStatusResultStateHistory**](disaster_recovery_get_replication_status_result_stateHistory.md) |  | [optional] 
**StateTimestamp** | [**time.Time**](time.Time.md) |  | [optional] 
**VcoIp** | **string** |  | 
**VcoReplicationIp** | **string** |  | [optional] 
**VcoUuid** | **string** |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


