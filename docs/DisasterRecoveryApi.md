# \DisasterRecoveryApi

All URIs are relative to *http://localhost/portal/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisasterRecoveryConfigureActiveForReplication**](DisasterRecoveryApi.md#DisasterRecoveryConfigureActiveForReplication) | **Post** /disasterRecovery/configureActiveForReplication | Configure current VCO for disaster recovery
[**DisasterRecoveryDemoteActive**](DisasterRecoveryApi.md#DisasterRecoveryDemoteActive) | **Post** /disasterRecovery/demoteActive | Demote current VCO from active to zombie
[**DisasterRecoveryGetReplicationBlob**](DisasterRecoveryApi.md#DisasterRecoveryGetReplicationBlob) | **Post** /disasterRecovery/getReplicationBlob | Get blob needed to configure VCO replication on standby
[**DisasterRecoveryGetReplicationStatus**](DisasterRecoveryApi.md#DisasterRecoveryGetReplicationStatus) | **Post** /disasterRecovery/getReplicationStatus | Get disaster recovery status
[**DisasterRecoveryPrepareForStandby**](DisasterRecoveryApi.md#DisasterRecoveryPrepareForStandby) | **Post** /disasterRecovery/prepareForStandby | Designate current VCO as DR standby candidate
[**DisasterRecoveryPromoteStandbyToActive**](DisasterRecoveryApi.md#DisasterRecoveryPromoteStandbyToActive) | **Post** /disasterRecovery/promoteStandbyToActive | Promote standby VCO to active
[**DisasterRecoveryRemoveStandby**](DisasterRecoveryApi.md#DisasterRecoveryRemoveStandby) | **Post** /disasterRecovery/removeStandby | Remove VCO disaster recovery on current server
[**DisasterRecoveryTransitionToStandby**](DisasterRecoveryApi.md#DisasterRecoveryTransitionToStandby) | **Post** /disasterRecovery/transitionToStandby | Transition VCO to standby



## DisasterRecoveryConfigureActiveForReplication

> DisasterRecoveryConfigureActiveForReplicationResult DisasterRecoveryConfigureActiveForReplication(ctx, body)
Configure current VCO for disaster recovery

Configures the current Orchestrator to be active and the specified Orchestrator to be standby for Orchestrator disaster recovery replication. Required attributes: 1) `standbyList`, a single-entry array containing the `standbyAddress` and `standbyUuid`, 2) `drVCOUser`, a Orchestrator super user available on both the active and standby VCOs, and 3) `drVCOPassword`, the password of `drVCOUser` on the standby Orchestrator (unless the `autoConfigStandby` option is specified as `false`). The call sets up the active Orchestrator to allow replication from the standby and then (unless `autoConfigStandby` is `false`) makes a `transitionToStandby` API call to the specified standby, expected to have been previously placed in `STANDBY_CANDIDATE` state via `prepareForStandby`. After this call, the active and standby VCOs should be polled via `getReplicationStatus` until they  both reach `STANDBY_RUNNING` `drState` (or a configuration error is reported).  (Note: `drVCOPassword` is not persisted.)  Privileges required:  `CREATE` `REPLICATION`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**DisasterRecoveryConfigureActiveForReplication**](DisasterRecoveryConfigureActiveForReplication.md)|  | 

### Return type

[**DisasterRecoveryConfigureActiveForReplicationResult**](disaster_recovery_configure_active_for_replication_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisasterRecoveryDemoteActive

> DisasterRecoveryDemoteActiveResult DisasterRecoveryDemoteActive(ctx, body)
Demote current VCO from active to zombie

Demotes the current VCO from active to zombie. No input parameters are required. The active server is expected to be in the `drState` `FAILURE_GET_STANDBY_STATUS` or `FAILURE_MYSQL_ACTIVE_STATUS`, meaning that DR protection had been engaged (with the last successful replication status observed at `lastDRProtectedTime`) but that active failed a health check since that time.  If the active server is in the `drState` `STANDBY_RUNNING`, meaning that it has detected no problems in interacting with the standby server, the operator can force demotion of the active using the optional `force` parameter passed with value of `true`; in this case, the operator must ensure that the standby server has already been successfully promoted.  Privileges required:  `CREATE` `REPLICATION`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**DisasterRecoveryDemoteActive**](DisasterRecoveryDemoteActive.md)|  | 

### Return type

[**DisasterRecoveryDemoteActiveResult**](disaster_recovery_demote_active_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisasterRecoveryGetReplicationBlob

> DisasterRecoveryGetReplicationBlobResult DisasterRecoveryGetReplicationBlob(ctx, body)
Get blob needed to configure VCO replication on standby

Gets from the active Orchestrator the blob needed to configure replication on the standby. Used only when `configureActiveForReplication` was called with `autoConfigStandby` set to `false` [`true` by default].  Privileges required:  `CREATE` `REPLICATION`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**DisasterRecoveryGetReplicationBlob**](DisasterRecoveryGetReplicationBlob.md)|  | 

### Return type

[**DisasterRecoveryGetReplicationBlobResult**](disaster_recovery_get_replication_blob_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisasterRecoveryGetReplicationStatus

> DisasterRecoveryGetReplicationStatusResult DisasterRecoveryGetReplicationStatus(ctx, body)
Get disaster recovery status

Gets the disaster recovery replication status, optionally with client contact, state transition history, and storage information. No input parameters are required.  Can optionally specify one or more of the following `with` parameters: `clientContact`, `clientCount`, `stateHistory`, and `storageInfo`.  Privileges required:  `READ` `REPLICATION`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**DisasterRecoveryGetReplicationStatus**](DisasterRecoveryGetReplicationStatus.md)|  | 

### Return type

[**DisasterRecoveryGetReplicationStatusResult**](disaster_recovery_get_replication_status_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisasterRecoveryPrepareForStandby

> DisasterRecoveryPrepareForStandbyResult DisasterRecoveryPrepareForStandby(ctx, body)
Designate current VCO as DR standby candidate

Transitions the current Orchestrator to a quiesced state, ready to be configured as a standby system. No input parameters are required.  After this call, the Orchestrator will be restarted in standby mode. The caller should subsequently poll `getReplicationStatus` until `drState` is `STANDBY_CANDIDATE`.  This is the first step in configuring Orchestrator disaster recovery.  Privileges required:  `CREATE` `REPLICATION`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**DisasterRecoveryPrepareForStandby**](DisasterRecoveryPrepareForStandby.md)|  | 

### Return type

[**DisasterRecoveryPrepareForStandbyResult**](disaster_recovery_prepare_for_standby_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisasterRecoveryPromoteStandbyToActive

> DisasterRecoveryPromoteStandbyToActiveResult DisasterRecoveryPromoteStandbyToActive(ctx, body)
Promote standby VCO to active

Promotes the current server to take over as the single standalone VCO. The current server is expected to be a standby in the `drState` `FAILURE_MYSQL_STANDBY_STATUS`, meaning that DR protection had been engaged (with the last successful replication status observed at `lastDRProtectedTime`), but that standby has been unable to replicate since that time.   If the standby server is in the `drState` `STANDBY_RUNNING`, meaning that it has detected no problems in replicating from the active server, the operator can force promotion of the standby using the optional `force` parameter passed with a value of `true`. In this case, the standby server will call `demoteActive/force` on the active server.  The operator should, if possible, ensure that the formerly active server is demoted by running `demoteServer` directly on that server if the standby server was unable to do so successfully.  Privileges required:  `CREATE` `REPLICATION`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**DisasterRecoveryPromoteStandbyToActive**](DisasterRecoveryPromoteStandbyToActive.md)|  | 

### Return type

[**DisasterRecoveryPromoteStandbyToActiveResult**](disaster_recovery_promote_standby_to_active_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisasterRecoveryRemoveStandby

> DisasterRecoveryRemoveStandbyResult DisasterRecoveryRemoveStandby(ctx, body)
Remove VCO disaster recovery on current server

Removes disaster recovery on the current server. In addition, makes a best-effort call to `removeStandby` on the paired disaster recovery server. No input parameters are required.  Privileges required:  `CREATE` `REPLICATION`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**DisasterRecoveryRemoveStandby**](DisasterRecoveryRemoveStandby.md)|  | 

### Return type

[**DisasterRecoveryRemoveStandbyResult**](disaster_recovery_remove_standby_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisasterRecoveryTransitionToStandby

> DisasterRecoveryTransitionToStandbyResult DisasterRecoveryTransitionToStandby(ctx, body)
Transition VCO to standby

Configures the current Orchestrator to transition to standby in a disaster recovery active/standby pair. Requires the `activeAccessFromStandby` parameter that contains the data needed to configure standby. This data is produced by `configureActiveForReplication`, which by default automatically calls `transitionToStandby`; an explicit call is needed (with a blob obtained from `getReplicationBlob`), only if `configureActiveForReplication` is called with `autoConfigStandby` set to `false`.  Privileges required:  `CREATE` `REPLICATION`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**DisasterRecoveryTransitionToStandby**](DisasterRecoveryTransitionToStandby.md)|  | 

### Return type

[**DisasterRecoveryTransitionToStandbyResult**](disaster_recovery_transition_to_standby_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

