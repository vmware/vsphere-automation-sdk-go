
/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Client stubs for service: Resource
 * Functions that implement the generated ResourceClient interface
 */


package metamodel

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/log"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
)

type DefaultResourceClient struct {
	interfaceName       string
	interfaceDefinition core.InterfaceDefinition
	methodIdentifiers   []core.MethodIdentifier
	methodNameToDefMap  map[string]*core.MethodDefinition
	errorBindingMap     map[string]bindings.BindingType
	interfaceIdentifier core.InterfaceIdentifier
	connector           client.Connector
}

func NewDefaultResourceClient(connector client.Connector) *DefaultResourceClient {
	interfaceName := "com.vmware.vapi.metadata.metamodel.resource"
	interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
	methodIdentifiers := []core.MethodIdentifier{
		core.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorBindingMap := make(map[string]bindings.BindingType)
	errorBindingMap[errors.AlreadyExists{}.Error()] = errors.AlreadyExistsBindingType()
	errorBindingMap[errors.AlreadyInDesiredState{}.Error()] = errors.AlreadyInDesiredStateBindingType()
	errorBindingMap[errors.Canceled{}.Error()] = errors.CanceledBindingType()
	errorBindingMap[errors.ConcurrentChange{}.Error()] = errors.ConcurrentChangeBindingType()
	errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errorBindingMap[errors.FeatureInUse{}.Error()] = errors.FeatureInUseBindingType()
	errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	errorBindingMap[errors.InvalidElementConfiguration{}.Error()] = errors.InvalidElementConfigurationBindingType()
	errorBindingMap[errors.InvalidElementType{}.Error()] = errors.InvalidElementTypeBindingType()
	errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errorBindingMap[errors.NotAllowedInCurrentState{}.Error()] = errors.NotAllowedInCurrentStateBindingType()
	errorBindingMap[errors.OperationNotFound{}.Error()] = errors.OperationNotFoundBindingType()
	errorBindingMap[errors.ResourceBusy{}.Error()] = errors.ResourceBusyBindingType()
	errorBindingMap[errors.ResourceInUse{}.Error()] = errors.ResourceInUseBindingType()
	errorBindingMap[errors.ResourceInaccessible{}.Error()] = errors.ResourceInaccessibleBindingType()
	errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errorBindingMap[errors.TimedOut{}.Error()] = errors.TimedOutBindingType()
	errorBindingMap[errors.UnableToAllocateResource{}.Error()] = errors.UnableToAllocateResourceBindingType()
	errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errorBindingMap[errors.UnexpectedInput{}.Error()] = errors.UnexpectedInputBindingType()
	errorBindingMap[errors.Unsupported{}.Error()] = errors.UnsupportedBindingType()
	errorBindingMap[errors.UnverifiedPeer{}.Error()] = errors.UnverifiedPeerBindingType()


	rIface := DefaultResourceClient{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier: interfaceIdentifier, connector: connector}
	rIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
	rIface.methodNameToDefMap["list"] = rIface.listMethodDefinition()
	return &rIface
}

func (rIface *DefaultResourceClient) List() (map[string]bool, error) {
	typeConverter := rIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(rIface.interfaceIdentifier, "list")
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
	executionContext := rIface.connector.NewExecutionContext()
	methodResult := rIface.Invoke(executionContext, methodIdentifier, inputDataValue)
	var emptyOutput map[string]bool
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), resourceListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(map[string]bool), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}


func (rIface *DefaultResourceClient) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
	methodResult := rIface.connector.GetApiProvider().Invoke(rIface.interfaceName, methodId.Name(), inputDataValue, ctx)
	return methodResult
}


func (rIface *DefaultResourceClient) listMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(rIface.interfaceName)
	typeConverter := rIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(resourceListInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(resourceListOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultResourceClient.list method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultResourceClient.list method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "list")
	errorDefinitions := make([]data.ErrorDefinition, 0)

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}
