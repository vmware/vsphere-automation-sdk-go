// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Skeleton file for service: Greeter
// Implementation vapiCore_.ApiInterface which adapts API service implementation -
// Functions that implement the generated GreeterProvider interface

package greeter

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

type GreeterApiInterface struct {
	interfaceName       string
	interfaceDefinition vapiCore_.InterfaceDefinition
	methodIdentifiers   map[string]vapiCore_.MethodIdentifier
	methodMetadataMap   map[string]vapiProtocol_.OperationMetadata
	impl                GreeterProvider
}

func NewGreeterApiInterface(impl GreeterProvider) *GreeterApiInterface {
	interfaceName := "vmodl.examples.greeter.greeter"
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier(interfaceName)
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"greet": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "greet"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	gIface := GreeterApiInterface{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, impl: impl}
	gIface.methodMetadataMap = map[string]vapiProtocol_.OperationMetadata{}
	gIface.methodMetadataMap["greet"] = gIface.GreetMetadata(GreeterGreetOutputType())
	return &gIface
}

func (gIface *GreeterApiInterface) greet(ctx *vapiCore_.ExecutionContext, methodId vapiCore_.MethodIdentifier, inputDataValue vapiData_.DataValue) vapiCore_.MethodResult {
	acceptOptions := &vapiCore_.BindingAcceptOptions{
		MethodName: methodId.Name(),
		IsStream:   false,
		IsTaskOnly: false,
	}
	if ok, mr := vapiCore_.ValidateServerBindings(ctx, acceptOptions); !ok {
		return mr
	}

	typeConverter := vapiBindings_.NewTypeConverter()
	opMetadata := gIface.methodMetadataMap["greet"]

	result, err := gIface.impl.Greet(ctx)

	if err != nil {
		errType, exist := opMetadata.ErrorBindingMap()[err.Error()]
		if !exist {
			return createGreeterErrorResult([]error{err}, "Invalid error returned by server")
		}
		errorValue, errorInError := typeConverter.ConvertToVapi(err, errType)
		if errorInError != nil {
			return createGreeterErrorResult(errorInError, "Invalid error")
		}
		return vapiCore_.NewMethodResult(nil, errorValue.(*vapiData_.ErrorValue))
	}

	outputDataValue, errorInOutput := typeConverter.ConvertToVapi(result, opMetadata.OutputType())
	if errorInOutput != nil {
		return createGreeterErrorResult(errorInOutput, "Invalid output")
	}
	return vapiCore_.NewMethodResult(outputDataValue, nil)
}

func (gIface *GreeterApiInterface) GreetMetadata(outputType vapiBindings_.BindingType) vapiProtocol_.OperationMetadata {
	methodDefinition := gIface.GreetMethodDefinition(outputType)

	errorBindingMap := map[string]vapiBindings_.BindingType{}
	errorBindingMap[vapiStdErrors_.Unauthenticated{}.Error()] = vapiStdErrors_.UnauthenticatedBindingType()

	return vapiProtocol_.NewOperationMetadata(
		methodDefinition,
		greeterGreetInputType(),
		outputType,
		errorBindingMap,
		greeterGreetRestMetadata())
}

func (gIface *GreeterApiInterface) GreetMethodDefinition(outputType vapiBindings_.BindingType) *vapiCore_.MethodDefinition {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier(gIface.interfaceName)
	typeConverter := vapiBindings_.NewTypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(greeterGreetInputType())
	if inputError != nil {
		vapiLog_.Errorf("Error in ConvertToDataDefinition for GreeterApiInterface.greet method's input - %s",
			vapiBindings_.VAPIerrorsToError(inputError).Error())
		return nil
	}

	output, outputError := typeConverter.ConvertToDataDefinition(outputType)
	if outputError != nil {
		vapiLog_.Errorf("Error in ConvertToDataDefinition for GreeterApiInterface.greet method's output - %s",
			vapiBindings_.VAPIerrorsToError(outputError).Error())
		return nil
	}

	methodIdentifier := vapiCore_.NewMethodIdentifier(interfaceIdentifier, "greet")
	errorDefinitions := make([]vapiData_.ErrorDefinition, 0)
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(vapiStdErrors_.UnauthenticatedBindingType())
	if errError1 != nil {
		vapiLog_.Errorf("Error in ConvertToDataDefinition for GreeterApiInterface.greet method's vapiStdErrors_.Unauthenticated error - %s",
			vapiBindings_.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(vapiData_.ErrorDefinition))

	methodDefinition := vapiCore_.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func createGreeterErrorResult(errs []error, msg string) vapiCore_.MethodResult {
	vapiLog_.Error(msg)
	errValue := vapiBindings_.CreateErrorValueFromMessages(vapiBindings_.INTERNAL_SERVER_ERROR_DEF, errs)
	return vapiCore_.NewMethodResult(nil, errValue)
}

func createGreeterErrorInput(errs []error, msg string) vapiCore_.MethodResult {
	vapiLog_.Error(msg)
	errValue := vapiBindings_.CreateErrorValueFromMessages(vapiBindings_.INVALID_ARGUMENT_ERROR_DEF, errs)
	return vapiCore_.NewMethodResult(nil, errValue)
}

func (gIface *GreeterApiInterface) Identifier() vapiCore_.InterfaceIdentifier {
	return vapiCore_.NewInterfaceIdentifier(gIface.interfaceName)
}

func (gIface *GreeterApiInterface) Definition() vapiCore_.InterfaceDefinition {
	return gIface.interfaceDefinition
}

func (gIface *GreeterApiInterface) MethodDefinition(mi vapiCore_.MethodIdentifier) *vapiCore_.MethodDefinition {
	if val, ok := gIface.methodMetadataMap[mi.Name()]; ok {
		return val.MethodDefinition()
	}
	return nil
}

func (gIface *GreeterApiInterface) Invoke(ctx *vapiCore_.ExecutionContext, methodId vapiCore_.MethodIdentifier, inputDataValue vapiData_.DataValue) vapiCore_.MethodResult {
	switch methodId.Name() {
	case "greet":
		return gIface.greet(ctx, methodId, inputDataValue)
	}
	return createGreeterErrorResult([]error{}, "Invalid method call")
}

func (gIface *GreeterApiInterface) RegisterRoutesHandlers(router vapiProtocolServerRest_.Router, provider vapiCore_.APIProvider, appCtxBuilders []vapiProtocolServerRestContextbuilder_.ApplicationContextBuilder, secCtxBuilders []vapiProtocolServerRestContextbuilder_.SecurityContextBuilder, opts ...vapiProtocolServerRest_.RequestHandlerOption) {
	opts = append(opts, vapiProtocolServerRest_.WithApplicationContextBuilders(appCtxBuilders))
	opts = append(opts, vapiProtocolServerRest_.WithSecurityContextBuilders(secCtxBuilders))

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
