
/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Client stubs for service: Disk
 * Functions that implement the generated DiskClient interface
 */


package hardware

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vsphere/vcenter/vm/hardware/Disk"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/log"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/client"
)

type DefaultDiskClient struct {
	interfaceName       string
	interfaceDefinition core.InterfaceDefinition
	methodIdentifiers   []core.MethodIdentifier
	methodNameToDefMap  map[string]*core.MethodDefinition
	errorBindingMap     map[string]bindings.BindingType
	interfaceIdentifier core.InterfaceIdentifier
	connector           client.Connector
}

func NewDefaultDiskClient(connector client.Connector) *DefaultDiskClient {
	interfaceName := "com.vmware.vcenter.vm.hardware.disk"
	interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
	methodIdentifiers := []core.MethodIdentifier{
		core.NewMethodIdentifier(interfaceIdentifier, "list"),
		core.NewMethodIdentifier(interfaceIdentifier, "get"),
		core.NewMethodIdentifier(interfaceIdentifier, "create"),
		core.NewMethodIdentifier(interfaceIdentifier, "update"),
		core.NewMethodIdentifier(interfaceIdentifier, "delete"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorBindingMap := make(map[string]bindings.BindingType)
	errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	errorBindingMap[errors.OperationNotFound{}.Error()] = errors.OperationNotFoundBindingType()
	errorBindingMap[errors.UnexpectedInput{}.Error()] = errors.UnexpectedInputBindingType()
	errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errorBindingMap[errors.TimedOut{}.Error()] = errors.TimedOutBindingType()
	dIface := DefaultDiskClient{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier: interfaceIdentifier, connector: connector}
	dIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
	dIface.methodNameToDefMap["list"] = dIface.listMethodDefinition()
	dIface.methodNameToDefMap["get"] = dIface.getMethodDefinition()
	dIface.methodNameToDefMap["create"] = dIface.createMethodDefinition()
	dIface.methodNameToDefMap["update"] = dIface.updateMethodDefinition()
	dIface.methodNameToDefMap["delete"] = dIface.deleteMethodDefinition()
	return &dIface
}

func (dIface *DefaultDiskClient) List(vmParam string) ([]DiskSummary, error) {
	typeConverter := dIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(dIface.interfaceIdentifier, "list")
	sv := bindings.NewStructValueBuilder(diskListInputType(), typeConverter)
	sv.AddStructField("Vm", vmParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []DiskSummary
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := diskListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	dIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := dIface.Invoke(dIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput []DiskSummary
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), diskListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.([]DiskSummary), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (dIface *DefaultDiskClient) Get(vmParam string, diskParam string) (Disk.DiskInfo, error) {
	typeConverter := dIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(dIface.interfaceIdentifier, "get")
	sv := bindings.NewStructValueBuilder(diskGetInputType(), typeConverter)
	sv.AddStructField("Vm", vmParam)
	sv.AddStructField("Disk", diskParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput Disk.DiskInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := diskGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	dIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := dIface.Invoke(dIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput Disk.DiskInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), diskGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(Disk.DiskInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (dIface *DefaultDiskClient) Create(vmParam string, specParam Disk.DiskCreateSpec) (string, error) {
	typeConverter := dIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(dIface.interfaceIdentifier, "create")
	sv := bindings.NewStructValueBuilder(diskCreateInputType(), typeConverter)
	sv.AddStructField("Vm", vmParam)
	sv.AddStructField("Spec", specParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput string
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := diskCreateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	dIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := dIface.Invoke(dIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput string
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), diskCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(string), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (dIface *DefaultDiskClient) Update(vmParam string, diskParam string, specParam DiskUpdateSpec) error {
	typeConverter := dIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(dIface.interfaceIdentifier, "update")
	sv := bindings.NewStructValueBuilder(diskUpdateInputType(), typeConverter)
	sv.AddStructField("Vm", vmParam)
	sv.AddStructField("Disk", diskParam)
	sv.AddStructField("Spec", specParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := diskUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	dIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := dIface.Invoke(dIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (dIface *DefaultDiskClient) Delete(vmParam string, diskParam string) error {
	typeConverter := dIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(dIface.interfaceIdentifier, "delete")
	sv := bindings.NewStructValueBuilder(diskDeleteInputType(), typeConverter)
	sv.AddStructField("Vm", vmParam)
	sv.AddStructField("Disk", diskParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := diskDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	dIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := dIface.Invoke(dIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}


func (dIface *DefaultDiskClient) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
	methodResult := dIface.connector.GetApiProvider().Invoke(dIface.interfaceName, methodId.Name(), inputDataValue, ctx)
	return methodResult
}


func (dIface *DefaultDiskClient) listMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(dIface.interfaceName)
	typeConverter := dIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(diskListInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(diskListOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.list method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.list method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "list")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	dIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.list method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	dIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.list method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	dIface.errorBindingMap[errors.ResourceInaccessible{}.Error()] = errors.ResourceInaccessibleBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.ResourceInaccessibleBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.list method's errors.ResourceInaccessible error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	dIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.list method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	dIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.list method's errors.Unauthenticated error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
	dIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError6 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.list method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError6).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (dIface *DefaultDiskClient) getMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(dIface.interfaceName)
	typeConverter := dIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(diskGetInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(diskGetOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.get method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.get method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "get")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	dIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.get method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	dIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.get method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	dIface.errorBindingMap[errors.ResourceInaccessible{}.Error()] = errors.ResourceInaccessibleBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.ResourceInaccessibleBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.get method's errors.ResourceInaccessible error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	dIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.get method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	dIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.get method's errors.Unauthenticated error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
	dIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError6 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.get method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError6).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (dIface *DefaultDiskClient) createMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(dIface.interfaceName)
	typeConverter := dIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(diskCreateInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(diskCreateOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.create method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.create method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "create")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	dIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.create method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	dIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.create method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	dIface.errorBindingMap[errors.NotAllowedInCurrentState{}.Error()] = errors.NotAllowedInCurrentStateBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.NotAllowedInCurrentStateBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.create method's errors.NotAllowedInCurrentState error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	dIface.errorBindingMap[errors.UnableToAllocateResource{}.Error()] = errors.UnableToAllocateResourceBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.UnableToAllocateResourceBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.create method's errors.UnableToAllocateResource error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	dIface.errorBindingMap[errors.ResourceInUse{}.Error()] = errors.ResourceInUseBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.ResourceInUseBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.create method's errors.ResourceInUse error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
	dIface.errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.InvalidArgumentBindingType())
	if errError6 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.create method's errors.InvalidArgument error - %s",
			bindings.VAPIerrorsToError(errError6).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))
	dIface.errorBindingMap[errors.ResourceBusy{}.Error()] = errors.ResourceBusyBindingType()
	errDef7, errError7 := typeConverter.ConvertToDataDefinition(errors.ResourceBusyBindingType())
	if errError7 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.create method's errors.ResourceBusy error - %s",
			bindings.VAPIerrorsToError(errError7).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef7.(data.ErrorDefinition))
	dIface.errorBindingMap[errors.ResourceInaccessible{}.Error()] = errors.ResourceInaccessibleBindingType()
	errDef8, errError8 := typeConverter.ConvertToDataDefinition(errors.ResourceInaccessibleBindingType())
	if errError8 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.create method's errors.ResourceInaccessible error - %s",
			bindings.VAPIerrorsToError(errError8).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef8.(data.ErrorDefinition))
	dIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef9, errError9 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError9 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.create method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError9).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef9.(data.ErrorDefinition))
	dIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errDef10, errError10 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	if errError10 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.create method's errors.Unauthenticated error - %s",
			bindings.VAPIerrorsToError(errError10).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef10.(data.ErrorDefinition))
	dIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef11, errError11 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError11 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.create method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError11).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef11.(data.ErrorDefinition))
	dIface.errorBindingMap[errors.Unsupported{}.Error()] = errors.UnsupportedBindingType()
	errDef12, errError12 := typeConverter.ConvertToDataDefinition(errors.UnsupportedBindingType())
	if errError12 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.create method's errors.Unsupported error - %s",
			bindings.VAPIerrorsToError(errError12).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef12.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (dIface *DefaultDiskClient) updateMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(dIface.interfaceName)
	typeConverter := dIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(diskUpdateInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(diskUpdateOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.update method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.update method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "update")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	dIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.update method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	dIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.update method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	dIface.errorBindingMap[errors.NotAllowedInCurrentState{}.Error()] = errors.NotAllowedInCurrentStateBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.NotAllowedInCurrentStateBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.update method's errors.NotAllowedInCurrentState error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	dIface.errorBindingMap[errors.ResourceBusy{}.Error()] = errors.ResourceBusyBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.ResourceBusyBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.update method's errors.ResourceBusy error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	dIface.errorBindingMap[errors.ResourceInaccessible{}.Error()] = errors.ResourceInaccessibleBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.ResourceInaccessibleBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.update method's errors.ResourceInaccessible error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
	dIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError6 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.update method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError6).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))
	dIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errDef7, errError7 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	if errError7 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.update method's errors.Unauthenticated error - %s",
			bindings.VAPIerrorsToError(errError7).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef7.(data.ErrorDefinition))
	dIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef8, errError8 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError8 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.update method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError8).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef8.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (dIface *DefaultDiskClient) deleteMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(dIface.interfaceName)
	typeConverter := dIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(diskDeleteInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(diskDeleteOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.delete method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.delete method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "delete")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	dIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.delete method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	dIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.delete method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	dIface.errorBindingMap[errors.NotAllowedInCurrentState{}.Error()] = errors.NotAllowedInCurrentStateBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.NotAllowedInCurrentStateBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.delete method's errors.NotAllowedInCurrentState error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	dIface.errorBindingMap[errors.ResourceBusy{}.Error()] = errors.ResourceBusyBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.ResourceBusyBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.delete method's errors.ResourceBusy error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	dIface.errorBindingMap[errors.ResourceInaccessible{}.Error()] = errors.ResourceInaccessibleBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.ResourceInaccessibleBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.delete method's errors.ResourceInaccessible error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
	dIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError6 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.delete method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError6).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))
	dIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errDef7, errError7 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	if errError7 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.delete method's errors.Unauthenticated error - %s",
			bindings.VAPIerrorsToError(errError7).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef7.(data.ErrorDefinition))
	dIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef8, errError8 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError8 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDiskClient.delete method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError8).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef8.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}
