
/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Client stubs for service: Hardware
 * Functions that implement the generated HardwareClient interface
 */


package vm

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vsphere/vcenter/vm/Hardware"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/log"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/client"
)

type DefaultHardwareClient struct {
	interfaceName       string
	interfaceDefinition core.InterfaceDefinition
	methodIdentifiers   []core.MethodIdentifier
	methodNameToDefMap  map[string]*core.MethodDefinition
	errorBindingMap     map[string]bindings.BindingType
	interfaceIdentifier core.InterfaceIdentifier
	connector           client.Connector
}

func NewDefaultHardwareClient(connector client.Connector) *DefaultHardwareClient {
	interfaceName := "com.vmware.vcenter.vm.hardware"
	interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
	methodIdentifiers := []core.MethodIdentifier{
		core.NewMethodIdentifier(interfaceIdentifier, "get"),
		core.NewMethodIdentifier(interfaceIdentifier, "update"),
		core.NewMethodIdentifier(interfaceIdentifier, "upgrade"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorBindingMap := make(map[string]bindings.BindingType)
	errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	errorBindingMap[errors.OperationNotFound{}.Error()] = errors.OperationNotFoundBindingType()
	errorBindingMap[errors.UnexpectedInput{}.Error()] = errors.UnexpectedInputBindingType()
	errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errorBindingMap[errors.TimedOut{}.Error()] = errors.TimedOutBindingType()
	hIface := DefaultHardwareClient{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier: interfaceIdentifier, connector: connector}
	hIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
	hIface.methodNameToDefMap["get"] = hIface.getMethodDefinition()
	hIface.methodNameToDefMap["update"] = hIface.updateMethodDefinition()
	hIface.methodNameToDefMap["upgrade"] = hIface.upgradeMethodDefinition()
	return &hIface
}

func (hIface *DefaultHardwareClient) Get(vmParam string) (HardwareInfo, error) {
	typeConverter := hIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(hIface.interfaceIdentifier, "get")
	sv := bindings.NewStructValueBuilder(hardwareGetInputType(), typeConverter)
	sv.AddStructField("Vm", vmParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput HardwareInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := hardwareGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	hIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := hIface.Invoke(hIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput HardwareInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), hardwareGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(HardwareInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), hIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (hIface *DefaultHardwareClient) Update(vmParam string, specParam HardwareUpdateSpec) error {
	typeConverter := hIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(hIface.interfaceIdentifier, "update")
	sv := bindings.NewStructValueBuilder(hardwareUpdateInputType(), typeConverter)
	sv.AddStructField("Vm", vmParam)
	sv.AddStructField("Spec", specParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := hardwareUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	hIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := hIface.Invoke(hIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), hIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (hIface *DefaultHardwareClient) Upgrade(vmParam string, versionParam *Hardware.HardwareVersion) error {
	typeConverter := hIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(hIface.interfaceIdentifier, "upgrade")
	sv := bindings.NewStructValueBuilder(hardwareUpgradeInputType(), typeConverter)
	sv.AddStructField("Vm", vmParam)
	sv.AddStructField("Version", versionParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := hardwareUpgradeRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	hIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := hIface.Invoke(hIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), hIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}


func (hIface *DefaultHardwareClient) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
	methodResult := hIface.connector.GetApiProvider().Invoke(hIface.interfaceName, methodId.Name(), inputDataValue, ctx)
	return methodResult
}


func (hIface *DefaultHardwareClient) getMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(hIface.interfaceName)
	typeConverter := hIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(hardwareGetInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(hardwareGetOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultHardwareClient.get method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultHardwareClient.get method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "get")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	hIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultHardwareClient.get method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	hIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultHardwareClient.get method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	hIface.errorBindingMap[errors.ResourceInaccessible{}.Error()] = errors.ResourceInaccessibleBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.ResourceInaccessibleBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultHardwareClient.get method's errors.ResourceInaccessible error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	hIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultHardwareClient.get method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	hIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultHardwareClient.get method's errors.Unauthenticated error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
	hIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError6 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultHardwareClient.get method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError6).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (hIface *DefaultHardwareClient) updateMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(hIface.interfaceName)
	typeConverter := hIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(hardwareUpdateInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(hardwareUpdateOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultHardwareClient.update method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultHardwareClient.update method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "update")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	hIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultHardwareClient.update method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	hIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultHardwareClient.update method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	hIface.errorBindingMap[errors.AlreadyInDesiredState{}.Error()] = errors.AlreadyInDesiredStateBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.AlreadyInDesiredStateBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultHardwareClient.update method's errors.AlreadyInDesiredState error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	hIface.errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.InvalidArgumentBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultHardwareClient.update method's errors.InvalidArgument error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	hIface.errorBindingMap[errors.Unsupported{}.Error()] = errors.UnsupportedBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.UnsupportedBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultHardwareClient.update method's errors.Unsupported error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
	hIface.errorBindingMap[errors.ResourceBusy{}.Error()] = errors.ResourceBusyBindingType()
	errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.ResourceBusyBindingType())
	if errError6 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultHardwareClient.update method's errors.ResourceBusy error - %s",
			bindings.VAPIerrorsToError(errError6).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))
	hIface.errorBindingMap[errors.ResourceInaccessible{}.Error()] = errors.ResourceInaccessibleBindingType()
	errDef7, errError7 := typeConverter.ConvertToDataDefinition(errors.ResourceInaccessibleBindingType())
	if errError7 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultHardwareClient.update method's errors.ResourceInaccessible error - %s",
			bindings.VAPIerrorsToError(errError7).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef7.(data.ErrorDefinition))
	hIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef8, errError8 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError8 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultHardwareClient.update method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError8).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef8.(data.ErrorDefinition))
	hIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errDef9, errError9 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	if errError9 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultHardwareClient.update method's errors.Unauthenticated error - %s",
			bindings.VAPIerrorsToError(errError9).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef9.(data.ErrorDefinition))
	hIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef10, errError10 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError10 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultHardwareClient.update method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError10).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef10.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (hIface *DefaultHardwareClient) upgradeMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(hIface.interfaceName)
	typeConverter := hIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(hardwareUpgradeInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(hardwareUpgradeOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultHardwareClient.upgrade method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultHardwareClient.upgrade method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "upgrade")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	hIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultHardwareClient.upgrade method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	hIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultHardwareClient.upgrade method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	hIface.errorBindingMap[errors.NotAllowedInCurrentState{}.Error()] = errors.NotAllowedInCurrentStateBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.NotAllowedInCurrentStateBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultHardwareClient.upgrade method's errors.NotAllowedInCurrentState error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	hIface.errorBindingMap[errors.AlreadyInDesiredState{}.Error()] = errors.AlreadyInDesiredStateBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.AlreadyInDesiredStateBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultHardwareClient.upgrade method's errors.AlreadyInDesiredState error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	hIface.errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.InvalidArgumentBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultHardwareClient.upgrade method's errors.InvalidArgument error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
	hIface.errorBindingMap[errors.Unsupported{}.Error()] = errors.UnsupportedBindingType()
	errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.UnsupportedBindingType())
	if errError6 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultHardwareClient.upgrade method's errors.Unsupported error - %s",
			bindings.VAPIerrorsToError(errError6).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))
	hIface.errorBindingMap[errors.ResourceBusy{}.Error()] = errors.ResourceBusyBindingType()
	errDef7, errError7 := typeConverter.ConvertToDataDefinition(errors.ResourceBusyBindingType())
	if errError7 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultHardwareClient.upgrade method's errors.ResourceBusy error - %s",
			bindings.VAPIerrorsToError(errError7).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef7.(data.ErrorDefinition))
	hIface.errorBindingMap[errors.ResourceInaccessible{}.Error()] = errors.ResourceInaccessibleBindingType()
	errDef8, errError8 := typeConverter.ConvertToDataDefinition(errors.ResourceInaccessibleBindingType())
	if errError8 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultHardwareClient.upgrade method's errors.ResourceInaccessible error - %s",
			bindings.VAPIerrorsToError(errError8).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef8.(data.ErrorDefinition))
	hIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef9, errError9 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError9 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultHardwareClient.upgrade method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError9).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef9.(data.ErrorDefinition))
	hIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errDef10, errError10 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	if errError10 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultHardwareClient.upgrade method's errors.Unauthenticated error - %s",
			bindings.VAPIerrorsToError(errError10).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef10.(data.ErrorDefinition))
	hIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef11, errError11 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError11 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultHardwareClient.upgrade method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError11).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef11.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}
