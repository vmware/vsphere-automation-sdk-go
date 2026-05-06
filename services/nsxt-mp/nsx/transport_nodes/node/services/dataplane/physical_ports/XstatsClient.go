// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Xstats
// Used by client-side stubs.

package physical_ports

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type XstatsClient interface {

	// Returns the extended statistics for the specified physical port. Statistics information includes Tx and Rx packets, bytes, drops and errors. Output will differ depending upon the physical port hardware type.
	//
	// @param physicalPortParam Name of the physical-port to get the xstats for (required)
	// @param transportNodeIdParam (required)
	// @return com.vmware.nsx.model.PhysicalPortStats
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(physicalPortParam string, transportNodeIdParam string) (nsxModel.PhysicalPortStats, error)
}

type xstatsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewXstatsClient(connector vapiProtocolClient_.Connector) *xstatsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.transport_nodes.node.services.dataplane.physical_ports.xstats")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	xIface := xstatsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &xIface
}

func (xIface *xstatsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := xIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (xIface *xstatsClient) Get(physicalPortParam string, transportNodeIdParam string) (nsxModel.PhysicalPortStats, error) {
	typeConverter := xIface.connector.TypeConverter()
	executionContext := xIface.connector.NewExecutionContext()
	operationRestMetaData := xstatsGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(xstatsGetInputType(), typeConverter)
	sv.AddStructField("PhysicalPort", physicalPortParam)
	sv.AddStructField("TransportNodeId", transportNodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.PhysicalPortStats
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := xIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_nodes.node.services.dataplane.physical_ports.xstats", "get", inputDataValue, executionContext)
	var emptyOutput nsxModel.PhysicalPortStats
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), XstatsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.PhysicalPortStats), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), xIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
