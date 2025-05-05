// Copyright (c) 2019-2025 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: DistributedVxlanConnections
// Used by client-side stubs.

package infra

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type DistributedVxlanConnectionsClient interface {

	// Delete a distributed vxlan connection
	//
	// @param distributedVxlanConnectionIdParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(distributedVxlanConnectionIdParam string) error

	// Read a distributed vxlan connection.
	//
	// @param distributedVxlanConnectionIdParam (required)
	// @return com.vmware.nsx_policy.model.DistributedVxlanConnection
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(distributedVxlanConnectionIdParam string) (nsx_policyModel.DistributedVxlanConnection, error)

	// List distributed vxlan connection
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includeMarkForDeleteObjectsParam If true, resources that are marked for deletion will be included in the results. By default, these resources are not included. (optional, default to false)
	// @param includedFieldsParam Note - this parameter currently only works when used with the search APIs /policy/api/v1/search/query and /policy/api/v1/search/dsl. It is ignored for other list APIs. (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam If true, results are sorted in ascending order (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_policy.model.DistributedVxlanConnectionListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.DistributedVxlanConnectionListResult, error)

	// If a distributed vxlan connection is not already present, create a new distributed vxlan connection. If it already exists, update the distributed vxlan connection.
	//
	// @param distributedVxlanConnectionIdParam (required)
	// @param distributedVxlanConnectionParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(distributedVxlanConnectionIdParam string, distributedVxlanConnectionParam nsx_policyModel.DistributedVxlanConnection) error

	// If a distributed vxlan connection is not already present, create a new distributed vxlan connection. If it already exists, update the distributed vxlan connection. This operation will fully replace the object.
	//
	// @param distributedVxlanConnectionIdParam (required)
	// @param distributedVxlanConnectionParam (required)
	// @return com.vmware.nsx_policy.model.DistributedVxlanConnection
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(distributedVxlanConnectionIdParam string, distributedVxlanConnectionParam nsx_policyModel.DistributedVxlanConnection) (nsx_policyModel.DistributedVxlanConnection, error)
}

type distributedVxlanConnectionsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewDistributedVxlanConnectionsClient(connector vapiProtocolClient_.Connector) *distributedVxlanConnectionsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.distributed_vxlan_connections")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"patch":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	dIface := distributedVxlanConnectionsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &dIface
}

func (dIface *distributedVxlanConnectionsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := dIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (dIface *distributedVxlanConnectionsClient) Delete(distributedVxlanConnectionIdParam string) error {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	operationRestMetaData := distributedVxlanConnectionsDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(distributedVxlanConnectionsDeleteInputType(), typeConverter)
	sv.AddStructField("DistributedVxlanConnectionId", distributedVxlanConnectionIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.distributed_vxlan_connections", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (dIface *distributedVxlanConnectionsClient) Get(distributedVxlanConnectionIdParam string) (nsx_policyModel.DistributedVxlanConnection, error) {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	operationRestMetaData := distributedVxlanConnectionsGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(distributedVxlanConnectionsGetInputType(), typeConverter)
	sv.AddStructField("DistributedVxlanConnectionId", distributedVxlanConnectionIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.DistributedVxlanConnection
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.distributed_vxlan_connections", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.DistributedVxlanConnection
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), DistributedVxlanConnectionsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.DistributedVxlanConnection), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (dIface *distributedVxlanConnectionsClient) List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.DistributedVxlanConnectionListResult, error) {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	operationRestMetaData := distributedVxlanConnectionsListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(distributedVxlanConnectionsListInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludeMarkForDeleteObjects", includeMarkForDeleteObjectsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.DistributedVxlanConnectionListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.distributed_vxlan_connections", "list", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.DistributedVxlanConnectionListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), DistributedVxlanConnectionsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.DistributedVxlanConnectionListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (dIface *distributedVxlanConnectionsClient) Patch(distributedVxlanConnectionIdParam string, distributedVxlanConnectionParam nsx_policyModel.DistributedVxlanConnection) error {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	operationRestMetaData := distributedVxlanConnectionsPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(distributedVxlanConnectionsPatchInputType(), typeConverter)
	sv.AddStructField("DistributedVxlanConnectionId", distributedVxlanConnectionIdParam)
	sv.AddStructField("DistributedVxlanConnection", distributedVxlanConnectionParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.distributed_vxlan_connections", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (dIface *distributedVxlanConnectionsClient) Update(distributedVxlanConnectionIdParam string, distributedVxlanConnectionParam nsx_policyModel.DistributedVxlanConnection) (nsx_policyModel.DistributedVxlanConnection, error) {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	operationRestMetaData := distributedVxlanConnectionsUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(distributedVxlanConnectionsUpdateInputType(), typeConverter)
	sv.AddStructField("DistributedVxlanConnectionId", distributedVxlanConnectionIdParam)
	sv.AddStructField("DistributedVxlanConnection", distributedVxlanConnectionParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.DistributedVxlanConnection
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.distributed_vxlan_connections", "update", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.DistributedVxlanConnection
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), DistributedVxlanConnectionsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.DistributedVxlanConnection), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
