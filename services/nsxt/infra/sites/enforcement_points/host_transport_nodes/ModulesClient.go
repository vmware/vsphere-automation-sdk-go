// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Modules
// Used by client-side stubs.

package host_transport_nodes

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type ModulesClient interface {

	// Get the module details of a host transport node
	//
	// @param siteIdParam (required)
	// @param enforcementpointIdParam (required)
	// @param hostTransportNodeIdParam (required)
	// @return com.vmware.nsx_policy.model.SoftwareModuleResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(siteIdParam string, enforcementpointIdParam string, hostTransportNodeIdParam string) (model.SoftwareModuleResult, error)
}

type modulesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewModulesClient(connector client.Connector) *modulesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.sites.enforcement_points.host_transport_nodes.modules")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get": core.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	mIface := modulesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &mIface
}

func (mIface *modulesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := mIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (mIface *modulesClient) Get(siteIdParam string, enforcementpointIdParam string, hostTransportNodeIdParam string) (model.SoftwareModuleResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(modulesGetInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("HostTransportNodeId", hostTransportNodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.SoftwareModuleResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := modulesGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.host_transport_nodes.modules", "get", inputDataValue, executionContext)
	var emptyOutput model.SoftwareModuleResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), modulesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.SoftwareModuleResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
