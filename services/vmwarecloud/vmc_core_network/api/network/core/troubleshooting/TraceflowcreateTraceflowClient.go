// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: TraceflowcreateTraceflow
// Used by client-side stubs.

package troubleshooting

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	vmwarecloudVmc_core_networkModel "github.com/vmware/vsphere-automation-sdk-go/services/vmwarecloud/vmc_core_network/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type TraceflowcreateTraceflowClient interface {

	// create a sddc traceflow
	//
	// @param orgIdParam organization identifier
	// @param createNsxTraceflowRequestParam sddc traceflow request
	// @return OK
	//
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	CreateSddcTraceflowActivity(orgIdParam string, createNsxTraceflowRequestParam vmwarecloudVmc_core_networkModel.CreateNsxTraceflowRequest) (vmwarecloudVmc_core_networkModel.Activity, error)
}

type traceflowcreateTraceflowClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewTraceflowcreateTraceflowClient(connector vapiProtocolClient_.Connector) *traceflowcreateTraceflowClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vmwarecloud.vmc_core_network.api.network.core.troubleshooting.traceflowcreate_traceflow")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"create_sddc_traceflow_activity": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "create_sddc_traceflow_activity"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	tIface := traceflowcreateTraceflowClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &tIface
}

func (tIface *traceflowcreateTraceflowClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := tIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (tIface *traceflowcreateTraceflowClient) CreateSddcTraceflowActivity(orgIdParam string, createNsxTraceflowRequestParam vmwarecloudVmc_core_networkModel.CreateNsxTraceflowRequest) (vmwarecloudVmc_core_networkModel.Activity, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := traceflowcreateTraceflowCreateSddcTraceflowActivityRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(traceflowcreateTraceflowCreateSddcTraceflowActivityInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("CreateNsxTraceflowRequest", createNsxTraceflowRequestParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmwarecloudVmc_core_networkModel.Activity
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_core_network.api.network.core.troubleshooting.traceflowcreate_traceflow", "create_sddc_traceflow_activity", inputDataValue, executionContext)
	var emptyOutput vmwarecloudVmc_core_networkModel.Activity
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TraceflowcreateTraceflowCreateSddcTraceflowActivityOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmwarecloudVmc_core_networkModel.Activity), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
