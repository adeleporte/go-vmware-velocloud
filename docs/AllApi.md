# \AllApi

All URIs are relative to *http://localhost/portal/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfigurationCloneAndConvertConfiguration**](AllApi.md#ConfigurationCloneAndConvertConfiguration) | **Post** /configuration/cloneAndConvertConfiguration | Create new segment-based profile from existing network-based profile
[**ConfigurationCloneConfiguration**](AllApi.md#ConfigurationCloneConfiguration) | **Post** /configuration/cloneConfiguration | Clone configuration profile
[**ConfigurationCloneEnterpriseTemplate**](AllApi.md#ConfigurationCloneEnterpriseTemplate) | **Post** /configuration/cloneEnterpriseTemplate | Clone default enterprise configuration profile
[**ConfigurationDeleteConfiguration**](AllApi.md#ConfigurationDeleteConfiguration) | **Post** /configuration/deleteConfiguration | Delete configuration profile
[**ConfigurationGetConfiguration**](AllApi.md#ConfigurationGetConfiguration) | **Post** /configuration/getConfiguration | Get configuration profile
[**ConfigurationGetConfigurationModules**](AllApi.md#ConfigurationGetConfigurationModules) | **Post** /configuration/getConfigurationModules | Get configuration profile modules
[**ConfigurationGetIdentifiableApplications**](AllApi.md#ConfigurationGetIdentifiableApplications) | **Post** /configuration/getIdentifiableApplications | Get the applications that are identifiable through DPI.
[**ConfigurationGetRoutableApplications**](AllApi.md#ConfigurationGetRoutableApplications) | **Post** /configuration/getRoutableApplications | Get first packet routable applications
[**ConfigurationInsertConfigurationModule**](AllApi.md#ConfigurationInsertConfigurationModule) | **Post** /configuration/insertConfigurationModule | Create configuration module
[**ConfigurationUpdateConfiguration**](AllApi.md#ConfigurationUpdateConfiguration) | **Post** /configuration/updateConfiguration | Update configuration profile
[**ConfigurationUpdateConfigurationModule**](AllApi.md#ConfigurationUpdateConfigurationModule) | **Post** /configuration/updateConfigurationModule | Update configuration module
[**DisasterRecoveryConfigureActiveForReplication**](AllApi.md#DisasterRecoveryConfigureActiveForReplication) | **Post** /disasterRecovery/configureActiveForReplication | Configure current VCO for disaster recovery
[**DisasterRecoveryDemoteActive**](AllApi.md#DisasterRecoveryDemoteActive) | **Post** /disasterRecovery/demoteActive | Demote current VCO from active to zombie
[**DisasterRecoveryGetReplicationBlob**](AllApi.md#DisasterRecoveryGetReplicationBlob) | **Post** /disasterRecovery/getReplicationBlob | Get blob needed to configure VCO replication on standby
[**DisasterRecoveryGetReplicationStatus**](AllApi.md#DisasterRecoveryGetReplicationStatus) | **Post** /disasterRecovery/getReplicationStatus | Get disaster recovery status
[**DisasterRecoveryPrepareForStandby**](AllApi.md#DisasterRecoveryPrepareForStandby) | **Post** /disasterRecovery/prepareForStandby | Designate current VCO as DR standby candidate
[**DisasterRecoveryPromoteStandbyToActive**](AllApi.md#DisasterRecoveryPromoteStandbyToActive) | **Post** /disasterRecovery/promoteStandbyToActive | Promote standby VCO to active
[**DisasterRecoveryRemoveStandby**](AllApi.md#DisasterRecoveryRemoveStandby) | **Post** /disasterRecovery/removeStandby | Remove VCO disaster recovery on current server
[**DisasterRecoveryTransitionToStandby**](AllApi.md#DisasterRecoveryTransitionToStandby) | **Post** /disasterRecovery/transitionToStandby | Transition VCO to standby
[**EdgeDeleteEdge**](AllApi.md#EdgeDeleteEdge) | **Post** /edge/deleteEdge | Delete Edge
[**EdgeDeleteEdgeBgpNeighborRecords**](AllApi.md#EdgeDeleteEdgeBgpNeighborRecords) | **Post** /edge/deleteEdgeBgpNeighborRecords | Delete Edge BGP neighbor records
[**EdgeEdgeCancelReactivation**](AllApi.md#EdgeEdgeCancelReactivation) | **Post** /edge/edgeCancelReactivation | Cancel pending Edge reactivation request
[**EdgeEdgeProvision**](AllApi.md#EdgeEdgeProvision) | **Post** /edge/edgeProvision | Provision Edge
[**EdgeEdgeRequestReactivation**](AllApi.md#EdgeEdgeRequestReactivation) | **Post** /edge/edgeRequestReactivation | Prepare Edge for reactivation
[**EdgeGetClientVisibilityMode**](AllApi.md#EdgeGetClientVisibilityMode) | **Post** /edge/getClientVisibilityMode | Get an edge&#39;s client visibility mode
[**EdgeGetEdge**](AllApi.md#EdgeGetEdge) | **Post** /edge/getEdge | Get Edge
[**EdgeGetEdgeConfigurationModules**](AllApi.md#EdgeGetEdgeConfigurationModules) | **Post** /edge/getEdgeConfigurationModules | Get edge configuration modules
[**EdgeGetEdgeConfigurationStack**](AllApi.md#EdgeGetEdgeConfigurationStack) | **Post** /edge/getEdgeConfigurationStack | Get Edge configuration stack
[**EdgeSetEdgeEnterpriseConfiguration**](AllApi.md#EdgeSetEdgeEnterpriseConfiguration) | **Post** /edge/setEdgeEnterpriseConfiguration | Apply configuration profile to Edge
[**EdgeSetEdgeHandOffGateways**](AllApi.md#EdgeSetEdgeHandOffGateways) | **Post** /edge/setEdgeHandOffGateways | Set Edge on premise hand-off gateway(s)
[**EdgeSetEdgeOperatorConfiguration**](AllApi.md#EdgeSetEdgeOperatorConfiguration) | **Post** /edge/setEdgeOperatorConfiguration | Apply operator profile to Edge
[**EdgeUpdateEdgeAdminPassword**](AllApi.md#EdgeUpdateEdgeAdminPassword) | **Post** /edge/updateEdgeAdminPassword | Update Edge local UI authentication credentials
[**EdgeUpdateEdgeAttributes**](AllApi.md#EdgeUpdateEdgeAttributes) | **Post** /edge/updateEdgeAttributes | Update Edge attributes
[**EdgeUpdateEdgeCredentialsByConfiguration**](AllApi.md#EdgeUpdateEdgeCredentialsByConfiguration) | **Post** /edge/updateEdgeCredentialsByConfiguration | Update Edge local UI credentials by configuration profile
[**EnterpriseDecodeEnterpriseKey**](AllApi.md#EnterpriseDecodeEnterpriseKey) | **Post** /enterprise/decodeEnterpriseKey | Decrypt an enterprise key
[**EnterpriseDeleteEnterprise**](AllApi.md#EnterpriseDeleteEnterprise) | **Post** /enterprise/deleteEnterprise | Delete enterprise
[**EnterpriseDeleteEnterpriseGatewayRecords**](AllApi.md#EnterpriseDeleteEnterpriseGatewayRecords) | **Post** /enterprise/deleteEnterpriseGatewayRecords | Delete gateway BGP neighbor record(s) for enterprise
[**EnterpriseDeleteEnterpriseNetworkAllocation**](AllApi.md#EnterpriseDeleteEnterpriseNetworkAllocation) | **Post** /enterprise/deleteEnterpriseNetworkAllocation | Delete enterprise network allocation
[**EnterpriseDeleteEnterpriseNetworkSegment**](AllApi.md#EnterpriseDeleteEnterpriseNetworkSegment) | **Post** /enterprise/deleteEnterpriseNetworkSegment | Delete an enterprise network segment
[**EnterpriseDeleteEnterpriseService**](AllApi.md#EnterpriseDeleteEnterpriseService) | **Post** /enterprise/deleteEnterpriseService | Delete enterprise network service
[**EnterpriseEncodeEnterpriseKey**](AllApi.md#EnterpriseEncodeEnterpriseKey) | **Post** /enterprise/encodeEnterpriseKey | Encrypt an enterprise key
[**EnterpriseGetEnterprise**](AllApi.md#EnterpriseGetEnterprise) | **Post** /enterprise/getEnterprise | Get enterprise
[**EnterpriseGetEnterpriseAddresses**](AllApi.md#EnterpriseGetEnterpriseAddresses) | **Post** /enterprise/getEnterpriseAddresses | Get enterprise IP address(es)
[**EnterpriseGetEnterpriseAlertConfigurations**](AllApi.md#EnterpriseGetEnterpriseAlertConfigurations) | **Post** /enterprise/getEnterpriseAlertConfigurations | Get enterprise alert configuration(s)
[**EnterpriseGetEnterpriseAlerts**](AllApi.md#EnterpriseGetEnterpriseAlerts) | **Post** /enterprise/getEnterpriseAlerts | Get triggered enterprise alerts
[**EnterpriseGetEnterpriseAllAlertRecipients**](AllApi.md#EnterpriseGetEnterpriseAllAlertRecipients) | **Post** /enterprise/getEnterpriseAllAlertsRecipients | Get recipients receiving all enterprise alerts
[**EnterpriseGetEnterpriseCapabilities**](AllApi.md#EnterpriseGetEnterpriseCapabilities) | **Post** /enterprise/getEnterpriseCapabilities | Get enterprise capabilities
[**EnterpriseGetEnterpriseConfigurations**](AllApi.md#EnterpriseGetEnterpriseConfigurations) | **Post** /enterprise/getEnterpriseConfigurations | Get enterprise configuration profiles
[**EnterpriseGetEnterpriseEdges**](AllApi.md#EnterpriseGetEnterpriseEdges) | **Post** /enterprise/getEnterpriseEdges | Get enterprise Edges
[**EnterpriseGetEnterpriseGatewayHandoff**](AllApi.md#EnterpriseGetEnterpriseGatewayHandoff) | **Post** /enterprise/getEnterpriseGatewayHandoff | Get enterprise gateway handoff configuration
[**EnterpriseGetEnterpriseMaximumSegments**](AllApi.md#EnterpriseGetEnterpriseMaximumSegments) | **Post** /enterprise/getEnterpriseMaximumSegments | Get enterprise maximum segment
[**EnterpriseGetEnterpriseNetworkAllocation**](AllApi.md#EnterpriseGetEnterpriseNetworkAllocation) | **Post** /enterprise/getEnterpriseNetworkAllocation | Get enterprise network allocation
[**EnterpriseGetEnterpriseNetworkAllocations**](AllApi.md#EnterpriseGetEnterpriseNetworkAllocations) | **Post** /enterprise/getEnterpriseNetworkAllocations | Get enterprise network allocations
[**EnterpriseGetEnterpriseNetworkSegments**](AllApi.md#EnterpriseGetEnterpriseNetworkSegments) | **Post** /enterprise/getEnterpriseNetworkSegments | Get all network segment objects defined on an enterprise
[**EnterpriseGetEnterpriseProperty**](AllApi.md#EnterpriseGetEnterpriseProperty) | **Post** /enterprise/getEnterpriseProperty | Get enterprise property
[**EnterpriseGetEnterpriseRouteConfiguration**](AllApi.md#EnterpriseGetEnterpriseRouteConfiguration) | **Post** /enterprise/getEnterpriseRouteConfiguration | Get enterprise global routing preferences
[**EnterpriseGetEnterpriseRouteTable**](AllApi.md#EnterpriseGetEnterpriseRouteTable) | **Post** /enterprise/getEnterpriseRouteTable | Get enterprise route table
[**EnterpriseGetEnterpriseServices**](AllApi.md#EnterpriseGetEnterpriseServices) | **Post** /enterprise/getEnterpriseServices | Get enterprise network services
[**EnterpriseGetEnterpriseUsers**](AllApi.md#EnterpriseGetEnterpriseUsers) | **Post** /enterprise/getEnterpriseUsers | Get enterprise users
[**EnterpriseInsertEnterprise**](AllApi.md#EnterpriseInsertEnterprise) | **Post** /enterprise/insertEnterprise | Create enterprise
[**EnterpriseInsertEnterpriseNetworkAllocation**](AllApi.md#EnterpriseInsertEnterpriseNetworkAllocation) | **Post** /enterprise/insertEnterpriseNetworkAllocation | Create enterprise network allocation
[**EnterpriseInsertEnterpriseNetworkSegment**](AllApi.md#EnterpriseInsertEnterpriseNetworkSegment) | **Post** /enterprise/insertEnterpriseNetworkSegment | Create enterprise network segment
[**EnterpriseInsertEnterpriseService**](AllApi.md#EnterpriseInsertEnterpriseService) | **Post** /enterprise/insertEnterpriseService | Create enterprise service
[**EnterpriseInsertEnterpriseUser**](AllApi.md#EnterpriseInsertEnterpriseUser) | **Post** /enterprise/insertEnterpriseUser | Create enterprise user
[**EnterpriseInsertOrUpdateEnterpriseAlertConfigurations**](AllApi.md#EnterpriseInsertOrUpdateEnterpriseAlertConfigurations) | **Post** /enterprise/insertOrUpdateEnterpriseAlertConfigurations | Create, update, or delete enterprise alert configurations
[**EnterpriseInsertOrUpdateEnterpriseCapability**](AllApi.md#EnterpriseInsertOrUpdateEnterpriseCapability) | **Post** /enterprise/insertOrUpdateEnterpriseCapability | Create or update enterprise capability
[**EnterpriseInsertOrUpdateEnterpriseGatewayHandoff**](AllApi.md#EnterpriseInsertOrUpdateEnterpriseGatewayHandoff) | **Post** /enterprise/insertOrUpdateEnterpriseGatewayHandoff | Create or update enterprise gateway handoff configuration
[**EnterpriseInsertOrUpdateEnterpriseProperty**](AllApi.md#EnterpriseInsertOrUpdateEnterpriseProperty) | **Post** /enterprise/insertOrUpdateEnterpriseProperty | Insert or update an enterprise property
[**EnterpriseProxyDeleteEnterpriseProxyUser**](AllApi.md#EnterpriseProxyDeleteEnterpriseProxyUser) | **Post** /enterpriseProxy/deleteEnterpriseProxyUser | Delete partner admin user
[**EnterpriseProxyGetEnterpriseProxyEdgeInventory**](AllApi.md#EnterpriseProxyGetEnterpriseProxyEdgeInventory) | **Post** /enterpriseProxy/getEnterpriseProxyEdgeInventory | Get partner enterprises and associated Edge inventory
[**EnterpriseProxyGetEnterpriseProxyEnterprises**](AllApi.md#EnterpriseProxyGetEnterpriseProxyEnterprises) | **Post** /enterpriseProxy/getEnterpriseProxyEnterprises | Get partner enterprises
[**EnterpriseProxyGetEnterpriseProxyGatewayPools**](AllApi.md#EnterpriseProxyGetEnterpriseProxyGatewayPools) | **Post** /enterpriseProxy/getEnterpriseProxyGatewayPools | Get gateway pools
[**EnterpriseProxyGetEnterpriseProxyGateways**](AllApi.md#EnterpriseProxyGetEnterpriseProxyGateways) | **Post** /enterpriseProxy/getEnterpriseProxyGateways | Get gateways for enterprise proxy
[**EnterpriseProxyGetEnterpriseProxyOperatorProfiles**](AllApi.md#EnterpriseProxyGetEnterpriseProxyOperatorProfiles) | **Post** /enterpriseProxy/getEnterpriseProxyOperatorProfiles | Get partner operator profiles
[**EnterpriseProxyGetEnterpriseProxyProperty**](AllApi.md#EnterpriseProxyGetEnterpriseProxyProperty) | **Post** /enterprise/getEnterpriseProxyProperty | Get enterprise proxy property
[**EnterpriseProxyGetEnterpriseProxyUser**](AllApi.md#EnterpriseProxyGetEnterpriseProxyUser) | **Post** /enterpriseProxy/getEnterpriseProxyUser | Get enterprise proxy user
[**EnterpriseProxyGetEnterpriseProxyUsers**](AllApi.md#EnterpriseProxyGetEnterpriseProxyUsers) | **Post** /enterpriseProxy/getEnterpriseProxyUsers | Get enterprise proxy admin users
[**EnterpriseProxyInsertEnterpriseProxyEnterprise**](AllApi.md#EnterpriseProxyInsertEnterpriseProxyEnterprise) | **Post** /enterpriseProxy/insertEnterpriseProxyEnterprise | Create enterprise owned by MSP
[**EnterpriseProxyInsertEnterpriseProxyUser**](AllApi.md#EnterpriseProxyInsertEnterpriseProxyUser) | **Post** /enterpriseProxy/insertEnterpriseProxyUser | Create partner admin user
[**EnterpriseProxyInsertOrUpdateEnterpriseProxyProperty**](AllApi.md#EnterpriseProxyInsertOrUpdateEnterpriseProxyProperty) | **Post** /enterpriseProxy/insertOrUpdateEnterpriseProxyProperty | Insert or update an enterprise proxy (MSP) property
[**EnterpriseProxyUpdateEnterpriseProxyUser**](AllApi.md#EnterpriseProxyUpdateEnterpriseProxyUser) | **Post** /enterpriseProxy/updateEnterpriseProxyUser | Update enterprise proxy admin user
[**EnterpriseSetEnterpriseAllAlertRecipients**](AllApi.md#EnterpriseSetEnterpriseAllAlertRecipients) | **Post** /enterprise/setEnterpriseAllAlertsRecipients | Specify recipients for all enterprise alerts
[**EnterpriseSetEnterpriseMaximumSegments**](AllApi.md#EnterpriseSetEnterpriseMaximumSegments) | **Post** /enterprise/setEnterpriseMaximumSegments | Set enterprise maximum segments
[**EnterpriseSetEnterpriseOperatorConfiguration**](AllApi.md#EnterpriseSetEnterpriseOperatorConfiguration) | **Post** /enterprise/setEnterpriseOperatorConfiguration | Apply operator profile to given enterprise
[**EnterpriseUpdateEnterprise**](AllApi.md#EnterpriseUpdateEnterprise) | **Post** /enterprise/updateEnterprise | Update enterprise
[**EnterpriseUpdateEnterpriseNetworkAllocation**](AllApi.md#EnterpriseUpdateEnterpriseNetworkAllocation) | **Post** /enterprise/updateEnterpriseNetworkAllocation | Update enterprise network allocation
[**EnterpriseUpdateEnterpriseNetworkSegment**](AllApi.md#EnterpriseUpdateEnterpriseNetworkSegment) | **Post** /enterprise/updateEnterpriseNetworkSegment | Update an enterprise network segment
[**EnterpriseUpdateEnterpriseRoute**](AllApi.md#EnterpriseUpdateEnterpriseRoute) | **Post** /enterprise/updateEnterpriseRoute | Update preferred VPN exits for a route
[**EnterpriseUpdateEnterpriseRouteConfiguration**](AllApi.md#EnterpriseUpdateEnterpriseRouteConfiguration) | **Post** /enterprise/updateEnterpriseRouteConfiguration | Update enterprise global routing preferences
[**EnterpriseUpdateEnterpriseSecurityPolicy**](AllApi.md#EnterpriseUpdateEnterpriseSecurityPolicy) | **Post** /enterprise/updateEnterpriseSecurityPolicy | Update enterprise security policy
[**EnterpriseUpdateEnterpriseService**](AllApi.md#EnterpriseUpdateEnterpriseService) | **Post** /enterprise/updateEnterpriseService | Update enterprise network service
[**EnterpriseUserDeleteEnterpriseUser**](AllApi.md#EnterpriseUserDeleteEnterpriseUser) | **Post** /enterpriseUser/deleteEnterpriseUser | Delete enterprise user
[**EnterpriseUserGetEnterpriseUser**](AllApi.md#EnterpriseUserGetEnterpriseUser) | **Post** /enterpriseUser/getEnterpriseUser | Get enterprise user
[**EnterpriseUserUpdateEnterpriseUser**](AllApi.md#EnterpriseUserUpdateEnterpriseUser) | **Post** /enterpriseUser/updateEnterpriseUser | Update enterprise user
[**EventGetEnterpriseEvents**](AllApi.md#EventGetEnterpriseEvents) | **Post** /event/getEnterpriseEvents | Get Edge events
[**EventGetOperatorEvents**](AllApi.md#EventGetOperatorEvents) | **Post** /event/getOperatorEvents | Get operator events
[**EventGetProxyEvents**](AllApi.md#EventGetProxyEvents) | **Post** /event/getProxyEvents | Fetch Partner events
[**FirewallGetEnterpriseFirewallLogs**](AllApi.md#FirewallGetEnterpriseFirewallLogs) | **Post** /firewall/getEnterpriseFirewallLogs | Get enterprise firewall logs
[**GatewayDeleteGateway**](AllApi.md#GatewayDeleteGateway) | **Post** /gateway/deleteGateway | Delete a gateway
[**GatewayGatewayProvision**](AllApi.md#GatewayGatewayProvision) | **Post** /gateway/gatewayProvision | Provision gateway
[**GatewayGetGatewayEdgeAssignments**](AllApi.md#GatewayGetGatewayEdgeAssignments) | **Post** /gateway/getGatewayEdgeAssignments | Get edge assignments for a gateway
[**GatewayUpdateGatewayAttributes**](AllApi.md#GatewayUpdateGatewayAttributes) | **Post** /gateway/updateGatewayAttributes | Update gateway attributes
[**LinkQualityEventGetLinkQualityEvents**](AllApi.md#LinkQualityEventGetLinkQualityEvents) | **Post** /linkQualityEvent/getLinkQualityEvents | Get link quality events
[**LoginEnterpriseLogin**](AllApi.md#LoginEnterpriseLogin) | **Post** /login/enterpriseLogin | Authenticate enterprise or partner (MSP) user
[**LoginOperatorLogin**](AllApi.md#LoginOperatorLogin) | **Post** /login/operatorLogin | Authenticate operator user
[**Logout**](AllApi.md#Logout) | **Post** /logout | Logout and invalidate authorization session cookie
[**Meta**](AllApi.md#Meta) | **Post** /meta/{apiPath} | Get Swagger specification for any VCO API call
[**MetricsGetEdgeAppLinkMetrics**](AllApi.md#MetricsGetEdgeAppLinkMetrics) | **Post** /metrics/getEdgeAppLinkMetrics | Get flow metric aggregate data
[**MetricsGetEdgeAppLinkSeries**](AllApi.md#MetricsGetEdgeAppLinkSeries) | **Post** /metrics/getEdgeAppLinkSeries | Get flow metric time series data
[**MetricsGetEdgeAppMetrics**](AllApi.md#MetricsGetEdgeAppMetrics) | **Post** /metrics/getEdgeAppMetrics | Get flow metric aggregate data by application
[**MetricsGetEdgeAppSeries**](AllApi.md#MetricsGetEdgeAppSeries) | **Post** /metrics/getEdgeAppSeries | Get flow metric time series data by application
[**MetricsGetEdgeCategoryMetrics**](AllApi.md#MetricsGetEdgeCategoryMetrics) | **Post** /metrics/getEdgeCategoryMetrics | Get flow metric aggregate data by application category
[**MetricsGetEdgeCategorySeries**](AllApi.md#MetricsGetEdgeCategorySeries) | **Post** /metrics/getEdgeCategorySeries | Get flow metric time series data by application category
[**MetricsGetEdgeDestMetrics**](AllApi.md#MetricsGetEdgeDestMetrics) | **Post** /metrics/getEdgeDestMetrics | Get flow metric aggregate data by destination
[**MetricsGetEdgeDestSeries**](AllApi.md#MetricsGetEdgeDestSeries) | **Post** /metrics/getEdgeDestSeries | Get flow metric time series data by destination
[**MetricsGetEdgeDeviceMetrics**](AllApi.md#MetricsGetEdgeDeviceMetrics) | **Post** /metrics/getEdgeDeviceMetrics | Get flow metric aggregate data by client device
[**MetricsGetEdgeDeviceSeries**](AllApi.md#MetricsGetEdgeDeviceSeries) | **Post** /metrics/getEdgeDeviceSeries | Get flow metric time series data by client device
[**MetricsGetEdgeLinkMetrics**](AllApi.md#MetricsGetEdgeLinkMetrics) | **Post** /metrics/getEdgeLinkMetrics | Get advanced flow metric aggregate data by link
[**MetricsGetEdgeLinkSeries**](AllApi.md#MetricsGetEdgeLinkSeries) | **Post** /metrics/getEdgeLinkSeries | Get advanced flow metric time series data by link
[**MetricsGetEdgeOsMetrics**](AllApi.md#MetricsGetEdgeOsMetrics) | **Post** /metrics/getEdgeOsMetrics | Get flow metric aggregate data by client OS
[**MetricsGetEdgeOsSeries**](AllApi.md#MetricsGetEdgeOsSeries) | **Post** /metrics/getEdgeOsSeries | Get flow metric time series data by client OS
[**MetricsGetEdgeSegmentMetrics**](AllApi.md#MetricsGetEdgeSegmentMetrics) | **Post** /metrics/getEdgeSegmentMetrics | Get flow metric aggregate data by segment Id
[**MetricsGetEdgeSegmentSeries**](AllApi.md#MetricsGetEdgeSegmentSeries) | **Post** /metrics/getEdgeSegmentSeries | Get flow metric time series data by segment id
[**MetricsGetEdgeStatusMetrics**](AllApi.md#MetricsGetEdgeStatusMetrics) | **Post** /metrics/getEdgeStatusMetrics | Get Edge healthStats metrics for an interval
[**MetricsGetEdgeStatusSeries**](AllApi.md#MetricsGetEdgeStatusSeries) | **Post** /metrics/getEdgeStatusSeries | Get Edge healthStats time series for an interval 
[**MetricsGetGatewayStatusMetrics**](AllApi.md#MetricsGetGatewayStatusMetrics) | **Post** /metrics/getGatewayStatusMetrics | Get Gateway health metric summaries for an interval
[**MetricsGetGatewayStatusSeries**](AllApi.md#MetricsGetGatewayStatusSeries) | **Post** /metrics/getGatewayStatusSeries | Get Gateway health metric time series for an interval
[**MonitoringGetAggregateEdgeLinkMetrics**](AllApi.md#MonitoringGetAggregateEdgeLinkMetrics) | **Post** /monitoring/getAggregateEdgeLinkMetrics | Get aggregate Edge link metrics across enterprises
[**MonitoringGetAggregateEnterpriseEvents**](AllApi.md#MonitoringGetAggregateEnterpriseEvents) | **Post** /monitoring/getAggregateEnterpriseEvents | Get events across all enterprises
[**MonitoringGetAggregates**](AllApi.md#MonitoringGetAggregates) | **Post** /monitoring/getAggregates | Get aggregate enterprise and Edge information
[**MonitoringGetEnterpriseBgpPeerStatus**](AllApi.md#MonitoringGetEnterpriseBgpPeerStatus) | **Post** /monitoring/getEnterpriseBgpPeerStatus | Get gateway BGP peer status for all enterprise gateways
[**MonitoringGetEnterpriseEdgeBgpPeerStatus**](AllApi.md#MonitoringGetEnterpriseEdgeBgpPeerStatus) | **Post** /monitoring/getEnterpriseEdgeBgpPeerStatus | Get Edge BGP peer status for all enterprise Edges
[**MonitoringGetEnterpriseEdgeClusterStatus**](AllApi.md#MonitoringGetEnterpriseEdgeClusterStatus) | **Post** /monitoring/getEnterpriseEdgeClusterStatus | Get Edge Cluster status
[**MonitoringGetEnterpriseEdgeLinkStatus**](AllApi.md#MonitoringGetEnterpriseEdgeLinkStatus) | **Post** /monitoring/getEnterpriseEdgeLinkStatus | Get Edge and link status data
[**MonitoringGetEnterpriseEdgeStatus**](AllApi.md#MonitoringGetEnterpriseEdgeStatus) | **Post** /monitoring/getEnterpriseEdgeStatus | Get Enterprise Edge monitoring status
[**MonitoringGetNetworkGatewayStatus**](AllApi.md#MonitoringGetNetworkGatewayStatus) | **Post** /monitoring/getNetworkGatewayStatus | Get Network Gateway monitoring status
[**NetworkDeleteNetworkGatewayPool**](AllApi.md#NetworkDeleteNetworkGatewayPool) | **Post** /network/deleteNetworkGatewayPool | Delete gateway pool
[**NetworkGetNetworkConfigurations**](AllApi.md#NetworkGetNetworkConfigurations) | **Post** /network/getNetworkConfigurations | Get operator configuration profiles
[**NetworkGetNetworkEnterprises**](AllApi.md#NetworkGetNetworkEnterprises) | **Post** /network/getNetworkEnterprises | Get enterprises on network
[**NetworkGetNetworkGatewayPools**](AllApi.md#NetworkGetNetworkGatewayPools) | **Post** /network/getNetworkGatewayPools | Get gateway pools on network
[**NetworkGetNetworkGateways**](AllApi.md#NetworkGetNetworkGateways) | **Post** /network/getNetworkGateways | Get gateways on network
[**NetworkGetNetworkOperatorUsers**](AllApi.md#NetworkGetNetworkOperatorUsers) | **Post** /network/getNetworkOperatorUsers | Get operator users for network
[**NetworkInsertNetworkGatewayPool**](AllApi.md#NetworkInsertNetworkGatewayPool) | **Post** /network/insertNetworkGatewayPool | Create gateway pool
[**NetworkUpdateNetworkGatewayPoolAttributes**](AllApi.md#NetworkUpdateNetworkGatewayPoolAttributes) | **Post** /network/updateNetworkGatwayPoolAttributes | Update gateway pool attributes
[**OperatorUserDeleteOperatorUser**](AllApi.md#OperatorUserDeleteOperatorUser) | **Post** /operatorUser/deleteOperatorUser | Delete operator user
[**OperatorUserGetOperatorUser**](AllApi.md#OperatorUserGetOperatorUser) | **Post** /operatorUser/getOperatorUser | Get operator user
[**OperatorUserInsertOperatorUser**](AllApi.md#OperatorUserInsertOperatorUser) | **Post** /operatorUser/insertOperatorUser | Create operator user
[**OperatorUserUpdateOperatorUser**](AllApi.md#OperatorUserUpdateOperatorUser) | **Post** /operatorUser/updateOperatorUser | Update operator user
[**RoleCreateRoleCustomization**](AllApi.md#RoleCreateRoleCustomization) | **Post** /role/createRoleCustomization | Create role customization
[**RoleDeleteRoleCustomization**](AllApi.md#RoleDeleteRoleCustomization) | **Post** /role/deleteRoleCustomization | Delete role customization
[**RoleGetUserTypeRoles**](AllApi.md#RoleGetUserTypeRoles) | **Post** /role/getUserTypeRoles | Get roles per user type
[**RoleSetEnterpriseDelegatedToEnterpriseProxy**](AllApi.md#RoleSetEnterpriseDelegatedToEnterpriseProxy) | **Post** /role/setEnterpriseDelegatedToEnterpriseProxy | Grant enterprise access to partner
[**RoleSetEnterpriseDelegatedToOperator**](AllApi.md#RoleSetEnterpriseDelegatedToOperator) | **Post** /role/setEnterpriseDelegatedToOperator | Grant enterprise access to network operator
[**RoleSetEnterpriseProxyDelegatedToOperator**](AllApi.md#RoleSetEnterpriseProxyDelegatedToOperator) | **Post** /role/setEnterpriseProxyDelegatedToOperator | Grant enterprise proxy access to network operator
[**RoleSetEnterpriseUserManagementDelegatedToOperator**](AllApi.md#RoleSetEnterpriseUserManagementDelegatedToOperator) | **Post** /role/setEnterpriseUserManagementDelegatedToOperator | Grant enterprise user access to network operator
[**SetClientDeviceHostName**](AllApi.md#SetClientDeviceHostName) | **Post** /clientDevice/setClientDeviceHostName | Set hostname for client device
[**SystemPropertyGetSystemProperties**](AllApi.md#SystemPropertyGetSystemProperties) | **Post** /systemProperty/getSystemProperties | Get all system properties
[**SystemPropertyGetSystemProperty**](AllApi.md#SystemPropertyGetSystemProperty) | **Post** /systemProperty/getSystemProperty | Get system property
[**SystemPropertyInsertOrUpdateSystemProperty**](AllApi.md#SystemPropertyInsertOrUpdateSystemProperty) | **Post** /systemProperty/insertOrUpdateSystemProperty | Create or update system property
[**SystemPropertyInsertSystemProperty**](AllApi.md#SystemPropertyInsertSystemProperty) | **Post** /systemProperty/insertSystemProperty | Create system property
[**SystemPropertyUpdateSystemProperty**](AllApi.md#SystemPropertyUpdateSystemProperty) | **Post** /systemProperty/updateSystemProperty | Update system property
[**VcoDiagnosticsGetVcoDbDiagnostics**](AllApi.md#VcoDiagnosticsGetVcoDbDiagnostics) | **Post** /vcoDiagnostics/getVcoDbDiagnostics | Get VCO Database Diagnostics
[**VcoInventoryAssociateEdge**](AllApi.md#VcoInventoryAssociateEdge) | **Post** /vcoInventory/associateEdge | Assign Edge to enterprise
[**VcoInventoryGetInventoryItems**](AllApi.md#VcoInventoryGetInventoryItems) | **Post** /vcoInventory/getInventoryItems | Get available VCO inventory items
[**VpnGenerateVpnGatewayConfiguration**](AllApi.md#VpnGenerateVpnGatewayConfiguration) | **Post** /vpn/generateVpnGatewayConfiguration | Provision a non-SD-WAN VPN site



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


## EnterpriseDecodeEnterpriseKey

> EnterpriseDecodeEnterpriseKeyResult EnterpriseDecodeEnterpriseKey(ctx, body)
Decrypt an enterprise key

Decrypt an enterprise key  Privileges required:  `READ` `ENTERPRISE_KEYS`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseDecodeEnterpriseKey**](EnterpriseDecodeEnterpriseKey.md)|  | 

### Return type

[**EnterpriseDecodeEnterpriseKeyResult**](enterprise_decode_enterprise_key_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseDeleteEnterprise

> EnterpriseDeleteEnterpriseResult EnterpriseDeleteEnterprise(ctx, body)
Delete enterprise

Deletes the specified enterprise (by `id` or `enterpriseId`).  Privileges required:  `DELETE` `ENTERPRISE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseDeleteEnterprise**](EnterpriseDeleteEnterprise.md)|  | 

### Return type

[**EnterpriseDeleteEnterpriseResult**](enterprise_delete_enterprise_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseDeleteEnterpriseGatewayRecords

> EnterpriseDeleteEnterpriseGatewayRecordsResult EnterpriseDeleteEnterpriseGatewayRecords(ctx, body)
Delete gateway BGP neighbor record(s) for enterprise

Deletes the enterprise gateway BGP neighbor record(s) matching the specified gateway id(s) (`gatewayId`) and neighbor IP addresses (`neighborIp`).  Privileges required:  `DELETE` `NETWORK_SERVICE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseDeleteEnterpriseGatewayRecords**](EnterpriseDeleteEnterpriseGatewayRecords.md)|  | 

### Return type

[**EnterpriseDeleteEnterpriseGatewayRecordsResult**](enterprise_delete_enterprise_gateway_records_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseDeleteEnterpriseNetworkAllocation

> EnterpriseDeleteEnterpriseNetworkAllocationResult EnterpriseDeleteEnterpriseNetworkAllocation(ctx, body)
Delete enterprise network allocation

Deletes the specified enterprise network allocation (by `id`).  Privileges required:  `DELETE` `NETWORK_ALLOCATION`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseDeleteEnterpriseNetworkAllocation**](EnterpriseDeleteEnterpriseNetworkAllocation.md)|  | 

### Return type

[**EnterpriseDeleteEnterpriseNetworkAllocationResult**](enterprise_delete_enterprise_network_allocation_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseDeleteEnterpriseNetworkSegment

> DeletionConfirmation EnterpriseDeleteEnterpriseNetworkSegment(ctx, body)
Delete an enterprise network segment

Delete an enterprise network segment, by id.  Privileges required:  `DELETE` `NETWORK_ALLOCATION`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseDeleteEnterpriseNetworkSegment**](EnterpriseDeleteEnterpriseNetworkSegment.md)|  | 

### Return type

[**DeletionConfirmation**](deletion_confirmation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseDeleteEnterpriseService

> EnterpriseDeleteEnterpriseServiceResult EnterpriseDeleteEnterpriseService(ctx, body)
Delete enterprise network service

Deletes the specified enterprise network service (by `id`).  Privileges required:  `DELETE` `NETWORK_SERVICE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseDeleteEnterpriseService**](EnterpriseDeleteEnterpriseService.md)|  | 

### Return type

[**EnterpriseDeleteEnterpriseServiceResult**](enterprise_delete_enterprise_service_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseEncodeEnterpriseKey

> EnterpriseEncodeEnterpriseKeyResult EnterpriseEncodeEnterpriseKey(ctx, body)
Encrypt an enterprise key

Encrypt an enterprise key  Privileges required:  `CREATE` `ENTERPRISE_KEYS`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseEncodeEnterpriseKey**](EnterpriseEncodeEnterpriseKey.md)|  | 

### Return type

[**EnterpriseEncodeEnterpriseKeyResult**](enterprise_encode_enterprise_key_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseGetEnterprise

> EnterpriseGetEnterpriseResult EnterpriseGetEnterprise(ctx, body)
Get enterprise

Gets data for the specified enterprise with optional proxy (partner) detail.  Privileges required:  `READ` `ENTERPRISE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseGetEnterprise**](EnterpriseGetEnterprise.md)|  | 

### Return type

[**EnterpriseGetEnterpriseResult**](enterprise_get_enterprise_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseGetEnterpriseAddresses

> []EnterpriseGetEnterpriseAddressesResultItem EnterpriseGetEnterpriseAddresses(ctx, body)
Get enterprise IP address(es)

Gets all public IP address information for the management and control entities associated with the specified enterprise, including Orchestrator(s), Gateway(s), and datacenter(s).  Privileges required:  `READ` `ENTERPRISE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseGetEnterpriseAddresses**](EnterpriseGetEnterpriseAddresses.md)|  | 

### Return type

[**[]EnterpriseGetEnterpriseAddressesResultItem**](enterprise_get_enterprise_addresses_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseGetEnterpriseAlertConfigurations

> []EnterpriseGetEnterpriseAlertConfigurationsResultItem EnterpriseGetEnterpriseAlertConfigurations(ctx, body)
Get enterprise alert configuration(s)

Gets all alert configuration(s) for the specified enterprise.  Privileges required:  `READ` `ENTERPRISE_ALERT`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseGetEnterpriseAlertConfigurations**](EnterpriseGetEnterpriseAlertConfigurations.md)|  | 

### Return type

[**[]EnterpriseGetEnterpriseAlertConfigurationsResultItem**](enterprise_get_enterprise_alert_configurations_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseGetEnterpriseAlerts

> EnterpriseGetEnterpriseAlertsResult EnterpriseGetEnterpriseAlerts(ctx, body)
Get triggered enterprise alerts

Gets past triggered alerts for the specified enterprise.  Privileges required:  `READ` `ENTERPRISE_ALERT`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseGetEnterpriseAlerts**](EnterpriseGetEnterpriseAlerts.md)|  | 

### Return type

[**EnterpriseGetEnterpriseAlertsResult**](enterprise_get_enterprise_alerts_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseGetEnterpriseAllAlertRecipients

> EnterpriseGetEnterpriseAllAlertRecipientsResult EnterpriseGetEnterpriseAllAlertRecipients(ctx, body)
Get recipients receiving all enterprise alerts

Gets all recipients configured to receive all alerts for the specified enterprise.  Privileges required:  `READ` `ENTERPRISE_ALERT`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseGetEnterpriseAllAlertRecipients**](EnterpriseGetEnterpriseAllAlertRecipients.md)|  | 

### Return type

[**EnterpriseGetEnterpriseAllAlertRecipientsResult**](enterprise_get_enterprise_all_alert_recipients_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseGetEnterpriseCapabilities

> EnterpriseGetEnterpriseCapabilitiesResult EnterpriseGetEnterpriseCapabilities(ctx, body)
Get enterprise capabilities

Gets all enterprise capabilities (e.g. BGP, COS mapping, PKI, etc.) currently enabled/disabled for the specified enterprise.  Privileges required:  `READ` `ENTERPRISE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseGetEnterpriseCapabilities**](EnterpriseGetEnterpriseCapabilities.md)|  | 

### Return type

[**EnterpriseGetEnterpriseCapabilitiesResult**](enterprise_get_enterprise_capabilities_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseGetEnterpriseConfigurations

> []EnterpriseGetEnterpriseConfigurationsResultItem EnterpriseGetEnterpriseConfigurations(ctx, body)
Get enterprise configuration profiles

Gets all configuration profiles, with optional Edge and/or module details, for the specified enterprise.  Privileges required:  `READ` `ENTERPRISE_PROFILE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseGetEnterpriseConfigurations**](EnterpriseGetEnterpriseConfigurations.md)|  | 

### Return type

[**[]EnterpriseGetEnterpriseConfigurationsResultItem**](enterprise_get_enterprise_configurations_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseGetEnterpriseEdges

> []EnterpriseGetEnterpriseEdgesResultItem EnterpriseGetEnterpriseEdges(ctx, body)
Get enterprise Edges

Gets all Edges associated with the specified enterprise, including optional site, link, and configuration details.  Privileges required:  `READ` `EDGE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseGetEnterpriseEdges**](EnterpriseGetEnterpriseEdges.md)|  | 

### Return type

[**[]EnterpriseGetEnterpriseEdgesResultItem**](enterprise_get_enterprise_edges_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseGetEnterpriseGatewayHandoff

> EnterpriseGetEnterpriseGatewayHandoffResult EnterpriseGetEnterpriseGatewayHandoff(ctx, body)
Get enterprise gateway handoff configuration

Gets the gateway handoff configuration for the specified enterprise.  Privileges required:  `READ` `ENTERPRISE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseGetEnterpriseGatewayHandoff**](EnterpriseGetEnterpriseGatewayHandoff.md)|  | 

### Return type

[**EnterpriseGetEnterpriseGatewayHandoffResult**](enterprise_get_enterprise_gateway_handoff_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseGetEnterpriseMaximumSegments

> EnterpriseGetEnterprisePropertyResult EnterpriseGetEnterpriseMaximumSegments(ctx, body)
Get enterprise maximum segment

Gets maximum segments for the specified enterprise.  Privileges required:  `READ` `ENTERPRISE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseGetEnterpriseMaximumSegments**](EnterpriseGetEnterpriseMaximumSegments.md)|  | 

### Return type

[**EnterpriseGetEnterprisePropertyResult**](enterprise_get_enterprise_property_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseGetEnterpriseNetworkAllocation

> EnterpriseGetEnterpriseNetworkAllocationResult EnterpriseGetEnterpriseNetworkAllocation(ctx, body)
Get enterprise network allocation

Gets the specified network allocation object (by `id`) for the specified enterprise.  Privileges required:  `READ` `NETWORK_ALLOCATION`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseGetEnterpriseNetworkAllocation**](EnterpriseGetEnterpriseNetworkAllocation.md)|  | 

### Return type

[**EnterpriseGetEnterpriseNetworkAllocationResult**](enterprise_get_enterprise_network_allocation_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseGetEnterpriseNetworkAllocations

> []EnterpriseGetEnterpriseNetworkAllocationsResultItem EnterpriseGetEnterpriseNetworkAllocations(ctx, body)
Get enterprise network allocations

Gets all network allocations for the specified enterprise.  Privileges required:  `READ` `NETWORK_ALLOCATION`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseGetEnterpriseNetworkAllocations**](EnterpriseGetEnterpriseNetworkAllocations.md)|  | 

### Return type

[**[]EnterpriseGetEnterpriseNetworkAllocationsResultItem**](enterprise_get_enterprise_network_allocations_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseGetEnterpriseNetworkSegments

> []EnterpriseGetEnterpriseNetworkSegmentsResultItem EnterpriseGetEnterpriseNetworkSegments(ctx, body)
Get all network segment objects defined on an enterprise

Retrieve a list of all of the network segments defined forthe given enterprise.  Privileges required:  `READ` `NETWORK_ALLOCATION`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseGetEnterpriseNetworkSegments**](EnterpriseGetEnterpriseNetworkSegments.md)|  | 

### Return type

[**[]EnterpriseGetEnterpriseNetworkSegmentsResultItem**](enterprise_get_enterprise_network_segments_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseGetEnterpriseProperty

> EnterpriseGetEnterprisePropertyResult EnterpriseGetEnterpriseProperty(ctx, body)
Get enterprise property

Get a enterprise property by object id or other attribute.  Privileges required:  `READ` `ENTERPRISE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseGetEnterpriseProperty**](EnterpriseGetEnterpriseProperty.md)|  | 

### Return type

[**EnterpriseGetEnterprisePropertyResult**](enterprise_get_enterprise_property_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseGetEnterpriseRouteConfiguration

> EnterpriseGetEnterpriseRouteConfigurationResult EnterpriseGetEnterpriseRouteConfiguration(ctx, body)
Get enterprise global routing preferences

Gets all global routing preferences for the specified enterprise, incuding enterprise route advertisement, routing preferences, OSPF, and BGP advertisement policy.  Privileges required:  `READ` `ENTERPRISE_PROFILE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseGetEnterpriseRouteConfiguration**](EnterpriseGetEnterpriseRouteConfiguration.md)|  | 

### Return type

[**EnterpriseGetEnterpriseRouteConfigurationResult**](enterprise_get_enterprise_route_configuration_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseGetEnterpriseRouteTable

> EnterpriseGetEnterpriseRouteTableResult EnterpriseGetEnterpriseRouteTable(ctx, body)
Get enterprise route table

Gets the composite enterprise route table, optionally scoped by profile(s). The returned routes include static routes, directly connected routes, and learned routes.  Privileges required:  `READ` `ENTERPRISE_PROFILE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseGetEnterpriseRouteTable**](EnterpriseGetEnterpriseRouteTable.md)|  | 

### Return type

[**EnterpriseGetEnterpriseRouteTableResult**](enterprise_get_enterprise_route_table_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseGetEnterpriseServices

> []EnterpriseGetEnterpriseServicesResultItem EnterpriseGetEnterpriseServices(ctx, body)
Get enterprise network services

Gets all network service JSON objects defined for the specified enterprise.  Privileges required:  `READ` `NETWORK_SERVICE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseGetEnterpriseServices**](EnterpriseGetEnterpriseServices.md)|  | 

### Return type

[**[]EnterpriseGetEnterpriseServicesResultItem**](enterprise_get_enterprise_services_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseGetEnterpriseUsers

> []EnterpriseGetEnterpriseUsersResultItem EnterpriseGetEnterpriseUsers(ctx, body)
Get enterprise users

Gets all enterprise users for the specified enterprise (by `enterpriseId`).  Privileges required:  `READ` `ENTERPRISE`  `READ` `ENTERPRISE_USER`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseGetEnterpriseUsers**](EnterpriseGetEnterpriseUsers.md)|  | 

### Return type

[**[]EnterpriseGetEnterpriseUsersResultItem**](enterprise_get_enterprise_users_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseInsertEnterprise

> EnterpriseInsertEnterpriseResult EnterpriseInsertEnterprise(ctx, body)
Create enterprise

Creates a new enterprise, which is owned by the operator.  Privileges required:  `CREATE` `ENTERPRISE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseInsertEnterprise**](EnterpriseInsertEnterprise.md)|  | 

### Return type

[**EnterpriseInsertEnterpriseResult**](enterprise_insert_enterprise_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseInsertEnterpriseNetworkAllocation

> EnterpriseInsertEnterpriseNetworkAllocationResult EnterpriseInsertEnterpriseNetworkAllocation(ctx, body)
Create enterprise network allocation

Creates a new network allocation for the specified enterprise.  Privileges required:  `CREATE` `NETWORK_ALLOCATION`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseInsertEnterpriseNetworkAllocation**](EnterpriseInsertEnterpriseNetworkAllocation.md)|  | 

### Return type

[**EnterpriseInsertEnterpriseNetworkAllocationResult**](enterprise_insert_enterprise_network_allocation_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseInsertEnterpriseNetworkSegment

> EnterpriseInsertEnterpriseNetworkSegmentResult EnterpriseInsertEnterpriseNetworkSegment(ctx, body)
Create enterprise network segment

Creates a new network segment for the specified enterprise.  Privileges required:  `CREATE` `NETWORK_ALLOCATION`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseInsertEnterpriseNetworkSegment**](EnterpriseInsertEnterpriseNetworkSegment.md)|  | 

### Return type

[**EnterpriseInsertEnterpriseNetworkSegmentResult**](enterprise_insert_enterprise_network_segment_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseInsertEnterpriseService

> EnterpriseInsertEnterpriseServiceResult EnterpriseInsertEnterpriseService(ctx, body)
Create enterprise service

Creates a new enterprise service for the specified enterprise.  Privileges required:  `CREATE` `NETWORK_SERVICE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseInsertEnterpriseService**](EnterpriseInsertEnterpriseService.md)|  | 

### Return type

[**EnterpriseInsertEnterpriseServiceResult**](enterprise_insert_enterprise_service_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseInsertEnterpriseUser

> EnterpriseInsertEnterpriseUserResult EnterpriseInsertEnterpriseUser(ctx, body)
Create enterprise user

Creates a new enterprise user.  Privileges required:  `CREATE` `ENTERPRISE_USER`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**NewEnterpriseUser**](NewEnterpriseUser.md)|  | 

### Return type

[**EnterpriseInsertEnterpriseUserResult**](enterprise_insert_enterprise_user_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseInsertOrUpdateEnterpriseAlertConfigurations

> EnterpriseInsertOrUpdateEnterpriseAlertConfigurationsResult EnterpriseInsertOrUpdateEnterpriseAlertConfigurations(ctx, body)
Create, update, or delete enterprise alert configurations

Creates, updates, or deletes alert configurations for the specified enterprise. Returns the array of alert configurations submitted, with ids added for the entries that have been successfully created. If an entry is not successfully created or updated, an `error` property is included in the response.  Privileges required:  `CREATE` `ENTERPRISE_ALERT`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseInsertOrUpdateEnterpriseAlertConfigurations**](EnterpriseInsertOrUpdateEnterpriseAlertConfigurations.md)|  | 

### Return type

[**EnterpriseInsertOrUpdateEnterpriseAlertConfigurationsResult**](enterprise_insert_or_update_enterprise_alert_configurations_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseInsertOrUpdateEnterpriseCapability

> EnterpriseInsertOrUpdateEnterpriseCapabilityResult EnterpriseInsertOrUpdateEnterpriseCapability(ctx, body)
Create or update enterprise capability

Creates or updates the specified capability (`name`, `value`) for the specified enterprise (by `enterpriseId`).  Privileges required:  `UPDATE` `ENTERPRISE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseInsertOrUpdateEnterpriseCapability**](EnterpriseInsertOrUpdateEnterpriseCapability.md)|  | 

### Return type

[**EnterpriseInsertOrUpdateEnterpriseCapabilityResult**](enterprise_insert_or_update_enterprise_capability_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseInsertOrUpdateEnterpriseGatewayHandoff

> EnterpriseInsertOrUpdateEnterpriseGatewayHandoffResult EnterpriseInsertOrUpdateEnterpriseGatewayHandoff(ctx, body)
Create or update enterprise gateway handoff configuration

Creates or updates a gateway handoff configuration for the specified enterprise (by `enterpriseId`).  Privileges required:  `UPDATE` `ENTERPRISE`  `UPDATE` `GATEWAY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseInsertOrUpdateEnterpriseGatewayHandoff**](EnterpriseInsertOrUpdateEnterpriseGatewayHandoff.md)|  | 

### Return type

[**EnterpriseInsertOrUpdateEnterpriseGatewayHandoffResult**](enterprise_insert_or_update_enterprise_gateway_handoff_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseInsertOrUpdateEnterpriseProperty

> EnterpriseInsertOrUpdateEnterprisePropertyResult EnterpriseInsertOrUpdateEnterpriseProperty(ctx, body)
Insert or update an enterprise property

Insert a enterprise property. If property with the given name already exists, the property will be updated.  Privileges required:  `UPDATE` `ENTERPRISE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseInsertOrUpdateEnterpriseProperty**](EnterpriseInsertOrUpdateEnterpriseProperty.md)|  | 

### Return type

[**EnterpriseInsertOrUpdateEnterprisePropertyResult**](enterprise_insert_or_update_enterprise_property_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseProxyDeleteEnterpriseProxyUser

> EnterpriseProxyDeleteEnterpriseProxyUserResult EnterpriseProxyDeleteEnterpriseProxyUser(ctx, body)
Delete partner admin user

Deletes the specified enterprise proxy user (by `id` or `username`). Note: `enterpriseProxyId` is a required parameter when invoking this method as an operator or as a partner user.  Privileges required:  `DELETE` `PROXY_USER`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseProxyDeleteEnterpriseProxyUser**](EnterpriseProxyDeleteEnterpriseProxyUser.md)|  | 

### Return type

[**EnterpriseProxyDeleteEnterpriseProxyUserResult**](enterprise_proxy_delete_enterprise_proxy_user_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseProxyGetEnterpriseProxyEdgeInventory

> []EnterpriseProxyGetEnterpriseProxyEdgeInventoryResultItem EnterpriseProxyGetEnterpriseProxyEdgeInventory(ctx, body)
Get partner enterprises and associated Edge inventory

Gets all partner enterprises and their associated Edge inventory.  Privileges required:  `READ` `ENTERPRISE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseProxyGetEnterpriseProxyEdgeInventory**](EnterpriseProxyGetEnterpriseProxyEdgeInventory.md)|  | 

### Return type

[**[]EnterpriseProxyGetEnterpriseProxyEdgeInventoryResultItem**](enterprise_proxy_get_enterprise_proxy_edge_inventory_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseProxyGetEnterpriseProxyEnterprises

> []EnterpriseProxyGetEnterpriseProxyEnterprisesResultItem EnterpriseProxyGetEnterpriseProxyEnterprises(ctx, body)
Get partner enterprises

Gets all partner enterprises. Optionally includes all Edges or Edge counts.  Privileges required:  `READ` `ENTERPRISE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseProxyGetEnterpriseProxyEnterprises**](EnterpriseProxyGetEnterpriseProxyEnterprises.md)|  | 

### Return type

[**[]EnterpriseProxyGetEnterpriseProxyEnterprisesResultItem**](enterprise_proxy_get_enterprise_proxy_enterprises_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseProxyGetEnterpriseProxyGatewayPools

> []EnterpriseProxyGetEnterpriseProxyGatewayPoolsResultItem EnterpriseProxyGetEnterpriseProxyGatewayPools(ctx, body)
Get gateway pools

Gets all gateway pools associated with the specified enterprise proxy. Optionally includes all gateways or enterprises belonging to each pool.  Privileges required:  `READ` `GATEWAY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseProxyGetEnterpriseProxyGatewayPools**](EnterpriseProxyGetEnterpriseProxyGatewayPools.md)|  | 

### Return type

[**[]EnterpriseProxyGetEnterpriseProxyGatewayPoolsResultItem**](enterprise_proxy_get_enterprise_proxy_gateway_pools_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseProxyGetEnterpriseProxyGateways

> []EnterpriseProxyGetEnterpriseProxyGatewaysResultItem EnterpriseProxyGetEnterpriseProxyGateways(ctx, body)
Get gateways for enterprise proxy

Gets all gateways associated with the specified enterprise proxy.  Privileges required:  `READ` `GATEWAY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseProxyGetEnterpriseProxyGateways**](EnterpriseProxyGetEnterpriseProxyGateways.md)|  | 

### Return type

[**[]EnterpriseProxyGetEnterpriseProxyGatewaysResultItem**](enterprise_proxy_get_enterprise_proxy_gateways_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseProxyGetEnterpriseProxyOperatorProfiles

> []EnterpriseProxyGetEnterpriseProxyOperatorProfilesResultItem EnterpriseProxyGetEnterpriseProxyOperatorProfiles(ctx, body)
Get partner operator profiles

Gets all operator profiles associated with the specified partner (MSP), as assigned by the operator.  Privileges required:  `READ` `OPERATOR_PROFILE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseProxyGetEnterpriseProxyOperatorProfiles**](EnterpriseProxyGetEnterpriseProxyOperatorProfiles.md)|  | 

### Return type

[**[]EnterpriseProxyGetEnterpriseProxyOperatorProfilesResultItem**](enterprise_proxy_get_enterprise_proxy_operator_profiles_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseProxyGetEnterpriseProxyProperty

> EnterpriseProxyGetEnterpriseProxyPropertyResult EnterpriseProxyGetEnterpriseProxyProperty(ctx, body)
Get enterprise proxy property

Get a enterprise proxy property by object id or other attribute.  Privileges required:  `READ` `PROXY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseProxyGetEnterpriseProxyProperty**](EnterpriseProxyGetEnterpriseProxyProperty.md)|  | 

### Return type

[**EnterpriseProxyGetEnterpriseProxyPropertyResult**](enterprise_proxy_get_enterprise_proxy_property_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseProxyGetEnterpriseProxyUser

> EnterpriseProxyGetEnterpriseProxyUserResult EnterpriseProxyGetEnterpriseProxyUser(ctx, body)
Get enterprise proxy user

Gets the specified enterprise proxy user (by `id` or `username`).  Privileges required:  `READ` `PROXY_USER`  `READ` `PROXY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseProxyGetEnterpriseProxyUser**](EnterpriseProxyGetEnterpriseProxyUser.md)|  | 

### Return type

[**EnterpriseProxyGetEnterpriseProxyUserResult**](enterprise_proxy_get_enterprise_proxy_user_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseProxyGetEnterpriseProxyUsers

> []EnterpriseProxyGetEnterpriseProxyUsersResultItem EnterpriseProxyGetEnterpriseProxyUsers(ctx, body)
Get enterprise proxy admin users

Gets all proxy admin users for the specified enterprise.  Privileges required:  `READ` `ENTERPRISE`  `READ` `PROXY_USER`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseProxyGetEnterpriseProxyUsers**](EnterpriseProxyGetEnterpriseProxyUsers.md)|  | 

### Return type

[**[]EnterpriseProxyGetEnterpriseProxyUsersResultItem**](enterprise_proxy_get_enterprise_proxy_users_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseProxyInsertEnterpriseProxyEnterprise

> EnterpriseProxyInsertEnterpriseProxyEnterpriseResult EnterpriseProxyInsertEnterpriseProxyEnterprise(ctx, body)
Create enterprise owned by MSP

Creates a new enterprise owned by an MSP. Whereas the `insertEnterprise` method creates an enterprise in the global or network context with no MSP association, this method creates one owned by an MSP, as determined by the credentials of the caller.  Privileges required:  `CREATE` `ENTERPRISE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseProxyInsertEnterpriseProxyEnterprise**](EnterpriseProxyInsertEnterpriseProxyEnterprise.md)|  | 

### Return type

[**EnterpriseProxyInsertEnterpriseProxyEnterpriseResult**](enterprise_proxy_insert_enterprise_proxy_enterprise_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseProxyInsertEnterpriseProxyUser

> EnterpriseProxyInsertEnterpriseProxyUserResult EnterpriseProxyInsertEnterpriseProxyUser(ctx, body)
Create partner admin user

Creates a new partner admin user.  Privileges required:  `CREATE` `PROXY_USER`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**NewEnterpriseProxyUser**](NewEnterpriseProxyUser.md)|  | 

### Return type

[**EnterpriseProxyInsertEnterpriseProxyUserResult**](enterprise_proxy_insert_enterprise_proxy_user_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseProxyInsertOrUpdateEnterpriseProxyProperty

> EnterpriseProxyInsertOrUpdateEnterpriseProxyPropertyResult EnterpriseProxyInsertOrUpdateEnterpriseProxyProperty(ctx, body)
Insert or update an enterprise proxy (MSP) property

Insert a enterprise proxy property. If property with the given name already exists, the property will be updated.  Privileges required:  `UPDATE` `PROXY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseProxyInsertOrUpdateEnterpriseProxyProperty**](EnterpriseProxyInsertOrUpdateEnterpriseProxyProperty.md)|  | 

### Return type

[**EnterpriseProxyInsertOrUpdateEnterpriseProxyPropertyResult**](enterprise_proxy_insert_or_update_enterprise_proxy_property_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseProxyUpdateEnterpriseProxyUser

> EnterpriseProxyUpdateEnterpriseProxyUserResult EnterpriseProxyUpdateEnterpriseProxyUser(ctx, body)
Update enterprise proxy admin user

Updates the specified enterprise proxy admin user (by object `id` or other identifying attribute). Expects an `_update` object containing the field(s) to be updated and their corresponding value(s).  Privileges required:  `UPDATE` `PROXY_USER`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseProxyUpdateEnterpriseProxyUser**](EnterpriseProxyUpdateEnterpriseProxyUser.md)|  | 

### Return type

[**EnterpriseProxyUpdateEnterpriseProxyUserResult**](enterprise_proxy_update_enterprise_proxy_user_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseSetEnterpriseAllAlertRecipients

> EnterpriseSetEnterpriseAllAlertRecipientsResult EnterpriseSetEnterpriseAllAlertRecipients(ctx, body)
Specify recipients for all enterprise alerts

Specifies the recipients who should receive all alerts for the specified enterprise.  Privileges required:  `UPDATE` `ENTERPRISE_ALERT`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseSetEnterpriseAllAlertRecipients**](EnterpriseSetEnterpriseAllAlertRecipients.md)|  | 

### Return type

[**EnterpriseSetEnterpriseAllAlertRecipientsResult**](enterprise_set_enterprise_all_alert_recipients_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseSetEnterpriseMaximumSegments

> EnterpriseSetEnterpriseMaximumSegmentsResult EnterpriseSetEnterpriseMaximumSegments(ctx, body)
Set enterprise maximum segments

Set maximum segments for the specified enterprise (by `enterpriseId`).  Privileges required:  `UPDATE` `ENTERPRISE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseSetEnterpriseMaximumSegments**](EnterpriseSetEnterpriseMaximumSegments.md)|  | 

### Return type

[**EnterpriseSetEnterpriseMaximumSegmentsResult**](enterprise_set_enterprise_maximum_segments_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseSetEnterpriseOperatorConfiguration

> EnterpriseSetEnterpriseOperatorConfigurationResult EnterpriseSetEnterpriseOperatorConfiguration(ctx, body)
Apply operator profile to given enterprise

Sets the operator configuration for the specified enterprise (by `enterpriseId`). This overrides any network-assigned operator configuration and the network default operator configuration.  Privileges required:  `UPDATE` `ENTERPRISE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseSetEnterpriseOperatorConfiguration**](EnterpriseSetEnterpriseOperatorConfiguration.md)|  | 

### Return type

[**EnterpriseSetEnterpriseOperatorConfigurationResult**](enterprise_set_enterprise_operator_configuration_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseUpdateEnterprise

> RowsModifiedConfirmation EnterpriseUpdateEnterprise(ctx, body)
Update enterprise

Updates the specified enterprise (by `enterpriseId` or `name`). Expects an `_update` object containing the field(s) to be updated and their corresponding value(s).  Privileges required:  `UPDATE` `ENTERPRISE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseUpdateEnterprise1**](EnterpriseUpdateEnterprise1.md)|  | 

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


## EnterpriseUpdateEnterpriseNetworkAllocation

> EnterpriseUpdateEnterpriseNetworkAllocationResult EnterpriseUpdateEnterpriseNetworkAllocation(ctx, body)
Update enterprise network allocation

Updates the specified enterprise network allocation (by `id`). Expects an `_update` object containing the field(s) to be updated and their corresponding value(s).  Privileges required:  `UPDATE` `NETWORK_ALLOCATION`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseUpdateEnterpriseNetworkAllocation**](EnterpriseUpdateEnterpriseNetworkAllocation.md)|  | 

### Return type

[**EnterpriseUpdateEnterpriseNetworkAllocationResult**](enterprise_update_enterprise_network_allocation_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseUpdateEnterpriseNetworkSegment

> EnterpriseUpdateEnterpriseNetworkSegmentResult EnterpriseUpdateEnterpriseNetworkSegment(ctx, body)
Update an enterprise network segment

Update an enterprise network segment.  Privileges required:  `UPDATE` `NETWORK_ALLOCATION`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseUpdateEnterpriseNetworkSegment**](EnterpriseUpdateEnterpriseNetworkSegment.md)|  | 

### Return type

[**EnterpriseUpdateEnterpriseNetworkSegmentResult**](enterprise_update_enterprise_network_segment_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseUpdateEnterpriseRoute

> EnterpriseUpdateEnterpriseRouteResult EnterpriseUpdateEnterpriseRoute(ctx, body)
Update preferred VPN exits for a route

Updates the specified enterprise route, set advertisement, and cost values. Required parameters include the original route, as returned by `enterprise/getEnterpriseRouteTable`, and the updated route with modified advertisement and route preference ordering.  Privileges required:  `UPDATE` `ENTERPRISE_PROFILE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseUpdateEnterpriseRoute**](EnterpriseUpdateEnterpriseRoute.md)|  | 

### Return type

[**EnterpriseUpdateEnterpriseRouteResult**](enterprise_update_enterprise_route_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseUpdateEnterpriseRouteConfiguration

> EnterpriseUpdateEnterpriseRouteConfigurationResult EnterpriseUpdateEnterpriseRouteConfiguration(ctx, body)
Update enterprise global routing preferences

Updates the specified enterprise global routing preferences as specified by configuration `id` or logicalId.  Privileges required:  `UPDATE` `ENTERPRISE_PROFILE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseUpdateEnterpriseRouteConfiguration**](EnterpriseUpdateEnterpriseRouteConfiguration.md)|  | 

### Return type

[**EnterpriseUpdateEnterpriseRouteConfigurationResult**](enterprise_update_enterprise_route_configuration_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseUpdateEnterpriseSecurityPolicy

> EnterpriseUpdateEnterpriseSecurityPolicyResult EnterpriseUpdateEnterpriseSecurityPolicy(ctx, body)
Update enterprise security policy

Updates the security policy for the specified enterprise (by `enterpriseId`) using the passed in `ipsec` settings.  Privileges required:  `UPDATE` `ENTERPRISE_PROFILE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseUpdateEnterpriseSecurityPolicy**](EnterpriseUpdateEnterpriseSecurityPolicy.md)|  | 

### Return type

[**EnterpriseUpdateEnterpriseSecurityPolicyResult**](enterprise_update_enterprise_security_policy_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseUpdateEnterpriseService

> EnterpriseUpdateEnterpriseServiceResult EnterpriseUpdateEnterpriseService(ctx, body)
Update enterprise network service

Updates the enterprise service for the specified enterprise (by `id` or `enterpriseId`). Expects an `_update` object containing the field(s) to be updated and their corresponding value(s).  Privileges required:  `UPDATE` `NETWORK_SERVICE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseUpdateEnterpriseService**](EnterpriseUpdateEnterpriseService.md)|  | 

### Return type

[**EnterpriseUpdateEnterpriseServiceResult**](enterprise_update_enterprise_service_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseUserDeleteEnterpriseUser

> EnterpriseUserDeleteEnterpriseUserResult EnterpriseUserDeleteEnterpriseUser(ctx, body)
Delete enterprise user

Deletes the specified enterprise user (by `id` or `username`). Note: `enterpriseId` is a required parameter when invoking this method as an operator or as a partner user.  Privileges required:  `DELETE` `ENTERPRISE_USER`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseUserDeleteEnterpriseUser**](EnterpriseUserDeleteEnterpriseUser.md)|  | 

### Return type

[**EnterpriseUserDeleteEnterpriseUserResult**](enterprise_user_delete_enterprise_user_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseUserGetEnterpriseUser

> EnterpriseUserGetEnterpriseUserResult EnterpriseUserGetEnterpriseUser(ctx, body)
Get enterprise user

Gets the specified enterprise user (by `id` or `username`).  Privileges required:  `READ` `ENTERPRISE_USER`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseUserGetEnterpriseUser**](EnterpriseUserGetEnterpriseUser.md)|  | 

### Return type

[**EnterpriseUserGetEnterpriseUserResult**](enterprise_user_get_enterprise_user_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseUserUpdateEnterpriseUser

> EnterpriseUserUpdateEnterpriseUserResult EnterpriseUserUpdateEnterpriseUser(ctx, body)
Update enterprise user

Updates the specified enterprise user (by object `id` or other identifying attribute). Expects an `_update` object containing the field(s) to be updated and their corresponding value(s).  Privileges required:  `UPDATE` `ENTERPRISE_USER`, or  `UPDATE` `OPERATOR_USER`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EnterpriseUserUpdateEnterpriseUser1**](EnterpriseUserUpdateEnterpriseUser1.md)|  | 

### Return type

[**EnterpriseUserUpdateEnterpriseUserResult**](enterprise_user_update_enterprise_user_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventGetEnterpriseEvents

> EventGetEnterpriseEventsResult EventGetEnterpriseEvents(ctx, body)
Get Edge events

Gets Edge events in an enterprise or Edge context. Returns an array of Edge events sorted by `eventTime`..  Privileges required:  `READ` `ENTERPRISE_EVENT`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EventGetEnterpriseEvents**](EventGetEnterpriseEvents.md)|  | 

### Return type

[**EventGetEnterpriseEventsResult**](event_get_enterprise_events_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventGetOperatorEvents

> EventGetOperatorEventsResult EventGetOperatorEvents(ctx, body)
Get operator events

Gets operator events by network ID (optional). If not specified, will be taken for the caller's security context. Optionally, use a filter object to limit the number of events returned. Additionally, specify a time interval with an interval object. If no end date is specified, then the default is the current date. Specify a `gatewayId` to filter events for the specified gateway.  Privileges required:  `READ` `OPERATOR_EVENT`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EventGetOperatorEvents**](EventGetOperatorEvents.md)|  | 

### Return type

[**EventGetOperatorEventsResult**](event_get_operator_events_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventGetProxyEvents

> EventGetProxyEventsResult EventGetProxyEvents(ctx, body)
Fetch Partner events

Fetch Partner Events by enterprise ID. Optionally, a filter object to specify search criteria and limit on number of records. A time interval could be specified with an interval object. If no end date is given, it will default to the current date. Enterprise ID can be supplied to filter events to those of a specific enterprises.  Privileges required:  `READ` `PROXY_EVENT`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EventGetProxyEvents**](EventGetProxyEvents.md)|  | 

### Return type

[**EventGetProxyEventsResult**](event_get_proxy_events_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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


## GatewayDeleteGateway

> GatewayDeleteGatewayResult GatewayDeleteGateway(ctx, body)
Delete a gateway

Delete a gateway by id.  Privileges required:  `DELETE` `GATEWAY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**GatewayDeleteGateway**](GatewayDeleteGateway.md)|  | 

### Return type

[**GatewayDeleteGatewayResult**](gateway_delete_gateway_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayGatewayProvision

> GatewayGatewayProvisionResult GatewayGatewayProvision(ctx, body)
Provision gateway

Provisions a gateway into the specified operator network.  Privileges required:  `CREATE` `GATEWAY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**GatewayGatewayProvision**](GatewayGatewayProvision.md)|  | 

### Return type

[**GatewayGatewayProvisionResult**](gateway_gateway_provision_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayGetGatewayEdgeAssignments

> []GatewayGetGatewayEdgeAssignmentsResultItem GatewayGetGatewayEdgeAssignments(ctx, body)
Get edge assignments for a gateway

Get edge assignments for a gateway  Privileges required:  `READ` `GATEWAY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**GatewayGetGatewayEdgeAssignments**](GatewayGetGatewayEdgeAssignments.md)|  | 

### Return type

[**[]GatewayGetGatewayEdgeAssignmentsResultItem**](gateway_get_gateway_edge_assignments_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayUpdateGatewayAttributes

> GatewayUpdateGatewayAttributesResult GatewayUpdateGatewayAttributes(ctx, body)
Update gateway attributes

Updates gateway attributes (`name`, `ipAddress`, on-premise parametrization, and `description`) and associated site attributes for the specified gateway (by `id`).  Privileges required:  `UPDATE` `GATEWAY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**GatewayUpdateGatewayAttributes**](GatewayUpdateGatewayAttributes.md)|  | 

### Return type

[**GatewayUpdateGatewayAttributesResult**](gateway_update_gateway_attributes_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkQualityEventGetLinkQualityEvents

> LinkQualityEventGetLinkQualityEventsResult LinkQualityEventGetLinkQualityEvents(ctx, body)
Get link quality events

Gets all link quality scores per link for the specified Edge within the specified time interval. Rolls up link quality events to provide an aggregate score for the Edge. Returns an empty array if no link quality events are available in the specified timeframe.  Privileges required:  `READ` `EDGE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**LinkQualityEventGetLinkQualityEvents**](LinkQualityEventGetLinkQualityEvents.md)|  | 

### Return type

[**LinkQualityEventGetLinkQualityEventsResult**](link_quality_event_get_link_quality_events_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoginEnterpriseLogin

> LoginEnterpriseLogin(ctx, authorization)
Authenticate enterprise or partner (MSP) user

Authenticates an enterprise or partner (MSP) user and, upon successful login, returns a velocloud.session cookie. Pass this session cookie in the authentication header in subsequent VCO calls.  If you are using an HTTP client (e.g. Postman) that is configured to automatically follow HTTP redirects, a successful authentication request will cause your client to follow an HTTP 302 redirect to the portal 'Home' web page. Your session cookie can then be used to make VCO API calls.  Note that session cookies expire after a period of time specified in the VCO configuration (default is 24 hours).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | [**AuthObject**](AuthObject.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoginOperatorLogin

> LoginOperatorLogin(ctx, authorization)
Authenticate operator user

Authenticates an operator user and, upon successful login, returns a velocloud.session cookie. Pass this session cookie in the authentication header in subsequent VCO calls.  If you are using an HTTP client (e.g. Postman) that is configured to automatically follow HTTP redirects, a successful authentication request will cause your client to follow an HTTP 302 redirect to the portal 'Home' web page. Your session cookie can then be used to make VCO API calls.   Note that session cookies expire after a period of time specified in the VCO configuration (default is 24 hours).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | [**AuthObject**](AuthObject.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Logout

> Logout(ctx, )
Logout and invalidate authorization session cookie

Logs out the VCO API user and invalidates the session cookie.

### Required Parameters

This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Meta

> InlineResponse200 Meta(ctx, apiPath)
Get Swagger specification for any VCO API call

Gets the Swagger specification for any VCO API call.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiPath** | **string**| the path to another api method, starting after /rest/ | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsGetEdgeAppLinkMetrics

> []MetricsGetEdgeAppLinkMetricsResultItem MetricsGetEdgeAppLinkMetrics(ctx, body)
Get flow metric aggregate data

Gets flow metric summaries for the specified time interval by link. On success, this method returns an array of flow data in which each entry corresponds to a link on the specified Edge. In the request body, the `id` and `edgeId` property names are interchangeable. The `enterpriseId` property is required when this method is invoked in the operator context.  Privileges required:  `READ` `EDGE`  `VIEW_FLOW_STATS` `undefined`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MetricsGetEdgeAppLinkMetrics**](MetricsGetEdgeAppLinkMetrics.md)|  | 

### Return type

[**[]MetricsGetEdgeAppLinkMetricsResultItem**](metrics_get_edge_app_link_metrics_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsGetEdgeAppLinkSeries

> []MetricsGetEdgeAppLinkSeriesResultItem MetricsGetEdgeAppLinkSeries(ctx, body)
Get flow metric time series data

Gets flow metric time series data for the specified time interval by link. On success, this method returns an array of flow data in which each entry corresponds to a link on the specified Edge. In the request body, the `id` and `edgeId` property names are interchangeable. The `enterpriseId` property is required when this method is invoked in the operator context.  Privileges required:  `READ` `EDGE`  `VIEW_FLOW_STATS` `undefined`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MetricsGetEdgeAppLinkSeries**](MetricsGetEdgeAppLinkSeries.md)|  | 

### Return type

[**[]MetricsGetEdgeAppLinkSeriesResultItem**](metrics_get_edge_app_link_series_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsGetEdgeAppMetrics

> []MetricsGetEdgeAppMetricsResultItem MetricsGetEdgeAppMetrics(ctx, body)
Get flow metric aggregate data by application

Gets flow metric summaries for the specified time interval by application. On success, this method returns an array of flow data in which each entry corresponds to a single application. In the request body, the `id` and `edgeId` property names are interchangeable. The `enterpriseId` property is required when this method is invoked in the operator context.  Privileges required:  `READ` `EDGE`  `VIEW_FLOW_STATS` `undefined`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MetricsGetEdgeAppMetrics**](MetricsGetEdgeAppMetrics.md)|  | 

### Return type

[**[]MetricsGetEdgeAppMetricsResultItem**](metrics_get_edge_app_metrics_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsGetEdgeAppSeries

> []MetricsGetEdgeAppSeriesResultItem MetricsGetEdgeAppSeries(ctx, body)
Get flow metric time series data by application

Gets flow metric time series data for the specified time interval by application. On success, this method returns an array of flow data in which each entry corresponds to a single application. In the request body, the `id` and `edgeId` property names are interchangeable. The `enterpriseId` property is required when this method is invoked in the operator context.  Privileges required:  `READ` `EDGE`  `VIEW_FLOW_STATS` `undefined`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MetricsGetEdgeAppSeries**](MetricsGetEdgeAppSeries.md)|  | 

### Return type

[**[]MetricsGetEdgeAppSeriesResultItem**](metrics_get_edge_app_series_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsGetEdgeCategoryMetrics

> []MetricsGetEdgeCategoryMetricsResultItem MetricsGetEdgeCategoryMetrics(ctx, body)
Get flow metric aggregate data by application category

Gets flow metric summaries for the specified time interval by application category. On success, this method returns an array of flow data in which each entry corresponds to a category of application traffic that has traversed the specified Edge. In the request body, the `id` and `edgeId` property names are interchangeable. The `enterpriseId` property is required when this method is invoked in the operator context.  Privileges required:  `READ` `EDGE`  `VIEW_FLOW_STATS` `undefined`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MetricsGetEdgeCategoryMetrics**](MetricsGetEdgeCategoryMetrics.md)|  | 

### Return type

[**[]MetricsGetEdgeCategoryMetricsResultItem**](metrics_get_edge_category_metrics_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsGetEdgeCategorySeries

> []MetricsGetEdgeCategorySeriesResultItem MetricsGetEdgeCategorySeries(ctx, body)
Get flow metric time series data by application category

Gets flow metric time series data for the specified time interval by application category. On success, this method returns an array of flow data in which each entry corresponds to a category of application traffic that has traversed the specified Edge. In the request body, the `id` and `edgeId` property names are interchangeable. The `enterpriseId` property is required when this method is invoked in the operator context.  Privileges required:  `READ` `EDGE`  `VIEW_FLOW_STATS` `undefined`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MetricsGetEdgeCategorySeries**](MetricsGetEdgeCategorySeries.md)|  | 

### Return type

[**[]MetricsGetEdgeCategorySeriesResultItem**](metrics_get_edge_category_series_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsGetEdgeDestMetrics

> []MetricsGetEdgeDestMetricsResultItem MetricsGetEdgeDestMetrics(ctx, body)
Get flow metric aggregate data by destination

Gets flow metric summaries for the specified time interval by destination. On success, this method returns an array of flow data in which each entry corresponds to a distinct traffic destination. In the request body, the `id` and `edgeId` property names are interchangeable. The `enterpriseId` property is required when this method is invoked in the operator context.  Privileges required:  `READ` `EDGE`  `VIEW_FLOW_STATS` `undefined`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MetricsGetEdgeDestMetrics**](MetricsGetEdgeDestMetrics.md)|  | 

### Return type

[**[]MetricsGetEdgeDestMetricsResultItem**](metrics_get_edge_dest_metrics_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsGetEdgeDestSeries

> []MetricsGetEdgeDestSeriesResultItem MetricsGetEdgeDestSeries(ctx, body)
Get flow metric time series data by destination

Gets flow metric time series data for the specified time interval by destination. On success, this method returns an array of flow data in which each entry corresponds to a distinct traffic destination. In the request body, the `id` and `edgeId` property names are interchangeable. The `enterpriseId` property is required when this method is invoked in the operator context.  Privileges required:  `READ` `EDGE`  `VIEW_FLOW_STATS` `undefined`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MetricsGetEdgeDestSeries**](MetricsGetEdgeDestSeries.md)|  | 

### Return type

[**[]MetricsGetEdgeDestSeriesResultItem**](metrics_get_edge_dest_series_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsGetEdgeDeviceMetrics

> []MetricsGetEdgeDeviceMetricsResultItem MetricsGetEdgeDeviceMetrics(ctx, body)
Get flow metric aggregate data by client device

Gets flow metric summaries for the specified time interval by client device. On success, this method returns an array of flow data in which each entry corresponds to a distinct device. In the request body, the `id` and `edgeId` property names are interchangeable. The `enterpriseId` property is required when this method is invoked in the operator context.  Privileges required:  `READ` `EDGE`  `VIEW_USER_IDENTIFIABLE_FLOW_STATS` `undefined`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MetricsGetEdgeDeviceMetrics**](MetricsGetEdgeDeviceMetrics.md)|  | 

### Return type

[**[]MetricsGetEdgeDeviceMetricsResultItem**](metrics_get_edge_device_metrics_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsGetEdgeDeviceSeries

> []MetricsGetEdgeDeviceSeriesResultItem MetricsGetEdgeDeviceSeries(ctx, body)
Get flow metric time series data by client device

Gets flow metric time series data for the specified time interval by client device. On success, this method returns an array of flow data in which each entry corresponds to a distinct device. In the request body, the `id` and `edgeId` property names are interchangeable. The `enterpriseId` property is required when this method is invoked in the operator context.  Privileges required:  `READ` `EDGE`  `VIEW_USER_IDENTIFIABLE_FLOW_STATS` `undefined`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MetricsGetEdgeDeviceSeries**](MetricsGetEdgeDeviceSeries.md)|  | 

### Return type

[**[]MetricsGetEdgeDeviceSeriesResultItem**](metrics_get_edge_device_series_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsGetEdgeLinkMetrics

> []MetricsGetEdgeLinkMetricsResultItem MetricsGetEdgeLinkMetrics(ctx, body)
Get advanced flow metric aggregate data by link

Gets advanced flow metric summaries for the specified time interval by link. On success, this method returns an array of flow data in which each entry corresponds to a link on the specified Edge. In the request body, the `id` and `edgeId` property names are interchangeable. The `enterpriseId` property is required when this method is invoked in the operator context.  Privileges required:  `READ` `EDGE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MetricsGetEdgeLinkMetrics**](MetricsGetEdgeLinkMetrics.md)|  | 

### Return type

[**[]MetricsGetEdgeLinkMetricsResultItem**](metrics_get_edge_link_metrics_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsGetEdgeLinkSeries

> []MetricsGetEdgeLinkSeriesResultItem MetricsGetEdgeLinkSeries(ctx, body)
Get advanced flow metric time series data by link

Gets advanced flow metric time series data for the specified time interval by link. On success, this method returns an array of flow data in which each entry corresponds to a link on the specified Edge. In the request body, the `id` and `edgeId` property names are interchangeable. The `enterpriseId` property is required when this method is invoked in the operator context.  Privileges required:  `READ` `EDGE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MetricsGetEdgeLinkSeries**](MetricsGetEdgeLinkSeries.md)|  | 

### Return type

[**[]MetricsGetEdgeLinkSeriesResultItem**](metrics_get_edge_link_series_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsGetEdgeOsMetrics

> []MetricsGetEdgeOsMetricsResultItem MetricsGetEdgeOsMetrics(ctx, body)
Get flow metric aggregate data by client OS

Gets flow metric summaries for the specified time interval by client OS. On success, this method returns an array of flow data in which each entry corresponds to a distinct OS on a client device. In the request body, the `id` and `edgeId` property names are interchangeable. The `enterpriseId` property is required when this method is invoked in the operator context.  Privileges required:  `READ` `EDGE`  `VIEW_FLOW_STATS` `undefined`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MetricsGetEdgeOsMetrics**](MetricsGetEdgeOsMetrics.md)|  | 

### Return type

[**[]MetricsGetEdgeOsMetricsResultItem**](metrics_get_edge_os_metrics_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsGetEdgeOsSeries

> []MetricsGetEdgeOsSeriesResultItem MetricsGetEdgeOsSeries(ctx, body)
Get flow metric time series data by client OS

Gets flow metric time series data for the specified time interval by client OS. On success, this method returns an array of flow data in which each entry corresponds to a distinct OS on a client device. In the request body, the `id` and `edgeId` property names are interchangeable. The `enterpriseId` property is required when this method is invoked in the operator context.  Privileges required:  `READ` `EDGE`  `VIEW_FLOW_STATS` `undefined`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MetricsGetEdgeOsSeries**](MetricsGetEdgeOsSeries.md)|  | 

### Return type

[**[]MetricsGetEdgeOsSeriesResultItem**](metrics_get_edge_os_series_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsGetEdgeSegmentMetrics

> []MetricsGetEdgeSegmentMetricsResultItem MetricsGetEdgeSegmentMetrics(ctx, body)
Get flow metric aggregate data by segment Id

Gets flow metric summaries for the specified time interval by segment id. On success, this method returns an array of flow data where each entry corresponds to a segment id traffic that has traversed the specified Edge. In the request body, the `id` and `edgeId` property names are interchangeable. The `enterpriseId` property is required when this method is invoked in the operator context.  Privileges required:  `READ` `EDGE`  `VIEW_FLOW_STATS` `undefined`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MetricsGetEdgeSegmentMetrics**](MetricsGetEdgeSegmentMetrics.md)|  | 

### Return type

[**[]MetricsGetEdgeSegmentMetricsResultItem**](metrics_get_edge_segment_metrics_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsGetEdgeSegmentSeries

> []MetricsGetEdgeSegmentSeriesResultItem MetricsGetEdgeSegmentSeries(ctx, body)
Get flow metric time series data by segment id

Gets flow metric time series data for the specified time interval by segment id. On success, this method returns an array of flow data in which each entry corresponds to a segment id of traffic that has traversed the specified Edge. In the request body, the `id` and `edgeId` property names are interchangeable. The `enterpriseId` property is required when this method is invoked in the operator context.  Privileges required:  `READ` `EDGE`  `VIEW_FLOW_STATS` `undefined`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MetricsGetEdgeSegmentSeries**](MetricsGetEdgeSegmentSeries.md)|  | 

### Return type

[**[]MetricsGetEdgeSegmentSeriesResultItem**](metrics_get_edge_segment_series_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsGetEdgeStatusMetrics

> EdgeStatusMetricsSummary MetricsGetEdgeStatusMetrics(ctx, body)
Get Edge healthStats metrics for an interval

Fetch healthStats metric summaries for the given edge, time interval and list of metrics. On success, this method returns anhealthsStats object with min, max and average values of requested metrics  Privileges required:  `READ` `EDGE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MetricsGetEdgeStatusMetrics**](MetricsGetEdgeStatusMetrics.md)|  | 

### Return type

[**EdgeStatusMetricsSummary**](edge_status_metrics_summary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsGetEdgeStatusSeries

> GetEdgeStatusMetricsTimeSeriesResult MetricsGetEdgeStatusSeries(ctx, body)
Get Edge healthStats time series for an interval 

Fetch healthStats time series for the given edge, time interval and list of metrics. On success, this method returns anarray of healthsStats series of requested metrics  Privileges required:  `READ` `EDGE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MetricsGetEdgeStatusSeries**](MetricsGetEdgeStatusSeries.md)|  | 

### Return type

[**GetEdgeStatusMetricsTimeSeriesResult**](get_edge_status_metrics_time_series_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsGetGatewayStatusMetrics

> GatewayStatusMetricsSummary MetricsGetGatewayStatusMetrics(ctx, body)
Get Gateway health metric summaries for an interval

Fetch health metric summaries given a target gateway, time interval and list of metrics. On success, this method returns an object containing min, max and average values of requested metrics  Privileges required:  `READ` `GATEWAY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MetricsGetGatewayStatusMetrics**](MetricsGetGatewayStatusMetrics.md)|  | 

### Return type

[**GatewayStatusMetricsSummary**](gateway_status_metrics_summary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsGetGatewayStatusSeries

> GetGatewayStatusMetricsTimeSeriesResult MetricsGetGatewayStatusSeries(ctx, body)
Get Gateway health metric time series for an interval

Fetch health metric time series given a target gateway, time interval and list of metrics. On success, this method returns an array of time series, one per metric.  Privileges required:  `READ` `GATEWAY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MetricsGetGatewayStatusSeries**](MetricsGetGatewayStatusSeries.md)|  | 

### Return type

[**GetGatewayStatusMetricsTimeSeriesResult**](get_gateway_status_metrics_time_series_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitoringGetAggregateEdgeLinkMetrics

> []MonitoringGetAggregateEdgeLinkMetricsResultItem MonitoringGetAggregateEdgeLinkMetrics(ctx, body)
Get aggregate Edge link metrics across enterprises

Gets aggregate link metrics for the request interval for all active links across all enterprises, where a link is considered to be active if an Edge has reported any activity for it in the last 24 hours. On success, returns an array of network utilization metrics, one per link.  Privileges required:  `READ` `ENTERPRISE`  `READ` `EDGE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MonitoringGetAggregateEdgeLinkMetrics**](MonitoringGetAggregateEdgeLinkMetrics.md)|  | 

### Return type

[**[]MonitoringGetAggregateEdgeLinkMetricsResultItem**](monitoring_get_aggregate_edge_link_metrics_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitoringGetAggregateEnterpriseEvents

> MonitoringGetAggregateEnterpriseEventsResult MonitoringGetAggregateEnterpriseEvents(ctx, body)
Get events across all enterprises

Gets events across all enterprises in a paginated list. When called in the MSP/Partner context, queries only enterprises managed by the MSP.  Privileges required:  `READ` `ENTERPRISE`  `READ` `EDGE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MonitoringGetAggregateEnterpriseEvents**](MonitoringGetAggregateEnterpriseEvents.md)|  | 

### Return type

[**MonitoringGetAggregateEnterpriseEventsResult**](monitoring_get_aggregate_enterprise_events_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitoringGetAggregates

> MonitoringGetAggregatesResult MonitoringGetAggregates(ctx, body)
Get aggregate enterprise and Edge information

Gets a comprehensive listing of all enterprises and Edges on a network. Returns an object containing an aggregate `edgeCount`, a list (`enterprises`) containing enterprise objects, and a map (`edges`) that provides Edge counts per enterprise.  Privileges required:  `READ` `ENTERPRISE`  `READ` `EDGE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MonitoringGetAggregates**](MonitoringGetAggregates.md)|  | 

### Return type

[**MonitoringGetAggregatesResult**](monitoring_get_aggregates_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitoringGetEnterpriseBgpPeerStatus

> []MonitoringGetEnterpriseBgpPeerStatusResultItem MonitoringGetEnterpriseBgpPeerStatus(ctx, body)
Get gateway BGP peer status for all enterprise gateways

Gets the gateway BGP peer status for all enterprise gateways. Returns an array in which each entry corresponds to a gateway and contains an associated set of BGP peers with state records.  Privileges required:  `READ` `NETWORK_SERVICE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**InlineObject**](InlineObject.md)|  | 

### Return type

[**[]MonitoringGetEnterpriseBgpPeerStatusResultItem**](monitoring_get_enterprise_bgp_peer_status_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitoringGetEnterpriseEdgeBgpPeerStatus

> []MonitoringGetEnterpriseEdgeBgpPeerStatusResultItem MonitoringGetEnterpriseEdgeBgpPeerStatus(ctx, body)
Get Edge BGP peer status for all enterprise Edges

Gets the Edge BGP peer status for all enterprise Edges. Returns an array in which each entry corresponds to an Edge and contains an associated set of BGP peers and state records.  Privileges required:  `READ` `EDGE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**InlineObject1**](InlineObject1.md)|  | 

### Return type

[**[]MonitoringGetEnterpriseEdgeBgpPeerStatusResultItem**](monitoring_get_enterprise_edge_bgp_peer_status_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitoringGetEnterpriseEdgeClusterStatus

> []InlineResponse2001 MonitoringGetEnterpriseEdgeClusterStatus(ctx, body)
Get Edge Cluster status

Get edge`s utilization data by clusters  Privileges required:  `READ` `EDGE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MonitoringGetEnterpriseEdgeClusterStatus**](MonitoringGetEnterpriseEdgeClusterStatus.md)|  | 

### Return type

[**[]InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitoringGetEnterpriseEdgeLinkStatus

> []MonitoringGetEnterpriseEdgeLinkStatusResultItem MonitoringGetEnterpriseEdgeLinkStatus(ctx, body)
Get Edge and link status data

Gets the current Edge and Edge-link status for all enterprises. edges.  Privileges required:  `READ` `ENTERPRISE`  `READ` `EDGE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**InlineObject2**](InlineObject2.md)|  | 

### Return type

[**[]MonitoringGetEnterpriseEdgeLinkStatusResultItem**](monitoring_get_enterprise_edge_link_status_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitoringGetEnterpriseEdgeStatus

> MonitoringGetEnterpriseEdgeStatusResult MonitoringGetEnterpriseEdgeStatus(ctx, body)
Get Enterprise Edge monitoring status

Fetch Enterprise Edge monitoring status time series for the given time interval, all for the enterprise or list of edgess, and list of metrics. On success, this method returns anarray of edge monitoring statuses.  Privileges required:  `READ` `EDGE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MonitoringGetEnterpriseEdgeStatus**](MonitoringGetEnterpriseEdgeStatus.md)|  | 

### Return type

[**MonitoringGetEnterpriseEdgeStatusResult**](monitoring_get_enterprise_edge_status_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitoringGetNetworkGatewayStatus

> MonitoringGetNetworkGatewayStatusResult MonitoringGetNetworkGatewayStatus(ctx, body)
Get Network Gateway monitoring status

Fetch Network Gateway monitoring status time series for the given time interval, all for the network or list of gateways, and list of metrics. On success, this method returns anarray of gateway monitoring statuses.  Privileges required:  `READ` `GATEWAY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MonitoringGetNetworkGatewayStatus**](MonitoringGetNetworkGatewayStatus.md)|  | 

### Return type

[**MonitoringGetNetworkGatewayStatusResult**](monitoring_get_network_gateway_status_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkDeleteNetworkGatewayPool

> NetworkDeleteNetworkGatewayPoolResult NetworkDeleteNetworkGatewayPool(ctx, body)
Delete gateway pool

Deletes the specified gateway pool (by `id`).  Privileges required:  `DELETE` `GATEWAY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**NetworkDeleteNetworkGatewayPool**](NetworkDeleteNetworkGatewayPool.md)|  | 

### Return type

[**NetworkDeleteNetworkGatewayPoolResult**](network_delete_network_gateway_pool_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkGetNetworkConfigurations

> []NetworkGetNetworkConfigurationsResultItem NetworkGetNetworkConfigurations(ctx, body)
Get operator configuration profiles

Gets all operator configuration profiles associated with an operator's network. Optionally includes the modules associated with each profile. This call does not return templates.  Privileges required:  `READ` `OPERATOR_PROFILE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**NetworkGetNetworkConfigurations**](NetworkGetNetworkConfigurations.md)|  | 

### Return type

[**[]NetworkGetNetworkConfigurationsResultItem**](network_get_network_configurations_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkGetNetworkEnterprises

> []NetworkGetNetworkEnterprisesResultItem NetworkGetNetworkEnterprises(ctx, body)
Get enterprises on network

Gets all enterprises existing on a network, optionally including all Edges or Edge counts. The `edgeConfigUpdate` \"with\" option may also be passed to check whether the application of configuration updates to Edges is enabled for each enterprise.  Privileges required:  `READ` `ENTERPRISE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**NetworkGetNetworkEnterprises**](NetworkGetNetworkEnterprises.md)|  | 

### Return type

[**[]NetworkGetNetworkEnterprisesResultItem**](network_get_network_enterprises_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkGetNetworkGatewayPools

> []NetworkGetNetworkGatewayPoolsResultItem NetworkGetNetworkGatewayPools(ctx, body)
Get gateway pools on network

Gets all gateway pools associated with the specified network. Optionally includes the gateways or enterprises belonging to each pool.  Privileges required:  `READ` `GATEWAY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**NetworkGetNetworkGatewayPools**](NetworkGetNetworkGatewayPools.md)|  | 

### Return type

[**[]NetworkGetNetworkGatewayPoolsResultItem**](network_get_network_gateway_pools_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkGetNetworkGateways

> []NetworkGetNetworkGatewaysResultItem NetworkGetNetworkGateways(ctx, body)
Get gateways on network

Gets all gateways associated with the specified network.  Privileges required:  `READ` `GATEWAY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**NetworkGetNetworkGateways**](NetworkGetNetworkGateways.md)|  | 

### Return type

[**[]NetworkGetNetworkGatewaysResultItem**](network_get_network_gateways_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkGetNetworkOperatorUsers

> []NetworkGetNetworkOperatorUsersResultItem NetworkGetNetworkOperatorUsers(ctx, body)
Get operator users for network

Gets all operator users associated with the specified network.  Privileges required:  `READ` `OPERATOR_USER`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**NetworkGetNetworkOperatorUsers**](NetworkGetNetworkOperatorUsers.md)|  | 

### Return type

[**[]NetworkGetNetworkOperatorUsersResultItem**](network_get_network_operator_users_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkInsertNetworkGatewayPool

> NetworkInsertNetworkGatewayPoolResult NetworkInsertNetworkGatewayPool(ctx, body)
Create gateway pool

Creates a gateway pool for the specified network.  Privileges required:  `CREATE` `GATEWAY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**NetworkInsertNetworkGatewayPool**](NetworkInsertNetworkGatewayPool.md)|  | 

### Return type

[**NetworkInsertNetworkGatewayPoolResult**](network_insert_network_gateway_pool_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkUpdateNetworkGatewayPoolAttributes

> NetworkUpdateNetworkGatewayPoolAttributesResult NetworkUpdateNetworkGatewayPoolAttributes(ctx, body)
Update gateway pool attributes

Updates the configurable attributes (`name`, `description`, and `handOffType`) of the specified gateway pool (by `id` or `gatewayPoolId`).  Privileges required:  `UPDATE` `GATEWAY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**NetworkUpdateNetworkGatewayPoolAttributes**](NetworkUpdateNetworkGatewayPoolAttributes.md)|  | 

### Return type

[**NetworkUpdateNetworkGatewayPoolAttributesResult**](network_update_network_gateway_pool_attributes_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperatorUserDeleteOperatorUser

> OperatorUserDeleteOperatorUserResult OperatorUserDeleteOperatorUser(ctx, body)
Delete operator user

Deletes the specified operator user (by `id` or `username`).  Privileges required:  `DELETE` `OPERATOR_USER`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**OperatorUserDeleteOperatorUser**](OperatorUserDeleteOperatorUser.md)|  | 

### Return type

[**OperatorUserDeleteOperatorUserResult**](operator_user_delete_operator_user_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperatorUserGetOperatorUser

> OperatorUserGetOperatorUserResult OperatorUserGetOperatorUser(ctx, body)
Get operator user

Gets the specified operator user object (by `id` or `username`).  Privileges required:  `READ` `OPERATOR_USER`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**OperatorUserGetOperatorUser**](OperatorUserGetOperatorUser.md)|  | 

### Return type

[**OperatorUserGetOperatorUserResult**](operator_user_get_operator_user_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperatorUserInsertOperatorUser

> OperatorUserGetOperatorUserResult OperatorUserInsertOperatorUser(ctx, body)
Create operator user

Creates an operator user and associates it with the operator's network.  Privileges required:  `CREATE` `OPERATOR_USER`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**OperatorUserInsertOperatorUser**](OperatorUserInsertOperatorUser.md)|  | 

### Return type

[**OperatorUserGetOperatorUserResult**](operator_user_get_operator_user_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperatorUserUpdateOperatorUser

> OperatorUserUpdateOperatorUserResult OperatorUserUpdateOperatorUser(ctx, body)
Update operator user

Updates the specified operator user (by object `id` or `username`). Expects an `_update` object containing the field(s) to be updated and their corresponding value(s).  Privileges required:  `UPDATE` `OPERATOR_USER`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**OperatorUserUpdateOperatorUser1**](OperatorUserUpdateOperatorUser1.md)|  | 

### Return type

[**OperatorUserUpdateOperatorUserResult**](operator_user_update_operator_user_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RoleCreateRoleCustomization

> RoleCreateRoleCustomizationResult RoleCreateRoleCustomization(ctx, body)
Create role customization

Creates a role customization specified by `roleId` and an array of `privilegeIds`.  Privileges required:  `UPDATE` `NETWORK`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**RoleCreateRoleCustomization**](RoleCreateRoleCustomization.md)|  | 

### Return type

[**RoleCreateRoleCustomizationResult**](role_create_role_customization_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RoleDeleteRoleCustomization

> RoleDeleteRoleCustomizationResult RoleDeleteRoleCustomization(ctx, body)
Delete role customization

Deletes the specified role customization (by `name` or `forRoleId`).  Privileges required:  `UPDATE` `NETWORK`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**RoleDeleteRoleCustomization**](RoleDeleteRoleCustomization.md)|  | 

### Return type

[**RoleDeleteRoleCustomizationResult**](role_delete_role_customization_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RoleGetUserTypeRoles

> []RoleGetUserTypeRolesResultItem RoleGetUserTypeRoles(ctx, body)
Get roles per user type

Gets all roles defined for the specified `userType`.  Privileges required:  `READ` `ENTERPRISE_USER`, or  `READ` `PROXY_USER`, or  `READ` `OPERATOR_USER`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**RoleGetUserTypeRoles**](RoleGetUserTypeRoles.md)|  | 

### Return type

[**[]RoleGetUserTypeRolesResultItem**](role_get_user_type_roles_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RoleSetEnterpriseDelegatedToEnterpriseProxy

> RoleSetEnterpriseDelegatedToEnterpriseProxyResult RoleSetEnterpriseDelegatedToEnterpriseProxy(ctx, body)
Grant enterprise access to partner

Grants enterprise access to the specified enterprise proxy (partner). When an enterprise is delegated to a proxy, proxy users are granted access to view, configure, and troubleshoot Edges owned by the enterprise. As a security consideration, proxy Support users cannot view personally identifiable information.  Privileges required:  `UPDATE` `ENTERPRISE_DELEGATION`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**InlineObject3**](InlineObject3.md)|  | 

### Return type

[**RoleSetEnterpriseDelegatedToEnterpriseProxyResult**](role_set_enterprise_delegated_to_enterprise_proxy_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RoleSetEnterpriseDelegatedToOperator

> RoleSetEnterpriseDelegatedToOperatorResult RoleSetEnterpriseDelegatedToOperator(ctx, body)
Grant enterprise access to network operator

Grants enterprise access to the network operator. When an enterprise is delegated to the operator, operator users are permitted to view, configure, and troubleshoot Edges owned by the enterprise. As a security consideration, operator users cannot view personally identifiable information.  Privileges required:  `UPDATE` `ENTERPRISE_DELEGATION`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**InlineObject4**](InlineObject4.md)|  | 

### Return type

[**RoleSetEnterpriseDelegatedToOperatorResult**](role_set_enterprise_delegated_to_operator_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RoleSetEnterpriseProxyDelegatedToOperator

> RoleSetEnterpriseProxyDelegatedToOperatorResult RoleSetEnterpriseProxyDelegatedToOperator(ctx, body)
Grant enterprise proxy access to network operator

Grants enterprise proxy access to the network operator. When an enterprise proxy is delegated to the operator, operator users are granted access to view, configure and troubleshoot objects owned by the proxy.  Privileges required:  `UPDATE` `ENTERPRISE_PROXY_DELEGATION`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**InlineObject5**](InlineObject5.md)|  | 

### Return type

[**RoleSetEnterpriseProxyDelegatedToOperatorResult**](role_set_enterprise_proxy_delegated_to_operator_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RoleSetEnterpriseUserManagementDelegatedToOperator

> RoleSetEnterpriseUserManagementDelegatedToOperatorResult RoleSetEnterpriseUserManagementDelegatedToOperator(ctx, body)
Grant enterprise user access to network operator

Grants enterprise user access to the specified network operator. When enterprise user management is delegated to the operator, operator users are granted enterprise-level user management capabilities (user creation, password resets, etc.).  Privileges required:  `UPDATE` `ENTERPRISE_DELEGATION`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**InlineObject6**](InlineObject6.md)|  | 

### Return type

[**RoleSetEnterpriseUserManagementDelegatedToOperatorResult**](role_set_enterprise_user_management_delegated_to_operator_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetClientDeviceHostName

> SetClientDeviceHostNameResult SetClientDeviceHostName(ctx, body)
Set hostname for client device

Sets the `hostName` for client device  Privileges required:  `UPDATE` `CLIENT_DEVICE`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**SetClientDeviceHostName**](SetClientDeviceHostName.md)|  | 

### Return type

[**SetClientDeviceHostNameResult**](set_client_device_host_name_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemPropertyGetSystemProperties

> []SystemPropertyGetSystemPropertiesResultItem SystemPropertyGetSystemProperties(ctx, body)
Get all system properties

Gets all configured system properties.  Privileges required:  `READ` `SYSTEM_PROPERTY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**SystemPropertyGetSystemProperties**](SystemPropertyGetSystemProperties.md)|  | 

### Return type

[**[]SystemPropertyGetSystemPropertiesResultItem**](system_property_get_system_properties_result_item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemPropertyGetSystemProperty

> SystemPropertyGetSystemPropertyResult SystemPropertyGetSystemProperty(ctx, body)
Get system property

Gets the specified system property. Specify by object `id` or `name`.  Privileges required:  `READ` `SYSTEM_PROPERTY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**SystemPropertyGetSystemProperty**](SystemPropertyGetSystemProperty.md)|  | 

### Return type

[**SystemPropertyGetSystemPropertyResult**](system_property_get_system_property_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemPropertyInsertOrUpdateSystemProperty

> SystemPropertyInsertOrUpdateSystemPropertyResult SystemPropertyInsertOrUpdateSystemProperty(ctx, body)
Create or update system property

Creates the specified system property. If the system property with the specified `name` already exists, then the API call updates it.  Privileges required:  `CREATE` `SYSTEM_PROPERTY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**SystemPropertyInsertOrUpdateSystemProperty**](SystemPropertyInsertOrUpdateSystemProperty.md)|  | 

### Return type

[**SystemPropertyInsertOrUpdateSystemPropertyResult**](system_property_insert_or_update_system_property_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemPropertyInsertSystemProperty

> SystemPropertyInsertSystemPropertyResult SystemPropertyInsertSystemProperty(ctx, body)
Create system property

Creates a new system property with the specified `name`, `value`, and optionally other attributes.  Privileges required:  `CREATE` `SYSTEM_PROPERTY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**SystemPropertyInsertSystemProperty**](SystemPropertyInsertSystemProperty.md)|  | 

### Return type

[**SystemPropertyInsertSystemPropertyResult**](system_property_insert_system_property_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemPropertyUpdateSystemProperty

> SystemPropertyUpdateSystemPropertyResult SystemPropertyUpdateSystemProperty(ctx, body)
Update system property

Updates the specified system property (by object `id` or `name`). Expects an `_update` object containing the field(s) to be updated and their corresponding value(s).  Privileges required:  `UPDATE` `SYSTEM_PROPERTY`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**SystemPropertyUpdateSystemProperty**](SystemPropertyUpdateSystemProperty.md)|  | 

### Return type

[**SystemPropertyUpdateSystemPropertyResult**](system_property_update_system_property_result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VcoDiagnosticsGetVcoDbDiagnostics

> []InlineResponse2002 VcoDiagnosticsGetVcoDbDiagnostics(ctx, body)
Get VCO Database Diagnostics

Gets the diagnostic information of the VCO database.  Privileges required:  `READ` `VCO_DIAGNOSTICS`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | **map[string]interface{}**|  | 

### Return type

[**[]InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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

