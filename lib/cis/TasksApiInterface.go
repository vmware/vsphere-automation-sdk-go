// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Skeleton file for service: Tasks
// Implementation vapiCore_.ApiInterface which adapts API service implementation -
// Functions that implement the generated TasksProvider interface

package cis

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiDataSerializersCleanjson_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data/serializers/cleanjson"
	vapiL10n_ "github.com/vmware/vsphere-automation-sdk-go/runtime/l10n"
	vapiLib_ "github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	vapiLog_ "github.com/vmware/vsphere-automation-sdk-go/runtime/log"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	vapiProtocolServerRest_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/server/rest"
	vapiProtocolServerRestContextbuilder_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/server/rest/contextbuilder"
	"net/http"
)

type TasksApiInterface struct {
	interfaceName       string
	interfaceDefinition vapiCore_.InterfaceDefinition
	methodIdentifiers   map[string]vapiCore_.MethodIdentifier
	methodMetadataMap   map[string]vapiProtocol_.OperationMetadata
	impl                TasksProvider
}

func NewTasksApiInterface(impl TasksProvider) *TasksApiInterface {
	interfaceName := "com.vmware.cis.tasks"
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier(interfaceName)
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"cancel": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "cancel"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	tIface := TasksApiInterface{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, impl: impl}
	tIface.methodMetadataMap = map[string]vapiProtocol_.OperationMetadata{}
	tIface.methodMetadataMap["get"] = tIface.getMetadata(TasksGetOutputType())
	tIface.methodMetadataMap["list"] = tIface.listMetadata(TasksListOutputType())
	tIface.methodMetadataMap["cancel"] = tIface.cancelMetadata(TasksCancelOutputType())
	return &tIface
}

func (tIface *TasksApiInterface) get(ctx *vapiCore_.ExecutionContext, methodId vapiCore_.MethodIdentifier, inputDataValue vapiData_.DataValue) vapiCore_.MethodResult {
	acceptOptions := &vapiCore_.BindingAcceptOptions{
		MethodName: methodId.Name(),
		IsStream:   false,
		IsTaskOnly: false,
	}
	if ok, mr := vapiCore_.ValidateServerBindings(ctx, acceptOptions); !ok {
		return mr
	}

	typeConverter := vapiBindings_.NewTypeConverter()
	opMetadata := tIface.methodMetadataMap["get"]
	structValue := inputDataValue.(*vapiData_.StructValue)
	extr := vapiBindings_.NewStructValueExtractor(opMetadata.InputType(), structValue, typeConverter)
	inputInterface, taskParamErr := extr.ExtractValue("task")
	if taskParamErr != nil {
		return createTasksErrorInput(taskParamErr, "Invalid input")
	}
	var taskParam string
	taskParam, _ = inputInterface.(string)
	inputInterface, specParamErr := extr.ExtractValue("spec")
	if specParamErr != nil {
		return createTasksErrorInput(specParamErr, "Invalid input")
	}
	var specParam *TasksGetSpec
	specParam, _ = inputInterface.(*TasksGetSpec)

	result, err := tIface.impl.Get(taskParam, specParam, ctx)

	if err != nil {
		errType, exist := opMetadata.ErrorBindingMap()[err.Error()]
		if !exist {
			return createTasksErrorResult([]error{err}, "Invalid error returned by server")
		}
		errorValue, errorInError := typeConverter.ConvertToVapi(err, errType)
		if errorInError != nil {
			return createTasksErrorResult(errorInError, "Invalid error")
		}
		return vapiCore_.NewMethodResult(nil, errorValue.(*vapiData_.ErrorValue))
	}

	outputDataValue, errorInOutput := typeConverter.ConvertToVapi(result, opMetadata.OutputType())
	if errorInOutput != nil {
		return createTasksErrorResult(errorInOutput, "Invalid output")
	}
	return vapiCore_.NewMethodResult(outputDataValue, nil)
}

func (tIface *TasksApiInterface) list(ctx *vapiCore_.ExecutionContext, methodId vapiCore_.MethodIdentifier, inputDataValue vapiData_.DataValue) vapiCore_.MethodResult {
	acceptOptions := &vapiCore_.BindingAcceptOptions{
		MethodName: methodId.Name(),
		IsStream:   false,
		IsTaskOnly: false,
	}
	if ok, mr := vapiCore_.ValidateServerBindings(ctx, acceptOptions); !ok {
		return mr
	}

	typeConverter := vapiBindings_.NewTypeConverter()
	opMetadata := tIface.methodMetadataMap["list"]
	structValue := inputDataValue.(*vapiData_.StructValue)
	extr := vapiBindings_.NewStructValueExtractor(opMetadata.InputType(), structValue, typeConverter)
	inputInterface, filterSpecParamErr := extr.ExtractValue("filter_spec")
	if filterSpecParamErr != nil {
		return createTasksErrorInput(filterSpecParamErr, "Invalid input")
	}
	var filterSpecParam *TasksFilterSpec
	filterSpecParam, _ = inputInterface.(*TasksFilterSpec)
	inputInterface, resultSpecParamErr := extr.ExtractValue("result_spec")
	if resultSpecParamErr != nil {
		return createTasksErrorInput(resultSpecParamErr, "Invalid input")
	}
	var resultSpecParam *TasksGetSpec
	resultSpecParam, _ = inputInterface.(*TasksGetSpec)

	result, err := tIface.impl.List(filterSpecParam, resultSpecParam, ctx)

	if err != nil {
		errType, exist := opMetadata.ErrorBindingMap()[err.Error()]
		if !exist {
			return createTasksErrorResult([]error{err}, "Invalid error returned by server")
		}
		errorValue, errorInError := typeConverter.ConvertToVapi(err, errType)
		if errorInError != nil {
			return createTasksErrorResult(errorInError, "Invalid error")
		}
		return vapiCore_.NewMethodResult(nil, errorValue.(*vapiData_.ErrorValue))
	}

	outputDataValue, errorInOutput := typeConverter.ConvertToVapi(result, opMetadata.OutputType())
	if errorInOutput != nil {
		return createTasksErrorResult(errorInOutput, "Invalid output")
	}
	return vapiCore_.NewMethodResult(outputDataValue, nil)
}

func (tIface *TasksApiInterface) cancel(ctx *vapiCore_.ExecutionContext, methodId vapiCore_.MethodIdentifier, inputDataValue vapiData_.DataValue) vapiCore_.MethodResult {
	acceptOptions := &vapiCore_.BindingAcceptOptions{
		MethodName: methodId.Name(),
		IsStream:   false,
		IsTaskOnly: false,
	}
	if ok, mr := vapiCore_.ValidateServerBindings(ctx, acceptOptions); !ok {
		return mr
	}

	typeConverter := vapiBindings_.NewTypeConverter()
	opMetadata := tIface.methodMetadataMap["cancel"]
	structValue := inputDataValue.(*vapiData_.StructValue)
	extr := vapiBindings_.NewStructValueExtractor(opMetadata.InputType(), structValue, typeConverter)
	inputInterface, taskParamErr := extr.ExtractValue("task")
	if taskParamErr != nil {
		return createTasksErrorInput(taskParamErr, "Invalid input")
	}
	var taskParam string
	taskParam, _ = inputInterface.(string)

	var result interface{} = nil

	err := tIface.impl.Cancel(taskParam, ctx)

	if err != nil {
		errType, exist := opMetadata.ErrorBindingMap()[err.Error()]
		if !exist {
			return createTasksErrorResult([]error{err}, "Invalid error returned by server")
		}
		errorValue, errorInError := typeConverter.ConvertToVapi(err, errType)
		if errorInError != nil {
			return createTasksErrorResult(errorInError, "Invalid error")
		}
		return vapiCore_.NewMethodResult(nil, errorValue.(*vapiData_.ErrorValue))
	}

	outputDataValue, errorInOutput := typeConverter.ConvertToVapi(result, opMetadata.OutputType())
	if errorInOutput != nil {
		return createTasksErrorResult(errorInOutput, "Invalid output")
	}
	return vapiCore_.NewMethodResult(outputDataValue, nil)
}

func (tIface *TasksApiInterface) getMetadata(outputType vapiBindings_.BindingType) vapiProtocol_.OperationMetadata {
	methodDefinition := tIface.getMethodDefinition(outputType)

	errorBindingMap := map[string]vapiBindings_.BindingType{}
	errorBindingMap[vapiStdErrors_.Error{}.Error()] = vapiStdErrors_.ErrorBindingType()
	errorBindingMap[vapiStdErrors_.NotFound{}.Error()] = vapiStdErrors_.NotFoundBindingType()
	errorBindingMap[vapiStdErrors_.ResourceInaccessible{}.Error()] = vapiStdErrors_.ResourceInaccessibleBindingType()
	errorBindingMap[vapiStdErrors_.ServiceUnavailable{}.Error()] = vapiStdErrors_.ServiceUnavailableBindingType()
	errorBindingMap[vapiStdErrors_.Unauthenticated{}.Error()] = vapiStdErrors_.UnauthenticatedBindingType()
	errorBindingMap[vapiStdErrors_.Unauthorized{}.Error()] = vapiStdErrors_.UnauthorizedBindingType()

	return vapiProtocol_.NewOperationMetadata(
		methodDefinition,
		tasksGetInputType(),
		outputType,
		errorBindingMap,
		tasksGetRestMetadata())
}
func (tIface *TasksApiInterface) listMetadata(outputType vapiBindings_.BindingType) vapiProtocol_.OperationMetadata {
	methodDefinition := tIface.listMethodDefinition(outputType)

	errorBindingMap := map[string]vapiBindings_.BindingType{}
	errorBindingMap[vapiStdErrors_.InvalidArgument{}.Error()] = vapiStdErrors_.InvalidArgumentBindingType()
	errorBindingMap[vapiStdErrors_.ResourceInaccessible{}.Error()] = vapiStdErrors_.ResourceInaccessibleBindingType()
	errorBindingMap[vapiStdErrors_.ServiceUnavailable{}.Error()] = vapiStdErrors_.ServiceUnavailableBindingType()
	errorBindingMap[vapiStdErrors_.Unauthenticated{}.Error()] = vapiStdErrors_.UnauthenticatedBindingType()
	errorBindingMap[vapiStdErrors_.Unauthorized{}.Error()] = vapiStdErrors_.UnauthorizedBindingType()

	return vapiProtocol_.NewOperationMetadata(
		methodDefinition,
		tasksListInputType(),
		outputType,
		errorBindingMap,
		tasksListRestMetadata())
}
func (tIface *TasksApiInterface) cancelMetadata(outputType vapiBindings_.BindingType) vapiProtocol_.OperationMetadata {
	methodDefinition := tIface.cancelMethodDefinition(outputType)

	errorBindingMap := map[string]vapiBindings_.BindingType{}
	errorBindingMap[vapiStdErrors_.Error{}.Error()] = vapiStdErrors_.ErrorBindingType()
	errorBindingMap[vapiStdErrors_.NotAllowedInCurrentState{}.Error()] = vapiStdErrors_.NotAllowedInCurrentStateBindingType()
	errorBindingMap[vapiStdErrors_.NotFound{}.Error()] = vapiStdErrors_.NotFoundBindingType()
	errorBindingMap[vapiStdErrors_.ResourceInaccessible{}.Error()] = vapiStdErrors_.ResourceInaccessibleBindingType()
	errorBindingMap[vapiStdErrors_.ServiceUnavailable{}.Error()] = vapiStdErrors_.ServiceUnavailableBindingType()
	errorBindingMap[vapiStdErrors_.Unauthenticated{}.Error()] = vapiStdErrors_.UnauthenticatedBindingType()
	errorBindingMap[vapiStdErrors_.Unauthorized{}.Error()] = vapiStdErrors_.UnauthorizedBindingType()
	errorBindingMap[vapiStdErrors_.Unsupported{}.Error()] = vapiStdErrors_.UnsupportedBindingType()

	return vapiProtocol_.NewOperationMetadata(
		methodDefinition,
		tasksCancelInputType(),
		outputType,
		errorBindingMap,
		tasksCancelRestMetadata())
}

func (tIface *TasksApiInterface) getMethodDefinition(outputType vapiBindings_.BindingType) *vapiCore_.MethodDefinition {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier(tIface.interfaceName)
	typeConverter := vapiBindings_.NewTypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(tasksGetInputType())
	if inputError != nil {
		vapiLog_.Errorf("Error in ConvertToDataDefinition for TasksApiInterface.get method's input - %s",
			vapiBindings_.VAPIerrorsToError(inputError).Error())
		return nil
	}

	output, outputError := typeConverter.ConvertToDataDefinition(outputType)
	if outputError != nil {
		vapiLog_.Errorf("Error in ConvertToDataDefinition for TasksApiInterface.get method's output - %s",
			vapiBindings_.VAPIerrorsToError(outputError).Error())
		return nil
	}

	methodIdentifier := vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get")
	errorDefinitions := make([]vapiData_.ErrorDefinition, 0)
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(vapiStdErrors_.ErrorBindingType())
	if errError1 != nil {
		vapiLog_.Errorf("Error in ConvertToDataDefinition for TasksApiInterface.get method's vapiStdErrors_.Error error - %s",
			vapiBindings_.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(vapiData_.ErrorDefinition))
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(vapiStdErrors_.NotFoundBindingType())
	if errError2 != nil {
		vapiLog_.Errorf("Error in ConvertToDataDefinition for TasksApiInterface.get method's vapiStdErrors_.NotFound error - %s",
			vapiBindings_.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(vapiData_.ErrorDefinition))
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(vapiStdErrors_.ResourceInaccessibleBindingType())
	if errError3 != nil {
		vapiLog_.Errorf("Error in ConvertToDataDefinition for TasksApiInterface.get method's vapiStdErrors_.ResourceInaccessible error - %s",
			vapiBindings_.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(vapiData_.ErrorDefinition))
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(vapiStdErrors_.ServiceUnavailableBindingType())
	if errError4 != nil {
		vapiLog_.Errorf("Error in ConvertToDataDefinition for TasksApiInterface.get method's vapiStdErrors_.ServiceUnavailable error - %s",
			vapiBindings_.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(vapiData_.ErrorDefinition))
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(vapiStdErrors_.UnauthenticatedBindingType())
	if errError5 != nil {
		vapiLog_.Errorf("Error in ConvertToDataDefinition for TasksApiInterface.get method's vapiStdErrors_.Unauthenticated error - %s",
			vapiBindings_.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(vapiData_.ErrorDefinition))
	errDef6, errError6 := typeConverter.ConvertToDataDefinition(vapiStdErrors_.UnauthorizedBindingType())
	if errError6 != nil {
		vapiLog_.Errorf("Error in ConvertToDataDefinition for TasksApiInterface.get method's vapiStdErrors_.Unauthorized error - %s",
			vapiBindings_.VAPIerrorsToError(errError6).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef6.(vapiData_.ErrorDefinition))

	methodDefinition := vapiCore_.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}
func (tIface *TasksApiInterface) listMethodDefinition(outputType vapiBindings_.BindingType) *vapiCore_.MethodDefinition {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier(tIface.interfaceName)
	typeConverter := vapiBindings_.NewTypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(tasksListInputType())
	if inputError != nil {
		vapiLog_.Errorf("Error in ConvertToDataDefinition for TasksApiInterface.list method's input - %s",
			vapiBindings_.VAPIerrorsToError(inputError).Error())
		return nil
	}

	output, outputError := typeConverter.ConvertToDataDefinition(outputType)
	if outputError != nil {
		vapiLog_.Errorf("Error in ConvertToDataDefinition for TasksApiInterface.list method's output - %s",
			vapiBindings_.VAPIerrorsToError(outputError).Error())
		return nil
	}

	methodIdentifier := vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list")
	errorDefinitions := make([]vapiData_.ErrorDefinition, 0)
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(vapiStdErrors_.InvalidArgumentBindingType())
	if errError1 != nil {
		vapiLog_.Errorf("Error in ConvertToDataDefinition for TasksApiInterface.list method's vapiStdErrors_.InvalidArgument error - %s",
			vapiBindings_.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(vapiData_.ErrorDefinition))
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(vapiStdErrors_.ResourceInaccessibleBindingType())
	if errError2 != nil {
		vapiLog_.Errorf("Error in ConvertToDataDefinition for TasksApiInterface.list method's vapiStdErrors_.ResourceInaccessible error - %s",
			vapiBindings_.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(vapiData_.ErrorDefinition))
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(vapiStdErrors_.ServiceUnavailableBindingType())
	if errError3 != nil {
		vapiLog_.Errorf("Error in ConvertToDataDefinition for TasksApiInterface.list method's vapiStdErrors_.ServiceUnavailable error - %s",
			vapiBindings_.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(vapiData_.ErrorDefinition))
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(vapiStdErrors_.UnauthenticatedBindingType())
	if errError4 != nil {
		vapiLog_.Errorf("Error in ConvertToDataDefinition for TasksApiInterface.list method's vapiStdErrors_.Unauthenticated error - %s",
			vapiBindings_.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(vapiData_.ErrorDefinition))
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(vapiStdErrors_.UnauthorizedBindingType())
	if errError5 != nil {
		vapiLog_.Errorf("Error in ConvertToDataDefinition for TasksApiInterface.list method's vapiStdErrors_.Unauthorized error - %s",
			vapiBindings_.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(vapiData_.ErrorDefinition))

	methodDefinition := vapiCore_.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}
func (tIface *TasksApiInterface) cancelMethodDefinition(outputType vapiBindings_.BindingType) *vapiCore_.MethodDefinition {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier(tIface.interfaceName)
	typeConverter := vapiBindings_.NewTypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(tasksCancelInputType())
	if inputError != nil {
		vapiLog_.Errorf("Error in ConvertToDataDefinition for TasksApiInterface.cancel method's input - %s",
			vapiBindings_.VAPIerrorsToError(inputError).Error())
		return nil
	}

	output, outputError := typeConverter.ConvertToDataDefinition(outputType)
	if outputError != nil {
		vapiLog_.Errorf("Error in ConvertToDataDefinition for TasksApiInterface.cancel method's output - %s",
			vapiBindings_.VAPIerrorsToError(outputError).Error())
		return nil
	}

	methodIdentifier := vapiCore_.NewMethodIdentifier(interfaceIdentifier, "cancel")
	errorDefinitions := make([]vapiData_.ErrorDefinition, 0)
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(vapiStdErrors_.ErrorBindingType())
	if errError1 != nil {
		vapiLog_.Errorf("Error in ConvertToDataDefinition for TasksApiInterface.cancel method's vapiStdErrors_.Error error - %s",
			vapiBindings_.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(vapiData_.ErrorDefinition))
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(vapiStdErrors_.NotAllowedInCurrentStateBindingType())
	if errError2 != nil {
		vapiLog_.Errorf("Error in ConvertToDataDefinition for TasksApiInterface.cancel method's vapiStdErrors_.NotAllowedInCurrentState error - %s",
			vapiBindings_.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(vapiData_.ErrorDefinition))
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(vapiStdErrors_.NotFoundBindingType())
	if errError3 != nil {
		vapiLog_.Errorf("Error in ConvertToDataDefinition for TasksApiInterface.cancel method's vapiStdErrors_.NotFound error - %s",
			vapiBindings_.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(vapiData_.ErrorDefinition))
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(vapiStdErrors_.ResourceInaccessibleBindingType())
	if errError4 != nil {
		vapiLog_.Errorf("Error in ConvertToDataDefinition for TasksApiInterface.cancel method's vapiStdErrors_.ResourceInaccessible error - %s",
			vapiBindings_.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(vapiData_.ErrorDefinition))
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(vapiStdErrors_.ServiceUnavailableBindingType())
	if errError5 != nil {
		vapiLog_.Errorf("Error in ConvertToDataDefinition for TasksApiInterface.cancel method's vapiStdErrors_.ServiceUnavailable error - %s",
			vapiBindings_.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(vapiData_.ErrorDefinition))
	errDef6, errError6 := typeConverter.ConvertToDataDefinition(vapiStdErrors_.UnauthenticatedBindingType())
	if errError6 != nil {
		vapiLog_.Errorf("Error in ConvertToDataDefinition for TasksApiInterface.cancel method's vapiStdErrors_.Unauthenticated error - %s",
			vapiBindings_.VAPIerrorsToError(errError6).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef6.(vapiData_.ErrorDefinition))
	errDef7, errError7 := typeConverter.ConvertToDataDefinition(vapiStdErrors_.UnauthorizedBindingType())
	if errError7 != nil {
		vapiLog_.Errorf("Error in ConvertToDataDefinition for TasksApiInterface.cancel method's vapiStdErrors_.Unauthorized error - %s",
			vapiBindings_.VAPIerrorsToError(errError7).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef7.(vapiData_.ErrorDefinition))
	errDef8, errError8 := typeConverter.ConvertToDataDefinition(vapiStdErrors_.UnsupportedBindingType())
	if errError8 != nil {
		vapiLog_.Errorf("Error in ConvertToDataDefinition for TasksApiInterface.cancel method's vapiStdErrors_.Unsupported error - %s",
			vapiBindings_.VAPIerrorsToError(errError8).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef8.(vapiData_.ErrorDefinition))

	methodDefinition := vapiCore_.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func createTasksErrorResult(errs []error, msg string) vapiCore_.MethodResult {
	vapiLog_.Error(msg)
	errValue := vapiBindings_.CreateErrorValueFromMessages(vapiBindings_.INTERNAL_SERVER_ERROR_DEF, errs)
	return vapiCore_.NewMethodResult(nil, errValue)
}

func createTasksErrorInput(errs []error, msg string) vapiCore_.MethodResult {
	vapiLog_.Error(msg)
	errValue := vapiBindings_.CreateErrorValueFromMessages(vapiBindings_.INVALID_ARGUMENT_ERROR_DEF, errs)
	return vapiCore_.NewMethodResult(nil, errValue)
}

func (tIface *TasksApiInterface) Identifier() vapiCore_.InterfaceIdentifier {
	return vapiCore_.NewInterfaceIdentifier(tIface.interfaceName)
}

func (tIface *TasksApiInterface) Definition() vapiCore_.InterfaceDefinition {
	return tIface.interfaceDefinition
}

func (tIface *TasksApiInterface) MethodDefinition(mi vapiCore_.MethodIdentifier) *vapiCore_.MethodDefinition {
	if val, ok := tIface.methodMetadataMap[mi.Name()]; ok {
		return val.MethodDefinition()
	}
	return nil
}

func (tIface *TasksApiInterface) Invoke(ctx *vapiCore_.ExecutionContext, methodId vapiCore_.MethodIdentifier, inputDataValue vapiData_.DataValue) vapiCore_.MethodResult {
	switch methodId.Name() {
	case "get":
		return tIface.get(ctx, methodId, inputDataValue)
	case "list":
		return tIface.list(ctx, methodId, inputDataValue)
	case "cancel":
		return tIface.cancel(ctx, methodId, inputDataValue)
	}
	return createTasksErrorResult([]error{}, "Invalid method call")
}

func (tIface *TasksApiInterface) RegisterRoutesHandlers(router vapiProtocolServerRest_.Router, provider vapiCore_.APIProvider, appCtxBuilders []vapiProtocolServerRestContextbuilder_.ApplicationContextBuilder, secCtxBuilders []vapiProtocolServerRestContextbuilder_.SecurityContextBuilder, opts ...vapiProtocolServerRest_.RequestHandlerOption) {
	opts = append(opts, vapiProtocolServerRest_.WithApplicationContextBuilders(appCtxBuilders))
	opts = append(opts, vapiProtocolServerRest_.WithSecurityContextBuilders(secCtxBuilders))

	router.AddRoute().
		Methods("POST").
		Path("/cis/tasks").
		MatcherFunc(func(r *http.Request, m *vapiProtocolServerRest_.RouteMatch) bool {
			queryMap := vapiProtocolServerRest_.GetQuery(r)
			if val, ok := queryMap["action"]; ok && len(val) == 1 && val[0] != nil && *val[0] == "list" {
				return true
			}
			return false
		}).
		Handler(vapiProtocolServerRest_.NewRequestHandler(tIface.methodMetadataMap["list"], provider, opts...))

	router.AddRoute().
		Methods("POST").
		Path("/cis/tasks/{task}").
		MatcherFunc(func(r *http.Request, m *vapiProtocolServerRest_.RouteMatch) bool {
			queryMap := vapiProtocolServerRest_.GetQuery(r)
			if val, ok := queryMap["action"]; ok && len(val) == 1 && val[0] != nil && *val[0] == "cancel" {
				return true
			}
			return false
		}).
		Handler(vapiProtocolServerRest_.NewRequestHandler(tIface.methodMetadataMap["cancel"], provider, opts...))

	router.AddRoute().
		Methods("GET").
		Path("/cis/tasks/{task}").
		Handler(vapiProtocolServerRest_.NewRequestHandler(tIface.methodMetadataMap["get"], provider, opts...))

	router.AddNotFoundHandler(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		errValue := vapiBindings_.CreateErrorValueFromMessages(vapiBindings_.NOT_FOUND_ERROR_DEF, []error{vapiL10n_.NewRuntimeError("vapi.protocol.server.rest.http_not_found", nil)})
		jsonResponse, _ := vapiDataSerializersCleanjson_.NewDataValueToJsonEncoder().Encode(errValue)
		rw.Header().Set(vapiLib_.HTTP_CONTENT_TYPE_HEADER, vapiLib_.JSON_CONTENT_TYPE)
		rw.WriteHeader(404)
		rw.Write([]byte(jsonResponse))
	}))

	router.AddMethodNotAllowedHandler(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		errValue := vapiBindings_.CreateErrorValueFromMessages(vapiBindings_.NOT_FOUND_ERROR_DEF, []error{vapiL10n_.NewRuntimeError("vapi.protocol.server.rest.http_not_found", nil)})
		jsonResponse, _ := vapiDataSerializersCleanjson_.NewDataValueToJsonEncoder().Encode(errValue)
		rw.Header().Set(vapiLib_.HTTP_CONTENT_TYPE_HEADER, vapiLib_.JSON_CONTENT_TYPE)
		rw.WriteHeader(404)
		rw.Write([]byte(jsonResponse))
	}))
}
