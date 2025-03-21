// Copyright (c) 2019-2025 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Skeleton file for service: Session
// Implementation vapiCore_.ApiInterface which adapts API service implementation -
// Functions that implement the generated SessionProvider interface

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

type SessionApiInterface struct {
	interfaceName       string
	interfaceDefinition vapiCore_.InterfaceDefinition
	methodIdentifiers   map[string]vapiCore_.MethodIdentifier
	methodMetadataMap   map[string]vapiProtocol_.OperationMetadata
	impl                SessionProvider
}

func NewSessionApiInterface(impl SessionProvider) *SessionApiInterface {
	interfaceName := "com.vmware.cis.session"
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier(interfaceName)
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"create": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "create"),
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	sIface := SessionApiInterface{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, impl: impl}
	sIface.methodMetadataMap = map[string]vapiProtocol_.OperationMetadata{}
	sIface.methodMetadataMap["create"] = sIface.createMetadata(SessionCreateOutputType())
	sIface.methodMetadataMap["delete"] = sIface.deleteMetadata(SessionDeleteOutputType())
	sIface.methodMetadataMap["get"] = sIface.getMetadata(SessionGetOutputType())
	return &sIface
}

func (sIface *SessionApiInterface) create(ctx *vapiCore_.ExecutionContext, methodId vapiCore_.MethodIdentifier, inputDataValue vapiData_.DataValue) vapiCore_.MethodResult {
	acceptOptions := &vapiCore_.BindingAcceptOptions{
		MethodName: methodId.Name(),
		IsStream:   false,
		IsTaskOnly: false,
	}
	if ok, mr := vapiCore_.ValidateServerBindings(ctx, acceptOptions); !ok {
		return mr
	}

	typeConverter := vapiBindings_.NewTypeConverter()
	opMetadata := sIface.methodMetadataMap["create"]

	result, err := sIface.impl.Create(ctx)

	if err != nil {
		errType, exist := opMetadata.ErrorBindingMap()[err.Error()]
		if !exist {
			return createSessionErrorResult([]error{err}, "Invalid error returned by server")
		}
		errorValue, errorInError := typeConverter.ConvertToVapi(err, errType)
		if errorInError != nil {
			return createSessionErrorResult(errorInError, "Invalid error")
		}
		return vapiCore_.NewMethodResult(nil, errorValue.(*vapiData_.ErrorValue))
	}

	outputDataValue, errorInOutput := typeConverter.ConvertToVapi(result, opMetadata.OutputType())
	if errorInOutput != nil {
		return createSessionErrorResult(errorInOutput, "Invalid output")
	}
	return vapiCore_.NewMethodResult(outputDataValue, nil)
}

func (sIface *SessionApiInterface) delete(ctx *vapiCore_.ExecutionContext, methodId vapiCore_.MethodIdentifier, inputDataValue vapiData_.DataValue) vapiCore_.MethodResult {
	acceptOptions := &vapiCore_.BindingAcceptOptions{
		MethodName: methodId.Name(),
		IsStream:   false,
		IsTaskOnly: false,
	}
	if ok, mr := vapiCore_.ValidateServerBindings(ctx, acceptOptions); !ok {
		return mr
	}

	typeConverter := vapiBindings_.NewTypeConverter()
	opMetadata := sIface.methodMetadataMap["delete"]

	var result interface{} = nil

	err := sIface.impl.Delete(ctx)

	if err != nil {
		errType, exist := opMetadata.ErrorBindingMap()[err.Error()]
		if !exist {
			return createSessionErrorResult([]error{err}, "Invalid error returned by server")
		}
		errorValue, errorInError := typeConverter.ConvertToVapi(err, errType)
		if errorInError != nil {
			return createSessionErrorResult(errorInError, "Invalid error")
		}
		return vapiCore_.NewMethodResult(nil, errorValue.(*vapiData_.ErrorValue))
	}

	outputDataValue, errorInOutput := typeConverter.ConvertToVapi(result, opMetadata.OutputType())
	if errorInOutput != nil {
		return createSessionErrorResult(errorInOutput, "Invalid output")
	}
	return vapiCore_.NewMethodResult(outputDataValue, nil)
}

func (sIface *SessionApiInterface) get(ctx *vapiCore_.ExecutionContext, methodId vapiCore_.MethodIdentifier, inputDataValue vapiData_.DataValue) vapiCore_.MethodResult {
	acceptOptions := &vapiCore_.BindingAcceptOptions{
		MethodName: methodId.Name(),
		IsStream:   false,
		IsTaskOnly: false,
	}
	if ok, mr := vapiCore_.ValidateServerBindings(ctx, acceptOptions); !ok {
		return mr
	}

	typeConverter := vapiBindings_.NewTypeConverter()
	opMetadata := sIface.methodMetadataMap["get"]

	result, err := sIface.impl.Get(ctx)

	if err != nil {
		errType, exist := opMetadata.ErrorBindingMap()[err.Error()]
		if !exist {
			return createSessionErrorResult([]error{err}, "Invalid error returned by server")
		}
		errorValue, errorInError := typeConverter.ConvertToVapi(err, errType)
		if errorInError != nil {
			return createSessionErrorResult(errorInError, "Invalid error")
		}
		return vapiCore_.NewMethodResult(nil, errorValue.(*vapiData_.ErrorValue))
	}

	outputDataValue, errorInOutput := typeConverter.ConvertToVapi(result, opMetadata.OutputType())
	if errorInOutput != nil {
		return createSessionErrorResult(errorInOutput, "Invalid output")
	}
	return vapiCore_.NewMethodResult(outputDataValue, nil)
}

func (sIface *SessionApiInterface) createMetadata(outputType vapiBindings_.BindingType) vapiProtocol_.OperationMetadata {
	methodDefinition := sIface.createMethodDefinition(outputType)

	errorBindingMap := map[string]vapiBindings_.BindingType{}
	errorBindingMap[vapiStdErrors_.Unauthenticated{}.Error()] = vapiStdErrors_.UnauthenticatedBindingType()
	errorBindingMap[vapiStdErrors_.ServiceUnavailable{}.Error()] = vapiStdErrors_.ServiceUnavailableBindingType()

	return vapiProtocol_.NewOperationMetadata(
		methodDefinition,
		sessionCreateInputType(),
		outputType,
		errorBindingMap,
		sessionCreateRestMetadata())
}
func (sIface *SessionApiInterface) deleteMetadata(outputType vapiBindings_.BindingType) vapiProtocol_.OperationMetadata {
	methodDefinition := sIface.deleteMethodDefinition(outputType)

	errorBindingMap := map[string]vapiBindings_.BindingType{}
	errorBindingMap[vapiStdErrors_.Unauthenticated{}.Error()] = vapiStdErrors_.UnauthenticatedBindingType()
	errorBindingMap[vapiStdErrors_.ServiceUnavailable{}.Error()] = vapiStdErrors_.ServiceUnavailableBindingType()

	return vapiProtocol_.NewOperationMetadata(
		methodDefinition,
		sessionDeleteInputType(),
		outputType,
		errorBindingMap,
		sessionDeleteRestMetadata())
}
func (sIface *SessionApiInterface) getMetadata(outputType vapiBindings_.BindingType) vapiProtocol_.OperationMetadata {
	methodDefinition := sIface.getMethodDefinition(outputType)

	errorBindingMap := map[string]vapiBindings_.BindingType{}
	errorBindingMap[vapiStdErrors_.Unauthenticated{}.Error()] = vapiStdErrors_.UnauthenticatedBindingType()
	errorBindingMap[vapiStdErrors_.ServiceUnavailable{}.Error()] = vapiStdErrors_.ServiceUnavailableBindingType()

	return vapiProtocol_.NewOperationMetadata(
		methodDefinition,
		sessionGetInputType(),
		outputType,
		errorBindingMap,
		sessionGetRestMetadata())
}

func (sIface *SessionApiInterface) createMethodDefinition(outputType vapiBindings_.BindingType) *vapiCore_.MethodDefinition {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier(sIface.interfaceName)
	typeConverter := vapiBindings_.NewTypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(sessionCreateInputType())
	if inputError != nil {
		vapiLog_.Errorf("Error in ConvertToDataDefinition for SessionApiInterface.create method's input - %s",
			vapiBindings_.VAPIerrorsToError(inputError).Error())
		return nil
	}

	output, outputError := typeConverter.ConvertToDataDefinition(outputType)
	if outputError != nil {
		vapiLog_.Errorf("Error in ConvertToDataDefinition for SessionApiInterface.create method's output - %s",
			vapiBindings_.VAPIerrorsToError(outputError).Error())
		return nil
	}

	methodIdentifier := vapiCore_.NewMethodIdentifier(interfaceIdentifier, "create")
	errorDefinitions := make([]vapiData_.ErrorDefinition, 0)
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(vapiStdErrors_.UnauthenticatedBindingType())
	if errError1 != nil {
		vapiLog_.Errorf("Error in ConvertToDataDefinition for SessionApiInterface.create method's vapiStdErrors_.Unauthenticated error - %s",
			vapiBindings_.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(vapiData_.ErrorDefinition))
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(vapiStdErrors_.ServiceUnavailableBindingType())
	if errError2 != nil {
		vapiLog_.Errorf("Error in ConvertToDataDefinition for SessionApiInterface.create method's vapiStdErrors_.ServiceUnavailable error - %s",
			vapiBindings_.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(vapiData_.ErrorDefinition))

	methodDefinition := vapiCore_.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}
func (sIface *SessionApiInterface) deleteMethodDefinition(outputType vapiBindings_.BindingType) *vapiCore_.MethodDefinition {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier(sIface.interfaceName)
	typeConverter := vapiBindings_.NewTypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(sessionDeleteInputType())
	if inputError != nil {
		vapiLog_.Errorf("Error in ConvertToDataDefinition for SessionApiInterface.delete method's input - %s",
			vapiBindings_.VAPIerrorsToError(inputError).Error())
		return nil
	}

	output, outputError := typeConverter.ConvertToDataDefinition(outputType)
	if outputError != nil {
		vapiLog_.Errorf("Error in ConvertToDataDefinition for SessionApiInterface.delete method's output - %s",
			vapiBindings_.VAPIerrorsToError(outputError).Error())
		return nil
	}

	methodIdentifier := vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete")
	errorDefinitions := make([]vapiData_.ErrorDefinition, 0)
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(vapiStdErrors_.UnauthenticatedBindingType())
	if errError1 != nil {
		vapiLog_.Errorf("Error in ConvertToDataDefinition for SessionApiInterface.delete method's vapiStdErrors_.Unauthenticated error - %s",
			vapiBindings_.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(vapiData_.ErrorDefinition))
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(vapiStdErrors_.ServiceUnavailableBindingType())
	if errError2 != nil {
		vapiLog_.Errorf("Error in ConvertToDataDefinition for SessionApiInterface.delete method's vapiStdErrors_.ServiceUnavailable error - %s",
			vapiBindings_.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(vapiData_.ErrorDefinition))

	methodDefinition := vapiCore_.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}
func (sIface *SessionApiInterface) getMethodDefinition(outputType vapiBindings_.BindingType) *vapiCore_.MethodDefinition {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier(sIface.interfaceName)
	typeConverter := vapiBindings_.NewTypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(sessionGetInputType())
	if inputError != nil {
		vapiLog_.Errorf("Error in ConvertToDataDefinition for SessionApiInterface.get method's input - %s",
			vapiBindings_.VAPIerrorsToError(inputError).Error())
		return nil
	}

	output, outputError := typeConverter.ConvertToDataDefinition(outputType)
	if outputError != nil {
		vapiLog_.Errorf("Error in ConvertToDataDefinition for SessionApiInterface.get method's output - %s",
			vapiBindings_.VAPIerrorsToError(outputError).Error())
		return nil
	}

	methodIdentifier := vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get")
	errorDefinitions := make([]vapiData_.ErrorDefinition, 0)
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(vapiStdErrors_.UnauthenticatedBindingType())
	if errError1 != nil {
		vapiLog_.Errorf("Error in ConvertToDataDefinition for SessionApiInterface.get method's vapiStdErrors_.Unauthenticated error - %s",
			vapiBindings_.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(vapiData_.ErrorDefinition))
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(vapiStdErrors_.ServiceUnavailableBindingType())
	if errError2 != nil {
		vapiLog_.Errorf("Error in ConvertToDataDefinition for SessionApiInterface.get method's vapiStdErrors_.ServiceUnavailable error - %s",
			vapiBindings_.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(vapiData_.ErrorDefinition))

	methodDefinition := vapiCore_.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func createSessionErrorResult(errs []error, msg string) vapiCore_.MethodResult {
	vapiLog_.Error(msg)
	errValue := vapiBindings_.CreateErrorValueFromMessages(vapiBindings_.INTERNAL_SERVER_ERROR_DEF, errs)
	return vapiCore_.NewMethodResult(nil, errValue)
}

func createSessionErrorInput(errs []error, msg string) vapiCore_.MethodResult {
	vapiLog_.Error(msg)
	errValue := vapiBindings_.CreateErrorValueFromMessages(vapiBindings_.INVALID_ARGUMENT_ERROR_DEF, errs)
	return vapiCore_.NewMethodResult(nil, errValue)
}

func (sIface *SessionApiInterface) Identifier() vapiCore_.InterfaceIdentifier {
	return vapiCore_.NewInterfaceIdentifier(sIface.interfaceName)
}

func (sIface *SessionApiInterface) Definition() vapiCore_.InterfaceDefinition {
	return sIface.interfaceDefinition
}

func (sIface *SessionApiInterface) MethodDefinition(mi vapiCore_.MethodIdentifier) *vapiCore_.MethodDefinition {
	if val, ok := sIface.methodMetadataMap[mi.Name()]; ok {
		return val.MethodDefinition()
	}
	return nil
}

func (sIface *SessionApiInterface) Invoke(ctx *vapiCore_.ExecutionContext, methodId vapiCore_.MethodIdentifier, inputDataValue vapiData_.DataValue) vapiCore_.MethodResult {
	switch methodId.Name() {
	case "create":
		return sIface.create(ctx, methodId, inputDataValue)
	case "delete":
		return sIface.delete(ctx, methodId, inputDataValue)
	case "get":
		return sIface.get(ctx, methodId, inputDataValue)
	}
	return createSessionErrorResult([]error{}, "Invalid method call")
}

func (sIface *SessionApiInterface) RegisterRoutesHandlers(router vapiProtocolServerRest_.Router, provider vapiCore_.APIProvider, appCtxBuilders []vapiProtocolServerRestContextbuilder_.ApplicationContextBuilder, secCtxBuilders []vapiProtocolServerRestContextbuilder_.SecurityContextBuilder, opts ...vapiProtocolServerRest_.RequestHandlerOption) {
	opts = append(opts, vapiProtocolServerRest_.WithApplicationContextBuilders(appCtxBuilders))
	opts = append(opts, vapiProtocolServerRest_.WithSecurityContextBuilders(secCtxBuilders))

	router.AddRoute().
		Methods("POST").
		Path("/session").
		Handler(vapiProtocolServerRest_.NewRequestHandler(sIface.methodMetadataMap["create"], provider, opts...))

	router.AddRoute().
		Methods("DELETE").
		Path("/session").
		Handler(vapiProtocolServerRest_.NewRequestHandler(sIface.methodMetadataMap["delete"], provider, opts...))

	router.AddRoute().
		Methods("GET").
		Path("/session").
		Handler(vapiProtocolServerRest_.NewRequestHandler(sIface.methodMetadataMap["get"], provider, opts...))

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
