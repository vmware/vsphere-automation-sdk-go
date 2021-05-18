// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Resource
// Used by client-side stubs.

package metamodel

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
)

const _ = core.SupportedByRuntimeVersion1

// The Resource interface provides methods to retrieve information about resource types.
//
//  A service is a logical grouping of operations that operate on an entity. Each entity is identifier by a namespace (or resource type) and an unique identifier.
type ResourceClient interface {

	// Returns the set of resource types present across all the service elements contained in all the package elements.
	// @return Set of resource types
	// The return value will contain identifiers for the resource type: ``com.vmware.vapi.resource``.
	List() (map[string]bool, error)
}

type resourceClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewResourceClient(connector client.Connector) *resourceClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.vapi.metadata.metamodel.resource")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"list": core.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	rIface := resourceClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &rIface
}

func (rIface *resourceClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := rIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (rIface *resourceClient) List() (map[string]bool, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(resourceListInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput map[string]bool
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := resourceListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	rIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.metamodel.resource", "list", inputDataValue, executionContext)
	var emptyOutput map[string]bool
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), resourceListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(map[string]bool), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
