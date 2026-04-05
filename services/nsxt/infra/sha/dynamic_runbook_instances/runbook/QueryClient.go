// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Query
// Used by client-side stubs.

package runbook

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type QueryClient interface {

	// Retrieve the Dynamic Runbook path.
	//
	// @param odsDynamicRunbookQueryParam (required)
	// @return com.vmware.nsx_policy.model.OdsDynamicRunbookQueryResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(odsDynamicRunbookQueryParam nsx_policyModel.OdsDynamicRunbookQuery) (nsx_policyModel.OdsDynamicRunbookQueryResult, error)
}

type queryClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewQueryClient(connector vapiProtocolClient_.Connector) *queryClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.sha.dynamic_runbook_instances.runbook.query")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"create": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "create"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	qIface := queryClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &qIface
}

func (qIface *queryClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := qIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (qIface *queryClient) Create(odsDynamicRunbookQueryParam nsx_policyModel.OdsDynamicRunbookQuery) (nsx_policyModel.OdsDynamicRunbookQueryResult, error) {
	typeConverter := qIface.connector.TypeConverter()
	executionContext := qIface.connector.NewExecutionContext()
	operationRestMetaData := queryCreateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(queryCreateInputType(), typeConverter)
	sv.AddStructField("OdsDynamicRunbookQuery", odsDynamicRunbookQueryParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.OdsDynamicRunbookQueryResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := qIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sha.dynamic_runbook_instances.runbook.query", "create", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.OdsDynamicRunbookQueryResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), QueryCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.OdsDynamicRunbookQueryResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), qIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
