// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Servers
// Used by client-side stubs.

package dhcp

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type ServersClient interface {

	// Create a logical DHCP server with v4 and/or v6 servers.
	//  Please use Policy API: PUT /policy/api/v1/infra/tier-0s/{tier-0-id} PUT /policy/api/v1/infra/tier-1s/{tier-1-id} { \"dhcp_config_path\" : \"/infra/dhcp-server-configs/dhcp-server-profile\", \"_revision\" : 0 } PUT /policy/api/v1/infra/segments/{segment-id} { \"display_name\" : \"segment-1\", \"subnets\": [ { \"gateway_address\": \"40.1.1.1/16\", \"dhcp_ranges\": [ \"40.1.2.2-40.1.2.20\" ], \"dhcp_config\": { \"resource_type\": \"SegmentDhcpV4Config\", \"lease_time\": 84600 }, \"network\": \"40.1.2.0/24\" } ], \"dhcp_config_path\" : \"/infra/dhcp-server-configs/dhcp-server-profile\", \"connectivity_path\": \"/infra/tier-1s/mgw\", \"_revision\" : 0 }
	//
	// Deprecated: This API element is deprecated.
	//
	// @param logicalDhcpServerParam (required)
	// @return com.vmware.nsx.model.LogicalDhcpServer
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(logicalDhcpServerParam nsxModel.LogicalDhcpServer) (nsxModel.LogicalDhcpServer, error)

	// Delete a logical DHCP server specified by server id.
	//  Please use Policy API: PUT /policy/api/v1/infra/tier-0s/{tier-0-id} PUT /policy/api/v1/infra/tier-1s/{tier-1-id} { \"display_name\" : \"mgw\", \"_revision\" : 0 } PUT /policy/api/v1/infra/segments/{segment-id} { \"display_name\" : \"segment-1\", \"subnets\": [ { \"gateway_address\": \"40.1.1.1/16\", \"network\": \"40.1.2.0/24\" } ], \"connectivity_path\": \"/infra/tier-1s/mgw\", \"_revision\" : 0 }
	//
	// Deprecated: This API element is deprecated.
	//
	// @param serverIdParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(serverIdParam string) error

	// Retrieve a logical DHCP server specified by server id.
	//  Please use Policy API: GET /policy/api/v1/infra/tier-0s/{tier-0-id} GET /policy/api/v1/infra/tier-1s/{tier-1-id} GET /policy/api/v1/infra/segments/{segment-id}
	//
	// Deprecated: This API element is deprecated.
	//
	// @param serverIdParam (required)
	// @return com.vmware.nsx.model.LogicalDhcpServer
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(serverIdParam string) (nsxModel.LogicalDhcpServer, error)

	// List logical DHCP servers with pagination support.
	//  Please use Policy API: GET /policy/api/v1/infra/tier-0s/{tier-0-id} GET /policy/api/v1/infra/tier-1s/{tier-1-id} GET /policy/api/v1/infra/segments/{segment-id}
	//
	// Deprecated: This API element is deprecated.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Note - this parameter currently only works when used with the search APIs /policy/api/v1/search/query and /policy/api/v1/search/dsl. It is ignored for other list APIs. (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam If true, results are sorted in ascending order (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx.model.LogicalDhcpServerListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsxModel.LogicalDhcpServerListResult, error)

	// Update a logical DHCP server with new configurations.
	//  Please use Policy API: PUT /policy/api/v1/infra/tier-0s/{tier-0-id} PUT /policy/api/v1/infra/tier-1s/{tier-1-id} { \"dhcp_config_path\" : \"/infra/dhcp-server-configs/dhcp-server-profile\", \"_revision\" : 0 } PUT /policy/api/v1/infra/segments/{segment-id} { \"display_name\" : \"segment-1\", \"subnets\": [ { \"gateway_address\": \"40.1.1.1/16\", \"dhcp_ranges\": [ \"40.1.2.2-40.1.2.20\" ], \"dhcp_config\": { \"resource_type\": \"SegmentDhcpV4Config\", \"lease_time\": 84600 }, \"network\": \"40.1.2.0/24\" } ], \"dhcp_config_path\" : \"/infra/dhcp-server-configs/dhcp-server-profile\", \"connectivity_path\": \"/infra/tier-1s/mgw\", \"_revision\" : 0 }
	//
	// Deprecated: This API element is deprecated.
	//
	// @param serverIdParam (required)
	// @param logicalDhcpServerParam (required)
	// @return com.vmware.nsx.model.LogicalDhcpServer
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(serverIdParam string, logicalDhcpServerParam nsxModel.LogicalDhcpServer) (nsxModel.LogicalDhcpServer, error)
}

type serversClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewServersClient(connector vapiProtocolClient_.Connector) *serversClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.dhcp.servers")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"create": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "create"),
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	sIface := serversClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *serversClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *serversClient) Create(logicalDhcpServerParam nsxModel.LogicalDhcpServer) (nsxModel.LogicalDhcpServer, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := serversCreateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(serversCreateInputType(), typeConverter)
	sv.AddStructField("LogicalDhcpServer", logicalDhcpServerParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.LogicalDhcpServer
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.dhcp.servers", "create", inputDataValue, executionContext)
	var emptyOutput nsxModel.LogicalDhcpServer
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ServersCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.LogicalDhcpServer), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *serversClient) Delete(serverIdParam string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := serversDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(serversDeleteInputType(), typeConverter)
	sv.AddStructField("ServerId", serverIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.dhcp.servers", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *serversClient) Get(serverIdParam string) (nsxModel.LogicalDhcpServer, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := serversGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(serversGetInputType(), typeConverter)
	sv.AddStructField("ServerId", serverIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.LogicalDhcpServer
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.dhcp.servers", "get", inputDataValue, executionContext)
	var emptyOutput nsxModel.LogicalDhcpServer
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ServersGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.LogicalDhcpServer), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *serversClient) List(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsxModel.LogicalDhcpServerListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := serversListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(serversListInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.LogicalDhcpServerListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.dhcp.servers", "list", inputDataValue, executionContext)
	var emptyOutput nsxModel.LogicalDhcpServerListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ServersListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.LogicalDhcpServerListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *serversClient) Update(serverIdParam string, logicalDhcpServerParam nsxModel.LogicalDhcpServer) (nsxModel.LogicalDhcpServer, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := serversUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(serversUpdateInputType(), typeConverter)
	sv.AddStructField("ServerId", serverIdParam)
	sv.AddStructField("LogicalDhcpServer", logicalDhcpServerParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.LogicalDhcpServer
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.dhcp.servers", "update", inputDataValue, executionContext)
	var emptyOutput nsxModel.LogicalDhcpServer
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ServersUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.LogicalDhcpServer), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
