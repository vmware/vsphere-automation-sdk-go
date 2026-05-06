// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: FlowMon
// Used by client-side stubs.

package dataplane

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type FlowMonClient interface {

	// Run flow monitor for *timeout* seconds for all or certain CPU core(s) and return top 10 flows.
	//
	// @param transportNodeIdParam (required)
	// @param coreIdParam CPU core on which the flows are to be monitored (optional)
	// @param fieldsParam Comma-separated field names to include in query result (optional)
	// @param timeoutParam Timeout for flow monitor in seconds (optional, default to 30)
	// @return com.vmware.nsx.model.EdgeDataplaneTopKFlows
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(transportNodeIdParam string, coreIdParam *int64, fieldsParam *string, timeoutParam *int64) (nsxModel.EdgeDataplaneTopkFlows, error)

	// Starts NSX Edge dataplane flow monitor on all or certain CPU core(s) with a timeout. Stops flow monitor after timeout and dumps the flow file on local file store on edge. If *top_10* argument is set to *true* top 10 flows are collected, else all flows are collected.
	//
	// @param transportNodeIdParam (required)
	// @param edgeDataplaneFlowMonitorStartSettingParam (required)
	// @param fieldsParam Comma-separated field names to include in query result (optional)
	// @param top10Param Collect top 10 flows when set to true, else collect all flows. (optional, default to false)
	// @return com.vmware.nsx.model.EdgeDataplaneFlowMonitorMessage
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(transportNodeIdParam string, edgeDataplaneFlowMonitorStartSettingParam nsxModel.EdgeDataplaneFlowMonitorStartSetting, fieldsParam *string, top10Param *bool) (nsxModel.EdgeDataplaneFlowMonitorMessage, error)
}

type flowMonClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewFlowMonClient(connector vapiProtocolClient_.Connector) *flowMonClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.transport_nodes.node.services.dataplane.flow_mon")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	fIface := flowMonClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &fIface
}

func (fIface *flowMonClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := fIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (fIface *flowMonClient) Get(transportNodeIdParam string, coreIdParam *int64, fieldsParam *string, timeoutParam *int64) (nsxModel.EdgeDataplaneTopkFlows, error) {
	typeConverter := fIface.connector.TypeConverter()
	executionContext := fIface.connector.NewExecutionContext()
	operationRestMetaData := flowMonGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(flowMonGetInputType(), typeConverter)
	sv.AddStructField("TransportNodeId", transportNodeIdParam)
	sv.AddStructField("CoreId", coreIdParam)
	sv.AddStructField("Fields", fieldsParam)
	sv.AddStructField("Timeout", timeoutParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.EdgeDataplaneTopkFlows
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := fIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_nodes.node.services.dataplane.flow_mon", "get", inputDataValue, executionContext)
	var emptyOutput nsxModel.EdgeDataplaneTopkFlows
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), FlowMonGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.EdgeDataplaneTopkFlows), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), fIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (fIface *flowMonClient) Update(transportNodeIdParam string, edgeDataplaneFlowMonitorStartSettingParam nsxModel.EdgeDataplaneFlowMonitorStartSetting, fieldsParam *string, top10Param *bool) (nsxModel.EdgeDataplaneFlowMonitorMessage, error) {
	typeConverter := fIface.connector.TypeConverter()
	executionContext := fIface.connector.NewExecutionContext()
	operationRestMetaData := flowMonUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(flowMonUpdateInputType(), typeConverter)
	sv.AddStructField("TransportNodeId", transportNodeIdParam)
	sv.AddStructField("EdgeDataplaneFlowMonitorStartSetting", edgeDataplaneFlowMonitorStartSettingParam)
	sv.AddStructField("Fields", fieldsParam)
	sv.AddStructField("Top10", top10Param)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.EdgeDataplaneFlowMonitorMessage
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := fIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_nodes.node.services.dataplane.flow_mon", "update", inputDataValue, executionContext)
	var emptyOutput nsxModel.EdgeDataplaneFlowMonitorMessage
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), FlowMonUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.EdgeDataplaneFlowMonitorMessage), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), fIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
