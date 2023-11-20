// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Traceflow
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

type TraceflowClient interface {

	// gets all sddc traceflow matching the given query parameters
	//
	// @param orgIdParam organization identifier
	// @param networkConnectivityConfigIdParam NetworkConnectivityConfig Id
	// @return OK
	//
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	GetAllSddcTraceflows(orgIdParam string, networkConnectivityConfigIdParam *string) ([]vmwarecloudVmc_core_networkModel.SddcTraceflow, error)

	// create a sddc traceflow
	//
	// @param orgIdParam organization identifier
	// @param createNsxTraceflowRequestParam sddc traceflow request
	// @return OK
	//
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	CreateSddcTraceflowActivityV1(orgIdParam string, createNsxTraceflowRequestParam vmwarecloudVmc_core_networkModel.CreateNsxTraceflowRequest) (vmwarecloudVmc_core_networkModel.Activity, error)

	// gets sddc traceflow of given id
	//
	// @param orgIdParam organization identifier
	// @param traceflowIdParam Sddc Traceflow Id
	// @return OK
	//
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	// @throws NotFound Cannot find the sddc traceflow with given identifier
	GetSddcTraceflowById(orgIdParam string, traceflowIdParam string) (vmwarecloudVmc_core_networkModel.SddcTraceflow, error)

	// retrace a sddc traceflow
	//
	// @param orgIdParam organization identifier
	// @param traceflowIdParam Sddc Traceflow Id
	// @param actionParam Action to be performed
	// @return OK
	//
	// @throws InvalidRequest Bad Request
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	// @throws NotFound Cannot find the sddc traceflow with given identifier
	RetraceSddcTraceflowActivityV1(orgIdParam string, traceflowIdParam string, actionParam *string) (vmwarecloudVmc_core_networkModel.Activity, error)

	// Delete a specific sddc traceflow
	//
	// @param orgIdParam organization identifier
	// @param traceflowIdParam Sddc Traceflow Id
	// @return Accepted
	//
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	// @throws NotFound Cannot find the sddc traceflow with given identifier
	DeleteSddcTraceflowById(orgIdParam string, traceflowIdParam string) error
}

type traceflowClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewTraceflowClient(connector vapiProtocolClient_.Connector) *traceflowClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vmwarecloud.vmc_core_network.api.network.core.troubleshooting.traceflow")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get_all_sddc_traceflows":            vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get_all_sddc_traceflows"),
		"create_sddc_traceflow_activity_v1":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "create_sddc_traceflow_activity_v1"),
		"get_sddc_traceflow_by_id":           vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get_sddc_traceflow_by_id"),
		"retrace_sddc_traceflow_activity_v1": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "retrace_sddc_traceflow_activity_v1"),
		"delete_sddc_traceflow_by_id":        vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete_sddc_traceflow_by_id"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	tIface := traceflowClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &tIface
}

func (tIface *traceflowClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := tIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (tIface *traceflowClient) GetAllSddcTraceflows(orgIdParam string, networkConnectivityConfigIdParam *string) ([]vmwarecloudVmc_core_networkModel.SddcTraceflow, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := traceflowGetAllSddcTraceflowsRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(traceflowGetAllSddcTraceflowsInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("NetworkConnectivityConfigId", networkConnectivityConfigIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []vmwarecloudVmc_core_networkModel.SddcTraceflow
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_core_network.api.network.core.troubleshooting.traceflow", "get_all_sddc_traceflows", inputDataValue, executionContext)
	var emptyOutput []vmwarecloudVmc_core_networkModel.SddcTraceflow
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TraceflowGetAllSddcTraceflowsOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.([]vmwarecloudVmc_core_networkModel.SddcTraceflow), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *traceflowClient) CreateSddcTraceflowActivityV1(orgIdParam string, createNsxTraceflowRequestParam vmwarecloudVmc_core_networkModel.CreateNsxTraceflowRequest) (vmwarecloudVmc_core_networkModel.Activity, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := traceflowCreateSddcTraceflowActivityV1RestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(traceflowCreateSddcTraceflowActivityV1InputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("CreateNsxTraceflowRequest", createNsxTraceflowRequestParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmwarecloudVmc_core_networkModel.Activity
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_core_network.api.network.core.troubleshooting.traceflow", "create_sddc_traceflow_activity_v1", inputDataValue, executionContext)
	var emptyOutput vmwarecloudVmc_core_networkModel.Activity
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TraceflowCreateSddcTraceflowActivityV1OutputType())
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

func (tIface *traceflowClient) GetSddcTraceflowById(orgIdParam string, traceflowIdParam string) (vmwarecloudVmc_core_networkModel.SddcTraceflow, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := traceflowGetSddcTraceflowByIdRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(traceflowGetSddcTraceflowByIdInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("TraceflowId", traceflowIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmwarecloudVmc_core_networkModel.SddcTraceflow
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_core_network.api.network.core.troubleshooting.traceflow", "get_sddc_traceflow_by_id", inputDataValue, executionContext)
	var emptyOutput vmwarecloudVmc_core_networkModel.SddcTraceflow
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TraceflowGetSddcTraceflowByIdOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmwarecloudVmc_core_networkModel.SddcTraceflow), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *traceflowClient) RetraceSddcTraceflowActivityV1(orgIdParam string, traceflowIdParam string, actionParam *string) (vmwarecloudVmc_core_networkModel.Activity, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := traceflowRetraceSddcTraceflowActivityV1RestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(traceflowRetraceSddcTraceflowActivityV1InputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("TraceflowId", traceflowIdParam)
	sv.AddStructField("Action", actionParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmwarecloudVmc_core_networkModel.Activity
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_core_network.api.network.core.troubleshooting.traceflow", "retrace_sddc_traceflow_activity_v1", inputDataValue, executionContext)
	var emptyOutput vmwarecloudVmc_core_networkModel.Activity
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TraceflowRetraceSddcTraceflowActivityV1OutputType())
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

func (tIface *traceflowClient) DeleteSddcTraceflowById(orgIdParam string, traceflowIdParam string) error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := traceflowDeleteSddcTraceflowByIdRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(traceflowDeleteSddcTraceflowByIdInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("TraceflowId", traceflowIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_core_network.api.network.core.troubleshooting.traceflow", "delete_sddc_traceflow_by_id", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
