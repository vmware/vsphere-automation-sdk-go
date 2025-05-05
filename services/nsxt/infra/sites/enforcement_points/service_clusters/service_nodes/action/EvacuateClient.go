// Copyright (c) 2019-2025 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Evacuate
// Used by client-side stubs.

package action

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type EvacuateClient interface {

	// This API is designed to evacuate service contexts from a specified edge node.To successfully evacuate the service contexts from the edge service node with the given ID, it is crucial that this node only contains auto-allocated service contexts. If there are any manually allocated service contexts present on the service cluster member, the evacuation task will not be executed. It is also advisable to place the edge service node into maintenance mode before performing the evacuation through the API. If the edge service node is not already in maintenance mode, the API will automatically transition it into maintenance mode before carrying out the evacuation operation. To ensure high availability during the evacuation process, the service cluster should have at least two healthy service nodes.
	//
	// @param siteIdParam (required)
	// @param enforcementpointIdParam (required)
	// @param serviceClusterIdParam (required)
	// @param serviceNodeIdParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(siteIdParam string, enforcementpointIdParam string, serviceClusterIdParam string, serviceNodeIdParam string) error
}

type evacuateClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewEvacuateClient(connector vapiProtocolClient_.Connector) *evacuateClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.sites.enforcement_points.service_clusters.service_nodes.action.evacuate")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"create": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "create"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	eIface := evacuateClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &eIface
}

func (eIface *evacuateClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := eIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (eIface *evacuateClient) Create(siteIdParam string, enforcementpointIdParam string, serviceClusterIdParam string, serviceNodeIdParam string) error {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	operationRestMetaData := evacuateCreateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(evacuateCreateInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("ServiceClusterId", serviceClusterIdParam)
	sv.AddStructField("ServiceNodeId", serviceNodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.service_clusters.service_nodes.action.evacuate", "create", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
