// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ManagementPlaneAPINetworkingServicesLoadbalancer
// Used by client-side stubs.

package api

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-mp/model"
)

const _ = core.SupportedByRuntimeVersion1

type ManagementPlaneAPINetworkingServicesLoadbalancerClient interface {

	// Create a load balancer application profile.
	//
	// @param lbAppProfileParam (required)
	// The parameter must contain all the properties defined in model.LbAppProfile.
	// @return com.vmware.model.LbAppProfile
	// The return value will contain all the properties defined in model.LbAppProfile.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CreateLoadBalancerApplicationProfile(lbAppProfileParam *data.StructValue) (*data.StructValue, error)

	// Create a load balancer client-ssl profile.
	//
	// @param lbClientSslProfileParam (required)
	// @return com.vmware.model.LbClientSslProfile
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CreateLoadBalancerClientSslProfile(lbClientSslProfileParam model.LbClientSslProfile) (model.LbClientSslProfile, error)

	// Create a load balancer monitor.
	//
	// @param lbMonitorParam (required)
	// The parameter must contain all the properties defined in model.LbMonitor.
	// @return com.vmware.model.LbMonitor
	// The return value will contain all the properties defined in model.LbMonitor.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CreateLoadBalancerMonitor(lbMonitorParam *data.StructValue) (*data.StructValue, error)

	// Create a load balancer persistence profile.
	//
	// @param lbPersistenceProfileParam (required)
	// The parameter must contain all the properties defined in model.LbPersistenceProfile.
	// @return com.vmware.model.LbPersistenceProfile
	// The return value will contain all the properties defined in model.LbPersistenceProfile.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CreateLoadBalancerPersistenceProfile(lbPersistenceProfileParam *data.StructValue) (*data.StructValue, error)

	// Create a load balancer pool.
	//
	// @param lbPoolParam (required)
	// @return com.vmware.model.LbPool
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CreateLoadBalancerPool(lbPoolParam model.LbPool) (model.LbPool, error)

	// Create a load balancer rule.
	//
	// @param lbRuleParam (required)
	// @return com.vmware.model.LbRule
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CreateLoadBalancerRule(lbRuleParam model.LbRule) (model.LbRule, error)

	// Create a load balancer server-ssl profile.
	//
	// @param lbServerSslProfileParam (required)
	// @return com.vmware.model.LbServerSslProfile
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CreateLoadBalancerServerSslProfile(lbServerSslProfileParam model.LbServerSslProfile) (model.LbServerSslProfile, error)

	// Create a load balancer service.
	//
	// @param lbServiceParam (required)
	// @return com.vmware.model.LbService
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CreateLoadBalancerService(lbServiceParam model.LbService) (model.LbService, error)

	// Create a load balancer virtual server.
	//
	// @param lbVirtualServerParam (required)
	// @return com.vmware.model.LbVirtualServer
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CreateLoadBalancerVirtualServer(lbVirtualServerParam model.LbVirtualServer) (model.LbVirtualServer, error)

	// It is used to create virtual servers, the associated rules and bind the rules to the virtual server. To add new rules, make sure the rules which have no identifier specified, the new rules are automatically generated and associated to the virtual server. If the virtual server need to consume some existed rules without change, those rules should not be specified in this array, otherwise, the rules are updated.
	//
	// @param lbVirtualServerWithRuleParam (required)
	// @return com.vmware.model.LbVirtualServerWithRule
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CreateLoadBalancerVirtualServerWithRulesCreateWithRules(lbVirtualServerWithRuleParam model.LbVirtualServerWithRule) (model.LbVirtualServerWithRule, error)

	// Delete a load balancer application profile.
	//
	// @param applicationProfileIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteLoadBalancerApplicationProfile(applicationProfileIdParam string) error

	// Delete a load balancer client-ssl profile.
	//
	// @param clientSslProfileIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteLoadBalancerClientSslProfile(clientSslProfileIdParam string) error

	// Delete a load balancer monitor.
	//
	// @param monitorIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteLoadBalancerMonitor(monitorIdParam string) error

	// Delete a load balancer persistence profile.
	//
	// @param persistenceProfileIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteLoadBalancerPersistenceProfile(persistenceProfileIdParam string) error

	// Delete a load balancer pool.
	//
	// @param poolIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteLoadBalancerPool(poolIdParam string) error

	// Delete a load balancer rule.
	//
	// @param ruleIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteLoadBalancerRule(ruleIdParam string) error

	// Delete a load balancer server-ssl profile.
	//
	// @param serverSslProfileIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteLoadBalancerServerSslProfile(serverSslProfileIdParam string) error

	// Delete a load balancer service.
	//
	// @param serviceIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteLoadBalancerService(serviceIdParam string) error

	// Delete a load balancer virtual server.
	//
	// @param virtualServerIdParam (required)
	// @param deleteAssociatedRulesParam Delete associated rules (optional, default to false)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteLoadBalancerVirtualServer(virtualServerIdParam string, deleteAssociatedRulesParam *bool) error

	// Returns the statistics of the given load balancer pool by given load balancer serives id and load balancer pool id. Currently, only realtime mode is supported.
	//
	// @param serviceIdParam (required)
	// @param poolIdParam (required)
	// @param sourceParam Data source type. (optional)
	// @return com.vmware.model.LbPoolStatistics
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetLoadBalancerPoolStatistics(serviceIdParam string, poolIdParam string, sourceParam *string) (model.LbPoolStatistics, error)

	// Returns the status of the given load balancer pool by given load balancer serives id and load balancer pool id.
	//
	// @param serviceIdParam (required)
	// @param poolIdParam (required)
	// @param sourceParam Data source type. (optional)
	// @return com.vmware.model.LbPoolStatus
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetLoadBalancerPoolStatus(serviceIdParam string, poolIdParam string, sourceParam *string) (model.LbPoolStatus, error)

	// Returns the statistics of the given load balancer service.
	//
	// @param serviceIdParam (required)
	// @param sourceParam Data source type. (optional)
	// @return com.vmware.model.LbServiceStatistics
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetLoadBalancerServiceStatistics(serviceIdParam string, sourceParam *string) (model.LbServiceStatistics, error)

	// Returns the status of the given load balancer service.
	//
	// @param serviceIdParam (required)
	// @param includeInstanceDetailsParam Flag to indicate whether include detail information (optional, default to false)
	// @param sourceParam Data source type. (optional)
	// @param transportNodeIdsParam The UUIDs of transport nodes (optional)
	// @return com.vmware.model.LbServiceStatus
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetLoadBalancerServiceStatus(serviceIdParam string, includeInstanceDetailsParam *bool, sourceParam *string, transportNodeIdsParam *string) (model.LbServiceStatus, error)

	// Returns the statistics of the load balancer virtual server by given load balancer serives id and load balancer virtual server id.
	//
	// @param serviceIdParam (required)
	// @param virtualServerIdParam (required)
	// @param sourceParam Data source type. (optional)
	// @return com.vmware.model.LbVirtualServerStatistics
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetLoadBalancerVirtualServerStatistics(serviceIdParam string, virtualServerIdParam string, sourceParam *string) (model.LbVirtualServerStatistics, error)

	// Returns the status of the virtual server by given load balancer serives id and load balancer virtual server id.
	//
	// @param serviceIdParam (required)
	// @param virtualServerIdParam (required)
	// @param sourceParam Data source type. (optional)
	// @return com.vmware.model.LbVirtualServerStatus
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetLoadBalancerVirtualServerStatus(serviceIdParam string, virtualServerIdParam string, sourceParam *string) (model.LbVirtualServerStatus, error)

	// Retrieve a paginated list of load balancer application profiles.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param type_Param application profile type (optional)
	// @return com.vmware.model.LbAppProfileListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListLoadBalancerApplicationProfiles(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, type_Param *string) (model.LbAppProfileListResult, error)

	// Retrieve a paginated list of load balancer client-ssl profiles.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.model.LbClientSslProfileListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListLoadBalancerClientSslProfiles(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.LbClientSslProfileListResult, error)

	// Retrieve a paginated list of load balancer monitors.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param type_Param monitor query type (optional)
	// @return com.vmware.model.LbMonitorListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListLoadBalancerMonitors(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, type_Param *string) (model.LbMonitorListResult, error)

	// Retrieve a paginated list of load balancer persistence profiles.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param type_Param persistence profile type (optional)
	// @return com.vmware.model.LbPersistenceProfileListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListLoadBalancerPersistenceProfiles(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, type_Param *string) (model.LbPersistenceProfileListResult, error)

	// Returns the statistics list of load balancer pools in given load balancer service. Currently, only realtime mode is supported.
	//
	// @param serviceIdParam (required)
	// @param sourceParam Data source type. (optional)
	// @return com.vmware.model.LbPoolStatisticsListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListLoadBalancerPoolStatistics(serviceIdParam string, sourceParam *string) (model.LbPoolStatisticsListResult, error)

	// Returns the status list of load balancer pools in given load balancer service.
	//
	// @param serviceIdParam (required)
	// @param sourceParam Data source type. (optional)
	// @return com.vmware.model.LbPoolStatusListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListLoadBalancerPoolStatuses(serviceIdParam string, sourceParam *string) (model.LbPoolStatusListResult, error)

	// Retrieve a paginated list of load balancer pools.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.model.LbPoolListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListLoadBalancerPools(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.LbPoolListResult, error)

	// Retrieve a paginated list of load balancer rules.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.model.LbRuleListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListLoadBalancerRules(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.LbRuleListResult, error)

	// Retrieve a paginated list of load balancer server-ssl profiles.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.model.LbServerSslProfileListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListLoadBalancerServerSslProfiles(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.LbServerSslProfileListResult, error)

	// Retrieve a paginated list of load balancer services. When logical_router_id is specified in request parameters, the associated load balancer services which are related to the given logical router returned.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param logicalRouterIdParam Logical router identifier (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.model.LbServiceListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListLoadBalancerServices(cursorParam *string, includedFieldsParam *string, logicalRouterIdParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.LbServiceListResult, error)

	// Retrieve a list of supported SSL ciphers and protocols.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.model.LbSslCipherAndProtocolListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListLoadBalancerSslCiphersAndProtocols(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.LbSslCipherAndProtocolListResult, error)

	// Returns the status list of virtual servers in given load balancer service.
	//
	// @param serviceIdParam (required)
	// @param sourceParam Data source type. (optional)
	// @return com.vmware.model.LbVirtualServerStatusListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListLoadBalancerVirtualServerStatuses(serviceIdParam string, sourceParam *string) (model.LbVirtualServerStatusListResult, error)

	// Retrieve a paginated list of load balancer virtual servers.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.model.LbVirtualServerListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListLoadBalancerVirtualServers(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.LbVirtualServerListResult, error)

	// Returns the statistics list of virtual servers in given load balancer service. Currently, only realtime mode is supported.
	//
	// @param serviceIdParam (required)
	// @param sourceParam Data source type. (optional)
	// @return com.vmware.model.LbVirtualServerStatisticsListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListLoadBalancerVirtualServersStatistics(serviceIdParam string, sourceParam *string) (model.LbVirtualServerStatisticsListResult, error)

	// For ADD_MEMBERS, pool members will be created and added to load balancer pool. This action is only valid for static pool members. For REMOVE_MEMBERS, pool members will be removed from load balancer pool via IP and port in pool member settings. This action is only valid for static pool members. For UPDATE_MEMBERS, pool members admin state will be updated. This action is valid for both static pool members and dynamic pool members. For dynamic pool members, this update will be stored in customized_members field in load balancer pool member group.
	//
	// @param poolIdParam (required)
	// @param poolMemberSettingListParam (required)
	// @param actionParam Specifies addition, removal and modification action (required)
	// @return com.vmware.model.LbPool
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	PerformPoolMemberAction(poolIdParam string, poolMemberSettingListParam model.PoolMemberSettingList, actionParam string) (model.LbPool, error)

	// Retrieve a load balancer application profile.
	//
	// @param applicationProfileIdParam (required)
	// @return com.vmware.model.LbAppProfile
	// The return value will contain all the properties defined in model.LbAppProfile.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReadLoadBalancerApplicationProfile(applicationProfileIdParam string) (*data.StructValue, error)

	// Retrieve a load balancer client-ssl profile.
	//
	// @param clientSslProfileIdParam (required)
	// @return com.vmware.model.LbClientSslProfile
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReadLoadBalancerClientSslProfile(clientSslProfileIdParam string) (model.LbClientSslProfile, error)

	// Retrieve a load balancer monitor.
	//
	// @param monitorIdParam (required)
	// @return com.vmware.model.LbMonitor
	// The return value will contain all the properties defined in model.LbMonitor.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReadLoadBalancerMonitor(monitorIdParam string) (*data.StructValue, error)

	// API is used to retrieve the usage of load balancer entities which include current number and remaining number of credits, virtual Servers, pools, pool Members and different size of LB services from the given node. Currently only Edge node is supported.
	//
	// @param nodeIdParam (required)
	// @return com.vmware.model.LbNodeUsage
	// The return value will contain all the properties defined in model.LbNodeUsage.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReadLoadBalancerNodeUsage(nodeIdParam string) (*data.StructValue, error)

	// API is used to retrieve the load balancer node usage summary for all nodes.
	//
	// @param includeUsagesParam Whether to include node usages (optional)
	// @return com.vmware.model.LbNodeUsageSummary
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReadLoadBalancerNodeUsageSummary(includeUsagesParam *bool) (model.LbNodeUsageSummary, error)

	// Retrieve a load balancer persistence profile.
	//
	// @param persistenceProfileIdParam (required)
	// @return com.vmware.model.LbPersistenceProfile
	// The return value will contain all the properties defined in model.LbPersistenceProfile.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReadLoadBalancerPersistenceProfile(persistenceProfileIdParam string) (*data.StructValue, error)

	// Retrieve a load balancer pool.
	//
	// @param poolIdParam (required)
	// @return com.vmware.model.LbPool
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReadLoadBalancerPool(poolIdParam string) (model.LbPool, error)

	// Retrieve a load balancer rule.
	//
	// @param ruleIdParam (required)
	// @return com.vmware.model.LbRule
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReadLoadBalancerRule(ruleIdParam string) (model.LbRule, error)

	// Retrieve a load balancer server-ssl profile.
	//
	// @param serverSslProfileIdParam (required)
	// @return com.vmware.model.LbServerSslProfile
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReadLoadBalancerServerSslProfile(serverSslProfileIdParam string) (model.LbServerSslProfile, error)

	// Retrieve a load balancer service.
	//
	// @param serviceIdParam (required)
	// @return com.vmware.model.LbService
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReadLoadBalancerService(serviceIdParam string) (model.LbService, error)

	// API to download below information which will be used for debugging and troubleshooting. 1) Load balancer service 2) Load balancer associated virtual servers 3) Load balancer associated pools 4) Load balancer associated profiles such as persistence, SSL, application. 5) Load balancer associated monitors 6) Load balancer associated rules
	//
	// @param serviceIdParam (required)
	// @return com.vmware.model.LbServiceDebugInfo
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReadLoadBalancerServiceDebugInfo(serviceIdParam string) (model.LbServiceDebugInfo, error)

	// API to fetch the capacity and current usage of the given load balancer service.
	//
	// @param serviceIdParam (required)
	// @return com.vmware.model.LbServiceUsage
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReadLoadBalancerServiceUsage(serviceIdParam string) (model.LbServiceUsage, error)

	// Retrieve a load balancer virtual server.
	//
	// @param virtualServerIdParam (required)
	// @return com.vmware.model.LbVirtualServer
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReadLoadBalancerVirtualServer(virtualServerIdParam string) (model.LbVirtualServer, error)

	// Update a load balancer application profile.
	//
	// @param applicationProfileIdParam (required)
	// @param lbAppProfileParam (required)
	// The parameter must contain all the properties defined in model.LbAppProfile.
	// @return com.vmware.model.LbAppProfile
	// The return value will contain all the properties defined in model.LbAppProfile.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateLoadBalancerApplicationProfile(applicationProfileIdParam string, lbAppProfileParam *data.StructValue) (*data.StructValue, error)

	// Update a load balancer client-ssl profile.
	//
	// @param clientSslProfileIdParam (required)
	// @param lbClientSslProfileParam (required)
	// @return com.vmware.model.LbClientSslProfile
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateLoadBalancerClientSslProfile(clientSslProfileIdParam string, lbClientSslProfileParam model.LbClientSslProfile) (model.LbClientSslProfile, error)

	// Update a load balancer monitor.
	//
	// @param monitorIdParam (required)
	// @param lbMonitorParam (required)
	// The parameter must contain all the properties defined in model.LbMonitor.
	// @return com.vmware.model.LbMonitor
	// The return value will contain all the properties defined in model.LbMonitor.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateLoadBalancerMonitor(monitorIdParam string, lbMonitorParam *data.StructValue) (*data.StructValue, error)

	// Update a load balancer persistence profile.
	//
	// @param persistenceProfileIdParam (required)
	// @param lbPersistenceProfileParam (required)
	// The parameter must contain all the properties defined in model.LbPersistenceProfile.
	// @return com.vmware.model.LbPersistenceProfile
	// The return value will contain all the properties defined in model.LbPersistenceProfile.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateLoadBalancerPersistenceProfile(persistenceProfileIdParam string, lbPersistenceProfileParam *data.StructValue) (*data.StructValue, error)

	// Update a load balancer pool.
	//
	// @param poolIdParam (required)
	// @param lbPoolParam (required)
	// @return com.vmware.model.LbPool
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateLoadBalancerPool(poolIdParam string, lbPoolParam model.LbPool) (model.LbPool, error)

	// Update a load balancer rule.
	//
	// @param ruleIdParam (required)
	// @param lbRuleParam (required)
	// @return com.vmware.model.LbRule
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateLoadBalancerRule(ruleIdParam string, lbRuleParam model.LbRule) (model.LbRule, error)

	// Update a load balancer server-ssl profile.
	//
	// @param serverSslProfileIdParam (required)
	// @param lbServerSslProfileParam (required)
	// @return com.vmware.model.LbServerSslProfile
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateLoadBalancerServerSslProfile(serverSslProfileIdParam string, lbServerSslProfileParam model.LbServerSslProfile) (model.LbServerSslProfile, error)

	// Update a load balancer service.
	//
	// @param serviceIdParam (required)
	// @param lbServiceParam (required)
	// @return com.vmware.model.LbService
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateLoadBalancerService(serviceIdParam string, lbServiceParam model.LbService) (model.LbService, error)

	// Update a load balancer virtual server.
	//
	// @param virtualServerIdParam (required)
	// @param lbVirtualServerParam (required)
	// @return com.vmware.model.LbVirtualServer
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateLoadBalancerVirtualServer(virtualServerIdParam string, lbVirtualServerParam model.LbVirtualServer) (model.LbVirtualServer, error)

	// It is used to update virtual servers, the associated rules and update the binding of virtual server and rules. To add new rules, make sure the rules which have no identifier specified, the new rules are automatically generated and associated to the virtual server. To delete old rules, the rules should not be configured in new action, the UUID of deleted rules should be also removed from rule_ids. To update rules, the rules should be specified with new change and configured with identifier. If there are some rules which are not modified, those rule should not be specified in the rules list, the UUID list of rules should be specified in rule_ids of LbVirtualServer.
	//
	// @param virtualServerIdParam (required)
	// @param lbVirtualServerWithRuleParam (required)
	// @return com.vmware.model.LbVirtualServerWithRule
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateLoadBalancerVirtualServerWithRulesUpdateWithRules(virtualServerIdParam string, lbVirtualServerWithRuleParam model.LbVirtualServerWithRule) (model.LbVirtualServerWithRule, error)
}

type managementPlaneAPINetworkingServicesLoadbalancerClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewManagementPlaneAPINetworkingServicesLoadbalancerClient(connector client.Connector) *managementPlaneAPINetworkingServicesLoadbalancerClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.management_plane_API_networking_services_loadbalancer")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create_load_balancer_application_profile":                         core.NewMethodIdentifier(interfaceIdentifier, "create_load_balancer_application_profile"),
		"create_load_balancer_client_ssl_profile":                          core.NewMethodIdentifier(interfaceIdentifier, "create_load_balancer_client_ssl_profile"),
		"create_load_balancer_monitor":                                     core.NewMethodIdentifier(interfaceIdentifier, "create_load_balancer_monitor"),
		"create_load_balancer_persistence_profile":                         core.NewMethodIdentifier(interfaceIdentifier, "create_load_balancer_persistence_profile"),
		"create_load_balancer_pool":                                        core.NewMethodIdentifier(interfaceIdentifier, "create_load_balancer_pool"),
		"create_load_balancer_rule":                                        core.NewMethodIdentifier(interfaceIdentifier, "create_load_balancer_rule"),
		"create_load_balancer_server_ssl_profile":                          core.NewMethodIdentifier(interfaceIdentifier, "create_load_balancer_server_ssl_profile"),
		"create_load_balancer_service":                                     core.NewMethodIdentifier(interfaceIdentifier, "create_load_balancer_service"),
		"create_load_balancer_virtual_server":                              core.NewMethodIdentifier(interfaceIdentifier, "create_load_balancer_virtual_server"),
		"create_load_balancer_virtual_server_with_rules_create_with_rules": core.NewMethodIdentifier(interfaceIdentifier, "create_load_balancer_virtual_server_with_rules_create_with_rules"),
		"delete_load_balancer_application_profile":                         core.NewMethodIdentifier(interfaceIdentifier, "delete_load_balancer_application_profile"),
		"delete_load_balancer_client_ssl_profile":                          core.NewMethodIdentifier(interfaceIdentifier, "delete_load_balancer_client_ssl_profile"),
		"delete_load_balancer_monitor":                                     core.NewMethodIdentifier(interfaceIdentifier, "delete_load_balancer_monitor"),
		"delete_load_balancer_persistence_profile":                         core.NewMethodIdentifier(interfaceIdentifier, "delete_load_balancer_persistence_profile"),
		"delete_load_balancer_pool":                                        core.NewMethodIdentifier(interfaceIdentifier, "delete_load_balancer_pool"),
		"delete_load_balancer_rule":                                        core.NewMethodIdentifier(interfaceIdentifier, "delete_load_balancer_rule"),
		"delete_load_balancer_server_ssl_profile":                          core.NewMethodIdentifier(interfaceIdentifier, "delete_load_balancer_server_ssl_profile"),
		"delete_load_balancer_service":                                     core.NewMethodIdentifier(interfaceIdentifier, "delete_load_balancer_service"),
		"delete_load_balancer_virtual_server":                              core.NewMethodIdentifier(interfaceIdentifier, "delete_load_balancer_virtual_server"),
		"get_load_balancer_pool_statistics":                                core.NewMethodIdentifier(interfaceIdentifier, "get_load_balancer_pool_statistics"),
		"get_load_balancer_pool_status":                                    core.NewMethodIdentifier(interfaceIdentifier, "get_load_balancer_pool_status"),
		"get_load_balancer_service_statistics":                             core.NewMethodIdentifier(interfaceIdentifier, "get_load_balancer_service_statistics"),
		"get_load_balancer_service_status":                                 core.NewMethodIdentifier(interfaceIdentifier, "get_load_balancer_service_status"),
		"get_load_balancer_virtual_server_statistics":                      core.NewMethodIdentifier(interfaceIdentifier, "get_load_balancer_virtual_server_statistics"),
		"get_load_balancer_virtual_server_status":                          core.NewMethodIdentifier(interfaceIdentifier, "get_load_balancer_virtual_server_status"),
		"list_load_balancer_application_profiles":                          core.NewMethodIdentifier(interfaceIdentifier, "list_load_balancer_application_profiles"),
		"list_load_balancer_client_ssl_profiles":                           core.NewMethodIdentifier(interfaceIdentifier, "list_load_balancer_client_ssl_profiles"),
		"list_load_balancer_monitors":                                      core.NewMethodIdentifier(interfaceIdentifier, "list_load_balancer_monitors"),
		"list_load_balancer_persistence_profiles":                          core.NewMethodIdentifier(interfaceIdentifier, "list_load_balancer_persistence_profiles"),
		"list_load_balancer_pool_statistics":                               core.NewMethodIdentifier(interfaceIdentifier, "list_load_balancer_pool_statistics"),
		"list_load_balancer_pool_statuses":                                 core.NewMethodIdentifier(interfaceIdentifier, "list_load_balancer_pool_statuses"),
		"list_load_balancer_pools":                                         core.NewMethodIdentifier(interfaceIdentifier, "list_load_balancer_pools"),
		"list_load_balancer_rules":                                         core.NewMethodIdentifier(interfaceIdentifier, "list_load_balancer_rules"),
		"list_load_balancer_server_ssl_profiles":                           core.NewMethodIdentifier(interfaceIdentifier, "list_load_balancer_server_ssl_profiles"),
		"list_load_balancer_services":                                      core.NewMethodIdentifier(interfaceIdentifier, "list_load_balancer_services"),
		"list_load_balancer_ssl_ciphers_and_protocols":                     core.NewMethodIdentifier(interfaceIdentifier, "list_load_balancer_ssl_ciphers_and_protocols"),
		"list_load_balancer_virtual_server_statuses":                       core.NewMethodIdentifier(interfaceIdentifier, "list_load_balancer_virtual_server_statuses"),
		"list_load_balancer_virtual_servers":                               core.NewMethodIdentifier(interfaceIdentifier, "list_load_balancer_virtual_servers"),
		"list_load_balancer_virtual_servers_statistics":                    core.NewMethodIdentifier(interfaceIdentifier, "list_load_balancer_virtual_servers_statistics"),
		"perform_pool_member_action":                                       core.NewMethodIdentifier(interfaceIdentifier, "perform_pool_member_action"),
		"read_load_balancer_application_profile":                           core.NewMethodIdentifier(interfaceIdentifier, "read_load_balancer_application_profile"),
		"read_load_balancer_client_ssl_profile":                            core.NewMethodIdentifier(interfaceIdentifier, "read_load_balancer_client_ssl_profile"),
		"read_load_balancer_monitor":                                       core.NewMethodIdentifier(interfaceIdentifier, "read_load_balancer_monitor"),
		"read_load_balancer_node_usage":                                    core.NewMethodIdentifier(interfaceIdentifier, "read_load_balancer_node_usage"),
		"read_load_balancer_node_usage_summary":                            core.NewMethodIdentifier(interfaceIdentifier, "read_load_balancer_node_usage_summary"),
		"read_load_balancer_persistence_profile":                           core.NewMethodIdentifier(interfaceIdentifier, "read_load_balancer_persistence_profile"),
		"read_load_balancer_pool":                                          core.NewMethodIdentifier(interfaceIdentifier, "read_load_balancer_pool"),
		"read_load_balancer_rule":                                          core.NewMethodIdentifier(interfaceIdentifier, "read_load_balancer_rule"),
		"read_load_balancer_server_ssl_profile":                            core.NewMethodIdentifier(interfaceIdentifier, "read_load_balancer_server_ssl_profile"),
		"read_load_balancer_service":                                       core.NewMethodIdentifier(interfaceIdentifier, "read_load_balancer_service"),
		"read_load_balancer_service_debug_info":                            core.NewMethodIdentifier(interfaceIdentifier, "read_load_balancer_service_debug_info"),
		"read_load_balancer_service_usage":                                 core.NewMethodIdentifier(interfaceIdentifier, "read_load_balancer_service_usage"),
		"read_load_balancer_virtual_server":                                core.NewMethodIdentifier(interfaceIdentifier, "read_load_balancer_virtual_server"),
		"update_load_balancer_application_profile":                         core.NewMethodIdentifier(interfaceIdentifier, "update_load_balancer_application_profile"),
		"update_load_balancer_client_ssl_profile":                          core.NewMethodIdentifier(interfaceIdentifier, "update_load_balancer_client_ssl_profile"),
		"update_load_balancer_monitor":                                     core.NewMethodIdentifier(interfaceIdentifier, "update_load_balancer_monitor"),
		"update_load_balancer_persistence_profile":                         core.NewMethodIdentifier(interfaceIdentifier, "update_load_balancer_persistence_profile"),
		"update_load_balancer_pool":                                        core.NewMethodIdentifier(interfaceIdentifier, "update_load_balancer_pool"),
		"update_load_balancer_rule":                                        core.NewMethodIdentifier(interfaceIdentifier, "update_load_balancer_rule"),
		"update_load_balancer_server_ssl_profile":                          core.NewMethodIdentifier(interfaceIdentifier, "update_load_balancer_server_ssl_profile"),
		"update_load_balancer_service":                                     core.NewMethodIdentifier(interfaceIdentifier, "update_load_balancer_service"),
		"update_load_balancer_virtual_server":                              core.NewMethodIdentifier(interfaceIdentifier, "update_load_balancer_virtual_server"),
		"update_load_balancer_virtual_server_with_rules_update_with_rules": core.NewMethodIdentifier(interfaceIdentifier, "update_load_balancer_virtual_server_with_rules_update_with_rules"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	mIface := managementPlaneAPINetworkingServicesLoadbalancerClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &mIface
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := mIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) CreateLoadBalancerApplicationProfile(lbAppProfileParam *data.StructValue) (*data.StructValue, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerCreateLoadBalancerApplicationProfileInputType(), typeConverter)
	sv.AddStructField("LbAppProfile", lbAppProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerCreateLoadBalancerApplicationProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "create_load_balancer_application_profile", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerCreateLoadBalancerApplicationProfileOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) CreateLoadBalancerClientSslProfile(lbClientSslProfileParam model.LbClientSslProfile) (model.LbClientSslProfile, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerCreateLoadBalancerClientSslProfileInputType(), typeConverter)
	sv.AddStructField("LbClientSslProfile", lbClientSslProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbClientSslProfile
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerCreateLoadBalancerClientSslProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "create_load_balancer_client_ssl_profile", inputDataValue, executionContext)
	var emptyOutput model.LbClientSslProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerCreateLoadBalancerClientSslProfileOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbClientSslProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) CreateLoadBalancerMonitor(lbMonitorParam *data.StructValue) (*data.StructValue, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerCreateLoadBalancerMonitorInputType(), typeConverter)
	sv.AddStructField("LbMonitor", lbMonitorParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerCreateLoadBalancerMonitorRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "create_load_balancer_monitor", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerCreateLoadBalancerMonitorOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) CreateLoadBalancerPersistenceProfile(lbPersistenceProfileParam *data.StructValue) (*data.StructValue, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerCreateLoadBalancerPersistenceProfileInputType(), typeConverter)
	sv.AddStructField("LbPersistenceProfile", lbPersistenceProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerCreateLoadBalancerPersistenceProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "create_load_balancer_persistence_profile", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerCreateLoadBalancerPersistenceProfileOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) CreateLoadBalancerPool(lbPoolParam model.LbPool) (model.LbPool, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerCreateLoadBalancerPoolInputType(), typeConverter)
	sv.AddStructField("LbPool", lbPoolParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbPool
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerCreateLoadBalancerPoolRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "create_load_balancer_pool", inputDataValue, executionContext)
	var emptyOutput model.LbPool
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerCreateLoadBalancerPoolOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbPool), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) CreateLoadBalancerRule(lbRuleParam model.LbRule) (model.LbRule, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerCreateLoadBalancerRuleInputType(), typeConverter)
	sv.AddStructField("LbRule", lbRuleParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbRule
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerCreateLoadBalancerRuleRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "create_load_balancer_rule", inputDataValue, executionContext)
	var emptyOutput model.LbRule
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerCreateLoadBalancerRuleOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbRule), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) CreateLoadBalancerServerSslProfile(lbServerSslProfileParam model.LbServerSslProfile) (model.LbServerSslProfile, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerCreateLoadBalancerServerSslProfileInputType(), typeConverter)
	sv.AddStructField("LbServerSslProfile", lbServerSslProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbServerSslProfile
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerCreateLoadBalancerServerSslProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "create_load_balancer_server_ssl_profile", inputDataValue, executionContext)
	var emptyOutput model.LbServerSslProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerCreateLoadBalancerServerSslProfileOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbServerSslProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) CreateLoadBalancerService(lbServiceParam model.LbService) (model.LbService, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerCreateLoadBalancerServiceInputType(), typeConverter)
	sv.AddStructField("LbService", lbServiceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbService
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerCreateLoadBalancerServiceRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "create_load_balancer_service", inputDataValue, executionContext)
	var emptyOutput model.LbService
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerCreateLoadBalancerServiceOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbService), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) CreateLoadBalancerVirtualServer(lbVirtualServerParam model.LbVirtualServer) (model.LbVirtualServer, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerCreateLoadBalancerVirtualServerInputType(), typeConverter)
	sv.AddStructField("LbVirtualServer", lbVirtualServerParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbVirtualServer
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerCreateLoadBalancerVirtualServerRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "create_load_balancer_virtual_server", inputDataValue, executionContext)
	var emptyOutput model.LbVirtualServer
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerCreateLoadBalancerVirtualServerOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbVirtualServer), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) CreateLoadBalancerVirtualServerWithRulesCreateWithRules(lbVirtualServerWithRuleParam model.LbVirtualServerWithRule) (model.LbVirtualServerWithRule, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerCreateLoadBalancerVirtualServerWithRulesCreateWithRulesInputType(), typeConverter)
	sv.AddStructField("LbVirtualServerWithRule", lbVirtualServerWithRuleParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbVirtualServerWithRule
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerCreateLoadBalancerVirtualServerWithRulesCreateWithRulesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "create_load_balancer_virtual_server_with_rules_create_with_rules", inputDataValue, executionContext)
	var emptyOutput model.LbVirtualServerWithRule
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerCreateLoadBalancerVirtualServerWithRulesCreateWithRulesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbVirtualServerWithRule), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) DeleteLoadBalancerApplicationProfile(applicationProfileIdParam string) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerDeleteLoadBalancerApplicationProfileInputType(), typeConverter)
	sv.AddStructField("ApplicationProfileId", applicationProfileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerDeleteLoadBalancerApplicationProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "delete_load_balancer_application_profile", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) DeleteLoadBalancerClientSslProfile(clientSslProfileIdParam string) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerDeleteLoadBalancerClientSslProfileInputType(), typeConverter)
	sv.AddStructField("ClientSslProfileId", clientSslProfileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerDeleteLoadBalancerClientSslProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "delete_load_balancer_client_ssl_profile", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) DeleteLoadBalancerMonitor(monitorIdParam string) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerDeleteLoadBalancerMonitorInputType(), typeConverter)
	sv.AddStructField("MonitorId", monitorIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerDeleteLoadBalancerMonitorRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "delete_load_balancer_monitor", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) DeleteLoadBalancerPersistenceProfile(persistenceProfileIdParam string) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerDeleteLoadBalancerPersistenceProfileInputType(), typeConverter)
	sv.AddStructField("PersistenceProfileId", persistenceProfileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerDeleteLoadBalancerPersistenceProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "delete_load_balancer_persistence_profile", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) DeleteLoadBalancerPool(poolIdParam string) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerDeleteLoadBalancerPoolInputType(), typeConverter)
	sv.AddStructField("PoolId", poolIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerDeleteLoadBalancerPoolRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "delete_load_balancer_pool", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) DeleteLoadBalancerRule(ruleIdParam string) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerDeleteLoadBalancerRuleInputType(), typeConverter)
	sv.AddStructField("RuleId", ruleIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerDeleteLoadBalancerRuleRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "delete_load_balancer_rule", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) DeleteLoadBalancerServerSslProfile(serverSslProfileIdParam string) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerDeleteLoadBalancerServerSslProfileInputType(), typeConverter)
	sv.AddStructField("ServerSslProfileId", serverSslProfileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerDeleteLoadBalancerServerSslProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "delete_load_balancer_server_ssl_profile", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) DeleteLoadBalancerService(serviceIdParam string) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerDeleteLoadBalancerServiceInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerDeleteLoadBalancerServiceRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "delete_load_balancer_service", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) DeleteLoadBalancerVirtualServer(virtualServerIdParam string, deleteAssociatedRulesParam *bool) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerDeleteLoadBalancerVirtualServerInputType(), typeConverter)
	sv.AddStructField("VirtualServerId", virtualServerIdParam)
	sv.AddStructField("DeleteAssociatedRules", deleteAssociatedRulesParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerDeleteLoadBalancerVirtualServerRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "delete_load_balancer_virtual_server", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) GetLoadBalancerPoolStatistics(serviceIdParam string, poolIdParam string, sourceParam *string) (model.LbPoolStatistics, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerGetLoadBalancerPoolStatisticsInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("PoolId", poolIdParam)
	sv.AddStructField("Source", sourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbPoolStatistics
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerGetLoadBalancerPoolStatisticsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "get_load_balancer_pool_statistics", inputDataValue, executionContext)
	var emptyOutput model.LbPoolStatistics
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerGetLoadBalancerPoolStatisticsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbPoolStatistics), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) GetLoadBalancerPoolStatus(serviceIdParam string, poolIdParam string, sourceParam *string) (model.LbPoolStatus, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerGetLoadBalancerPoolStatusInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("PoolId", poolIdParam)
	sv.AddStructField("Source", sourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbPoolStatus
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerGetLoadBalancerPoolStatusRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "get_load_balancer_pool_status", inputDataValue, executionContext)
	var emptyOutput model.LbPoolStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerGetLoadBalancerPoolStatusOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbPoolStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) GetLoadBalancerServiceStatistics(serviceIdParam string, sourceParam *string) (model.LbServiceStatistics, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerGetLoadBalancerServiceStatisticsInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("Source", sourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbServiceStatistics
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerGetLoadBalancerServiceStatisticsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "get_load_balancer_service_statistics", inputDataValue, executionContext)
	var emptyOutput model.LbServiceStatistics
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerGetLoadBalancerServiceStatisticsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbServiceStatistics), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) GetLoadBalancerServiceStatus(serviceIdParam string, includeInstanceDetailsParam *bool, sourceParam *string, transportNodeIdsParam *string) (model.LbServiceStatus, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerGetLoadBalancerServiceStatusInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("IncludeInstanceDetails", includeInstanceDetailsParam)
	sv.AddStructField("Source", sourceParam)
	sv.AddStructField("TransportNodeIds", transportNodeIdsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbServiceStatus
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerGetLoadBalancerServiceStatusRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "get_load_balancer_service_status", inputDataValue, executionContext)
	var emptyOutput model.LbServiceStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerGetLoadBalancerServiceStatusOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbServiceStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) GetLoadBalancerVirtualServerStatistics(serviceIdParam string, virtualServerIdParam string, sourceParam *string) (model.LbVirtualServerStatistics, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerGetLoadBalancerVirtualServerStatisticsInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("VirtualServerId", virtualServerIdParam)
	sv.AddStructField("Source", sourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbVirtualServerStatistics
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerGetLoadBalancerVirtualServerStatisticsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "get_load_balancer_virtual_server_statistics", inputDataValue, executionContext)
	var emptyOutput model.LbVirtualServerStatistics
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerGetLoadBalancerVirtualServerStatisticsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbVirtualServerStatistics), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) GetLoadBalancerVirtualServerStatus(serviceIdParam string, virtualServerIdParam string, sourceParam *string) (model.LbVirtualServerStatus, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerGetLoadBalancerVirtualServerStatusInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("VirtualServerId", virtualServerIdParam)
	sv.AddStructField("Source", sourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbVirtualServerStatus
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerGetLoadBalancerVirtualServerStatusRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "get_load_balancer_virtual_server_status", inputDataValue, executionContext)
	var emptyOutput model.LbVirtualServerStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerGetLoadBalancerVirtualServerStatusOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbVirtualServerStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) ListLoadBalancerApplicationProfiles(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, type_Param *string) (model.LbAppProfileListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerListLoadBalancerApplicationProfilesInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Type_", type_Param)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbAppProfileListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerListLoadBalancerApplicationProfilesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "list_load_balancer_application_profiles", inputDataValue, executionContext)
	var emptyOutput model.LbAppProfileListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerListLoadBalancerApplicationProfilesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbAppProfileListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) ListLoadBalancerClientSslProfiles(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.LbClientSslProfileListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerListLoadBalancerClientSslProfilesInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbClientSslProfileListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerListLoadBalancerClientSslProfilesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "list_load_balancer_client_ssl_profiles", inputDataValue, executionContext)
	var emptyOutput model.LbClientSslProfileListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerListLoadBalancerClientSslProfilesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbClientSslProfileListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) ListLoadBalancerMonitors(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, type_Param *string) (model.LbMonitorListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerListLoadBalancerMonitorsInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Type_", type_Param)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbMonitorListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerListLoadBalancerMonitorsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "list_load_balancer_monitors", inputDataValue, executionContext)
	var emptyOutput model.LbMonitorListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerListLoadBalancerMonitorsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbMonitorListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) ListLoadBalancerPersistenceProfiles(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, type_Param *string) (model.LbPersistenceProfileListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerListLoadBalancerPersistenceProfilesInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Type_", type_Param)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbPersistenceProfileListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerListLoadBalancerPersistenceProfilesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "list_load_balancer_persistence_profiles", inputDataValue, executionContext)
	var emptyOutput model.LbPersistenceProfileListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerListLoadBalancerPersistenceProfilesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbPersistenceProfileListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) ListLoadBalancerPoolStatistics(serviceIdParam string, sourceParam *string) (model.LbPoolStatisticsListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerListLoadBalancerPoolStatisticsInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("Source", sourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbPoolStatisticsListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerListLoadBalancerPoolStatisticsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "list_load_balancer_pool_statistics", inputDataValue, executionContext)
	var emptyOutput model.LbPoolStatisticsListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerListLoadBalancerPoolStatisticsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbPoolStatisticsListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) ListLoadBalancerPoolStatuses(serviceIdParam string, sourceParam *string) (model.LbPoolStatusListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerListLoadBalancerPoolStatusesInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("Source", sourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbPoolStatusListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerListLoadBalancerPoolStatusesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "list_load_balancer_pool_statuses", inputDataValue, executionContext)
	var emptyOutput model.LbPoolStatusListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerListLoadBalancerPoolStatusesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbPoolStatusListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) ListLoadBalancerPools(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.LbPoolListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerListLoadBalancerPoolsInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbPoolListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerListLoadBalancerPoolsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "list_load_balancer_pools", inputDataValue, executionContext)
	var emptyOutput model.LbPoolListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerListLoadBalancerPoolsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbPoolListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) ListLoadBalancerRules(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.LbRuleListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerListLoadBalancerRulesInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbRuleListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerListLoadBalancerRulesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "list_load_balancer_rules", inputDataValue, executionContext)
	var emptyOutput model.LbRuleListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerListLoadBalancerRulesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbRuleListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) ListLoadBalancerServerSslProfiles(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.LbServerSslProfileListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerListLoadBalancerServerSslProfilesInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbServerSslProfileListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerListLoadBalancerServerSslProfilesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "list_load_balancer_server_ssl_profiles", inputDataValue, executionContext)
	var emptyOutput model.LbServerSslProfileListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerListLoadBalancerServerSslProfilesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbServerSslProfileListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) ListLoadBalancerServices(cursorParam *string, includedFieldsParam *string, logicalRouterIdParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.LbServiceListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerListLoadBalancerServicesInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("LogicalRouterId", logicalRouterIdParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbServiceListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerListLoadBalancerServicesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "list_load_balancer_services", inputDataValue, executionContext)
	var emptyOutput model.LbServiceListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerListLoadBalancerServicesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbServiceListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) ListLoadBalancerSslCiphersAndProtocols(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.LbSslCipherAndProtocolListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerListLoadBalancerSslCiphersAndProtocolsInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbSslCipherAndProtocolListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerListLoadBalancerSslCiphersAndProtocolsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "list_load_balancer_ssl_ciphers_and_protocols", inputDataValue, executionContext)
	var emptyOutput model.LbSslCipherAndProtocolListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerListLoadBalancerSslCiphersAndProtocolsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbSslCipherAndProtocolListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) ListLoadBalancerVirtualServerStatuses(serviceIdParam string, sourceParam *string) (model.LbVirtualServerStatusListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerListLoadBalancerVirtualServerStatusesInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("Source", sourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbVirtualServerStatusListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerListLoadBalancerVirtualServerStatusesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "list_load_balancer_virtual_server_statuses", inputDataValue, executionContext)
	var emptyOutput model.LbVirtualServerStatusListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerListLoadBalancerVirtualServerStatusesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbVirtualServerStatusListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) ListLoadBalancerVirtualServers(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.LbVirtualServerListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerListLoadBalancerVirtualServersInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbVirtualServerListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerListLoadBalancerVirtualServersRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "list_load_balancer_virtual_servers", inputDataValue, executionContext)
	var emptyOutput model.LbVirtualServerListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerListLoadBalancerVirtualServersOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbVirtualServerListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) ListLoadBalancerVirtualServersStatistics(serviceIdParam string, sourceParam *string) (model.LbVirtualServerStatisticsListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerListLoadBalancerVirtualServersStatisticsInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("Source", sourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbVirtualServerStatisticsListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerListLoadBalancerVirtualServersStatisticsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "list_load_balancer_virtual_servers_statistics", inputDataValue, executionContext)
	var emptyOutput model.LbVirtualServerStatisticsListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerListLoadBalancerVirtualServersStatisticsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbVirtualServerStatisticsListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) PerformPoolMemberAction(poolIdParam string, poolMemberSettingListParam model.PoolMemberSettingList, actionParam string) (model.LbPool, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerPerformPoolMemberActionInputType(), typeConverter)
	sv.AddStructField("PoolId", poolIdParam)
	sv.AddStructField("PoolMemberSettingList", poolMemberSettingListParam)
	sv.AddStructField("Action", actionParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbPool
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerPerformPoolMemberActionRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "perform_pool_member_action", inputDataValue, executionContext)
	var emptyOutput model.LbPool
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerPerformPoolMemberActionOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbPool), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) ReadLoadBalancerApplicationProfile(applicationProfileIdParam string) (*data.StructValue, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerReadLoadBalancerApplicationProfileInputType(), typeConverter)
	sv.AddStructField("ApplicationProfileId", applicationProfileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerReadLoadBalancerApplicationProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "read_load_balancer_application_profile", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerReadLoadBalancerApplicationProfileOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) ReadLoadBalancerClientSslProfile(clientSslProfileIdParam string) (model.LbClientSslProfile, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerReadLoadBalancerClientSslProfileInputType(), typeConverter)
	sv.AddStructField("ClientSslProfileId", clientSslProfileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbClientSslProfile
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerReadLoadBalancerClientSslProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "read_load_balancer_client_ssl_profile", inputDataValue, executionContext)
	var emptyOutput model.LbClientSslProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerReadLoadBalancerClientSslProfileOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbClientSslProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) ReadLoadBalancerMonitor(monitorIdParam string) (*data.StructValue, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerReadLoadBalancerMonitorInputType(), typeConverter)
	sv.AddStructField("MonitorId", monitorIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerReadLoadBalancerMonitorRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "read_load_balancer_monitor", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerReadLoadBalancerMonitorOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) ReadLoadBalancerNodeUsage(nodeIdParam string) (*data.StructValue, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerReadLoadBalancerNodeUsageInputType(), typeConverter)
	sv.AddStructField("NodeId", nodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerReadLoadBalancerNodeUsageRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "read_load_balancer_node_usage", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerReadLoadBalancerNodeUsageOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) ReadLoadBalancerNodeUsageSummary(includeUsagesParam *bool) (model.LbNodeUsageSummary, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerReadLoadBalancerNodeUsageSummaryInputType(), typeConverter)
	sv.AddStructField("IncludeUsages", includeUsagesParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbNodeUsageSummary
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerReadLoadBalancerNodeUsageSummaryRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "read_load_balancer_node_usage_summary", inputDataValue, executionContext)
	var emptyOutput model.LbNodeUsageSummary
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerReadLoadBalancerNodeUsageSummaryOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbNodeUsageSummary), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) ReadLoadBalancerPersistenceProfile(persistenceProfileIdParam string) (*data.StructValue, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerReadLoadBalancerPersistenceProfileInputType(), typeConverter)
	sv.AddStructField("PersistenceProfileId", persistenceProfileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerReadLoadBalancerPersistenceProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "read_load_balancer_persistence_profile", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerReadLoadBalancerPersistenceProfileOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) ReadLoadBalancerPool(poolIdParam string) (model.LbPool, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerReadLoadBalancerPoolInputType(), typeConverter)
	sv.AddStructField("PoolId", poolIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbPool
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerReadLoadBalancerPoolRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "read_load_balancer_pool", inputDataValue, executionContext)
	var emptyOutput model.LbPool
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerReadLoadBalancerPoolOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbPool), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) ReadLoadBalancerRule(ruleIdParam string) (model.LbRule, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerReadLoadBalancerRuleInputType(), typeConverter)
	sv.AddStructField("RuleId", ruleIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbRule
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerReadLoadBalancerRuleRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "read_load_balancer_rule", inputDataValue, executionContext)
	var emptyOutput model.LbRule
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerReadLoadBalancerRuleOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbRule), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) ReadLoadBalancerServerSslProfile(serverSslProfileIdParam string) (model.LbServerSslProfile, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerReadLoadBalancerServerSslProfileInputType(), typeConverter)
	sv.AddStructField("ServerSslProfileId", serverSslProfileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbServerSslProfile
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerReadLoadBalancerServerSslProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "read_load_balancer_server_ssl_profile", inputDataValue, executionContext)
	var emptyOutput model.LbServerSslProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerReadLoadBalancerServerSslProfileOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbServerSslProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) ReadLoadBalancerService(serviceIdParam string) (model.LbService, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerReadLoadBalancerServiceInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbService
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerReadLoadBalancerServiceRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "read_load_balancer_service", inputDataValue, executionContext)
	var emptyOutput model.LbService
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerReadLoadBalancerServiceOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbService), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) ReadLoadBalancerServiceDebugInfo(serviceIdParam string) (model.LbServiceDebugInfo, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerReadLoadBalancerServiceDebugInfoInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbServiceDebugInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerReadLoadBalancerServiceDebugInfoRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "read_load_balancer_service_debug_info", inputDataValue, executionContext)
	var emptyOutput model.LbServiceDebugInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerReadLoadBalancerServiceDebugInfoOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbServiceDebugInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) ReadLoadBalancerServiceUsage(serviceIdParam string) (model.LbServiceUsage, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerReadLoadBalancerServiceUsageInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbServiceUsage
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerReadLoadBalancerServiceUsageRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "read_load_balancer_service_usage", inputDataValue, executionContext)
	var emptyOutput model.LbServiceUsage
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerReadLoadBalancerServiceUsageOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbServiceUsage), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) ReadLoadBalancerVirtualServer(virtualServerIdParam string) (model.LbVirtualServer, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerReadLoadBalancerVirtualServerInputType(), typeConverter)
	sv.AddStructField("VirtualServerId", virtualServerIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbVirtualServer
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerReadLoadBalancerVirtualServerRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "read_load_balancer_virtual_server", inputDataValue, executionContext)
	var emptyOutput model.LbVirtualServer
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerReadLoadBalancerVirtualServerOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbVirtualServer), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) UpdateLoadBalancerApplicationProfile(applicationProfileIdParam string, lbAppProfileParam *data.StructValue) (*data.StructValue, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerUpdateLoadBalancerApplicationProfileInputType(), typeConverter)
	sv.AddStructField("ApplicationProfileId", applicationProfileIdParam)
	sv.AddStructField("LbAppProfile", lbAppProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerUpdateLoadBalancerApplicationProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "update_load_balancer_application_profile", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerUpdateLoadBalancerApplicationProfileOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) UpdateLoadBalancerClientSslProfile(clientSslProfileIdParam string, lbClientSslProfileParam model.LbClientSslProfile) (model.LbClientSslProfile, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerUpdateLoadBalancerClientSslProfileInputType(), typeConverter)
	sv.AddStructField("ClientSslProfileId", clientSslProfileIdParam)
	sv.AddStructField("LbClientSslProfile", lbClientSslProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbClientSslProfile
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerUpdateLoadBalancerClientSslProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "update_load_balancer_client_ssl_profile", inputDataValue, executionContext)
	var emptyOutput model.LbClientSslProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerUpdateLoadBalancerClientSslProfileOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbClientSslProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) UpdateLoadBalancerMonitor(monitorIdParam string, lbMonitorParam *data.StructValue) (*data.StructValue, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerUpdateLoadBalancerMonitorInputType(), typeConverter)
	sv.AddStructField("MonitorId", monitorIdParam)
	sv.AddStructField("LbMonitor", lbMonitorParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerUpdateLoadBalancerMonitorRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "update_load_balancer_monitor", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerUpdateLoadBalancerMonitorOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) UpdateLoadBalancerPersistenceProfile(persistenceProfileIdParam string, lbPersistenceProfileParam *data.StructValue) (*data.StructValue, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerUpdateLoadBalancerPersistenceProfileInputType(), typeConverter)
	sv.AddStructField("PersistenceProfileId", persistenceProfileIdParam)
	sv.AddStructField("LbPersistenceProfile", lbPersistenceProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerUpdateLoadBalancerPersistenceProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "update_load_balancer_persistence_profile", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerUpdateLoadBalancerPersistenceProfileOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) UpdateLoadBalancerPool(poolIdParam string, lbPoolParam model.LbPool) (model.LbPool, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerUpdateLoadBalancerPoolInputType(), typeConverter)
	sv.AddStructField("PoolId", poolIdParam)
	sv.AddStructField("LbPool", lbPoolParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbPool
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerUpdateLoadBalancerPoolRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "update_load_balancer_pool", inputDataValue, executionContext)
	var emptyOutput model.LbPool
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerUpdateLoadBalancerPoolOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbPool), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) UpdateLoadBalancerRule(ruleIdParam string, lbRuleParam model.LbRule) (model.LbRule, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerUpdateLoadBalancerRuleInputType(), typeConverter)
	sv.AddStructField("RuleId", ruleIdParam)
	sv.AddStructField("LbRule", lbRuleParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbRule
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerUpdateLoadBalancerRuleRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "update_load_balancer_rule", inputDataValue, executionContext)
	var emptyOutput model.LbRule
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerUpdateLoadBalancerRuleOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbRule), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) UpdateLoadBalancerServerSslProfile(serverSslProfileIdParam string, lbServerSslProfileParam model.LbServerSslProfile) (model.LbServerSslProfile, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerUpdateLoadBalancerServerSslProfileInputType(), typeConverter)
	sv.AddStructField("ServerSslProfileId", serverSslProfileIdParam)
	sv.AddStructField("LbServerSslProfile", lbServerSslProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbServerSslProfile
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerUpdateLoadBalancerServerSslProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "update_load_balancer_server_ssl_profile", inputDataValue, executionContext)
	var emptyOutput model.LbServerSslProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerUpdateLoadBalancerServerSslProfileOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbServerSslProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) UpdateLoadBalancerService(serviceIdParam string, lbServiceParam model.LbService) (model.LbService, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerUpdateLoadBalancerServiceInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("LbService", lbServiceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbService
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerUpdateLoadBalancerServiceRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "update_load_balancer_service", inputDataValue, executionContext)
	var emptyOutput model.LbService
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerUpdateLoadBalancerServiceOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbService), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) UpdateLoadBalancerVirtualServer(virtualServerIdParam string, lbVirtualServerParam model.LbVirtualServer) (model.LbVirtualServer, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerUpdateLoadBalancerVirtualServerInputType(), typeConverter)
	sv.AddStructField("VirtualServerId", virtualServerIdParam)
	sv.AddStructField("LbVirtualServer", lbVirtualServerParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbVirtualServer
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerUpdateLoadBalancerVirtualServerRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "update_load_balancer_virtual_server", inputDataValue, executionContext)
	var emptyOutput model.LbVirtualServer
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerUpdateLoadBalancerVirtualServerOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbVirtualServer), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesLoadbalancerClient) UpdateLoadBalancerVirtualServerWithRulesUpdateWithRules(virtualServerIdParam string, lbVirtualServerWithRuleParam model.LbVirtualServerWithRule) (model.LbVirtualServerWithRule, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesLoadbalancerUpdateLoadBalancerVirtualServerWithRulesUpdateWithRulesInputType(), typeConverter)
	sv.AddStructField("VirtualServerId", virtualServerIdParam)
	sv.AddStructField("LbVirtualServerWithRule", lbVirtualServerWithRuleParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbVirtualServerWithRule
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesLoadbalancerUpdateLoadBalancerVirtualServerWithRulesUpdateWithRulesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_loadbalancer", "update_load_balancer_virtual_server_with_rules_update_with_rules", inputDataValue, executionContext)
	var emptyOutput model.LbVirtualServerWithRule
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesLoadbalancerUpdateLoadBalancerVirtualServerWithRulesUpdateWithRulesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbVirtualServerWithRule), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
