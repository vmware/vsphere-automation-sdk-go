// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Ipv6IpPools
// Used by client-side stubs.

package servers

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type Ipv6IpPoolsClient interface {

	// Create an ip pool for a local DHCP IPv6 server
	//  Please use Policy API: PUT /policy/api/v1/infra/segments/{segment-id} { \"display_name\" : \"segment-1\", \"subnets\": [ { \"gateway_address\": \"2019::200.1.0.1/128\", \"dhcp_ranges\": [ \"2019::200.1.0.0-2019::200.255.255.255\" ], \"dhcp_config\": { \"resource_type\": \"SegmentDhcpV6Config\", \"lease_time\": 84600 }, \"network\": \"2019::200.1.1.1/128\" } ], \"dhcp_config_path\" : \"/infra/dhcp-server-configs/dhcp-server-profile\", \"connectivity_path\": \"/infra/tier-1s/mgw\", \"_revision\" : 0 }
	//
	// Deprecated: This API element is deprecated.
	//
	// @param serverIdParam (required)
	// @param dhcpV6IpPoolParam (required)
	// @return com.vmware.nsx.model.DhcpV6IpPool
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(serverIdParam string, dhcpV6IpPoolParam nsxModel.DhcpV6IpPool) (nsxModel.DhcpV6IpPool, error)

	// Return a specific ip pool of a given logical DHCP IPv6 server.
	//  Please use Policy API: GET /policy/api/v1/infra/segments/{segment-id}
	//
	// Deprecated: This API element is deprecated.
	//
	// @param serverIdParam (required)
	// @param poolIdParam (required)
	// @return com.vmware.nsx.model.DhcpV6IpPool
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(serverIdParam string, poolIdParam string) (nsxModel.DhcpV6IpPool, error)

	// List the ip pools of a logical DHCP IPv6 server with pagination support.
	//  Please use Policy API: GET /policy/api/v1/infra/segments/{segment-id}
	//
	// Deprecated: This API element is deprecated.
	//
	// @param serverIdParam (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Note - this parameter currently only works when used with the search APIs /policy/api/v1/search/query and /policy/api/v1/search/dsl. It is ignored for other list APIs. (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam If true, results are sorted in ascending order (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx.model.DhcpV6IpPoolListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(serverIdParam string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsxModel.DhcpV6IpPoolListResult, error)

	// Update a specific ip pool of a given logical DHCP IPv6 server.
	//  Please use Policy API: PUT /policy/api/v1/infra/segments/{segment-id} { \"display_name\" : \"segment-1\", \"subnets\": [ { \"gateway_address\": \"2019::200.1.0.1/128\", \"dhcp_ranges\": [ \"2019::200.1.0.0-2019::200.255.255.255\" ], \"dhcp_config\": { \"resource_type\": \"SegmentDhcpV6Config\", \"lease_time\": 84600 }, \"network\": \"2019::200.1.1.1/128\" } ], \"dhcp_config_path\" : \"/infra/dhcp-server-configs/dhcp-server-profile\", \"connectivity_path\": \"/infra/tier-1s/mgw\", \"_revision\" : 0 }
	//
	// Deprecated: This API element is deprecated.
	//
	// @param serverIdParam (required)
	// @param poolIdParam (required)
	// @param dhcpV6IpPoolParam (required)
	// @return com.vmware.nsx.model.DhcpV6IpPool
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(serverIdParam string, poolIdParam string, dhcpV6IpPoolParam nsxModel.DhcpV6IpPool) (nsxModel.DhcpV6IpPool, error)
}

type ipv6IpPoolsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewIpv6IpPoolsClient(connector vapiProtocolClient_.Connector) *ipv6IpPoolsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.dhcp.servers.ipv6_ip_pools")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"create": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "create"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	iIface := ipv6IpPoolsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &iIface
}

func (iIface *ipv6IpPoolsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := iIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (iIface *ipv6IpPoolsClient) Create(serverIdParam string, dhcpV6IpPoolParam nsxModel.DhcpV6IpPool) (nsxModel.DhcpV6IpPool, error) {
	typeConverter := iIface.connector.TypeConverter()
	executionContext := iIface.connector.NewExecutionContext()
	operationRestMetaData := ipv6IpPoolsCreateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(ipv6IpPoolsCreateInputType(), typeConverter)
	sv.AddStructField("ServerId", serverIdParam)
	sv.AddStructField("DhcpV6IpPool", dhcpV6IpPoolParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.DhcpV6IpPool
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := iIface.connector.GetApiProvider().Invoke("com.vmware.nsx.dhcp.servers.ipv6_ip_pools", "create", inputDataValue, executionContext)
	var emptyOutput nsxModel.DhcpV6IpPool
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), Ipv6IpPoolsCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.DhcpV6IpPool), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (iIface *ipv6IpPoolsClient) Get(serverIdParam string, poolIdParam string) (nsxModel.DhcpV6IpPool, error) {
	typeConverter := iIface.connector.TypeConverter()
	executionContext := iIface.connector.NewExecutionContext()
	operationRestMetaData := ipv6IpPoolsGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(ipv6IpPoolsGetInputType(), typeConverter)
	sv.AddStructField("ServerId", serverIdParam)
	sv.AddStructField("PoolId", poolIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.DhcpV6IpPool
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := iIface.connector.GetApiProvider().Invoke("com.vmware.nsx.dhcp.servers.ipv6_ip_pools", "get", inputDataValue, executionContext)
	var emptyOutput nsxModel.DhcpV6IpPool
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), Ipv6IpPoolsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.DhcpV6IpPool), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (iIface *ipv6IpPoolsClient) List(serverIdParam string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsxModel.DhcpV6IpPoolListResult, error) {
	typeConverter := iIface.connector.TypeConverter()
	executionContext := iIface.connector.NewExecutionContext()
	operationRestMetaData := ipv6IpPoolsListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(ipv6IpPoolsListInputType(), typeConverter)
	sv.AddStructField("ServerId", serverIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.DhcpV6IpPoolListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := iIface.connector.GetApiProvider().Invoke("com.vmware.nsx.dhcp.servers.ipv6_ip_pools", "list", inputDataValue, executionContext)
	var emptyOutput nsxModel.DhcpV6IpPoolListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), Ipv6IpPoolsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.DhcpV6IpPoolListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (iIface *ipv6IpPoolsClient) Update(serverIdParam string, poolIdParam string, dhcpV6IpPoolParam nsxModel.DhcpV6IpPool) (nsxModel.DhcpV6IpPool, error) {
	typeConverter := iIface.connector.TypeConverter()
	executionContext := iIface.connector.NewExecutionContext()
	operationRestMetaData := ipv6IpPoolsUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(ipv6IpPoolsUpdateInputType(), typeConverter)
	sv.AddStructField("ServerId", serverIdParam)
	sv.AddStructField("PoolId", poolIdParam)
	sv.AddStructField("DhcpV6IpPool", dhcpV6IpPoolParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.DhcpV6IpPool
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := iIface.connector.GetApiProvider().Invoke("com.vmware.nsx.dhcp.servers.ipv6_ip_pools", "update", inputDataValue, executionContext)
	var emptyOutput nsxModel.DhcpV6IpPool
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), Ipv6IpPoolsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.DhcpV6IpPool), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
