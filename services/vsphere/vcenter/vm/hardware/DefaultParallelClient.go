
/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Client stubs for service: Parallel
 * Functions that implement the generated ParallelClient interface
 */


package hardware

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/log"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/client"
)

type DefaultParallelClient struct {
	interfaceName       string
	interfaceDefinition core.InterfaceDefinition
	methodIdentifiers   []core.MethodIdentifier
	methodNameToDefMap  map[string]*core.MethodDefinition
	errorBindingMap     map[string]bindings.BindingType
	interfaceIdentifier core.InterfaceIdentifier
	connector           client.Connector
}

func NewDefaultParallelClient(connector client.Connector) *DefaultParallelClient {
	interfaceName := "com.vmware.vcenter.vm.hardware.parallel"
	interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
	methodIdentifiers := []core.MethodIdentifier{
		core.NewMethodIdentifier(interfaceIdentifier, "list"),
		core.NewMethodIdentifier(interfaceIdentifier, "get"),
		core.NewMethodIdentifier(interfaceIdentifier, "create"),
		core.NewMethodIdentifier(interfaceIdentifier, "update"),
		core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		core.NewMethodIdentifier(interfaceIdentifier, "connect"),
		core.NewMethodIdentifier(interfaceIdentifier, "disconnect"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorBindingMap := make(map[string]bindings.BindingType)
	errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	errorBindingMap[errors.OperationNotFound{}.Error()] = errors.OperationNotFoundBindingType()
	errorBindingMap[errors.UnexpectedInput{}.Error()] = errors.UnexpectedInputBindingType()
	errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errorBindingMap[errors.TimedOut{}.Error()] = errors.TimedOutBindingType()
	pIface := DefaultParallelClient{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier: interfaceIdentifier, connector: connector}
	pIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
	pIface.methodNameToDefMap["list"] = pIface.listMethodDefinition()
	pIface.methodNameToDefMap["get"] = pIface.getMethodDefinition()
	pIface.methodNameToDefMap["create"] = pIface.createMethodDefinition()
	pIface.methodNameToDefMap["update"] = pIface.updateMethodDefinition()
	pIface.methodNameToDefMap["delete"] = pIface.deleteMethodDefinition()
	pIface.methodNameToDefMap["connect"] = pIface.connectMethodDefinition()
	pIface.methodNameToDefMap["disconnect"] = pIface.disconnectMethodDefinition()
	return &pIface
}

func (pIface *DefaultParallelClient) List(vmParam string) ([]ParallelSummary, error) {
	typeConverter := pIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(pIface.interfaceIdentifier, "list")
	sv := bindings.NewStructValueBuilder(parallelListInputType(), typeConverter)
	sv.AddStructField("Vm", vmParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []ParallelSummary
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := parallelListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := pIface.Invoke(pIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput []ParallelSummary
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), parallelListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.([]ParallelSummary), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (pIface *DefaultParallelClient) Get(vmParam string, portParam string) (ParallelInfo, error) {
	typeConverter := pIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(pIface.interfaceIdentifier, "get")
	sv := bindings.NewStructValueBuilder(parallelGetInputType(), typeConverter)
	sv.AddStructField("Vm", vmParam)
	sv.AddStructField("Port", portParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput ParallelInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := parallelGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := pIface.Invoke(pIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput ParallelInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), parallelGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(ParallelInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (pIface *DefaultParallelClient) Create(vmParam string, specParam ParallelCreateSpec) (string, error) {
	typeConverter := pIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(pIface.interfaceIdentifier, "create")
	sv := bindings.NewStructValueBuilder(parallelCreateInputType(), typeConverter)
	sv.AddStructField("Vm", vmParam)
	sv.AddStructField("Spec", specParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput string
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := parallelCreateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := pIface.Invoke(pIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput string
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), parallelCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(string), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (pIface *DefaultParallelClient) Update(vmParam string, portParam string, specParam ParallelUpdateSpec) error {
	typeConverter := pIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(pIface.interfaceIdentifier, "update")
	sv := bindings.NewStructValueBuilder(parallelUpdateInputType(), typeConverter)
	sv.AddStructField("Vm", vmParam)
	sv.AddStructField("Port", portParam)
	sv.AddStructField("Spec", specParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := parallelUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := pIface.Invoke(pIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (pIface *DefaultParallelClient) Delete(vmParam string, portParam string) error {
	typeConverter := pIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(pIface.interfaceIdentifier, "delete")
	sv := bindings.NewStructValueBuilder(parallelDeleteInputType(), typeConverter)
	sv.AddStructField("Vm", vmParam)
	sv.AddStructField("Port", portParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := parallelDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := pIface.Invoke(pIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (pIface *DefaultParallelClient) Connect(vmParam string, portParam string) error {
	typeConverter := pIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(pIface.interfaceIdentifier, "connect")
	sv := bindings.NewStructValueBuilder(parallelConnectInputType(), typeConverter)
	sv.AddStructField("Vm", vmParam)
	sv.AddStructField("Port", portParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := parallelConnectRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := pIface.Invoke(pIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (pIface *DefaultParallelClient) Disconnect(vmParam string, portParam string) error {
	typeConverter := pIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(pIface.interfaceIdentifier, "disconnect")
	sv := bindings.NewStructValueBuilder(parallelDisconnectInputType(), typeConverter)
	sv.AddStructField("Vm", vmParam)
	sv.AddStructField("Port", portParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := parallelDisconnectRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := pIface.Invoke(pIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}


func (pIface *DefaultParallelClient) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
	methodResult := pIface.connector.GetApiProvider().Invoke(pIface.interfaceName, methodId.Name(), inputDataValue, ctx)
	return methodResult
}


func (pIface *DefaultParallelClient) listMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(pIface.interfaceName)
	typeConverter := pIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(parallelListInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(parallelListOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.list method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.list method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "list")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	pIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.list method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.list method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.ResourceInaccessible{}.Error()] = errors.ResourceInaccessibleBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.ResourceInaccessibleBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.list method's errors.ResourceInaccessible error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.list method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.list method's errors.Unauthenticated error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError6 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.list method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError6).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (pIface *DefaultParallelClient) getMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(pIface.interfaceName)
	typeConverter := pIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(parallelGetInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(parallelGetOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.get method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.get method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "get")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	pIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.get method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.get method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.ResourceInaccessible{}.Error()] = errors.ResourceInaccessibleBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.ResourceInaccessibleBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.get method's errors.ResourceInaccessible error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.get method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.get method's errors.Unauthenticated error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError6 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.get method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError6).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (pIface *DefaultParallelClient) createMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(pIface.interfaceName)
	typeConverter := pIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(parallelCreateInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(parallelCreateOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.create method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.create method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "create")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	pIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.create method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.create method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.NotAllowedInCurrentState{}.Error()] = errors.NotAllowedInCurrentStateBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.NotAllowedInCurrentStateBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.create method's errors.NotAllowedInCurrentState error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.UnableToAllocateResource{}.Error()] = errors.UnableToAllocateResourceBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.UnableToAllocateResourceBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.create method's errors.UnableToAllocateResource error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.ResourceBusy{}.Error()] = errors.ResourceBusyBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.ResourceBusyBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.create method's errors.ResourceBusy error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.ResourceInaccessible{}.Error()] = errors.ResourceInaccessibleBindingType()
	errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.ResourceInaccessibleBindingType())
	if errError6 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.create method's errors.ResourceInaccessible error - %s",
			bindings.VAPIerrorsToError(errError6).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef7, errError7 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError7 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.create method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError7).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef7.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errDef8, errError8 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	if errError8 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.create method's errors.Unauthenticated error - %s",
			bindings.VAPIerrorsToError(errError8).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef8.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef9, errError9 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError9 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.create method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError9).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef9.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (pIface *DefaultParallelClient) updateMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(pIface.interfaceName)
	typeConverter := pIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(parallelUpdateInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(parallelUpdateOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.update method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.update method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "update")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	pIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.update method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.update method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.NotAllowedInCurrentState{}.Error()] = errors.NotAllowedInCurrentStateBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.NotAllowedInCurrentStateBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.update method's errors.NotAllowedInCurrentState error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.ResourceBusy{}.Error()] = errors.ResourceBusyBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.ResourceBusyBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.update method's errors.ResourceBusy error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.ResourceInaccessible{}.Error()] = errors.ResourceInaccessibleBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.ResourceInaccessibleBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.update method's errors.ResourceInaccessible error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError6 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.update method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError6).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errDef7, errError7 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	if errError7 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.update method's errors.Unauthenticated error - %s",
			bindings.VAPIerrorsToError(errError7).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef7.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef8, errError8 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError8 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.update method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError8).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef8.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (pIface *DefaultParallelClient) deleteMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(pIface.interfaceName)
	typeConverter := pIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(parallelDeleteInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(parallelDeleteOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.delete method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.delete method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "delete")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	pIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.delete method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.delete method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.NotAllowedInCurrentState{}.Error()] = errors.NotAllowedInCurrentStateBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.NotAllowedInCurrentStateBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.delete method's errors.NotAllowedInCurrentState error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.ResourceBusy{}.Error()] = errors.ResourceBusyBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.ResourceBusyBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.delete method's errors.ResourceBusy error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.ResourceInaccessible{}.Error()] = errors.ResourceInaccessibleBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.ResourceInaccessibleBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.delete method's errors.ResourceInaccessible error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError6 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.delete method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError6).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errDef7, errError7 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	if errError7 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.delete method's errors.Unauthenticated error - %s",
			bindings.VAPIerrorsToError(errError7).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef7.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef8, errError8 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError8 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.delete method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError8).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef8.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (pIface *DefaultParallelClient) connectMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(pIface.interfaceName)
	typeConverter := pIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(parallelConnectInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(parallelConnectOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.connect method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.connect method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "connect")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	pIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.connect method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.connect method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.AlreadyInDesiredState{}.Error()] = errors.AlreadyInDesiredStateBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.AlreadyInDesiredStateBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.connect method's errors.AlreadyInDesiredState error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.NotAllowedInCurrentState{}.Error()] = errors.NotAllowedInCurrentStateBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.NotAllowedInCurrentStateBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.connect method's errors.NotAllowedInCurrentState error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.ResourceBusy{}.Error()] = errors.ResourceBusyBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.ResourceBusyBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.connect method's errors.ResourceBusy error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.ResourceInaccessible{}.Error()] = errors.ResourceInaccessibleBindingType()
	errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.ResourceInaccessibleBindingType())
	if errError6 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.connect method's errors.ResourceInaccessible error - %s",
			bindings.VAPIerrorsToError(errError6).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef7, errError7 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError7 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.connect method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError7).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef7.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errDef8, errError8 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	if errError8 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.connect method's errors.Unauthenticated error - %s",
			bindings.VAPIerrorsToError(errError8).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef8.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef9, errError9 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError9 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.connect method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError9).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef9.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (pIface *DefaultParallelClient) disconnectMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(pIface.interfaceName)
	typeConverter := pIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(parallelDisconnectInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(parallelDisconnectOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.disconnect method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.disconnect method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "disconnect")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	pIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.disconnect method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.disconnect method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.AlreadyInDesiredState{}.Error()] = errors.AlreadyInDesiredStateBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.AlreadyInDesiredStateBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.disconnect method's errors.AlreadyInDesiredState error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.NotAllowedInCurrentState{}.Error()] = errors.NotAllowedInCurrentStateBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.NotAllowedInCurrentStateBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.disconnect method's errors.NotAllowedInCurrentState error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.ResourceBusy{}.Error()] = errors.ResourceBusyBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.ResourceBusyBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.disconnect method's errors.ResourceBusy error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.ResourceInaccessible{}.Error()] = errors.ResourceInaccessibleBindingType()
	errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.ResourceInaccessibleBindingType())
	if errError6 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.disconnect method's errors.ResourceInaccessible error - %s",
			bindings.VAPIerrorsToError(errError6).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef7, errError7 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError7 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.disconnect method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError7).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef7.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errDef8, errError8 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	if errError8 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.disconnect method's errors.Unauthenticated error - %s",
			bindings.VAPIerrorsToError(errError8).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef8.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef9, errError9 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError9 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultParallelClient.disconnect method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError9).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef9.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}
