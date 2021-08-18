// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: State
// Used by client-side stubs.

package transport_node_collections

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type StateClient interface {

	// Returns the state of transport node collection based on the states of transport nodes of the hosts which are part of compute collection.
	//
	// @param siteIdParam (required)
	// @param enforcementpointIdParam (required)
	// @param transportNodeCollectionIdParam (required)
	// @return com.vmware.nsx_policy.model.TransportNodeCollectionState
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(siteIdParam string, enforcementpointIdParam string, transportNodeCollectionIdParam string) (model.TransportNodeCollectionState, error)
}

type stateClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewStateClient(connector client.Connector) *stateClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.sites.enforcement_points.transport_node_collections.state")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get": core.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := stateClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *stateClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *stateClient) Get(siteIdParam string, enforcementpointIdParam string, transportNodeCollectionIdParam string) (model.TransportNodeCollectionState, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(stateGetInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("TransportNodeCollectionId", transportNodeCollectionIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.TransportNodeCollectionState
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := stateGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.transport_node_collections.state", "get", inputDataValue, executionContext)
	var emptyOutput model.TransportNodeCollectionState
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), stateGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.TransportNodeCollectionState), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
