// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: NatRules
// Used by client-side stubs.

package statistics

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type NatRulesClient interface {

	// Returns the summation of statistics for all rules from all logical routers which are present on given transport node. Only cached statistics are supported. The query parameter \"source=realtime\" is not supported.
	//
	// @param nodeIdParam (required)
	// @param sourceParam Data source type. (optional)
	// @return com.vmware.nsx.model.NatStatisticsPerTransportNode
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(nodeIdParam string, sourceParam *string) (model.NatStatisticsPerTransportNode, error)
}

type natRulesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewNatRulesClient(connector client.Connector) *natRulesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.transport_nodes.statistics.nat_rules")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get": core.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	nIface := natRulesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &nIface
}

func (nIface *natRulesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := nIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (nIface *natRulesClient) Get(nodeIdParam string, sourceParam *string) (model.NatStatisticsPerTransportNode, error) {
	typeConverter := nIface.connector.TypeConverter()
	executionContext := nIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(natRulesGetInputType(), typeConverter)
	sv.AddStructField("NodeId", nodeIdParam)
	sv.AddStructField("Source", sourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NatStatisticsPerTransportNode
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := natRulesGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	nIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := nIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_nodes.statistics.nat_rules", "get", inputDataValue, executionContext)
	var emptyOutput model.NatStatisticsPerTransportNode
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), natRulesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NatStatisticsPerTransportNode), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
