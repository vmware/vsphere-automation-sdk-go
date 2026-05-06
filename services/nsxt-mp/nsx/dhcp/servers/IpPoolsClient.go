// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: IpPools
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

type IpPoolsClient interface {

	// Create an ip pool for a local DHCP server
	//  Please use Policy API: PUT /policy/api/v1/infra/segments/{segment-id} { \"display_name\" : \"segment-1\", \"subnets\": [ { \"gateway_address\": \"40.1.1.1/16\", \"dhcp_ranges\": [ \"40.1.2.2-40.1.2.20\" ], \"dhcp_config\": { \"resource_type\": \"SegmentDhcpV4Config\", \"lease_time\": 84600 }, \"network\": \"40.1.2.0/24\" } ], \"dhcp_config_path\" : \"/infra/dhcp-server-configs/dhcp-server-profile\", \"connectivity_path\": \"/infra/tier-1s/mgw\", \"_revision\" : 0 }
	//
	// Deprecated: This API element is deprecated.
	//
	// @param serverIdParam (required)
	// @param dhcpIpPoolParam (required)
	// @return com.vmware.nsx.model.DhcpIpPool
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(serverIdParam string, dhcpIpPoolParam nsxModel.DhcpIpPool) (nsxModel.DhcpIpPool, error)

	// Delete a specific ip pool of a given logical DHCP server.
	//  Please use Policy API: PUT /policy/api/v1/infra/segments/{segment-id} { \"display_name\" : \"segment-1\", \"subnets\": [ { \"gateway_address\": \"40.1.1.1/16\", \"network\": \"40.1.2.0/24\" } ], \"connectivity_path\": \"/infra/tier-1s/mgw\", \"_revision\" : 0 }
	//
	// Deprecated: This API element is deprecated.
	//
	// @param serverIdParam (required)
	// @param poolIdParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(serverIdParam string, poolIdParam string) error

	// Return a specific ip pool of a given logical DHCP server.
	//  Please use Policy API: GET /policy/api/v1/infra/segments/{segment-id}
	//
	// Deprecated: This API element is deprecated.
	//
	// @param serverIdParam (required)
	// @param poolIdParam (required)
	// @return com.vmware.nsx.model.DhcpIpPool
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(serverIdParam string, poolIdParam string) (nsxModel.DhcpIpPool, error)

	// List the ip pools of a logical DHCP server with pagination support.
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
	// @return com.vmware.nsx.model.DhcpIpPoolListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(serverIdParam string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsxModel.DhcpIpPoolListResult, error)

	// Update a specific ip pool of a given logical DHCP server.
	//  Please use Policy API: PUT /policy/api/v1/infra/segments/{segment-id} { \"display_name\" : \"segment-1\", \"subnets\": [ { \"gateway_address\": \"40.1.1.1/16\", \"dhcp_ranges\": [ \"40.1.2.2-40.1.2.20\" ], \"dhcp_config\": { \"resource_type\": \"SegmentDhcpV4Config\", \"lease_time\": 84600 }, \"network\": \"40.1.2.0/24\" } ], \"dhcp_config_path\" : \"/infra/dhcp-server-configs/dhcp-server-profile\", \"connectivity_path\": \"/infra/tier-1s/mgw\", \"_revision\" : 0 }
	//
	// Deprecated: This API element is deprecated.
	//
	// @param serverIdParam (required)
	// @param poolIdParam (required)
	// @param dhcpIpPoolParam (required)
	// @return com.vmware.nsx.model.DhcpIpPool
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(serverIdParam string, poolIdParam string, dhcpIpPoolParam nsxModel.DhcpIpPool) (nsxModel.DhcpIpPool, error)
}

type ipPoolsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewIpPoolsClient(connector vapiProtocolClient_.Connector) *ipPoolsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.dhcp.servers.ip_pools")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"create": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "create"),
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	iIface := ipPoolsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &iIface
}

func (iIface *ipPoolsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := iIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (iIface *ipPoolsClient) Create(serverIdParam string, dhcpIpPoolParam nsxModel.DhcpIpPool) (nsxModel.DhcpIpPool, error) {
	typeConverter := iIface.connector.TypeConverter()
	executionContext := iIface.connector.NewExecutionContext()
	operationRestMetaData := ipPoolsCreateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(ipPoolsCreateInputType(), typeConverter)
	sv.AddStructField("ServerId", serverIdParam)
	sv.AddStructField("DhcpIpPool", dhcpIpPoolParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.DhcpIpPool
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := iIface.connector.GetApiProvider().Invoke("com.vmware.nsx.dhcp.servers.ip_pools", "create", inputDataValue, executionContext)
	var emptyOutput nsxModel.DhcpIpPool
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), IpPoolsCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.DhcpIpPool), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (iIface *ipPoolsClient) Delete(serverIdParam string, poolIdParam string) error {
	typeConverter := iIface.connector.TypeConverter()
	executionContext := iIface.connector.NewExecutionContext()
	operationRestMetaData := ipPoolsDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(ipPoolsDeleteInputType(), typeConverter)
	sv.AddStructField("ServerId", serverIdParam)
	sv.AddStructField("PoolId", poolIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := iIface.connector.GetApiProvider().Invoke("com.vmware.nsx.dhcp.servers.ip_pools", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (iIface *ipPoolsClient) Get(serverIdParam string, poolIdParam string) (nsxModel.DhcpIpPool, error) {
	typeConverter := iIface.connector.TypeConverter()
	executionContext := iIface.connector.NewExecutionContext()
	operationRestMetaData := ipPoolsGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(ipPoolsGetInputType(), typeConverter)
	sv.AddStructField("ServerId", serverIdParam)
	sv.AddStructField("PoolId", poolIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.DhcpIpPool
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := iIface.connector.GetApiProvider().Invoke("com.vmware.nsx.dhcp.servers.ip_pools", "get", inputDataValue, executionContext)
	var emptyOutput nsxModel.DhcpIpPool
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), IpPoolsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.DhcpIpPool), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (iIface *ipPoolsClient) List(serverIdParam string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsxModel.DhcpIpPoolListResult, error) {
	typeConverter := iIface.connector.TypeConverter()
	executionContext := iIface.connector.NewExecutionContext()
	operationRestMetaData := ipPoolsListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(ipPoolsListInputType(), typeConverter)
	sv.AddStructField("ServerId", serverIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.DhcpIpPoolListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := iIface.connector.GetApiProvider().Invoke("com.vmware.nsx.dhcp.servers.ip_pools", "list", inputDataValue, executionContext)
	var emptyOutput nsxModel.DhcpIpPoolListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), IpPoolsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.DhcpIpPoolListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (iIface *ipPoolsClient) Update(serverIdParam string, poolIdParam string, dhcpIpPoolParam nsxModel.DhcpIpPool) (nsxModel.DhcpIpPool, error) {
	typeConverter := iIface.connector.TypeConverter()
	executionContext := iIface.connector.NewExecutionContext()
	operationRestMetaData := ipPoolsUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(ipPoolsUpdateInputType(), typeConverter)
	sv.AddStructField("ServerId", serverIdParam)
	sv.AddStructField("PoolId", poolIdParam)
	sv.AddStructField("DhcpIpPool", dhcpIpPoolParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.DhcpIpPool
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := iIface.connector.GetApiProvider().Invoke("com.vmware.nsx.dhcp.servers.ip_pools", "update", inputDataValue, executionContext)
	var emptyOutput nsxModel.DhcpIpPool
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), IpPoolsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.DhcpIpPool), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
