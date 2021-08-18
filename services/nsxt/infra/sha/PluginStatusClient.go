// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: PluginStatus
// Used by client-side stubs.

package sha

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type PluginStatusClient interface {

	// Show all the installed system health plugins on given node
	//
	// @param nodeIdParam The TN node id. (required)
	// @return com.vmware.nsx_policy.model.PerNodeShaPluginStatusListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(nodeIdParam string) (model.PerNodeShaPluginStatusListResult, error)
}

type pluginStatusClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewPluginStatusClient(connector client.Connector) *pluginStatusClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.sha.plugin_status")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"list": core.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	pIface := pluginStatusClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &pIface
}

func (pIface *pluginStatusClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := pIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (pIface *pluginStatusClient) List(nodeIdParam string) (model.PerNodeShaPluginStatusListResult, error) {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(pluginStatusListInputType(), typeConverter)
	sv.AddStructField("NodeId", nodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PerNodeShaPluginStatusListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := pluginStatusListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sha.plugin_status", "list", inputDataValue, executionContext)
	var emptyOutput model.PerNodeShaPluginStatusListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), pluginStatusListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PerNodeShaPluginStatusListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
