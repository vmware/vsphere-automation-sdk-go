
/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Client stubs for service: Restore
 * Functions that implement the generated RestoreClient interface
 */


package cluster

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/log"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
)

type DefaultRestoreClient struct {
	interfaceName       string
	interfaceDefinition core.InterfaceDefinition
	methodIdentifiers   []core.MethodIdentifier
	methodNameToDefMap  map[string]*core.MethodDefinition
	errorBindingMap     map[string]bindings.BindingType
	interfaceIdentifier core.InterfaceIdentifier
	connector           client.Connector
}

func NewDefaultRestoreClient(connector client.Connector) *DefaultRestoreClient {
	interfaceName := "com.vmware.nsx_policy.cluster.restore"
	interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
	methodIdentifiers := []core.MethodIdentifier{
		core.NewMethodIdentifier(interfaceIdentifier, "advance"),
		core.NewMethodIdentifier(interfaceIdentifier, "cancel"),
		core.NewMethodIdentifier(interfaceIdentifier, "retry"),
		core.NewMethodIdentifier(interfaceIdentifier, "start"),
		core.NewMethodIdentifier(interfaceIdentifier, "suspend"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorBindingMap := make(map[string]bindings.BindingType)
	errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	errorBindingMap[errors.OperationNotFound{}.Error()] = errors.OperationNotFoundBindingType()
	errorBindingMap[errors.UnexpectedInput{}.Error()] = errors.UnexpectedInputBindingType()
	errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errorBindingMap[errors.TimedOut{}.Error()] = errors.TimedOutBindingType()
	rIface := DefaultRestoreClient{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier: interfaceIdentifier, connector: connector}
	rIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
	rIface.methodNameToDefMap["advance"] = rIface.advanceMethodDefinition()
	rIface.methodNameToDefMap["cancel"] = rIface.cancelMethodDefinition()
	rIface.methodNameToDefMap["retry"] = rIface.retryMethodDefinition()
	rIface.methodNameToDefMap["start"] = rIface.startMethodDefinition()
	rIface.methodNameToDefMap["suspend"] = rIface.suspendMethodDefinition()
	return &rIface
}

func (rIface *DefaultRestoreClient) Advance(advanceClusterRestoreRequestParam model.AdvanceClusterRestoreRequest) (model.ClusterRestoreStatus, error) {
	typeConverter := rIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(rIface.interfaceIdentifier, "advance")
	sv := bindings.NewStructValueBuilder(restoreAdvanceInputType(), typeConverter)
	sv.AddStructField("AdvanceClusterRestoreRequest", advanceClusterRestoreRequestParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ClusterRestoreStatus
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := restoreAdvanceRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	rIface.connector.SetConnectionMetadata(connectionMetadata)
	executionContext := rIface.connector.NewExecutionContext()
	methodResult := rIface.Invoke(executionContext, methodIdentifier, inputDataValue)
	var emptyOutput model.ClusterRestoreStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), restoreAdvanceOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ClusterRestoreStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (rIface *DefaultRestoreClient) Cancel() (model.ClusterRestoreStatus, error) {
	typeConverter := rIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(rIface.interfaceIdentifier, "cancel")
	sv := bindings.NewStructValueBuilder(restoreCancelInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ClusterRestoreStatus
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := restoreCancelRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	rIface.connector.SetConnectionMetadata(connectionMetadata)
	executionContext := rIface.connector.NewExecutionContext()
	methodResult := rIface.Invoke(executionContext, methodIdentifier, inputDataValue)
	var emptyOutput model.ClusterRestoreStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), restoreCancelOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ClusterRestoreStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (rIface *DefaultRestoreClient) Retry() (model.ClusterRestoreStatus, error) {
	typeConverter := rIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(rIface.interfaceIdentifier, "retry")
	sv := bindings.NewStructValueBuilder(restoreRetryInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ClusterRestoreStatus
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := restoreRetryRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	rIface.connector.SetConnectionMetadata(connectionMetadata)
	executionContext := rIface.connector.NewExecutionContext()
	methodResult := rIface.Invoke(executionContext, methodIdentifier, inputDataValue)
	var emptyOutput model.ClusterRestoreStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), restoreRetryOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ClusterRestoreStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (rIface *DefaultRestoreClient) Start(initiateClusterRestoreRequestParam model.InitiateClusterRestoreRequest) (model.ClusterRestoreStatus, error) {
	typeConverter := rIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(rIface.interfaceIdentifier, "start")
	sv := bindings.NewStructValueBuilder(restoreStartInputType(), typeConverter)
	sv.AddStructField("InitiateClusterRestoreRequest", initiateClusterRestoreRequestParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ClusterRestoreStatus
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := restoreStartRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	rIface.connector.SetConnectionMetadata(connectionMetadata)
	executionContext := rIface.connector.NewExecutionContext()
	methodResult := rIface.Invoke(executionContext, methodIdentifier, inputDataValue)
	var emptyOutput model.ClusterRestoreStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), restoreStartOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ClusterRestoreStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (rIface *DefaultRestoreClient) Suspend() (model.ClusterRestoreStatus, error) {
	typeConverter := rIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(rIface.interfaceIdentifier, "suspend")
	sv := bindings.NewStructValueBuilder(restoreSuspendInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ClusterRestoreStatus
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := restoreSuspendRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	rIface.connector.SetConnectionMetadata(connectionMetadata)
	executionContext := rIface.connector.NewExecutionContext()
	methodResult := rIface.Invoke(executionContext, methodIdentifier, inputDataValue)
	var emptyOutput model.ClusterRestoreStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), restoreSuspendOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ClusterRestoreStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}


func (rIface *DefaultRestoreClient) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
	methodResult := rIface.connector.GetApiProvider().Invoke(rIface.interfaceName, methodId.Name(), inputDataValue, ctx)
	return methodResult
}


func (rIface *DefaultRestoreClient) advanceMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(rIface.interfaceName)
	typeConverter := rIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(restoreAdvanceInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(restoreAdvanceOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRestoreClient.advance method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRestoreClient.advance method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "advance")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	rIface.errorBindingMap[errors.ConcurrentChange{}.Error()] = errors.ConcurrentChangeBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ConcurrentChangeBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRestoreClient.advance method's errors.ConcurrentChange error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	rIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRestoreClient.advance method's errors.InvalidRequest error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	rIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRestoreClient.advance method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	rIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRestoreClient.advance method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	rIface.errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.InternalServerErrorBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRestoreClient.advance method's errors.InternalServerError error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
	rIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError6 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRestoreClient.advance method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError6).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (rIface *DefaultRestoreClient) cancelMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(rIface.interfaceName)
	typeConverter := rIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(restoreCancelInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(restoreCancelOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRestoreClient.cancel method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRestoreClient.cancel method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "cancel")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	rIface.errorBindingMap[errors.ConcurrentChange{}.Error()] = errors.ConcurrentChangeBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ConcurrentChangeBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRestoreClient.cancel method's errors.ConcurrentChange error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	rIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRestoreClient.cancel method's errors.InvalidRequest error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	rIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRestoreClient.cancel method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	rIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRestoreClient.cancel method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	rIface.errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.InternalServerErrorBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRestoreClient.cancel method's errors.InternalServerError error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
	rIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError6 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRestoreClient.cancel method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError6).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (rIface *DefaultRestoreClient) retryMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(rIface.interfaceName)
	typeConverter := rIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(restoreRetryInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(restoreRetryOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRestoreClient.retry method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRestoreClient.retry method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "retry")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	rIface.errorBindingMap[errors.ConcurrentChange{}.Error()] = errors.ConcurrentChangeBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ConcurrentChangeBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRestoreClient.retry method's errors.ConcurrentChange error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	rIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRestoreClient.retry method's errors.InvalidRequest error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	rIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRestoreClient.retry method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	rIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRestoreClient.retry method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	rIface.errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.InternalServerErrorBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRestoreClient.retry method's errors.InternalServerError error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
	rIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError6 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRestoreClient.retry method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError6).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (rIface *DefaultRestoreClient) startMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(rIface.interfaceName)
	typeConverter := rIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(restoreStartInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(restoreStartOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRestoreClient.start method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRestoreClient.start method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "start")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	rIface.errorBindingMap[errors.ConcurrentChange{}.Error()] = errors.ConcurrentChangeBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ConcurrentChangeBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRestoreClient.start method's errors.ConcurrentChange error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	rIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRestoreClient.start method's errors.InvalidRequest error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	rIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRestoreClient.start method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	rIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRestoreClient.start method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	rIface.errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.InternalServerErrorBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRestoreClient.start method's errors.InternalServerError error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
	rIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError6 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRestoreClient.start method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError6).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (rIface *DefaultRestoreClient) suspendMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(rIface.interfaceName)
	typeConverter := rIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(restoreSuspendInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(restoreSuspendOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRestoreClient.suspend method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRestoreClient.suspend method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "suspend")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	rIface.errorBindingMap[errors.ConcurrentChange{}.Error()] = errors.ConcurrentChangeBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ConcurrentChangeBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRestoreClient.suspend method's errors.ConcurrentChange error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	rIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRestoreClient.suspend method's errors.InvalidRequest error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	rIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRestoreClient.suspend method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	rIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRestoreClient.suspend method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	rIface.errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.InternalServerErrorBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRestoreClient.suspend method's errors.InternalServerError error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
	rIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError6 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRestoreClient.suspend method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError6).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}
