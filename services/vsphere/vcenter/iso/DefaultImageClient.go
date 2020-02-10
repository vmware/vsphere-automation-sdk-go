
/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Client stubs for service: Image
 * Functions that implement the generated ImageClient interface
 */


package iso

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/log"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/client"
)

type DefaultImageClient struct {
	interfaceName       string
	interfaceDefinition core.InterfaceDefinition
	methodIdentifiers   []core.MethodIdentifier
	methodNameToDefMap  map[string]*core.MethodDefinition
	errorBindingMap     map[string]bindings.BindingType
	interfaceIdentifier core.InterfaceIdentifier
	connector           client.Connector
}

func NewDefaultImageClient(connector client.Connector) *DefaultImageClient {
	interfaceName := "com.vmware.vcenter.iso.image"
	interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
	methodIdentifiers := []core.MethodIdentifier{
		core.NewMethodIdentifier(interfaceIdentifier, "mount"),
		core.NewMethodIdentifier(interfaceIdentifier, "unmount"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorBindingMap := make(map[string]bindings.BindingType)
	errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	errorBindingMap[errors.OperationNotFound{}.Error()] = errors.OperationNotFoundBindingType()
	errorBindingMap[errors.UnexpectedInput{}.Error()] = errors.UnexpectedInputBindingType()
	errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errorBindingMap[errors.TimedOut{}.Error()] = errors.TimedOutBindingType()
	iIface := DefaultImageClient{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier: interfaceIdentifier, connector: connector}
	iIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
	iIface.methodNameToDefMap["mount"] = iIface.mountMethodDefinition()
	iIface.methodNameToDefMap["unmount"] = iIface.unmountMethodDefinition()
	return &iIface
}

func (iIface *DefaultImageClient) Mount(libraryItemParam string, vmParam string) (string, error) {
	typeConverter := iIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(iIface.interfaceIdentifier, "mount")
	sv := bindings.NewStructValueBuilder(imageMountInputType(), typeConverter)
	sv.AddStructField("LibraryItem", libraryItemParam)
	sv.AddStructField("Vm", vmParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput string
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := imageMountRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	iIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := iIface.Invoke(iIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput string
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), imageMountOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(string), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (iIface *DefaultImageClient) Unmount(vmParam string, cdromParam string) error {
	typeConverter := iIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(iIface.interfaceIdentifier, "unmount")
	sv := bindings.NewStructValueBuilder(imageUnmountInputType(), typeConverter)
	sv.AddStructField("Vm", vmParam)
	sv.AddStructField("Cdrom", cdromParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := imageUnmountRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	iIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := iIface.Invoke(iIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}


func (iIface *DefaultImageClient) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
	methodResult := iIface.connector.GetApiProvider().Invoke(iIface.interfaceName, methodId.Name(), inputDataValue, ctx)
	return methodResult
}


func (iIface *DefaultImageClient) mountMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(iIface.interfaceName)
	typeConverter := iIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(imageMountInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(imageMountOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultImageClient.mount method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultImageClient.mount method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "mount")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	iIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultImageClient.mount method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	iIface.errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.InvalidArgumentBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultImageClient.mount method's errors.InvalidArgument error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	iIface.errorBindingMap[errors.NotAllowedInCurrentState{}.Error()] = errors.NotAllowedInCurrentStateBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.NotAllowedInCurrentStateBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultImageClient.mount method's errors.NotAllowedInCurrentState error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	iIface.errorBindingMap[errors.ResourceInaccessible{}.Error()] = errors.ResourceInaccessibleBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.ResourceInaccessibleBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultImageClient.mount method's errors.ResourceInaccessible error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (iIface *DefaultImageClient) unmountMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(iIface.interfaceName)
	typeConverter := iIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(imageUnmountInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(imageUnmountOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultImageClient.unmount method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultImageClient.unmount method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "unmount")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	iIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultImageClient.unmount method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	iIface.errorBindingMap[errors.NotAllowedInCurrentState{}.Error()] = errors.NotAllowedInCurrentStateBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotAllowedInCurrentStateBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultImageClient.unmount method's errors.NotAllowedInCurrentState error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}
