
/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Client stubs for service: DownloadSession
 * Functions that implement the generated DownloadSessionClient interface
 */


package item

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/log"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/client"
)

type DefaultDownloadSessionClient struct {
	interfaceName       string
	interfaceDefinition core.InterfaceDefinition
	methodIdentifiers   []core.MethodIdentifier
	methodNameToDefMap  map[string]*core.MethodDefinition
	errorBindingMap     map[string]bindings.BindingType
	interfaceIdentifier core.InterfaceIdentifier
	connector           client.Connector
}

func NewDefaultDownloadSessionClient(connector client.Connector) *DefaultDownloadSessionClient {
	interfaceName := "com.vmware.content.library.item.download_session"
	interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
	methodIdentifiers := []core.MethodIdentifier{
		core.NewMethodIdentifier(interfaceIdentifier, "create"),
		core.NewMethodIdentifier(interfaceIdentifier, "get"),
		core.NewMethodIdentifier(interfaceIdentifier, "list"),
		core.NewMethodIdentifier(interfaceIdentifier, "keepAlive"),
		core.NewMethodIdentifier(interfaceIdentifier, "cancel"),
		core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		core.NewMethodIdentifier(interfaceIdentifier, "fail"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorBindingMap := make(map[string]bindings.BindingType)
	errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	errorBindingMap[errors.OperationNotFound{}.Error()] = errors.OperationNotFoundBindingType()
	errorBindingMap[errors.UnexpectedInput{}.Error()] = errors.UnexpectedInputBindingType()
	errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errorBindingMap[errors.TimedOut{}.Error()] = errors.TimedOutBindingType()
	dIface := DefaultDownloadSessionClient{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier: interfaceIdentifier, connector: connector}
	dIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
	dIface.methodNameToDefMap["create"] = dIface.createMethodDefinition()
	dIface.methodNameToDefMap["get"] = dIface.getMethodDefinition()
	dIface.methodNameToDefMap["list"] = dIface.listMethodDefinition()
	dIface.methodNameToDefMap["keep_alive"] = dIface.keepAliveMethodDefinition()
	dIface.methodNameToDefMap["cancel"] = dIface.cancelMethodDefinition()
	dIface.methodNameToDefMap["delete"] = dIface.deleteMethodDefinition()
	dIface.methodNameToDefMap["fail"] = dIface.failMethodDefinition()
	return &dIface
}

func (dIface *DefaultDownloadSessionClient) Create(clientTokenParam *string, createSpecParam DownloadSessionModel) (string, error) {
	typeConverter := dIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(dIface.interfaceIdentifier, "create")
	sv := bindings.NewStructValueBuilder(downloadSessionCreateInputType(), typeConverter)
	sv.AddStructField("ClientToken", clientTokenParam)
	sv.AddStructField("CreateSpec", createSpecParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput string
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := downloadSessionCreateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	dIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := dIface.Invoke(dIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput string
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), downloadSessionCreateOutputType())
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

func (dIface *DefaultDownloadSessionClient) Get(downloadSessionIdParam string) (DownloadSessionModel, error) {
	typeConverter := dIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(dIface.interfaceIdentifier, "get")
	sv := bindings.NewStructValueBuilder(downloadSessionGetInputType(), typeConverter)
	sv.AddStructField("DownloadSessionId", downloadSessionIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput DownloadSessionModel
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := downloadSessionGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	dIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := dIface.Invoke(dIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput DownloadSessionModel
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), downloadSessionGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(DownloadSessionModel), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (dIface *DefaultDownloadSessionClient) List(libraryItemIdParam *string) ([]string, error) {
	typeConverter := dIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(dIface.interfaceIdentifier, "list")
	sv := bindings.NewStructValueBuilder(downloadSessionListInputType(), typeConverter)
	sv.AddStructField("LibraryItemId", libraryItemIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []string
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := downloadSessionListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	dIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := dIface.Invoke(dIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput []string
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), downloadSessionListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.([]string), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (dIface *DefaultDownloadSessionClient) KeepAlive(downloadSessionIdParam string, progressParam *int64) error {
	typeConverter := dIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(dIface.interfaceIdentifier, "keep_alive")
	sv := bindings.NewStructValueBuilder(downloadSessionKeepAliveInputType(), typeConverter)
	sv.AddStructField("DownloadSessionId", downloadSessionIdParam)
	sv.AddStructField("Progress", progressParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := downloadSessionKeepAliveRestMetadata()
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

func (dIface *DefaultDownloadSessionClient) Cancel(downloadSessionIdParam string) error {
	typeConverter := dIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(dIface.interfaceIdentifier, "cancel")
	sv := bindings.NewStructValueBuilder(downloadSessionCancelInputType(), typeConverter)
	sv.AddStructField("DownloadSessionId", downloadSessionIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := downloadSessionCancelRestMetadata()
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

func (dIface *DefaultDownloadSessionClient) Delete(downloadSessionIdParam string) error {
	typeConverter := dIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(dIface.interfaceIdentifier, "delete")
	sv := bindings.NewStructValueBuilder(downloadSessionDeleteInputType(), typeConverter)
	sv.AddStructField("DownloadSessionId", downloadSessionIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := downloadSessionDeleteRestMetadata()
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

func (dIface *DefaultDownloadSessionClient) Fail(downloadSessionIdParam string, clientErrorMessageParam string) error {
	typeConverter := dIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(dIface.interfaceIdentifier, "fail")
	sv := bindings.NewStructValueBuilder(downloadSessionFailInputType(), typeConverter)
	sv.AddStructField("DownloadSessionId", downloadSessionIdParam)
	sv.AddStructField("ClientErrorMessage", clientErrorMessageParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := downloadSessionFailRestMetadata()
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


func (dIface *DefaultDownloadSessionClient) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
	methodResult := dIface.connector.GetApiProvider().Invoke(dIface.interfaceName, methodId.Name(), inputDataValue, ctx)
	return methodResult
}


func (dIface *DefaultDownloadSessionClient) createMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(dIface.interfaceName)
	typeConverter := dIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(downloadSessionCreateInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(downloadSessionCreateOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDownloadSessionClient.create method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDownloadSessionClient.create method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "create")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	dIface.errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidArgumentBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDownloadSessionClient.create method's errors.InvalidArgument error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	dIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDownloadSessionClient.create method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (dIface *DefaultDownloadSessionClient) getMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(dIface.interfaceName)
	typeConverter := dIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(downloadSessionGetInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(downloadSessionGetOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDownloadSessionClient.get method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDownloadSessionClient.get method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "get")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	dIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDownloadSessionClient.get method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (dIface *DefaultDownloadSessionClient) listMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(dIface.interfaceName)
	typeConverter := dIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(downloadSessionListInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(downloadSessionListOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDownloadSessionClient.list method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDownloadSessionClient.list method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "list")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	dIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDownloadSessionClient.list method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (dIface *DefaultDownloadSessionClient) keepAliveMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(dIface.interfaceName)
	typeConverter := dIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(downloadSessionKeepAliveInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(downloadSessionKeepAliveOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDownloadSessionClient.keepAlive method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDownloadSessionClient.keepAlive method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "keepAlive")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	dIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDownloadSessionClient.keepAlive method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	dIface.errorBindingMap[errors.NotAllowedInCurrentState{}.Error()] = errors.NotAllowedInCurrentStateBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotAllowedInCurrentStateBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDownloadSessionClient.keepAlive method's errors.NotAllowedInCurrentState error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (dIface *DefaultDownloadSessionClient) cancelMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(dIface.interfaceName)
	typeConverter := dIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(downloadSessionCancelInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(downloadSessionCancelOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDownloadSessionClient.cancel method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDownloadSessionClient.cancel method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "cancel")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	dIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDownloadSessionClient.cancel method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	dIface.errorBindingMap[errors.NotAllowedInCurrentState{}.Error()] = errors.NotAllowedInCurrentStateBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotAllowedInCurrentStateBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDownloadSessionClient.cancel method's errors.NotAllowedInCurrentState error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (dIface *DefaultDownloadSessionClient) deleteMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(dIface.interfaceName)
	typeConverter := dIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(downloadSessionDeleteInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(downloadSessionDeleteOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDownloadSessionClient.delete method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDownloadSessionClient.delete method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "delete")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	dIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDownloadSessionClient.delete method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (dIface *DefaultDownloadSessionClient) failMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(dIface.interfaceName)
	typeConverter := dIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(downloadSessionFailInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(downloadSessionFailOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDownloadSessionClient.fail method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDownloadSessionClient.fail method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "fail")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	dIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDownloadSessionClient.fail method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	dIface.errorBindingMap[errors.NotAllowedInCurrentState{}.Error()] = errors.NotAllowedInCurrentStateBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotAllowedInCurrentStateBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultDownloadSessionClient.fail method's errors.NotAllowedInCurrentState error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}
