// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Operations
// Used by client-side stubs.

package core

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	vmwarecloudVmc_core_inventoryModel "github.com/vmware/vsphere-automation-sdk-go/services/vmwarecloud/vmc_core_inventory/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type OperationsClient interface {

	// Get a collection of operations scoped by the operator
	//
	// @param orgIdParam organization identifier
	// @param includeDeletedResourcesParam should the deleted resources be included in the results?
	// @param pageParam Zero-based page index (0..N)
	// @param sizeParam The size of the page to be returned
	// @param sortParam Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported.
	// @return OK
	//
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	GetCoreOperationsAsOperator(orgIdParam *string, includeDeletedResourcesParam *bool, pageParam *int64, sizeParam *int64, sortParam []string) (vmwarecloudVmc_core_inventoryModel.OperationResponse, error)

	// Get details for a specific core operation
	//
	// @param operationIdParam operation identifier
	// @param orgIdParam organization identifier
	// @return OK
	//
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	// @throws NotFound Not Found
	GetCoreOperationByIdAsOperator(operationIdParam string, orgIdParam *string) (vmwarecloudVmc_core_inventoryModel.Operation, error)
}

type operationsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewOperationsClient(connector vapiProtocolClient_.Connector) *operationsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vmwarecloud.vmc_core_inventory.api.operation.operator.core.operations")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get_core_operations_as_operator":      vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get_core_operations_as_operator"),
		"get_core_operation_by_id_as_operator": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get_core_operation_by_id_as_operator"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	oIface := operationsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &oIface
}

func (oIface *operationsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := oIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (oIface *operationsClient) GetCoreOperationsAsOperator(orgIdParam *string, includeDeletedResourcesParam *bool, pageParam *int64, sizeParam *int64, sortParam []string) (vmwarecloudVmc_core_inventoryModel.OperationResponse, error) {
	typeConverter := oIface.connector.TypeConverter()
	executionContext := oIface.connector.NewExecutionContext()
	operationRestMetaData := operationsGetCoreOperationsAsOperatorRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(operationsGetCoreOperationsAsOperatorInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("IncludeDeletedResources", includeDeletedResourcesParam)
	sv.AddStructField("Page", pageParam)
	sv.AddStructField("Size", sizeParam)
	sv.AddStructField("Sort", sortParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmwarecloudVmc_core_inventoryModel.OperationResponse
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := oIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_core_inventory.api.operation.operator.core.operations", "get_core_operations_as_operator", inputDataValue, executionContext)
	var emptyOutput vmwarecloudVmc_core_inventoryModel.OperationResponse
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), OperationsGetCoreOperationsAsOperatorOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmwarecloudVmc_core_inventoryModel.OperationResponse), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), oIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (oIface *operationsClient) GetCoreOperationByIdAsOperator(operationIdParam string, orgIdParam *string) (vmwarecloudVmc_core_inventoryModel.Operation, error) {
	typeConverter := oIface.connector.TypeConverter()
	executionContext := oIface.connector.NewExecutionContext()
	operationRestMetaData := operationsGetCoreOperationByIdAsOperatorRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(operationsGetCoreOperationByIdAsOperatorInputType(), typeConverter)
	sv.AddStructField("OperationId", operationIdParam)
	sv.AddStructField("OrgId", orgIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmwarecloudVmc_core_inventoryModel.Operation
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := oIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_core_inventory.api.operation.operator.core.operations", "get_core_operation_by_id_as_operator", inputDataValue, executionContext)
	var emptyOutput vmwarecloudVmc_core_inventoryModel.Operation
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), OperationsGetCoreOperationByIdAsOperatorOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmwarecloudVmc_core_inventoryModel.Operation), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), oIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
